package server

import (
	"net/http"
	"os"

	"task-tracker/pkg/api"
	"task-tracker/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(config *Config, store db.TaskStore) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(config.GetTimeout()))

	// Setup routes
	setupAPIRoutes(r, store)
	setupStaticRoutes(r, config)

	return r
}

func setupAPIRoutes(r *chi.Mux, store db.TaskStore) {
	r.Route("/api", func(apiRouter chi.Router) {
		// Initialize API handlers with database connection
		api.Init(apiRouter, store)
	})
}

func setupStaticRoutes(r *chi.Mux, config *Config) {
	// Static file server
	workDir, _ := os.Getwd()
	filesDir := http.Dir(workDir + "/" + config.WebDir)
	fileServer := http.FileServer(filesDir)
	r.Handle("/*", fileServer)
}
