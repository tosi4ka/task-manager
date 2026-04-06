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
		t.Fatalf("the first registration dropped: %v", err)
	}

	_, err = svc.Register(context.Background(), RegisterRequest{
		Name:     "Fedora",
		Email:    "ubuntu@gmail.com",
		Password: "password123",
	})
	if err == nil {
		t.Error("expected a duplicate email error")
	}
}

func TestLogin(t *testing.T) {
	test := []struct {
		name      string
		req       LoginRequest
		expectErr bool
	}{
		{name: "Successful registration",
			req: LoginRequest{
				Email:    "eby@gmail.com",
				Password: "password123",
			},
			expectErr: false},
		{
			name: "Successful login",
			req: LoginRequest{
				Email:    "eby@gmail.com",
				Password: "password123",
			},
			expectErr: false,
		},
		{
			name: "empty email",
			req: LoginRequest{
				Email:    "",
				Password: "password123",
			},
			expectErr: true,
		},
		{
			name: "empty password",
			req: LoginRequest{
				Email:    "eby@gmail.com",
				Password: "",
			},
			expectErr: true,
		},
		{
			name: "wrong password",
			req: LoginRequest{
				Email:    "eby@gmail.com",
				Password: "wrongpassword",
			},
			expectErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewUserService(newMockRepo())

			svc.Register(context.Background(), RegisterRequest{
				Name:     "Abbey",
				Email:    "eby@gmail.com",
				Password: "password123",
			})

			_, err := svc.Login(context.Background(), tt.req)

			if tt.expectErr && err == nil {
				t.Errorf("Error asked me to convey that she is busy.")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("didn't expect an error: %v", err)
			}
		})
	}
}
