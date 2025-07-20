package main

import (
	"fmt"
	"net/http"
	"task-tracker/pkg/db"
	server "task-tracker/pkg/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Initialize database
	log.Info("Opening DB connection")
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	taskStore := db.NewDataTaskStore(database)

	// Load configuration
	config := server.LoadConfig()

	router := server.SetupRouter(config, taskStore)

	// Start server
	log.Infof("Listening server on port %s\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router))
}
