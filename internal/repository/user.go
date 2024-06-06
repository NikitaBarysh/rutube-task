package repository

import (
	"context"
	"database/sql"
	"fmt"

	"rutube-task/internal/entity"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	rep *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		rep: db,
	}
}

func (r *AuthRepository) SetUserDB(ctx context.Context, user entity.User) (int, error) {
	var id int

	tx, err := r.rep.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("SetUserDB: err to begin transaction: %w", err)
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO users (login, password) VALUES ($1,$2)`,
		user.Login, user.Password)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return 0, fmt.Errorf("SetUserDB: err to do Rollback: %w", errRollback)
		}
		return 0, fmt.Errorf("SetUserDB: err to do exec in DB: %w", err)
	}

	row := tx.QueryRowxContext(ctx, "SELECT id FROM  users WHERE login=$1", user.Login)
	if row.Err() != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return 0, fmt.Errorf("SetUserDB: err to do Rollback: %w", errRollback)
		}
		return 0, fmt.Errorf("SetUserDB: err to get id: %w", row.Err())
	}

	if err = row.Scan(&id); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return 0, fmt.Errorf("SetUserDB: err to do Rollback: %w", errRollback)
		}
		return 0, fmt.Errorf("SetUserDB: err to scan id: %w", err)
	}

	return id, tx.Commit()
}

// GetUserFromDB - получаем id пользователя из базы
func (r *AuthRepository) GetUserFromDB(ctx context.Context, user entity.User) (int, error) {
	var id int

	row := r.rep.QueryRowxContext(ctx, `SELECT id FROM users WHERE login=$1 AND password=$2`, user.Login, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("GetUserFromDB: err to scan id: %w", err)
	}

	return id, nil

}

// Validate - проверяем на наличие логина в базе
func (r *AuthRepository) Validate(ctx context.Context, username string) error {
	var id int

	row := r.rep.QueryRowxContext(ctx, "SELECT id FROM users WHERE login=$1", username)
	if row.Err() != nil {
		return fmt.Errorf("Validate: err to get id: %w ", row.Err())
	}

	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("Validate: err to scan id: %w ", err)
	}

	return fmt.Errorf("Validate: err to get id: %w ", row.Err())
}
