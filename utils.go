package main

import (
	"reflect"
)

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
