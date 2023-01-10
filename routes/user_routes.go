package routes

import (
	"go_fiber/main/controllers"

	"github.com/gofiber/fiber/v2"

)

func UserRoute(app *fiber.App) {
	//userprofile
	app.Post("/api/profile", controllers.CreateUser)
	app.Get("/api/profile/:profileCode", controllers.GetAUser)
	app.Put("api/profile/:profileCode", controllers.EditAUser)

	//workingexp
	app.Get("/api/working-experience/:profileCode", controllers.GetWorkingExperience)
	app.Put("/api/working-experience/:profileCode", controllers.EditWorkingExperience)

	//employment
	app.Post("/api/employment/:profileCode", controllers.CreateEmployment)

}
