package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang"

	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

type Statusrole struct /*atomicsql-table:"statusRole"*/ {
	orm.Generic_MODEL

	ID         int32          `bson:"-"`
	StatusName sql.NullString `bson:"status_name"`
}
