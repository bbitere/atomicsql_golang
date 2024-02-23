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
    type Bar struct /*atomicsql-table:"bar"* / {
	        orm.Generic_MODEL
	        Id                  int32                         `json:"Id,omitempty"`
            FooId               sql.NullInt64                 `json:"foo_id"`
            UniqueInt           sql.NullInt32                 `json:"unique_int"`
            NonuniqueInt        sql.NullInt32                 `json:"nonunique_int"`
    }
*/
    func (model  Bar) GetID() int64 {
	    return int64( model.Id )
    }
	func (model  Bar) SetID( id int64 ) {
	    model.Id = int64(id)
    }

    type T_Bar struct {
	        
			orm.Generic_Def
	        Id                  string
            FooId               string
            UniqueInt           string
            NonuniqueInt        string
    }
	
	func (_this *T_Bar) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"bar",
            FnNewInst:	func(bFull bool)any{ 
				var model = new (Bar);
				if( bFull ){
					
				}
				return model;
			},
        }
    }
