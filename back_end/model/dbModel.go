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
		ID         uint      `gorm:"primaryKey" json:"id"`
		CustomerId uuid.UUID ` json:"customer_id" `
		LastName   string    `gorm:"not null; unique" validate:"required,min=1,max=100" json:"last_name" `
		FirstName  string    `gorm:"not null; unique" validate:"required,min=1,max=100" json:"first_name" `
		Email      string    `gorm:"->;<-:create; not null;unique" validate:"required,min=1,max=100"  json:"email" `
		Address    string    `gorm:"not null; unique" validate:"required,min=1,max=200" json:"address" `
		Contact    string    `gorm:"not null; unique" validate:"required,min=10,max=11" json:"contact" `
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Product struct {
		ID         uint `gorm:"primaryKey" json:"id"`
		ProducName string
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)
