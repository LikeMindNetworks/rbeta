package model

import (
	"strings"
)

const (
	SEP string = "-"
)

// Name related to reducers
// - source: ddb table
// - reducer functions: lambda functions

MakeDDBName(ns string, tableName string) string {
}

MakeReducerName(ns string, tableName string, reducerName string) string {
}
