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
    type Statusrole struct /*atomicsql-table:"statusrole"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            StatusName          sql.NullString                `json:"status_name"`
    }
	*/
    func (model  Statusrole) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Statusrole) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Statusrole struct {
	        
			orm.Generic_Def
	        ID                  string
            StatusName          string
    }
	
	func (_this *T_Statusrole) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"statusrole",
            FnNewInst:	func(bFull bool)any{ 
				var model = new (Statusrole);
				if( bFull ){
					
				}
				return model;
			},
        }
    }
