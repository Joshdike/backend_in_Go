package handlers

import "github.com/jackc/pgx/v5/pgxpool"

type handle struct {
}

func New(pool *pgxpool.Pool, jwtSecret string) *handle {
	return &handle{}
}
