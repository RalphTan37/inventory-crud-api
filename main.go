package main

import (
	"fmt"      //format package
	"net/http" //functioanality for creating HTTP clients & servers
)

func main() {
	//instantiates an HTTP server
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	//initializes and starts HTTP server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to Listen to Server", err)
	}
}

// handler for HTTP server
func basicHandler(w http.ResponseWriter, r *http.Request) { //write and request parameters
	w.Write([]byte("Inventory System Project")) //cast to []byte as Write expects it
}
