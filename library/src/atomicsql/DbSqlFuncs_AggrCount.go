package atomicsql

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"
	"database/sql"
	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
)

//----------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------
var  SQL_CountInt TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt( arr []int ) int64{

	var min int64= 0
    for i := 0; i < len(arr); i++ {
            min ++;//= arr[i]
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_CountInt16 TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt16( arr []int16 ) int64{

	var min int64 = 0
    for i := 0; i < len(arr); i++ {
            min ++;//+= int32(arr[i]);
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_CountInt32 TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt32( arr []int32 ) int64{

	var min int64 = 0
    for i := 0; i < len(arr); i++ {
            min ++;//+= arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountInt64 TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt64( arr []int64 ) int64{

	var min int64 = 0
    for i := 0; i < len(arr); i++ {
            min ++;//+= arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountF64 TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountF64( arr []float64 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
            min ++;//+= arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountF32 TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountF32( arr []float32 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
            min ++;//+= arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountStr TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountStr( arr []string ) int64{

	var min int64 = 0;
    for i :=0; i < len(arr); i++ {
        
    	min ++;//+= arr[i]        
    }
	return min;
}
/*
//----------------------------------------------------------------------------------------------
var  SQL_CountDate TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountDate( arr []time.Time ) time.Time{

	
	return min;
}
*/

//----------------------------------------------------------------------------------------------
var  SQL_CountInt16N TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt16N( arr []sql.NullInt16 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min ++;//= int32(arr[i].Int16)
		}
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_CountInt32N TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt32N( arr []sql.NullInt32 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min ++;//+= int32(arr[i].Int32)
		}
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountInt64N TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountInt64N( arr []sql.NullInt64 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min ++;//+= (arr[i].Int64)
		}
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_CountF64N TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountF64N( arr []sql.NullFloat64 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min ++;//+= arr[i].Float64
		}
    }
	return min;
}

//----------------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------------
var  SQL_CountStrN TTextSql = TTextSql{ 
	Postgres: 	"COUNT( @x1@ )",
	Mysql: 		"COUNT( @x1@ )",
	Mssql: 		"COUNT( @x1@ )",
}
func Sql_CountStrN( arr []sql.NullString ) int64{

	var min int64 = 0;
    for i :=0; i < len(arr); i++ {
        if  arr[i].Valid {
            min ++;//+= arr[i].String
        }
    }
	return min;
}

