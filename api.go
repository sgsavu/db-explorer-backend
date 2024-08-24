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

	v1.Post("/tables/", handleTables)
	v1.Put("/tables/", handleDuplicateTable)

	v1.Delete("/tables/:name/", handleDeleteTable)
	v1.Patch("/tables/:name/", handleRenameTable)

	v1.Post("/tables/:name/columns/", handleColumns)
	v1.Post("/tables/:name/primary-keys/", handlePrimaryKeys)
	v1.Post("/tables/:name/records/", handleRecords)

	v1.Put("/tables/:name/records/", handleInsertRecord)
	v1.Delete("/tables/:name/records/", handleRemoveRecord)
	v1.Patch("/tables/:name/records/", handleEditRecord)

	return app.Listen(fmt.Sprintf(":%s", *port))
}
