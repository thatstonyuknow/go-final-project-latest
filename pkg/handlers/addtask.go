package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"task-tracker/pkg/db"
)

// AddTaskHandler handles POST requests to add a new task
func AddTaskHandler(store db.TaskStore, w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON format"})
		return
	}

	// Validate title is not empty
	if strings.TrimSpace(task.Title) == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Task title not specified"})
		return
	}

	// Check and adjust date
	if err := checkDate(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid date format or repeat rule"})
		return
	}

	// Add task to database
	id, err := store.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to add task"})
		return
	}

	// Return success response with ID
	writeJson(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(id, 10)})
}
