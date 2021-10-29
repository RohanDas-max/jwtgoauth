package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/shoping-site/controller"
)

func Routes(app *fiber.App) {

	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
}
