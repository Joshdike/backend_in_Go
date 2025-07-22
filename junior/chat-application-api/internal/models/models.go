package models

import "time"

type User struct {
	ID           int       `json:"id" db:"user_id"`
	Username     string    `json:"name" db:"username"`
	PhoneNumber  string    `json:"phone_number" db:"phone_number"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Message struct {
	ID        int       `json:"id" db:"message_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Conversation struct {
	ID        int       `json:"id" db:"conversation_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Messages  []Message `json:"messages" db:"messages"`
}
