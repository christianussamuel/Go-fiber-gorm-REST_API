package routes

import (
	"go_fiber/main/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/api", handler.UserHandlerCreate)
}
