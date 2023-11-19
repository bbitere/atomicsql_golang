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
    type Employee1 struct /*atomicsql-table:"employee1"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            IsActive            bool                          `json:"isActive"`
            UserID              *User                         `json:"-"`
            User_ID             sql.NullInt32                 `json:"user_ID"`
            DepartmID           *Department                   `json:"-"`
            Departm_ID          sql.NullInt32                 `json:"departm_ID"`
            Departm2ID          *Department                   `json:"-"`
            Departm2_ID         sql.NullInt32                 `json:"departm2_ID"`
    }
	*/
    func (model  Employee1) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Employee1) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Employee1 struct {
	        
			orm.Generic_Def
	        ID                  string
            IsActive            string
            UserID              T_User
            User_ID             string
            DepartmID           T_Department
            Departm_ID          string
            Departm2ID          T_Department
            Departm2_ID         string
    }
	
	func (_this *T_Employee1) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"employee1",
            FnNewInst:	func()any{ return new (Employee1) },
        }
    }
