package main

import (
	"log"
	"os"
	"path"

	"ecommerce.com/api/config"
	"ecommerce.com/api/lib"
	"ecommerce.com/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	lib.Loaddotenv()
	//Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Allowed_origins(),
		AllowHeaders:     "Origin, Content-Type, Accept,Option,Authorization,authorization",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	app.Static("/", "./public")

	//!Sample get request
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	//Home routes
	routes.RootRoutes(app)

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
