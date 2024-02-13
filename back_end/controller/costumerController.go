package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/model"
)

func GetAllUsers(c *fiber.Ctx) error {
	customers := []model.Customer{
		{ID: 1,
			FirstName: "Igie",
			LastName:  "Baldesanso",
			Address:   "Marikina"},
		{ID: 2,
			FirstName: "Jose",
			LastName:  "Cruz",
			Address:   "Marikina"},
	}

	if len(customers) <= 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No ucustomers found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": customers,
	})
}
