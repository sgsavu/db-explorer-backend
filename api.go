package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initAPI() error {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	// v1.Use("/", func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		c.Locals("allowed", true)
	// 		return c.Next()
	// 	}
	// 	return fiber.ErrUpgradeRequired
	// })

	// v1.Get("/", websocket.New(handleWebSocket))

	v1.Post("/tables/", handleTables)
	v1.Post("/tables/:name/", handleTable)
	v1.Post("/tables/:name/", handleInsertRecord)
	v1.Delete("/tables/:name/:id", handleDeleteRecord)
	// v1.Put("/tables/:name/entries/:id", handleDeleteRecord)

	return app.Listen(":3000")
}
