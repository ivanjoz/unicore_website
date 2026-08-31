package main

import (
	"archive/zip"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func lambdaFunctionName(params DeployParams) string { return params.AppName + "-backend" }

func compiledBinaryPath() string { return GetBaseWD() + compiledRelativePath }

func compiledZipPath() string { return compiledBinaryPath() + ".zip" }

// CompileBackend produce el binario arm64 y lo empaqueta.
func CompileBackend(params DeployParams) {
	binaryPath := compiledBinaryPath()

	//   - CGO_ENABLED=0  porque provided.al2023 no trae las bibliotecas del sistema contra las que
	//     enlazaría, y un binario dinámico ahí no arranca.
	//   - lambda.norpc   descarta la vía de invocación local heredada de aws-lambda-go, y con ella
	//     net/rpc, que llama a reflect.Value.MethodByName con una cadena en tiempo de ejecución.
	//     Eso enciende una bandera global del enlazador que retiene todos los métodos exportados
	//     de todo tipo alcanzable, y desactiva la eliminación de código muerto del binario entero.
	//   - -trimpath      quita las rutas absolutas de la máquina que compila, que no aportan nada
	//     y hacen que el binario dependa de dónde se construyó.
	//   - -ldflags -s -w descarta símbolos y DWARF: no hay depurador en Lambda, y un binario más
	//     pequeño arranca en frío antes.
	command := fmt.Sprintf(
		`CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -trimpath -ldflags '-s -w' -o %v .`,
		binaryPath)

	fmt.Println("Compilando con::", command)
	buildCommand := exec.Command("bash", "-c", command)
	buildCommand.Dir = GetBaseWD()
	buildCommand.Stdout = os.Stdout
	buildCommand.Stderr = os.Stderr

	if err := buildCommand.Run(); err != nil {
		exitWithError("Error al generar el compilado: " + err.Error())
	}

	if err := CompressExecutable(binaryPath, compiledZipPath()); err != nil {
		exitWithError("Error al comprimir el compilado en .zip: " + err.Error())
	}

	info, err := os.Stat(compiledZipPath())
	if err == nil {
		fmt.Printf("Compilado listo: %v (%v kb)\n", compiledZipPath(), info.Size()/1024)
	}
}

// CompressExecutable escribe el .zip que espera Lambda: una sola entrada llamada "bootstrap".
//
// Se construye con archive/zip y no invocando a `zip` porque el permiso de ejecución es parte del
// contrato: sin el bit +x en los atributos externos de la entrada, el runtime no puede arrancar el
// binario y falla con un error que no menciona los permisos por ningún lado.
func CompressExecutable(binaryPath string, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	binaryBytes, err := os.ReadFile(binaryPath)
	if err != nil {
		return err
	}

	entry, err := zipWriter.CreateHeader(&zip.FileHeader{
		Name:           "bootstrap",
		Method:         zip.Deflate,
		CreatorVersion: 3 << 8,     // Unix
		ExternalAttrs:  0755 << 16, // -rwxr-xr-x
	})
	if err != nil {
		return err
	}

	_, err = entry.Write(binaryBytes)
	return err
}

// UploadCompiledToS3 deja el .zip donde la plantilla lo declara. Solo hace falta para crear el
// stack: la propiedad Code de una AWS::Lambda::Function apunta a S3 y CloudFormation lee de ahí.
func UploadCompiledToS3(params DeployParams) {
	awsConfig := MakeAwsConfig(params)
	client := s3.NewFromConfig(awsConfig)

	fmt.Printf("Enviando compilado a s3://%v/%v\n", params.AWS.DeploymentBucket, params.S3CompiledKey)
	if err := SendFileToS3(client, params.AWS.DeploymentBucket, params.S3CompiledKey, compiledZipPath()); err != nil {
		exitWithError("Error al enviar el compilado a S3: " + err.Error())
	}
	fmt.Println("Compilado subido.")
}

// DeployLambdaCode publica el código sin tocar CloudFormation, subiendo el .zip directamente.
//
// Es la acción 1, y es la que se usa a diario: un despliegue de infraestructura tarda minutos
// porque CloudFormation reconcilia todo el stack, mientras que esto son segundos y es lo único
// que hace falta cuando lo que cambió es el código de Go.
func DeployLambdaCode(params DeployParams) {
	zipBytes, err := os.ReadFile(compiledZipPath())
	if err != nil {
		exitWithError("Error al leer el compilado .zip: " + err.Error())
	}

	functionName := lambdaFunctionName(params)
	fmt.Printf("Publicando %v kb en la Lambda %v...\n", len(zipBytes)/1024, functionName)

	client := lambda.NewFromConfig(MakeAwsConfig(params))
	_, err = client.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
		FunctionName: &functionName,
		ZipFile:      zipBytes,
	})
	if err != nil {
		exitWithError("Error al actualizar el código de la Lambda: " + err.Error())
	}

	fmt.Println("Código de la Lambda actualizado.")
}

// lambdaEnvironmentVariables es el entorno completo de la Lambda desplegada.
//
// Es el SUPERCONJUNTO de lo que declara la plantilla, no solo los secretos, y eso es obligatorio:
// UpdateFunctionConfiguration reemplaza el bloque entero en vez de fusionarlo, así que una
// variable que la plantilla declare y este mapa omita desaparece de la Lambda en cuanto la acción
// 3 reinyecta el entorno. TestEveryTemplateEnvironmentVariableIsReinjected vigila esa relación.
func lambdaEnvironmentVariables(params DeployParams) map[string]string {
	variables := map[string]string{
		"APP_NAME":                    params.AppName,
		"APP_URL":                     params.Backend.AppURL,
		"DYNAMO_TABLE":                params.AppName + "-db",
		"ALLOWED_ORIGINS":             params.Backend.AllowedOrigins,
		"CLIENT_IP_HEADER":            params.Backend.ClientIPHeader,
		"CONTACT_EMAIL":               params.Contact.Email,
		"CONTACT_MAX_MESSAGES_PER_IP": strconv.FormatInt(params.Contact.MaxMessagesPerIP, 10),
		"CONTACT_WINDOW_MINUTES":      strconv.FormatInt(params.Contact.WindowMinutes, 10),
		"CONTACT_TTL_DAYS":            strconv.FormatInt(params.Contact.TTLDays, 10),
		"SES_FROM_EMAIL":              params.Mailer.SesFromEmail,
	}

	// Solo cuando SMTP es de verdad la vía de envío. core.SendEmail prefiere SES en cuanto
	// SES_FROM_EMAIL tiene valor, así que con los dos configurados estas variables no se leerían
	// nunca — y una de ellas es una contraseña. Un secreto que no se usa no debe viajar: replicar
	// aquí la precedencia del runtime es lo que mantiene el bloque SMTP como un respaldo que se
	// activa borrando ses_from_email, en vez de como un secreto desplegado de más.
	if len(params.Mailer.SMTPHost) > 0 && len(params.Mailer.SesFromEmail) == 0 {
		variables["SMTP_HOST"] = params.Mailer.SMTPHost
		variables["SMTP_PORT"] = strconv.FormatInt(params.Mailer.SMTPPort, 10)
		variables["SMTP_USER"] = params.Mailer.SMTPUser
		variables["SMTP_PASSWORD"] = params.Mailer.SMTPPassword
		variables["SMTP_FROM"] = params.Mailer.SMTPFrom
	}

	return variables
}

// UpdateEnvironmentVariables reemplaza el entorno de la Lambda con el conjunto completo.
//
// UpdateFunctionConfiguration REEMPLAZA el bloque, no lo fusiona: toda variable que declara la
// plantilla debe repetirse aquí o desaparece de la Lambda desplegada. Por eso el mapa de abajo es
// el superconjunto y no solo los secretos, y por eso la acción 3 termina llamando a esta función.
func UpdateEnvironmentVariables(params DeployParams) {
	functionName := lambdaFunctionName(params)
	fmt.Println("Actualizando variables de entorno de", functionName, "...")

	variables := lambdaEnvironmentVariables(params)

	client := lambda.NewFromConfig(MakeAwsConfig(params))
	_, err := client.UpdateFunctionConfiguration(context.TODO(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: &functionName,
		Environment:  &lambdaTypes.Environment{Variables: variables},
	})
	if err != nil {
		exitWithError("Error al actualizar las variables de la Lambda: " + err.Error())
	}

	fmt.Println("Variables actualizadas.")
}
