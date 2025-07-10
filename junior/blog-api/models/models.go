package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `json:"id" db:"user_id"`
	Username     string    `json:"name" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

type Post struct {
	ID        int       `json:"id" db:"post_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	AuthorID  int       `json:"author_id" db:"author_id"`
}

type Comment struct {
	ID        int       `json:"id" db:"comment_id"`
	PostID    int       `json:"post_id" db:"post_id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PostRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required,min=10"`
}

type CommentRequest struct {
	Body     string `json:"body"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
