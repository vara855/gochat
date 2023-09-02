package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vara855/gochat/app/routes"
	"github.com/vara855/gochat/app/services"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
	// db   = flag.String("db", "memory", "Which db to use? ")
)

func main() {
	flag.Parse()
	app := fiber.New(fiber.Config{
		AppName:           "varadotdev chat",
		Prefork:           *prod,
		EnablePrintRoutes: true,
	})

	userService := services.New()


	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Register routes
	v1 := app.Group("/api/v1")
	routes.RootRoutes(app, v1)
	routes.UserRoutes(v1, userService)

	log.Fatal(app.Listen(*port))
}
