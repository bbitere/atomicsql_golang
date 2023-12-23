/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /**/
	/*
    type Dbparams struct /*atomicsql-table:"dbparams"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            ParamName           string                        `json:"ParamName"`
            ParamValue          string                        `json:"ParamValue"`
    }
	*/
    func (model  Dbparams) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Dbparams) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Dbparams struct {
	        
			orm.Generic_Def
	        ID                  string
            ParamName           string
            ParamValue          string
    }
	
	func (_this *T_Dbparams) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"dbparams",
            FnNewInst:	func()any{ return new (Dbparams) },
        }
    }
