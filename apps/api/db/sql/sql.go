package sql

import _ "embed"

//go:embed schema.sql
var schema string

func GetSchema() string {
	return schema
}
