package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type PaginationInfo struct {
	XPaginationCurrentPage int    `json:"x_pagination_current_page"`
	XPaginationPerPage     int    `json:"x_pagination_per_page"`
	XPaginationTotalCount  int    `json:"x_pagination_total_count"`
	XPaginationPageCount   int    `json:"x_pagination_page_count"`
}

func GeneratePaginationFromRequest(c *fiber.Ctx) Pagination {
	// Initializing default
	limit, err := strconv.Atoi(c.Query("per-page", "10"))
	if err != nil {
		fmt.Print("Error loading param: per-page, from context.")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		fmt.Print("Error loading param: page, from context.")
	}

	sort := c.Query("sort", "id")
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func GeneratePaginationInfo(c int64, p Pagination) PaginationInfo {
	pc := float64(c) / float64(p.Limit)
	return PaginationInfo{
		XPaginationCurrentPage: p.Page,
		XPaginationPerPage:     p.Limit,
		XPaginationTotalCount:  int(c),
		XPaginationPageCount:   int(math.Ceil(pc)),
	}
}
