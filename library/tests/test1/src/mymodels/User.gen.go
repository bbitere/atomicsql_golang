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
    type User struct /*atomicsql-table:"user"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            UserName            string                        `json:"userName"`
            UUID                string                        `json:"UUID"`
            UserPsw             string                        `json:"user_psw"`
            UserRoleID          *UserRole                     `json:"-"`
            UserRole_ID                                       `json:"userRole_ID"`
            Time1               sql.NullTime                  `json:"time1"`
            Money               float64                       `json:"money"`
    }
	*/
    func (model  User) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  User) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_User struct {
	        
			orm.Generic_Def
	        ID                  string
            UserName            string
            UUID                string
            UserPsw             string
            UserRoleID          T_UserRole
            UserRole_ID         string
            Time1               string
            Money               string
    }
	
	func (_this *T_User) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"user",
            FnNewInst:	func()any{ return new (User) },
        }
    }
