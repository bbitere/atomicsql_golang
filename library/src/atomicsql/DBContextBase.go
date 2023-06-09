package atomicsql

import (
	"database/sql"
	base64 "encoding/base64"
	"fmt"
	log "log"
	"time"
)

type TESqlDialect struct {
	Postgress string
	MySql     string
	MsSql     string
}

var ESqlDialect TESqlDialect = TESqlDialect{
	Postgress: "Postgress",
	MySql:     "MySql",
}

type TSqlColumnDef struct {
	LangName   string
	SqlName    string
	SqlType    string
	LangType   string
	Flags      string
	IsPrimary  bool
	IsNullable bool
	ForeignKeyLangName []string
}

// language of database: mysql and postgressql
type TLangDataBase struct {
	Type_BOOL          string
	Type_VARCHAR       string
	Type_CHAR          string
	Type_SMALLINT      string
	Type_INTEGER       string
	Type_SERIAL        string
	Type_FLOAT         string
	Type_DOUBLE        string
	Type_DATATIME      string
	Type_DATATIME_NULL string
	VALUE_TRUE         string
	VALUE_FALSE        string
	VALUE_NULL         string
	END_COMMAND			string
	EMPTY_STRING       string
}

type TDefTable struct {
	SchemaTable           string
	SqlTableName          string
	PrimaryColumnLangName string
	PrimaryColumnSqlName  string
	Columns               []TSqlColumnDef
}



func (_this *TDefTable)getDictColumnByLangName() *map[string](TSqlColumnDef){
	var dict = make(map[string](TSqlColumnDef))

	for _, col := range( _this.Columns){

		dict[ col.LangName ] = col;
		
	}
	return &dict
}

type TSchemaDef map[string]TDefTable

type TForeignKey struct {
	TgtTable_sqlName string
	TgtFldID_sqlName string

	RootTable_sqlName string
	RootFldFk_sqlName string
	RootFldFk_langName string 
	RootFldFk_lang2Name string
}

type TConnectionString struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}


type TExternVar struct {
	VarName string
	VarType string
}
type TCompiledSqlQuery struct {

	CompiledQuery 		string
	SelectSqlFields		map[string]string
	//joins				[]string
	Fields				map[string]string
	ExternVar			[]TExternVar

	Tag					string
	File 				string
	StartOff			int
	EndOff				int
	Hash				string // for checking the integrity

}

type DBContextBase struct {
	ConnectionString TConnectionString
	Db               *sql.DB

	SCHEMA_SQL 				TSchemaDef
	SCHEMA_SQL_BySqlName 	TSchemaDef
	SCHEMA_SQL_Columns		map[string] map[string]string

	FOREIGN_KEYS 			map[string]TForeignKey
	LangDB     				TLangDataBase
	Dialect    				string //TESqlDialect

	CompiledSqlQueries		map[ string ]TCompiledSqlQuery

	currOperationDTime 		time.Time
	currOperationDTime2 	time.Time
	accumulatorDTimeMicroSec int64
	accumulatorDTimeMicroSec2 int64

	subTagCounter			int
	sqlItemDefCounter		int 

	AllTables				map[ string ]*DBTable[IGeneric_MODEL]

}
type IDBContext interface {
	GetContext()IDBContext
}


func (_this *DBContextBase) getSubTag() string{
	
	_this.subTagCounter++;
	return fmt.Sprintf("%d",_this.subTagCounter);
}
func (_this *DBContextBase) resetSubTag() {
	
	_this.subTagCounter = 0;
}


func (_this *DBContextBase) newSQL_ITEM(prefix string) string{

	_this.sqlItemDefCounter++;
	var cnt = fmt.Sprintf("%d", _this.sqlItemDefCounter );

	return prefix+cnt;
}

func (_this *DBContextBase) clearCachedSyntax(){

	_this.subTagCounter =0;
	_this.sqlItemDefCounter = 0;
}


func (_this *DBContextBase) updateDeltaTime2(){

	var time1 = time.Now()
	var delta = time1.Sub( _this.currOperationDTime2)
	_this.accumulatorDTimeMicroSec2 += delta.Microseconds()
}

func (_this *DBContextBase) updateDeltaTime(){

	var time1 = time.Now()
	var delta = time1.Sub( _this.currOperationDTime)
	_this.accumulatorDTimeMicroSec += delta.Microseconds()
}



func (_this *DBContextBase) GetTotalDeltaTime() float64{

	return float64(_this.accumulatorDTimeMicroSec) / 1000.0
	//return float64(_this.accumulatorDTimeMicroSec2) / 1000.0
}


func (_this *DBContextBase) Constr(dialect string, schemaSql TSchemaDef) (*DBContextBase, error){

	_this.SCHEMA_SQL = schemaSql

	var err error
	_this.SCHEMA_SQL_BySqlName, err = _this.convertSchema(schemaSql)
	if( err != nil){
		return nil, err
	}

	_this.SCHEMA_SQL_Columns, err = _this.convertSchemaLangColumns(schemaSql)

	if dialect == ESqlDialect.Postgress {

		_this.LangDB = TLangDataBase{
			Type_BOOL:          "boolean",
			Type_VARCHAR:       "VARCHAR",
			Type_CHAR:          "CHAR",
			Type_SMALLINT:      "SMALLINT",
			Type_INTEGER:       "INTEGER",
			Type_SERIAL:        "SERIAL",
			Type_FLOAT:         "FLOAT",
			Type_DOUBLE:        "REAL",
			Type_DATATIME:      "DATETIME",
			Type_DATATIME_NULL: "DATETIME NULL",
			VALUE_TRUE:         "true",
			VALUE_FALSE:        "false",
			EMPTY_STRING:       "''",
			VALUE_NULL:       	"null",
			END_COMMAND:		";",
		}
	} else if dialect == ESqlDialect.MsSql {

		_this.LangDB = TLangDataBase{
			Type_BOOL:          "BIT",
			Type_VARCHAR:       "VARCHAR",
			Type_CHAR:          "CHAR",
			Type_SMALLINT:      "SMALLINT",
			Type_INTEGER:       "INTEGER",
			Type_SERIAL:        "INTEGER",
			Type_FLOAT:         "FLOAT",
			Type_DOUBLE:        "REAL",
			Type_DATATIME:      "DATETIME",
			Type_DATATIME_NULL: "DATETIME NULL",
			VALUE_TRUE:         "1",
			VALUE_FALSE:        "0",
			EMPTY_STRING:       "''",
			VALUE_NULL:       	"null",
			END_COMMAND:		"GO",
		}
	} else if dialect == ESqlDialect.MySql  {

		_this.LangDB = TLangDataBase{
			Type_BOOL:          "BIT",
			Type_VARCHAR:       "VARCHAR",
			Type_CHAR:          "CHAR",
			Type_SMALLINT:      "SMALLINT",
			Type_INTEGER:       "INTEGER",
			Type_SERIAL:        "INTEGER",
			Type_FLOAT:         "FLOAT",
			Type_DOUBLE:        "REAL",
			Type_DATATIME:      "DATETIME",
			Type_DATATIME_NULL: "DATETIME NULL",
			VALUE_TRUE:         "1",
			VALUE_FALSE:        "0",
			EMPTY_STRING:       "''",
			VALUE_NULL:       	"null",
			END_COMMAND:		";",
		}
	}

	
	return _this, nil
}


func (_this *DBContextBase) convertSchema( schemaSql TSchemaDef) (TSchemaDef, error) {

	var newSchema = make( TSchemaDef )
	for _, val := range(schemaSql) {

		_, has := newSchema[ val.SqlTableName ];
		if( has ){
			return nil, fmt.Errorf("duplicate name table in diferent schema: '%s'", val.SqlTableName );
		}
		newSchema[ val.SqlTableName ] = val;
	}	
		
	return newSchema, nil
}



func (_this *DBContextBase) convertSchemaLangColumns( schemaSql TSchemaDef ) (map[string] map[string]string, error) {

	var newSchema = make( map[string] map[string]string )

	for tableName, val := range(schemaSql) {

		var newTable = make( map[string]string ) 
		newSchema[ tableName ] = newTable;
		
		for _, val1 := range(val.Columns) {
			newTable[val1.LangName] = val1.SqlName
		}
	}	
		
	return newSchema, nil
}


func (_this *DBContextBase) Log_Fatal( formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Fatalln(formatstr);
}

func (_this *DBContextBase) Log_Print( formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Println(formatstr);
}

func (_this *DBContextBase) ProcessCompiledQuery( compiledSqlQueries* map[string]TCompiledSqlQuery, bDoAllChecks bool ){

	for key, val:= range(*compiledSqlQueries) {

		var bytes, err = base64.StdEncoding.DecodeString( val.CompiledQuery );
    	if err != nil {
			log.Printf("loading compiled query error: %s", err.Error() );
			return;
		}
		val.CompiledQuery = string( bytes );
		(*compiledSqlQueries)[key] = val;
	}

	if( bDoAllChecks){

	}

}

