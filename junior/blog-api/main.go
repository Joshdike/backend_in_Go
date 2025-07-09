package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/auth"
	"github.com/Joshdike/backend_in_Go/junior/blog-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	h := handlers.New(pool, os.Getenv("JWT_SECRET"))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Get("/posts", h.GetAllPosts)
	r.Get("/posts/{postId}", h.GetPostByID)
	r.Get("/posts/users/{userId}", h.GetPostsByUserId)
	r.Get("/posts/{postId}/comments", h.GetComments)
	r.Get("/posts/{postId}/comments/{commentId}", h.GetComment)

	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware(os.Getenv("JWT_SECRET")))

		r.Post("/posts", h.CreatePost)
		r.Put("/posts/{postId}", h.UpdatePost)
		r.Delete("/posts/{postId}", h.DeletePost)
		r.Post("/posts/{postId}/comments", h.CreateComment)
		r.Delete("/posts/{postId}/comments/{commentId}", h.DeleteComment)
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
