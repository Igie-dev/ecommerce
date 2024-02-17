package utils

import (
	"errors"
	"fmt"
	"strings"

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
	// var validationErrors []*ErrorResponse
	var errField []string

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
			errField = append(errField, err.Field())
		}

		isOrare := "is"

		if len(errField) > 1 {
			isOrare = "are"
		}

		errmsg := fmt.Sprintf("%s field %s required", strings.Join(errField, " , "), isOrare)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errmsg})
	}
	return c.Next()
}
