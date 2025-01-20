package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebastian-abraham/recoverly/controller"
	"gorm.io/gorm"
)

func TaskRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/task", controller.CreateTask(db))
}
