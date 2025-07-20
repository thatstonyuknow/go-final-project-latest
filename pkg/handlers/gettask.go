package handlers

import (
	"net/http"

	"task-tracker/pkg/db"
)

// GetTaskHandler handles GET requests to retrieve a specific task by ID
func GetTaskHandler(store db.TaskStore, w http.ResponseWriter, r *http.Request) {
	// Get ID parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "ID not specified"})
		return
	}

	// Get task from database
	task, err := store.GetTask(id)
	if err != nil {
		writeJson(w, http.StatusNotFound, map[string]string{"error": "Task not found"})
		return
	}

	// Convert to JSON format with string ID
	taskJSON := convertTaskToJSON(task)

	writeJson(w, http.StatusOK, taskJSON)
}
