package db

import (
	"database/sql"
	"fmt"
)

// GetTask returns a task by ID
func GetTask(db *sql.DB, id string) (*Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	
	task := &Task{}
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	
	return task, nil
}

// Tasks returns a list of upcoming tasks sorted by date, optionally filtered by search
func Tasks(db *sql.DB, limit int, search string) ([]*Task, error) {
	var query string
	var args []interface{}
	
	if search != "" {
		// Try to parse date in format dd.mm.yyyy and convert to yyyymmdd
		searchDate := ""
		if len(search) == 10 && search[2] == '.' && search[5] == '.' {
			// Format: dd.mm.yyyy -> yyyymmdd
			day := search[0:2]
			month := search[3:5]
			year := search[6:10]
			searchDate = year + month + day
		}
		
		if searchDate != "" {
			// Search by converted date
			query = `SELECT id, date, title, comment, repeat FROM scheduler 
					 WHERE date = ? ORDER BY date ASC LIMIT ?`
			args = []interface{}{searchDate, limit}
		} else {
			// Search in title, comment, or original date
			query = `SELECT id, date, title, comment, repeat FROM scheduler 
					 WHERE title LIKE ? OR comment LIKE ? OR date LIKE ?
					 ORDER BY date ASC LIMIT ?`
			searchParam := "%" + search + "%"
			args = []interface{}{searchParam, searchParam, searchParam, limit}
		}
	} else {
		query = `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`
		args = []interface{}{limit}
	}
	
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %v", err)
	}
	defer rows.Close()
	
	var tasks []*Task
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %v", err)
		}
		tasks = append(tasks, task)
	}
	
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}
	
	// Ensure we return empty slice instead of nil
	if tasks == nil {
		tasks = []*Task{}
	}
	
	return tasks, nil
}