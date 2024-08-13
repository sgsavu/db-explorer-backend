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

	db, err := connectToDb(body.Connect)
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

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
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

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDeleteRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	columns, err := getColumns(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	_, err = removeRecord(db, tableName, columns, body.Record)
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

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleInsertRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
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

	_, err = addRecord(db, tableName, columns, body.Record)
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

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleEditRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(EditRecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	err = editRecord(db, tableName, body.RecordInfo.Column, body.RecordInfo.Value, body.Update.Column, body.Update.Value)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTable(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handlePrimaryKeys(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handlePrimaryKeys - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getPrimaryKeys(db, body.Connect.DBName, tableName)
	if err != nil {
		error := fmt.Sprintf("handlePrimaryKeys - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDuplicateTable(c *fiber.Ctx) error {
	body := new(DuplicateTableRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	err = duplicateTable(db, body.SourceTableName, body.NewTableName)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTables(db)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDeleteTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	err = deleteTable(db, tableName)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTables(db)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleRenameTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RenameTableRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := connectToDb(body.Connect)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	err = renameTable(db, tableName, body.NewTableName)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := getTables(db)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	defer db.Close()

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}
