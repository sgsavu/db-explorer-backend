package main

type ConnectIntent struct {
	Addr   string `json:"address"`
	DBName string `json:"dbName"`
	User   string `json:"user"`
	Passwd string `json:"password"`
}

type RequestBody struct {
	Connect ConnectIntent `json:"connect"`
}

type RecordRequestBody struct {
	Connect ConnectIntent `json:"connect"`
	Record  []interface{} `json:"record"`
}

type EditRecordRequestBody struct {
	Connect ConnectIntent `json:"connect"`
	Value   any           `json:"value"`
}

type DuplicateTableRequestBody struct {
	Connect         ConnectIntent `json:"connect"`
	SourceTableName string        `json:"sourceTableName"`
	NewTableName    string        `json:"newTableName"`
}

type RenameTableRequestBody struct {
	Connect      ConnectIntent `json:"connect"`
	NewTableName string        `json:"newTableName"`
}
