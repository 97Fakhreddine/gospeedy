package core

import (
	"fmt"
	"log"
	"net/http"
)

// App struct contains all components to run the application
type App struct {
	Server *http.Server
}

// NewApp initializes a new app instance
func NewApp(port string) *App {
	return &App{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: nil, // Add your router here
		},
	}
}

// Run starts the application server
func (a *App) Run() {
	log.Printf("App is running on port %s", a.Server.Addr)
	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}
