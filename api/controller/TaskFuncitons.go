package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		err := c.BodyParser(task)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse JSON",
			})
		}

		err = db.Create(task).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create task",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Task created",
			"data":    task,
		})
	}
}
