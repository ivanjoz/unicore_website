package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cfnTypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

// La plantilla viaja dentro del binario: el despliegue es una llamada a la API de CloudFormation,
// sin Node, sin npx y sin bootstrap stack. También significa que el binario y la plantilla que
// aplica no pueden desincronizarse.
//
//go:embed cloudformation.yml
var cloudFormationTemplate string

const stackEventsPollInterval = 4 * time.Second

func stackNameOf(params DeployParams) string { return params.AppName + "-stack" }

// DeployCloudFormation crea el stack si no existe, o lo actualiza.
func DeployCloudFormation(params DeployParams) {
	stackName := stackNameOf(params)
	client := cloudformation.NewFromConfig(MakeAwsConfig(params))
	ctx := context.TODO()

	templateParameters := makeTemplateParameters(params)

	// El stack crea su propio rol de ejecución con un nombre fijo, y eso es exactamente lo que
	// CAPABILITY_NAMED_IAM reconoce. Es una confirmación explícita, no un permiso: sin ella
	// CloudFormation rechaza la plantilla en vez de crear identidades sin que nadie lo pidiera.
	capabilities := []cfnTypes.Capability{cfnTypes.CapabilityCapabilityNamedIam}

	currentStatus, stackExists := DescribeStackStatus(ctx, client, stackName)

	// Un create fallido deja el stack en ROLLBACK_COMPLETE: en ese estado solo se puede borrar, y
	// cualquier update es rechazado. Se avisa en vez de fallar con un error opaco de la API.
	if stackExists && currentStatus == cfnTypes.StackStatusRollbackComplete {
		exitWithError(fmt.Sprintf(
			"El stack %v está en ROLLBACK_COMPLETE (creación fallida previa). Bórralo antes de reintentar.",
			stackName))
	}

	// Solo se imprimen los eventos posteriores al arranque; los de despliegues anteriores sobran.
	deployStartedAt := time.Now().Add(-stackEventsPollInterval)

	var err error
	if stackExists {
		fmt.Printf("Actualizando stack %v (estado actual: %v)...\n", stackName, currentStatus)
		_, err = client.UpdateStack(ctx, &cloudformation.UpdateStackInput{
			StackName:    &stackName,
			TemplateBody: &cloudFormationTemplate,
			Parameters:   templateParameters,
			Capabilities: capabilities,
		})
		// CloudFormation responde con un error cuando la plantilla no cambió nada. No es un fallo:
		// la acción 3 se ejecuta a menudo solo para reinyectar variables.
		if err != nil && strings.Contains(err.Error(), "No updates are to be performed") {
			fmt.Println("Sin cambios en la infraestructura.")
			PrintStackOutputsAndSyncConfig(ctx, client, stackName)
			return
		}
	} else {
		fmt.Printf("Creando stack %v...\n", stackName)
		_, err = client.CreateStack(ctx, &cloudformation.CreateStackInput{
			StackName:    &stackName,
			TemplateBody: &cloudFormationTemplate,
			Parameters:   templateParameters,
			Capabilities: capabilities,
			// Sin esto un create fallido queda en ROLLBACK_COMPLETE y hay que borrarlo a mano
			// antes de poder reintentar.
			OnFailure: cfnTypes.OnFailureDelete,
		})
	}

	if err != nil {
		exitWithError("Error al enviar la plantilla a CloudFormation: " + err.Error())
	}

	finalStatus, rootFailureReason := WaitForStackAndPrintEvents(ctx, client, stackName, deployStartedAt)

	// Un fallo real arrastra a los demás recursos con "Resource creation cancelled", así que se
	// repite la primera causa: es la única línea que dice de verdad qué hay que arreglar.
	if !isSuccessfulStackStatus(finalStatus) {
		message := fmt.Sprintf("El despliegue del stack %v FALLÓ (estado %v).", stackName, finalStatus)
		if len(rootFailureReason) > 0 {
			message += "\nCausa raíz: " + rootFailureReason
		}
		exitWithError(message)
	}

	fmt.Printf("\nStack %v desplegado (%v)\n", stackName, finalStatus)
	PrintStackOutputsAndSyncConfig(ctx, client, stackName)
}

// makeTemplateParameters es el puente entre config.toml y la plantilla. Está separada de
// DeployCloudFormation para que una prueba pueda comprobar que cada nombre de aquí existe de
// verdad en el YAML embebido: CloudFormation rechaza un parámetro que no declara, y descubrirlo
// en mitad de un despliegue es la forma cara de leer un error tipográfico.
func makeTemplateParameters(params DeployParams) []cfnTypes.Parameter {
	return []cfnTypes.Parameter{
		makeStackParameter("NamePrefix", params.AppName),
		makeStackParameter("DeploymentBucket", params.AWS.DeploymentBucket),
		makeStackParameter("CompiledS3Key", params.S3CompiledKey),
		makeStackParameter("AppUrl", params.Backend.AppURL),
		makeStackParameter("AllowedOrigins", params.Backend.AllowedOrigins),
		makeStackParameter("ClientIpHeader", params.Backend.ClientIPHeader),
		makeStackParameter("ContactEmail", params.Contact.Email),
		makeStackParameter("ContactMaxMessagesPerIp", strconv.FormatInt(params.Contact.MaxMessagesPerIP, 10)),
		makeStackParameter("ContactWindowMinutes", strconv.FormatInt(params.Contact.WindowMinutes, 10)),
		makeStackParameter("ContactTtlDays", strconv.FormatInt(params.Contact.TTLDays, 10)),
		makeStackParameter("SesFromEmail", params.Mailer.SesFromEmail),
		makeStackParameter("ReservedConcurrency", strconv.FormatInt(params.Backend.ReservedConcurrency, 10)),
	}
}

// templateParameterNames son los parámetros que declara el YAML embebido: las claves indentadas
// con exactamente dos espacios dentro del bloque Parameters.
func templateParameterNames() []string {
	names := []string{}
	insideParameters := false

	for _, line := range strings.Split(cloudFormationTemplate, "\n") {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		// Una clave sin indentar cierra el bloque: Conditions, Resources, Outputs.
		if line[0] != ' ' && line[0] != '\t' {
			insideParameters = strings.HasPrefix(line, "Parameters:")
			continue
		}
		if !insideParameters {
			continue
		}
		if match := templateParameterKeyPattern.FindStringSubmatch(line); match != nil {
			names = append(names, match[1])
		}
	}
	return names
}

var templateParameterKeyPattern = regexp.MustCompile(`^  ([A-Za-z0-9]+):\s*$`)

func makeStackParameter(key, value string) cfnTypes.Parameter {
	return cfnTypes.Parameter{ParameterKey: aws.String(key), ParameterValue: aws.String(value)}
}

// DescribeStackStatus devuelve el estado actual del stack y si existe. DescribeStacks falla cuando
// no existe, que es la forma que da la API de responder esa pregunta.
func DescribeStackStatus(
	ctx context.Context, client *cloudformation.Client, stackName string,
) (cfnTypes.StackStatus, bool) {
	result, err := client.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{StackName: &stackName})
	if err != nil || len(result.Stacks) == 0 {
		return "", false
	}
	return result.Stacks[0].StackStatus, true
}

// WaitForStackAndPrintEvents sondea hasta un estado terminal, imprimiendo cada evento nuevo.
//
// Sin esto un despliegue es una espera ciega de varios minutos que acaba en "FAILED" sin decir por
// qué. Con esto, el evento nombra el recurso exacto y la razón, en el momento en que ocurre.
// Devuelve además la primera causa real de fallo, para poder repetirla al final.
func WaitForStackAndPrintEvents(
	ctx context.Context, client *cloudformation.Client, stackName string, since time.Time,
) (cfnTypes.StackStatus, string) {

	printedEventIDs := map[string]bool{}
	rootFailureReason := ""

	recordFailure := func(reason string) {
		if len(rootFailureReason) == 0 {
			rootFailureReason = reason
		}
	}

	for {
		recordFailure(PrintNewStackEvents(ctx, client, stackName, since, printedEventIDs))

		status, exists := DescribeStackStatus(ctx, client, stackName)
		// OnFailureDelete borra el stack tras un create fallido: deja de existir en pleno sondeo.
		if !exists {
			return "DELETED_AFTER_FAILURE", rootFailureReason
		}
		if isTerminalStackStatus(status) {
			// Última pasada: los eventos finales suelen llegar después del cambio de estado.
			time.Sleep(time.Second)
			recordFailure(PrintNewStackEvents(ctx, client, stackName, since, printedEventIDs))
			return status, rootFailureReason
		}

		time.Sleep(stackEventsPollInterval)
	}
}

// PrintNewStackEvents imprime los eventos posteriores a "since" que aún no se mostraron, del más
// antiguo al más reciente (la API los devuelve al revés). Devuelve la primera causa real de fallo
// del lote, si la hay.
func PrintNewStackEvents(ctx context.Context, client *cloudformation.Client,
	stackName string, since time.Time, printedEventIDs map[string]bool) string {

	result, err := client.DescribeStackEvents(ctx, &cloudformation.DescribeStackEventsInput{StackName: &stackName})
	if err != nil {
		return ""
	}

	pendingEvents := []cfnTypes.StackEvent{}
	for _, event := range result.StackEvents {
		if event.Timestamp == nil || event.Timestamp.Before(since) {
			continue
		}
		if event.EventId == nil || printedEventIDs[*event.EventId] {
			continue
		}
		printedEventIDs[*event.EventId] = true
		pendingEvents = append(pendingEvents, event)
	}

	sort.Slice(pendingEvents, func(a, b int) bool {
		return pendingEvents[a].Timestamp.Before(*pendingEvents[b].Timestamp)
	})

	rootFailureReason := ""
	for _, event := range pendingEvents {
		logicalID := aws.ToString(event.LogicalResourceId)
		reason := aws.ToString(event.ResourceStatusReason)

		line := fmt.Sprintf("  %v  %-30v %v",
			event.Timestamp.Local().Format("15:04:05"), logicalID, event.ResourceStatus)
		if len(reason) > 0 {
			line += " | " + reason
		}
		fmt.Println(line)

		// "Resource creation cancelled" es el efecto dominó del primer fallo, no una causa.
		if strings.HasSuffix(string(event.ResourceStatus), "_FAILED") && len(rootFailureReason) == 0 &&
			!strings.Contains(reason, "cancelled") {
			rootFailureReason = logicalID + ": " + reason
		}
	}

	return rootFailureReason
}

func isTerminalStackStatus(status cfnTypes.StackStatus) bool {
	return !strings.HasSuffix(string(status), "_IN_PROGRESS")
}

func isSuccessfulStackStatus(status cfnTypes.StackStatus) bool {
	return status == cfnTypes.StackStatusCreateComplete || status == cfnTypes.StackStatusUpdateComplete
}

// PrintStackOutputsAndSyncConfig muestra los outputs y guarda aws.lambda_url en config.toml.
func PrintStackOutputsAndSyncConfig(ctx context.Context, client *cloudformation.Client, stackName string) {
	result, err := client.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{StackName: &stackName})
	if err != nil || len(result.Stacks) == 0 {
		fmt.Println("No se pudieron leer los outputs del stack:", err)
		return
	}

	outputs := map[string]string{}
	fmt.Println("\nOutputs del stack:")
	for _, output := range result.Stacks[0].Outputs {
		key, value := aws.ToString(output.OutputKey), aws.ToString(output.OutputValue)
		outputs[key] = value
		fmt.Printf("  %-16v %v\n", key, value)
	}

	backendUrl := outputs["BackendUrl"]
	if len(backendUrl) == 0 {
		fmt.Println("\nEl stack no devolvió BackendUrl; aws.lambda_url no se modificó.")
		return
	}
	SyncLambdaUrlInConfig(backendUrl)
}

// Ancla en la clave de la sección [aws]; en el archivo solo existe una lambda_url.
var lambdaUrlInTomlPattern = regexp.MustCompile(`(?m)^(\s*lambda_url\s*=\s*")([^"]*)(")`)

// SyncLambdaUrlInConfig escribe la Function URL recién desplegada en config.toml, para que el
// frontend la lea de ahí y nadie tenga que copiarla a mano de la consola de AWS.
//
// Se hace por reemplazo de texto y no re-serializando el TOML: el archivo se mantiene a mano y sus
// comentarios son la razón de ser del formato.
func SyncLambdaUrlInConfig(deployedLambdaUrl string) {
	configPath := GetConfigPath()

	configBytes, err := ReadFile(configPath)
	if err != nil {
		fmt.Println("\nNo se pudo leer "+configPath+" para actualizar aws.lambda_url:", err)
		return
	}

	configText := string(configBytes)
	currentMatch := lambdaUrlInTomlPattern.FindStringSubmatch(configText)
	if currentMatch == nil {
		fmt.Println("\nNo se encontró la clave lambda_url en " + configPath + "; no se modificó nada.")
		fmt.Println("URL del backend desplegado: " + deployedLambdaUrl)
		return
	}

	previousLambdaUrl := currentMatch[2]
	if previousLambdaUrl == deployedLambdaUrl {
		fmt.Println("\naws.lambda_url ya apuntaba al backend desplegado: " + deployedLambdaUrl)
		return
	}

	updatedText := lambdaUrlInTomlPattern.ReplaceAllString(configText, "${1}"+deployedLambdaUrl+"${3}")
	if err := os.WriteFile(configPath, []byte(updatedText), 0644); err != nil {
		fmt.Println("\nNo se pudo escribir "+configPath+":", err)
		return
	}

	fmt.Println("\naws.lambda_url actualizado en " + configPath + ":")
	fmt.Println("  anterior: " + previousLambdaUrl)
	fmt.Println("  nuevo:    " + deployedLambdaUrl)

	// La URL cruda de Lambda no pasa por un dominio propio: si el valor anterior no era una
	// Function URL, se acaba de perder un dominio personalizado y hay que reapuntarlo.
	if len(previousLambdaUrl) > 0 && !strings.Contains(previousLambdaUrl, ".lambda-url.") {
		fmt.Println("  AVISO: el valor anterior era un dominio propio, no una Function URL.")
		fmt.Println("         Reapunta ese dominio a la nueva URL, o restaura el valor anterior.")
	}
}
