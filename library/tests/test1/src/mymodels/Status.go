package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang"

	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Statusrole struct /*atomicsql-table:"statusRole"*/ {
	orm.Generic_MODEL   `bson:"-"`

	NoSqlID		primitive.ObjectID		`bson:"_id,omitempty"`
	ID         int32          `json:"ID,omitempty"      bson:"-"`
	StatusName sql.NullString `json:"status_name"       bson:"status_name"`
}
