package main

import (
	"github.com/google/uuid"
)

type ConnectIntent struct {
	Addr   string `json:"address"`
	DBName string `json:"dbName"`
	User   string `json:"user"`
	Passwd string `json:"password"`
}

type DeleteRecordIntent struct {
	Table string `json:"table"`
	ID    string `json:"id"`
}

type AddRecordIntent struct {
	Table  string        `json:"table"`
	Record []interface{} `json:"record"`
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

type Request struct {
	Payload interface{} `json:"payload"`
	Type    string      `json:"type"`
}

type Response struct {
	Payload    interface{} `json:"payload"`
	StatusCode int64       `json:"statusCode"`
	Type       string      `json:"type"`
}

type RequestBody struct {
	Connect ConnectIntent `json:"connect"`
}

type InsertRecordRequestBody struct {
	Connect ConnectIntent `json:"connect"`
	Record  []interface{} `json:"record"`
}
