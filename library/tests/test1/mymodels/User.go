package test1_atomicsql

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"
)



type User struct /*atomicsql-table:"user"*/ {

	//Generic_MODEL

	ID                  int32                         `json:"ID,omitempty"`
	UserName            string                        `json:"userName"`
	UUID                string                        `json:"UUID"`
	UserPsw             string                        `json:"user_psw"`
	UserRoleID          *UserRole                     `json:"-"`
	UserRole_ID         sql.NullInt32                 `json:"userRole_ID"`
	Misc1               []uint8                       `json:"misc1"`
	Time1               sql.NullTime                  `json:"time1"`
	Money               float64                       `json:"money"`

}