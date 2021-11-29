package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GeneratePaginationFromRequest(c *fiber.Ctx) Pagination {
	// Initializing default
	limit, err := strconv.Atoi(c.Query("per-page", "20"))
	if err != nil {
		fmt.Print("Error loading param: per-page, from context.")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		fmt.Print("Error loading param: page, from context.")
	}

	sort := c.Query("sort", "created_at")
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
