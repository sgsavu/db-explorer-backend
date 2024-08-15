package main

import (
	"github.com/sgsavu/sqlutils"
)

type RequestBody struct {
	Connect sqlutils.SQLConnectionInfo `json:"connect"`
}

type RecordRequestBody struct {
	Connect sqlutils.SQLConnectionInfo `json:"connect"`
	Record  []interface{}              `json:"record"`
}

type EditRecordRequestBody struct {
	Connect sqlutils.SQLConnectionInfo `json:"connect"`
	Update  struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	} `json:"update"`
	RecordInfo struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	} `json:"recordInfo"`
	Value any `json:"value"`
}

type DuplicateTableRequestBody struct {
	Connect         sqlutils.SQLConnectionInfo `json:"connect"`
	SourceTableName string                     `json:"sourceTableName"`
	NewTableName    string                     `json:"newTableName"`
}

type RenameTableRequestBody struct {
	Connect      sqlutils.SQLConnectionInfo `json:"connect"`
	NewTableName string                     `json:"newTableName"`
}
