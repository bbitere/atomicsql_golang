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
            IsActive            bool                          `json:"isActive"`
            RoleName            string                        `json:"roleName"`
            RoleStatusID        *Statusrole                   `json:"-"`
            Role_status_ID      sql.NullInt32                 `json:"statusrole_ID"`
            MyActive            bool                          `json:"my_active"`
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
            RoleStatusID        T_Statusrole
            Role_status_ID      string
            MyActive            string
    }
	
	func (_this *T_UserRole) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"user_role",
            FnNewInst:	func(bFull bool)any{ 
				var model = new (UserRole);
				if( bFull ){
					
                        var defRoleStatusID = T_Statusrole{}
				        model.RoleStatusID = (defRoleStatusID.Def().FnNewInst(bFull)).(*Statusrole)
				}
				return model;
			},
        }
    }
