package auth

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      UserRepo
	jwtSecret string
}

func NewUserService(repo UserRepo, jwtSecret string) *UserService {
	return &UserService{repo: repo, jwtSecret: jwtSecret}
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
		return AuthResponse{}, fmt.Errorf("token error: %s", err)
	}

	token, err := GenerateToken(user.ID.String(), s.jwtSecret)

	return AuthResponse{
		AccessToken: token,
		User:        user,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req LoginRequest) (AuthResponse, error) {
	if req.Email == "" {
		return AuthResponse{}, fmt.Errorf("email is required")
	}
	if req.Password == "" {
		return AuthResponse{}, fmt.Errorf("password is required")
	}

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("email didn't exists")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("invalid password")
	}

	token, err := GenerateToken(user.ID.String(), s.jwtSecret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("token error: %s", err)
	}

	return AuthResponse{
		AccessToken: token,
		User:        user,
	}, nil
}
