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
	}))

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load database")
	}

	certFile := "self-signed.crt"
	keyFile := "private.key"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	models.MigrateUser(db)

	routes.UserRoutes(db, app)

	log.Println("Starting HTTPS server on https://localhost:8080")
	if err := app.ListenTLS(":8080", certFile, keyFile); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
