package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcustut/personal/internal/route"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{Router: mux.NewRouter()}
	return a
}

func (a *App) initRoutes() {
	route.RegisterTodoRoute(a.Router)
}

func (a *App) Run(port uint) {
	a.initRoutes()

	log.Printf("App running at port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}
