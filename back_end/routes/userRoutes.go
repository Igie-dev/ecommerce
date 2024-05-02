package routes

import (
	"laptop-shop-api/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	api := app.Group("/user")

	//Routes
	api.Post("/", controller.CreateUser)
	api.Get("/", controller.GetAllUsers)
	api.Get("/:id", controller.GetUser)
	api.Patch("/:id", controller.UpdateUser)
	api.Delete("/:id", controller.DeleteUser)
}