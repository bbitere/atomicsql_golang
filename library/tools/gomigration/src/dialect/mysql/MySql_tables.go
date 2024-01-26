package gomigration_mysql

import (
	//"database/sql"
	"fmt"
	"strings"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"
)

func (md *MySqlDialect) ReadTables(SqlLang string) (map[string]*dialect.DbTable, error) {

	tables := make(map[string]*dialect.DbTable)
	cnn := md.connection
	packageImports := make(map[string]string)

	//myDatabaseName := md.connection.Database
	requestText := `
		SELECT 
			table_name,        /*0*/
			table_schema,      /*1*/
			column_name,       /*2*/ 
			data_type,         /*3*/
			table_catalog,     /*4*/
			is_nullable,       /*5*/
			extra,             /*6*/
			column_default,    /*7*/
			ordinal_position
		FROM information_schema.columns
		WHERE table_schema = DATABASE()
		ORDER BY table_schema, table_name, ordinal_position
	`

	rows, err := cnn.Query(requestText)
	if err != nil {
		//return fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		                 var tableName, tableSchema,  columnName,  colType,  colType2,  colIsNullable,  colIsIdentity string
		if err := rows.Scan(&tableName, &tableSchema, &columnName, &colType, &colType2, &colIsNullable, &colIsIdentity); err != nil {
			//return fmt.Errorf("error scanning row: %v", err)
		}

        //tableName = getString(reader, 0)
        //tableSchema := getString(reader, 1)

        var table = tables[tableName]
        if table == nil {
            table = &dialect.DbTable{}
        }

        table.LangTableNameModel = utils.Utils_ConvertToIdentGoLang(tableName, true)
        table.Schema = tableSchema
        table.SqlTableNameModel = tableName
        tables[table.SqlTableNameModel] = table

        //columnName := getString(reader, 2)
        //colType := getString(reader, 3)
        //colType2 := getString(reader, 4)
        //colIsNullable := getString(reader, 5)
        //colIsIdentity := getString(reader, 6)
        //colDefalut := getString(reader, 7)

        column := &dialect.DbColumn{
            LangName:   utils.Utils_ConvertToIdentGoLang(columnName, true),
            LangName2:  utils.Utils_ConvertToIdentGoLang(columnName, false),
            SqlName:    columnName,
            SqlType:    colType,
            LangType:   "",//md.GetGoLangType(column, &packageImports),
            IsIdentity: colIsIdentity != "" && strings.Contains( colIsIdentity, "auto_increment"),
            IsNullable: isYes(colIsNullable),
        }
        column.LangType =  md.GetGoLangType(column, &packageImports);

        table.Columns = append(table.Columns, column)

        if column.IsIdentity {
            if table.PrimaryColumn != nil {
                fmt.Println(table, "has already a primary key")
                return nil, fmt.Errorf("table %s has already a primary key", tableName)
            }
            table.PrimaryColumn = column
        }
		

		if table == nil {
			return nil, fmt.Errorf("table not loaded")
		}

		if table.PrimaryColumn == nil {
			return nil, fmt.Errorf("table %s doesn't have a primary key", table.LangTableNameModel)
		}

		return tables, nil
	} 

	return nil, nil
}

func (md *MySqlDialect) ReadConstraintors(tables map[string]*dialect.DbTable) error {
	cnn := md.connection
	requestText := `
		SELECT DISTINCT(TC.table_name) AS table_name,   /*0*/
			TC.table_schema AS table_schema,            /*1*/
			KC.column_name as column_name,              /*2*/  
			KC.REFERENCED_TABLE_NAME,                   /*3*/
			KC.REFERENCED_COLUMN_NAME,                  /*4*/
			KC.CONSTRAINT_NAME,
			TC. CONSTRAINT_TYPE AS ConstraintType
			FROM 
			information_schema.table_constraints TC,  
			information_schema.key_column_usage KC  
			WHERE
			TC.CONSTRAINT_TYPE IN ('FOREIGN KEY', 'PRIMARY KEY') 
			and KC.table_name = TC.table_name and KC.TABLE_SCHEMA = TC.TABLE_SCHEMA
			and KC.CONSTRAINT_NAME = TC.CONSTRAINT_NAME
			and KC.table_schema = DATABASE()
			ORDER BY table_name ASC, column_name
	`

	rows, err := cnn.Query(requestText)
	if err != nil {
		//return fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		                 var tableName, schemaName,  columnName,   colRefTable, colRefColumn string;
		if err := rows.Scan(&tableName, &schemaName, &columnName, &colRefTable, &colRefColumn); err != nil {
			//return fmt.Errorf("error scanning row: %v", err)
		}

        //var tableName = getString(reader, 0)
        //var schemaName = getString(reader, 1)

        //var columnName = getString(reader, 2)
        //var colRefTable = getString(reader, 3)
        //var colRefColumn = getString(reader, 4)

        if colRefTable != "" && colRefColumn != "" {
            if _, exists := tables[tableName]; exists {
                table := tables[tableName]
                if column := findColumnByName(table, columnName); column != nil {
                    if refTable, exists := tables[colRefTable]; exists {
                        if primary := refTable.PrimaryColumn; primary != nil && primary.LangName == colRefColumn {
                            column.ForeignKey = refTable
                        } else {
                            md.Base1.PrintError(fmt.Sprintf("Foreign key from table %s to %s is not pointed to primary key", tableName, colRefTable))
                            return nil
                        }
                    } else {
                        md.Base1.PrintError(fmt.Sprintf("Not found table %s", colRefTable))
                        return nil
                    }
                }
            } else {
                md.Base1.PrintError(fmt.Sprintf("Not found table %s", tableName))
                return nil
            }
        }        
        return nil
    } 

    return nil
}



func isYes(s string) bool {
    return s != "" && s == "YES"
}

func findColumnByName(table *dialect.DbTable, columnName string) *dialect.DbColumn {

    for _, column := range table.Columns {
        if column.SqlName == columnName {
            return column
        }
    }
    return nil
}

