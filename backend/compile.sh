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
[-f marvelexplorers] && rm marvelexplorers

# Uncomment if need to build manually for local development
# # Ensure go.mod and go.sum are proper
# go mod tidy

# # Build the Go binary for Alpine (Linux, amd64)
# echo "Building binary for Alpine Linux (GOOS=linux, GOARCH=amd64)..."
# go build -o "$OUTPUT_NAME" "$SOURCE_FILE"

# # Check if the build was successful
# if [ $? -eq 0 ]; then
#     echo "Build successful! The binary is saved as '$OUTPUT_NAME'."
# else
#     echo "Build failed. Please check the errors above."
#     exit 1
# fi

[-f .env] && { echo ".env exists. continuing with deployment"; source ".env"; } || { echo ".env file not found. exiting."; exit 1; }

# Get latest artifact from workflows
bash artifact.sh || { echo "Failed to extract artifact from workflow. Exiting."; exit 1; }

# Verify the binary format
echo "Verifying the binary..."
file "$OUTPUT_NAME"

docker rm -f marvelbackend
docker rmi -f marvelbackend

docker build -t marvelbackend .
docker create --name marvelbackend -p 8000:8000 --network container_setup_default marvelbackend
docker start marvelbackend