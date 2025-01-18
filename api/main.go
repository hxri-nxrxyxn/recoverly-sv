package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sebastian-abraham/recoverly/database"
	"github.com/sebastian-abraham/recoverly/models"
	"github.com/sebastian-abraham/recoverly/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	models.MigrateUser(db)
	models.MigrateChat(db)
	models.MigrateChatGroup(db)

	routes.UserRoutes(db, app)
	routes.ChatRoutes(db, app)

	app.Listen(":8080")
}
