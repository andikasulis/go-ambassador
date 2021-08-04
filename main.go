package main

import (
	"go-ambassador/src/database"
	"go-ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	routes.Setus(app)

	app.Listen(":3000")
}
