package main

import (
	"app-backend/internal/database"
	"app-backend/internal/middlewares"
	"app-backend/internal/router"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Settings of app
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  3 * time.Second,
	})

	// Cors middleware
	app.Use(cors.New())
	// Custom Timer middleware
	app.Use(middlewares.Timer())
	// Default middleware
	app.Use(pprof.New())

	// Open connection with DB
	database.Connect()

	// Migrations
	database.AutoMigrate()
	fmt.Println("Database Migrated")

	// Setup routes
	router.SetupRoutes(app)
	fmt.Println("Routes Setup")

	// Serve files from multiple directories with custom config
	app.Static("/", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Index:     "index.html",
	})

	// Serve files from "./files" directory:
	app.Static("/", "./files")

	// Start app
	err := app.Listen("localhost:3000")
	if err != nil {
		log.Fatal()
	}
}
