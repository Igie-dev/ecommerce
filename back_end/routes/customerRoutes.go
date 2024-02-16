package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/laptop-shop.api/model"
	"github.com/laptop-shop.api/utils"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CustomerRoute(app *fiber.App) {
	api := app.Group("/customer")
	MigrateCustomer(r.DB)
	api.Get("/", r.GetAllCustomer)
	api.Post("/", utils.ValidateCustomer, r.CreateNewCustomer)
}

//Migrate

func MigrateCustomer(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Customer{})
	return err
}

// Create
func (r *Repository) CreateNewCustomer(c *fiber.Ctx) error {
	var mr *utils.MalformedRequest

	u := uuid.New()
	//Create user from data and add cutomer id by generating randum string
	customer := model.Customer{
		CustomerId: u,
	}

	err := utils.DecodeJSONBody(c, &customer)

	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(fiber.Map{"message": mr.Msg})
		}
	}

	err = r.DB.Create(&customer).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed to create customer"})
		return nil
	}
	return c.Status(http.StatusCreated).JSON(&fiber.Map{"message": "Customer created successfully"})
}

// Read
func (r *Repository) GetAllCustomer(c *fiber.Ctx) error {

	customers := &[]model.Customer{}

	err := r.DB.Find(&customers).Error

	if err != nil {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "No customer found"})
		return nil
	}

	if len(*customers) <= 0 {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "No customer found"})
		return nil
	}

	return c.Status(200).JSON(fiber.Map{
		"data": customers,
	})
}
