package lib

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/laptop.api/config"
)

var CorsConfig = cors.Config{
	AllowOrigins:     config.Allowed_origins(),
	AllowHeaders:     "Origin, Content-Type, Accept,Option,Authorization,authorization",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowCredentials: true,
}
