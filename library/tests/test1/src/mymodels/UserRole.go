package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

type UserRole struct /*atomicsql-table:"user_role"*/ {
	orm.Generic_MODEL

	ID             int32         `json:"ID,omitempty"`
	IsActive       bool          `json:"isActive"`
	RoleName       string        `json:"roleName"`
	RoleStatusID   *Statusrole   `json:"-"`
	Role_status_ID sql.NullInt32 `json:"role_status_ID"`
	MyActive       bool          `json:"my_active"`
}

type Department struct /*atomicsql-table:"departm"*/ {
	orm.Generic_MODEL

	ID       int32  `json:"ID,omitempty"`
	IsActive bool   `json:"isActive"`
	NameDep  string `json:"name_dep"`
}

type Employee1 struct /*atomicsql-table:"employee1"*/ {
	orm.Generic_MODEL

	ID       int32 `json:"ID,omitempty"`
	IsActive bool  `json:"isActive"`

	UserID  *User `json:"-"`
	User_ID int   `json:"user_ID"`

	DepartmentID  *Department   `json:"-"`
	Department_ID sql.NullInt32 `json:"departm_ID"`

	Department2ID  *Department   `json:"-"`
	Department2_ID sql.NullInt32 `json:"departm2_ID"`
}

type Employee2 struct /*atomicsql-table:"employee2"*/ {
	orm.Generic_MODEL

	ID int32 `json:"ID,omitempty"`
}

type ProjectStatus struct /*atomicsql-table:"projstatus"*/ {
	orm.Generic_MODEL

	ID   int32  `json:"ID,omitempty"`
	Name string `json:"name"`
}

type Project struct /*atomicsql-table:"project"*/ {
	orm.Generic_MODEL

	ID   int32  `json:"ID,omitempty"`
	Name string `json:"name"`

	//ProjectStatus2ID		*ProjectStatus				`json:"-"`
	//ProjectStatus2_ID		int							`json:"projStatus2_ID"`
}
