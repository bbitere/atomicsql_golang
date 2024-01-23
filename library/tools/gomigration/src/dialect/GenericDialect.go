package gomigration_dialect

import (
	"fmt"
	"strings"
)

type TProcessData func(dialect *GenericDialect, arg1 *GenericDialectArg)

type TGenericDialect struct {
	currentFile string
	dictError   map[string]string
	fnProcessData TProcessData
}

type GenericDialect interface{

 getSql() string;

SqlSeparator();

dropColumn(table *DbTable, column *DbColumn) string;

addColumn(table *DbTable, column *DbColumn) string;

updateColumn(table *DbTable, column, columnPrev *DbColumn) ;

dropTable(table *DbTable) string;

addTable(table *DbTable) string;

updateTable(table, tablePrev *DbTable) string;

dropFKConstrictor(table *DbTable, column *DbColumn) string ;

addFKConstrictor(table *DbTable, column *DbColumn) string;

getSqlType(langType string, bIsNullable *bool, nameOfColumn string) string;

 startConnection(arg *GenericDialectArg) bool;

 getProperty(name, tableName, colName, colValue string) string;

insertProperty(propName, value, tableName, colName, colValue string);

 updateProperty(propName, value, tableName, colName, colValue string);

 execScript(scriptTxt string);

readConstraintors(tables map[string]*DbTable) bool;

readTables(SqlLang string) map[string]*DbTable;

getGoLangTypeIntFk(column *DbColumn, packageImports *map[string]string) string;

getGoLangType(column *DbColumn, importPackage *map[string]string) string;
}

func (gd TGenericDialect) setCurrentFile(file string) {
	gd.currentFile = file
}

func (gd *TGenericDialect) printError(err string) string {
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

func (gd *TGenericDialect) cleanNameGoStruct(name string) string {
	return TGenericDialect_CleanNameGoStruct(name)
}

func TGenericDialect_CleanNameGoStruct(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return name
}

func TGenericDialect_GetDialectByName(SqlLang string) *GenericDialect {
	if SqlLang == ELangSql.PostgresSql {
		return &PostgressDialect{}
	} else if SqlLang == ELangSql.MySql {
		return nil;//&MySqlDialect{}
	} else {
		fmt.Printf("Error: sql dialect '%s' not implemented\n", SqlLang)
		return nil
	}
}

func (gd *TGenericDialect) isLongType(sqlColumnName string) bool {
	if sqlColumnName != "" && strings.HasPrefix(sqlColumnName, "_") {
		return true
	}
	return false
}

