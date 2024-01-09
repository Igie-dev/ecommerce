package routes

import (
	"path"

	"github.com/gofiber/fiber/v2"
)

func RootRoutes(app *fiber.App) {
	router := app.Group("/")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile(path.Join(".", "views", "index.html"))
	})

	router.Get("/index.html", func(c *fiber.Ctx) error {
		// Send the index.html file
		return c.SendFile(path.Join(".", "views", "index.html"))
	})
}
