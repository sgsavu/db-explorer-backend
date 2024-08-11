package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func handleWebSocket(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read: %v", err)
			break
		}

		var request Request
		if err := json.Unmarshal(message, &request); err != nil {
			log.Printf("unmarshal: %v", err)
			break
		}

		response := Response{
			Type: request.Type,
		}

		switch request.Type {
		case "CONNECT":
			var connect ConnectIntent
			if err := mapToStruct(request.Payload, &connect); err != nil {
				log.Printf("CONNECT - mapToStruct - %v", err)
				response.StatusCode = fiber.StatusBadRequest
				break
			}

			if err := connectToDb(connect); err != nil {
				log.Printf("CONNECT - connectToDb - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = request.Payload
			response.StatusCode = fiber.StatusOK

		case "GET_TABLES":
			result, err := getTables(db)
			if err != nil {
				log.Printf("GET_TABLES - getTables - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

		case "GET_TABLE":
			table, ok := request.Payload.(string)
			if !ok {
				log.Println("GET_TABLE - payload is not a valid string")
				response.StatusCode = fiber.StatusBadRequest
				break
			}

			result, err := getTable(db, table)
			if err != nil {
				log.Printf("GET_TABLE - getTable - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

		case "DELETE_RECORD":
			var deleteRecordIntent DeleteRecordIntent
			if err := mapToStruct(request.Payload, &deleteRecordIntent); err != nil {
				log.Printf("DELETE_RECORD - mapToStruct - %v", err)
				response.StatusCode = fiber.StatusBadRequest
				break
			}

			result, err := removeRecord(db, deleteRecordIntent.Table, deleteRecordIntent.ID)
			if err != nil {
				log.Printf("DELETE_RECORD - removeRecord - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

			sendTableUpdate(c, deleteRecordIntent.Table)

		case "INSERT_RECORD":
			var addRecordIntent AddRecordIntent
			if err := mapToStruct(request.Payload, &addRecordIntent); err != nil {
				log.Printf("INSERT_RECORD - mapToStruct - %v", err)
				response.StatusCode = fiber.StatusBadRequest
				break
			}

			columns, err := getColumns(db, addRecordIntent.Table)
			if err != nil {
				log.Printf("INSERT_RECORD - getColumns - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			result, err := addRecord(db, addRecordIntent.Table, columns, addRecordIntent.Record)
			if err != nil {
				log.Printf("INSERT_RECORD - addRecord - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

			sendTableUpdate(c, addRecordIntent.Table)

		default:
			log.Printf("unknown request type: %s", request.Type)
		}

		if err := c.WriteJSON(response); err != nil {
			log.Printf("sending response: %v", err)
			break
		}
	}
}

func sendTableUpdate(c *websocket.Conn, table string) {
	result, err := getTable(db, table)
	if err != nil {
		log.Printf("getTable: %v", err)
	}

	response := Response{
		Payload:    result,
		StatusCode: fiber.StatusOK,
		Type:       "GET_TABLE",
	}

	if err := c.WriteJSON(response); err != nil {
		log.Printf("sending response: %v", err)
	}
}
