package repository

import (
	"fmt"
	"time"

	"laptop-shop-api/database"
	"laptop-shop-api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(b *model.CreateUser) error {
	u := uuid.New()
	//Create user from data and add cutomer id by generating randum string
	user := model.User{
		UserId: u,
	}
	err := repo.db.Create(&user).Error
	return err
}

func (repo *UserRepo) All(limit int, offset uint) ([]*model.User, error) {
	var users []*model.User
	query := `SELECT * FROM users`
	var err error

	if limit > 0 && offset > 0 {
		query = fmt.Sprintf("%s LIMIT = ? OFFSET = ?", query)
		result := repo.db.Raw(query, limit, offset).Scan(&users)
		err = result.Error
	} else {
		result := repo.db.Raw(query).Scan(&users)
		err = result.Error

	}
	return users, err
}

func (repo *UserRepo) Get(ID uuid.UUID) (*model.User, error) {
	user := model.User{}
	query := `SELECT * FROM user WHERE user_id = ?`
	result := repo.db.Raw(query, ID).Scan(&user)
	err := result.Error
	return &user, err
}

func (repo *UserRepo) Update(ID uuid.UUID, c *model.UpdateUser) error {

	query := `UPDATE user SET updated_at = ?, email = ?, first_name = ?, last_name = ?, contact = ? WHERE id = ?`
	result := repo.db.Exec(query, time.Now(), c.Email, c.FirstName, c.LastName, c.Contact, ID)
	return result.Error
}

func (repo *UserRepo) Delete(ID uuid.UUID) error {
	var result *gorm.DB
	query := `DELETE FROM user WHERE user_id = ? RETURNING id`
	queryDelAddress := `DELETE FROM address WHERE user_id = ?`
	result = repo.db.Exec(query, ID)

	if result.Error != nil {
		return result.Error
	}

	result = repo.db.Exec(queryDelAddress, ID)

	return result.Error
}

func (repo *UserRepo) Exists(email string) (bool, error) {
	findByEmail := len(email)
	if findByEmail <= 0 {
		return false, nil
	}

	query := `SELECT user_id FROM "user" WHERE email = ?`

	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	result := repo.db.Raw(query, email).Scan(&exists)

	if result.Error != nil {
		return false, nil
	}

	return exists, nil
}
