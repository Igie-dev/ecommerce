package repository

import (
	"laptop-shop-api/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(b *model.CreateUser) error
	All(limit int, offset uint) ([]*model.User, error)
	Get(ID uuid.UUID) (*model.User, error)
	Exists(email string) (bool, error)
	Update(ID uuid.UUID, user *model.UpdateUser) error
	Delete(ID uuid.UUID) error
}
