package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) SubscribeDB(ctx context.Context, userID int, name string) error {
	var employeeID int

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("SubscribeDB: err to start transaction: %w", err)
	}

	row := tx.QueryRowxContext(ctx, `SELECT id FROM employees WHERE name=$1`, name)

	err = row.Scan(&employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SubscribeDB: err to scan response: %w", err)
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO subscriptions (user_id, employee_id) VALUES ($1, $2)`,
		userID, employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SubscribeDB: err to insert subscription: %w", err)
	}

	return tx.Commit()
}

func (r *SubscriptionRepository) UnsubscribeDB(ctx context.Context, userID int, name string) error {
	var employeeID int

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("SubscribeDB: err to start transaction: %w", err)
	}

	row := tx.QueryRowxContext(ctx, `SELECT id FROM employees WHERE name=$1`, name)

	err = row.Scan(&employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SubscribeDB: err to scan response: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM subscriptions WHERE user_id=$1 AND employee_id=$2`, userID, employeeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("SubscribeDB: err to delete: %w", err)
	}

	return tx.Commit()
}
