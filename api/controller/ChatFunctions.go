package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func CreateChatGroup(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		chat := new(models.Chat)
		err := c.BodyParser(chat)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse JSON",
			})
		}

		err = db.Create(chat).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create chat",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Chat created",
			"data":    chat,
		})

	}
}
