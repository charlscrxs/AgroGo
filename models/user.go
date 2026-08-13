package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleMayordomo UserRole = "MAYORDOMO"
	RoleAdmin     UserRole = "ADMIN"
)

type User struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Email         string    `json:"email" db:"email"`
	Password_hash string    `json:"-" db:"password_hash"`
	Role          UserRole  `json:"role" db:"role"`
	Is_active     bool      `json:"is_active" db:"is_active"`
	Created_at    time.Time `json:"created_at" db:"created_at"`
	Updated_at    time.Time `json:"updated_at" db:"updated_at"`
}

type UserController interface {
	GetById(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}
