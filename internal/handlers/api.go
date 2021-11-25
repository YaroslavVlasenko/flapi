package handlers

import (
	"app-backend/internal/responses"
	"github.com/gofiber/fiber/v2"
)

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(responses.Success("success", "All fine!", nil))
}
