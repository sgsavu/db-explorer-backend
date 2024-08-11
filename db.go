package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectToDb(connect Connect) error {
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
		return fmt.Errorf("opening db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("pinging db: %w", err)
	}

	return nil
}

func getTables() ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("fetching tables: %w", err)
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}
	return tableNames, rows.Err()
}

func getTable(table string) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s;", table)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("retrieving columns: %w", err)
	}

	sliceType := reflect.SliceOf(reflect.StructOf(createFields(columns)))
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	for rows.Next() {
		elemValue := reflect.New(sliceType.Elem()).Elem()

		var fields []interface{}
		for i := 0; i < elemValue.NumField(); i++ {
			fields = append(fields, elemValue.Field(i).Addr().Interface())
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}

		sliceValue = reflect.Append(sliceValue, elemValue)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return sliceValue.Interface(), nil
}

func removeRecord(table, id string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", table)
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("deleting record: %w", err)
	}
	return result.RowsAffected()
}

func getColumns(db *sql.DB, table string) ([]string, error) {
	query := `
		SELECT COLUMN_NAME
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE()
		ORDER BY ORDINAL_POSITION;
	`

	rows, err := db.Query(query, table)
	if err != nil {
		return nil, fmt.Errorf("getColumns: %v", err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, fmt.Errorf("getColumns: %v", err)
		}
		columns = append(columns, column)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getColumns: %v", err)
	}

	return columns, nil
}

func addRecord(db *sql.DB, table string, columns []string, values []interface{}) (int64, error) {
	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return 0, fmt.Errorf("invalid columns or values length")
	}

	placeholders := strings.Repeat("?, ", len(values)-1) + "?"

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		placeholders,
	)

	result, err := db.Exec(query, values...)
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addRecord: %v", err)
	}

	return id, nil
}
