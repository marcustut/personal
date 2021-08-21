package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var (
	listTodoRe   = regexp.MustCompile(`^\/todos[\/]*$`)
	getTodoRe    = regexp.MustCompile(`^\/todos\/(d+)$`)
	createTodoRe = regexp.MustCompile(`^\/todos[\/]*$`)
)

type todo struct {
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type datastore struct {
	m map[string]todo
	*sync.RWMutex
}

type todoHandler struct {
	store *datastore
}

func (h *todoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listTodoRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getTodoRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createTodoRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *todoHandler) List(w http.ResponseWriter, r *http.Request) {
	// Read todos
	h.store.RLock()
	todos := make([]todo, 0, len(h.store.m))
	for _, v := range h.store.m {
		todos = append(todos, v)
	}
	h.store.RUnlock()

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *todoHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get\n")
	fmt.Println(r.URL.Path)
}

func (h *todoHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create\n")
	fmt.Println(r.URL.Path)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	th := &todoHandler{
		store: &datastore{
			m: map[string]todo{
				"1": {
					Id:          1,
					Title:       "Test Todo",
					Description: "Test Todo's Description",
					Completed:   false,
				},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux := http.NewServeMux()
	mux.Handle("/todos", th)
	mux.Handle("/todos/", th)
	http.ListenAndServe(":8080", mux)
}
