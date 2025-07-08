package db

import (
	"database/sql"
	"fmt"
)

// DeleteTask deletes a task by ID
func DeleteTask(db *sql.DB, id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`

	res, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}

// UpdateDate updates only the date field of a task
func UpdateDate(db *sql.DB, nextDate string, id string) error {
	query := `UPDATE scheduler SET date = ? WHERE id = ?`

	res, err := db.Exec(query, nextDate, id)
	if err != nil {
		return fmt.Errorf("failed to update task date: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if count == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
