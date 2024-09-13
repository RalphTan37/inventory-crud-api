package application

import (
	"context"
	"fmt"
	"net/http" //functionality for creating HTTP clients & servers

	"github.com/redis/go-redis/v9" //Import Go-Redis Client
)

// store any application dependencies
type App struct {
	router http.Handler
	rdb    *redis.Client //stores Redis Client
}

// application constructor
func New() *App { //returns a pointer to an instance of the application and an error
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}), //creates new instance of Redis Client
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

	//calls the Ping Method of the Redis Client - returns an err if it it not able to connect
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	fmt.Println("starting server")

	//initializes and starts HTTP server
	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil // everything works correctly
}
