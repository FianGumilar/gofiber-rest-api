package middleware

import (
	"github.com/FianGumilar/gofiber-rest-api/utils"
	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	token := c.Get("token") // Get token header
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token is required",
		})
	}

	//_, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not admin",
		})
	} 

	return c.Next()
}
