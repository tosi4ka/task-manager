package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	err := r.db.QueryRowContext(ctx, "INSERT INTO tasks (title, description, assigned_by, assigned_to, estimate) VALUES ($1, $2, $3, $4, $5) RETURNING created_at, id, status, updated_at, completed_at", t.Title, t.Description, t.AssignedBy, t.AssignedTo, t.Estimate).Scan(&t.CreatedAt, &t.ID, &t.Status, &t.UpdatedAt, &t.CompletedAt)

	if err != nil {
		return Task{}, fmt.Errorf("Task create error^ %s", err)
	}

	return t, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id uuid.UUID) (Task, error) {
	var t Task
	err := r.db.QueryRowContext(ctx, "SELECT title, description, assigned_by, assigned_to, created_at, estimate, id, status, updated_at, completed_at FROM tasks WHERE id = $1", id).Scan(&t.Title, &t.Description, &t.AssignedBy, &t.AssignedTo, &t.CreatedAt, &t.Estimate, &t.ID, &t.Status, &t.UpdatedAt, &t.CompletedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return Task{}, fmt.Errorf("task not found")
	}
	if err != nil {
		return Task{}, fmt.Errorf("get task error: %s", err)
	}

	return t, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, t Task) (Task, error) {
	err := r.db.QueryRowContext(ctx, "UPDATE tasks Set title=$1, description=$2, assigned_to=$3, estimate=$4, status=$5, completed_at=$6 WHERE id =$7 RETURNING updated_at", t.Title, t.Description, t.AssignedTo, t.Estimate, t.Status, t.CompletedAt, t.ID).Scan(&t.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return Task{}, fmt.Errorf("task not found")
	}
	if err != nil {
		return Task{}, fmt.Errorf("get task error: %s", err)
	}

	return t, nil
}
