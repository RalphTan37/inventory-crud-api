package application

import (
	"net/http"

	"github.com/RalphTan37/inventory-crud-api/handler" //Import Handler Package
	"github.com/go-chi/chi/v5"                         //Go-Chi Package
	"github.com/go-chi/chi/v5/middleware"              //Logging Middleware Package
)

func loadRoutes() *chi.Mux {
	//new instance of a Go-Chi router
	router := chi.NewRouter()

	router.Use(middleware.Logger) //logs HTTP requests & responses

	//HTTP handler for / path
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/inventory", loadInventoryRoutes) //sub router for the /inventory path

	return router
}

// loads inventory routes
func loadInventoryRoutes(router chi.Router) {
	//creates instance of inventory handler
	inventoryHandler := &handler.Inventory{}

	//HTTP Methods for CRUD Methods
	router.Post("/", inventoryHandler.Create)
	router.Get("/", inventoryHandler.List)
	router.Get("/{id}", inventoryHandler.GetByID)
	router.Put("/{id}", inventoryHandler.UpdateByID)
	router.Delete("/{id}", inventoryHandler.DeleteByID)
}
