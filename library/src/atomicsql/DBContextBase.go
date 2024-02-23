package atomicsql

import (
	"database/sql"
	base64 "encoding/base64"
	"fmt"
	ioutil "io/ioutil"
	log "log"
	reflect "reflect"
	Str "strings"
	"time"
)

type VESqlDialect string
type TESqlDialect struct {
	Postgres VESqlDialect
	MySql     VESqlDialect
	MsSql     VESqlDialect
}

var ESqlDialect TESqlDialect = TESqlDialect{
	Postgres: "Postgres",
	MySql:     "MySql",
}

type TSqlColumnDef struct {
	LangName           string
	SqlName            string
	SqlType            string
	LangType           string
	Flags              string
	IsPrimary          bool
	IsNullable         bool
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
	END_COMMAND        string
	EMPTY_STRING       string
}

type TDefTable struct {
	SchemaTable           string
	SqlTableName          string
	PrimaryColumnLangName string
	PrimaryColumnSqlName  string
	Columns               []TSqlColumnDef
}

func (_this *TDefTable) getDictColumnByLangName() *map[string](TSqlColumnDef) {
	var dict = make(map[string](TSqlColumnDef))

	for _, col := range _this.Columns {

		dict[col.LangName] = col

	}
	return &dict
}

type TSchemaDef map[string]TDefTable

type TForeignKey struct {
	TgtTable_sqlName string
	TgtFldID_sqlName string

	RootTable_sqlName   string
	RootFldFk_sqlName   string
	RootFldFk_langName  string
	RootFldFk_lang2Name string
}

type TConnectionString struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SqlLang  VESqlDialect
}

type TExternVar struct {
	VarName string
	VarType string
}

// type TSubQuery = func(_ctx *DBContextBase, staticsVars *map[string]any, tagQuery string) (string,string)
type TSubQueryArg struct {
	Value       any    // the value for statics, for sql fields is : 34563456 or '34563456'
	Orginal_val string // 34563456 or '34563456'
	SqlCode     string // u1.UserRoleID
	ArgName     string // userRoleID
}

// type TSubQuery = func(_ctx *DBContextBase, argNames []any, tagQuery string) string
type TSubQuery struct {
	VariableName string //variable name: ids := ctx.table.QryS().Where()...
}

type TCompiledSqlQuery struct {
	CompiledQuery   string
	SelectSqlFields map[string]string
	//joins				[]string
	OrderedFields	[]string
	Fields    map[string]string
	ExternVar []TExternVar

	Tag        string
	File       string
	StartOff   int
	EndOff     int
	Hash       string // for checking the integrity
	IsQryS     bool
	SubQueries []TSubQuery
}

// this is the struct of ORM data context.
// after the developer execute 1.update_db.cmd, this will generate 2 files
// DbContext.gen.go, and DbContext_lambdaQueries.gen.go
// DbContext will be used in project containing all tables.
type DBContextBase struct {
	ConnectionString TConnectionString
	Db               *sql.DB

	SCHEMA_SQL           TSchemaDef
	SCHEMA_SQL_BySqlName TSchemaDef
	SCHEMA_SQL_Columns   map[string]map[string]string

	FOREIGN_KEYS map[string]TForeignKey
	LangDB       TLangDataBase
	Dialect      VESqlDialect //TESqlDialect

	CompiledSqlQueries       map[string]TCompiledSqlQuery
	DictTablesIncludeRelDefs map[string]*TDefIncludeRelation

	currOperationDTime        time.Time
	currOperationDTime2       time.Time
	accumulatorDTimeMicroSec  int64
	accumulatorDTimeMicroSec2 int64

	subTagCounter     int
	sqlItemDefCounter int

	AllTables map[string]*DBTable[IGeneric_MODEL]

	stackTransactions []*Transaction
	hasError          error

	GenericContext any /* *DBContext */
}
type IDBContext interface {
	GetContext() IDBContext
}

// execute sql query
// see sql.Exec()
func (_this *DBContextBase) Exec(sqlQuery string) (sql.Result, error) {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.Exec(sqlQuery)
	} else {
		return _this.Db.Exec(sqlQuery)
	}
}
func resultClose(dbResult sql.Result) {
	if dbResult != nil {

	}
}

// get a Row executing a sql query
// see sql.QueryRow()
func (_this *DBContextBase) QueryRow(sqlQuery string) *sql.Row {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.QueryRow(sqlQuery)
	} else {
		return _this.Db.QueryRow(sqlQuery)
	}
}

func queryRowClose(result *sql.Row) {

}

// get Rows executing a sql query
// see sql.Query()
func (_this *DBContextBase) Query(query string, args ...any) (*sql.Rows, error) {

	var cnt = len(_this.stackTransactions)
	if cnt > 0 {

		var tx = _this.stackTransactions[cnt-1]
		return tx.tx.Query(query)
	} else {
		return _this.Db.Query(query)
	}
}

func queryClose(result *sql.Rows) {

	if result != nil {
		result.Close()
	}
}

func (_this *DBContextBase) getSubTag() string {

	_this.subTagCounter++
	return fmt.Sprintf("%d", _this.subTagCounter)
}
func (_this *DBContextBase) resetSubTag() {

	_this.subTagCounter = 0
}

func (_this *DBContextBase) newSQL_ITEM(prefix string) string {

	_this.sqlItemDefCounter++
	var cnt = fmt.Sprintf("%d", _this.sqlItemDefCounter)

	return prefix + cnt
}

func (_this *DBContextBase) clearCachedSyntax() {

	_this.subTagCounter = 0
	_this.sqlItemDefCounter = 0
}

func (_this *DBContextBase) updateDeltaTime2() {

	var time1 = time.Now()
	var delta = time1.Sub(_this.currOperationDTime2)
	_this.accumulatorDTimeMicroSec2 += delta.Microseconds()
}

func (_this *DBContextBase) updateDeltaTime() {

	var time1 = time.Now()
	var delta = time1.Sub(_this.currOperationDTime)
	_this.accumulatorDTimeMicroSec += delta.Microseconds()
}

func (_this *DBContextBase) GetTotalDeltaTime() float64 {

	return float64(_this.accumulatorDTimeMicroSec) / 1000.0
	//return float64(_this.accumulatorDTimeMicroSec2) / 1000.0
}

// the constructor
func (_this *DBContextBase) Constr(dialect VESqlDialect, schemaSql TSchemaDef, ctxGeneric any) (*DBContextBase, error) {

	_this.SCHEMA_SQL = schemaSql
	_this.GenericContext = ctxGeneric

	var err error
	_this.SCHEMA_SQL_BySqlName, err = _this.convertSchema(schemaSql)
	if err != nil {
		return nil, err
	}

	_this.SCHEMA_SQL_Columns, err = _this.convertSchemaLangColumns(schemaSql)

	if dialect == ESqlDialect.Postgres {

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
			VALUE_NULL:         "null",
			END_COMMAND:        "GO",
		}
	} else if dialect == ESqlDialect.MySql {

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
			VALUE_NULL:         "null",
			END_COMMAND:        ";",
		}
	}

	return _this, nil
}

func (_this *DBContextBase) convertSchema(schemaSql TSchemaDef) (TSchemaDef, error) {

	var newSchema = make(TSchemaDef)
	for _, val := range schemaSql {

		_, has := newSchema[val.SqlTableName]
		if has {
			return nil, fmt.Errorf("duplicate name table in diferent schema: '%s'", val.SqlTableName)
		}
		newSchema[val.SqlTableName] = val
	}

	return newSchema, nil
}

func (_this *DBContextBase) convertSchemaLangColumns(schemaSql TSchemaDef) (map[string]map[string]string, error) {

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
func (_this *DBContextBase) Log_Fatal(formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Fatalln(formatstr)
}

// write in log
func (_this *DBContextBase) Log_Print(formatstr string) {

	//var msg = fmt.Sprintf(formatstr, args)
	log.Println(formatstr)
}

// It must be called in DBcontext_lambdaQueries.gen.go file. internal use
func (_this *DBContextBase) ProcessCompiledQuery(compiledSqlQueries *map[string]TCompiledSqlQuery, bDoAllChecks bool) {

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

func (_this *DBContextBase) isDialectSupportMultipleStatementsAtOnce() bool {

	if _this.Dialect == ESqlDialect.Postgres ||
		_this.Dialect == ESqlDialect.MsSql {
		return true
	}
	return false
}

func DBContext_cleanSaveFlags[T IGeneric_MODEL](model *T, _this *DBContextBase) {

	var reflVal = reflect.ValueOf(model).Elem()
	_DBContext_cleanSaveFlagsReflVal(reflVal, _this)
}

func _DBContext_cleanSaveFlagsReflVal(modelRefl reflect.Value, _this *DBContextBase) {

	var numCols = modelRefl.NumField()

	for i := 0; i < numCols; i++ {

		var field = modelRefl.Field(i)

		var nameTypeFld = field.Type().Name()
		if nameTypeFld == Generic_MODEL_Name {

			//var model1 = field.Interface().(Generic_MODEL)
			var model1 = field.Addr().Interface().(*Generic_MODEL)
			model1.flagIsSaved = false
			continue
		}
		if field.Type().Kind() == reflect.Pointer {

			var model1 = field.Elem()
			//var ptrVal = field.InterfaceData();
			var ptrVal = field.Pointer()
			if ptrVal != 0 && field.CanAddr() {
				_DBContext_cleanSaveFlagsReflVal(model1, _this)
			}
		}
	}
}
func DBContext_MarkSaved[T IGeneric_MODEL](model *T, _this *DBContextBase) bool {

	var reflVal = reflect.ValueOf(model).Elem()
	return DBContext_MarkSaveReflVal(reflVal, _this)
}

func DBContext_MarkSaveReflVal1[T IGeneric_MODEL](reflVal reflect.Value, _this *DBContextBase) bool {

	var ret = DBContext_MarkSaveReflVal(reflVal, _this)

	//check here the value of model
	var mm = reflVal.Interface().(T)
	if mm.GetID() != 0 {
	}
	return ret
}
func DBContext_MarkSaveReflVal(reflVal reflect.Value, _this *DBContextBase) bool {

	var numCols = reflVal.NumField()

	for i := 0; i < numCols; i++ {

		var field = reflVal.Field(i)
		var nameTypeFld = field.Type().Name()
		if nameTypeFld == Generic_MODEL_Name {

			//var model1 = field.Interface().(Generic_MODEL)
			var model1 = field.Addr().Interface().(*Generic_MODEL)
			model1.flagIsSaved = true
			return true
		}
	}
	return false
}

type TInterval struct {
	i0 int
	i1 int
}

func isInIntervals(intervals *[]TInterval, offset int) bool {

	for i := 0; i < len(*intervals); i++ {

		if (*intervals)[i].i0 <= offset &&
			(*intervals)[i].i1 >= offset {
			return true
		}
	}
	return false
}

func checkComments(txt1 string, idx int, length int) []TInterval {

	var intervals = []TInterval{}

	var txt = txt1[idx:]
	var idx2 = length

	for i := 0; i+1 < len(txt) && i+1 < idx2; i++ {

		var ch0 = txt[i]
		var ch1 = txt[i+1]
		var ttt = string([]byte{ch0, ch1})
		if ttt == "aa" {

		}

		if ch0 == '/' && ch1 == '/' {

			var iStart = i
			for ; i+1 < len(txt) && i+1 < idx2; i++ {

				var _ch0 = txt[i]
				if _ch0 == '\r' || _ch0 == '\n' {

					Arr_Append(&intervals, TInterval{i0: iStart + idx, i1: idx + i})
					break
				}
			}
		} else if ch0 == '/' && ch1 == '*' {

			var iStart = i
			for ; i+1 < len(txt) && i+1 < idx2; i++ {

				var _ch0 = txt[i]
				var _ch1 = txt[i+1]
				if _ch0 == '*' && _ch1 == '/' {

					Arr_Append(&intervals, TInterval{i0: iStart + idx, i1: idx + i})
					break
				}
			}
		}
	}
	return intervals
}

func (_this *DBContextBase) checkLambdaIntegrity(rootDir string, keyLamda string, lambda TCompiledSqlQuery) string {

	var filePath = rootDir + lambda.File
	var content, err = ioutil.ReadFile(filePath)
	if err != nil {
		//fmt.Printf("Eroare la citirea fișierului: %v\n", err)
		return "file-not-found"
	}

	var idx = 0
	var txt = string(content)
	if lambda.IsQryS {

		var parts = Str.Split(lambda.Tag, ".")
		var TagParent = parts[0]
		var TagQryS = parts[1]

		idx = Str_Index(txt, ".Qry(\""+TagParent+"\"", 0)
		if idx < 0 {
			idx = Str_Index(txt, ".Qry( \""+TagParent+"\"", 0)
		}
		if idx >= 0 {
			var intervals = checkComments(txt, idx, lambda.EndOff+200)

			var idxStart = idx
			for iLoop := 0; iLoop < 10; iLoop++ {
				var idx1 = Str_Index(txt, ".QryS(\""+TagQryS+"\"", idxStart)
				if idx1 < 0 {
					idx1 = Str_Index(txt, ".QryS( \""+TagQryS+"\"", idxStart)
				}
				if idx1 >= 0 {

					if !isInIntervals(&intervals, idx1) {

						idx = idx1
						break
					}
					idxStart = idx1 + 5
				}
			}
		}
	} else {
		idx = Str_Index(txt, ".Qry(\""+lambda.Tag+"\"", 0)
		if idx < 0 {
			idx = Str_Index(txt, ".Qry( \""+lambda.Tag+"\"", 0)
		}
	}

	if idx >= 0 {
		var contentLambdaFunc = Str_SubString(txt, idx+lambda.StartOff, lambda.EndOff-lambda.StartOff)

		var encodedString = base64.StdEncoding.EncodeToString([]byte(contentLambdaFunc))
		if encodedString != lambda.Hash {
			return keyLamda
		}
	}

	return ""
}

func (_this *DBContextBase) CheckIntegrity(rootDir string) string {

	for key, value := range _this.CompiledSqlQueries {

		//var = _this.CompiledSqlQueries[ itLambda ];
		var ret = _this.checkLambdaIntegrity(rootDir, key, value)
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
