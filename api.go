package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initAPI() error {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	v1.Post("/tables/", handleTables)
	v1.Post("/tables/:name/records/", handleTable)
	v1.Put("/tables/:name/records/", handleInsertRecord)
	v1.Delete("/tables/:name/records/:id/", handleDeleteRecord)
	v1.Patch("/tables/:name/records/:id/:field/", handleEditRecord)

	return app.Listen(":3000")
}
