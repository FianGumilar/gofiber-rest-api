package middleware

import "github.com/gofiber/fiber/v2"

func UserAuth(c *fiber.Ctx) error {
	token := c.Get("token")
	if token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
