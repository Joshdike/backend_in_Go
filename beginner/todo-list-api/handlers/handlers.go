package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type handle struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *handle {
	return &handle{db}
}

func (h handle) GetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h handle) GetTaskById(w http.ResponseWriter, r *http.Request) {

}

func (h handle) CreateTask(w http.ResponseWriter, r *http.Request) {

}
func (h handle) UpdateTask(w http.ResponseWriter, r *http.Request) {

}
func (h handle) DeleteTask(w http.ResponseWriter, r *http.Request) {

}
