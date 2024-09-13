package handler

import (
	"fmt"
	"net/http"
)

type Inventory struct{}

//each method will have the same HTTP Handler interface

// adds a new item to the inventory
func (i *Inventory) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add a new inventory item")
}

// returns items in the inventory
func (i *Inventory) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all items in inventory")
}

// returns item by ID
func (i *Inventory) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get item by ID")
}

// updates an existing inventory item
func (i *Inventory) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update item by ID")
}

// removes item in inventory by ID
func (i *Inventory) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete item by ID")
}
