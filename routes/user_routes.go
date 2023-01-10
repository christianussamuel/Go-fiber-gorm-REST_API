package routes

import (
	"go_fiber/main/controllers"

	"github.com/gofiber/fiber/v2"

)

func UserRoute(app *fiber.App) {
	app.Post("/api/profile", controllers.CreateUser)
}
