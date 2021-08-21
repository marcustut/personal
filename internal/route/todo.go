package route

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marcustut/personal/internal/handler"
	"github.com/marcustut/personal/internal/repo"
)

func RegisterTodoRoute(router *mux.Router) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect db: %s", err)
	}

	repo := repo.NewRepo(pool)

	th := handler.NewTodoHandler(repo)

	router.HandleFunc("/todos", th.List).Methods("GET")
	router.HandleFunc("/todos", th.Create).Methods("POST")
	router.HandleFunc("/todos/{id}", th.Get).Methods("GET")
	router.HandleFunc("/todos/{id}", nil).Methods("DELETE")
}
