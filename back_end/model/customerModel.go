package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Customer struct {
		ID         uint       `gorm:"primaryKey" json:"id"`
		CustomerId uuid.UUID  `json:"customer_id" `
		LastName   string     `json:"last_name" `
		FirstName  string     `json:"first_name" `
		Email      string     `json:"email"`
		Contact    string     `json:"contact"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  *time.Time `json:"updated_at"`
		Address    Address    `gorm:"null foreignKey:CustomerId" json:"address,omitempty"`
		Password   string     `json:"password"`
	}
	CreateCustomer struct {
		CustomerId uuid.UUID `json:"customer_id" validate:"required"`
		LastName   string    `json:"last_name" gorm:"not null; unique" validate:"required,min=1,max=100"  `
		FirstName  string    `json:"first_name" gorm:"not null; unique" validate:"required,min=1,max=100"  `
		Email      string    `json:"email" gorm:"->;<-:create; not null; unique" validate:"required,min=1,max=100"   `
		Contact    string    `json:"contact" gorm:"not null; unique" validate:"required,min=10,max=11"  `
		Password   string    `json:"password" validate:"required,lte=100,gte=10"`
	}

	UpdateCustomer struct {
		CustomerId uuid.UUID `json:"customer_id" validate:"required"`
		LastName   string    `json:"last_name" gorm:"not null; unique" validate:"required,min=1,max=100"  `
		FirstName  string    `json:"first_name" gorm:"not null; unique" validate:"required,min=1,max=100"  `
		Email      string    `json:"email" gorm:"->;<-:create; not null; unique" validate:"required,min=1,max=100"   `
		Contact    string    `json:"contact" gorm:"not null; unique" validate:"required,min=10,max=11"  `
	}
)
