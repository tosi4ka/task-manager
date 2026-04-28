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
