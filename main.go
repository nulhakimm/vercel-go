package main

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {

	//Create fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := os.Getenv("NAME")
		if name == "" {
			name = "World"
		}
		return c.SendString("Hello " + name)
	})

	adaptor.FiberApp(app)(w, r)
}

func main() {
	// Create a simple server
	http.HandleFunc("/", Handler)

	// Start the server on port 3000
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
