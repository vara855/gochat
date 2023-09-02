package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vara855/gochat/app/handlers"
)

func RootRoutes(app *fiber.App, mainGroup fiber.Router) {

	app.Get("", handlers.RootHandler)

	mainGroup.Get("", handlers.RootHandler)
}
