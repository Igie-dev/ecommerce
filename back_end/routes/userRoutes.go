package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop_api/controller"
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