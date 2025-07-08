#!/bin/bash

# Build script for Docker image

echo "Building task-tracker application..."

# Build for Linux (required for Docker container)
echo "Compiling for Linux..."
GOOS=linux GOARCH=amd64 go build -o task-tracker .

if [ $? -ne 0 ]; then
    echo "Failed to compile application"
    exit 1
fi

echo "Building Docker image..."
docker build -t task-tracker:latest .

if [ $? -ne 0 ]; then
    echo "Failed to build Docker image"
    exit 1
fi

echo "Docker image built successfully!"
echo ""
echo "To run the container:"
echo "mkdir -p ./data"
echo "docker run -d --name task-tracker-app \\"
echo "  -p 7540:7540 \\"
echo "  -v \$(pwd)/data:/data \\"
echo "  task-tracker:latest"
echo ""
echo "Access the application at: http://localhost:7540"