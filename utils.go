package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func mapToStruct(input interface{}, output interface{}) error {
	bytes, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("marshal input: %w", err)
	}
	if err := json.Unmarshal(bytes, output); err != nil {
		return fmt.Errorf("unmarshal to output: %w", err)
	}
	return nil
}

func createFields(columns []string) []reflect.StructField {
	var fields []reflect.StructField
	for _, col := range columns {
		fields = append(fields, reflect.StructField{
			Name: col,
			Type: reflect.TypeOf(""),
		})
	}
	return fields
}
