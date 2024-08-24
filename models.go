package main

import (
	"github.com/sgsavu/sqlutils/v2"
)

type RequestBody struct {
	ConnectionInfo sqlutils.SQLConnectionInfo `json:"connectionInfo"`
}

type RecordRequestBody struct {
	ConnectionInfo sqlutils.SQLConnectionInfo `json:"connectionInfo"`
	Record         []interface{}              `json:"record"`
}

type EditRecordRequestBody struct {
	ConnectionInfo sqlutils.SQLConnectionInfo `json:"connectionInfo"`
	Update         struct {
		Column string `json:"column"`
		Value  any    `json:"value"`
	} `json:"update"`
	RecordInfo struct {
		Column string `json:"column"`
		Value  any    `json:"value"`
	} `json:"recordInfo"`
}

type DuplicateTableRequestBody struct {
	ConnectionInfo  sqlutils.SQLConnectionInfo `json:"connectionInfo"`
	SourceTableName string                     `json:"sourceTableName"`
	NewTableName    string                     `json:"newTableName"`
}

type RenameTableRequestBody struct {
	ConnectionInfo sqlutils.SQLConnectionInfo `json:"connectionInfo"`
	NewTableName   string                     `json:"newTableName"`
}
