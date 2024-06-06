package repository

import (
	"context"

	"rutube-task/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	AuthorizationRepository
	EmployeeRepositoryInterface
	SubscriptionRepositoryInterface
}

type AuthorizationRepository interface {
	SetUserDB(ctx context.Context, user entity.User) (int, error)
	GetUserFromDB(ctx context.Context, user entity.User) (int, error)
	Validate(ctx context.Context, username string) error
}

type EmployeeRepositoryInterface interface {
	SetEmployeeListDB(ctx context.Context, employeeInfo entity.Employee) error
	GetEmployeeListDB(ctx context.Context) ([]entity.Employee, error)
	GetEmployeeIDFromDB(ctx context.Context, name string) (int, error)
	GetEmployeeBirthdayDB() ([]entity.Employee, error)
}

type SubscriptionRepositoryInterface interface {
	SubscribeDB(ctx context.Context, userID int, name string) error
	UnsubscribeDB(ctx context.Context, userID int, name string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository:         NewAuthRepository(db),
		EmployeeRepositoryInterface:     NewEmployeeRepository(db),
		SubscriptionRepositoryInterface: NewSubscriptionRepository(db),
	}
}
