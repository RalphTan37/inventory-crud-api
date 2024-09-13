package application

import (
	"context"
	"fmt"
	"net/http" //functionality for creating HTTP clients & servers
	"time"

	"github.com/redis/go-redis/v9" //Import Go-Redis Client
)

// store any application dependencies
type App struct {
	router http.Handler
	rdb    *redis.Client //stores Redis Client
	config Config
}

// application constructor
func New(config Config) *App { //returns a pointer to an instance of the application and an error
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}), //creates new instance of Redis Client
		config: config,
	}

	app.loadRoutes()

	return app
}

// starting the application
func (a *App) Start(ctx context.Context) error {
	//instantiates an HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	//calls the Ping Method of the Redis Client - returns an err if it it not able to connect
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	//handles the multiple return pts
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("starting server")

	ch := make(chan error, 1) // initialize channel w/ error type & buffer sz 1

	/*
		goroutine - runs server concurrently:
		starts a new anonymous function in a new thread of execution & ensures that it doesn't block the main thread
	*/
	go func() {
		//initializes and starts HTTP server
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err) //publishes a value to the channel
		}
		close(ch) //closes channel when the function is done
	}()

	//receiver for channel - blocks the code execution until it either receives a value or the channel is closed
	select {
	case err = <-ch: //capture any value sent on this channel
		return err
	case <-ctx.Done(): //returns channel inside
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil // everything works correctly
}
