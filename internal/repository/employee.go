package repository

import (
	"context"
	"fmt"
	"time"

	"rutube-task/internal/entity"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) SetEmployeeListDB(ctx context.Context, employeeInfo entity.Employee) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO employees (name, birthday) VALUES ($1, $2)`,
		employeeInfo.Name, employeeInfo.Birthday)

	if err != nil {
		return fmt.Errorf("SetEmployeeListDB: err to exec query: %w", err)
	}

	return nil
}

func (r *EmployeeRepository) GetEmployeeListDB(ctx context.Context) ([]entity.Employee, error) {
	employees := make([]entity.Employee, 0)

	rows, err := r.db.QueryContext(ctx, `SELECT * FROM employees`)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeeListDB: err to exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee entity.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Birthday)
		if err != nil {
			return nil, fmt.Errorf("GetEmployeeListDB: err to scan row: %w", err)
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *EmployeeRepository) GetEmployeeIDFromDB(ctx context.Context, name string) (int, error) {
	var id int

	row := r.db.QueryRowxContext(ctx, `SELECT id FROM employees WHERE name=$1`, name)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("GetEmployeeIDFromDB: err to scan response: %w", err)
	}

	return id, nil
}

func (r *EmployeeRepository) GetEmployeeBirthdayDB() ([]entity.Employee, error) {
	employeeList := make([]entity.Employee, 0)

	today := time.Now().Format("01-02")

	rows, err := r.db.Queryx(`SELECT id, name FROM employees WHERE TO_CHAR(birthday, 'MM-DD') = $1`, today)
	if err != nil {
		return nil, fmt.Errorf("GetEmployeeBirthdayDB: err to exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee entity.Employee
		err := rows.Scan(&employee.ID, &employee.Name)
		if err != nil {
			return nil, fmt.Errorf("GetEmployeeBirthdayDB: err to scan row: %w", err)
		}
		employeeList = append(employeeList, employee)
	}

	return employeeList, nil
}
