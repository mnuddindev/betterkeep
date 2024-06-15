package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
)

func NotImplemented(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello there!!",
		"status":  fiber.StatusOK,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	if utils.IsEmpty(user.FirstName) || utils.IsEmpty(user.LastName) || utils.IsEmpty(user.Email) || utils.IsEmpty(user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "fields cannot be empty",
			"status": fiber.StatusUnauthorized,
		})
	}
	if !utils.IsEmail(user.Email) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "invalid email format",
			"status": fiber.StatusUnauthorized,
		})
	}
}
