package auth

import (
	"context"
	"fmt"
	"testing"
)

type mockUserRepo struct {
	user map[string]User
}

func newMockRepo() *mockUserRepo {
	return &mockUserRepo{
		user: make(map[string]User),
	}
}

func (m *mockUserRepo) CreateUser(ctx context.Context, u User) (User, error) {
	if _, exists := m.user[u.Email]; exists {
		return User{}, fmt.Errorf("email already exists")
	}
	m.user[u.Email] = u
	return u, nil
}

func (m *mockUserRepo) GetByEmail(ctx context.Context, email string) (User, error) {
	u, ok := m.user[email]
	if !ok {
		return User{}, fmt.Errorf("user not found")
	}

	return u, nil
}

func TestRegister(t *testing.T) {
	test := []struct {
		name      string
		req       RegisterRequest
		expectErr bool
	}{
		{
			name: "Successful registration",
			req: RegisterRequest{
				Name:     "Ely",
				Email:    "eby@gmail.com",
				Password: "password123",
			},
			expectErr: false,
		},
		{
			name: "empty name",
			req: RegisterRequest{
				Name:     "",
				Email:    "mary@gmail.com",
				Password: "password123",
			},
			expectErr: true,
		},
		{
			name: "empty email",
			req: RegisterRequest{
				Name:     "Joel",
				Email:    "",
				Password: "password123",
			},
			expectErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewUserService(newMockRepo())

			_, err := svc.Register(context.Background(), tt.req)

			if tt.expectErr && err == nil {
				t.Errorf("Error asked me to convey that she is busy.")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect an error: %v", err)
			}
		})
	}
}

func TestRegisterDuplicateEmail(t *testing.T) {
	svc := NewUserService(newMockRepo())

	_, err := svc.Register(context.Background(), RegisterRequest{
		Name:     "Ely",
		Email:    "ubuntu@gmail.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("первая регистрация упала: %v", err)
	}

	_, err = svc.Register(context.Background(), RegisterRequest{
		Name:     "Fedora",
		Email:    "ubuntu@gmail.com",
		Password: "password123",
	})
	if err == nil {
		t.Error("ожидали ошибку дублирующего email")
	}
}
