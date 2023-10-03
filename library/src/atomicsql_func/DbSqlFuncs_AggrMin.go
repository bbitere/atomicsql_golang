package atomicsql_func

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"
	"database/sql"

	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"

	time "time"
)

//----------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------
var  SQL_MinInt TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinInt( arr []int ) int{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MinInt16 TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinInt16( arr []int16 ) int16{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MinInt32 TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinInt32( arr []int32 ) int32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinInt64 TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinInt64( arr []int64 ) int64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinF64 TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinF64( arr []float64 ) float64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinF32 TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinF32( arr []float32 ) float32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinStr TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinStr( arr []string ) string{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min > arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinDate TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument
func Sql_MinDate( arr []time.Time ) time.Time{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Compare(arr[i]) > 0 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinInt16N TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinInt16N( arr []sql.NullInt16 ) sql.NullInt16{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Int16 > arr[i].Int16 {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MinInt32N TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinInt32N( arr []sql.NullInt32 ) sql.NullInt32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {		
        if  min.Valid && arr[i].Valid && min.Int32 > arr[i].Int32 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinInt64N TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinInt64N( arr []sql.NullInt64 ) sql.NullInt64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Int64 > arr[i].Int64 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinF64N TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinF64N( arr []sql.NullFloat64 ) sql.NullFloat64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Float64 > arr[i].Float64 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinStrN TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinStrN( arr []sql.NullString ) sql.NullString{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.String > arr[i].String {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MinDateN TTextSql = TTextSql{ 
	Postgres: 	"MIN( @x1@ )",
	Mysql: 		"MIN( @x1@ )",
	Mssql: 		"MIN( @x1@ )",
}
//returns the minum value from the list received as 'arr' argument, only for valid items
func Sql_MinDateN( arr []sql.NullTime ) sql.NullTime{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
		if  min.Valid && arr[i].Valid && min.Time.Compare( arr[i].Time ) > 0 {
        //if  min.Compare(arr[i]) > 0 {
            min = arr[i]
        }
    }
	return min;
}