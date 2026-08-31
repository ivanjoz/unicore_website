package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

// El deployer del backend. Es un módulo Go aparte del backend (cloud/go.mod) a propósito: los SDK
// de CloudFormation, S3 y Lambda solo hacen falta para desplegar, y arrastrarlos al go.mod del
// backend los metería en el binario que sube a Lambda, donde no se usan nunca.
//
// Habla directamente con la API de AWS. No hay dependencia de la CLI de aws, ni de CDK, ni de
// Node: el despliegue completo es este binario más la plantilla que lleva embebida.

// DeployParams refleja config.toml por secciones.
type DeployParams struct {
	// AppName prefija todo nombre físico: stack, Lambda, tabla y rol.
	AppName string `toml:"app_name"`
	// S3CompiledKey es la ruta del .zip dentro del bucket de despliegue. No lleva tag de archivo:
	// lo fija main(), no el usuario.
	S3CompiledKey string

	AWS struct {
		Profile          string `toml:"profile"`
		Region           string `toml:"region"`
		DeploymentBucket string `toml:"deployment_bucket"`
		// LambdaURL lo escribe el propio deployer tras desplegar la infraestructura. Está en el
		// archivo para que el frontend pueda leerlo sin consultar a AWS.
		LambdaURL string `toml:"lambda_url"`
	} `toml:"aws"`

	Backend struct {
		AppURL              string `toml:"app_url"`
		AllowedOrigins      string `toml:"allowed_origins"`
		ClientIPHeader      string `toml:"client_ip_header"`
		ReservedConcurrency int64  `toml:"reserved_concurrency"`
	} `toml:"backend"`

	Contact struct {
		Email            string `toml:"email"`
		MaxMessagesPerIP int64  `toml:"max_messages_per_ip"`
		WindowMinutes    int64  `toml:"window_minutes"`
		TTLDays          int64  `toml:"ttl_days"`
	} `toml:"contact"`

	Mailer struct {
		SesFromEmail string `toml:"ses_from_email"`
		SMTPHost     string `toml:"smtp_host"`
		SMTPPort     int64  `toml:"smtp_port"`
		SMTPUser     string `toml:"smtp_user"`
		// SMTPPassword es el único secreto de este archivo, y la razón por la que las variables de
		// entorno se reinyectan con la API de Lambda en vez de declararse en la plantilla: allí
		// serían legibles en la consola de CloudFormation para cualquiera con acceso de lectura.
		SMTPPassword string `toml:"smtp_password"`
		SMTPFrom     string `toml:"smtp_from"`
	} `toml:"mailer"`
}

// La clave es fija y no versionada por commit. CloudFormation solo redespliega el código de una
// Lambda cuando la clave cambia, así que la acción 3 no publica código nuevo por sí sola — eso
// lo hace UpdateFunctionCode al final, que sube el .zip directamente y no depende de S3.
const s3CompiledKey = "unicore-artifacts/backend-lambda.zip"

const compiledRelativePath = "/cloud/backend-compiled"

// GetBaseWD es la raíz del backend: este programa siempre corre desde cloud/.
func GetBaseWD() string {
	workingDirectory, _ := os.Getwd()
	parts := strings.Split(workingDirectory, "/")
	return strings.Join(parts[:len(parts)-1], "/")
}

// GetConfigPath deja que UNICORE_CONFIG_FILE seleccione el entorno, para no tener que editar el
// archivo de siempre cuando se despliega a otra cuenta.
func GetConfigPath() string {
	if configured := strings.TrimSpace(os.Getenv("UNICORE_CONFIG_FILE")); len(configured) > 0 {
		return configured
	}
	return GetBaseWD() + "/config.toml"
}

func main() {
	action := readAction()

	configPath := GetConfigPath()
	fmt.Println("Leyendo configuración desde:", configPath)
	configBytes, err := ReadFile(configPath)
	if err != nil {
		exitWithError("No se pudo leer " + configPath + ": " + err.Error())
	}

	params := DeployParams{S3CompiledKey: s3CompiledKey}
	if err := toml.Unmarshal(configBytes, &params); err != nil {
		exitWithError("El archivo " + configPath + " posee un formato erróneo: " + err.Error())
	}
	validateParams(&params)

	switch action {
	case "1":
		CompileBackend(params)
		DeployLambdaCode(params)
	case "2":
		UpdateEnvironmentVariables(params)
	case "3":
		DeployInfraestructure(params)
	default:
		fmt.Println("No se reconoció la opción seleccionada.")
		os.Exit(1)
	}
}

// readAction acepta la acción como argumento suelto ("./deploy.sh 3"), como accion=3, o la
// pregunta. Lo primero es lo que usa CI; lo último, una persona.
func readAction() string {
	validActions := map[string]bool{"1": true, "2": true, "3": true}

	for _, argument := range os.Args[1:] {
		if validActions[argument] {
			return argument
		}
		if strings.HasPrefix(argument, "accion=") {
			return strings.TrimPrefix(argument, "accion=")
		}
	}

	fmt.Println("Selecciona acción: [1] Publicar Código  [2] Actualizar Variables  [3] Desplegar Infraestructura")
	answer := ""
	if _, err := fmt.Scanln(&answer); err != nil {
		exitWithError("No se leyó ninguna acción.")
	}
	return strings.TrimSpace(answer)
}

// Se comprueba antes de tocar AWS y se nombran todas las que faltan de una vez: descubrirlas de
// una en una, cada una tras un despliegue fallido, es la forma lenta de leer esta lista.
func validateParams(params *DeployParams) {
	required := map[string]string{
		"app_name":              params.AppName,
		"aws.profile":           params.AWS.Profile,
		"aws.region":            params.AWS.Region,
		"aws.deployment_bucket": params.AWS.DeploymentBucket,
	}

	missing := []string{}
	for name, value := range required {
		if len(strings.TrimSpace(value)) == 0 {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		exitWithError("Faltan parámetros obligatorios en la configuración: " + strings.Join(missing, ", "))
	}

	// Un buzón vacío no impide desplegar —el backend responde 503 y lo dice—, pero desplegar un
	// formulario que no entrega a nadie casi nunca es lo que se pretendía.
	if len(strings.TrimSpace(params.Contact.Email)) == 0 {
		fmt.Println("AVISO: contact.email está vacío; el formulario responderá 503 hasta que se configure.")
	}
	if len(strings.TrimSpace(params.Mailer.SesFromEmail)) == 0 &&
		len(strings.TrimSpace(params.Mailer.SMTPHost)) == 0 {
		fmt.Println("AVISO: no hay remitente configurado; los mensajes se guardarán con estado 2 (no notificado).")
	}
}

// Termina con un mensaje legible en vez de un panic: un stack trace de Go no aporta nada cuando el
// error viene de AWS o de un archivo mal escrito, y además tapa la causa real.
func exitWithError(message string) {
	fmt.Println("\n" + message)
	os.Exit(1)
}

// DeployInfraestructure es el despliegue completo.
func DeployInfraestructure(params DeployParams) {
	CompileBackend(params)
	UploadCompiledToS3(params)

	fmt.Println("\nDesplegando infraestructura con CloudFormation...")
	DeployCloudFormation(params)

	// La plantilla declara el bloque Environment completo y CloudFormation lo reemplaza en vez de
	// fusionarlo, así que cada despliegue de infraestructura borra las variables que solo conoce
	// el deployer. Reinyectarlas aquí es lo que mantiene SMTP_PASSWORD fuera de la plantilla.
	UpdateEnvironmentVariables(params)
}
