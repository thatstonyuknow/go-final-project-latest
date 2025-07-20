# Dockerfile for building and running task-tracker application
# Use a multi-stage build to create a minimal final image

# Stage 1: Build the Go application
FROM golang:1.24.2-alpine AS builder

# Install build dependencies (git for go modules, ca-certificates)
RUN apk add --no-cache git ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the Go source files to the working directory
COPY . ./

# Build the application 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -trimpath \
    -o /task-tracker \
    .

# Stage 2: Create a minimal final image
FROM alpine:latest

# Install ca-certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Create non-root user for security
RUN adduser -D -u 1000 appuser

# Set the working directory inside the final image
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /task-tracker ./task-tracker

# Copy web directory with static files
COPY --from=builder /app/web ./web

# Create data directory and set permissions
RUN mkdir -p /data && chown -R appuser:appuser /app /data

# Set environment variables with defaults
ENV TODO_PORT=7540
ENV TODO_DB_FILE=/data/scheduler.db
ENV TODO_WEB_DIR=./web
ENV TODO_LOG_LEVEL=info

# Create volume mount point for database
VOLUME ["/data"]

# Switch to non-root user
USER appuser

# Set the command to run the application when the container starts
CMD ["./task-tracker"]

# Build command:
# docker build -t task-tracker:latest .
#
# Run command:
# docker run -d --name task-tracker \
#   -p 7540:7540 \
#   -v $(pwd)/data:/data \
#   task-tracker:latest
#
# To use a different port:
# docker run -d --name task-tracker \
#   -e TODO_PORT=8080 \
#   -p 8080:8080 \
#   -v $(pwd)/data:/data \
#   task-tracker:latest