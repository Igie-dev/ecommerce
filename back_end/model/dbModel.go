package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"enail"`
	Address   string    `json:"address"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	ProducName string
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
