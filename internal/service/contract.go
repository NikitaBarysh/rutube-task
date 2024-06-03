package service

import (
	"context"

	"rutube-task/internal/config"
	"rutube-task/internal/entity"
	"rutube-task/internal/repository"
)

type Service struct {
	AuthorizationService
}

type AuthorizationService interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	ValidateLogin(ctx context.Context, user entity.User) error
	CheckData(ctx context.Context, user entity.User) (int, error)
	GenerateJWTToken(userID int) (string, error)
	GetUserIDFromToken(tokenString string) int
	GeneratePasswordHash(pass string) string
}

func NewService(cfg *config.Config, rep *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewService(cfg, rep),
	}
}
