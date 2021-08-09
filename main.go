package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"go-ambassador/src/database"
	"go-ambassador/src/routes"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()

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

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	routes.Setus(app)

	app.Listen(":3000")
}
