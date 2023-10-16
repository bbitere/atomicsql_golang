package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)



type User struct /*atomicsql-table:"user"*/ {

	orm.Generic_MODEL

	ID                  int32                         `json:"ID,omitempty"`
	UserName            string                        `json:"userName"`
	UUID                string                        `json:"UUID"`
	UserPsw             string                        `json:"user_psw"`
	UserRoleID          *UserRole                     `json:"-"`
	UserRole_ID         sql.NullInt32                 `json:"userRole_ID"`
	/*Misc1               []uint8                       `json:"misc1"`*/
	Time1               sql.NullTime                  `json:"time1"`
	Money               float64                       `json:"money"`


}