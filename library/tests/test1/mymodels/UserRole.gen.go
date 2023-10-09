/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /*import (
	    sql "database/sql"
        )*/
	/*
    type UserRole struct /*atomicsql-table:"user_role"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            IsActive            byte                          `json:"isActive"`
            RoleName            string                        `json:"roleName"`
            RoleStatusID        *StatusRole                   `json:"-"`
            Role_status_ID      sql.NullInt32                 `json:"role_status_ID"`
            MyActive            byte                          `json:"my_active"`
    }
	*/
    func (model  UserRole) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  UserRole) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_UserRole struct {
	        
			orm.Generic_Def
	        ID                  string
            IsActive            string
            RoleName            string
            RoleStatusID        T_StatusRole
            Role_status_ID      string
            MyActive            string
    }
	
	func (_this *T_UserRole) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
            FnNewInst:	func()any{ return new (UserRole) },
        }
    }
