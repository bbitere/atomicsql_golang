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
    type Employee1 struct /*atomicsql-table:"employee1"* / {
	        orm.Generic_MODEL
	        ID                  int32                         `json:"ID,omitempty"`
            IsActive            int16                         `json:"isActive"`
            UserID              *User                         `json:"-"`
            User_ID                                           `json:"user_ID"`
            DepartmentID        *Department                   `json:"-"`
            Department_ID                                     `json:"departm_ID"`
            Department2ID       sql.NullInt32                 `json:"departm2_ID"`
    }
	*/
    func (model  Employee1) GetID() int64 {
	    return int64( model.ID )
    }
	func (model  Employee1) SetID( id int64 ) {
	    model.ID = int32(id)
    }

    type T_Employee1 struct {
	        
			orm.Generic_Def
	        ID                  string
            IsActive            string
            UserID              T_User
            User_ID             string
            DepartmentID        T_Department
            Department_ID       string
            Department2ID       string
    }
	
	func (_this *T_Employee1) Def() *orm.TDefIncludeRelation{
        return &orm.TDefIncludeRelation{ 
            
            ValueDef: reflect.ValueOf( *_this),
			SqlTable:	"employee1",
            FnNewInst:	func(bFull bool)any{ 
				var model = new (Employee1);
				if( bFull ){
					
                        var defUser = T_User{}
				        model.UserID = (defUser.Def().FnNewInst(bFull)).(*User)
            
                        var defDepartment = T_Department{}
				        model.DepartmentID = (defDepartment.Def().FnNewInst(bFull)).(*Department)
				}
				return model;
			},
        }
    }
