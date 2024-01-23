package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	//"sort"
	"strings"
	//"time"
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
	inputDir     string
	outDir       string
	typeOutFile  string
	GenericDialectArg
}

type MigrationDB struct
{

}

func (m *MigrationDB) DoMigration(sqlLang, connectionString, inputDir, outDir, typeOutFile string) bool {
	arg := &DialectArg{
		inputDir:    inputDir,
		outDir:      outDir,
		typeOutFile: typeOutFile,
		GenericDialectArg: GenericDialectArg{
			strConnection: connectionString,
		},
	}

	dialect := GetDialectByName(sqlLang)
	if dialect == nil {
		return false
	}

	dialect.fnProcessData = m.ProcessData
	return dialect.startConnection(arg)
}

func (m *MigrationDB) ProcessData(dialect GenericDialect, arg1 GenericDialectArg) {
	arg := arg1.(*DialectArg)
	history := dialect.getProperty(PROP_HISTORY, TABLE_MIGRATION, ParamName, ParamValue)
	if history == nil {
		m.setupMigration(dialect)
	}

	files := m.scanFiles(arg.inputDir)
	sortedFiles := files.sortedFiles()
	m.applyImports(dialect, sortedFiles)

	fmt.Println("End Transaction")
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
		} else if Utils_GetFileInfoExt(file) == ".sql" && regex.MatchString(fileInfo.Name()) {
			listFiles = append(listFiles, file)
		}
	}

	return listFiles
}

func (m *MigrationDB) applyImports(dialect GenericDialect, files []string) {
	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		m.ApplyImport(dialect, file, string(content))
	}
}

func (m *MigrationDB) ApplyImport(dialect GenericDialect, file string, content string) {
	//fileInfo, _ := os.Stat(file)
	property := strings.TrimSuffix( Utils_GetFileInfoName(file), Utils_GetFileInfoExt(file))

	fileApplied := dialect.getProperty(property, TABLE_MIGRATION, ParamName, ParamValue)
	if fileApplied == nil {
		fmt.Printf("Start script: %s\n", property)
		dialect.execScript(content)
		dialect.insertProperty(property, SQL_FILE_APPLIED, TABLE_MIGRATION, ParamName, ParamValue)

		historyPropValue := m.getHistoryName(property)
		dialect.updateProperty(PROP_HISTORY, historyPropValue, TABLE_MIGRATION, ParamName, ParamValue)

		fmt.Printf("End   script: %s\n", property)
	}
}

func (m *MigrationDB) getHistoryName(property string) string {
	parts := strings.Split(property, ".")
	parts = parts[:len(parts)-1]
	return strings.Join(parts, ".")
}

func (m *MigrationDB) setupMigration(dialect GenericDialect) {
	tableMigration := createMigrationTable(dialect)
	scriptTable := dialect.addTable(tableMigration)
	dialect.execScript(scriptTable)

	dialect.insertProperty(PROP_HISTORY, "0000.00.00", TABLE_MIGRATION, ParamName, ParamValue)
}


func createMigrationTable(dialect GenericDialect) DbTable {
	isNullable := false
	tableMigration := DbTable{}
	tableMigration.columns = append(tableMigration.columns, DbColumn{}.initSqlPrimary("ID"))
	tableMigration.columns = append(tableMigration.columns, DbColumn{}.initSql(ParamName, dialect.getSqlType("string", &isNullable, ""), isNullable))
	tableMigration.columns = append(tableMigration.columns, DbColumn{}.initSql(ParamValue, dialect.getSqlType("string", &isNullable, ""), isNullable))

	tableMigration.initSql(TABLE_MIGRATION, tableMigration.columns[0])
	return tableMigration
}

/*
// Helper function to sort files
func (files []string) []string {
	sort.Strings(files)
	return files
}
*/