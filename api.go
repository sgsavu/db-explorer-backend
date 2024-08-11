package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initAPI() error {
	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Use("/", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	v1.Get("/", websocket.New(handleWebSocket))

	return app.Listen(":3000")
}
