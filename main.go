package main

import (
	"context"
	"fmt"
	"os"

	"os/signal" //access to incoming signals

	"github.com/RalphTan37/inventory-crud-api/application"
)

func main() {
	//new instance of the application
	app := application.New(application.LoadConfig())
	//takes in a signal and a ctx, return another ctx that will be notified if the signal is created
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel() //called at the end of the function

	//start the application
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
}
