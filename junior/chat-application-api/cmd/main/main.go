package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	mw "github.com/Joshdike/backend_in_Go/junior/chat-application-api/internal/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	// h := handlers.New(pool, os.Getenv("JWT_SECRET"))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes

	// r.Post("/register", h.Register)
	// r.Post("/login", h.Login)

	r.Group(func(r chi.Router) {
		r.Use(mw.JWTMiddleware(os.Getenv("JWT_SECRET")))
		// private routes
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
