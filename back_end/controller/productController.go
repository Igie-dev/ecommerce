package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/lib"
	"github.com/laptop-shop.api/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []model.Product
	result := lib.DB.Find(&products)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	if result.RowsAffected <= 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No customers found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": products,
	})
}
