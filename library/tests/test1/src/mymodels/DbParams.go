package mymodels

import (
	sql "database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

type Dbparams struct /*atomicsql-table:"dbparams"*/ {
	orm.Generic_MODEL

	ID         int32  `json:"ID,omitempty"`
	ParamName  string `json:"ParamName"`
	ParamValue string `json:"ParamValue"`
}

type Bar2 struct /*atomicsql-table:"bar2"*/ {
	orm.Generic_MODEL

	Id           int64         `json:"Id,omitempty"`
	FooId        sql.NullInt64 `json:"foo_id"`
	UniqueInt    sql.NullInt32 `json:"unique_int"`
	NonuniqueInt sql.NullInt32 `json:"nonunique_int"`
}

type Bar1 struct /*atomicsql-table:"bar1"*/ {
	orm.Generic_MODEL

	Id           int64         `json:"Id,omitempty"`
	FooId        sql.NullInt64 `json:"foo_id"`
	UniqueInt    sql.NullInt32 `json:"unique_int"`
	NonuniqueInt sql.NullInt32 `json:"nonunique_int"`
}

type Bar struct /*atomicsql-table:"bar"*/ {
	orm.Generic_MODEL

	Id           int64         `json:"Id,omitempty"`
	FooId        sql.NullInt64 `json:"foo_id"`
	UniqueInt    sql.NullInt32 `json:"unique_int"`
	NonuniqueInt sql.NullInt32 `json:"nonunique_int"`
}
