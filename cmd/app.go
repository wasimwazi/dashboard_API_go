package cmd

import (
	"dashboard-api/router"
	"database/sql"
	"log"
	"net/http"
)

// App : Struct to represent this app
type App struct {
	router.Router
}

// NewApp : Function to get App Struct
func NewApp(db *sql.DB) *App {
	return &App{
		Router: router.NewRouter(db),
	}
}

// Serve : Function to Run API Server
func (a *App) Serve(addr string) {
	router := a.Router.Setup()
	log.Println("App : Server is listening...")
	http.ListenAndServe(addr, router)
}
