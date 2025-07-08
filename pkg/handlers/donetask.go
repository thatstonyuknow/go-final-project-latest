package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"task-tracker/pkg/db"
)

// DoneTaskHandler handles POST requests to mark a task as done
func DoneTaskHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
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
	
	// Check if task has repeat rule
	if strings.TrimSpace(task.Repeat) == "" {
		// No repeat rule - delete the task
		if err := db.DeleteTask(database, id); err != nil {
			writeJson(w, map[string]string{"error": "Failed to delete task"})
			return
		}
	} else {
		// Task has repeat rule - calculate next date using NextDate function
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error": "Failed to calculate next date"})
			return
		}
		
		// Update task date
		if err := db.UpdateDate(database, nextDate, id); err != nil {
			writeJson(w, map[string]string{"error": "Failed to update task date"})
			return
		}
	}
	
	// Return empty JSON on success - use map[string]string for consistency
	writeJson(w, map[string]string{})
}