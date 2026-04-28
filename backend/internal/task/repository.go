package task

import (
	"context"
	"database/sql"
)

type TaskRepository struct {
	db *sql.DB
}

type TaskRepo interface {
	CreateTask(ctx context.Context, t Task) (Task, error)
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, t Task) (Task, error) {
	return Task{}, nil
}
