#!/bin/bash

# Variables
IMAGE_NAME="myapp"
CONTAINER_NAME="myapp"
BINARY_NAME="app"

# Print log message
function printlog() {
    echo ""; echo "$1"
}

# Remove old container
printlog "Removing container"
docker rm -f $CONTAINER_NAME 2> /dev/null

# Build binary
printlog "Building app binary"
GOOS="linux" GOARCH="amd64" go build -o $BINARY_NAME .

# Build Docker image
printlog "Building docker image"
docker build --no-cache -t $CONTAINER_NAME .

# Run Docker image
printlog "Running app on localhost:8080"
docker run --name $CONTAINER_NAME -d -p 8080:8080 $CONTAINER_NAME