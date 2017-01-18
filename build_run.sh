#!/bin/bash

# Variables
IMAGE_NAME="myapp"
CONTAINER_NAME="myapp"
BINARY_NAME="chat"
PORT="8080"

# Print log message
function printlog() {
    echo ""; echo "$1"
}

# Remove old container
printlog "Removing container and binary"
docker rm -f $CONTAINER_NAME 2> /dev/null
rm $BINARY_NAME 2> /dev/null
docker rmi $(docker images -f dangling=true -q) 2> /dev/null

# Build binary
printlog "Building app binary"
GOOS="linux" GOARCH="amd64" go build -o chat/$BINARY_NAME ./chat/

# Build Docker image
printlog "Building docker image"
docker build --no-cache -t $CONTAINER_NAME .

# Run Docker image
docker run --name $CONTAINER_NAME -d -p $PORT:$PORT $CONTAINER_NAME
printlog "Running app on localhost:$PORT"