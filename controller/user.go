package controller

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
	"gorm.io/gorm"
)

// POST /users
// Create new user
func CreateUserRequest(ctx *fiber.Ctx) error {
	// Validate input
	var user = &model.User{}

	if err := ctx.BodyParser(user); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	if err := restapi.ValidateStruct(user); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}

	// Create user
	db := ctx.Locals("db").(*gorm.DB)
	if err := service.CreateUser(db, user); err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}

	return restapi.CreatedResponse(ctx, user)
}

func ListUsersRequest(ctx *fiber.Ctx) error {
	var pagination = restapi.NewPagination()
	if err := ctx.QueryParser(pagination); err != nil {
		return restapi.ValidationErrorRespone(ctx, err)
	}
	db := ctx.Locals("db").(*gorm.DB)
	users, count, err := service.ListUsers(db, pagination)
	if err != nil {
		return restapi.InternalServerErrorResponse(ctx)
	}
	return restapi.ListResponse(ctx, users, count, pagination)
}

func JWTAuthenticateUser(ctx *fiber.Ctx) error {
	tokenstr := restapi.ExtractBearerToken(ctx.Get("Authorization"))
	if tokenstr == "" {
		return restapi.UnauthorizedErrorResponse(ctx)
	}
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		publicKey := []byte("-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoIVw9/WlphgWveot3GKI8lmJWSGcB7e3PRM41792Uh2ENfZMCmEhyIdakNGqe2wuM4/nJmqaCjZvvrA+0AqawrpGY1K9ptLVfQZTfnveFC4V+jnDHpaF/XDIG1pyZfGJt/GSEBW5Y9DJ/l/Cndgv3Flr2tqeaGsae1KyqERZfqcRhk0Aw+G4WUdfxZKjoAjRe2fIePHWvoDudbWOxXa/jXkX7LipZY71Y6r2E27c+sdQsOws5UT0jnGVOqUAyIrzh42koZk1a5yhhL3Bquaoe86YriW16VJUfEDRqlrkEXrfS+m738wxV3Xh/nsRox4C8NarvtNuYALmiwjAFNS3dwIDAQAB\n-----END PUBLIC KEY-----")
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
	db := ctx.Locals("db").(*gorm.DB)
	user, err := service.RetrieveUserByEmail(db, email)
	if err != nil {
		return restapi.UnauthorizedErrorResponse(ctx)
	}
	ctx.Locals("user", user)
	return ctx.Next()
}
