package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vara855/gochat/app/models"
	"github.com/vara855/gochat/app/services"
)

func GetUsers(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.GetAll()
		if err != nil {
			return &fiber.Error{
				Message: "Failed",
			}
		}
		return c.JSON(fiber.Map{
			"success": true,
			"users":   users,
		})
	}
}

func PostUsers(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		newUser := new(models.User)

		fmt.Printf("newUser: %v\n", newUser)
		if err := c.BodyParser(newUser); err != nil {

			return err
		}
		inserted, err := service.Insert(newUser)
		if err != nil {
			return &fiber.Error{
				Message: "Failed",
			}
		}
		return c.JSON(fiber.Map{
			"success": true,
			"user":   inserted,
		})
	}
}
