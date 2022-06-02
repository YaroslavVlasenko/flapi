package middlewares

import (
	"github.com/YaroslavVlasenko/flapi/internal/configs"
	"github.com/YaroslavVlasenko/flapi/internal/responses"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(configs.Load("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(responses.Error("error", "Missing or malformed JWT", err.Error()))
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(responses.Error("error", "Invalid or expired JWT", err.Error()))
}
