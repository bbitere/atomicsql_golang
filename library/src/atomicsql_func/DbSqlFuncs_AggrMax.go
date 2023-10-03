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
var  SQL_MaxInt TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxInt( arr []int ) int{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MaxInt16 TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxInt16( arr []int16 ) int16{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MaxInt32 TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxInt32( arr []int32 ) int32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxInt64 TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxInt64( arr []int64 ) int64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxF64 TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxF64( arr []float64 ) float64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxF32 TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxF32( arr []float32 ) float32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxStr TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxStr( arr []string ) string{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min < arr[i] {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxDate TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument
func Sql_MaxDate( arr []time.Time ) time.Time{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Compare(arr[i]) < 0 {
            min = arr[i]
        }
    }
	return min;
}




//----------------------------------------------------------------------------------------------
var  SQL_MaxInt16N TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxInt16N( arr []sql.NullInt16 ) sql.NullInt16{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Int16 > arr[i].Int16 {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MaxInt32N TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxInt32N( arr []sql.NullInt32 ) sql.NullInt32{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {		
        if  min.Valid && arr[i].Valid && min.Int32 > arr[i].Int32 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxInt64N TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxInt64N( arr []sql.NullInt64 ) sql.NullInt64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Int64 > arr[i].Int64 {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxF64N TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxF64N( arr []sql.NullFloat64 ) sql.NullFloat64{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.Float64 > arr[i].Float64 {
            min = arr[i]
        }
    }
	return min;
}
//----------------------------------------------------------------------------------------------
var  SQL_MaxStrN TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxStrN( arr []sql.NullString ) sql.NullString{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
        if  min.Valid && arr[i].Valid && min.String > arr[i].String {
            min = arr[i]
        }
    }
	return min;
}

//----------------------------------------------------------------------------------------------
var  SQL_MaxDateN TTextSql = TTextSql{ 
	Postgres: 	"MAX( @x1@ )",
	Mysql: 		"MAX( @x1@ )",
	Mssql: 		"MAX( @x1@ )",
}
//returns the maximum value from the list received as 'arr' argument, only for valid items
func Sql_MaxDateN( arr []sql.NullTime ) sql.NullTime{

	var min = arr[0]
    for i :=1; i < len(arr); i++ {
		if  min.Valid && arr[i].Valid && min.Time.Compare( arr[i].Time ) < 0 {
        //if  min.Compare(arr[i]) > 0 {
            min = arr[i]
        }
    }
	return min;
}