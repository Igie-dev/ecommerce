package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/laptop-shop.api/model"
	"github.com/laptop-shop.api/utils"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func (r *CustomerRepository) CustomerRoutes(app *fiber.App) {

	api := app.Group("/customer")

	//Model
	customerModel := &model.Customer{}

	//Migration
	err := r.MigrateCustomer(customerModel)

	if err != nil {
		fmt.Println("Error migrating customer model:", err)
	}

	//JSONBody validation
	validateCustomerMiddleWare := func(ctx *fiber.Ctx) error {
		return utils.ValidateJSONBody(customerModel, ctx)
	}

	//Routes
	api.Get("/", r.GetAllCustomer)
	api.Post("/", validateCustomerMiddleWare, r.CreateNewCustomer)

}

//*Controllers

// Migrate
func (r *CustomerRepository) MigrateCustomer(model interface{}) error {
	err := r.DB.AutoMigrate(model)
	if err != nil {
		fmt.Println("Error migrating customer model:", err)
		// Handle the error, e.g., log it or return an error response
	}
	return err
}

// Create
func (r *CustomerRepository) CreateNewCustomer(ctx *fiber.Ctx) error {

	u := uuid.New()
	//Create user from data and add cutomer id by generating randum string
	customer := model.Customer{
		CustomerId: u,
	}

	err := r.DB.Create(&customer).Error

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed to create customer"})
		return nil
	}
	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{"message": "Customer created successfully"})
}

// Read
func (r *CustomerRepository) GetAllCustomer(ctx *fiber.Ctx) error {

	customers := &[]model.Customer{}

	err := r.DB.Find(&customers).Error

	if err != nil {
		ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "No customer found"})
		return nil
	}

	if len(*customers) <= 0 {
		ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "No customer found"})
		return nil
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": customers,
	})
}
