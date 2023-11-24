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
    type Bar2 struct /*atomicsql-table:"bar2"* / {
	        orm.Generic_MODEL
	        Id                  int32                         `json:"Id,omitempty"`
            FooId               sql.NullInt64                 `json:"foo_id"`
            UniqueInt           sql.NullInt32                 `json:"unique_int"`
            NonuniqueInt        sql.NullInt32                 `json:"nonunique_int"`
    }
	*/
    func (model  Bar2) GetID() int64 {
	    return int64( model.Id )
    }
	func (model  Bar2) SetID( id int64 ) {
	    model.Id = int32(id)
    }

    type T_Bar2 struct {
	        
			orm.Generic_Def
	        Id                  string
            FooId               string
            UniqueInt           string
            NonuniqueInt        string
    }
	
	func (_this *T_Bar2) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"bar2",
            FnNewInst:	func()any{ return new (Bar2) },
        }
    }
