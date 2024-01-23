package gomigration_pstgress

import (
	"database/sql"
	"fmt"
)

func (pd *PostgressDialect) readTables(SqlLang string) (map[string]*DbTable, error) {

	tables := make(map[string]*DbTable)
	cnn := pd.connection
	packageImports := make(map[string]string)

	requestText := `
		SELECT
			t.table_name AS table1,
			t.table_schema AS schema1,
			c.column_name AS column1,
			format('%s', udt_name::regtype) AS colType2,
			c.data_type AS colType1,
			c.is_nullable AS isNullable1,
			c.ordinal_position AS ordinal_position1,
			(
				SELECT 'YES'
				FROM information_schema.table_constraints tc1
				JOIN information_schema.constraint_column_usage AS ccu USING (constraint_schema, constraint_name)
				JOIN information_schema.columns AS c1 ON c1.table_schema = tc1.constraint_schema
					AND tc1.table_name = c1.table_name AND ccu.column_name = c1.column_name
				WHERE constraint_type = 'PRIMARY KEY'
					AND tc1.table_name = t.table_name
					AND c1.column_name = c.column_name
				LIMIT 1
			) AS is_identity1
		FROM
			information_schema.columns c
		INNER JOIN information_schema.tables t ON
			t.table_name = c.table_name
		WHERE
			t.table_schema = 'public'
			AND t.table_type = 'BASE TABLE'
			AND c.table_schema = 'public'
		GROUP BY t.table_name, t.table_schema,
			c.column_name, udt_name::regtype, c.data_type,
			c.is_nullable, c.is_identity,
			c.ordinal_position
		ORDER BY c.ordinal_position;
	`

	rows, err := cnn.Query(requestText)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var table *DbTable
	for rows.Next() {
		var tableName, tableSchema string
		var columnName, colType, colType2, colIsNullable, colIsIdentity string
		if err := rows.Scan(&tableName, &tableSchema, &columnName, &colType2, &colType, &colIsNullable, &colIsIdentity); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		table = tables[tableName]
		if table == nil {
			table = &DbTable{
				LangTableNameModel: ConvertToIdentGoLang(tableName),
				Schema:             tableSchema,
				SqlTableNameModel:  tableName,
				PrimaryColumn:      nil,
				Columns:            make([]*DbColumn, 0),
			}
			tables[tableName] = table
		}

		column := &DbColumn{
			LangName:      ConvertToIdentGoLang(columnName, true),
			LangName2:     ConvertToIdentGoLang(columnName, false),
			SqlName:       columnName,
			SqlType:       colType,
			LangType:      pd.getGoLangType(column, &packageImports),
			BIsIdentity:   pd.isYes(colIsIdentity),
			BIsNullable:   pd.isYes(colIsNullable),
		}

		table.Columns = append(table.Columns, column)

		if column.BIsIdentity {
			if table.PrimaryColumn != nil {
				return nil, fmt.Errorf("%s has already a primary key", table.SqlTableNameModel)
			}
			table.PrimaryColumn = column
		}
	}

	if table == nil {
		return nil, fmt.Errorf("table not loaded")
	}
	if table.PrimaryColumn == nil {
		return nil, fmt.Errorf("table %s doesn't have a primary key", table.LangTableNameModel)
	}

	return tables, nil
}

func (pd *PostgressDialect) readConstraintors(tables map[string]*DbTable) error {
	cnn := pd.connection
	requestText := `
		SELECT * FROM (
			SELECT
				ccu.table_schema AS table_schema,
				kcu.table_name AS table_name,
				CASE WHEN (pgc.contype = 'f') THEN kcu.column_name ELSE ccu.column_name END AS column_name,
				CASE WHEN (pgc.contype = 'f') THEN ccu.table_name ELSE (NULL) END AS reference_table,
				CASE WHEN (pgc.contype = 'f') THEN ccu.column_name ELSE (NULL) END AS reference_col,
				CASE WHEN (pgc.contype = 'p') THEN 'yes' ELSE 'no' END AS auto_inc,
				CASE WHEN (pgc.contype = 'p') THEN 'no' ELSE 'yes' END AS is_nullable,
				'integer' AS data_type,
				'0' AS numeric_scale,
				'32' AS numeric_precision
			FROM
				pg_constraint AS pgc
				J
                JOIN pg_namespace nsp ON nsp.oid = pgc.connamespace
				JOIN pg_class cls ON pgc.conrelid = cls.oid
				JOIN information_schema.key_column_usage kcu ON kcu.constraint_name = pgc.conname
				LEFT JOIN information_schema.constraint_column_usage ccu ON pgc.conname = ccu.constraint_name
					AND nsp.nspname = ccu.constraint_schema
					AND ccu.constraint_schema = 'public'
			UNION
			SELECT
				NULL AS table_schema,
				table_name,
				column_name,
				NULL AS reference_table,
				NULL AS reference_col,
				'no' AS auto_inc,
				is_nullable,
				data_type,
				numeric_scale,
				numeric_precision
			FROM information_schema.columns cols
			WHERE
				table_schema = 'public'
				AND CONCAT(table_name, column_name) NOT IN (
					SELECT CONCAT(kcu.table_name, kcu.column_name)
					FROM
						pg_constraint AS pgc
						JOIN pg_namespace nsp ON nsp.oid = pgc.connamespace
						JOIN pg_class cls ON pgc.conrelid = cls.oid
						JOIN information_schema.key_column_usage kcu ON kcu.constraint_name = pgc.conname
						LEFT JOIN information_schema.constraint_column_usage ccu ON pgc.conname = ccu.constraint_name
							AND nsp.nspname = ccu.constraint_schema
				)
		) AS foo
		WHERE foo.table_schema IS NOT NULL
		ORDER BY table_name ASC, column_name
	`

	rows, err := cnn.Query(requestText)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var schemaName, tableName, columnName, colRefTable, colRefColumn string
		if err := rows.Scan(&schemaName, &tableName, &columnName, &colRefTable, &colRefColumn); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		if colRefTable != "" && colRefColumn != "" {
			if table, ok := tables[tableName]; ok {
				if column := findColumn(table, columnName); column != nil {
					if refTable, ok := tables[colRefTable]; ok {
						if primary := refTable.PrimaryColumn; primary != nil && primary.LangName == colRefColumn {
							column.ForeignKey = refTable
						} else {
							return fmt.Errorf("Foreign key from table %s to %s is not pointed to primary key", tableName, colRefTable)
						}
					} else {
						return fmt.Errorf("not found table %s", colRefTable)
					}
				}
			} else {
				return fmt.Errorf("not found table %s", tableName)
			}
		}
	}

	return nil
}

func (pd *PostgressDialect) getString(reader *sql.Rows, index int) string {
	var result sql.NullString
	if err := reader.Scan(&result); err != nil {
		return ""
	}

	if result.Valid {
		return result.String
	}
	return ""
}

func (pd *PostgressDialect) isYes(s string) bool {
	return s == "YES"
}

