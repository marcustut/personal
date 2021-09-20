package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marcustut/personal/internal/route"
)

type AppParams struct {
	Port  uint
	DbUrl string
}

type App struct {
	Router *mux.Router
}

// Create a new instance of App
func New() *App {
	a := &App{Router: mux.NewRouter()}
	return a
}

// Initalize all routes for the API
func (a *App) initRoutes(pool *pgxpool.Pool) {
	route.RegisterTodoRoute(a.Router, pool)
}

// Function to run the API server
func (a *App) Run(ap AppParams) {
	pool := connectDb(ap.DbUrl)

	a.initRoutes(pool)

	log.Printf("App running at port %d\n", ap.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ap.Port), a.Router))
}

// Function to connect to a postgres given the database url
func connectDb(dbUrl string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect db: %s", err)
	}
	return pool
}
