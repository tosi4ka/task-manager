package task

import (
	"context"
	"testing"
	"time"

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

func (m *mockTaskRepo) GetByID(ctx context.Context, id uuid.UUID) (Task, error) {
	t, ok := m.task[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return t, nil
}

func (m *mockTaskRepo) UpdateTask(ctx context.Context, t Task) (Task, error) {
	m.task[t.ID] = t
	return t, nil
}

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

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

func TestUpdateTask(t *testing.T) {
	test := []struct {
		name      string
		id        uuid.UUID
		req       UpdateTaskRequest
		expectErr bool
	}{
		{
			name:      "Updated successfully",
			id:        uuid.New(),
			req:       UpdateTaskRequest{Title: strPtr("New title")},
			expectErr: false,
		},
		{
			name:      "All fields empty",
			id:        uuid.New(),
			req:       UpdateTaskRequest{},
			expectErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			repo.task[tt.id] = Task{ID: tt.id}
			svc := NewTaskService(repo)

			_, err := svc.UpdateTask(context.Background(), tt.id, tt.req)

			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect error: %v", err)
			}
		})
	}
}

func (m *mockTaskRepo) ListTasks(ctx context.Context, id uuid.UUID) ([]Task, error) {
	var tasks []Task
	for _, t := range m.task {
		if t.AssignedTo == id {
			tasks = append(tasks, t)
		}
	}
	return tasks, nil
}

func TestListTask(t *testing.T) {
	userID := uuid.New()
	completedAt := time.Now()

	test := []struct {
		name          string
		id            uuid.UUID
		seedTasks     []Task
		expectedCount int
		expectErr     bool
	}{
		{
			name: "user has tasks",
			id:   userID,
			seedTasks: []Task{
				{
					Title:       "Assemble an intergalactic antimatter condenser",
					Description: "We need to fly to Blips and Cheets, get three Class IX crystals, and not lose a single one. Don't screw up, Morty.",
					AssignedBy:  uuid.New(),
					AssignedTo:  userID,
					CreatedAt:   time.Now(),
					Estimate:    10,
					ID:          uuid.New(),
					Status:      StatusInProgress,
					UpdatedAt:   time.Now(),
					CompletedAt: nil,
				},
				{
					Title:       "Fix megaseeds admin interface",
					Description: "The 'Destroy All' button is back on the home page. Remove it from the settings and add a confirmation. Well... no, let there be a confirmation.",
					AssignedBy:  uuid.New(),
					AssignedTo:  userID,
					CreatedAt:   time.Now(),
					Estimate:    15,
					ID:          uuid.New(),
					Status:      StatusInProgress,
					UpdatedAt:   time.Now(),
					CompletedAt: &completedAt,
				},
			},
			expectedCount: 2,
			expectErr:     false,
		},
		{
			name:          "user not found",
			id:            uuid.UUID{},
			seedTasks:     nil,
			expectedCount: 0,
			expectErr:     true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()

			for _, task := range tt.seedTasks {
				repo.task[task.ID] = task
			}

			svc := NewTaskService(repo)

			result, err := svc.ListTasks(context.Background(), tt.id)

			if len(result) != tt.expectedCount {
				t.Errorf("expected %d tasks, got %d", tt.expectedCount, len(result))
			}
			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect error: %v", err)
			}
		})
	}
}

func (m *mockTaskRepo) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, ok := m.task[id]
	if !ok {
		return ErrTaskNotFound
	}

	delete(m.task, id)

	return nil
}

func TestDeleteTask(t *testing.T) {
	test := []struct {
		name      string
		id        uuid.UUID
		expectErr bool
	}{
		{
			name:      "Delete successful",
			id:        uuid.New(),
			expectErr: false,
		},
		{
			name:      "Error delete",
			id:        uuid.UUID{},
			expectErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			repo.task[tt.id] = Task{ID: tt.id}

			svc := NewTaskService(repo)

			ctx := context.Background()
			err := svc.DeleteTask(ctx, tt.id)

			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect error: %v", err)
			}

		})
	}
}

func TestGetByID(t *testing.T) {
	test := []struct {
		name      string
		id        uuid.UUID
		expectErr bool
	}{
		{
			name:      "task has been found",
			id:        uuid.New(),
			expectErr: false,
		},
		{
			name:      "task not find",
			id:        uuid.UUID{},
			expectErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			repo.task[tt.id] = Task{ID: tt.id}

			svc := NewTaskService(repo)

			ctx := context.Background()
			_, err := svc.GetByID(ctx, tt.id)

			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect error: %v", err)
			}

		})
	}
}
