package gomigration_dialect

import (
	"fmt"
	"strings"
)

type TProcessData func(dialect GenericDialect, arg1 IGenericDialectArg)

type TGenericDialect struct {
	currentFile string
	dictError   map[string]string
	FnProcessData TProcessData
}

type GenericDialectArg struct {
	Connection_String string
}


type IGenericDialectArg interface{

    GetGenericDialectArg() *GenericDialectArg;
}




type GenericDialect interface{

GetGenericDialect() *TGenericDialect;
GetSql() string;

SqlSeparator() string;

DropColumn(table *DbTable, column *DbColumn) string;

AddColumn(table *DbTable, column *DbColumn) string;

UpdateColumn(table *DbTable, column, columnPrev *DbColumn) string;

DropTable(table *DbTable) string;

AddTable(table *DbTable) string;

UpdateTable(table, tablePrev *DbTable) string;

DropFKConstrictor(table *DbTable, column *DbColumn) string ;

AddFKConstrictor(table *DbTable, column *DbColumn) string;

GetSqlType(langType string, bIsNullable *bool, nameOfColumn string) string;

StartConnection(arg IGenericDialectArg) bool;

GetProperty(name, tableName, colName, colValue string) string;

InsertProperty(propName, value, tableName, colName, colValue string);

UpdateProperty(propName, value, tableName, colName, colValue string);

ExecScript(scriptTxt string);

ReadConstraintors(tables map[string]*DbTable) error;


ReadTables(SqlLang string) (map[string] *DbTable, error);

GetGoLangTypeIntFk(column *DbColumn, packageImports *map[string]string) string;

GetGoLangType(column *DbColumn, importPackage *map[string]string) string;

PrintError(err string) string ;
SetCurrentFile(file string);

CleanNameGoStruct(name string) string;
}

//-------------------------------------------------------------------------------------------------------------------------


//-------------------------------------------------------------------------------------------------------------------------

func (gd *TGenericDialect) SetCurrentFile(file string) {
	gd.currentFile = file
}

func (gd *TGenericDialect) PrintError(err string) string {
	if _, ok := gd.dictError[err]; !ok {
		gd.dictError[err] = err
		fmt.Println(err)

		if gd.currentFile != "" {
			fmt.Printf("Error In File: %s\n", gd.currentFile)
		}

		fmt.Println("")
	}
	return ""
}

func (gd *TGenericDialect) CleanNameGoStruct(name string) string {
	return TGenericDialect_CleanNameGoStruct(name)
}

func TGenericDialect_CleanNameGoStruct(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return name
}



func (gd *TGenericDialect) IsLongType1(sqlColumnName string) bool {
	if sqlColumnName != "" && strings.HasPrefix(sqlColumnName, "_") {
		return true
	}
	return false
}
func (gd *TGenericDialect) IsLongType(sqlColumnName string, type1 string, type2 string) string {
	var b = gd.IsLongType1(sqlColumnName)
    if( b ){
        return type1;
    }else{
        return type2;
    }	
}

