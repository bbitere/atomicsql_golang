package test1_atomicsql

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"
	//orm "github.com/bbitere/atomicsql_golang/atomicsql"
)


type UserRole struct /*atomicsql-table:"user_role"*/ {

	//Generic_MODEL

	ID                  int32                         `json:"ID,omitempty"`
	IsActive            bool                          `json:"isActive"`
	RoleName            string                        `json:"roleName"`
	RoleStatusID        *StatusRole                   `json:"-"`
	Role_status_ID      sql.NullInt32                 `json:"role_status_ID"`
	
}