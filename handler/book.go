package handler

import (
	"gorest/filter"
	"gorest/model"
	"gorest/restapi"
	"gorest/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /books
// Get all books
func ListBooksRequest(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	var pagination = restapi.NewPagination()
	if err := ctx.QueryParser(pagination); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	var filter filter.BookFilter
	if err := ctx.QueryParser(&filter); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	db := ctx.Locals("db").(*gorm.DB)
	books, count, err := service.ListBooks(db, user, pagination, &filter)
	if err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}
	return restapi.ListResponse(ctx, books, count, pagination)
}

// POST /books
// Create new book
func CreateBookRequest(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	// Validate input
	var book = &model.Book{
		UserID: user.ID,
	}
	if err := ctx.BodyParser(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	if err := restapi.ValidateStruct(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Create book
	db := ctx.Locals("db").(*gorm.DB)
	if err := service.CreateBook(db, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.CreatedResponse(ctx, book)
}

// GET /books/:id
// Find a book
func RetrieveBookRequest(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	db := ctx.Locals("db").(*gorm.DB)
	book, _ := service.RetrieveBook(db, id, user)
	if book == nil {
		return restapi.NotFoundRespone(ctx)
	}
	return restapi.DetailResponse(ctx, book)
}

// PATCH /books/:id
// Update a book
func UpdateBookRequest(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	db := ctx.Locals("db").(*gorm.DB)
	book, _ := service.RetrieveBook(db, id, user)
	if book == nil {
		return restapi.NotFoundRespone(ctx)
	}

	// Parse
	if err := ctx.BodyParser(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Validate
	if err := restapi.ValidateStruct(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Update
	if err := service.UpdateBook(db, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.DetailResponse(ctx, book)
}

// DELETE /books/:id
// Delete a book
func DeleteBookRequest(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	db := ctx.Locals("db").(*gorm.DB)
	book, _ := service.RetrieveBook(db, id, user)
	if book == nil {
		return restapi.NotFoundRespone(ctx)
	}
	if err := service.DeleteBook(db, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.DeletedResponse(ctx)
}
