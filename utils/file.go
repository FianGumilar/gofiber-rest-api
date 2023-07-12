package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SingleFile(c *fiber.Ctx) error {
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("Error", errFile)
	}

	var filename *string

	if file != nil {
		filename = &file.Filename

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/cover/%s", *filename))
		if errSaveFile != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to save file",
			})
		}
	}
	c.Locals("filename", *filename)

	return c.Next()
}
