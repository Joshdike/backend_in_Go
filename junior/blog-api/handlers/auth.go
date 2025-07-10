package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Joshdike/backend_in_Go/junior/blog-api/models"
	"github.com/Joshdike/backend_in_Go/junior/blog-api/repo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type handle struct {
	db        *pgxpool.Pool
	userRepo  *repo.UserRepo
	jwtSecret string
}

func New(db *pgxpool.Pool, jwtSecret string) *handle {
	return &handle{
		db:        db,
		userRepo:  repo.NewUserRepo(db),
		jwtSecret: jwtSecret,
	}
}

func (h handle) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, `{"error":"username, email, and password are required"}`, http.StatusBadRequest)
		return
	}
	if len(req.Password) < 8 {
		http.Error(w, `{"error":"password must be at least 8 characters"}`, http.StatusBadRequest)
		return
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
	}
	if err := user.SetPassword(req.Password); err != nil {
		http.Error(w, `{"error":"failed to hash password"}`, http.StatusInternalServerError)
		return
	}
	if err := h.userRepo.Create(user); err != nil {
		http.Error(w, `{"error":"failed to create user"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	})
}

func (h handle) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if creds.Username == "" || creds.Password == "" {
		http.Error(w, `{"error":"username and password are required"}`, http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.FindByUsername(creds.Username)
	if err != nil {
		http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
		return
	}
	if !user.CheckPassword(creds.Password) {
		http.Error(w, `{"error":"invalid password"}`, http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		http.Error(w, `{"error":"Token generation failed"}`, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
