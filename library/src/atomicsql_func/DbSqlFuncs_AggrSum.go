package atomicsql_func

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

// internal use
var  SQL_SumInt TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
// returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt( arr []int ) int{

	var min = 0
    for i := 0; i < len(arr); i++ {
            min += arr[i]
    }
	return min;
}
//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt16 TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt16( arr []int16 ) int32{

	var min int32 = 0
    for i := 0; i < len(arr); i++ {
            min += int32(arr[i]);
    }
	return min;
}
//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt32 TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt32( arr []int32 ) int32{

	var min int32 = 0
    for i := 0; i < len(arr); i++ {
            min += arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt64 TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt64( arr []int64 ) int64{

	var min int64 = 0
    for i := 0; i < len(arr); i++ {
            min += arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumF64 TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumF64( arr []float64 ) float64{

	var min float64 = 0
    for i :=0; i < len(arr); i++ {
            min += arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumF32 TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumF32( arr []float32 ) float32{

	var min float32 = 0
    for i :=0; i < len(arr); i++ {
            min += arr[i]
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_ConcatStr TTextSql = TTextSql{ 
	Postgres: 	"CONCAT( @x1@ )",
	Mysql: 		"CONCAT( @x1@ )",
	Mssql: 		"CONCAT( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_ConcatStr( arr []string ) string{

	var min = "";
    for i :=0; i < len(arr); i++ {
        
    	min += arr[i]        
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt16N TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument, only for valid items
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt16N( arr []sql.NullInt16 ) int32{

	var min int32 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min += int32(arr[i].Int16)
		}
    }
	return min;
}
//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt32N TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument, only for valid items
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt32N( arr []sql.NullInt32 ) int32{

	var min int32 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min += int32(arr[i].Int32)
		}
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumInt64N TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument, only for valid items
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumInt64N( arr []sql.NullInt64 ) int64{

	var min int64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min += (arr[i].Int64)
		}
    }
	return min;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_SumF64N TTextSql = TTextSql{ 
	Postgres: 	"SUM( @x1@ )",
	Mysql: 		"SUM( @x1@ )",
	Mssql: 		"SUM( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument, only for valid items
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_SumF64N( arr []sql.NullFloat64 ) float64{

	var min float64 = 0
    for i :=0; i < len(arr); i++ {
		if  arr[i].Valid {
            min += arr[i].Float64
		}
    }
	return min;
}
//----------------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------------

// internal use
var  SQL_ConcatStrN TTextSql = TTextSql{ 
	Postgres: 	"CONCAT( @x1@ )",
	Mysql: 		"CONCAT( @x1@ )",
	Mssql: 		"CONCAT( @x1@ )",
}
//returns the sum of values from the list received as 'arr' argument, only for valid items
// 
// this method is translated in sql query, when it is using in sequence of Select( Aggregate() ) methods, to aggregate data
func Sql_ConcatStrN( arr []sql.NullString ) string{

	var min = "";
    for i :=0; i < len(arr); i++ {
        if  arr[i].Valid {
            min += arr[i].String
        }
    }
	return min;
}

