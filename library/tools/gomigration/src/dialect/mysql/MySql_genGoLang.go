package gomigration_mysql

import (
	//"database/sql"
	"fmt"
	"strings"

	// "time"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
)

// GetGoLangTypeIntFK returns the equivalent GoLang type for an integer foreign key in MySQL
func (m *MySqlDialect) GetGoLangTypeIntFk(column *dialect.DbColumn, packageImports *map[string]string) string {
	nameSqlType := column.SqlType
	//var nameArr string

	if column.IsNullable {
		(*packageImports)["sql"] = "database/sql"

		switch nameSqlType {
		case "smallint", "smallserial":
			return "sql.NullInt16"
		case "integer", "serial":
			return "sql.NullInt32"
		case "bigint", "bigserial":
			return "sql.NullInt64"
		default:
			fmt.Printf("%s: *** not implemented ***\n", nameSqlType)
			return ""
		}
	}

	switch nameSqlType {
	case "smallint", "smallserial":
		return "int16"
	case "integer", "serial":
		return "int32"
	case "bigint", "bigserial":
		return "int64"
	default:
		fmt.Printf("%s: *** not implemented ***\n", nameSqlType)
		return ""
	}
}

// GetGoLangType returns the equivalent GoLang type for a column in MySQL
func (m *MySqlDialect) GetGoLangType(column *dialect.DbColumn, _importPackage *map[string]string) string {
	nameSqlType := column.SqlType
	//var nameArr string
    importPackage := *_importPackage

	if strings.Contains(nameSqlType, "[]") {
		if column.IsNullable {
			return "[]uint8"
		} else {
			return "[]uint8"
		}
	}

	if column.ForeignKey != nil {
		return fmt.Sprintf("*%s", column.ForeignKey.LangTableNameModel)
	}

	switch nameSqlType {
	case "tinyint", "smallint":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt16"
		} else {
			return "int16"
		}
	case "smallserial":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt16"
		} else {
			return "int16"
		}
	case "int", "integer":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt32"
		} else {
			return "int32"
		}
	case "serial":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt32"
		} else {
			return "int32"
		}
	case "bigint", "int4":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt64"
		} else {
			return "int64"
		}
	case "bigserial":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullInt64"
		} else {
			return "int64"
		}
	case "uuid":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullString"
		} else {
			return "string"
		}
	case "timestamp", "timestamp without time zone", "date":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullTime"
		} else {
			importPackage["time"] = "time"
			return "time.Time"
		}
	case "bit", "boolean", "bool":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullBool"
		} else {
			return "bool"
		}
	case "decimal", "double", "numeric":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullFloat64"
		} else {
			return "float64"
		}
	case "real":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			fmt.Println("sql.NullFloat32 is not exist in golang database/sql. Change the type of field '", column.SqlName, "' to double precision")
			return "sql.NullFloat32"
		} else {
			return "float32"
		}
	case "double precision":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullFloat64"
		} else {
			return "float64"
		}
	case "json":
		if column.IsNullable {
			return "string"
		} else {
			return "string"
		}
	case "array":
		fmt.Printf("%s: *** not implemented ***\n", nameSqlType)
		return "[]string"
	case "character", "char":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullString"
		} else {
			return "string"
		}
	case "character varying", "varchar":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullString"
		} else {
			return "string"
		}
	case "text":
		if column.IsNullable {
			importPackage["sql"] = "database/sql"
			return "sql.NullString"
		} else {
			return "string"
		}
	default:
		//debug.Assert(false, "Wait Debugger")

        fmt.Printf("Sql Type '%s': *** not implemented ***. Use text or varchar(1024)\n", nameSqlType)
		return ""
	}
}
