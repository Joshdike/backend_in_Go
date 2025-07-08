package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/posts", handlers.GetPosts)
	r.Get("/posts/{postId}", handlers.GetPost)
	r.Post("/posts", handlers.CreatePost)
	r.Put("/posts/{postId}", handlers.UpdatePost)
	r.Delete("/posts/{postId}", handlers.DeletePost)
	r.Get("/posts/{postId}/comments", handlers.GetComments)
	r.Get("/posts/{postId}/comments/{commentId}", handlers.GetComment)
	r.Post("/posts/{postId}/comments", handlers.CreateComment)
	r.Put("/posts/{postId}/comments/{commentId}", handlers.UpdateComment)
	r.Delete("/posts/{postId}/comments/{commentId}", handlers.DeleteComment)

	port := ":8080"
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
