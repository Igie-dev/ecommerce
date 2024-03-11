package model

import (
	"github.com/google/uuid"
)

type (
	Address struct {
		CustomerId  uuid.UUID `json:"customer_id"`
		PostalCode  string    `json:"postal_code"`
		City        string    `json:"city"`
		Country     string    `json:"country"`
		Street      string    `json:"street"`
		HouseNumber string    `json:"house_no"`
		Barangay    string    `json:"barangay"`
		Region      string    `json:"region"`
		Province    string    `json:"province"`
	}
 
CreateAddress struct {
     CustomerId uuid.UUID `json:"customer_id" validate:"required"`
    City string `json:"city" validate:"required"`
    PostalCode string `json:"postal_code" validate:"required"`
    Country string `json:"country" validate:"required"`
    Street string `json:"street" validate:"required"`
    HouseNumber string `json:"house_no" validate:"required"`
    Barangay string `json:"barangay" validate:"required"`
    Region string `json:"region" validate:"required"`
    Province string `json:"province validate:"required"`

  }
)
