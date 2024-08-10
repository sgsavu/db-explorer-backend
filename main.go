package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var db *sql.DB

type Connect struct {
	Addr   string `json:"address"`
	DBName string `json:"dbName"`
	User   string `json:"user"`
	Passwd string `json:"password"`
}

type DeleteRecord struct {
	Table string `json:"table"`
	ID    string `json:"id"`
}

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	DOB          string
	Balance      int64
	Address      string
	CreationDate string
}

type Transaction struct {
	ID          uuid.UUID
	SenderID    uuid.UUID
	RecipientID uuid.UUID
	Amount      float32
	Currency    string
	Timestamp   string
	Description string
}

type Message struct {
	Protocol string `json:"protocol"`
	Payload  any    `json:"payload"`
}

func getTables() ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println("Error fetching tables:", err)
		return nil, err
	}
	defer rows.Close()

	var tableNames []string

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}

	return tableNames, nil
}

func getTable(table string, resultType interface{}) (interface{}, error) {
	sliceType := reflect.SliceOf(reflect.TypeOf(resultType))
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	query := fmt.Sprintf("SELECT * FROM %s;", table)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		elemValue := reflect.New(reflect.TypeOf(resultType)).Elem()

		var fields []interface{}
		for i := 0; i < elemValue.NumField(); i++ {
			fields = append(fields, elemValue.Field(i).Addr().Interface())
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, fmt.Errorf("scan: %v", err)
		}

		sliceValue = reflect.Append(sliceValue, elemValue)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %v", err)
	}

	return sliceValue.Interface(), nil
}

func removeRecord(table string, id string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", table)

	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("removeRecord: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("removeRecord: %v", err)
	}

	return rowsAffected, nil
}

func connectToDb(connect Connect) {
	cfg := mysql.Config{
		User:   connect.User,
		Passwd: connect.Passwd,
		Net:    "tcp",
		Addr:   connect.Addr,
		DBName: connect.DBName,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func handleGetTable(tableName string) (any, error) {
	var err error
	var result any

	switch tableName {
	case "users":
		result, err = getTable(tableName, User{})
	case "transactions":
		result, err = getTable(tableName, Transaction{})
	default:
		log.Fatalf("Unknown table: %s", tableName)
		return nil, err
	}

	if err != nil {
		log.Println("error @ getTable:", err)
		return nil, err
	}

	return result, nil
}

func spinUpAPI() {
	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Use("/", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	v1.Get("/", websocket.New(func(c *websocket.Conn) {
		var (
			_   int
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			message := new(Message)

			if err := json.Unmarshal(msg, message); err != nil {
				log.Println("error @ unmarshal:", err)
				break
			}

			resp := Message{
				Protocol: message.Protocol,
				Payload:  nil,
			}

			log.Printf("msg: %s", msg)
			log.Printf("message: %s", message)

			switch message.Protocol {
			case "CONNECT":
				// Check if Payload is of type map[string]interface{}
				payloadMap, ok := message.Payload.(map[string]interface{})
				if !ok {
					log.Println("error: payload is not a valid map")
					break
				}

				// Convert the map to a Connect struct (assuming Connect is a struct)
				var connect Connect
				payloadBytes, _ := json.Marshal(payloadMap) // Convert map to JSON
				if err := json.Unmarshal(payloadBytes, &connect); err != nil {
					log.Println("error @ convert to Connect:", err)
					break
				}

				connectToDb(connect)

			case "GET_TABLES":
				tables, err := getTables()
				if err != nil {
					log.Println("error @ getTables:", err)
					break
				}
				resp.Payload = tables

			case "GET_TABLE":
				tableName, ok := message.Payload.(string)
				if !ok {
					log.Println("error: payload is not a valid string")
					break
				}

				result, err := handleGetTable(tableName)
				if err != nil {
					log.Println("error @ getTables:", err)
					break
				}

				resp.Payload = result

			case "DELETE_RECORD":
				// Check if Payload is of type map[string]interface{}
				payloadMap, ok := message.Payload.(map[string]interface{})
				if !ok {
					log.Println("error: payload is not a valid map")
					break
				}

				// Convert the map to a Connect struct (assuming Connect is a struct)
				var deleteRecord DeleteRecord
				payloadBytes, _ := json.Marshal(payloadMap) // Convert map to JSON
				if err := json.Unmarshal(payloadBytes, &deleteRecord); err != nil {
					log.Println("error @ convert to DeleteRecord:", err)
					break
				}

				rowsAffected, err := removeRecord(deleteRecord.Table, deleteRecord.ID)
				if err != nil {
					log.Println("error @ delete record:", err)
					break
				}

				resp.Payload = rowsAffected

				result, err := handleGetTable(deleteRecord.Table)
				if err != nil {
					log.Println("error @ getTables:", err)
					break
				}

				tableUpdate := Message{
					Protocol: "GET_TABLE",
					Payload:  result,
				}

				if err = c.WriteJSON(tableUpdate); err != nil {
					log.Println("write:", err)
					break
				}

			default:
				log.Println("unknown protocol:", message.Protocol)
			}

			if err = c.WriteJSON(resp); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	app.Listen(":3000")
}

func main() {
	spinUpAPI()
}
