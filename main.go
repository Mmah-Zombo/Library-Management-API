package main

import (
	"go-fiber-api/book"
	"go-fiber-api/database"

	"github.com/gofiber/fiber/v2"
)

func setupRoute(app *fiber.App) {
	app.Route("/api/v1", func(api fiber.Router) {
		api.Get("/books", book.GetBooks)
		api.Get("/books/:id", book.GetBook)
		api.Post("/books", book.AddBook)
		api.Put("/books/:id", book.UpdateBook)
		api.Delete("/books/:id", book.DeleteBook)
	})
}

func main() {
	app := fiber.New()

	database.Init()
	setupRoute(app)

	app.Listen(":3000")
}
