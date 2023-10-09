package test1_where

import (
	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

func Test1_GetConnectionString() atmsql.TConnectionString {

	var connString = atmsql.TConnectionString{
		Host:     "localhost",
		Port:     3306,
		User:     "rbitere",
		Password: "xz021621",
		DbName:   "test1",
	}
	return connString
}
