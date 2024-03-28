package mymodels

import (
	//orm	  "github.com/bbitere/atomicsql_golang/tree/main/library/src/atomicsql"
	"database/sql"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type User struct /*atomicsql-table:"user"*/ {

	orm.Generic_MODEL  `bson:"-"`
	NoSqlID				primitive.ObjectID				`json:"-"             bson:"_id,omitempty"`
	

	ID                  int                           	`json:"ID"            bson:"-"`
	UserName            string                        	`json:"userName"      bson:"userName"`
	UUID                string                        	`json:"UUID"          bson:"UUID"`
	UserPsw             string                        	`json:"user_psw"      bson:"user_psw"`
	UserRoleID          *UserRole                     	`json:"-"             bson:"userRole_ID"`
	UserRole_ID         sql.NullInt32                 	`json:"userRole_ID"   bson:"-"`
	/*Misc1               []uint8                       `json:"misc1"`*/
	Time1               sql.NullTime                  	`json:"time1"         bson:"time1"`
	Money               float64                       	`json:"money"         bson:"money"`


}