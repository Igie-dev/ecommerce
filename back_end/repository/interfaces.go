package repository

import (
	"github.com/google/uuid"
	"github.com/laptop-shop.api/model"
)

type CustomerRepository interface {
	Create(b *model.CreateCustomer) error
	All(limit int, offset uint) ([]*model.Customer, error)
	Get(ID uuid.UUID) (*model.Customer, error)
	Exists(email string) (bool, error)
	Update(ID uuid.UUID, customer *model.UpdateCustomer) error
	Delete(ID uuid.UUID) error
}
