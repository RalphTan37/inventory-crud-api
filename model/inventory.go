package model

import (
	"time" //Import Time Package

	"github.com/google/uuid" //Import Google UUID Package
)

// inventory properties
type Inventory struct {
	ItemID         uuid.UUID  `json:"item_ID"`
	Name           string     `json:"name"`
	Category       string     `json:"category"`
	Quantity       int        `json:"quantity"`
	Price          float64    `json:"price"`
	Supplier       string     `json:"supplier"`
	Location       string     `json:"location"`
	Status         string     `json:"status"`
	ExpirationDate *time.Time `json:"expiration_date"`
	CreatedAt      time.Time  `json:"created_At"`
	UpdatedAt      time.Time  `json:"updated_At"`
}
