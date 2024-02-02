package gomigration_mysql

import (
	"database/sql"
	"fmt"
	"strings"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDialect struct {
	connection *sql.DB
    Base1 dialect.TGenericDialect;
}

// GetSql returns the SQL dialect used by MySQL
func (m *MySqlDialect) GetSql() string {
	return "mysql"
}

// SqlSeparator returns the SQL separator used by MySQL
func (m *MySqlDialect) SqlSeparator() string {
	return ";"
}

// TokenizeIdentif tokenizes an identifier for MySQL
func (m *MySqlDialect) TokenizeIdentif(identif string) string {
	return fmt.Sprintf("`%s`", identif)
}

// TokenizeTable tokenizes a table name for MySQL
func (m *MySqlDialect) TokenizeTable(table *dialect.DbTable) string {
	return m.TokenizeIdentif(table.SqlTableNameModel)
}

// DropTable generates a DROP TABLE statement for MySQL
func (m *MySqlDialect) DropTable(table *dialect.DbTable) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", m.TokenizeIdentif(table.SqlTableNameModel))
}

// AddTable generates a CREATE TABLE statement for MySQL
func (m *MySqlDialect) AddTable(table *dialect.DbTable) string {
	var NL = "\n"
	var constraintorList []string
	var columnsArr []string
	sqlTableName := table.SqlTableNameModel

	for _, col := range table.Columns {
		columnsArr = append(columnsArr, m._AddColumn(col, &constraintorList))
	}

	columnsDefs := strings.Join(columnsArr, ","+NL)
	colIDName := table.PrimaryColumn.SqlName

	constraintors := strings.Join(constraintorList, ","+NL)
	if constraintors != "" {
		constraintors = fmt.Sprintf(",%s%s", NL, constraintors)
	}

	tableAdd := fmt.Sprintf(`
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS %s
            (
                %s

                ,PRIMARY KEY (%s)
                %s
            ) ENGINE=InnoDB`, m.TokenizeIdentif(sqlTableName), columnsDefs, m.TokenizeIdentif(colIDName), constraintors)

	return tableAdd
}

// DropColumn generates a DROP COLUMN statement for MySQL
func (m *MySqlDialect) DropColumn(table *dialect.DbTable, column *dialect.DbColumn) string {
	return fmt.Sprintf(`
            ALTER TABLE %s
            DROP COLUMN %s`, m.TokenizeIdentif(table.SqlTableNameModel), m.TokenizeIdentif(column.SqlName))
}

// AddColumn generates an ADD COLUMN statement for MySQL
func (m *MySqlDialect) AddColumn(table *dialect.DbTable, column *dialect.DbColumn) string {
	listConstraintors := make([]string, 0)
	colData := m._AddColumn(column, &listConstraintors)
	return fmt.Sprintf(`
            ALTER TABLE %s
            ADD COLUMN %s`, m.TokenizeIdentif(table.SqlTableNameModel), colData)
}

// _AddColumn generates the column definition for MySQL
func (m *MySqlDialect) _AddColumn(column *dialect.DbColumn, listConstraintors *[]string) string {
	colName := column.SqlName

	if column.IsIdentity {
		if column.LangType == "int64" {
			return fmt.Sprintf("%s BIGINT NOT NULL AUTO_INCREMENT", m.TokenizeIdentif(colName))
		} else {
			return fmt.Sprintf("%s INT NOT NULL AUTO_INCREMENT", m.TokenizeIdentif(colName))
		}
	} else {
		sqlType := column.SqlType
		if column.ForeignKey != nil {
			targetTableSqlName := m.TokenizeTable(column.ForeignKey)
			targetTableID := m.TokenizeIdentif(column.ForeignKey.PrimaryColumn.SqlName)

			fkName := m.TokenizeIdentif(column.SqlName)
			fkFKName := m.TokenizeIdentif(fmt.Sprintf("fk_%s", column.SqlName))

			*listConstraintors = append(*listConstraintors, fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)", fkFKName, fkName, targetTableSqlName, targetTableID))

			return fmt.Sprintf("%s INT NULL", fkName)
		} else {
			if column.IsNullable {
				return fmt.Sprintf("%s %s NULL", m.TokenizeIdentif(colName), sqlType)
			} else {
				return fmt.Sprintf("%s %s NOT NULL", m.TokenizeIdentif(colName), sqlType)
			}
		}
	}
}

// UpdateTable generates an UPDATE TABLE statement for MySQL
func (m *MySqlDialect) UpdateTable(table *dialect.DbTable, tablePrev *dialect.DbTable) string {
	if table.SqlTableNameModel != tablePrev.SqlTableNameModel {
		return fmt.Sprintf(`
                RENAME TABLE %s
                TO %s`, m.TokenizeIdentif(tablePrev.SqlTableNameModel), m.TokenizeIdentif(table.SqlTableNameModel))
	}
	return ""
}

// UpdateColumn generates an UPDATE COLUMN statement for MySQL
func (m *MySqlDialect) UpdateColumn(table *dialect.DbTable, column *dialect.DbColumn, columnPrev *dialect.DbColumn) string {
	if columnPrev.LangName != column.LangName {
		return m.Base1.PrintError("internal error")
	}

	if column.SqlType != columnPrev.SqlType {
		return fmt.Sprintf(`
                ALTER TABLE %s 
                MODIFY %s %s`, m.TokenizeIdentif(table.SqlTableNameModel), m.TokenizeIdentif(column.SqlName), column.SqlType)
	}

	if column.SqlName != columnPrev.SqlName {
		return fmt.Sprintf(`
                ALTER TABLE %s
                CHANGE COLUMN %s %s %s`, m.TokenizeIdentif(table.SqlTableNameModel), m.TokenizeIdentif(columnPrev.SqlName), m.TokenizeIdentif(column.SqlName), column.SqlType)
	}

	return ""
}

// AddFKConstraint generates an ADD FOREIGN KEY statement for MySQL
func (m *MySqlDialect) AddFKConstrictor(table *dialect.DbTable, column *dialect.DbColumn) string {
	tableName := m.TokenizeTable(column.ForeignKey)
	tableID := m.TokenizeIdentif(column.ForeignKey.PrimaryColumn.SqlName)

	myTableName := m.TokenizeTable(table)
	fkName := m.TokenizeIdentif(column.SqlName)
	fkFKName := m.TokenizeIdentif(fmt.Sprintf("fk_%s", column.SqlName))

	return fmt.Sprintf(`
                ALTER TABLE %s
                ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s);`, myTableName, fkFKName, fkName, tableName, tableID)
}

// DropFKConstraint generates a DROP FOREIGN KEY statement for MySQL
func (m *MySqlDialect) DropFKConstrictor(table *dialect.DbTable, column *dialect.DbColumn) string {
	myTableName := m.TokenizeTable(table)
	fkFKName := m.TokenizeIdentif(fmt.Sprintf("fk_%s", column.SqlName))

	return fmt.Sprintf(`
                ALTER TABLE %s
                DROP FOREIGN KEY %s;`, myTableName, fkFKName)
}

// GetSqlType returns the equivalent MySQL data type for a given Go type
func (m *MySqlDialect) GetSqlType(langType string, bIsNullable *bool, nameOfColumn string) string {
	langType = m.Base1.CleanNameGoStruct(langType)
	if strings.HasPrefix(langType, "[]") {
		return m.Base1.PrintError(fmt.Sprintf("type %s not supported!", langType))
	}

	switch langType {
	case "NullString":
		*bIsNullable = true
		return m.IsLongType(nameOfColumn, "TEXT", "VARCHAR(255)" )
	case "NullBool":
		*bIsNullable = true
		return "BOOLEAN"
	case "NullByte", "NullInt16", "NullInt", "NullInt32":
		*bIsNullable = true
		return "SMALLINT"
	case "NullInt64":
		*bIsNullable = true
		return "BIGINT"
	case "NullFloat64":
		*bIsNullable = true
		return "DOUBLE PRECISION"
	case "time.NullTime", "NullTime":
		*bIsNullable = true
		return "TIMESTAMP"
	case "string":
		return m.IsLongType(nameOfColumn, "TEXT", "VARCHAR(255)" )
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
		return "FLOAT"
	case "float64":
		return "DOUBLE PRECISION"
	default:
		return m.Base1.PrintError(fmt.Sprintf("not supported type %s", langType))
	}
}


// getProperty retrieves a property from the database for a specific table, column, and value
func (m *MySqlDialect) GetProperty(propName, tableName, colName, colValue string) string {
	requestText := fmt.Sprintf(`
                SELECT %s FROM %s
                WHERE %s = '%s'`, m.TokenizeIdentif(colValue), m.TokenizeIdentif(tableName), m.TokenizeIdentif(colName), propName)

	row := m.connection.QueryRow(requestText)
	var value string
	if err := row.Scan(&value); err != nil {
		return ""
	}
	return value
}

// UpdateProperty updates a property in the database for a specific table, column, and value
func (m *MySqlDialect) UpdateProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
                UPDATE %s
                SET   %s  = '%s'
                WHERE %s   = '%s'`, m.TokenizeIdentif(tableName), m.TokenizeIdentif(colValue), value, m.TokenizeIdentif(colName), propName)

	_, err := m.connection.Exec(requestText)
	if err != nil {
		fmt.Println("Error updating property:", err)
	}
}

// InsertProperty inserts a property into the database for a specific table, column, and value
func (m *MySqlDialect) InsertProperty(propName, value, tableName, colName, colValue string) {
	requestText := fmt.Sprintf(`
                INSERT INTO %s ( %s, %s, %s) 
                VALUES (  default, '%s', '%s' )`, m.TokenizeIdentif(tableName), m.TokenizeIdentif("ID"), m.TokenizeIdentif(colName), m.TokenizeIdentif(colValue), propName, value)

	_, err := m.connection.Exec(requestText)
	if err != nil {
		fmt.Println("Error inserting property:", err)
	}
}

// ExecScript executes a SQL script in MySQL
func (m *MySqlDialect) ExecScript(scriptTxt string) {
	parts := strings.Split(scriptTxt, ";")
	for i, part := range parts {
		script1 := strings.TrimSpace(part)
		if len(script1) > 0 {
			_, err := m.connection.Exec(script1)
			if err != nil {
				msg := err.Error()
				fmt.Printf("Error exec script part %d: %s\n", i, msg)
				panic(err)
			}
		}
	}
}

// StartConnection initiates a connection to the MySQL database
func (m *MySqlDialect) StartConnection(arg dialect.IGenericDialectArg) bool {
	var err error

    var connString = arg.GetGenericDialectArg().Connection_String;

    var connStr = utils.Utils_parseConnectionString(connString);
	var	dataSource = fmt.Sprintf(
			"Data Source=%s;"+
				"Initial Catalog=%s;"+
				"User id=%s;"+
				"password=%s;",
			connStr.Host,
			connStr.DbName,
			connStr.User,
			connStr.Password)

	m.connection, err = sql.Open("mysql", dataSource )
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return false
	}
	defer m.connection.Close()

	err = m.connection.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return false
	}

	if true {
		m.Base1.FnProcessData( m, arg)
		return true
	}

	return false
}



func (pd *MySqlDialect) IsLongType(sqlColumnName string, type1 string, type2 string) string {

    return pd.Base1.IsLongType(sqlColumnName, type1, type2);
}

func (pd *MySqlDialect) PrintError(err string) string {

    return pd.Base1.PrintError(err)
}
func (pd *MySqlDialect) SetCurrentFile(file string)  {
    
    pd.Base1.SetCurrentFile(file)
}

func (pd *MySqlDialect) CleanNameGoStruct(name string) string {

    return pd.Base1.CleanNameGoStruct(name )
}
func (pd *MySqlDialect) GetGenericDialect() *dialect.TGenericDialect{

    return &pd.Base1;
}