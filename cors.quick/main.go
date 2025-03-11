package main

import (
	"github.com/jeffotoni/quick"
	"github.com/jeffotoni/quick/middleware/cors"
)

func main() {
	q := quick.New()

	q.Use(cors.New(cors.Config{
		AllowedOrigins:   []string{"*"}, // Allow any origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow any header
		ExposedHeaders:   []string{"*"}, // Show any header
		AllowCredentials: true,          // Allow cookies and authentication via CORS
		Debug:            true,
	}))

	q.Get("/v1/user/:id", func(c *quick.Ctx) error {
		c.Set("Content-type", "application/json")
		return c.Status(quick.StatusOK).String("Hello, Quick in action!!")
	})

	// start dir files
	q.Static("/static", "./static")

	// server files
	q.Get("/", func(c *quick.Ctx) error {
		c.File("./static/index.html")
		return nil
	})

	q.Listen(":8080")
}
