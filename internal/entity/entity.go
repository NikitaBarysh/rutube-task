package entity

import "errors"

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" required:"true"`
	Password string `json:"password" required:"true"`
}

var (
	// ErrNotUniqueLogin - ошибка, если занят логин
	ErrNotUniqueLogin = errors.New("login is busy")
)
