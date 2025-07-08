package models

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type CommentRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}