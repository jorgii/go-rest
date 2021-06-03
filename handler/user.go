package handler

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"gorest/model"
	"gorest/restapi"
	"gorest/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// POST /users
// Create new user
func (h *Handler) CreateUserRequest(ctx *fiber.Ctx) error {
	// Validate input
	var user = &model.User{}

	if err := ctx.BodyParser(user); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	if err := restapi.ValidateStruct(user); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Create user
	if err := service.CreateUser(h.DB, user); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.CreatedResponse(ctx, user)
}

func (h *Handler) ListUsersRequest(ctx *fiber.Ctx) error {
	var pagination = restapi.NewPagination()
	if err := ctx.QueryParser(pagination); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	users, count, err := service.ListUsers(h.DB, pagination)
	if err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}
	return restapi.ListResponse(ctx, users, count, pagination)
}

func (h *Handler) JWTAuthenticateUser(ctx *fiber.Ctx) error {
	tokenstr := restapi.ExtractBearerToken(ctx.Get("Authorization"))
	if tokenstr == "" {
		return restapi.UnauthorizedErrorResponse(ctx)
	}
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		publicKey := []byte(h.Config.JWTPublicKey)
		block, _ := pem.Decode(publicKey)
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pub, ok := pub.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("could not convert key")
		}
		return pub, nil
	})
	if err != nil {
		return restapi.UnauthorizedErrorResponse(ctx)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return restapi.InternalServerErrorResponse(ctx)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return restapi.InternalServerErrorResponse(ctx)
	}
	user, err := service.RetrieveUserByEmail(h.DB, email)
	if err != nil {
		return restapi.UnauthorizedErrorResponse(ctx)
	}
	h.User = user
	return ctx.Next()
}
