package repository

import (
	"context"
	"my-todo/internal/domain/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user_register *models.User) error
	// GetByID(ctx context.Context, id int) (*models.User, error)
	// Update(ctx context.Context, user_update *models.User ) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool)UserRepository{
	return &userRepository{db: db}
}

func (ur *userRepository) Create(ctx context.Context, user_register *models.User) error{
	err := ur.db.QueryRow(ctx, "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", user_register.Username, user_register.Email, user_register.PasswordHash).Scan(&user_register.ID)
	return err
}																			