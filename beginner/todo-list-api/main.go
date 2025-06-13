package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Joshdike/backend_in_Go/beginner/todo-list-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	h := handlers.New(pool)

	r.Get("/tasks", h.GetAllTasks)
	r.Get("/task/{id}", h.GetTaskById)
	r.Post("/task", h.CreateTask)
	r.Put("/task/{id}", h.UpdateTask)
	r.Delete("/task/{id}", h.DeleteTask)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

}
