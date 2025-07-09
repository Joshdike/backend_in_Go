package repo

import (
	"context"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING user_id, created_at, updated_at
	`
	return r.db.QueryRow(context.Background(), query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

}

func (r *UserRepo) FindByUsername(username string) (*models.User, error) {
	query := `
		SELECT user_id, username, email, password_hash, created_at, updated_at
		FROM users
		WHERE username = $1
	`
	var user models.User
	err := r.db.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
