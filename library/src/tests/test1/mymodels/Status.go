package tests_test1_mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"

	orm "atomicsql"
)


type StatusRole struct /*atomicsql-table:"statusRole"*/ {

	orm.Generic_MODEL

	ID                  int32                         `json:"ID,omitempty"`
	StatusName          sql.NullString                `json:"status_name"`

}