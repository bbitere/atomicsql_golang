package gomigration_general

import (
	//"database/sql"
	"fmt"
	//"strings"
	dialect "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect"
	postgres "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/dialect/postgres"
)

func GetDialectByName(SqlLang string) dialect.GenericDialect {
	
	if SqlLang == dialect.ELangSql.PostgresSql {
		return new (postgres.PostgressDialect);
	} else if SqlLang == dialect.ELangSql.MySql {
		return nil;//&MySqlDialect{}
	} else {
		fmt.Printf("Error: sql dialect '%s' not implemented\n", SqlLang)
		return nil
	}
}