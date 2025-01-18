package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := new(models.User)

		err := c.BodyParser(user)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse Body",
				"error":   err.Error(),
			})
		}

		if user.Email == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Email or Password is missing",
			})
		}

		hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not hash password",
				"error":   err.Error(),
			})
		}

		user.Password = string(hasedPassword)

		err = db.Create(user).Error
		if err != nil {

			if strings.Contains(err.Error(), "SQLSTATE 23505") {
				return c.Status(400).JSON(fiber.Map{
					"message": "Email already exists",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create user",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "User created",
			"data":    user,
		})
	}
}

func GetUsers(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		users := new([]models.User)

		err := db.Find(users).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve users",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Retrieved users",
			"data":    users,
		})
	}
}

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
