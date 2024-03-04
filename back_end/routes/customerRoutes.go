package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/controller"
)

func CustomerRoutes(app *fiber.App) {

	api := app.Group("/customer")

	//Routes
	api.Post("/", controller.CreateCustomer)
	api.Get("/", controller.GetAllCustomers)
	api.Get("/:id", controller.GetCustomer)
	api.Patch("/:id", controller.UpdateCustomer)
	api.Delete("/:id", controller.DeleteCustomer)
}
