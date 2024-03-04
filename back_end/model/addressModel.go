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
)
