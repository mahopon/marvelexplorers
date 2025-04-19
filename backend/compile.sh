#!/bin/bash

# Define variables
OUTPUT_NAME="marvelexplorers"
SOURCE_FILE="main.go"

# Ensure the Go environment is set correctly
export PATH=$PATH:/usr/local/go/bin
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Remove current file
rm marvelexplorers

# Build the Go binary for Alpine (Linux, amd64)
echo "Building binary for Alpine Linux (GOOS=linux, GOARCH=amd64)..."
go build -o "$OUTPUT_NAME" "$SOURCE_FILE"

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! The binary is saved as '$OUTPUT_NAME'."
else
    echo "Build failed. Please check the errors above."
    exit 1
fi

# Verify the binary format
echo "Verifying the binary..."
file "$OUTPUT_NAME"

docker stop marvelbackend
docker rm marvelbackend
docker rmi marvelbackend

docker build -t marvelbackend .
docker create --name marvelbackend -p 8000:8000 --network container_setup_default marvelbackend
docker start marvelbackend