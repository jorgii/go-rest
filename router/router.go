package router

import (
	"gorest/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	users := app.Group("/users")
	users.Post("", controller.CreateUserRequest)
	users.Get("", controller.ListUsersRequest)

	books := app.Group("/books", controller.JWTAuthenticateUser)
	books.Get("", controller.ListBooksRequest)
	books.Post("", controller.CreateBookRequest)
	books.Get("/:id", controller.RetrieveBookRequest)
	books.Patch("/:id", controller.UpdateBookRequest)
	books.Delete("/:id", controller.DeleteBookRequest)
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
