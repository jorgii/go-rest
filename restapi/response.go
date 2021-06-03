package restapi

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DetailResponse(ctx *fiber.Ctx, body interface{}) error {
	return ctx.Status(http.StatusOK).JSON(body)
}

func ListResponse(ctx *fiber.Ctx, data interface{}, count int64, pagination *Pagination) error {
	return ctx.Status(http.StatusOK).JSON(NewPageFromPagination(pagination, data, count))
}

func CreatedResponse(ctx *fiber.Ctx, body interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(body)
}

func DeletedResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNoContent).Send(nil)
}

func NotFoundRespone(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNotFound).JSON(NewNotFoundError())
}

func ValidationErrorRespone(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusBadRequest).JSON(NewValidationError(err))
}

func InternalServerErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(NewInternalServerError())
}

func UnauthorizedErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusUnauthorized).JSON(NewUnauthorizedError())
}
