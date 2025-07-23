package models

import "time"

type Conversation struct {
	ID           string    `json:"id" db:"id"`
	Type         string    `json:"type" db:"type"`
	Title        *string   `json:"title,omitempty" db:"title"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	MemberPhones []string  `json:"member_phones" db:"member_phones"`
}

type ConversationMember struct {
	ConversationID string      `json:"conversation_id" db:"conversation_id"`
	UserPhone      PhoneNumber `json:"user_phone" db:"user_phone"`
	Role           MemberRole  `json:"role" db:"role"`
	JoinedAt       time.Time   `json:"joined_at" db:"joined_at,default:now()"`
}

type MemberRole string

const (
	RoleAdmin  MemberRole = "admin"
	RoleMember MemberRole = "member"
)

func (m MemberRole) IsValid() bool {
	switch m {
	case RoleAdmin, RoleMember:
		return true
	default:
		return false
	}
}
