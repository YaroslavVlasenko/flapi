package router

import (
	"github.com/YaroslavVlasenko/flapi/internal/handlers"
	"github.com/YaroslavVlasenko/flapi/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// SetupRoutes setup router `api
func SetupRoutes(app *fiber.App) {
	// Monitor Middleware page
	app.Get("/dashboard", monitor.New())

	// Middleware
	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")
	v1.Get("/", handlers.Hello)

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	user := v1.Group("/user")
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Patch("/:id", middlewares.Protected(), handlers.UpdateUser)
	user.Delete("/:id", middlewares.Protected(), handlers.DeleteUser)

	// Product
	product := v1.Group("/product")
	product.Get("/all", handlers.GetAllProducts)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", middlewares.Protected(), handlers.CreateProduct)
	product.Patch("/:id", middlewares.Protected(), handlers.UpdateProduct)
	product.Delete("/:id", middlewares.Protected(), handlers.DeleteProduct)
}
