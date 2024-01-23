package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	//"sort"
	"strings"
	//"time"
	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	dialect_general "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect/general"
	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"
)

const (
	RECOGNIZE_FILE_PATTERN = `^\d{4}\.\d{2}\.\d{2}\.(\d{2}|\d{3}|\d{4}|\d{6})\.(.+)$`
	TABLE_MIGRATION        = "DBParams"
	ParamName              = "ParamName"
	ParamValue             = "ParamValue"
	SQL_FILE_APPLIED       = "SQL_APPL"
	PROP_HISTORY           = "version"
)

type DialectArg struct {
	Base dialect.GenericDialectArg
	inputDir     string
	outDir       string
	typeOutFile  string	
}

func (m *DialectArg) GetGenericDialectArg() *dialect.GenericDialectArg{

	return &m.Base;
}

type MigrationDB struct
{

}

func (m *MigrationDB) DoMigration(sqlLang, connectionString, inputDir, outDir, typeOutFile string) bool {

	arg := &DialectArg{
		inputDir:    inputDir,
		outDir:      outDir,
		typeOutFile: typeOutFile,
		Base: dialect.GenericDialectArg{
			Connection_String: connectionString,
		},
	}

	var dialectInst = dialect_general.GetDialectByName(sqlLang)
	if dialectInst == nil {
		return false
	}
	var genericDialect = dialectInst.GetGenericDialect();

	genericDialect.FnProcessData = m.ProcessData
	return dialectInst.StartConnection( arg );
}

func (m *MigrationDB) ProcessData(dialect dialect.GenericDialect, arg1 dialect.IGenericDialectArg) {
	arg := arg1.(*DialectArg)
	history := dialect.GetProperty(PROP_HISTORY, TABLE_MIGRATION, ParamName, ParamValue)
	if history == "" {
		m.setupMigration(dialect)
	}

	files := m.scanFiles(arg.inputDir)
	//sortedFiles := m.sortedFilesByName( &files );
	var sortedFiles = sort.StringSlice(files);
	
	m.applyImports(dialect, sortedFiles)

	fmt.Println("End Transaction")
}

func (m *MigrationDB) sortedFilesByName(files *[]string){

}

func (m *MigrationDB) scanFiles(dir string) []string {
	var listFiles []string
	regex := regexp.MustCompile(RECOGNIZE_FILE_PATTERN)

	allFiles, _ := filepath.Glob(filepath.Join(dir, "*"))
	for _, file := range allFiles {
		fileInfo, _ := os.Stat(file)
		if fileInfo.IsDir() {
			subdirFiles := m.scanFiles(file)
			listFiles = append(listFiles, subdirFiles...)
		} else if utils.Utils_GetFileInfoExt(file) == ".sql" && regex.MatchString(fileInfo.Name()) {
			listFiles = append(listFiles, file)
		}
	}

	return listFiles
}

func (m *MigrationDB) applyImports(dialect dialect.GenericDialect, files []string) {
	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		m.ApplyImport(dialect, file, string(content))
	}
}

func (m *MigrationDB) ApplyImport(dialect dialect.GenericDialect, file string, content string) {
	//fileInfo, _ := os.Stat(file)
	property := strings.TrimSuffix( utils.Utils_GetFileInfoName(file), utils.Utils_GetFileInfoExt(file))

	fileApplied := dialect.GetProperty(property, TABLE_MIGRATION, ParamName, ParamValue)
	if fileApplied == "" {
		fmt.Printf("Start script: %s\n", property)
		dialect.ExecScript(content)
		dialect.InsertProperty(property, SQL_FILE_APPLIED, TABLE_MIGRATION, ParamName, ParamValue)

		historyPropValue := m.getHistoryName(property)
		dialect.UpdateProperty(PROP_HISTORY, historyPropValue, TABLE_MIGRATION, ParamName, ParamValue)

		fmt.Printf("End   script: %s\n", property)
	}
}

func (m *MigrationDB) getHistoryName(property string) string {
	parts := strings.Split(property, ".")
	parts = parts[:len(parts)-1]
	return strings.Join(parts, ".")
}

func (m *MigrationDB) setupMigration(dialectInst dialect.GenericDialect) {
	tableMigration := createMigrationTable(dialectInst)
	scriptTable := dialectInst.AddTable(tableMigration)
	dialectInst.ExecScript(scriptTable)

	dialectInst.InsertProperty(PROP_HISTORY, "0000.00.00", TABLE_MIGRATION, ParamName, ParamValue)
}

func newDbColumn() *dialect.DbColumn{
	return new (dialect.DbColumn);
}

func createMigrationTable(dialectInst dialect.GenericDialect) *dialect.DbTable {
	isNullable := false
	tableMigration := &dialect.DbTable{}
	tableMigration.Columns = append(tableMigration.Columns, newDbColumn().InitSqlPrimary("ID"))
	tableMigration.Columns = append(tableMigration.Columns, newDbColumn().InitSql(ParamName, dialectInst.GetSqlType("string", &isNullable, ""), isNullable))
	tableMigration.Columns = append(tableMigration.Columns, newDbColumn().InitSql(ParamValue, dialectInst.GetSqlType("string", &isNullable, ""), isNullable))

	tableMigration.InitSql(TABLE_MIGRATION, tableMigration.Columns[0])
	return tableMigration
}

/*
// Helper function to sort files
func (files []string) []string {
	sort.Strings(files)
	return files
}
*/