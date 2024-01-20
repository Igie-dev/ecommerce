package config

import (
	"fmt"
	"os"
	"strings"
)

func Allowed_origins() string {
	var frontEndUrl = os.Getenv("FRONT_END_ENDPOINT")
	var allowedOrigins = []string{fmt.Sprintf("%s", frontEndUrl)}
	allowed := strings.Join(allowedOrigins, ",")
	return allowed
}
