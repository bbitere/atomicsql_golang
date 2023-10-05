package atomicsql_func

import (
	//"reflect"
	//"unsafe"

	time "time"
)

// this struct should not be used in your code
//
// this is a mirror of time.Time. use time.Time
type TTime struct {
	time.Time
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_AddDate TTextSql = TTextSql{ 
	Postgres: 	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
	Mysql:    	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
	Mssql:   	"ADDDATE(@x1@, @x2@, @x3@, @x4@)",
}
//return a date = `time` + (year, month, day)
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_AddDate( time time.Time, years int, months int, days int)  time.Time{

	return time.AddDate( years, months, days);
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Now TTextSql = TTextSql{ 
	Postgres: 	"NOW()::Timestamp",
	Mysql:    	"NOW()",//UTC_TIMESTAMP()
	Mssql:   	"CURRENT_TIMESTAMP()",
}
//return a time with value of current time.
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.SQL_IFF]
func Sql_Now() time.Time{

	return time.Now();
}
//--------------------------------------------------------------------------------------------------

//internal use
var  SQL_Day TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Day FROM @THIS@)",
	Mysql:    	"DAY(@THIS@)",
	Mssql:   	"DAY(@THIS@)",
}
//return time.day. TTime is identic with time.Time. In your code use time.Time to use time type in Where()
// Example:
//  var count = ctx.Table.Qry("tag1").Where( func(x *Table) bool { 
//  					return SQL_IFF( x.Table.TimeAccess.Day() == 0, "Mon", "" ) == x.Table.DayStart.GetCount()
//  })
// It will be translate in sql as:
//  Select COUNT(*) FROM Table WHERE (CASE WHEN (Day(TimeAccess) = 0 ) THEN 'Mon' ELSE '' END) = DayStart
// 
func (_this *TTime) Day()  int{

	return _this.Time.Day();
}
//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Year TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Year FROM @THIS@)",
	Mysql:    	"YEAR(@THIS@)",
	Mssql:   	"YEAR(@THIS@)",
}
//return time.year
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Year()  int{

	return _this.Time.Year();
}


//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Month TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Month FROM @THIS@)",
	Mysql:    	"MONTH(@THIS@)",
	Mssql:   	"MONTH(@THIS@)",
}
//return time.month
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Month()  int{

	return int(_this.Time.Month());
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Hour TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Hour FROM @THIS@)",
	Mysql:    	"HOUR(@THIS@)",
	Mssql:   	"HOUR(@THIS@)",
}
//return time.hour
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Hour()  int{

	return (_this.Time.Hour());
}
//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Minute TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Minute FROM @THIS@)",
	Mysql:    	"Minute(@THIS@)",
	Mssql:   	"Minute(@THIS@)",
}
//return time.minute
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Minute()  int{

	return (_this.Time.Minute());
}
//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Second TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(Second FROM @THIS@)",
	Mysql:    	"Second(@THIS@)",
	Mssql:   	"Second(@THIS@)",
}
//return time.second
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Second()  int{

	return (_this.Time.Second());
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_UnixMilli TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(MILLIS FROM @THIS@)",
	Mysql:    	"MILLIS(@THIS@)",
	Mssql:   	"MILLIS(@THIS@)",
}
//return time.millisecond
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) UnixMilli() int64{

	return (_this.Time.UnixMilli());
}


//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_YearDay TTextSql = TTextSql{ 
	Postgres: 	"EXTRACT(YearDay FROM @THIS@)",
	Mysql:    	"YearDay(@THIS@)",
	Mssql:   	"YearDay(@THIS@)",
}
//return the day of year
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) YearDay()  int{

	return (_this.Time.YearDay());
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Compare TTextSql = TTextSql{ 
	Postgres: 	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
	Mysql:    	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
	Mssql:   	"CASE WHEN @THIS@ = @x1@ THEN 0 WHEN @THIS@ < @x1@ THEN -1 ELSE 1 END",
}
//return compare this with `t` struct of time.
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Compare(t time.Time)  int{

	return int(_this.Time.Compare(t));
}

//--------------------------------------------------------------------------------------------------

// internal use
var  SQL_Format TTextSql = TTextSql{ 
	Postgres: 	"to_char(@THIS@, @x1@)",
	Mysql:    	"DATE_FORMAT(@THIS@, @x1@)",
	Mssql:   	"FORMAT(@THIS@, @x1@)",
}
//return a string formated like Time.Format
// 
// this method is translated in sql query, when it is using in Where() and Select() methods. For more info see usage of [atomicsql_func.TTime.Day]
func (_this *TTime) Format(formatStr string)  string{

	return _this.Time.Format(formatStr);
}


