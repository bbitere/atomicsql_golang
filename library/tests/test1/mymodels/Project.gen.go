/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /**/
	/*
    type Project struct /*atomicsql-table:"project"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            Name                string                        `json:"name"`
    }
	*/
    func (model  Project) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Project) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Project struct {
	        
			orm.Generic_Def
	        ID                  string
            Name                string
    }
	
	func (_this *T_Project) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
            FnNewInst:	func()any{ return new (Project) },
        }
    }
