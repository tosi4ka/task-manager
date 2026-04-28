package task

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type TaskRepository struct {
	db *sql.DB
}

type TaskRepo interface {
	CreateTask(ctx context.Context, t Task) (Task, error)
	UpdateTask(ctx context.Context, t Task) (Task, error)
	GetByID(ctx context.Context, id uuid.UUID) (Task, error)
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, t Task) (Task, error) {
	return Task{}, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id uuid.UUID) (Task, error) {
	return Task{}, nil
}
