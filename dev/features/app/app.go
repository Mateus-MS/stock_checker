package app

import (
	"database/sql"
	"net/http"
	"sync"
)

var instance *Application
var once sync.Once

type Application struct {
	DB     *sql.DB
	Router *Router
}

func GetInstance() *Application {
	once.Do(func() {
		instance = newApplication()
	})
	return instance
}

func newApplication() *Application {
	// Create the router
	router := CreateRouter()

	// Serve static files from the "frontend" directory
	router.Mux.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("dev/frontend"))))

	// Return the application instance
	return &Application{
		DB:     StartDBConnection(),
		Router: &router,
	}
}
