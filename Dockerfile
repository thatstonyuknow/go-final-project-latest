# Use Ubuntu as base image
FROM ubuntu:latest

# Install necessary packages
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy the compiled binary
COPY task-tracker /app/task-tracker

# Copy web directory
COPY web /app/web

# Set environment variables
ENV TODO_PORT=7540
ENV TODO_DB_FILE=/data/scheduler.db
ENV TODO_WEB_DIR=./web
ENV TODO_LOG_LEVEL=info

# Expose the port
EXPOSE 7540

# Create volume mount point for database
VOLUME ["/data"]

# Make binary executable
RUN chmod +x /app/task-tracker


# Run the application
CMD ["/app/task-tracker"]