package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sgsavu/sqlutils/v2"
)

func handleTables(c *fiber.Ctx) error {
	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleTables - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetTables(db, body.ConnectionInfo.DBName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleTables - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleColumns(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleColumns - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetColumns(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleColumns - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handlePrimaryKeys(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handlePrimaryKeys - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetPrimaryKeys(db, body.ConnectionInfo.DBName, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handlePrimaryKeys - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleRecords(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetTable(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDuplicateTable(c *fiber.Ctx) error {
	body := new(DuplicateTableRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.DuplicateTable(db, body.SourceTableName, body.NewTableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, body.ConnectionInfo.DBName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDeleteTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.DeleteTable(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, body.ConnectionInfo.DBName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleRenameTable(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RenameTableRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.RenameTable(db, tableName, body.NewTableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, body.ConnectionInfo.DBName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleRemoveRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	_, err = sqlutils.RemoveRecord(db, body.ConnectionInfo.DBName, tableName, body.ConnectionInfo.Type, body.Record)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, body.ConnectionInfo.Type)
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

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	columns, err := sqlutils.GetColumns(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	_, err = sqlutils.InsertRecord(db, tableName, columns, body.Record, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleEditRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	body := new(EditRecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(&body.ConnectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.EditRecord(db, tableName, body.RecordInfo.Column, body.RecordInfo.Value, body.Update.Column, body.Update.Value, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, body.ConnectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}
