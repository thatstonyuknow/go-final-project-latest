package db

import (
	"database/sql"
	"fmt"
)

// AddTask adds a new task to the database and returns the ID
func AddTask(db *sql.DB, task *Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("failed to insert task: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %v", err)
	}

	return id, nil
}
