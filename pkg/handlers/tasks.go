package handlers

import (
	"net/http"

	"task-tracker/pkg/db"
)

// TasksResp represents the response structure for tasks list
type TasksResp struct {
	Tasks []TaskJSON `json:"tasks"`
}

// TasksHandler handles GET requests to retrieve list of tasks
func TasksHandler(store db.TaskStore, w http.ResponseWriter, r *http.Request) {
	// Get search parameter
	search := r.URL.Query().Get("search")

	tasks, err := store.Tasks(100, search)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "failed to get tasks"})
		return
	}

	// Convert to JSON format with string IDs
	jsonTasks := make([]TaskJSON, 0, len(tasks))
	for _, task := range tasks {
		jsonTasks = append(jsonTasks, convertTaskToJSON(task))
	}

	writeJson(w, http.StatusOK, TasksResp{
		Tasks: jsonTasks,
	})
}
