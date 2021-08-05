package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
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

	app.Use(requestid.New(requestid.Config{
		Next:   nil,
		Header: fiber.HeaderXRequestID,
		Generator: func() string {
			return utils.UUID()
		},
		ContextKey: "requestid",
	}))

	routes.Setus(app)

	app.Listen(":3000")
}
