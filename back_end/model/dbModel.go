package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Admin struct {
		ID        uint      `gorm:"primaryKey" json:"id"`
		LastName  string    `gorm:"not null; unique" json:"last_name" validate:"required,min=1,max=100"`
		FirstName string    `gorm:"not null; unique" json:"first_name" validate:"required,min=1,max=100"`
		Email     string    `gorm:"->;<-:create; not null; unique" json:"email" validate:"required,min=1,max=100"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Customer struct {
		ID         uint              `gorm:"primaryKey" json:"id"`
		CustomerId uuid.UUID         ` json:"customer_id" `
		LastName   string            `gorm:"not null; unique" validate:"required,min=1,max=100" json:"last_name" `
		FirstName  string            `gorm:"not null; unique" validate:"required,min=1,max=100" json:"first_name" `
		Email      string            `gorm:"->;<-:create; not null; unique" validate:"required,min=1,max=100"  json:"email" `
		Contact    string            `gorm:"not null; unique" validate:"required,min=10,max=11" json:"contact" `
		CreatedAt  time.Time         `json:"created_at"`
		UpdatedAt  time.Time         `json:"updated_at"`
		Addresses  []CustomerAddress `gorm:"null foreignKey:CustomerId" json:"addresses,omitempty"`
	}
	CustomerAddress struct {
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

	Product struct {
		ID         uint `gorm:"primaryKey" json:"id"`
		ProducName string
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)
