/* this class is generated automatically by DB_Tool.exe exporter*/

    package mymodels
	import (
        "reflect"
        orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
    )
	
	
        /**/
	/*
    type Department struct /*atomicsql-table:"departm"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            IsActive            bool                          `json:"isActive"`
            NameDep             string                        `json:"name_dep"`
    }
	*/
    func (model  Department) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Department) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Department struct {
	        
			orm.Generic_Def
	        ID                  string
            IsActive            string
            NameDep             string
    }
	
	func (_this *T_Department) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"departm",
            FnNewInst:	func()any{ return new (Department) },
        }
    }
