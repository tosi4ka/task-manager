package auth

import (
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepo interface {
	CreateUser(ctx context.Context, u User) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}

func (r *UserRepository) CreateUser(ctx context.Context, u User) (User, error) {
	err := r.db.QueryRowContext(ctx, "INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id, name, email, created_at, updated_at", u.Name, u.Password, u.Email).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return User{}, fmt.Errorf("User create error^ %s", err)
	}

	return u, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password,created_at, updated_at FROM users WHERE email = $1", email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return User{}, fmt.Errorf("get user error: %s", err)
	}

	return u, nil
}
