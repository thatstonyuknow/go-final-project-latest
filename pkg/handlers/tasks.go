package handlers

import (
	"database/sql"
	"net/http"

	"task-tracker/pkg/db"
)

// TasksResp represents the response structure for tasks list
type TasksResp struct {
	Tasks []TaskJSON `json:"tasks"`
}

// TasksHandler handles GET requests to retrieve list of tasks
func TasksHandler(database *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Get search parameter
	search := r.URL.Query().Get("search")

	tasks, err := db.Tasks(database, 100, search)
	if err != nil {
		writeJson(w, map[string]string{"error": "failed to get tasks"})
		return
	}

	// Convert to JSON format with string IDs
	jsonTasks := make([]TaskJSON, 0, len(tasks))
	for _, task := range tasks {
		jsonTasks = append(jsonTasks, convertTaskToJSON(task))
	}

	writeJson(w, TasksResp{
		Tasks: jsonTasks,
	})
}