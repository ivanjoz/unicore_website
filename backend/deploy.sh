#!/usr/bin/env bash
# Wrapper del deployer. Toda la lógica vive en cloud/ (Go), que lleva la plantilla de
# CloudFormation embebida y habla directo con la API de AWS: sin CLI de aws, sin CDK, sin Node.
#
#   ./deploy.sh          -> pregunta la acción
#   ./deploy.sh 1        -> Publicar Código (compila y sube el .zip a la Lambda)
#   ./deploy.sh 2        -> Actualizar Variables de entorno
#   ./deploy.sh 3        -> Desplegar Infraestructura (CloudFormation) y reinyectar variables
#
# UNICORE_CONFIG_FILE selecciona otro config.toml:
#   UNICORE_CONFIG_FILE=/ruta/config.prod.toml ./deploy.sh 3
set -euo pipefail

cd "$(dirname "$0")/cloud"

GO_PATH="go"
if [ -x /usr/local/go/bin/go ]; then
    GO_PATH="/usr/local/go/bin/go"
fi

exec "$GO_PATH" run . "$@"
