package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func initAPI(port *string) error {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")

	v1.Get("/tables/", handleTables)
	v1.Put("/tables/", handleDuplicateTable)

	v1.Patch("/tables/:name/", handleRenameTable)
	v1.Delete("/tables/:name/", handleDeleteTable)

	v1.Get("/tables/:name/columns/", handleColumns)
	v1.Get("/tables/:name/primary-keys/", handlePrimaryKeys)

	v1.Get("/tables/:name/records/", handleRecords)
	v1.Post("/tables/:name/records/", handleInsertRecord)
	v1.Patch("/tables/:name/records/", handleEditRecord)
	v1.Delete("/tables/:name/records/", handleRemoveRecord)

	return app.Listen(fmt.Sprintf(":%s", *port))
}
