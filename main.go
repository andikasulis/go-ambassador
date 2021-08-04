package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-ambassador/src/database"
	"go-ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setus(app)

	app.Listen(":3000")
}
