package app

import (
	"database/sql"
	"net/http"
	"sync"
)

var app_instance *Application
var app_once sync.Once

type Application struct {
	DB     *sql.DB
	Router *Router
}

func GetInstance() *Application {
	app_once.Do(func() {
		app_instance = newApplication()
	})
	return app_instance
}

func newApplication() *Application {
	// Create the router
	router := CreateRouter()

	// Serve static files from the "frontend" directory
	router.Mux.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("dev/frontend"))))

	// Return the application instance
	return &Application{
		// DB:     GetInstance(),
		Router: &router,
	}
}
