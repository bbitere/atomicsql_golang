/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /**/
	/*
    type Employee2 struct /*atomicsql-table:"employee2"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
    }
	*/
    func (model  Employee2) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Employee2) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Employee2 struct {
	        
			orm.Generic_Def
	        ID                  string
    }
	
	func (_this *T_Employee2) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"employee2",
            FnNewInst:	func()any{ return new (Employee2) },
        }
    }
