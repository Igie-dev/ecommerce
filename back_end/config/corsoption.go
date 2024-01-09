package config

import (
	"fmt"
	"os"
	"strings"
)

func Allowed_origins() string {
	var front_end_enpoint = os.Getenv("FRONT_END_ENDPOINT")
	var allowedOrigins = []string{fmt.Sprintf("%s", front_end_enpoint), "oots"}
	allowed := strings.Join(allowedOrigins, ",")
	return allowed
}
