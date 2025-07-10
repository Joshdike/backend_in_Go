package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi"
)

func (h handle) GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user_id = r.URL.Query().Get("user_id")
	if user_id == "" {
		http.Error(w, `{"error":"user_id is required"}`, http.StatusBadRequest)
		return
	}
	author_id, err := strconv.Atoi(user_id)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}
	query, params, err := sq.Select("*").From("posts").Where("author_id = ?", author_id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	rows, err := h.db.Query(r.Context(), query, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error":"no post found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	result := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			http.Error(w, `{"error":"error retrieving posts"}`, http.StatusInternalServerError)
			return
		}
		result = append(result, post)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}

}

func (h handle) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query, params, err := sq.Select("*").From("posts").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	rows, err := h.db.Query(r.Context(), query, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error":"no post found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	result := make([]models.Post, 0)
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			http.Error(w, `{"error": "error retrieving posts"}`, http.StatusInternalServerError)
			return
		}
		result = append(result, post)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func (h handle) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		http.Error(w, `{"error":"invalid post ID"}`, http.StatusBadRequest)
		return
	}

	query, params, err := sq.Select("*").From("posts").Where("post_id=?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	var post models.Post
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func (h handle) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.Context().Value("userID").(string))
	if err != nil {
		http.Error(w, `{"error":"invalid user ID"}`, http.StatusUnauthorized)
		return
	}
	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}

	query, params, err := sq.Insert("posts").Columns("title", "content", "author_id").Values(post.Title, post.Content, post.AuthorID).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error creating post"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"successful"}`))
}

func (h handle) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		http.Error(w, `{"error":"invalid post ID"}`, http.StatusBadRequest)
		return
	}
	var req models.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.Context().Value("userID").(string))
	if err != nil {
		http.Error(w, `{"error":"invalid user ID"}`, http.StatusUnauthorized)
		return
	}
	var existingAuthorID int
	query, params, err := sq.Select("author_id").From("posts").Where("post_id=?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&existingAuthorID)
	if err != nil {
		http.Error(w, `{"error":"post not found"}`, http.StatusNotFound)
		return
	}
	if existingAuthorID != userID {
		http.Error(w, `{"error":"unauthorized to update post"}`, http.StatusForbidden)
		return
	}
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	query, params, err = sq.Update("posts").Set("title", post.Title).Set("content", post.Content).Where("post_id=?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error updating post"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"successful"}`))

}

func (h handle) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		http.Error(w, `{"error":"invalid post ID"}`, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.Context().Value("userID").(string))
	if err != nil {
		http.Error(w, `{"error":"invalid user ID"}`, http.StatusUnauthorized)
		return
	}
	var existingAuthorID int
	query, params, err := sq.Select("author_id").From("posts").Where("post_id=?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&existingAuthorID)
	if err != nil {
		http.Error(w, `{"error":"post not found"}`, http.StatusNotFound)
		return
	}
	if existingAuthorID != userID {
		http.Error(w, `{"error":"unauthorized to delete post"}`, http.StatusForbidden)
		return
	}
	query, params, err = sq.Delete("posts").Where("post_id=?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error deleting post"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
