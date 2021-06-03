package router

import (
	"gorest/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.Handler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	users := app.Group("/users")
	users.Post("", h.CreateUserRequest)
	users.Get("", h.ListUsersRequest)

	books := app.Group("/books", h.JWTAuthenticateUser)
	books.Get("", h.ListBooksRequest)
	books.Post("", h.CreateBookRequest)
	books.Get("/:id", h.RetrieveBookRequest)
	books.Patch("/:id", h.UpdateBookRequest)
	books.Delete("/:id", h.DeleteBookRequest)
}
