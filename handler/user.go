package handler

import (
	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/model"
	"github.com/gofiber/fiber/v2"
)

// Select all users
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB

	var users []model.User
	err := db.Find(&users).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Data:    users,
		Message: "all users",
	})
}

// Select user by id
func GetUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
	id := c.Params("id")

	// SELECT * FROM users WHERE ID equals to id, and allocate that in user var
	err := db.Table("users").Find(&user, "ID = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Data:    user,
		Message: "userFound",
	})
}

// Creates a new user on Database
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Sorry cannot insert empty fields",
			Data:    nil,
		})
	}

	// TODO:HASHING PASSWORD

	// INSERT QUERY
	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
