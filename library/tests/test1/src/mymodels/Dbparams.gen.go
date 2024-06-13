/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /**/
	
    type DBParams struct /*atomicsql-table:"DBParams"*/ {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            ParamName           string                        `json:"ParamName"`
            ParamValue          string                        `json:"ParamValue"`
    }
	
    func (model  DBParams) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  DBParams) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_DBParams struct {
	        
			orm.Generic_Def
	        ID                  string
            ParamName           string
            ParamValue          string
    }
	
	func (_this *T_DBParams) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"DBParams",
            FnNewInst:	func(bFull bool)any{ return new (DBParams) },
        }
    }
