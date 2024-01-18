package controller

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        int32  `json:"age"`
	Address    string `json:"address"`
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []User{
		{Id: "1",
			First_name: "Igie",
			Last_name:  "Baldesanso",
			Age:        24,
			Address:    "Marikina"},
		{Id: "2",
			First_name: "Jose",
			Last_name:  "Cruz",
			Age:        24,
			Address:    "Marikina"},
	}

	if len(users) <= 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No users found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})
}
