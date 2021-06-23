package middleware

import (
	"github.com/adrienBdx/chore-todos/gofiber/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("SECRET"), //config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"message": "Invalid or expired JWT"})
}

func GetUserClaims(ctx *fiber.Ctx) (fiber.Handler, error) {
	err := handlers.Authorize(ctx)

	if err != nil {
		return nil, err
	}

	return nil, err
}
