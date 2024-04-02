package atomicsql

import (
	//"database/sql"
	base64 "encoding/base64"
	"fmt"
	ioutil "io/ioutil"

	//log "log"
	reflect "reflect"
	Str "strings"
	//"time"
)
	

type VESqlDialect string
type TESqlDialect struct {
	Postgres  VESqlDialect
	MySql     VESqlDialect
	MsSql     VESqlDialect
	MongoDB   VESqlDialect
}

var ESqlDialect TESqlDialect = TESqlDialect{
	Postgres: 	"Postgres",
	MySql:     	"MySql",
	MongoDB:	"mongo",
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
	NosqlQuery [] any
}


type IDBContext interface {
	GetContext() IDBContext
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

	return CheckLambdaIntegrity(rootDir, keyLamda, lambda) 
}

func CheckLambdaIntegrity(rootDir string, keyLamda string, lambda TCompiledSqlQuery) string {

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
