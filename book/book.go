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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "Could not fetch books",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   nil,
		"message": "All books fetched successfully",
		"data":    books,
	})
}

func GetBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")
	var book database.Book

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Integer expected for book ID, got string: " + c.Params("id"),
			"data":    nil,
		})
	}

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "No book with ID: " + c.Params("id"),
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"error":   nil,
		"message": fmt.Sprintf("Book with ID %s found", c.Params("id")),
		"data":    book,
	})
}

func AddBook(c *fiber.Ctx) error {
	var book database.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid request data",
			"data":    nil,
		})
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if result := database.Db.Create(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "Failed to add book",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   nil,
		"message": fmt.Sprintf("Book with title: '%s' sucessfully added", book.Title),
		"data":    book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Integer expected for book ID, got string: " + c.Params("id"),
			"data":    nil,
		})
	}

	var book database.Book

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": fmt.Sprintf("Book with ID: %d not found.", bookID),
			"data":    nil,
		})
	}

	if result = database.Db.Delete(&book, bookID); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "Failed to delete book",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   nil,
		"message": fmt.Sprintf("Book with ID '%d' deleted", bookID),
		"data":    nil,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	bookID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Integer expected for book ID, got string: " + c.Params("id"),
			"data":    nil,
		})
	}

	var updateData map[string]interface{}

	if err = c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errror":  err.Error(),
			"message": "Invalid request data",
			"data":    nil,
		})
	}

	var book database.Book

	result := database.Db.First(&book, bookID)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": fmt.Sprintf("Book with ID: %d not found.", bookID),
			"data":    nil,
		})
	}

	if dateStr, ok := updateData["publish_date"].(string); ok {
		date, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   err.Error(),
				"message": "Invalid date format for publish_date",
				"data":    nil,
			})
		}

		updateData["publish_date"] = date
	}
	result = database.Db.Model(&book).Updates(updateData)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   result.Error.Error(),
			"message": "Failed to update book",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"error":   nil,
		"message": fmt.Sprintf("Book with ID: %d updated", bookID),
		"data":    book,
	})
}
