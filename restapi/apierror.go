package restapi

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Message  string            `json:"message,omitempty"`
	Property string            `json:"property,omitempty"`
	Code     string            `json:"code,omitempty"`
	Context  map[string]string `json:"context,omitempty"`
}

type ApiErrors struct {
	Errors []ApiError `json:"errors"`
}

func NewValidationError(err error) *ApiErrors {
	apiErrors := &ApiErrors{
		Errors: []ApiError{},
	}
	switch err := err.(type) {
	case validator.ValidationErrors:
		for _, fieldErr := range err {
			apiError := ApiError{
				Message:  getValidationErrorMessage(fieldErr),
				Property: fieldErr.Field(),
			}
			apiErrors.Errors = append(apiErrors.Errors, apiError)
		}
	case *json.SyntaxError:
		apiError := ApiError{
			Message: "Invalid JSON object.",
		}
		apiErrors.Errors = append(apiErrors.Errors, apiError)
	default:
		log.Println(err.Error())
		apiError := ApiError{
			Message: "Unknown error.",
		}
		apiErrors.Errors = append(apiErrors.Errors, apiError)
	}

	return apiErrors
}

func NewNotFoundError() *ApiErrors {
	var apiError = ApiError{
		Message: "Not found.",
		Code:    "not_found",
	}
	return &ApiErrors{
		Errors: []ApiError{apiError},
	}
}

func NewInternalServerError() *ApiErrors {
	var apiError = ApiError{
		Message: "Internal server error.",
		Code:    "internal_server_error",
	}
	return &ApiErrors{
		Errors: []ApiError{apiError},
	}
}

func NewUnauthorizedError() *ApiErrors {
	var apiError = ApiError{
		Message: "Unauthorized.",
		Code:    "unauthorized",
	}
	return &ApiErrors{
		Errors: []ApiError{apiError},
	}
}

func getValidationErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be at maximum %s in length", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be at minimum %s in length", e.Field(), e.Param())
	}
	if e.Field() != "" {
		return fmt.Sprintf("%s is not valid", e.Field())
	}
	return e.Error()
}
