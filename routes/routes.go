package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/shoping-site/controller"
)

func Routes(app *fiber.App) {

	app.Get("/", controller.Hello)
}
