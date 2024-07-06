package handler

import (
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/vercel/go-bridge/go/bridge"
)

func Handler() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		message := os.Getenv("HELLO_MESSAGE")
		if message == "" {
			message = "Hello, World!"
		}
		return c.SendString(message)
	})

	bridge.Start(adaptor.FiberApp(app))
}
