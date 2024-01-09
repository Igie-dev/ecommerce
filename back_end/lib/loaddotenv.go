package lib

import (
	"log"

	"github.com/joho/godotenv"
)

func Loaddotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
