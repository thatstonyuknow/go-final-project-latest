package handlers

import (
	"net/http"
	"strings"

	"task-tracker/pkg/db"
)

// DeleteTaskHandler handles DELETE requests to remove a task
func DeleteTaskHandler(store db.TaskStore, w http.ResponseWriter, r *http.Request) {
	// Get ID parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "ID not specified"})
		return
	}

	// Delete task from database
	if err := store.DeleteTask(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJson(w, http.StatusNotFound, map[string]string{"error": "Task not found"})
		} else {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
		}
		return
	}

	// Return empty JSON on success
	writeJson(w, http.StatusOK, map[string]interface{}{})
}
