package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var requestBody map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Simulate authentication logic
	if requestBody["username"] == "user" && requestBody["password"] == "password" {
		return c.JSON(fiber.Map{"message": "Login successful"})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}
