package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Allowed_origins() string {
	var frontEndUrl = os.Getenv("FRONT_END_ENDPOINT")
	var allowedOrigins = []string{frontEndUrl}
	allowed := strings.Join(allowedOrigins, ",")
	return allowed
}

var CorsConfig = cors.Config{
	AllowOrigins:     Allowed_origins(),
	AllowHeaders:     "Origin, Content-Type, Accept,Option,Authorization,authorization",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowCredentials: true,
}
