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
			var connect Connect
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
			result, err := getTables()
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

			result, err := getTable(table)
			if err != nil {
				log.Printf("GET_TABLE - getTable - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

		case "DELETE_RECORD":
			var deleteRecord DeleteRecord
			if err := mapToStruct(request.Payload, &deleteRecord); err != nil {
				log.Printf("DELETE_RECORD - mapToStruct - %v", err)
				response.StatusCode = fiber.StatusBadRequest
				break
			}

			result, err := removeRecord(deleteRecord.Table, deleteRecord.ID)
			if err != nil {
				log.Printf("DELETE_RECORD - removeRecord - %v", err)
				response.StatusCode = fiber.StatusInternalServerError
				break
			}

			response.Payload = result
			response.StatusCode = fiber.StatusOK

			sendTableUpdate(c, deleteRecord.Table)

		// case "INSERT_RECORD":

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
	result, err := getTable(table)
	if err != nil {
		log.Printf("getTable: %v", err)
	}

	response := Response{
		Type:    "GET_TABLE",
		Payload: result,
	}

	if err := c.WriteJSON(response); err != nil {
		log.Printf("sending response: %v", err)
	}
}
