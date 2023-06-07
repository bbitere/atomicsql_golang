package atomicsql

import (
	//"reflect"
	//"unsafe"
	//sql "database/sql"
	//fmt "fmt"
	//reflect "reflect"

	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"

	time "time"
)
type TTime struct {
	time.Time
}

//--------------------------------------------------------------------------------------------------
var  SQL_AddDate TTextSql = TTextSql{ 
	Postgres: 	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
	Mysql:    	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
	Mssql:   	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
}
func Sql_AddDate( time time.Time, years int, months int, days int)  time.Time{

	return time.AddDate( years, months, days);
}

//--------------------------------------------------------------------------------------------------
var  SQL_Now TTextSql = TTextSql{ 
	Postgres: 	"NOW()::Timestamp",
	Mysql:    	"NOW()",//UTC_TIMESTAMP()
	Mssql:   	"CURRENT_TIMESTAMP()",
}
func Sql_Now() time.Time{

	return time.Now();
}
//--------------------------------------------------------------------------------------------------


var  SQL_Day TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Day FROM @THIS@)",
	Mysql:    	"DAY(@THIS@)",
	Mssql:   	"DAY(@THIS@)",
}
func (_this *TTime) Day()  int{

	return _this.Time.Day();
}
//--------------------------------------------------------------------------------------------------
var  SQL_Year TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Year FROM @THIS@)",
	Mysql:    	"YEAR(@THIS@)",
	Mssql:   	"YEAR(@THIS@)",
}
func (_this *TTime) Year()  int{

	return _this.Time.Year();
}


//--------------------------------------------------------------------------------------------------
var  SQL_Month TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Month FROM @THIS@)",
	Mysql:    	"MONTH(@THIS@)",
	Mssql:   	"MONTH(@THIS@)",
}
func (_this *TTime) Month()  int{

	return int(_this.Time.Month());
}

//--------------------------------------------------------------------------------------------------
var  SQL_Hour TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Hour FROM @THIS@)",
	Mysql:    	"HOUR(@THIS@)",
	Mssql:   	"HOUR(@THIS@)",
}
func (_this *TTime) Hour()  int{

	return (_this.Time.Hour());
}
//--------------------------------------------------------------------------------------------------
var  SQL_Minute TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Minute FROM @THIS@)",
	Mysql:    	"Minute(@THIS@)",
	Mssql:   	"Minute(@THIS@)",
}
func (_this *TTime) Minute()  int{

	return (_this.Time.Minute());
}
//--------------------------------------------------------------------------------------------------
var  SQL_Second TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Second FROM @THIS@)",
	Mysql:    	"Second(@THIS@)",
	Mssql:   	"Second(@THIS@)",
}
func (_this *TTime) Second()  int{

	return (_this.Time.Second());
}

//--------------------------------------------------------------------------------------------------
var  SQL_UnixMilli TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(MILLIS FROM @THIS@)",
	Mysql:    	"MILLIS(@THIS@)",
	Mssql:   	"MILLIS(@THIS@)",
}
func (_this *TTime) UnixMilli() int64{

	return (_this.Time.UnixMilli());
}


//--------------------------------------------------------------------------------------------------
var  SQL_YearDay TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(YearDay FROM @THIS@)",
	Mysql:    	"YearDay(@THIS@)",
	Mssql:   	"YearDay(@THIS@)",
}
func (_this *TTime) YearDay()  int{

	return (_this.Time.YearDay());
}

//--------------------------------------------------------------------------------------------------
var  SQL_Compare TTextSql = TTextSql{ 
	Postgres: 	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
	Mysql:    	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
	Mssql:   	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
}
func (_this *TTime) Compare(t time.Time)  int{

	return int(_this.Time.Compare(t));
}

//--------------------------------------------------------------------------------------------------
var  SQL_Format TTextSql = TTextSql{ 
	Postgres: 	"to_char(@THIS@, @x1@)",
	Mysql:    	"DATE_FORMAT(@THIS@, @x1@)",
	Mssql:   	"FORMAT(@THIS@, @x1@)",
}
func (_this *TTime) Format(formatStr string)  string{

	return _this.Time.Format(formatStr);
}


