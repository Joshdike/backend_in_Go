package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/models"
	"github.com/go-chi/chi"

	sq "github.com/Masterminds/squirrel"
)

func (h handle) GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post_id, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		http.Error(w, `{"error":"invalid post ID"}`, http.StatusBadRequest)
		return
	}
	query, params, err := sq.Select("*").From("comments").Where("post_id=?", post_id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	rows, err := h.db.Query(r.Context(), query, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error":"no comment found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"error retrieving comments"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	result := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.Body, &comment.CreatedAt)
		if err != nil {
			http.Error(w, `{"error": "error retrieving comments"}`, http.StatusInternalServerError)
			return
		}
		result = append(result, comment)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}

}

func (h handle) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	comment_id, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		http.Error(w, `{"error":"invalid comment ID"}`, http.StatusBadRequest)
		return
	}

	query, params, err := sq.Select("*").From("comments").Where("comment_id=?", comment_id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	var comment models.Comment
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.Body, &comment.CreatedAt)
	if err != nil {
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}

}

func (h handle) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.Context().Value("userID").(string))
	if err != nil {
		http.Error(w, `{"error":"invalid user ID"}`, http.StatusUnauthorized)
		return
	}

	post_id, err := strconv.Atoi(chi.URLParam(r, "postId"))
	if err != nil {
		http.Error(w, `{"error":"invalid post ID"}`, http.StatusBadRequest)
		return
	}

	var postExists bool
	err = h.db.QueryRow(r.Context(), "SELECT EXISTS(SELECT 1 FROM posts WHERE post_id = $1)", post_id).Scan(&postExists)
	if err != nil || !postExists {
		http.Error(w, `{"error":"post not found"}`, http.StatusNotFound)
		return
	}
	comment := models.Comment{
		PostID:   post_id,
		AuthorID: userID,
		Body:     req.Body,
	}

	query, params, err := sq.Insert("comments").Columns("post_id", "author_id", "body").Values(comment.PostID, comment.AuthorID, comment.Body).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error creating comment"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"comment created successfully"}`))
}

func (h handle) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	comment_id, err := strconv.Atoi(chi.URLParam(r, "commentId"))
	if err != nil {
		http.Error(w, `{"error":"invalid comment ID"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(r.Context().Value("userID").(string))
	if err != nil {
		http.Error(w, `{"error":"invalid user ID"}`, http.StatusUnauthorized)
		return
	}
	var existingAuthorID int
	query, params, err := sq.Select("author_id").From("comments").Where("comment_id=?", comment_id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&existingAuthorID)
	if err != nil {
		http.Error(w, `{"error":"comment not found"}`, http.StatusNotFound)
		return
	}
	if existingAuthorID != userID {
		http.Error(w, `{"error":"unauthorized to delete comment"}`, http.StatusForbidden)
		return
	}
	query, params, err = sq.Delete("comments").Where("comment_id=?", comment_id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error deleting comment"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
