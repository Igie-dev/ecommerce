package routes

import (
	"ecommerce.com/api/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	router := app.Group("/users")

	router.Get("/", controller.GetAllUsers)
}
