package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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
	h := handlers.New(pool)
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/posts", h.GetPosts)
	r.Get("/posts/{postId}", h.GetPost)
	r.Post("/posts", h.CreatePost)
	r.Put("/posts/{postId}", h.UpdatePost)
	r.Delete("/posts/{postId}", h.DeletePost)
	r.Get("/posts/{postId}/comments", h.GetComments)
	r.Get("/posts/{postId}/comments/{commentId}", h.GetComment)
	r.Post("/posts/{postId}/comments", h.CreateComment)
	r.Delete("/posts/{postId}/comments/{commentId}", h.DeleteComment)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
