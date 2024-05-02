package controller

import (
	"laptop-shop-api/database"
	"laptop-shop-api/dto"
	"laptop-shop-api/model"
	repo "laptop-shop-api/repository"
	vldt "laptop-shop-api/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create new user
func CreateUser(c *fiber.Ctx) error {
	user := &model.CreateUser{}

	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	// Create a new validator for a User model.
	validate := vldt.NewValidator()
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": vldt.ValidatorErrors(err),
		})
	}

	userRepo := repo.NewUserRepo(database.GetDB())
	// check user already exists

	exists, err := userRepo.Exists(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "User with this email  already exists",
		})
	}

	//Generate password
	user.Password, _ = GeneratePasswordHash([]byte(user.Password))
	if err := userRepo.Create(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	dbUser, err := userRepo.Get(user.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": dto.ToUser(dbUser),
	})

}

// Get one user
func GetUser(c *fiber.Ctx) error {
	ID := c.Params("id")

	userId := uuid.MustParse(ID)
	userRepo := repo.NewUserRepo(database.GetDB())
	user, err := userRepo.Get(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user were not found",
		})
	}
	return c.JSON(fiber.Map{
		"user": dto.ToUser(user),
	})
}

// Get all user
func GetAllUsers(c *fiber.Ctx) error {
	pageNo, pageSize := GetPagination(c)
	userRepo := repo.NewUserRepo(database.GetDB())

	users, err := userRepo.All(pageSize, uint(pageSize*(pageNo-1)))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "users were not found",
		})
	}

	return c.JSON(fiber.Map{
		"page":      pageNo,
		"page_size": pageSize,
		"count":     len(users),
		"user":  dto.ToUsers(users),
	})
}

// Update a user
func UpdateUser(c *fiber.Ctx) error {
	ID := c.Params("id")

	userId := uuid.MustParse(ID)

	userRepo := repo.NewUserRepo(database.GetDB())
	_, err := userRepo.Get(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user were not found",
		})
	}
	user := &model.UpdateUser{}
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	// Create a new validator for a User model.
	validate := vldt.NewValidator()
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "invalid input found",
			"errors": vldt.ValidatorErrors(err),
		})
	}
	if err := userRepo.Update(userId, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}
	dbUser, err := userRepo.Get(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": dto.ToUser(dbUser),
	})
}

// Delete a user
func DeleteUser(c *fiber.Ctx) error {
	ID := c.Params("id")

	userId := uuid.MustParse(ID)

	userRepo := repo.NewUserRepo(database.GetDB())

	_, err := userRepo.Get(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "user were not found",
		})
	}

	err = userRepo.Delete(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{})
}
