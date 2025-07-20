package db

import (
    "fmt"
)

// UpdateTask updates an existing task
func (d *DataTaskStore) UpdateTask(task *Task) error {
    query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
    
    res, err := d.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
    if err != nil {
        return fmt.Errorf("failed to update task: %v", err)
    }
    
    count, err := res.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get affected rows: %v", err)
    }
    
    if count == 0 {
        return fmt.Errorf("incorrect id for updating task")
    }
    
    return nil
}