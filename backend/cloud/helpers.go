package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// MakeAwsConfig resuelve credenciales por el perfil nombrado en config.toml, y no por la cadena
// por defecto: en una máquina con varias cuentas configuradas, "la que esté activa en el entorno"
// es la forma de desplegar a la cuenta equivocada sin enterarse.
func MakeAwsConfig(params DeployParams) aws.Config {
	awsConfig, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithSharedConfigProfile(params.AWS.Profile),
		awsconfig.WithRegion(params.AWS.Region))

	if err != nil {
		exitWithError("Error al cargar la configuración de AWS: " + err.Error())
	}
	return awsConfig
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// SendFileToS3 sube un archivo por streaming, sin cargarlo entero en memoria.
func SendFileToS3(client *s3.Client, bucket string, key string, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	return err
}
