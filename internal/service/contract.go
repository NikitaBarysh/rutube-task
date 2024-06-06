package service

import (
	"context"

	"rutube-task/internal/config"
	"rutube-task/internal/entity"
	"rutube-task/internal/repository"
)

type Service struct {
	AuthorizationService
	EmployeeServiceInterface
	SubscriptionServiceInterface
	AlertServiceInterface
}

type AuthorizationService interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	ValidateLogin(ctx context.Context, user entity.User) error
	CheckData(ctx context.Context, user entity.User) (int, error)
	GenerateJWTToken(userID int) (string, error)
	GetUserIDFromToken(tokenString string) int
	GeneratePasswordHash(pass string) string
}

type EmployeeServiceInterface interface {
	SetEmployeeList(ctx context.Context, employees []entity.Employee) error
	GetEmployee(ctx context.Context) ([]entity.Employee, error)
}

type SubscriptionServiceInterface interface {
	Subscribe(ctx context.Context, userID int, name string) error
	Unsubscribe(ctx context.Context, userID int, name string) error
}

type AlertServiceInterface interface {
	Scheduler()
}

func NewService(cfg *config.Config, rep *repository.Repository) *Service {
	return &Service{
		AuthorizationService:         NewAuthService(cfg, rep),
		EmployeeServiceInterface:     NewEmployeeService(rep),
		SubscriptionServiceInterface: NewSubscriptionService(rep),
		AlertServiceInterface:        NewAlert(cfg, rep),
	}
}
