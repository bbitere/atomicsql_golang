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
    type StatusRole struct /*atomicsql-table:"statusrole"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            StatusName          sql.NullString                `json:"status_name"`
    }
	*/
    func (model  StatusRole) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  StatusRole) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_StatusRole struct {
	        
			orm.Generic_Def
	        ID                  string
            StatusName          string
    }
	
	func (_this *T_StatusRole) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
            FnNewInst:	func()any{ return new (StatusRole) },
        }
    }
