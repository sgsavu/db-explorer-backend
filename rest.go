package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sgsavu/sqlutils/v4"
)

func handleTables(c *fiber.Ctx) error {
	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleTables - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetTables(db, connectionInfo.Name, connectionInfo.Type)
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

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleColumns - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetColumns(db, tableName, connectionInfo.Type)
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

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handlePrimaryKeys - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetPrimaryKeys(db, connectionInfo.Name, tableName, connectionInfo.Type)
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

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	result, err := sqlutils.GetTable(db, tableName, connectionInfo.Type)
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
	tableName := c.Params("name")

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(DuplicateTableRequestBody)

	bodyBytes := c.Body()
	if len(bodyBytes) != 0 {
		err := c.BodyParser(body)
		if err != nil {
			return err
		}
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.DuplicateTable(db, tableName, body.NewTableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, connectionInfo.Name, connectionInfo.Type)
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

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.DeleteTable(db, tableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDeleteTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, connectionInfo.Name, connectionInfo.Type)
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

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(RenameTableRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.RenameTable(db, tableName, body.NewTableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTables(db, connectionInfo.Name, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleRenameTable - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleInsertRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	_, err = sqlutils.InsertRecord(db, tableName, body.Record, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleInsertRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleDuplicateRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.DuplicateRecord(db, connectionInfo.Name, tableName, body.Record, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDuplicateRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleEditRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(EditRecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	err = sqlutils.EditRecord(db, tableName, body.Record, body.Update.Column, body.Update.Value, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleEditRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}

func handleRemoveRecord(c *fiber.Ctx) error {
	tableName := c.Params("name")

	connectionInfo := new(sqlutils.DBConnection)
	if err := c.ReqHeaderParser(connectionInfo); err != nil {
		return err
	}

	body := new(RecordRequestBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	db, err := sqlutils.ConnectDB(connectionInfo)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}
	defer db.Close()

	_, err = sqlutils.RemoveRecord(db, connectionInfo.Name, tableName, connectionInfo.Type, body.Record)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	result, err := sqlutils.GetTable(db, tableName, connectionInfo.Type)
	if err != nil {
		error := fmt.Sprintf("handleDeleteRecord - %v", err)
		log.Println(error)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"error": error})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"result": result})
}
