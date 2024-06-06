package service

import (
	"context"
	"fmt"

	"rutube-task/internal/entity"
	"rutube-task/internal/repository"
)

type EmployeeService struct {
	rep *repository.Repository
}

func NewEmployeeService(rep *repository.Repository) *EmployeeService {
	return &EmployeeService{rep: rep}
}

func (s *EmployeeService) SetEmployeeList(ctx context.Context, employees []entity.Employee) error {
	for _, employee := range employees {
		err := s.rep.EmployeeRepositoryInterface.SetEmployeeListDB(ctx, employee)
		if err != nil {
			return fmt.Errorf("SetEmployeeList: %w ", err)
		}
	}
	return nil
}

func (s *EmployeeService) GetEmployee(ctx context.Context) ([]entity.Employee, error) {
	list, err := s.rep.GetEmployeeListDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeeList: %w", err)
	}
	return list, nil
}
