package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joshdike/backend_in_Go/beginner/todo-list-api/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
)

type handle struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *handle {
	return &handle{db}
}

func (h handle) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query, params, err := sq.Select("*").From("todo_list").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)

	rows, err := h.db.Query(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	result := make([]models.Task, 0)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskId, &task.Task)
		if err != nil {
			http.Error(w, `{"error":"error retrieving tasks"}`, http.StatusInternalServerError)
			return
		}
		result = append(result, task)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func (h handle) GetTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"conversion error"}`, http.StatusBadRequest)
		return
	}

	query, params, err := sq.Select("*").From("todo_list").Where("task_id = ?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Println(query, params)
	err = h.db.QueryRow(r.Context(), query, params...).Scan(&task.TaskId, &task.Task)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error":"task not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"error retrieving data"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func (h handle) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, `{"error":"error decoding payload"}`, http.StatusBadRequest)
		return
	}
	query, params, err := sq.Insert("todo_list").Columns("task").Values(task.Task).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Printf("success: query: %s, params: %v\n", query, params)
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error inserting task"}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"successful"}`))
}
func (h handle) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"conversion error"}`, http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, `{"error":"error decoding payload"}`, http.StatusBadRequest)
		return
	}
	query, params, err := sq.Update("todo_list").Where("task_id = ?", id).Set("task", task.Task).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Printf("success: query: %s, params: %v\n", query, params)
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error updating task"}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"successful"}`))
}
func (h handle) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, `{"error":"conversion error"}`, http.StatusBadRequest)
		return
	}
	query, params, err := sq.Delete("todo_list").Where("task_id = ?", id).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"internal sql error"}`, http.StatusInternalServerError)
		return
	}
	fmt.Printf("success: query: %s, params: %v\n", query, params)
	_, err = h.db.Exec(r.Context(), query, params...)
	if err != nil {
		http.Error(w, `{"error":"error deleting task"}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message":"successful"}`))
}
