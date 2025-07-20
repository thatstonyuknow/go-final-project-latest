package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"task-tracker/pkg/db"
)

// TaskJSON represents a task for JSON response with all fields as strings
type TaskJSON struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// writeJson writes JSON response to the client
func writeJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// convertTaskToJSON converts a db.Task to TaskJSON with string ID
func convertTaskToJSON(task *db.Task) TaskJSON {
	return TaskJSON{
		ID:      strconv.FormatInt(task.ID, 10),
		Date:    task.Date,
		Title:   task.Title,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}
}

// checkDate validates and adjusts task date according to requirements
func checkDate(task *db.Task) error {
	now := time.Now()

	// If date is empty, use today's date
	if task.Date == "" {
		task.Date = now.Format("20060102")
		return nil
	}

	// Parse the date to check if it's valid
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}

	// Check if we have a repeat rule
	if task.Repeat != "" {
		// Validate repeat rule and get next date
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		// If task date is less than today, use the calculated next date
		if t.Before(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())) {
			task.Date = next
		}
	} else {
		// No repeat rule - if date is less than today, use today's date
		if t.Before(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())) {
			task.Date = now.Format("20060102")
		}
	}

	return nil
}
