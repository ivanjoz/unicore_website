package main

import (
	"archive/zip"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// El fallo que esta prueba evita: un parámetro renombrado en un sitio y no en el otro.
// CloudFormation rechaza un parámetro que no declara, y acepta en silencio no recibir uno que sí
// declara —usando su Default—, así que la segunda mitad es la que de verdad cuesta descubrir.
func TestDeployerAndTemplateAgreeOnEveryParameter(t *testing.T) {
	declared := templateParameterNames()
	if len(declared) == 0 {
		t.Fatal("no se leyó ningún parámetro del YAML embebido; ¿cambió su formato?")
	}

	sent := []string{}
	for _, parameter := range makeTemplateParameters(DeployParams{}) {
		sent = append(sent, aws.ToString(parameter.ParameterKey))
	}

	sort.Strings(declared)
	sort.Strings(sent)

	if strings.Join(declared, ",") != strings.Join(sent, ",") {
		t.Fatalf("la plantilla y el deployer no coinciden:\n  plantilla: %v\n  deployer:  %v",
			declared, sent)
	}
}

// Toda variable que declara la plantilla tiene que repetirse en UpdateEnvironmentVariables:
// UpdateFunctionConfiguration reemplaza el bloque entero, así que una que falte aquí desaparece
// de la Lambda en cuanto la acción 3 reinyecta el entorno.
func TestEveryTemplateEnvironmentVariableIsReinjected(t *testing.T) {
	declared := templateEnvironmentVariableNames(t)
	if len(declared) == 0 {
		t.Fatal("no se leyó ninguna variable de entorno del YAML embebido")
	}

	reinjected := map[string]bool{}
	for name := range lambdaEnvironmentVariables(DeployParams{}) {
		reinjected[name] = true
	}

	for _, name := range declared {
		if !reinjected[name] {
			t.Fatalf("la plantilla declara %v pero UpdateEnvironmentVariables no la reinyecta: "+
				"se borraría de la Lambda en el próximo despliegue de infraestructura", name)
		}
	}
}

// Lee las claves del bloque Environment.Variables de la plantilla, que están indentadas con diez
// espacios bajo "Variables:".
func templateEnvironmentVariableNames(t *testing.T) []string {
	t.Helper()

	names := []string{}
	insideVariables := false
	for _, line := range strings.Split(cloudFormationTemplate, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "Variables:" {
			insideVariables = true
			continue
		}
		if !insideVariables {
			continue
		}
		if !strings.HasPrefix(line, "          ") || len(trimmed) == 0 || trimmed[0] == '#' {
			break
		}
		names = append(names, strings.TrimSpace(strings.Split(trimmed, ":")[0]))
	}
	return names
}

// El .zip lo construye Go y no el comando `zip` justamente por el bit de ejecución: sin él el
// runtime no arranca el binario y el error no menciona los permisos por ningún lado.
func TestCompressedExecutableIsABootstrapWithTheExecutableBit(t *testing.T) {
	temporaryDirectory := t.TempDir()
	binaryPath := filepath.Join(temporaryDirectory, "backend-compiled")
	if err := os.WriteFile(binaryPath, []byte("no soy un binario de verdad"), 0644); err != nil {
		t.Fatalf("no se pudo escribir el binario de prueba: %v", err)
	}

	if err := CompressExecutable(binaryPath, binaryPath+".zip"); err != nil {
		t.Fatalf("CompressExecutable falló: %v", err)
	}

	reader, err := zip.OpenReader(binaryPath + ".zip")
	if err != nil {
		t.Fatalf("el .zip no se puede abrir: %v", err)
	}
	defer reader.Close()

	if len(reader.File) != 1 {
		t.Fatalf("el .zip tiene %v entradas, se esperaba 1", len(reader.File))
	}
	entry := reader.File[0]
	if entry.Name != "bootstrap" {
		t.Fatalf("la entrada se llama %q; Lambda espera exactamente \"bootstrap\"", entry.Name)
	}
	if mode := entry.Mode(); mode&0111 == 0 {
		t.Fatalf("la entrada no es ejecutable: %v", mode)
	}
}

func TestSyncLambdaUrlRewritesOnlyTheValue(t *testing.T) {
	temporaryDirectory := t.TempDir()
	configPath := filepath.Join(temporaryDirectory, "config.toml")
	original := "app_name = \"unicore\"\n\n[aws]\n# un comentario que debe sobrevivir\nregion = \"us-east-1\"\nlambda_url = \"\"\n"
	if err := os.WriteFile(configPath, []byte(original), 0644); err != nil {
		t.Fatalf("no se pudo escribir el config de prueba: %v", err)
	}

	t.Setenv("UNICORE_CONFIG_FILE", configPath)
	SyncLambdaUrlInConfig("https://abc123.lambda-url.us-east-1.on.aws/")

	updated, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("no se pudo releer el config: %v", err)
	}
	text := string(updated)

	if !strings.Contains(text, `lambda_url = "https://abc123.lambda-url.us-east-1.on.aws/"`) {
		t.Fatalf("la URL no se escribió:\n%v", text)
	}
	// La razón de reemplazar texto en vez de re-serializar el TOML.
	if !strings.Contains(text, "# un comentario que debe sobrevivir") {
		t.Fatalf("se perdieron los comentarios del archivo:\n%v", text)
	}
	if !strings.Contains(text, `region = "us-east-1"`) {
		t.Fatalf("se perdió otra clave del archivo:\n%v", text)
	}
}
