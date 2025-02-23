package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Custom Middleware for logging
func Logger(c *fiber.Ctx) error {
	fmt.Println("Request URL:", c.OriginalURL())
	return c.Next() // Proceed to the next handler
}

func main() {
	app := fiber.New()

	// Apply Middleware
	app.Use(Logger)

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Middleware!")
	})

	app.Listen(":3000")
}
