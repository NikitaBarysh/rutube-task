package repository

import (
	"context"

	"rutube-task/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	AuthorizationRepository
}

type AuthorizationRepository interface {
	SetUserDB(ctx context.Context, user entity.User) (int, error)
	GetUserFromDB(ctx context.Context, user entity.User) (int, error)
	Validate(ctx context.Context, username string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthRepository(db),
	}
}
