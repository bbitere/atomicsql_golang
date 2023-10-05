package atomicsql_func

import (
	"strconv"
	"strings"

	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
	"regexp"
	Str "strings"

	"github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_Like TTextSql = TTextSql{ 
	Postgres: "@x1@ like '@x2@'",
	Mysql:    "@x1@ like '@x2@'",
	Mssql:    "@x1@ like '@x2@'",
}
// it is similar with LIKE operator from sql
// Example:
//  var count = ctx.Table.Qry("tag1").Where( func(x *Table) bool { 
//  					return SQL_Like( x.Table.Field1, "%acasa%") }).GetCount()
//  })
// It will be translate in sql as:
//  Select COUNT(*) FROM Table WHERE Field1 like '%acasa%'
// if( userName LIKE '%jing%') -> will return all rows that userName contains `jing`
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_Like( val string, likeFormat string) bool {

	var regExFormat = strings.ReplaceAll( likeFormat, "%", ".*");
	
	var match, _ = regexp.MatchString(regExFormat, val)
	return match;
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_MatchRegex TTextSql = TTextSql{ 
	Postgres: "@x1@ ~ '@x2@'",
	//Postgres: "@x1@ ~* '@x2@'", insensitive
	Mysql:    "@x1@ REGEXP '@x2@'",
	Mssql:    "@x1@ LIKE '@x2@'",
}
// match `val` arg with a reg expression `regExFormat`.
// return true if is match
// Example:
//  var count = ctx.Table.Qry("tag1").Where( func(x *Table) bool { 
//  					return SQL_MatchRegex( x.Table.Field1, "*.acasa*.") }).GetCount()
//  })
// It will be translate in sql as:
//  Select COUNT(*) FROM Table WHERE Field1 ~ '*.acasa*.'
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_MatchRegex( val string, regExFormat string) bool {

	var match, _ = regexp.MatchString(regExFormat, val)
	return match;
}


//----------------------------------------------------------------------------------------------

// internal use
var  SQL_Int2Str TTextSql = TTextSql{ 
	Postgres: "CAST(@x1@ AS NUMERIC )",
	Mysql: "CONVERT(@x1@, INT )",
	Mssql: "CONVERT(@x1@, INT )",
}
// convert integer value to string
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_Int2Str( val int) string{

	return strconv.Itoa(val)
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_StrTrim TTextSql = TTextSql{ 
	Postgres: "TRIM(@x1@)",
	Mysql:    "TRIM(@x1@)",
	Mssql:   	"TRIM(@x1@)",
}
//return the string trimmed (left and right)
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_StrTrim( str string) string {

	return Str.Trim( str, " " );
}


//----------------------------------------------------------------------------------------------

// internal use
var  SQL_StrLen TTextSql = TTextSql{ 
	Postgres: 	"LEN(@x1@)",
	Mysql:    	"LEN(@x1@)",
	Mssql:   	"LEN(@x1@)",
}
//return the lenght of `str`
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_StrLen( str string) int {

	return len( str );
}

//----------------------------------------------------------------------------------------------

// internal use
var  SQL_Concat TTextSql = TTextSql{ 
	Postgres: 	"CONCAT( @xx1@ )",
	Mysql: 		"CONCAT( @xx1@ )",
	Mssql: 		"CONCAT( @xx1@ )",
}
//return concat values
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_Concat( arr ... string ) string{

	var min = "";
    for i :=0; i < len(arr); i++ {
        
    	min += arr[i]        
    }
	return min;
}

//----------------------------------------------------------------------------------------------
// internal use
var  SQL_ArrayContain TTextSql = TTextSql{ 
	Postgres: 	"@x2@ IN ( @x1@ )",
	Mysql: 		"@x2@ IN ( @x1@ )",
	Mssql: 		"@x2@ IN ( @x1@ )",
}
//return true: if `val` param is contained in array `arr` parameter
// 
// this method is translated in sql query, when it is using in Where() and Select() methods
func Sql_ArrayContain[T comparable]( arr []T, val T ) bool{

	return atomicsql.Arr_Contains( &arr, val );
}


