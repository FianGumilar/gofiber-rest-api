package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
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

func HandleMultipleFile(c *fiber.Ctx) error {
	forms, errForms := c.MultipartForm()
	if errForms != nil {
		log.Println("error processing multipart", errForms)
	}

	files := forms.File["photo"]

	var filenames []string

	for i, file := range files {
		var filename string

		if file != nil {
			filename = fmt.Sprintf("%s-%d", file.Filename, i)

			errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/cover/%s", filename))
			if errSaveFile != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Failed to save file",
				})
			} else {
				fmt.Println("Nothing files to upload")
			}
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	c.Locals("filenames", filenames)

	return c.Next()
}
