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

func (h handle) GetPosts(w http.ResponseWriter, r *http.Request) {

}

func (h handle) GetPost(w http.ResponseWriter, r *http.Request) {

}

func (h handle) CreatePost(w http.ResponseWriter, r *http.Request) {

}

func (h handle) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (h handle) DeletePost(w http.ResponseWriter, r *http.Request) {

}

func (h handle) GetComments(w http.ResponseWriter, r *http.Request) {

}

func (h handle) GetComment(w http.ResponseWriter, r *http.Request) {

}

func (h handle) CreateComment(w http.ResponseWriter, r *http.Request) {

}

func (h handle) DeleteComment(w http.ResponseWriter, r *http.Request) {

}
