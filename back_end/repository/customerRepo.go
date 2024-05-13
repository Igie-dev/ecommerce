package repository

import (
	"fmt"
	"time"

	"laptop-shop/database"
	"laptop-shop/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *database.DB
}

func NewCustomerRepo(db *database.DB) CustomerRepository {
	return &CustomerRepo{db}
}

func (repo *CustomerRepo) Create(b *model.CreateCustomer) error {
	u := uuid.New()
	//Create customer from data and add cutomer id by generating randum string
	customer := model.Customer{
		CustomerId: u,
	}
	err := repo.db.Create(&customer).Error
	return err
}

func (repo *CustomerRepo) All(limit int, offset uint) ([]*model.Customer, error) {
	var customers []*model.Customer
	query := `SELECT * FROM customers`
	var err error

	if limit > 0 && offset > 0 {
		query = fmt.Sprintf("%s LIMIT = ? OFFSET = ?", query)
		result := repo.db.Raw(query, limit, offset).Scan(&customers)
		err = result.Error
	} else {
		result := repo.db.Raw(query).Scan(&customers)
		err = result.Error

	}
	return customers, err
}

func (repo *CustomerRepo) Get(ID uuid.UUID) (*model.Customer, error) {
	customer := model.Customer{}
	query := `SELECT * FROM customer WHERE customer_id = ?`
	result := repo.db.Raw(query, ID).Scan(&customer)
	err := result.Error
	return &customer, err
}

func (repo *CustomerRepo) Update(ID uuid.UUID, c *model.UpdateCustomer) error {

	query := `UPDATE customer SET updated_at = ?, email = ?, first_name = ?, last_name = ?, contact = ? WHERE id = ?`
	result := repo.db.Exec(query, time.Now(), c.Email, c.FirstName, c.LastName, c.Contact, ID)
	return result.Error
}

func (repo *CustomerRepo) Delete(ID uuid.UUID) error {
	var result *gorm.DB
	query := `DELETE FROM customer WHERE customer_id = ? RETURNING id`
	queryDelAddress := `DELETE FROM address WHERE customer_id = ?`
	result = repo.db.Exec(query, ID)

	if result.Error != nil {
		return result.Error
	}

	result = repo.db.Exec(queryDelAddress, ID)

	return result.Error
}

func (repo *CustomerRepo) Exists(email string) (bool, error) {
	findByEmail := len(email)
	if findByEmail <= 0 {
		return false, nil
	}

	query := `SELECT customer_id FROM "customer" WHERE email = ?`

	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	result := repo.db.Raw(query, email).Scan(&exists)

	if result.Error != nil {
		return false, nil
	}

	return exists, nil
}
