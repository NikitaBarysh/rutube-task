package entity

import (
	"errors"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" required:"true"`
	Password string `json:"password" required:"true"`
}

type SubscriptionRequest struct {
	Name string `json:"name"`
}

type Employee struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" required:"true"`
	Birthday time.Time `json:"birthday" required:"true"`
}

type Subscription struct {
	UserID     int `json:"user_id"`
	EmployeeID int `json:"employee_id"`
}

var (
	ErrNotUniqueLogin = errors.New("login is busy")
)
