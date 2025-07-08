package handlers

import (
	"database/sql"
	"net/http"

	"task-tracker/pkg/db"
)

// GetTaskHandler handles GET requests to retrieve a specific task by ID
func GetTaskHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Get ID parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, map[string]string{"error": "ID not specified"})
		return
	}

	// Get task from database
	task, err := db.GetTask(database, id)
	if err != nil {
		writeJson(w, map[string]string{"error": "Task not found"})
		return
	}

	// Convert to JSON format with string ID
	taskJSON := convertTaskToJSON(task)

	writeJson(w, taskJSON)
}
