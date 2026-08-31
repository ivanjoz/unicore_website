# Unicore — backend

Backend en Go para el formulario de contacto del sitio. Corre como una sola función Lambda
(`provided.al2023`, arm64) detrás de una Function URL, con DynamoDB como única base de datos.

El diseño está tomado del backend de genix (`backend/security/contact.go`): mismas validaciones,
mismo orden de operaciones, mismo namespace de acciones para el limitador y el mismo empaquetado
de la IP del cliente. Lo único que cambia es *cómo* se hace atómico el conteo, y esa diferencia se
explica abajo.

## Rutas

| Método | Ruta                 | Descripción |
| ------ | -------------------- | ----------- |
| `POST` | `/p-contact-message` | Guarda un mensaje y lo notifica al buzón configurado. |

El prefijo `p-` marca una ruta pública: sin token y sin cuenta, igual que en genix.

Cuerpo:

```json
{ "Name": "Ada Lovelace", "Email": "ada@example.com", "Company": "opcional", "Message": "…" }
```

Respuestas:

| Código | Cuerpo                                  | Cuándo |
| ------ | --------------------------------------- | ------ |
| `200`  | `{"Received":true,"Notified":true}`     | Guardado y notificado. |
| `200`  | `{"Received":true,"Notified":false}`    | Guardado; el correo falló. El mensaje **no** se pierde. |
| `400`  | `{"error":"…"}`                         | Cuerpo inválido u origen indeterminable. |
| `429`  | `{"error":"…"}` + `Retry-After`         | Límite por IP alcanzado. |
| `503`  | `{"error":"…"}`                         | Sin `CONTACT_EMAIL`, o el limitador no responde. |

## El limitador de tasa

El problema que resuelve es una carrera de lectura y escritura: varias Lambdas sirviendo al mismo
cliente leen el mismo contador, todas concluyen que están por debajo del límite y todas escriben.
Peticiones en paralelo desde un mismo origen no son la excepción, son *el patrón de abuso*, así
que un límite que se cae justo ahí no vale nada.

genix lo resuelve con un lock distribuido contra un daemon en Rust: se toma la llave por
`(acción, IP)`, se cuentan las filas de la ventana, se compara y se inserta. Aquí ese daemon no
existe, y DynamoDB ofrece la primitiva que lo hace innecesario. Un `UpdateItem` aplica su
`ConditionExpression` y su `ADD` como una sola operación atómica sobre un solo ítem:

```
Key:        pk = "RL#<acción>#<identificador>",  sk = "W#<inicioDeVentana>"
Update:     SET ttl = :ttl ADD hits :one
Condition:  attribute_not_exists(hits) OR hits < :limit
```

De N llamadas simultáneas pasan exactamente las primeras `limit`; el resto recibe
`ConditionalCheckFailedException`. Un viaje de ida y vuelta, sin llave que tomar, sin llave que
soltar, y sin nada retenido si la Lambda muere a mitad de la petición.

Lo que se conserva del diseño de genix:

- **Namespace de acciones** (`core.LimitAction`). La acción encabeza la clave, así dos
  funcionalidades nunca comparten presupuesto aunque cuenten sobre el mismo número.
- **`ClientIPKey`**. IPv4 va tal cual; IPv6 se reduce al prefijo /63, porque a un cliente
  doméstico se le entrega un /64 entero y limitar por dirección sería gratis de evadir.
- **Fallo cerrado.** Si DynamoDB no responde, la petición se rechaza con 503. Sin límite
  aplicable, un endpoint sin autenticar es un relé abierto hacia el buzón de alguien.
- **Escribir primero, enviar después.** La fila *es* el mensaje; perderlo por un fallo de SMTP
  sería perder lo que el visitante escribió. Un envío fallido marca estado 2 y consume igual el
  presupuesto, que es lo que impide que un servidor de correo roto convierta esto en un endpoint
  sin medir.

Lo que cambia, y su coste: la ventana es **fija**, no deslizante. Cada identificador reinicia su
contador en el mismo borde de reloj, de modo que un llamante puede gastar su presupuesto al final
de una ventana y otra vez al principio de la siguiente — 2×límite alcanzable a caballo del borde.
Se acepta porque la alternativa que mantiene una ventana deslizante real es guardar un ítem por
petición y sumarlos, que es una lectura seguida de una escritura: exactamente la carrera que este
diseño existe para evitar. Léase el límite como «como mucho 2×límite por ventana en el peor caso».

## Esquema

Una sola tabla, `pk` + `sk`, con TTL habilitado.

| Entidad  | `pk`                        | `sk`                          |
| -------- | --------------------------- | ----------------------------- |
| Mensaje  | `MSG#<códigoDeSemanaISO>`   | `<ip>#<creado>#<secuencia>`   |
| Ventana  | `RL#<acción>#<id>`          | `W#<inicioDeVentana>`         |

Los mensajes se particionan por semana ISO porque no pertenecen a ninguna cuenta —quien escribe no
tiene una—, lo que evita una partición única que crece sin límite y permite purgar dejando caer
semanas enteras. La IP encabeza la clave de ordenación, así que «qué más envió esta dirección» es
un rango sobre una partición y no un recorrido de la semana. La cola de la clave es el contador que
devolvió el limitador: como su incremento es atómico, dos escritores concurrentes nunca reciben el
mismo valor, y dos mensajes enviados dentro del mismo segundo no pueden pisarse.

Las ventanas del limitador llevan TTL y se borran solas.

## Desarrollo

```bash
podman run -d --rm -p 8000:8000 docker.io/amazon/dynamodb-local

export DYNAMO_ENDPOINT=http://localhost:8000
export AWS_REGION=us-east-1 AWS_ACCESS_KEY_ID=fake AWS_SECRET_ACCESS_KEY=fake

go test ./...          # las pruebas de integración se omiten sin DYNAMO_ENDPOINT
go run .               # servidor local en :3333
```

```bash
curl -s localhost:3333/p-contact-message -H 'Content-Type: application/json' -d '{
  "Name":"Ada Lovelace","Email":"ada@example.com",
  "Message":"Quisiera conversar sobre una iniciativa de código abierto."}'
```

Las pruebas del limitador corren contra un DynamoDB real, no contra un doble: lo que se está
comprobando es la atomicidad de DynamoDB, y un simulador solo se comprobaría a sí mismo.
`TestCheckAndCountIsAtomicUnderConcurrency` lanza 40 llamadas simultáneas contra un límite de 5 y
exige que pasen exactamente 5, cada una con un contador distinto.

## Despliegue

```bash
cp config.example.toml config.toml   # completar
./deploy.sh 3                        # crear/actualizar toda la infraestructura
```

`deploy.sh` es un wrapper mínimo; toda la lógica vive en `cloud/`, un **módulo Go aparte** del
backend. Lo es a propósito: los SDK de CloudFormation, S3 y Lambda solo hacen falta para
desplegar, y meterlos en el `go.mod` del backend los metería en el binario que sube a Lambda.

El deployer lleva `cloudformation.yml` embebido con `//go:embed` y habla directo con la API de
AWS. No hay dependencia de la CLI de `aws`, ni de CDK, ni de Node: el despliegue completo es este
binario más su plantilla, que por venir dentro no pueden desincronizarse.

| Acción | Qué hace | Cuándo |
| ------ | -------- | ------ |
| `./deploy.sh 1` | Compila arm64 y sube el `.zip` con `UpdateFunctionCode`. | A diario. Segundos: no toca CloudFormation. |
| `./deploy.sh 2` | Reescribe las variables de entorno de la Lambda. | Al cambiar `config.toml` sin tocar código. |
| `./deploy.sh 3` | Compila → S3 → CloudFormation → reinyecta variables. | Al cambiar la infraestructura. Minutos. |

Sin argumento pregunta la acción. `UNICORE_CONFIG_FILE=/ruta/config.prod.toml ./deploy.sh 3`
selecciona otro entorno sin editar el archivo de siempre.

Mientras el stack se despliega, el deployer sondea `DescribeStackEvents` e imprime cada evento
nuevo con su recurso y su razón. Si falla, repite al final la **causa raíz**: un fallo real
arrastra a los demás recursos con «Resource creation cancelled», y esa cascada tapa la única línea
que dice qué hay que arreglar. Al terminar escribe la Function URL en `aws.lambda_url` de
`config.toml`, por reemplazo de texto y no re-serializando el TOML, para no perder los comentarios.

El bucket de `aws.deployment_bucket` debe existir antes del primer despliegue: el stack lee de ahí
el `.zip`, así que crearlo desde el propio stack sería circular.

### El bloque de variables de entorno

`UpdateFunctionConfiguration` **reemplaza** el entorno de la Lambda, no lo fusiona, y
CloudFormation hace lo mismo con su bloque `Environment`. De ahí dos consecuencias que el código
respeta y dos pruebas que las vigilan:

- La acción 3 termina llamando a la 2. Sin eso, cada despliegue de infraestructura borraría las
  variables que solo conoce el deployer.
- `lambdaEnvironmentVariables` es el superconjunto de lo que declara la plantilla, no solo los
  secretos. `TestEveryTemplateEnvironmentVariableIsReinjected` falla si una variable de la
  plantilla no está ahí, porque desaparecería de la Lambda en el siguiente despliegue.

`mailer.smtp_password` es el único secreto y por eso no aparece en la plantilla: allí sería legible
en la consola de CloudFormation para cualquiera con acceso de lectura. Llega solo por la API de
Lambda.

`TestDeployerAndTemplateAgreeOnEveryParameter` compara los parámetros que envía el deployer con los
que declara el YAML embebido. CloudFormation rechaza uno que no declara —error ruidoso— pero acepta
en silencio no recibir uno que sí declara, usando su `Default`; esa segunda mitad es la que cuesta
descubrir.

### Correo

Para que salga, el remitente de `mailer.ses_from_email` debe estar verificado en SES en la misma
región, y la cuenta fuera del *sandbox* si el destinatario no está verificado también. La política
IAM se acota a esa identidad y a su dominio —un remitente puede estar verificado de cualquiera de
las dos formas— en vez de a `*`. Mientras tanto el formulario funciona: los mensajes se guardan con
estado 2.

## Variables de entorno

Las escribe el deployer desde `config.toml`; esta tabla es lo que lee el backend en ejecución.

| Variable                      | Por defecto        | Descripción |
| ----------------------------- | ------------------ | ----------- |
| `APP_NAME`                    | `unicore`          | Prefijo de recursos y firma del correo. |
| `APP_URL`                     | `https://un.pe`    | Solo aparece en el correo de notificación. |
| `DYNAMO_TABLE`                | `<APP_NAME>-db`    | Tabla única. |
| `DYNAMO_ENDPOINT`             | —                  | DynamoDB local, para pruebas. |
| `ALLOWED_ORIGINS`             | `*`                | Lista blanca de CORS, separada por comas. |
| `CONTACT_EMAIL`               | —                  | Buzón destino. Vacío deshabilita el endpoint. |
| `CONTACT_MAX_MESSAGES_PER_IP` | `3`                | Mensajes por ventana. |
| `CONTACT_WINDOW_MINUTES`      | `10`               | Duración de la ventana. |
| `CONTACT_TTL_DAYS`            | `0`                | Caducidad de los mensajes. 0 = para siempre. |
| `SES_FROM_EMAIL`              | —                  | Remitente SES. Preferido sobre SMTP. |
| `SMTP_HOST` … `SMTP_FROM`     | —                  | Alternativa a SES. |
| `CLIENT_IP_HEADER`            | —                  | Ver abajo. |
| `LOGS_FULL`                   | —                  | `1` imprime todas las líneas de log. |

### `CLIENT_IP_HEADER`

Por defecto la IP se toma de `RequestContext.HTTP.SourceIP`, que escribe AWS desde el par TCP y no
se puede falsificar. Las cabeceras sí: a `X-Forwarded-For` se le *añade*, de modo que un cliente
que manda la suya queda primero en la lista, y cualquier límite basado en ese valor sería evadible
con un flag de curl. Por eso no se lee ninguna.

La excepción es poner un CDN delante: entonces `SourceIP` es la dirección del CDN, todos los
visitantes comparten una clave y los tres primeros agotan el límite de todo el mundo. En ese caso
`CLIENT_IP_HEADER=cloudfront-viewer-address` lee la cabecera del CDN. **Solo es seguro cuando la
Function URL no se puede alcanzar directamente**, porque desde ese momento la cabecera es tan
fiable como quien pueda llegar al origen.
