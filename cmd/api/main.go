package main

import (
	"fmt"
	"github.com/YaroslavVlasenko/flapi/internal/database"
	"github.com/YaroslavVlasenko/flapi/internal/middlewares"
	"github.com/YaroslavVlasenko/flapi/internal/router"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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
		ReadTimeout:   7 * time.Second,
		WriteTimeout:  3 * time.Second,
		AppName: "Test Application v1.0.1",
	})

	// Cors middleware
	app.Use(cors.New())
	// Compression level middleware
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	// Custom Timer middleware
	app.Use(middlewares.Timer())
	// Default middleware
	app.Use(pprof.New())
	// CSRF middleware
	//app.Use(csrf.New())
	// Favicon middleware
	app.Use(favicon.New())

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
