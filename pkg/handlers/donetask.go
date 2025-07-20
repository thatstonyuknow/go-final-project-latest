package handlers

import (
	"net/http"
	"strings"
	"time"

	"task-tracker/pkg/db"
)

// DoneTaskHandler handles POST requests to mark a task as done
func DoneTaskHandler(store db.TaskStore, w http.ResponseWriter, r *http.Request) {
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
	
	// Check if task has repeat rule
	if strings.TrimSpace(task.Repeat) == "" {
		// No repeat rule - delete the task
		if err := store.DeleteTask(id); err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
			return
		}
	} else {
		// Task has repeat rule - calculate next date using NextDate function
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to calculate next date"})
			return
		}
		
		// Update task date
		if err := store.UpdateDate(nextDate, id); err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update task date"})
			return
		}
	}
	
	// Return empty JSON on success - use map[string]string for consistency
	writeJson(w, http.StatusOK, map[string]string{})
}