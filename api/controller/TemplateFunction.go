package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/models"
	"gorm.io/gorm"
)

func CreateTemplate(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {	
		template := new(models.Template)
		err := c.BodyParser(template)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Could not parse JSON",
			})
		}

		err = db.Create(template).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not create task",
				"error":   err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Task created",
			"data":    template,
		})
	}
}
