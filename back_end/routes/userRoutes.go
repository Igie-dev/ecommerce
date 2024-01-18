package routes

import (
	"ecommerce.com/api/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	router := app.Group("/user")

	router.Get("/", controller.GetAllUsers)
}
