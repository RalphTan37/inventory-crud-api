package main

import (
	"fmt"      //format package
	"net/http" //functionality for creating HTTP clients & servers

	"github.com/go-chi/chi/v5"            //Go-Chi Package
	"github.com/go-chi/chi/v5/middleware" //Logging Middleware Package
)

func main() {
	router := chi.NewRouter()     //initialize router, constructor
	router.Use(middleware.Logger) //logs HTTP requests & responses

	router.Get("/inventory", basicHandler) //route for HTTP Get Request

	//instantiates an HTTP server
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	//initializes and starts HTTP server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to Listen to Server", err)
	}
}

// handler for HTTP server
func basicHandler(w http.ResponseWriter, r *http.Request) { //write and request parameters
	w.Write([]byte("Inventory Management System Project")) //cast to []byte as Write expects it
}
