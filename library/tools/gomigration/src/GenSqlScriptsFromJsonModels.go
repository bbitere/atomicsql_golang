package main

import (
	"encoding/json"
	"filepath"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	dialect_general "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect/general"
	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"
)

const (
	TAG_Type       = "Type"
	TAG_tags       = "tags"
	TABLE_SQLNAME  = "_tblsql"
	EXTENSION_JSON = ".json"
	EXTENSION_SQL  = ".sql"
	PATTERN_NAME_SQL = ".models-"
)

const (
	CONCAT_DOT = "."
)

type DialectArg1 struct {
	Base1 dialect.GenericDialectArg
	dirJSONs      string
	dirSQL        string
	delimeter     string
	bApplyScripts bool
}

func (m *DialectArg1) GetGenericDialectArg() *dialect.GenericDialectArg{

	return &m.Base1;
}

func  NewDialectArg(strConnection, dirJSONs, dirSQL, delimeter string, bApplyScripts bool) *DialectArg1 {

	return &DialectArg1{
		Base1: dialect.GenericDialectArg{ Connection_String:  strConnection,},
		dirJSONs:      dirJSONs,
		dirSQL:        dirSQL,
		delimeter:     delimeter,
		bApplyScripts: bApplyScripts,
	}
}



type TDbTable struct {
	counterOrder int
	script       string
	nameDbTable  string
	tableInst    *dialect.DbTable
}

func (This*TDbTable) Constructor1( nameDbTable string,  script string) *TDbTable{
            
	This.nameDbTable = nameDbTable;
	This.script = script;
	This.tableInst = nil;
	return This;
}

func (This*TDbTable) Constructor2( tableInst* dialect.DbTable,  script string) *TDbTable{
            
	This.nameDbTable = tableInst.LangTableNameModel;
	This.script = script;
	This.tableInst = tableInst;
	return This;
}

    

type TUpdColumn struct {
	script    string
	tableInst *dialect.DbTable
}

func (This*TUpdColumn) Constructor( tableInst* dialect.DbTable,  script string) *TUpdColumn{
            
	This.script = script;
	This.tableInst = tableInst;
	return This;
}


type PreGenerateScript struct {
	listAddTables  []*TDbTable
	listDropTables []*TDbTable
	listUpdTables  []*TUpdColumn
}

type GenSqlScriptsFromJsonModels struct {
	dictForeignKeys map[string]*dialect.FKRootTgt
	tables          map[string]*dialect.DbTable
}


func NewTDbTable(nameDbTable, script string) *TDbTable {
	return &TDbTable{
		nameDbTable: nameDbTable,
		script:      script,
		tableInst:   nil,
	}
}

func NewTDbTableWithInst(tableInst *dialect.DbTable, script string) *TDbTable {
	return &TDbTable{
		nameDbTable: tableInst.LangTableNameModel,
		script:      script,
		tableInst:   tableInst,
	}
}

func NewTUpdColumn(tableInst *dialect.DbTable, script string) *TUpdColumn {
	return &TUpdColumn{
		tableInst: tableInst,
		script:    script,
	}
}

func NewPreGenerateScript() *PreGenerateScript {
	return &PreGenerateScript{
		listAddTables:  make([]*TDbTable, 0),
		listDropTables: make([]*TDbTable, 0),
		listUpdTables:  make([]*TUpdColumn, 0),
	}
}

func NewGenSqlScriptsFromJsonModels() *GenSqlScriptsFromJsonModels {
	return &GenSqlScriptsFromJsonModels{
		dictForeignKeys: make(map[string]*dialect.FKRootTgt),
		tables:          make(map[string]*dialect.DbTable),
	}
}

func (d *GenSqlScriptsFromJsonModels) GenerateScripts(
		sqlLang, dirJSONs, dirSQL, 
		connection_string, delimeter string, bApplyScripts bool) bool {

	arg := NewDialectArg(connection_string, dirJSONs, dirSQL, delimeter, bApplyScripts)
	dialectInst := dialect_general.GetDialectByName(sqlLang)
	if dialectInst == nil {
		return false;
	}

	dialectInst.GetGenericDialect().FnProcessData = d.ProcessData;
	return dialectInst.StartConnection( arg );
}


func (d *GenSqlScriptsFromJsonModels) ProcessData(dialectInst dialect.GenericDialect, arg1 dialect.IGenericDialectArg){
	arg, ok := arg1.(*DialectArg1)
	if !ok {
		fmt.Println("Invalid DialectArg type")
		return
	}

	inst := &MigrationDB{};

	sqlFiles, err := ioutil.ReadDir(arg.dirSQL)
	if err != nil {
		dialectInst.PrintError(err.Error())
		return
	}

	for _, sqlFile := range sqlFiles {
		if !strings.Contains(sqlFile.Name(), PATTERN_NAME_SQL) {
			continue
		}

		dialectInst.SetCurrentFile(sqlFile.Name())
		err := os.Remove(sqlFile.Name())
		if err != nil {
			dialectInst.PrintError(err.Error())
		}
	}

	files, err := d.readFilesFromDir( dialectInst, arg.dirSQL );
	if err != nil {
		dialectInst.PrintError(err.Error())
		return
	}

	//files = sort.Strings(files);
	sort.Strings(files);

	var prevFile string
	for iFile, file := range files {
		if strings.Contains(file, "2023.09.25.125837.models.json") {
			utils.Utils_Utils_Nop();
		}

		var bMustDeleteJSON = []bool{false};
		var tagFile string
		content := d.generateScript(prevFile, file, arg.delimeter, dialectInst, &tagFile, bMustDeleteJSON)
		if bMustDeleteJSON[0] {
			err := os.Remove(file)
			if err != nil {
				dialectInst.PrintError(err.Error())
			}
			continue
		}
		if content == "" {
			continue
		}

		if arg.bApplyScripts {
			if strings.TrimSpace(content) != "" {
				inst.ApplyImport(dialectInst, file, content)
			}
		} else {
			etichet := tagFile
			if etichet != "" {
				etichet = "-" + etichet
			}
			fileName := utils.Utils_GetFileInfoName(file)
			sqlFile := fmt.Sprintf("%s\\%s%s%s", arg.dirSQL, fileName, etichet, EXTENSION_SQL)
			err := ioutil.WriteFile(sqlFile, []byte(content), 0644)
			if err != nil {
				dialectInst.PrintError(err.Error())
			}
			name := utils.Utils_GetFileInfoName(sqlFile)
			fmt.Printf("%03d. Generate file: %s\n", iFile, name)
		}

		prevFile = file
	}
}

func (m *GenSqlScriptsFromJsonModels) readFilesFromDir(dialectInst dialect.GenericDialect, dir string) ([]string, error) {

	//var files []string

	// Încercați să obțineți fișierele SQL și să le ștergeți
	sqlFiles, err := filepath.Glob(filepath.Join( dir, "*"+EXTENSION_SQL))
	if err != nil {
		fmt.Println("Eroare la obținerea fișierelor SQL:", err)
		return nil, err;
	}

	for _, sqlFile := range sqlFiles {
		if !strings.Contains(sqlFile, PATTERN_NAME_SQL) {
			continue
		}
		dialectInst.SetCurrentFile(sqlFile)
		err := os.Remove(sqlFile)
		if err != nil {
			fmt.Printf("Eroare la ștergerea fișierului SQL %s: %s\n", sqlFile, err)
		}
	}

	// Încercați să obțineți fișierele JSON
	jsonFiles, err := filepath.Glob(filepath.Join( dir, "*"+EXTENSION_JSON))
	if err != nil {
		fmt.Println("Eroare la obținerea fișierelor JSON:", err)
		return nil, err
	}

	return jsonFiles, nil
}


func (d *GenSqlScriptsFromJsonModels) generateScript(prevFile, file, delimeter string, dialectInst dialect.GenericDialect, tagFile *string, bMustDeleteJSON []bool) string {

	var dbTables0, dbTables1 []*dialect.DbTable
	//var err error

	dbTables0, _ = ReadCustomJson(prevFile, delimeter, dialectInst, bMustDeleteJSON)
	dbTables1, _ = ReadCustomJson(file, delimeter, dialectInst, bMustDeleteJSON)
	if bMustDeleteJSON[0] {
		return ""
	}

	var preGenerate *PreGenerateScript
	if dbTables0 != nil {
		preGenerate = d.generatePartialSQL(dbTables1, dbTables0, dialectInst, tagFile)
	} else {
		preGenerate = d.generateNewSQLScript(dbTables1, dialectInst, tagFile)
	}

	return preGenerate.generateScript(dialectInst)
}


func (d *GenSqlScriptsFromJsonModels) generateNewSqlScript(
	dbTables []*dialect.DbTable, 
	dialectInst dialect.GenericDialect, 
	tagFile *string) PreGenerateScript {


	ret := PreGenerateScript{}
	tableMigration := MigrationDB.CreateMigrationTable(dialectInst)
	scriptTable := dialectInst.AddTable(tableMigration)
	ret.listAddTables = append(ret.listAddTables, (new (TDbTable)).Constructor2( MigrationDB.TABLE_MIGRATION, scriptTable) )

	for _, table := range dbTables {
		txt := dialectInst.AddTable(table)
		if txt != "" {
			ret.listAddTables = append(ret.listAddTables, (new (TDbTable)).Constructor2(table, txt))
		}
	}
	*tagFile = "initDb"
	return ret
}

func (d *GenSqlScriptsFromJsonModels)  generatePartialSql(
	dbTables1 []*dialect.DbTable,
	 dbTables0 []*dialect.DbTable, 
	dialectInst dialect.GenericDialect, 
	tagFile *string) PreGenerateScript {


	ret := PreGenerateScript{}
	s := ""
	dict1 := utils.Utils_GetDictFromList(dbTables1, func(x *dialect.DbTable) string { return x.LangTableNameModel })
	dict0 := utils.Utils_GetDictFromList(dbTables0, func(x *dialect.DbTable) string { return x.LangTableNameModel })

	dict1Sql := utils.Utils_GetDictFromList(dbTables1, func(x *dialect.DbTable) string { return x.SqlTableNameModel })
	dict0Sql := utils.Utils_GetDictFromList(dbTables0, func(x *dialect.DbTable) string { return x.SqlTableNameModel })

	var _tagFile, _tagFile2 string
	diffTables := utils.Utils_GetDictionaryDifference(dict0, dict1)
	diffTblSql := utils.Utils_GetDictionaryDifference(dict0Sql, dict1Sql)

	for _, table := range diffTables {

		var _, hasLang = diffTables[table.LangTableNameModel];
		var _, hasSql  = diffTblSql[table.SqlTableNameModel];

		if hasLang && hasSql {
			txt := dialectInst.DropTable(table)
			ret.listDropTables = append(ret.listDropTables, 
							(new (TDbTable)).Constructor2(table, txt) )
			_tagFile += table.LangTableNameModel
		} else {
			_tagFile2 += table.LangTableNameModel
		}
	}
	s1 := strings.Join(ret.listDropTables, dialectInst.SqlSeparator())
	s += s1
	if _tagFile2 != "" {
		*tagFile = "renTable" + _tagFile2
	}
	if _tagFile != "" {
		*tagFile = "dropTable" + _tagFile
	}

	diffTables = utils.Utils_GetDictionaryDifference(dict1, dict0)
	diffTblSql = utils.Utils_GetDictionaryDifference(dict1Sql, dict0Sql)
	_tagFile, _tagFile2 = "", ""
	arr := []string{}
	for _, table := range diffTables {

		var _, hasLang = diffTables[table.LangTableNameModel];
		var _, hasSql  = diffTblSql[table.SqlTableNameModel];

		if hasLang && hasSql {
		//if diffTables[table.LangTableNameModel] && diffTblSql[table.SqlTableNameModel] {
			txt := dialectInst.AddTable(table)
			if txt != "" {
				ret.listDropTables = append(ret.listDropTables, (new (TDbTable)).Constructor2(table, txt))
				_tagFile += table.LangTableNameModel
			}
		} else {
			_tagFile2 += table.LangTableNameModel
		}
	}
	s1 = strings.Join(arr, dialectInst.SqlSeparator())
	s += s1

	if _tagFile2 != "" {
		*tagFile = "renTable" + _tagFile2
	}
	if _tagFile != "" {
		*tagFile = "addTable" + _tagFile
	}

	_tagFile, _tagFile2 = "", ""
	commonTables := utils.Utils_GetDictionaryUnion(dict1, dict0)

	arr = []string{}
	tagUpdate := ""
	for _, pair := range commonTables {
		txt := generateDiffInsideTable(pair.Item1, pair.Item2, dialectInst, &tagUpdate)
		if txt != "" {
			_tagFile += pair.Item1.LangTableNameModel
		}

		sql := dialectInst.UpdateTable(pair.Item1, pair.Item2)
		if sql != "" {
			if txt != "" {
				txt += dialectInst.SqlSeparator() + sql
			} else {
				txt += sql
			}

			_tagFile2 += pair.Item1.LangTableNameModel
		}
		if txt != "" {
			ret.listUpdTables = append(ret.listUpdTables, (new (TUpdColumn)).Constructor(pair.Item1, txt) )
		}
	}
	s1 = strings.Join(arr, dialectInst.SqlSeparator())
	s += s1

	if _tagFile2 != "" {
		*tagFile = "renTable" + _tagFile2
	}

	if _tagFile != "" {
		if len(tagUpdate) < 50 {
			*tagFile = "upd" + _tagFile + tagUpdate
		} else {
			*tagFile = "updTable" + _tagFile + "_multiple"
		}
	}

	return ret
}
func generateDiffInsideTable( table1 *dialect.DbTable, table2 *dialect.DbTable,
	 dialectInst dialect.GenericDialect, tagUpdate *string) string {

	if table1.LangTableNameModel != table2.LangTableNameModel {
		return dialectInst.PrintError("not the same name")
	}

	if table1.LangTableNameModel == "Project" {
		utils.Utils_Utils_Nop()
	}

	var s string
	dict1 := utils.Utils_GetDictFromList(table1.Columns, func(x *dialect.DbColumn) string { return x.LangName })
	dict0 := utils.Utils_GetDictFromList(table2.Columns, func(x *dialect.DbColumn) string { return x.LangName })

	dict1Sql := utils.Utils_GetDictFromList(table1.Columns, func(x *dialect.DbColumn) string { return x.SqlName })
	dict0Sql := utils.Utils_GetDictFromList(table2.Columns, func(x *dialect.DbColumn) string { return x.SqlName })

	commonColumns    := utils.Utils_GetDictionaryUnion(dict1, dict0)
	commonColumnsSql := utils.Utils_GetDictionaryUnion(dict1Sql, dict0Sql)

	var arr []string
	{
		diffColumns := utils.Utils_GetDictionaryDifference(dict1, dict0)
		diffColSql  := utils.Utils_GetDictionaryDifference(dict1Sql, dict0Sql)

		var _tagAddUpdate, _tagAddConstraint string
		for _, column := range diffColumns {

			_, hasLangName := diffColumns[column.LangName];
			_, hasSqlName  := diffColSql[column.SqlName];

			if hasLangName && hasSqlName{

				txt := dialectInst.AddColumn(table1, column)
				arr = append(arr, txt)
				_tagAddUpdate += column.LangName
			} else {
				var fkeyField, hasKfld = commonColumnsSql[column.SqlName]
				if hasKfld {
					if fkeyField.Item1.LangName2 == fkeyField.Item2.LangName && fkeyField.Item1.LangName != "" &&
						fkeyField.Item1.ForeignKey != nil && fkeyField.Item2.ForeignKey == nil {
						txt := dialectInst.AddFKConstrictor(table1, fkeyField.Item1)
						if txt != "" {
							arr = append(arr, txt)
							_tagAddConstraint += fkeyField.Item1.LangName
						}
					}
				}
			}
		}

		if _tagAddUpdate != "" {
			*tagUpdate += "_addCol" + _tagAddUpdate
		}
		if _tagAddConstraint != "" {
			*tagUpdate += "_addFk" + _tagAddConstraint
		}
	}

	{
		var _tagDropUpdate, _tagDropConstraint string

		diffColumns := utils.Utils_GetDictionaryDifference(dict0, dict1)
		diffColSql := utils.Utils_GetDictionaryDifference(dict0Sql, dict1Sql)
		for _, column := range diffColumns {

			_, hasLangName := diffColumns[column.LangName];
			_, hasSqlName  := diffColSql[ column.SqlName];

			if hasLangName &&hasSqlName {
				txt := dialectInst.DropColumn(table1, column)
				if txt != "" {
					arr = append(arr, txt)
					_tagDropUpdate += column.LangName
				}
			} else {

				fkeyField, hasKFld := commonColumnsSql[column.SqlName]
				if hasKFld {
					if fkeyField.Item2.LangName2 == fkeyField.Item1.LangName && fkeyField.Item2.LangName != "" &&
						fkeyField.Item2.ForeignKey != nil && fkeyField.Item1.ForeignKey == nil {
						txt := dialectInst.DropFKConstrictor(table1, fkeyField.Item2)
						if txt != "" {
							arr = append(arr, txt)
							_tagDropUpdate += fkeyField.Item2.LangName
						}
					}
				}
			}
		}
		if _tagDropUpdate != "" {
			*tagUpdate += "_delCol" + _tagDropUpdate
		}
		if _tagDropConstraint != "" {
			*tagUpdate += "_delFk" + _tagDropConstraint
		}
	}

	{
		var _tagUpdColUpdate string
		for _, pair := range commonColumns {
			txt := dialectInst.UpdateColumn(table1, pair.Item1, pair.Item2)
			if txt != "" {
				arr = append(arr, txt)
				_tagUpdColUpdate += pair.Item1.LangName
			}
		}
		if _tagUpdColUpdate != "" {
			*tagUpdate += "_updCol" + _tagUpdColUpdate
		}
	}

	s1 := strings.Join(arr, dialectInst.SqlSeparator())
	s += s1

	return s
}

/**
 * 
 * read the json where is the definition of structs/tables.
 */ 
 func ReadCustomJson(file string, delimiter string, dialectInst dialect.GenericDialect,
	bMustDeleteJsonFile []bool) ([]*dialect.DbTable, error) {

	if file == "" {
		return nil, nil
	}

	var dbTables []*dialect.DbTable

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(content), delimiter)

	for _, part := range parts {
		text := strings.TrimSpace(part)
		if text == "" {
			continue
		}

		idx := strings.Index(text, "={")
		if idx >= 0 {
			tableName := text[:idx]
			jsonContent := text[idx+2:]

			var dbTable = new (dialect.DbTable)
			dbTable.InitSql(tableName, nil)
			dbTable.Json = jsonContent // store here, and use it after
			dbTables = append(dbTables, dbTable)
		} else {
			fmt.Println("json definition wrong")
		}
	}

	dictSqlTables := make(map[string]*dialect.DbTable)

	for _, dbTable := range dbTables {
		if dbTable.LangTableNameModel == "Project" {
			utils.Utils_Utils_Nop()
		}

		var nextShouldBeFK string
		var nextShouldBeFKPointerType string

		obj := make(map[string]map[string]string)
		err := json.Unmarshal([]byte(dbTable.Json), &obj)
		if err != nil {
			return nil, err
		}

		sqlTableName := dbTable.LangTableNameModel
		if val, ok := obj[TABLE_SQLNAME]; ok {
			sqlTableName = decodeSqlTableName(val[TAG_tags])
		}
		if sqlTableName != "" {
			dbTable.SqlTableNameModel = sqlTableName
			if _, ok := dictSqlTables[dbTable.SqlTableNameModel]; !ok {
				dictSqlTables[dbTable.SqlTableNameModel] = dbTable
			} else {
				bMustDeleteJsonFile[0] = true
				dialectInst.PrintError(fmt.Sprintf("duplicate table at sql Name: '%s'. check if is duplicate or you miss to do in 2 steps the renaming of a table", dbTable.SqlTableNameModel))
			}
		}

		dictSqlColumns := make(map[string] *dialect.DbColumn)

		for key, val := range obj {
			if key == TABLE_SQLNAME {
				continue
			}

			var column = new (dialect.DbColumn)
			colType := val[TAG_Type]
			tags := val[TAG_tags]
			var langName2 string
			langName := key
			var tblPointer *dialect.DbTable

			sqlName := getSqlName(tags)
			if sqlName == "" {
				continue
			}
			if sqlName == "-" {
				nextShouldBeFK = key
				nextShouldBeFKPointerType = colType
				continue
			} else if nextShouldBeFK != "" {
				compressedName := utils.Utils_ConvertToIdentGoLang(key, true)
				if nextShouldBeFK != compressedName {
					bMustDeleteJsonFile[0] = true
					dialectInst.PrintError(fmt.Sprintf("names dont match %s - %s", key, nextShouldBeFK))
					return nil, nil
				}
				nextShouldBeFK = ""
				tableName := dialectInst.CleanNameGoStruct(nextShouldBeFKPointerType)
				tblPointer = findTableByName(dbTables, tableName)
				if tblPointer == nil {
					bMustDeleteJsonFile[0] = true
					dialectInst.PrintError(fmt.Sprintf("not found table %s", tableName))
					return nil, nil
				}
				const __ID = utils.POSTFIX_FOREIGNKEY
				if !strings.HasSuffix(sqlName, __ID) || !strings.HasSuffix(langName, __ID) {
					bMustDeleteJsonFile[0] = true
					dialectInst.PrintError(fmt.Sprintf("in table '%s' name of foreign key field (%s -> %s) should end in '%s'", dbTable.LangTableNameModel, langName, sqlName, __ID))
					return nil, nil
				}
				langName2 = langName
				langName = compressedName
			}

			bIsIdentity := strings.Contains(tags, "omitempty")
			if bIsIdentity {
				dbTable.PrimaryColumn = column
			}

			bIsNullable := false
			sqlType := dialectInst.GetSqlType(colType, &bIsNullable, sqlName)

			column.InitLangSql(langName, langName2, colType, sqlName, sqlType, bIsNullable, bIsIdentity, tblPointer)

			if _, ok := dictSqlColumns[column.SqlName]; !ok {
				dictSqlColumns[column.SqlName] = column
			} else {
				bMustDeleteJsonFile[0] = true
				dialectInst.PrintError(fmt.Sprintf("duplicate column at sql Name: '%s'. check if is duplicate or you miss to do in 2 steps the renaming of a column", column.SqlName))
			}

			dbTable.Columns = append(dbTable.Columns, column)
		}
	}

	return dbTables, nil
}

/**
 * 
 * `json:"ID,omitempty"`
 * `json:"UUID"`
 *  `json:"-"`
 */ 
func getSqlName(tags string) string {
	if tags == "" {
		return ""
	}

	if tags == "-" {
		return ""
	}

	if strings.HasPrefix(tags, "\"json:\\\"") {
		tags1 := strings.Replace(tags, "\"json:\\\"", "", -1)
		tags1 = strings.Replace(tags1, "\\\"\"", "", -1)
		p := strings.Split(tags1, ",")
		return strings.TrimSpace(p[0])
	} else {
		fmt.Printf("json definition of description field is incomplete: %s\n", tags)
	}
	return ""
}

/**
 * 
 * atomicsql-table:"user_role"
 */ 
func decodeSqlTableName(tags string) string {
	if tags == "" {
		return ""
	}

	if strings.HasPrefix(tags, "/*atomicsql-table:\"") {
		tags1 := strings.Replace(tags, "/*atomicsql-table:\"", "", -1)
		tags1 = strings.Replace(tags1, "\"*/", "", -1)
		return tags1
	} else {
		fmt.Printf("json definition of atomicsql-table incorrect: %s\n", tags)
	}
	return ""
}


func findTableByName(dbTables []*dialect.DbTable, tableName string) *dialect.DbTable {
	for _, dbTable := range dbTables {
		if dbTable.LangTableNameModel == tableName {
			return dbTable
		}
	}
	return nil
}
