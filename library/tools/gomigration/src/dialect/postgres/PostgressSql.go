package gomigration_pstgress

import (
	"database/sql"
	"fmt"
	"strings"
	//dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration_dialect"
)

type PostgressDialect struct {
	connection *sql.DB
    base1 TGenericDialect;
}

func (pd *PostgressDialect) getSql() string {
	return "postgres"
}

func (pd *PostgressDialect) SqlSeparator() string {
	return ";"
}

func (pd *PostgressDialect) tokenizIdentif(identif string) string {
	return "\"" + identif + "\""
}

func (pd *PostgressDialect) tokenizTable(table *DbTable) string {
	return "public." + pd.tokenizIdentif(table.SqlTableNameModel)
}

func (pd *PostgressDialect) tokenizTableName(tableName string) string {
	return "public." + pd.tokenizIdentif(tableName)
}

func (pd *PostgressDialect) dropTable(table *DbTable) string {
	s := fmt.Sprintf("DROP TABLE %s", pd.tokenizTable(table))
	return s
}

func (pd *PostgressDialect) addTable(table *DbTable) string {
	NL := "\n"
	sqlTableName := table.SqlTableNameModel
	var columnsArr []string
	for _, col := range table.columns {
		columnsArr = append(columnsArr, pd._addColumn(col))
	}
	columnsDefs := strings.Join(columnsArr, ","+NL)
	colIDName := table.PrimaryColumn.sqlName

	tableAdd := fmt.Sprintf(`
	-------------------------------------------------------------------	
	CREATE TABLE IF NOT EXISTS %s
	(
		%s
	)
	TABLESPACE pg_default;`,
		pd.tokenizTable(table), columnsDefs)

	return tableAdd
}

func (pd *PostgressDialect) dropColumn(table *DbTable, column *DbColumn) string {
	s := fmt.Sprintf(`
	ALTER TABLE %s
	DROP COLUMN %s`, pd.tokenizTable(table), pd.tokenizIdentif(column.sqlName))
	return s
}

func (pd *PostgressDialect) addColumn(table *DbTable, column *DbColumn) string {
	colData := pd._addColumn(column)
	s := fmt.Sprintf(`
	ALTER TABLE %s
	ADD COLUMN %s`, pd.tokenizTable(table), colData)
	return s
}

func (pd *PostgressDialect) _addColumn(column *DbColumn) string {
	colName := column.sqlName
	if column.bIsIdentity {
		if column.langType == "int64" {
			return fmt.Sprintf("%s bigserial NOT NULL AUTO_INCREMENT", pd.tokenizIdentif(colName))
		}
		return fmt.Sprintf("%s serial PRIMARY KEY NOT NULL", pd.tokenizIdentif(colName))
	} else {
		sqlType := column.sqlType
		if column.ForeignKey != nil {
			targetTableID := column.ForeignKey.PrimaryColumn.sqlName
			fkFkName := pd.tokenizIdentif(fmt.Sprintf("fk_%s", column.sqlName))
			fkName := pd.tokenizIdentif(colName)
			return fmt.Sprintf("%s integer NULL CONSTRAINT %s REFERENCES %s (%s)", fkName, fkFkName, pd.tokenizTable(column.ForeignKey), targetTableID)
		} else {
			if column.bIsNullable {
				return fmt.Sprintf("%s %s NULL", pd.tokenizIdentif(colName), sqlType)
			}
			return fmt.Sprintf("%s %s NOT NULL", pd.tokenizIdentif(colName), sqlType)
		}
	}
}

func (pd *PostgressDialect) updateTable(table *DbTable, tablePrev *DbTable) string {
	if table.SqlTableNameModel != tablePrev.SqlTableNameModel {
		s := fmt.Sprintf(`
	ALTER TABLE public.%s
	RENAME TO %s`, pd.tokenizIdentif(tablePrev.SqlTableNameModel), pd.tokenizIdentif(table.SqlTableNameModel))
		return s
	}
	return ""
}

func (pd *PostgressDialect) updateColumn(table *DbTable, column *DbColumn, columnPrev *DbColumn) string {
	if columnPrev.langName != column.langName {
		return pd.printError("internal error")
	}

	if column.sqlType != columnPrev.sqlType {
		r := fmt.Sprintf(`
	ALTER TABLE %s 
	ALTER COLUMN %s TYPE %s`, pd.tokenizTable(table), pd.tokenizIdentif(column.sqlName), column.sqlType)
		return r
	}

	if column.sqlName != columnPrev.sqlName {
		s := fmt.Sprintf(`
	ALTER TABLE %s
	RENAME COLUMN %s TO %s`, pd.tokenizTable(table), pd.tokenizIdentif(columnPrev.sqlName), pd.tokenizIdentif(column.sqlName))
		return s
	}

	return ""
}

func (pd *PostgressDialect) addFKConstrictor(table *DbTable, column *DbColumn) string {
	tablename := pd.tokenizTable(column.ForeignKey)
	tableID := pd.tokenizIdentif(column.ForeignKey.PrimaryColumn.sqlName)
	myTableName := pd.tokenizTable(table)
	fkName := pd.tokenizIdentif(column.sqlName)
	fkFkName := pd.tokenizIdentif(fmt.Sprintf("fk_%s", column.sqlName))
	s := fmt.Sprintf(`
	ALTER TABLE %s
	ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s);`,
		myTableName, fkFkName, fkName, tablename, tableID)
	return s
}

func (pd *PostgressDialect) dropFKConstrictor(table *DbTable, column *DbColumn) string {
	myTableName := pd.tokenizTable(table)
	fkFkName := pd.tokenizIdentif(fmt.Sprintf("fk_%s", column.sqlName))
	s := fmt.Sprintf(`
	ALTER TABLE %s
	DROP CONSTRAINT %s;`, myTableName, fkFkName)
	return s
}

func (pd *PostgressDialect) getSqlType(langType string, bIsNullable *bool, nameOfColumn string) string {
	langType = pd.cleanNameGoStruct(langType)
	if strings.HasPrefix(langType, "[]") {
		return pd.printError(fmt.Sprintf("type %s not supported!", langType))
	}

	switch langType {
	case "NullString":
		*bIsNullable = true
		return pd.isLongType(nameOfColumn, "TEXT", "VARCHAR")
	case "NullBool":
		*bIsNullable = true
		return "BOOLEAN"
	case "NullByte", "NullInt16":
		*bIsNullable = true
		return "SMALLINT"
	case "NullInt", "NullInt32":
		*bIsNullable = true
		return "INT"
	case "NullInt64":
		*bIsNullable = true
		return "BIGINT"
	case "NullFloat64":
		*bIsNullable = true
		return "double precision"
	case "time.NullTime", "NullTime":
		*bIsNullable = true
		return "TIMESTAMP"
	case "string":
		return pd.isLongType(nameOfColumn, "TEXT", "VARCHAR")
	case "char":
		return "CHAR"
	case "bool":
		return "BOOLEAN"
	case "byte", "int16":
		return "SMALLINT"
	case "int", "int32":
		return "INT"
	case "int64":
		return "BIGINT"
	case "Time":
		return "TIMESTAMP"
	case "UUID", "uuid.UUID":
		return "UUID"
	case "float32":
		return "real"
	case "float64":
		return "double precision"
	default:
		return pd.printError(fmt.Sprintf("not supported type %s", langType))
	}
}

func (pd *PostgressDialect) getProperty(propName, tableName, colName, colValue string) string {
	requestText := fmt.Sprintf(`
	SELECT %s FROM %s
	WHERE %s = '%s'`, pd.tokenizIdentif(colValue), pd.tokenizTableName(tableName), pd.tokenizIdentif(colName), propName)

	row := pd.connection.QueryRow(requestText)
	var obj interface{}
	if err := row.Scan(&obj); err != nil {
		return ""
	}
	return fmt.Sprintf("%v", obj)
}

func (pd *PostgressDialect) updateProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
	UPDATE %s
	SET %s = '%s'
	WHERE %s = '%s'`, pd.tokenizIdentif(tableName), pd.tokenizIdentif(colValue), value, pd.tokenizIdentif(colName), propName)

	pd.connection.Exec(requestText)
}

func (pd *PostgressDialect) insertProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
	INSERT INTO %s (%s, %s, %s) VALUES (default, '%s', '%s')`, pd.tokenizTableName(tableName), pd.tokenizIdentif("ID"), pd.tokenizIdentif(colName), pd.tokenizIdentif(colValue), propName, value)

	pd.connection.Exec(requestText)
}

func (pd *PostgressDialect) execScript(scriptTxt string) {
	parts := strings.Split(scriptTxt, ";")
	for i, part := range parts {
		script1 := strings.TrimSpace(part)
		if len(script1) > 0 {
			_, err := pd.connection.Exec(script1)
			if err != nil {
				msg := err.Error()
				fmt.Printf("Error exec script part %d: %s\n", i, msg)
				panic(err)
			}
		}
	}
}

func (pd *PostgressDialect) startConnection(arg *GenericDialectArg) bool {
	connectionString := arg.connection_string
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return false
	}

	pd.connection = conn

	if true {
		pd.fnProcessData(pd, arg)
		return true
	}

	tx, err := conn.Begin()
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return false
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	pd.fnProcessData(pd, arg)

	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		conn.Close()
		return false
	}

	return true
}

