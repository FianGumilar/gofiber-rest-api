package main

import (
	"log"

	"github.com/FianGumilar/gofiber-rest-api/config"
	"github.com/FianGumilar/gofiber-rest-api/database"
	"github.com/FianGumilar/gofiber-rest-api/database/migrations"
	"github.com/FianGumilar/gofiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// DATABASE
	database.DatabaseInit()

	// MIGRATION
	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	// ROUTE INIT
	routes.RouteInit(app)

	app.Listen(config.Config("PORT"))
}
