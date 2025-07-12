package repo

import "github.com/jackc/pgx/v5/pgxpool"

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type User struct {
	ID           int    `json:"id" db:"user_id"`
	Username     string `json:"name" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password_hash"`
}
