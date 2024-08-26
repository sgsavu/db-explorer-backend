package main

import (
	"github.com/sgsavu/sqlutils/v4"
)

type RecordRequestBody struct {
	Record sqlutils.TableRecord `json:"record"`
}

type EditRecordRequestBody struct {
	Record sqlutils.TableRecord `json:"record"`
	Update struct {
		Column string `json:"column"`
		Value  any    `json:"value"`
	} `json:"update"`
}

type DuplicateTableRequestBody struct {
	SourceTableName string `json:"sourceTableName"`
	NewTableName    string `json:"newTableName"`
}

type RenameTableRequestBody struct {
	NewTableName string `json:"newTableName"`
}
