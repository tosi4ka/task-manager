package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, req RegisterRequest) (AuthResponse, error) {
	if req.Name == "" {
		return AuthResponse{}, fmt.Errorf("name is required")
	}
	if req.Email == "" {
		return AuthResponse{}, fmt.Errorf("email is required")
	}
	if req.Password == "" {
		return AuthResponse{}, fmt.Errorf("password is required")
	}

	_, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return AuthResponse{}, fmt.Errorf("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("hash error: %s", err)
	}

	u := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("create user error: %s", err)
	}

	return AuthResponse{User: user}, nil
}
