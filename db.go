package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func connectToDb(connect ConnectIntent) (*sql.DB, error) {
	cfg := mysql.Config{
		User:   connect.User,
		Passwd: connect.Passwd,
		Net:    "tcp",
		Addr:   connect.Addr,
		DBName: connect.DBName,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("connectToDb - opening db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("connectToDb - pinging db: %w", err)
	}

	return db, nil
}

func getTables(db *sql.DB) ([]string, error) {
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

func getTable(db *sql.DB, tableName string) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s;", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("getTable - query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("getTable - retrieving columns: %w", err)
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

func removeRecord(db *sql.DB, tableName, id string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", tableName)
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("deleting record: %w", err)
	}
	return result.RowsAffected()
}

func getColumns(db *sql.DB, tableName string) ([]string, error) {
	query := `
		SELECT COLUMN_NAME
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE()
		ORDER BY ORDINAL_POSITION;
	`

	rows, err := db.Query(query, tableName)
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

func addRecord(db *sql.DB, tableName string, columns []string, values []interface{}) (int64, error) {
	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return 0, fmt.Errorf("invalid columns or values length")
	}

	placeholders := strings.Repeat("?, ", len(values)-1) + "?"

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
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

func editRecord(db *sql.DB, tableName string, field string, value any, recordId string) error {
	query := fmt.Sprintf("UPDATE %s SET %s = ?  WHERE ID = ?", tableName, field)

	result, err := db.Exec(query, value, recordId)
	if err != nil {
		return fmt.Errorf("editRecord: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated, check if the employee_id exists")
	}

	return nil
}
