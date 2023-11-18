/* this class is generated automatically by DB_Tool.exe exporter*/

package mymodels

import (
	"reflect"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

/**/

/*

    type ProjectStatus struct /*atomicsql-table:"projstatus"* / {

	        orm.Generic_MODEL

	        ID                  int32                         `json:"ID,omitempty"`
            Name                string                        `json:"name"`

    }

*/

func (model ProjectStatus) GetID() int64 {

	return int64(model.ID)

}

func (model ProjectStatus) SetID(id int64) {

	model.ID = int32(id)

}

type T_ProjectStatus struct {
	orm.Generic_Def

	ID   string
	Name string
}

func (_this *T_ProjectStatus) Def() *orm.TDefIncludeRelation {

	return &orm.TDefIncludeRelation{

		ValueDef: reflect.ValueOf(*_this),

		FnNewInst: func() any { return new(ProjectStatus) },
	}

}
