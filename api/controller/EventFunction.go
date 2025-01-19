package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func GetEvent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		event := new(models.Event)

		err := db.Where("event_id = ?", id).First(event).Error
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"message": "Event not found",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Event found",
			"data":    event,
		})
	}
}

func CreateEvent(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		event := new(models.Event)
		err := c.BodyParser(event)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse JSON",
			})
		}

		err = db.Create(event).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create chat",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Event created",
			"data":    event,
		})

	}
}

func GetEvents(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		events := new([]models.Event)

		err := db.Find(events).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not retrieve chats",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Retrieved chats",
			"data":    events,
		})
	}
}
