# Task Tracker

## Project Description

Task Tracker is a Go-based web server for managing tasks:

* Create, delete, and update tasks
* Retrieve the list of tasks
* Support for task repetition rules
* Calculate the next execution date based on a repetition rule

## Completed Features

* Create task
* Delete task
* Update task
* Retrieve task list
* Next execution date calculation based on repetition rules \*

*Advanced feature (marked with an asterisk) implemented: calculation of the next execution date according to repetition rules.*

## Local Setup

1. Clone the repository and navigate to the project directory:

   ```bash
   git clone <your-repository-url>
   cd task-tracker
   ```
2. Create a `.env` file in the project root (example):

   ```env
   TODO_DB_FILE=../scheduler.db
   PORT=7540
   ```
3. Run the server:

   ```bash
   go run ./pkg/api
   ```
4. Open in your browser: [http://localhost:7540](http://localhost:7540)

## Running Tests

1. Ensure the settings in `tests/settings.go` match:

   ```go
   var Port = 7540
   var DBFile = "../scheduler.db"
   var FullNextDate = true
   var Search = true
   var Token = ""
   ```
2. Run the tests:

   ```bash
   go test -count=1 ./tests
   ```

## Docker

1. Build the Docker image:

   ```bash
   docker build -t task-tracker .
   ```
2. Create the data directory and run the container:

   ```bash
   mkdir -p ./data
   docker run -d --name task-tracker-app \
     -p 7540:7540 \
     -v "$(pwd)/data:/data" \
     task-tracker:latest
   ```
3. The server will be available at [http://localhost:7540](http://localhost:7540)



# Task Tracker Docker Setup

## Requirements
- Docker installed on the system
- Go 1.19+ (for compilation)

## Building Docker Image

### Method 1: Automatic Build (Recommended)
```bash
# Make script executable
chmod +x build-docker.sh

# Run build
./build-docker.sh
```

### Method 2: Manual Build
```bash
# 1. Compile application for Linux
GOOS=linux GOARCH=amd64 go build -o task-tracker .

# 2. Build Docker image
docker build -t task-tracker:latest .
```

## Running Container

### Create Database Directory
```bash
mkdir -p ./data
```

### Method 1: Docker Command
```bash
docker run -d --name task-tracker-app \
  -p 7540:7540 \
  -v $(pwd)/data:/data \
  task-tracker:latest
```

## Testing the Application

1. Open browser and navigate to: http://localhost:7540
2. Test API: `curl http://localhost:7540/api/tasks`



## Environment Variables

- `TODO_PORT` - web server port (default: 7540)
- `TODO_DB_FILE` - database file path (default: /data/scheduler.db)
- `TODO_WEB_DIR` - web files directory (default: ./web)
- `TODO_LOG_LEVEL` - logging level (default: info)

## Notes

- SQLite database is persisted on host in `./data` directory
- Container uses port 7540
- Web interface available at http://localhost:7540
