package controller

import (
	"laptop-shop/database"
	"laptop-shop/dto"
	"laptop-shop/model"
	repo "laptop-shop/repository"
	vldt "laptop-shop/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create new customer
func CreateCustomer(c *fiber.Ctx) error {
	customer := &model.CreateCustomer{}

	if err := c.BodyParser(customer); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	// Create a new validator for a Customer model.
	validate := vldt.NewValidator()
	if err := validate.Struct(customer); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": vldt.ValidatorErrors(err),
		})
	}

	customerRepo := repo.NewCustomerRepo(database.GetDB())
	// check customer already exists

	exists, err := customerRepo.Exists(customer.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "Customer with this email  already exists",
		})
	}

	//Generate password
	customer.Password, _ = GeneratePasswordHash([]byte(customer.Password))
	if err := customerRepo.Create(customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	dbCustomer, err := customerRepo.Get(customer.CustomerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"customer": dto.ToCustomer(dbCustomer),
	})

}

// Get one customer
func GetCustomer(c *fiber.Ctx) error {
	ID := c.Params("id")

	customerId := uuid.MustParse(ID)
	customerRepo := repo.NewCustomerRepo(database.GetDB())
	customer, err := customerRepo.Get(customerId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "customer were not found",
		})
	}
	return c.JSON(fiber.Map{
		"customer": dto.ToCustomer(customer),
	})
}

// Get all customer
func GetAllCustomers(c *fiber.Ctx) error {
	pageNo, pageSize := GetPagination(c)
	customerRepo := repo.NewCustomerRepo(database.GetDB())

	customers, err := customerRepo.All(pageSize, uint(pageSize*(pageNo-1)))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "customers were not found",
		})
	}

	return c.JSON(fiber.Map{
		"page":      pageNo,
		"page_size": pageSize,
		"count":     len(customers),
		"customer":  dto.ToCustomers(customers),
	})
}

// Update a customer
func UpdateCustomer(c *fiber.Ctx) error {
	ID := c.Params("id")

	customerId := uuid.MustParse(ID)

	customerRepo := repo.NewCustomerRepo(database.GetDB())
	_, err := customerRepo.Get(customerId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "customer were not found",
		})
	}
	customer := &model.UpdateCustomer{}
	if err := c.BodyParser(customer); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	// Create a new validator for a Customer model.
	validate := vldt.NewValidator()
	if err := validate.Struct(customer); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": vldt.ValidatorErrors(err),
		})
	}
	if err := customerRepo.Update(customerId, customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	dbCustomer, err := customerRepo.Get(customerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"customer": dto.ToCustomer(dbCustomer),
	})
}

// Delete a customer
func DeleteCustomer(c *fiber.Ctx) error {
	ID := c.Params("id")

	customerId := uuid.MustParse(ID)

	customerRepo := repo.NewCustomerRepo(database.GetDB())

	_, err := customerRepo.Get(customerId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "customer were not found",
		})
	}

	err = customerRepo.Delete(customerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}
