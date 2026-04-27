package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AssignedBy  uuid.UUID `json:"assigned_by"`
	AssignedTo  uuid.UUID `json:"assigned_to"`
	CreatedAt   time.Time `json:"created_at"`
	Estimate    int       `json:"estimate"`

	ID          uuid.UUID  `json:"id"`
	Status      string     `json:"status"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AssignedBy  uuid.UUID `json:"assigned_by"`
	AssignedTo  uuid.UUID `json:"assigned_to"`
	Estimate    int       `json:"estimate"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
	Estimate    *int       `json:"estimate"`

	Status      *string    `json:"status"`
	CompletedAt *time.Time `json:"completed_at"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)
