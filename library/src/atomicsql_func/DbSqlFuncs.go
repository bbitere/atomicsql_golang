package atomicsql_func

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"
	"database/sql"

	atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"

	time "time"
)

type TTextSql struct{

	Postgres string
	Mysql string
	Mssql string

}


//----------------------------------------------------------------------------------------------
var  SQL_IIF TTextSql = TTextSql{ 
	Postgres: "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
	Mysql:    "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
	Mssql:    "(CASE WHEN (@x1@) THEN @x2@ ELSE @x3@ END)",
}
//Sql_IIF() - it replace the conditional operator 
func Sql_IIF[T any](  b bool, s1 T, s2 T) T {

	if( b){
		return s1
	}else {
		return s2
	}
}
//----------------------------------------------------------------------------------------------
// not used
func Sql_AA[T any]( ctx atomicsql.IDBContext)  {
}


//----------------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------------
var  SQL_IsNilI08 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilI08( val sql.NullByte, defval byte) byte {

	if( val.Valid ){
		return (val.Byte);
	}else{
		return defval;
	}	
}
//----------------------------------------------------------------------------------------------
var  SQL_IsNilI16 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}

//return val or defval if val is nil
func Sql_IsNilI16( val sql.NullInt16, defval int16) int16 {

	if( val.Valid ){
		return val.Int16;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------
var  SQL_IsNilI32 TTextSql = TTextSql{ 
	Postgres: 	"NULLIF(@x1@, @x2@)",
	Mysql:    	"ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilI32( val sql.NullInt32, defval int32) int32 {

	if( val.Valid ){
		return val.Int32;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------
var  SQL_IsNilI64 TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilI64( val sql.NullInt64, defval int64) int64 {

	if( val.Valid ){
		return val.Int64;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------
var  SQL_IsNilStr TTextSql = TTextSql{ 
	Postgres: 	"NULLIF(@x1@, @x2@)",
	Mysql:    	"ISNULL(@x1@, @x2@)",
	Mssql:  	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilStr( val sql.NullString, defval string) string {

	if( val.Valid ){
		return val.String;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------
var  SQL_IsNilBool TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilBool( val sql.NullBool, defval bool) bool {

	if( val.Valid ){
		return val.Bool;
	}else{
		return defval;
	}	
}

//----------------------------------------------------------------------------------------------
var  SQL_IsNilDate TTextSql = TTextSql{ 
	Postgres: "NULLIF(@x1@, @x2@)",
	Mysql:    "ISNULL(@x1@, @x2@)",
	Mssql:   	"ISNULL(@x1@, @x2@)",
}
//return val or defval if val is nil
func Sql_IsNilDate( val sql.NullTime, defval time.Time) time.Time {

	if( val.Valid ){
		return val.Time;
	}else{
		return defval;
	}	
}

