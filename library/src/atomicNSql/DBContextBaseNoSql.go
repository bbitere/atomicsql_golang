package atomicNSql

import (
	//"database/sql"
	base64 "encoding/base64"
	"fmt"

	//ioutil "io/ioutil"
	log "log"
	//reflect "reflect"
	//Str "strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

type TDefTable = atomicsql.TDefTable
type TSqlColumnDef = atomicsql.TSqlColumnDef
type TDefIncludeRelation = atomicsql.TDefIncludeRelation


type DBContextBaseNoSql struct {
	ConnectionString atomicsql.TConnectionString
	//Db               *sql.DB

	SCHEMA_SQL           atomicsql.TSchemaDef
	SCHEMA_SQL_BySqlName atomicsql.TSchemaDef
	SCHEMA_SQL_Columns   map[string]map[string]string

	//FOREIGN_KEYS map[string]TForeignKey
	LangDB       atomicsql.TLangDataBase
	Dialect      atomicsql.VESqlDialect //TESqlDialect

	CompiledSqlQueries       map[string]atomicsql.TCompiledSqlQuery
	DictTablesIncludeRelDefs map[string]*atomicsql.TDefIncludeRelation

	currOperationDTime        time.Time
	currOperationDTime2       time.Time
	accumulatorDTimeMicroSec  int64
	accumulatorDTimeMicroSec2 int64

	subTagCounter     int
	sqlItemDefCounter int

	AllTables map[string]*DBTableNoSql[atomicsql.IGeneric_MODEL]

	stackTransactions []*atomicsql.Transaction
	hasError          error

	GenericContext any /* *DBContext */

	//ConnString TConnectionString
	ClientDB *mongo.Client;
	Database *mongo.Database;
}

// the constructor
func (_this *DBContextBaseNoSql) Constr(dialect atomicsql.VESqlDialect, schemaSql atomicsql.TSchemaDef, ctxGeneric any) (*DBContextBaseNoSql, error) {

	_this.SCHEMA_SQL = schemaSql
	_this.GenericContext = ctxGeneric

	var err error
	_this.SCHEMA_SQL_BySqlName, err = _this.convertSchema(schemaSql)
	if err != nil {
		return nil, err
	}

	_this.SCHEMA_SQL_Columns, err = _this.convertSchemaLangColumns(schemaSql)

	if dialect == atomicsql.ESqlDialect.MongoDB {

		/*
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
			VALUE_NULL:         "null",
			END_COMMAND:        ";",
		}*/
	} 

	return _this, nil
}

/*
// execute sql query
// see sql.Exec()
func (_this *DBContextBaseNoSql) Exec(sqlQuery string) (sql.Result, error) {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.Exec(sqlQuery)
	} else {
		return _this.Db.Exec(sqlQuery)
	}
}


// get a Row executing a sql query
// see sql.QueryRow()
func (_this *DBContextBaseNoSql) QueryRow(sqlQuery string) *sql.Row {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.QueryRow(sqlQuery)
	} else {
		return _this.Db.QueryRow(sqlQuery)
	}
}



// get Rows executing a sql query
// see sql.Query()
func (_this *DBContextBaseNoSql) Query(query string, args ...any) (*sql.Rows, error) {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.Query(query)
	} else {
		return _this.Db.Query(query)
	}
}
*/


func (_this *DBContextBaseNoSql) getSubTag() string {

	_this.subTagCounter++
	return fmt.Sprintf("%d", _this.subTagCounter)
}
func (_this *DBContextBaseNoSql) resetSubTag() {

	_this.subTagCounter = 0
}

func (_this *DBContextBaseNoSql) newSQL_ITEM(prefix string) string {

	_this.sqlItemDefCounter++
	var cnt = fmt.Sprintf("%d", _this.sqlItemDefCounter)

	return prefix + cnt
}

func (_this *DBContextBaseNoSql) clearCachedSyntax() {

	_this.subTagCounter = 0
	_this.sqlItemDefCounter = 0
}

func (_this *DBContextBaseNoSql) updateDeltaTime2() {

	var time1 = time.Now()
	var delta = time1.Sub(_this.currOperationDTime2)
	_this.accumulatorDTimeMicroSec2 += delta.Microseconds()
}

func (_this *DBContextBaseNoSql) updateDeltaTime() {

	var time1 = time.Now()
	var delta = time1.Sub(_this.currOperationDTime)
	_this.accumulatorDTimeMicroSec += delta.Microseconds()
}

func (_this *DBContextBaseNoSql) GetTotalDeltaTime() float64 {

	return float64(_this.accumulatorDTimeMicroSec) / 1000.0
	//return float64(_this.accumulatorDTimeMicroSec2) / 1000.0
}


func (_this *DBContextBaseNoSql) convertSchema(schemaSql atomicsql.TSchemaDef) (atomicsql.TSchemaDef, error) {

	var newSchema = make(atomicsql.TSchemaDef)
	for _, val := range schemaSql {

		_, has := newSchema[val.SqlTableName]
		if has {
			return nil, fmt.Errorf("duplicate name table in diferent schema: '%s'", val.SqlTableName)
		}
		newSchema[val.SqlTableName] = val
	}

	return newSchema, nil
}

func (_this *DBContextBaseNoSql) convertSchemaLangColumns(schemaSql atomicsql.TSchemaDef) (map[string]map[string]string, error) {

	var newSchema = make(map[string]map[string]string)

	for tableName, val := range schemaSql {

		var newTable = make(map[string]string)
		newSchema[tableName] = newTable

		for _, val1 := range val.Columns {
			newTable[val1.LangName] = val1.SqlName
		}
	}

	return newSchema, nil
}

// write in log
func (_this *DBContextBaseNoSql) Log_Fatal(formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Fatalln(formatstr)
}

// write in log
func (_this *DBContextBaseNoSql) Log_Print(formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Println(formatstr)
}

// It must be called in DBcontext_lambdaQueries.gen.go file. internal use
func (_this *DBContextBaseNoSql) ProcessCompiledQuery(compiledSqlQueries *map[string]atomicsql.TCompiledSqlQuery, bDoAllChecks bool) {

	for key, val := range *compiledSqlQueries {

		var bytes, err = base64.StdEncoding.DecodeString(val.CompiledQuery)
		if err != nil {
			log.Printf("loading compiled query error: %s", err.Error())
			return
		}
		val.CompiledQuery = string(bytes)
		(*compiledSqlQueries)[key] = val
	}

	if bDoAllChecks {
	}
}

func (_this *DBContextBaseNoSql) isDialectSupportMultipleStatementsAtOnce() bool {
	
	return false
}


func (_this *DBContextBaseNoSql) CheckIntegrity(rootDir string) string {

	for key, value := range _this.CompiledSqlQueries {

		//var = _this.CompiledSqlQueries[ itLambda ];
		var ret = atomicsql.CheckLambdaIntegrity(rootDir, key, value)
		if ret == "file-not-found" {
			continue
		}
		if ret != "" {
			fmt.Printf("A lambda expression (%s) is modified. Recompile the project!", key)
			fmt.Println("")
			return ret
		}
	}
	return ""
}




