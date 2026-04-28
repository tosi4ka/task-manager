package task

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

type mockTaskRepo struct {
	task map[uuid.UUID]Task
}

func newMockRepo() *mockTaskRepo {
	return &mockTaskRepo{
		task: make(map[uuid.UUID]Task),
	}
}

func (m *mockTaskRepo) CreateTask(ctx context.Context, t Task) (Task, error) {
	m.task[t.ID] = t
	return t, nil
}

func TestCreateTask(t *testing.T) {
	test := []struct {
		name      string
		req       CreateTaskRequest
		expectErr bool
	}{
		{name: "Created successfully",
			req: CreateTaskRequest{
				Title:       "Title",
				Description: "You need to do something.",
				AssignedBy:  uuid.New(),
				AssignedTo:  uuid.New(),
				Estimate:    20,
			},
			expectErr: false},
		{name: "empty title",
			req: CreateTaskRequest{
				Title:       "",
				Description: "You need to do something.",
				AssignedBy:  uuid.New(),
				AssignedTo:  uuid.New(),
				Estimate:    20,
			},
			expectErr: true},
		{name: "empty description",
			req: CreateTaskRequest{
				Title:       "Title",
				Description: "",
				AssignedBy:  uuid.New(),
				AssignedTo:  uuid.New(),
				Estimate:    20,
			},
			expectErr: true},
		{name: "empty assignedBy",
			req: CreateTaskRequest{
				Title:       "Title",
				Description: "You need to do something.",
				AssignedBy:  uuid.UUID{},
				AssignedTo:  uuid.New(),
				Estimate:    20,
			},
			expectErr: true},
		{name: "empty assignedTo",
			req: CreateTaskRequest{
				Title:       "Title",
				Description: "You need to do something.",
				AssignedBy:  uuid.New(),
				AssignedTo:  uuid.UUID{},
				Estimate:    20,
			},
			expectErr: true},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewTaskService(newMockRepo())

			_, err := svc.CreateTask(context.Background(), tt.req)

			if tt.expectErr && err == nil {
				t.Errorf("Error asked me to convey that she is busy.")
			}

			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect an error: %v", err)
			}
		})
	}
}

func TestAppError(t *testing.T) {
	err := ErrTitleRequired
	if err.Error() != "title is required" {
		t.Errorf("expected 'title is required', got %s", err.Error())
	}
}
