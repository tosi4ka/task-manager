package auth

import (
	"context"

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
		return AuthResponse{}, ErrNameRequired
	}
	if req.Email == "" {
		return AuthResponse{}, ErrEmailRequired
	}
	if req.Password == "" {
		return AuthResponse{}, ErrPasswordRequired
	}

	_, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil {
		return AuthResponse{}, ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, ErrInternal
	}

	u := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return AuthResponse{}, ErrInternal
	}

	token, err := GenerateToken(user.ID.String(), s.jwtSecret)

	return AuthResponse{
		AccessToken: token,
		User:        user,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req LoginRequest) (AuthResponse, error) {
	if req.Email == "" {
		return AuthResponse{}, ErrEmailRequired
	}
	if req.Password == "" {
		return AuthResponse{}, ErrPasswordRequired
	}

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return AuthResponse{}, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return AuthResponse{}, ErrInvalidCredentials
	}

	token, err := GenerateToken(user.ID.String(), s.jwtSecret)
	if err != nil {
		return AuthResponse{}, ErrInternal
	}

	return AuthResponse{
		AccessToken: token,
		User:        user,
	}, nil
}
