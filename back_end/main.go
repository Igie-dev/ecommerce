package main

import (
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/laptop-shop.api/lib"
	"github.com/laptop-shop.api/routes"
)

func main() {
	app := fiber.New()
	lib.Loaddotenv()

	//Open database
	lib.DbConnection()
	defer lib.CloseDatabase()

	//Cors
	app.Use(cors.New(lib.CorsConfig))

	app.Static("/", "./public")

	//Home routes
	routes.RootRoutes(app)
	//User routes
	routes.CustomerRoute(app)

	// Define a middleware to handle all routes
	app.Use(func(c *fiber.Ctx) error {
		// Set the status code to 404

		c.Status(fiber.StatusNotFound)
		//Error routes
		switch c.Accepts("html", "json", "text") {
		case "html":
			return c.SendFile(path.Join(".", "views", "error.html"))
		case "json":
			return c.JSON(fiber.Map{"message": "Not found!"})
		default:
			return c.SendString("404 not found!")
		}
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
