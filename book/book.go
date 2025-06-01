package book

import (
	"fmt"
	"go-fiber-api/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []database.Book
	result := database.Db.Find(&books)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")
	var book database.Book

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Integer expected for book ID, got string: " + c.Params("id"),
		})
	}

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusFound).JSON(book)
}

func AddBook(c *fiber.Ctx) error {
	var book database.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if result := database.Db.Create(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var book database.Book

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		errMessage := fmt.Sprintf("Book with ID: %d not found.", bookID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": errMessage,
		})
	}

	if result = database.Db.Delete(&book, bookID); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bookId":  bookID,
		"message": "Deleted book",
	})
}

func UpdateBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var updateData map[string]interface{}

	if err = c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errror": "Invalid JSON",
		})
	}

	var book database.Book

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Book with ID: %d not found.", bookID),
		})
	}

	result = database.Db.Model(&book).Updates(updateData)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "Success",
		"message": fmt.Sprintf("Book with ID: %d updated", bookID),
		"data":    book,
	})
}
