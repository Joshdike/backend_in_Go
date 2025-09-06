package models

import "time"

type User struct {
	ID                  int        `json:"id" db:"user_id"`
	Username            string     `json:"name" db:"username"`
	Email               string     `json:"email" db:"email"`
	PasswordHash        string     `json:"-" db:"password_hash"`
	PasswordResetToken  string     `json:"-" db:"password_reset_token"`
	PasswordResetExpiry *time.Time `json:"-" db:"password_reset_expiry"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PasswordRecoveryRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type PasswordResetRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Token    string `json:"token" validate:"required"`
}

type AuthResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
