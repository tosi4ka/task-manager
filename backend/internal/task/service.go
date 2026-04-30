package task

import (
	"context"

	"github.com/google/uuid"
)

type TaskService struct {
	repo TaskRepo
}

func NewTaskService(repo TaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, req CreateTaskRequest) (Task, error) {
	if req.Title == "" {
		return Task{}, ErrTitleRequired
	}

	if req.Description == "" {
		return Task{}, ErrDescriptionRequired
	}

	if req.AssignedBy == (uuid.UUID{}) {
		return Task{}, ErrAssignedRequired
	}

	if req.AssignedTo == (uuid.UUID{}) {
		return Task{}, ErrAssignedRequired
	}

	t := Task{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		AssignedBy:  req.AssignedBy,
		AssignedTo:  req.AssignedTo,
		Estimate:    req.Estimate,
		Status:      StatusTodo,
	}

	task, err := s.repo.CreateTask(ctx, t)
	if err != nil {
		return Task{}, ErrInternal
	}

	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id uuid.UUID, req UpdateTaskRequest) (Task, error) {
	if req.Title == nil &&
		req.Description == nil &&
		req.AssignedTo == nil &&
		req.Estimate == nil &&
		req.Status == nil &&
		req.CompletedAt == nil {
		return Task{}, ErrNothingToUpdate
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return Task{}, ErrTaskNotFound
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.AssignedTo != nil {
		task.AssignedTo = *req.AssignedTo
	}
	if req.Estimate != nil {
		task.Estimate = *req.Estimate
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.CompletedAt != nil {
		task.CompletedAt = req.CompletedAt
	}

	updatedTask, err := s.repo.UpdateTask(ctx, task)
	if err != nil {
		return Task{}, ErrInternal
	}

	return updatedTask, nil
}

func (s *TaskService) ListTasks(ctx context.Context, id uuid.UUID) ([]Task, error) {
	if id == (uuid.UUID{}) {
		return []Task{}, ErrUserNotFound
	}

	tasks, err := s.repo.ListTasks(ctx, id)
	if err != nil {
		return []Task{}, ErrInternal
	}

	return tasks, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if id == (uuid.UUID{}) {
		return ErrTaskNotFound
	}

	err := s.repo.DeleteTask(ctx, id)
	if err != nil {
		return ErrInternal
	}

	return nil
}
