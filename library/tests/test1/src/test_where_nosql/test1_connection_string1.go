package test1_where

import (
	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)



func Test1_GetConnectionString() atmsql.TConnectionString{

	var connString = atmsql.TConnectionString{
		Host:     "localhost",		
		User:     "test1usr",
		Password: "123456",
		DbName:   "test1",
		SqlLang:  atmsql.ESqlDialect.MongoDB,
	}
	return connString;
}
