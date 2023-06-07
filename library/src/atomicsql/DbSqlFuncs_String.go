package atomicsql

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"

	"strconv"
	"strings"

	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
	"regexp"
	Str "strings"
)

//----------------------------------------------------------------------------------------------
var  SQL_Like TTextSql = TTextSql{ 
	Postgres: "@x1@ like '@x2@'",
	Mysql:    "@x1@ like '@x2@'",
	Mssql:    "@x1@ like '@x2@'",
}
func Sql_Like( val string, likeFormat string) bool {

	var regExFormat = strings.ReplaceAll( likeFormat, "%", ".*");
	
	var match, _ = regexp.MatchString(regExFormat, val)
	return match;
}

//----------------------------------------------------------------------------------------------
var  SQL_MatchRegex TTextSql = TTextSql{ 
	Postgres: "@x1@ ~ '@x2@'",
	//Postgres: "@x1@ ~* '@x2@'", insensitive
	Mysql:    "@x1@ REGEXP '@x2@'",
	Mssql:    "@x1@ LIKE '@x2@'",
}
func Sql_MatchRegex( val string, regExFormat string) bool {

	var match, _ = regexp.MatchString(regExFormat, val)
	return match;
}


//----------------------------------------------------------------------------------------------
var  SQL_Int2Str TTextSql = TTextSql{ 
	Postgres: "CAST(@x1@ AS NUMERIC )",
	Mysql: "CONVERT(@x1@, INT )",
	Mssql: "CONVERT(@x1@, INT )",
}
func Sql_Int2Str( val int) string{

	return strconv.Itoa(val)
}

//----------------------------------------------------------------------------------------------
var  SQL_StrTrim TTextSql = TTextSql{ 
	Postgres: "TRIM(@x1@)",
	Mysql:    "TRIM(@x1@)",
	Mssql:   	"TRIM(@x1@)",
}

func Sql_StrTrim( str string) string {

	return Str.Trim( str, " " );
}


//----------------------------------------------------------------------------------------------
var  SQL_StrLen TTextSql = TTextSql{ 
	Postgres: 	"LEN(@x1@)",
	Mysql:    	"LEN(@x1@)",
	Mssql:   	"LEN(@x1@)",
}

func Sql_StrLen( str string) int {

	return len( str );
}

//----------------------------------------------------------------------------------------------
var  SQL_Concat TTextSql = TTextSql{ 
	Postgres: 	"CONCAT( @xx1@ )",
	Mysql: 		"CONCAT( @xx1@ )",
	Mssql: 		"CONCAT( @xx1@ )",
}
func Sql_Concat( arr ... string ) string{

	var min = "";
    for i :=0; i < len(arr); i++ {
        
    	min += arr[i]        
    }
	return min;
}