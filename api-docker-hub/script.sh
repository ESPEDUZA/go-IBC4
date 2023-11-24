#!/bin/bash
echo "Creating container..."

# Vérifier si un tag d'image a été passé en argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <image-tag>"
    exit 1
fi

# Nom et tag de l'image
IMAGE_TAG=$1
IMAGE_NAME="espeduza/nangaparbat:$IMAGE_TAG"
CONTAINER_NAME="nangaparbat_container"

# Arrêter et supprimer le conteneur existant (si existant)
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

# Pull l'image Docker depuis Docker Hub
echo "Pulling image $IMAGE_NAME..."
docker pull $IMAGE_NAME
if [ $? -ne 0 ]; then
    echo "Failed to pull image $IMAGE_NAME"
    exit 1
fi

# Démarrer un nouveau conteneur avec l'image pullée
echo "Running image $IMAGE_NAME..."
docker run -d -p 3000:3000 --name $CONTAINER_NAME $IMAGE_NAME