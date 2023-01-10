package handler

import (
	"github.com/gofiber/fiber/v2"

)

func UserHandlerCreate(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"welcome_message": "hello dunia",
	})

}
