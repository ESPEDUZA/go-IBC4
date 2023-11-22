#!/bin/bash

# Vérifier si un tag d'image a été passé en argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <image-tag>"
    exit 1
fi

IMAGE_TAG=$1

IMAGE_NAME="espeduza/nangaparbat:$IMAGE_TAG"

echo "Pulling image $IMAGE_NAME..."
docker pull $IMAGE_NAME

if [ $? -ne 0 ]; then
    echo "Failed to pull image $IMAGE_NAME"
    exit 1
fi

echo "Running image $IMAGE_NAME..."
docker run -p 3000:3000 $IMAGE_NAME