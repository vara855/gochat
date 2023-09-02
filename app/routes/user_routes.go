package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vara855/gochat/app/handlers"
	"github.com/vara855/gochat/app/services"
)

func UserRoutes(router fiber.Router, userService services.UserService) {

	router.Get("/users", handlers.GetUsers(userService))

	router.Post("/users", handlers.PostUsers(userService))
}
