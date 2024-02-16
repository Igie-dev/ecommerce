package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Repository) ProductRoutes(app *fiber.App) {
	api := app.Group("/product")
	api.Get("/", r.GetAllCustomer)
}

// Create

func (r *Repository) CreateProduct(app *fiber.App) {

}
