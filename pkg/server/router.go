package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"task-tracker/pkg/api"
)

func SetupRouter(config *Config, db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(config.GetTimeout()))

	// Setup routes
	setupAPIRoutes(r, db)
	setupStaticRoutes(r, config)

	return r
}

func setupAPIRoutes(r *chi.Mux, db *sql.DB) {
	r.Route("/api", func(apiRouter chi.Router) {
		// Initialize API handlers with database connection
		api.Init(apiRouter, db)
	})
}

func setupStaticRoutes(r *chi.Mux, config *Config) {
	// Static file server
	workDir, _ := os.Getwd()
	filesDir := http.Dir(workDir + "/" + config.WebDir)
	fileServer := http.FileServer(filesDir)
	r.Handle("/*", fileServer)
}