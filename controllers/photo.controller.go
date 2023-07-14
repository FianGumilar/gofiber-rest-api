package controllers

import (
	"fmt"
	"log"

	"github.com/FianGumilar/gofiber-rest-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreatePhoto(c *fiber.Ctx) error {
	photos := new(models.Photo)

	if err := c.BodyParser(photos); err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}

	// Validator
	validate := validator.New()
	errValidate := validate.Struct(photos)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
		})
	}

	// Validation image requires
	var filenameString string

	filenames := c.Locals("filenames")

	log.Println("filename", filenames)

	if filenames == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "Image cover required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenames)
	}
	log.Println(filenameString)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
