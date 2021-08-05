package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-ambassador/src/controllers"
)

func Setus(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)
	admin.Get("user", controllers.User)
	admin.Post("logout", controllers.Logou)
}
