package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"task-tracker/pkg/db"
)

// DeleteTaskHandler handles DELETE requests to remove a task
func DeleteTaskHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Get ID parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, map[string]string{"error": "ID not specified"})
		return
	}

	// Delete task from database
	if err := db.DeleteTask(database, id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJson(w, map[string]string{"error": "Task not found"})
		} else {
			writeJson(w, map[string]string{"error": "Failed to delete task"})
		}
		return
	}

	// Return empty JSON on success
	writeJson(w, map[string]interface{}{})
}
