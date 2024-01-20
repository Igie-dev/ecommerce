package controller

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
	Address   string `json:"address"`
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []User{
		{Id: "1",
			FirstName: "Igie",
			LastName:  "Baldesanso",
			Age:       24,
			Address:   "Marikina"},
		{Id: "2",
			FirstName: "Jose",
			LastName:  "Cruz",
			Age:       24,
			Address:   "Marikina"},
	}

	if len(users) <= 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No users found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})
}
