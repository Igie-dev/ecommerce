package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/controller"
)

func CustomerRoute(app *fiber.App) {
	router := app.Group("/customers")

	router.Get("/", controller.GetAllUsers)
}
