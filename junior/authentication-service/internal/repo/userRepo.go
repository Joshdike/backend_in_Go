package repo

import (
	"context"

	"github.com/Joshdike/backend_in_Go/junior/authentication_service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email, token string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	SetPasswordResetToken(ctx context.Context, email, token string) error
	FindUserByResetToken(ctx context.Context, token string) (*models.User, error)
}

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) UserRepository {
	return &userRepo{pool: pool}
}

func (u *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepo) FindUserByEmail(ctx context.Context, email, token string) (*models.User, error) {
	return nil, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepo) SetPasswordResetToken(ctx context.Context, email, token string) error {
	return nil
}

func (u *userRepo) FindUserByResetToken(ctx context.Context, token string) (*models.User, error) {
	return nil, nil
}