package main

import (
	"livestreaming/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Streaming",
	})

	// CORS settings for cross-origin requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	// Register the /streaming routes
	streamingGroup := app.Group("/")
	handler.Streaming(streamingGroup)

	// Start the server and listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
