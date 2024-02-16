package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/laptop-shop.api/model"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate = validator.New()

// Validate customer data
func ValidateCustomer(c *fiber.Ctx) error {
	var validationErrors []*ErrorResponse
	var mr *MalformedRequest

	body := new(model.Customer)

	err := DecodeJSONBody(c, &body)

	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(fiber.Map{"message": mr.Msg})
		}
	}
	err = validate.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ErrorResponse
			el.FailedField = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	return c.Next()
}
