package main

import (
	"github.com/google/uuid"
)

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

type Request struct {
	Payload interface{} `json:"payload"`
	Type    string      `json:"type"`
}

type Response struct {
	Payload    interface{} `json:"payload"`
	StatusCode int64       `json:"statusCode"`
	Type       string      `json:"type"`
}
