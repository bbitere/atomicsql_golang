package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type User struct /*atomicsql-table:"user"*/ {

	orm.Generic_MODEL  `bson:"-"`
	NoSqlID				primitive.ObjectID				`bson:"_id,omitempty"`
	
	ID                  int                           	`bson:"-"`
	UserName            string                        	`bson:"userName"`
	UUID                string                        	`bson:"UUID"`
	UserPsw             string                        	`bson:"user_psw"`
	UserRoleID          *UserRole                     	`bson:"userRole_ID"`
	UserRole_ID         sql.NullInt32                 	`bson:"-"`
	/*Misc1               []uint8                       `json:"misc1"`*/
	Time1               sql.NullTime                  	`bson:"time1"`
	Money               float64                       	`bson:"money"`


}