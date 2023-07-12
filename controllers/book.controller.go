package controllers

import (
	"github.com/FianGumilar/gofiber-rest-api/database"
	"github.com/FianGumilar/gofiber-rest-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllBook(c *fiber.Ctx) error {
	var books []*models.Book

	database.DB.Debug().Find(&books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"books":   books,
		"message": "Success Get All Book",
	})
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Validate Request
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
		})
	}

	filename := c.Locals("filename").(string)

	NewBook := models.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
	}

	database.DB.Debug().Create(&NewBook)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"book":    NewBook,
	})
}
