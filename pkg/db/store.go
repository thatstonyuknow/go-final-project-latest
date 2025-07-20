package db

import "database/sql"

type TaskStore interface {
    GetTask(id string) (*Task, error)
    AddTask(task *Task) (int64, error)
    UpdateTask(task *Task) error
    DeleteTask(id string) error
    Tasks(limit int, search string) ([]*Task, error)
    UpdateDate(nextDate string, id string) error
}

type DataTaskStore struct {
    db *sql.DB
}

func NewDataTaskStore(db *sql.DB) TaskStore {
    return &DataTaskStore{db: db}
}