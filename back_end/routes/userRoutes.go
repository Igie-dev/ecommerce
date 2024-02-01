package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/controller"
)

func UserRoutes(app *fiber.App) {
	router := app.Group("/users")

	router.Get("/", controller.GetAllUsers)
}
