package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/tietangle/web/handlers"
)

func main() {
	app := fiber.New()

	app.Static("/", "./tie-tangle/build")

	api := app.Group("/api")
	api.Post("/login", handlers.LoginHandler)

	app.Listen(":8080")
}
