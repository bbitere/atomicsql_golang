package gomigration_pstgress

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"
)

type PostgressDialect struct {
	Connection *sql.DB
    Base1 dialect.TGenericDialect;
}

func (pd *PostgressDialect) GetGenericDialect() *dialect.TGenericDialect{

    return &pd.Base1;
}

func (pd *PostgressDialect) GetSql() string {
	return "postgres"
}

func (pd *PostgressDialect) SqlSeparator() string {
	return ";"
}

func (pd *PostgressDialect) TokenizIdentif(identif string) string {
	return "\"" + identif + "\""
}

func (pd *PostgressDialect) TokenizTable(table *dialect.DbTable) string {
	return "public." + pd.TokenizIdentif(table.SqlTableNameModel)
}

func (pd *PostgressDialect) TokenizTableName(tableName string) string {
	return "public." + pd.TokenizIdentif(tableName)
}

func (pd *PostgressDialect) DropTable(table *dialect.DbTable) string {
	s := fmt.Sprintf("DROP TABLE %s", pd.TokenizTable(table))
	return s
}

func (pd *PostgressDialect) AddTable(table *dialect.DbTable) string {
	NL := "\n"
	//sqlTableName := table.SqlTableNameModel
	var columnsArr []string
	for _, col := range table.Columns {
		columnsArr = append(columnsArr, pd._addColumn(col))
	}
	columnsDefs := strings.Join(columnsArr, ","+NL)
	//colIDName := table.PrimaryColumn.SqlName

	tableAdd := fmt.Sprintf(`
	-------------------------------------------------------------------	
	CREATE TABLE IF NOT EXISTS %s
	(
		%s
	)
	TABLESPACE pg_default;`,
		pd.TokenizTable(table), columnsDefs)

	return tableAdd
}

func (pd *PostgressDialect) DropColumn(table *dialect.DbTable, column *dialect.DbColumn) string {
	s := fmt.Sprintf(`
	ALTER TABLE %s
	DROP COLUMN %s`, pd.TokenizTable(table), pd.TokenizIdentif(column.SqlName))
	return s
}

func (pd *PostgressDialect) AddColumn(table *dialect.DbTable, column *dialect.DbColumn) string {
	colData := pd._addColumn(column)
	s := fmt.Sprintf(`
	ALTER TABLE %s
	ADD COLUMN %s`, pd.TokenizTable(table), colData)
	return s
}

func (pd *PostgressDialect) _addColumn(column *dialect.DbColumn) string {
	colName := column.SqlName
	if column.IsIdentity {
		if column.LangType == "int64" {
			return fmt.Sprintf("%s bigserial NOT NULL AUTO_INCREMENT", pd.TokenizIdentif(colName))
		}
		return fmt.Sprintf("%s serial PRIMARY KEY NOT NULL", pd.TokenizIdentif(colName))
	} else {
		sqlType := column.SqlType
		if column.ForeignKey != nil {
			targetTableID := column.ForeignKey.PrimaryColumn.SqlName
			fkFkName := pd.TokenizIdentif(fmt.Sprintf("fk_%s", column.SqlName))
			fkName := pd.TokenizIdentif(colName)
			return fmt.Sprintf("%s integer NULL CONSTRAINT %s REFERENCES %s (%s)", fkName, fkFkName, pd.TokenizTable(column.ForeignKey), targetTableID)
		} else {
			if column.IsNullable {
				return fmt.Sprintf("%s %s NULL", pd.TokenizIdentif(colName), sqlType)
			}
			return fmt.Sprintf("%s %s NOT NULL", pd.TokenizIdentif(colName), sqlType)
		}
	}
}

func (pd *PostgressDialect) UpdateTable(table *dialect.DbTable, tablePrev *dialect.DbTable) string {
	if table.SqlTableNameModel != tablePrev.SqlTableNameModel {
		s := fmt.Sprintf(`
	ALTER TABLE public.%s
	RENAME TO %s`, pd.TokenizIdentif(tablePrev.SqlTableNameModel), pd.TokenizIdentif(table.SqlTableNameModel))
		return s
	}
	return ""
}

func (pd *PostgressDialect) UpdateColumn(table *dialect.DbTable, column *dialect.DbColumn, columnPrev *dialect.DbColumn) string {
	if columnPrev.LangName != column.LangName {
		return pd.Base1.PrintError("internal error")
	}

	if column.SqlType != columnPrev.SqlType {
		r := fmt.Sprintf(`
	ALTER TABLE %s 
	ALTER COLUMN %s TYPE %s`, pd.TokenizTable(table), pd.TokenizIdentif(column.SqlName), column.SqlType)
		return r
	}

	if column.SqlName != columnPrev.SqlName {
		s := fmt.Sprintf(`
	ALTER TABLE %s
	RENAME COLUMN %s TO %s`, pd.TokenizTable(table), pd.TokenizIdentif(columnPrev.SqlName), pd.TokenizIdentif(column.SqlName))
		return s
	}

	return ""
}

func (pd *PostgressDialect) AddFKConstrictor(table *dialect.DbTable, column *dialect.DbColumn) string {
	tablename := pd.TokenizTable(column.ForeignKey)
	tableID := pd.TokenizIdentif(column.ForeignKey.PrimaryColumn.SqlName)
	myTableName := pd.TokenizTable(table)
	fkName := pd.TokenizIdentif(column.SqlName)
	fkFkName := pd.TokenizIdentif(fmt.Sprintf("fk_%s", column.SqlName))
	s := fmt.Sprintf(`
	ALTER TABLE %s
	ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s);`,
		myTableName, fkFkName, fkName, tablename, tableID)
	return s
}

func (pd *PostgressDialect) DropFKConstrictor(table *dialect.DbTable, column *dialect.DbColumn) string {
	myTableName := pd.TokenizTable(table)
	fkFkName := pd.TokenizIdentif(fmt.Sprintf("fk_%s", column.SqlName))
	s := fmt.Sprintf(`
	ALTER TABLE %s
	DROP CONSTRAINT %s;`, myTableName, fkFkName)
	return s
}

func (pd *PostgressDialect) GetSqlType(langType string, bIsNullable *bool, nameOfColumn string) string {

	langType = pd.Base1.CleanNameGoStruct(langType)
	if strings.HasPrefix(langType, "[]") {
		return pd.Base1.PrintError(fmt.Sprintf("type %s not supported!", langType))
	}

	switch langType {
	case "NullString":
		*bIsNullable = true
		return pd.Base1.IsLongType(nameOfColumn, "TEXT", "VARCHAR")
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
		return pd.IsLongType(nameOfColumn, "TEXT", "VARCHAR")
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
		return pd.PrintError(fmt.Sprintf("not supported type %s", langType))
	}
}

func (pd *PostgressDialect) GetProperty(propName, tableName, colName, colValue string) string {
	requestText := fmt.Sprintf(`
	SELECT %s FROM %s
	WHERE %s = '%s'`, pd.TokenizIdentif(colValue), pd.TokenizTableName(tableName), pd.TokenizIdentif(colName), propName)

	row := pd.Connection.QueryRow(requestText)
	var obj interface{}
	if err := row.Scan(&obj); err != nil {
		return ""
	}
	return fmt.Sprintf("%v", obj)
}

func (pd *PostgressDialect) UpdateProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
	UPDATE %s
	SET %s = '%s'
	WHERE %s = '%s'`, pd.TokenizIdentif(tableName), pd.TokenizIdentif(colValue), value, pd.TokenizIdentif(colName), propName)

	pd.Connection.Exec(requestText)
}

func (pd *PostgressDialect) InsertProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
	INSERT INTO %s (%s, %s, %s) VALUES (default, '%s', '%s')`, pd.TokenizTableName(tableName), pd.TokenizIdentif("ID"), pd.TokenizIdentif(colName), pd.TokenizIdentif(colValue), propName, value)

	pd.Connection.Exec(requestText)
}

func (pd *PostgressDialect) ExecScript(scriptTxt string) {
	parts := strings.Split(scriptTxt, ";")
	for i, part := range parts {
		script1 := strings.TrimSpace(part)
		if len(script1) > 0 {
			_, err := pd.Connection.Exec(script1)
			if err != nil {
				msg := err.Error()
				fmt.Printf("Error exec script part %d: %s\n", i, msg)
				panic(err)
			}
		}
	}
}


func (pd *PostgressDialect) StartConnection(arg dialect.IGenericDialectArg) bool {

	connectionString := arg.GetGenericDialectArg().Connection_String
    var connStr = utils.Utils_parseConnectionString( connectionString );
    
	var dataSource = "";
    if( connStr.Port == ""){

        dataSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			connStr.Host, connStr.User, connStr.Password, connStr.DbName)
    }else{
        dataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			connStr.Host, connStr.Port, connStr.User, connStr.Password, connStr.DbName)
    }

	conn, err := sql.Open("postgres", dataSource)
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

	pd.Connection = conn

	if true {
		pd.Base1.FnProcessData(pd, arg)
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

	pd.Base1.FnProcessData(pd, arg)

	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		conn.Close()
		return false
	}

	return true
}

func (pd *PostgressDialect) IsYes(s string) bool {
	return s != "" && s == "YES"
}


func (pd *PostgressDialect) IsLongType(sqlColumnName string, type1 string, type2 string) string {

    return pd.Base1.IsLongType(sqlColumnName, type1, type2);
}

func (pd *PostgressDialect) PrintError(err string) string {

    return pd.Base1.PrintError(err)
}
func (pd *PostgressDialect) SetCurrentFile(file string)  {
    
    pd.Base1.SetCurrentFile(file)
}

func (pd *PostgressDialect) CleanNameGoStruct(name string) string {

    return pd.Base1.CleanNameGoStruct(name )
}