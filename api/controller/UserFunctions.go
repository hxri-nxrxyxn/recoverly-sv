package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userData := new(models.User)

		err := c.BodyParser(userData)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse Body",
				"error":   err.Error(),
			})
		}

		if userData.Email == "" || userData.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Email or Password is missing",
			})
		}

		user := new(models.User)
		err = db.Where("email = ?", userData.Email).First(user).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return c.Status(404).JSON(fiber.Map{
					"message": "User not found",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve user",
				"error":   err.Error(),
			})
		}

		if user.Password != userData.Password {
			return c.Status(401).JSON(fiber.Map{
				"message": "Password is incorrect",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "User logged in",
		})
	}
}
