package api

import (
	"net/http"

	"task-tracker/pkg/db"
	"task-tracker/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

// Init API routes
func Init(r chi.Router, store db.TaskStore) {
	// GET endpoints
	r.Get("/nextdate", handlers.NextDateHandler)
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.TasksHandler(store, w, r)
	})

	r.Route("/task", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.AddTaskHandler(store, w, r)
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.GetTaskHandler(store, w, r)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateTaskHandler(store, w, r)
		})
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteTaskHandler(store, w, r)
		})
		r.Post("/done", func(w http.ResponseWriter, r *http.Request) {
			handlers.DoneTaskHandler(store, w, r)
		})
	})
}
