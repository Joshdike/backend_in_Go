package models

import "time"

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID   int    `json:"id"`
	PostID int    `json:"post_id"`
	Name string `json:"name"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CommentRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}