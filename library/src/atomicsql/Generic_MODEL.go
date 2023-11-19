package atomicsql

import (
	"reflect"
	"strconv"
)

/*
*
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
*/
type IGeneric_MODEL interface {
	GetID() int64
	//SetID(int64)

	//ReadRowSqlResult(sqlResult sql.Rows)
}

type Generic_MODEL struct {
	//ID string
	flagIsSaved bool
}

func (_this Generic_MODEL) GetID() int64 {
	return 0
}

//func (_this Generic_MODEL) SetID(id int64) {
//
//}

func (_this Generic_MODEL) GetUID() string {
	return strconv.Itoa(int(_this.GetID()))
}

/*
func (_this Generic_MODEL) DetacheModel()  {

	_this.SetID( 0 )
}*/

type IGeneric_Def interface {
	GetFK_IDs(m any) []int64
	Def() Generic_Def
}

type Generic_Def struct {
	FOREIGN_KEY_DEF string
}

type TDefIncludeRelation struct {
	ValueDef  reflect.Value //reflect value of this include Rel model
	FnNewInst func() any    //create the new instance of model
	PathFK    string        //full path of include
	RankFK    int           //rank: how many relations are in this include
	KeyFK     string        // last 2 items, defining the key in dict of ForeinKeys
	SqlTable  string        //sql table
	//FOREIGN_KEY string
}

func (_this *Generic_Def) Def() TDefIncludeRelation {
	return TDefIncludeRelation{ValueDef: reflect.ValueOf(_this), FnNewInst: nil}
}

/*
func (_this *Generic_Def) GetFK_IDs(m any )int64{
	return 0
}*/
