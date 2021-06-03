package handler

import (
	"gorest/filter"
	"gorest/model"
	"gorest/restapi"
	"gorest/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GET /books
// Get all books
func (h *Handler) ListBooksRequest(ctx *fiber.Ctx) error {
	var pagination = restapi.NewPagination()
	if err := ctx.QueryParser(pagination); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	var filter filter.BookFilter
	if err := ctx.QueryParser(&filter); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	books, count, err := service.ListBooks(h.DB, h.User, pagination, &filter)
	if err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}
	return restapi.ListResponse(ctx, books, count, pagination)
}

// POST /books
// Create new book
func (h *Handler) CreateBookRequest(ctx *fiber.Ctx) error {
	// Validate input
	var book = &model.Book{
		UserID: h.User.ID,
	}
	if err := ctx.BodyParser(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	if err := restapi.ValidateStruct(book); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Create book
	if err := service.CreateBook(h.DB, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.CreatedResponse(ctx, book)
}

// GET /books/:id
// Find a book
func (h *Handler) RetrieveBookRequest(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	book, _ := service.RetrieveBook(h.DB, id, h.User)
	if book == nil {
		return restapi.NotFoundRespone(ctx)
	}
	return restapi.DetailResponse(ctx, book)
}

// PATCH /books/:id
// Update a book
func (h *Handler) UpdateBookRequest(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	book, _ := service.RetrieveBook(h.DB, id, h.User)
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
	if err := service.UpdateBook(h.DB, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.DetailResponse(ctx, book)
}

// DELETE /books/:id
// Delete a book
func (h *Handler) DeleteBookRequest(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	// Get model if exist
	book, _ := service.RetrieveBook(h.DB, id, h.User)
	if book == nil {
		return restapi.NotFoundRespone(ctx)
	}
	if err := service.DeleteBook(h.DB, book); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.DeletedResponse(ctx)
}
