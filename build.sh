#!/bin/bash

# Build script for Docker image

echo "Building task-tracker application..."

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