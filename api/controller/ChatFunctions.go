package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func CreateChatGroup(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		chatGroup := new(models.ChatGroup)
		err := c.BodyParser(chatGroup)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse JSON",
			})
		}

		chatGroup.Chats = []models.Chat{
			{
				UserID:  1,
				Message: "Welcome to the chat group",
				Sent:    time.Now(),
			},
		}

		err = db.Create(chatGroup).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create chat group",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Chat group created",
			"data":    chatGroup,
		})

	}
}
