package main

import (
	"go_fiber/main/configs"
	"go_fiber/main/routes"

	"github.com/gofiber/fiber/v2"

)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	routes.RouteInit(app)
	routes.UserRoute(app)

	app.Listen(":6000")
}
