/* this class is generated automatically by DB_Tool.exe exporter*/

package mymodels

import (
	"reflect"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

/*import (
	    sql "database/sql"
        )*/
/*
    type Bar1 struct /*atomicsql-table:"bar1"* / {
	        orm.Generic_MODEL
	        Id                  int32                         `json:"Id,omitempty"`
            FooId               sql.NullInt64                 `json:"foo_id"`
            UniqueInt           sql.NullInt32                 `json:"unique_int"`
            NonuniqueInt        sql.NullInt32                 `json:"nonunique_int"`
    }
*/
    func (model  Bar1) GetID() int64 {
	    return int64( model.Id )
    }
	func (model  Bar1) SetID( id int64 ) {
	    model.Id = int64(id)
    }

    type T_Bar1 struct {
	        
			orm.Generic_Def
	        Id                  string
            FooId               string
            UniqueInt           string
            NonuniqueInt        string
    }
	
	func (_this *T_Bar1) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"bar1",
            FnNewInst:	func(bFull bool)any{ 
				var model = new (Bar1);
				if( bFull ){
					
				}
				return model;
			},
        }
    }
