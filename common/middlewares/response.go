package middlewares

import "github.com/gofiber/fiber/v2"

func unauthorizedResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "message": message})
}