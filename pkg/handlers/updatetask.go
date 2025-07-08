package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"task-tracker/pkg/db"
)

// UpdateTaskHandler handles PUT requests to update an existing task
func UpdateTaskHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
	var taskData struct {
		ID      string `json:"id"`
		Date    string `json:"date"`
		Title   string `json:"title"`
		Comment string `json:"comment"`
		Repeat  string `json:"repeat"`
	}

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&taskData); err != nil {
		writeJson(w, map[string]string{"error": "invalid JSON format"})
		return
	}

	// Validate ID is provided
	if strings.TrimSpace(taskData.ID) == "" {
		writeJson(w, map[string]string{"error": "ID not specified"})
		return
	}

	// Convert ID to int64
	id, err := strconv.ParseInt(taskData.ID, 10, 64)
	if err != nil {
		writeJson(w, map[string]string{"error": "Invalid ID format"})
		return
	}

	// Create task struct
	task := &db.Task{
		ID:      id,
		Date:    taskData.Date,
		Title:   taskData.Title,
		Comment: taskData.Comment,
		Repeat:  taskData.Repeat,
	}

	// Validate title is not empty
	if strings.TrimSpace(task.Title) == "" {
		writeJson(w, map[string]string{"error": "Task title not specified"})
		return
	}

	// Check and adjust date
	if err := checkDate(task); err != nil {
		writeJson(w, map[string]string{"error": "invalid date format or repeat rule"})
		return
	}

	// Update task in database
	if err := db.UpdateTask(database, task); err != nil {
		if strings.Contains(err.Error(), "incorrect id") {
			writeJson(w, map[string]string{"error": "Task not found"})
		} else {
			writeJson(w, map[string]string{"error": "failed to update task"})
		}
		return
	}

	// Return empty JSON on success
	writeJson(w, map[string]interface{}{})
}
