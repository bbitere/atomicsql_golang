package atomicsql_func

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"
	"database/sql"

	time "time"
)

type TTextSql struct{

	Postgres string
	Mysql string
	Mssql string

}


//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IIF TTextSql = TTextSql{ 
	Postgres: "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
	Mysql:    "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
	Mssql:    "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
}
// Sql_IIF() - it replace the conditional operator 
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
// Example:
//  var count = ctx.Table.Qry("tag1").Where( func(x *Table) bool { 
//  					return SQL_IFF( x.Table.Field1 != nil, x.Table.Field1, "") != "" }).GetCount()
//  })
// It will be translate in sql as:
//  Select COUNT(*) FROM Table WHERE (CASE WHEN (Field1 IS NOT NULL ) THEN Field1 ELSE '' END) <> ''
// 
func Sql_IIF[T any](  b bool, s1 T, s2 T) T {

	if( b){
		return s1
	}else {
		return s2
	}
}

//----------------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilI08 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilI08( val sql.NullByte, defval byte) byte {

	if( val.Valid ){
		return (val.Byte);
	}else{
		return defval;
	}	
}
//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilI16 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}

//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilI16( val sql.NullInt16, defval int16) int16 {

	if( val.Valid ){
		return val.Int16;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilI32 TTextSql = TTextSql{ 
	Postgres: 	"NULLIF(@x1@, @x2@)",
	Mysql:    	"ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilI32( val sql.NullInt32, defval int32) int32 {

	if( val.Valid ){
		return val.Int32;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilI64 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
// return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilI64( val sql.NullInt64, defval int64) int64 {

	if( val.Valid ){
		return val.Int64;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilStr TTextSql = TTextSql{ 
	Postgres: 	"NULLIF(@x1@, @x2@)",
	Mysql:    	"ISNULL(@x1@, @x2@)",
	Mssql:  	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilStr( val sql.NullString, defval string) string {

	if( val.Valid ){
		return val.String;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilBool TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilBool( val sql.NullBool, defval bool) bool {

	if( val.Valid ){
		return val.Bool;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_IsNilDate TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_IsNilDate( val sql.NullTime, defval time.Time) time.Time {

	if( val.Valid ){
		return val.Time;
	}else{
		return defval;
	}	
}

