package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang"

	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

type Statusrole struct /*atomicsql-table:"statusRole"*/ {
	orm.Generic_MODEL

	ID         int32          `json:"ID,omitempty"`
	StatusName sql.NullString `json:"status_name"`
}
