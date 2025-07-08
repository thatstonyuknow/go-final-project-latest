package api

import (
	"database/sql"
	"net/http"

	"task-tracker/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

// taskHandler handles different HTTP methods for /api/task
func taskHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlers.AddTaskHandler(db, w, r)
	case http.MethodGet:
		handlers.GetTaskHandler(db, w, r)
	case http.MethodPut:
		handlers.UpdateTaskHandler(db, w, r)
	case http.MethodDelete:
		handlers.DeleteTaskHandler(db, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Init initializes API routes
func Init(r chi.Router, db *sql.DB) {
	r.Get("/nextdate", handlers.NextDateHandler)
	r.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		taskHandler(db, w, r)
	})
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.TasksHandler(db, w, r)
	})
	r.HandleFunc("/task/done", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.DoneTaskHandler(db, w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
