package router

import (
	"gorest/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	users := app.Group("/users")
	users.Post("", handler.CreateUserRequest)
	users.Get("", handler.ListUsersRequest)

	books := app.Group("/books", handler.JWTAuthenticateUser)
	books.Get("", handler.ListBooksRequest)
	books.Post("", handler.CreateBookRequest)
	books.Get("/:id", handler.RetrieveBookRequest)
	books.Patch("/:id", handler.UpdateBookRequest)
	books.Delete("/:id", handler.DeleteBookRequest)
}

func NewApp(db *gorm.DB) *fiber.App {
	app := fiber.New()

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("db", db.Session(&gorm.Session{}))
		return ctx.Next()
	})
	app.Use(logger.New())
	SetupRoutes(app)
	return app
}
