package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func handleTables(c *fiber.Ctx) error {
	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleTables - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTables(db)
	if err != nil {
		error := fmt.Sprintf("handleTables - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTable(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDeleteRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")
	recordId := c.Params("id")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	_, err = removeRecord(db, tableName, recordId)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTable(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleInsertRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(InsertRecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	columns, err := getColumns(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	_, err = addRecord(db, tableName, columns, body.Payload)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTable(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}
