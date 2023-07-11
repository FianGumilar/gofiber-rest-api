package controllers

import (
	"time"

	"github.com/FianGumilar/gofiber-rest-api/database"
	"github.com/FianGumilar/gofiber-rest-api/models"
	"github.com/FianGumilar/gofiber-rest-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.Login)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Login",
			"error":   errValidate.Error(),
		})
	}

	var user models.User

	// Check Validation Email
	errEmail := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error
	if errEmail != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check Password Validation
	isValid := utils.CheckHash(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password is not valid",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	// Claims or Payloads Definitions
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	token, errGenerate := utils.GenerateToken(&claims)

	if errGenerate != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Token is not valid",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
