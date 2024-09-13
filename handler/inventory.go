package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/RalphTan37/inventory-crud-api/model"
	"github.com/RalphTan37/inventory-crud-api/repository/inventory"
	"github.com/google/uuid"
)

type Inventory struct {
	Repo *inventory.RedisRepo
}

//each method will have the same HTTP Handler interface

// adds a new item to the inventory
func (i *Inventory) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ItemID         uuid.UUID  `json:"item_ID"`
		Name           string     `json:"name"`
		Category       string     `json:"category"`
		Quantity       int        `json:"quantity"`
		Price          float64    `json:"price"`
		Supplier       string     `json:"supplier"`
		Location       string     `json:"location"`
		Status         string     `json:"status"`
		ExpirationDate *time.Time `json:"expiration_date"`
	}

	//decode request body into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//generate new UUID
	if body.ItemID == uuid.Nil {
		body.ItemID = uuid.New()
	}

	now := time.Now().UTC() // get current time for CreatedAt & UpdatedAt

	//create new inventory item w/ data from the request body
	item := model.Inventory{
		ItemID:         body.ItemID,
		Name:           body.Name,
		Category:       body.Category,
		Quantity:       body.Quantity,
		Price:          body.Price,
		Supplier:       body.Supplier,
		Location:       body.Location,
		Status:         body.Status,
		ExpirationDate: body.ExpirationDate,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	//insert new item in repo
	err := i.Repo.Insert(r.Context(), item)
	if err != nil {
		fmt.Println("failed to insert item:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// marshal created item into JSON
	res, err := json.Marshal(item)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)                      //write JSON reponse
	w.WriteHeader(http.StatusCreated) //status code 201
}

// returns items in the inventory
func (i *Inventory) List(w http.ResponseWriter, r *http.Request) {
	//retrieve cursor from query params - if not provided, default to "0"
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	// convert cursor from string to uint64
	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//define page size
	const size = 50

	//fetch list of items from repo
	res, err := i.Repo.FindAll(r.Context(), inventory.FindAllPage{
		Offset: uint(cursor),
		Size:   size,
	})
	if err != nil {
		fmt.Println("failed to find all items:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// holds retrieved items & cursor for pagination
	var response struct {
		Items  []model.Inventory `json:"items"`
		Cursor uint64            `json:"cursor"`
	}

	// assign response items with fetched items
	response.Items = res.Items
	response.Cursor = res.Cursor

	//marshal response struct into JSON
	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data) //write JSON response - status code 200
}

// returns item by ID
func (i *Inventory) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id") //extract URL param from HTTP request

	//parse UUID from URL param
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//find item by id
	it, err := i.Repo.FindByID(r.Context(), itemID)
	if errors.Is(err, inventory.ErrDNE) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find item by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//encode & send item as JSON
	if err := json.NewEncoder(w).Encode(it); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// updates an existing inventory item
func (i *Inventory) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var body struct { //holds body request
		Status string `json:"status"`
	}

	//decode JSON request into body struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//extract URL param from HTTP request
	idParam := chi.URLParam(r, "id")

	// parse UUID from URL param
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// find existing item by ID
	theItem, err := i.Repo.FindByID(r.Context(), itemID)
	if errors.Is(err, inventory.ErrDNE) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find item by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// updates item's status
	theItem.Status = body.Status
	theItem.UpdatedAt = time.Now().UTC()

	// save updated item back to Go-Redis
	if err := i.Repo.Insert(r.Context(), theItem); err != nil {
		fmt.Println("failed to update item:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//respond w/ udpated item
	if err := json.NewEncoder(w).Encode(theItem); err != nil {
		fmt.Println("failed to marshal updated item:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// removes item in inventory by ID
func (i *Inventory) DeleteByID(w http.ResponseWriter, r *http.Request) {
	//extract URL param from HTTP request
	idParam := chi.URLParam(r, "id")

	//parse UUID from URL param
	itemID, err := uuid.Parse(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// delete item from repo
	err = i.Repo.DeleteByID(r.Context(), itemID)
	if errors.Is(err, inventory.ErrDNE) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find item by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
