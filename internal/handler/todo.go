package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcustut/personal/internal/repo"
)

type TodoHandler struct {
	repo *repo.Repo
}

func NewTodoHandler(repo *repo.Repo) *TodoHandler {
	return &TodoHandler{
		repo: repo,
	}
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.ListTodos(context.Background())
	if err != nil {
		log.Printf("Error fetching todos %s\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error converting id to int %s\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	todo, err := h.repo.GetTodo(context.Background(), int32(id))
	if err != nil {
		log.Printf("Error getting todo with id %d: %s", id, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create is called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create"))
}
