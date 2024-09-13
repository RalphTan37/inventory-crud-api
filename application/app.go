package application

import (
	"context"
	"fmt"
	"net/http" //functionality for creating HTTP clients & servers
)

// store any application dependencies
type App struct {
	router http.Handler
}

// application constructor
func New() *App { //returns a pointer to an instance of the application
	app := &App{
		router: loadRoutes(),
	}
	return app
}

// starting the application
func (a *App) Start(ctx context.Context) error {
	//instantiates an HTTP server
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	//initializes and starts HTTP server
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil // everything works correctly
}
