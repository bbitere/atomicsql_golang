package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

type DialectArg struct {
	strConnection string
	dirJSONs      string
	dirSQL        string
	delimeter     string
	bApplyScripts bool
}

type TDbTable struct {
	counterOrder int
	script       string
	nameDbTable  string
	tableInst    *DbTable
}

type TUpdColumn struct {
	script    string
	tableInst *DbTable
}

type PreGenerateScript struct {
	listAddTables  []*TDbTable
	listDropTables []*TDbTable
	listUpdTables  []*TUpdColumn
}

type GenSqlScriptsFromJsonModels struct {
	dictForeignKeys map[string]*FKRootTgt
	tables          map[string]*DbTable
}

func (d *DialectArg) NewDialectArg(strConnection, dirJSONs, dirSQL, delimeter string, bApplyScripts bool) *DialectArg {
	return &DialectArg{
		strConnection: strConnection,
		dirJSONs:      dirJSONs,
		dirSQL:        dirSQL,
		delimeter:     delimeter,
		bApplyScripts: bApplyScripts,
	}
}

func NewTDbTable(nameDbTable, script string) *TDbTable {
	return &TDbTable{
		nameDbTable: nameDbTable,
		script:      script,
		tableInst:   nil,
	}
}

func NewTDbTableWithInst(tableInst *DbTable, script string) *TDbTable {
	return &TDbTable{
		nameDbTable: tableInst.LangTableNameModel,
		script:      script,
		tableInst:   tableInst,
	}
}

func NewTUpdColumn(tableInst *DbTable, script string) *TUpdColumn {
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
		dictForeignKeys: make(map[string]*FKRootTgt),
		tables:          make(map[string]*DbTable),
	}
}

func (d *GenSqlScriptsFromJsonModels) GenerateScripts(sqlLang, dirJSONs, dirSQL, connection_string, delimeter string, bApplyScripts bool) {
	arg := NewDialectArg(connection_string, dirJSONs, dirSQL, delimeter, bApplyScripts)
	dialect := GetDialectByName(sqlLang)
	if dialect == nil {
		return
	}

	dialect.ProcessData(arg)
}

func (d *GenSqlScriptsFromJsonModels) ProcessData(dialect GenericDialect, arg1 GenericDialectArg) {
	arg, ok := arg1.(*DialectArg)
	if !ok {
		fmt.Println("Invalid DialectArg type")
		return
	}

	inst := NewMigrationDB()

	sqlFiles, err := ioutil.ReadDir(arg.dirSQL)
	if err != nil {
		dialect.PrintError(err.Error())
		return
	}

	for _, sqlFile := range sqlFiles {
		if !strings.Contains(sqlFile.Name(), PATTERN_NAME_SQL) {
			continue
		}

		dialect.SetCurrentFile(sqlFile.Name())
		err := os.Remove(sqlFile.Name())
		if err != nil {
			dialect.PrintError(err.Error())
		}
	}

	files, err := ioutil.ReadDir(arg.dirJSONs)
	if err != nil {
		dialect.PrintError(err.Error())
		return
	}

	files = sortFiles(files)

	var prevFile string
	for iFile, file := range files {
		if strings.Contains(file.Name(), "2023.09.25.125837.models.json") {
			UtilsNop()
		}

		var bMustDeleteJSON bool
		var tagFile string
		content := d.generateScript(prevFile, file.Name(), arg.delimeter, dialect, &tagFile, &bMustDeleteJSON)
		if bMustDeleteJSON {
			err := os.Remove(file.Name())
			if err != nil {
				dialect.PrintError(err.Error())
			}
			continue
		}
		if content == "" {
			continue
		}

		if arg.bApplyScripts {
			if strings.TrimSpace(content) != "" {
				inst.ApplyImport(dialect, file.Name(), content)
			}
		} else {
			etichet := tagFile
			if etichet != "" {
				etichet = "-" + etichet
			}
			fileName := UtilsGetFileInfoName(file)
			sqlFile := fmt.Sprintf("%s\\%s%s%s", arg.dirSQL, fileName, etichet, EXTENSION_SQL)
			err := ioutil.WriteFile(sqlFile, []byte(content), 0644)
			if err != nil {
				dialect.PrintError(err.Error())
			}
			name := UtilsGetFileInfoName(sqlFile)
			fmt.Printf("%03d. Generate file: %s\n", iFile, name)
		}

		prevFile = file.Name()
	}
}

func (d *GenSqlScriptsFromJsonModels) generateScript(prevFile, file, delimeter string, dialect GenericDialect, tagFile *string, bMustDeleteJSON *bool) string {
	var dbTables0, dbTables1 []*DbTable
	var err error

	if prevFile != "" {
		dbTables0, *bMustDeleteJSON = ReadCustomJSON(prevFile, delimeter, dialect)
		if *bMustDeleteJSON {
			return ""
		}
	}

	dbTables1, *bMustDeleteJSON = ReadCustomJSON(file, delimeter, dialect)
	if *bMustDeleteJSON {
		return ""
	}

	var preGenerate *PreGenerateScript
	if dbTables0 != nil {
		preGenerate = d.generatePartialSQL(dbTables1, dbTables0, dialect, tagFile)
	} else {
		preGenerate = d.generateNewSQLScript(dbTables1, dialect, tagFile)
	}

	return preGenerate.generateScript(dialect)
}


func (d *GenSqlScriptsFromJsonModels) generateNewSqlScript(dbTables []DbTable, dialect GenericDialect, tagFile *string) PreGenerateScript {
	ret := PreGenerateScript{}
	tableMigration := MigrationDB.createMigrationTable(dialect)
	scriptTable := dialect.addTable(tableMigration)
	ret.listAddTables = append(ret.listAddTables, TDbTable{MigrationDB.TABLE_MIGRATION, scriptTable})

	for _, table := range dbTables {
		txt := dialect.addTable(table)
		if txt != "" {
			ret.listAddTables = append(ret.listAddTables, TDbTable{table, txt})
		}
	}
	*tagFile = "initDb"
	return ret
}

func generatePartialSql(dbTables1, dbTables0 []DbTable, dialect GenericDialect, tagFile *string) PreGenerateScript {
	ret := PreGenerateScript{}
	s := ""
	dict1 := Utils_getDictFromList(dbTables1, func(x DbTable) string { return x.LangTableNameModel })
	dict0 := Utils_getDictFromList(dbTables0, func(x DbTable) string { return x.LangTableNameModel })

	dict1Sql := Utils_getDictFromList(dbTables1, func(x DbTable) string { return x.SqlTableNameModel })
	dict0Sql := Utils_getDictFromList(dbTables0, func(x DbTable) string { return x.SqlTableNameModel })

	var _tagFile, _tagFile2 string
	diffTables := Utils_getDictionaryDifference(dict0, dict1)
	diffTblSql := Utils_getDictionaryDifference(dict0Sql, dict1Sql)

	for _, table := range diffTables {
		if diffTables[table.Value.LangTableNameModel] && diffTblSql[table.Value.SqlTableNameModel] {
			txt := dialect.dropTable(table.Value)
			ret.listDropTables = append(ret.listDropTables, TDbTable{table.Value, txt})
			_tagFile += table.Value.LangTableNameModel
		} else {
			_tagFile2 += table.Value.LangTableNameModel
		}
	}
	s1 := strings.Join(ret.listDropTables, dialect.SqlSeparator())
	s += s1
	if _tagFile2 != "" {
		*tagFile = "renTable" + _tagFile2
	}
	if _tagFile != "" {
		*tagFile = "dropTable" + _tagFile
	}

	diffTables = Utils_getDictionaryDifference(dict1, dict0)
	diffTblSql = Utils_getDictionaryDifference(dict1Sql, dict0Sql)
	_tagFile, _tagFile2 = "", ""
	arr := []string{}
	for _, table := range diffTables {
		if diffTables[table.Value.LangTableNameModel] && diffTblSql[table.Value.SqlTableNameModel] {
			txt := dialect.addTable(table.Value)
			if txt != "" {
				ret.listDropTables = append(ret.listDropTables, TDbTable{table.Value, txt})
				_tagFile += table.Value.LangTableNameModel
			}
		} else {
			_tagFile2 += table.Value.LangTableNameModel
		}
	}
	s1 = strings.Join(arr, dialect.SqlSeparator())
	s += s1

	if _tagFile2 != "" {
		*tagFile = "renTable" + _tagFile2
	}
	if _tagFile != "" {
		*tagFile = "addTable" + _tagFile
	}

	_tagFile, _tagFile2 = "", ""
	commonTables := Utils_getDictionaryUnion(dict1, dict0)
	arr = []string{}
	tagUpdate := ""
	for _, pair := range commonTables {
		txt := generateDiffInsideTable(pair.Value.Item1, pair.Value.Item2, dialect, &tagUpdate)
		if txt != "" {
			_tagFile += pair.Value.Item1.LangTableNameModel
		}

		sql := dialect.updateTable(pair.Value.Item1, pair.Value.Item2)
		if sql != "" {
			if txt != "" {
				txt += dialect.SqlSeparator() + sql
			} else {
				txt += sql
			}

			_tagFile2 += pair.Value.Item1.LangTableNameModel
		}
		if txt != "" {
			ret.listUpdTables = append(ret.listUpdTables, TUpdColumn{pair.Value.Item1, txt})
		}
	}
	s1 = strings.Join(arr, dialect.SqlSeparator())
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
func generateDiffInsideTable(table1, table2 DbTable, dialect GenericDialect, tagUpdate *string) string {
	if table1.LangTableNameModel != table2.LangTableNameModel {
		return dialect.printError("not the same name")
	}

	if table1.LangTableNameModel == "Project" {
		Utils_Nop()
	}

	var s string
	dict1 := Utils_getDictFromList(table1.columns, func(x Column) string { return x.langName })
	dict0 := Utils_getDictFromList(table2.columns, func(x Column) string { return x.langName })

	dict1Sql := Utils_getDictFromList(table1.columns, func(x Column) string { return x.sqlName })
	dict0Sql := Utils_getDictFromList(table2.columns, func(x Column) string { return x.sqlName })

	commonColumns := Utils_getDictionaryUnion(dict1, dict0)
	commonColumnsSql := Utils_getDictionaryUnion(dict1Sql, dict0Sql)

	var arr []string
	{
		diffColumns := Utils_getDictionaryDifference(dict1, dict0)
		diffColSql := Utils_getDictionaryDifference(dict1Sql, dict0Sql)

		var _tagAddUpdate, _tagAddConstraint string
		for _, column := range diffColumns {
			if diffColumns[column.Value.langName] && diffColSql[column.Value.sqlName] {
				txt := dialect.addColumn(table1, column.Value)
				arr = append(arr, txt)
				_tagAddUpdate += column.Value.langName
			} else {
				fkeyField := commonColumnsSql[column.Value.sqlName]
				if fkeyField != nil {
					if fkeyField.Item1.langName2 == fkeyField.Item2.langName && fkeyField.Item1.langName != "" &&
						fkeyField.Item1.ForeignKey != nil && fkeyField.Item2.ForeignKey == nil {
						txt := dialect.addFKConstrictor(table1, fkeyField.Item1)
						if txt != "" {
							arr = append(arr, txt)
							_tagAddConstraint += fkeyField.Item1.langName
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

		diffColumns := Utils_getDictionaryDifference(dict0, dict1)
		diffColSql := Utils_getDictionaryDifference(dict0Sql, dict1Sql)
		for _, column := range diffColumns {
			if diffColumns[column.Value.langName] && diffColSql[column.Value.sqlName] {
				txt := dialect.dropColumn(table1, column.Value)
				if txt != "" {
					arr = append(arr, txt)
					_tagDropUpdate += column.Value.langName
				}
			} else {
				fkeyField := commonColumnsSql[column.Value.sqlName]
				if fkeyField != nil {
					if fkeyField.Item2.langName2 == fkeyField.Item1.langName && fkeyField.Item2.langName != "" &&
						fkeyField.Item2.ForeignKey != nil && fkeyField.Item1.ForeignKey == nil {
						txt := dialect.dropFKConstrictor(table1, fkeyField.Item2)
						if txt != "" {
							arr = append(arr, txt)
							_tagDropUpdate += fkeyField.Item2.langName
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
			txt := dialect.updateColumn(table1, pair.Value.Item1, pair.Value.Item2)
			if txt != "" {
				arr = append(arr, txt)
				_tagUpdColUpdate += pair.Value.Item1.langName
			}
		}
		if _tagUpdColUpdate != "" {
			*tagUpdate += "_updCol" + _tagUpdColUpdate
		}
	}

	s1 := strings.Join(arr, dialect.SqlSeparator())
	s += s1

	return s
}

/**
 * 
 * read the json where is the definition of structs/tables.
 */ 
 func ReadCustomJson(file, delimiter string, dialect GenericDialect, bMustDeleteJsonFile *bool) ([]DbTable, error) {
	if file == "" {
		return nil, nil
	}

	var dbTables []DbTable

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

			dbTable := DbTable{}
			dbTable.initSql(tableName, nil)
			dbTable.json = jsonContent // store here, and use it after
			dbTables = append(dbTables, dbTable)
		} else {
			fmt.Println("json definition wrong")
		}
	}

	dictSqlTables := make(map[string]DbTable)

	for _, dbTable := range dbTables {
		if dbTable.LangTableNameModel == "Project" {
			Utils_Nop()
		}

		var nextShouldBeFK string
		var nextShouldBeFKPointerType string

		obj := make(map[string]map[string]string)
		err := json.Unmarshal([]byte(dbTable.json), &obj)
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
				*bMustDeleteJsonFile = true
				dialect.printError(fmt.Sprintf("duplicate table at sql Name: '%s'. check if is duplicate or you miss to do in 2 steps the renaming of a table", dbTable.SqlTableNameModel))
			}
		}

		dictSqlColumns := make(map[string]DbColumn)

		for key, val := range obj {
			if key == TABLE_SQLNAME {
				continue
			}

			column := DbColumn{}
			colType := val[TAG_Type]
			tags := val[TAG_tags]
			var langName2 string
			langName := key
			var tblPointer DbTable

			sqlName := getSqlName(tags)
			if sqlName == "" {
				continue
			}
			if sqlName == "-" {
				nextShouldBeFK = key
				nextShouldBeFKPointerType = colType
				continue
			} else if nextShouldBeFK != "" {
				compressedName := GoModelTemplate.ConvertToIdent_GoLang(key)
				if nextShouldBeFK != compressedName {
					*bMustDeleteJsonFile = true
					dialect.printError(fmt.Sprintf("names dont match %s - %s", key, nextShouldBeFK))
					return nil, nil
				}
				nextShouldBeFK = ""
				tableName := GenericDialect.CleanNameGoStruct(nextShouldBeFKPointerType)
				tblPointer = findTableByName(dbTables, tableName)
				if tblPointer == nil {
					*bMustDeleteJsonFile = true
					dialect.printError(fmt.Sprintf("not found table %s", tableName))
					return nil, nil
				}
				const __ID = GoModelTemplate.POSTFIX_FOREIGNKEY
				if !strings.HasSuffix(sqlName, __ID) || !strings.HasSuffix(langName, __ID) {
					*bMustDeleteJsonFile = true
					dialect.printError(fmt.Sprintf("in table '%s' name of foreign key field (%s -> %s) should end in '%s'", dbTable.LangTableNameModel, langName, sqlName, __ID))
					return nil, nil
				}
				langName2 = langName
				langName = compressedName
			}

			bIsIdentity := strings.Contains(tags, "omitempty")
			if bIsIdentity {
				dbTable.PrimaryColumn = &column
			}

			bIsNullable := false
			sqlType := dialect.getSqlType(colType, &bIsNullable, sqlName)

			column.initLangSql(langName, langName2, colType, sqlName, sqlType, bIsNullable, bIsIdentity, tblPointer)

			if _, ok := dictSqlColumns[column.sqlName]; !ok {
				dictSqlColumns[column.sqlName] = column
			} else {
				*bMustDeleteJsonFile = true
				dialect.printError(fmt.Sprintf("duplicate column at sql Name: '%s'. check if is duplicate or you miss to do in 2 steps the renaming of a column", column.sqlName))
			}

			dbTable.columns = append(dbTable.columns, column)
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
		p := strings.Split(tags1, ',')
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

func isYes(s string) bool {
	return s != "" && s == "YES"
}

func findTableByName(dbTables []DbTable, tableName string) *DbTable {
	for _, dbTable := range dbTables {
		if dbTable.LangTableNameModel == tableName {
			return &dbTable
		}
	}
	return nil
}
