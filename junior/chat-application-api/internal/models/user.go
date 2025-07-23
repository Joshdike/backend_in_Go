package models

import (
	"fmt"
	"time"

	"github.com/nyaruka/phonenumbers"
)

type User struct {
	Phone        PhoneNumber `json:"phone_number" db:"phone_number"`
	Username     string      `json:"name" db:"username"`
	PasswordHash string      `json:"-" db:"password_hash"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	LastSeen     time.Time   `json:"last_seen" db:"last_seen"`
}

type PasswordRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type PhoneNumber string

func (p PhoneNumber) Validate() error {
	parsed, err := phonenumbers.Parse(string(p), "US")
	if err != nil {
		return err
	}

	if !phonenumbers.IsValidNumber(parsed) {
		return fmt.Errorf("invalid phone number")
	}
	return nil
}
