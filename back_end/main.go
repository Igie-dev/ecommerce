package main

import (
	"log"
	"os"
	"path"

	"laptop-shop-api/config"
	"laptop-shop-api/database"
	"laptop-shop-api/lib"
	"laptop-shop-api/middleware"
	"laptop-shop-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	lib.Loaddotenv()

	//Open database
	err := database.ConnectDB()

	if err != nil {
		log.Panicf("Cannot connect to db: %v", err)
	}

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	//Cors
	app.Use(cors.New(config.CorsConfig))

	app.Static("/", "./public")

	//Home routes
	routes.RootRoutes(app)

	//user routes

	routes.CustomerRoutes(app)

	// Define a middleware to handle all routes
	app.Use(func(c *fiber.Ctx) error {
		// Set the status code to 404

		c.Status(fiber.StatusNotFound)
		//Error routes
		switch c.Accepts("html", "json", "text") {
		case "html":
			return c.SendFile(path.Join(".", "views","html", "error.html"))
		case "json":
			return c.JSON(fiber.Map{"message": "Not found!"})
		default:
			return c.SendString("404 not found!")
		}
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
