package gomigration_pstgress

import (
	"fmt"
	"strings"

	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
)


func (pd *PostgressDialect) GetGoLangTypeIntFk(column *dialect.DbColumn, packageImports *map[string]string) string {
	nameSqlType := column.SqlType

	if column.IsNullable {
		(*packageImports)["sql"] = "sql \"database/sql\""
		switch nameSqlType {
		case "smallint":
			return "sql.NullInt16"
		case "smallserial":
			return "sql.NullInt16"
		case "integer":
			return "sql.NullInt32"
		case "serial":
			return "sql.NullInt32"
		case "bigint":
			return "sql.NullInt64"
		case "bigserial":
			return "sql.NullInt64"
		default:
			fmt.Printf("%s: *** not implemented ***\n", nameSqlType)
			return ""
		}
	} else {
		switch nameSqlType {
		case "smallint":
			return "int16"
		case "smallserial":
			return "int16"
		case "integer":
			return "int32"
		case "serial":
			return "int32"
		case "bigint":
			return "int64"
		case "bigserial":
			return "int64"
		default:
			fmt.Printf("%s: *** not implemented ***\n", nameSqlType)
			return ""
		}
	}
}

func (pd *PostgressDialect) GetGoLangType(column *dialect.DbColumn, importPackage *map[string]string) string {
	nameSqlType := column.SqlType

	if strings.Contains(nameSqlType, "[]") {
		if column.IsNullable {
			return "[]uint8"
		} else {
			return "[]uint8"
		}
	}

	if column.ForeignKey != nil {
		return "*" + column.ForeignKey.LangTableNameModel
	}

	switch nameSqlType {
	case "smallint":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt16"
		} else {
			return "int16"
		}
	case "smallserial":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt16"
		} else {
			return "int16"
		}
	case "integer":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt32"
		} else {
			return "int32"
		}
	case "serial":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt32"
		} else {
			return "int32"
		}
	case "bigint", "int4":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt64"
		} else {
			return "int64"
		}
	case "bigserial":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullInt64"
		} else {
			return "int64"
		}
	case "uuid":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullString"
		} else {
			return "string"
		}
	case "timestamp", "timestamp without time zone", "date":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullTime"
		} else {
			(*importPackage)["time"] = "time \"time\""
			return "time.Time"
		}
	case "bit", "boolean", "bool":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullBool"
		} else {
			return "bool"
		}
	case "decimal", "numeric":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullFloat64"
		} else {
			return "float64"
		}
	case "real":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			fmt.Printf("sql.NullFloat32 is not exist in golang database/sql. Change the type of field '%s' to double precision\n", column.SqlName)
			return "sql.NullFloat32"
		} else {
			return "float32"
		}
	case "double precision":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
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
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullString"
		} else {
			return "string"
		}
	case "character varying", "varchar":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullString"
		} else {
			return "string"
		}
	case "text":
		if column.IsNullable {
			(*importPackage)["sql"] = "sql \"database/sql\""
			return "sql.NullString"
		} else {
			return "string"
		}
	default:
		fmt.Printf("Sql Type '%s': *** not implemented ***. Use text or varchar(1024)\n", nameSqlType)
		return ""
	}
}
