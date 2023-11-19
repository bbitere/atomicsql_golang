/* this class is generated automatically by DB_Tool.exe exporter*/

		package atomicsql_sqldefs

		import (
		
			
			m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
			orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
		)
//------------------------------------------------------

		type DBContext struct {
			orm.DBContextBase

			
			Bar                 *orm.DBTable[m.Bar]
			Bar_                  m.T_Bar
		
			Bar1                *orm.DBTable[m.Bar1]
			Bar1_                 m.T_Bar1
		
			Bar2                *orm.DBTable[m.Bar2]
			Bar2_                 m.T_Bar2
		
			Dbparams            *orm.DBTable[m.Dbparams]
			Dbparams_             m.T_Dbparams
		
			Department          *orm.DBTable[m.Department]
			Department_           m.T_Department
		
			Employee1           *orm.DBTable[m.Employee1]
			Employee1_            m.T_Employee1
		
			Employee2           *orm.DBTable[m.Employee2]
			Employee2_            m.T_Employee2
		
			Project             *orm.DBTable[m.Project]
			Project_              m.T_Project
		
			ProjectStatus       *orm.DBTable[m.ProjectStatus]
			ProjectStatus_        m.T_ProjectStatus
		
			Statusrole          *orm.DBTable[m.Statusrole]
			Statusrole_           m.T_Statusrole
		
			User                *orm.DBTable[m.User]
			User_                 m.T_User
		
			UserRole            *orm.DBTable[m.UserRole]
			UserRole_             m.T_UserRole
		
		}
		
		func (_this *DBContext) Constr(dbBase orm.DBContextBase)  (*DBContext,error) {
		
			_this.DBContextBase = dbBase;
			_this.DBContextBase.AllTables = make( map[string]*orm.DBTable[orm.IGeneric_MODEL])
			
			var sqlSchemaDef orm.TSchemaDef = orm.TSchemaDef{
				
	"Bar": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "bar",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "FooId",
					SqlName:    "foo_id",
					SqlType:    "bigint",
					LangType:   "sql.NullInt64",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UniqueInt",
					SqlName:    "unique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Bar1": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "bar1",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "FooId",
					SqlName:    "foo_id",
					SqlType:    "bigint",
					LangType:   "sql.NullInt64",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UniqueInt",
					SqlName:    "unique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Bar2": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "bar2",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "FooId",
					SqlName:    "foo_id",
					SqlType:    "bigint",
					LangType:   "sql.NullInt64",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UniqueInt",
					SqlName:    "unique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Dbparams": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "dbparams",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "ParamName",
					SqlName:    "ParamName",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "ParamValue",
					SqlName:    "ParamValue",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Department": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "departm",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "tinyint",
					LangType:   "int16",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NameDep",
					SqlName:    "name_dep",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Employee1": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "employee1",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "tinyint",
					LangType:   "int16",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "User_ID",
					SqlName:    "user_ID",
					SqlType:    "int",
					LangType:   "*User",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "UserID", "User" },
				},
				{
					LangName:   "Departm_ID",
					SqlName:    "departm_ID",
					SqlType:    "int",
					LangType:   "*Department",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "DepartmID", "Department" },
				},
				{
					LangName:   "Departm2ID",
					SqlName:    "departm2_ID",
					SqlType:    "int",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Employee2": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "employee2",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Project": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "project",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"ProjectStatus": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "projstatus",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Statusrole": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "statusrole",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "StatusName",
					SqlName:    "status_name",
					SqlType:    "varchar",
					LangType:   "sql.NullString",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"User": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "user",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserName",
					SqlName:    "userName",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UUID",
					SqlName:    "UUID",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserPsw",
					SqlName:    "user_psw",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserRole_ID",
					SqlName:    "userRole_ID",
					SqlType:    "int",
					LangType:   "*UserRole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "UserRoleID", "UserRole" },
				},
				{
					LangName:   "Time1",
					SqlName:    "time1",
					SqlType:    "timestamp",
					LangType:   "sql.NullTime",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Money",
					SqlName:    "money",
					SqlType:    "double",
					LangType:   "float64",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"UserRole": orm.TDefTable{
			SchemaTable:           "test1",
			SqlTableName:          "user_role",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "int",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "tinyint",
					LangType:   "int16",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "RoleName",
					SqlName:    "roleName",
					SqlType:    "varchar",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Role_status_ID",
					SqlName:    "role_status_ID",
					SqlType:    "int",
					LangType:   "*Statusrole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "RoleStatusID", "Statusrole" },
				},
				{
					LangName:   "MyActive",
					SqlName:    "my_active",
					SqlType:    "tinyint",
					LangType:   "int16",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

			}
			
			_this.Bar_ = m.T_Bar {
	        Id:                  "Id",
            FooId:               "foo_id",
            UniqueInt:           "unique_int",
            NonuniqueInt:        "nonunique_int",
			}

		
			_this.Bar1_ = m.T_Bar1 {
	        Id:                  "Id",
            FooId:               "foo_id",
            UniqueInt:           "unique_int",
            NonuniqueInt:        "nonunique_int",
			}

		
			_this.Bar2_ = m.T_Bar2 {
	        Id:                  "Id",
            FooId:               "foo_id",
            UniqueInt:           "unique_int",
            NonuniqueInt:        "nonunique_int",
			}

		
			_this.Dbparams_ = m.T_Dbparams {
	        ID:                  "ID",
            ParamName:           "ParamName",
            ParamValue:          "ParamValue",
			}

		
			_this.Department_ = m.T_Department {
	        ID:                  "ID",
            IsActive:            "isActive",
            NameDep:             "name_dep",
			}

		
			_this.Employee1_ = m.T_Employee1 {
	        ID:                  "ID",
            IsActive:            "isActive",
            User_ID:             "user_ID",
            UserID:              m.T_User {
                        ID:                  "user_ID.ID",
                        UserName:            "user_ID.userName",
                        UUID:                "user_ID.UUID",
                        UserPsw:             "user_ID.user_psw",
                        UserRole_ID:         "user_ID.userRole_ID",
                        UserRoleID:          m.T_UserRole {
                        ID:                  "user_ID.userRole_ID.ID",
                        IsActive:            "user_ID.userRole_ID.isActive",
                        RoleName:            "user_ID.userRole_ID.roleName",
                        Role_status_ID:      "user_ID.userRole_ID.role_status_ID",
                        RoleStatusID:        m.T_Statusrole {
                        ID:                  "user_ID.userRole_ID.role_status_ID.ID",
                        StatusName:          "user_ID.userRole_ID.role_status_ID.status_name", 
                        },
                        MyActive:            "user_ID.userRole_ID.my_active", 
                        },
                        Time1:               "user_ID.time1",
                        Money:               "user_ID.money", 
                        },
            Departm_ID:          "departm_ID",
            DepartmID:           m.T_Department {
                        ID:                  "departm_ID.ID",
                        IsActive:            "departm_ID.isActive",
                        NameDep:             "departm_ID.name_dep", 
                        },
            Departm2ID:          "departm2_ID",
			}

		
			_this.Employee2_ = m.T_Employee2 {
	        ID:                  "ID",
			}

		
			_this.Project_ = m.T_Project {
	        ID:                  "ID",
            Name:                "name",
			}

		
			_this.ProjectStatus_ = m.T_ProjectStatus {
	        ID:                  "ID",
            Name:                "name",
			}

		
			_this.Statusrole_ = m.T_Statusrole {
	        ID:                  "ID",
            StatusName:          "status_name",
			}

		
			_this.User_ = m.T_User {
	        ID:                  "ID",
            UserName:            "userName",
            UUID:                "UUID",
            UserPsw:             "user_psw",
            UserRole_ID:         "userRole_ID",
            UserRoleID:          m.T_UserRole {
                        ID:                  "userRole_ID.ID",
                        IsActive:            "userRole_ID.isActive",
                        RoleName:            "userRole_ID.roleName",
                        Role_status_ID:      "userRole_ID.role_status_ID",
                        RoleStatusID:        m.T_Statusrole {
                        ID:                  "userRole_ID.role_status_ID.ID",
                        StatusName:          "userRole_ID.role_status_ID.status_name", 
                        },
                        MyActive:            "userRole_ID.my_active", 
                        },
            Time1:               "time1",
            Money:               "money",
			}

		
			_this.UserRole_ = m.T_UserRole {
	        ID:                  "ID",
            IsActive:            "isActive",
            RoleName:            "roleName",
            Role_status_ID:      "role_status_ID",
            RoleStatusID:        m.T_Statusrole {
                        ID:                  "role_status_ID.ID",
                        StatusName:          "role_status_ID.status_name", 
                        },
            MyActive:            "my_active",
			}

		
			
			_this.DBContextBase.FOREIGN_KEYS = map[string]orm.TForeignKey{
								
				"employee1.user_ID":
				{  
					TgtTable_sqlName:"user", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"user_ID", 
					RootTable_sqlName:"employee1", 
					
					RootFldFk_langName:"UserID", 
					RootFldFk_lang2Name:"User_ID", 
				},
				
				
				"user.userRole_ID":
				{  
					TgtTable_sqlName:"user_role", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"userRole_ID", 
					RootTable_sqlName:"user", 
					
					RootFldFk_langName:"UserRoleID", 
					RootFldFk_lang2Name:"UserRole_ID", 
				},
				
				
				"user_role.role_status_ID":
				{  
					TgtTable_sqlName:"statusrole", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"role_status_ID", 
					RootTable_sqlName:"user_role", 
					
					RootFldFk_langName:"RoleStatusID", 
					RootFldFk_lang2Name:"Role_status_ID", 
				},
				
				
				"employee1.departm_ID":
				{  
					TgtTable_sqlName:"departm", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"departm_ID", 
					RootTable_sqlName:"employee1", 
					
					RootFldFk_langName:"DepartmID", 
					RootFldFk_lang2Name:"Departm_ID", 
				},
				

			}
			
			_, err := _this.DBContextBase.Constr( dbBase.Dialect, sqlSchemaDef)
			
			_this.LoadCompiledQuery();
			
			return _this, err
		}

		func New_DBContext(dbBase orm.DBContextBase) (*DBContext, error) {
			
			ret, err := (new(DBContext)).Constr(dbBase)
			
						
			ret.Bar =                     (new(orm.DBTable[m.Bar])).Constr("bar", "Bar", &ret.DBContextBase)			
			ret.Bar1 =                    (new(orm.DBTable[m.Bar1])).Constr("bar1", "Bar1", &ret.DBContextBase)			
			ret.Bar2 =                    (new(orm.DBTable[m.Bar2])).Constr("bar2", "Bar2", &ret.DBContextBase)			
			ret.Dbparams =                (new(orm.DBTable[m.Dbparams])).Constr("dbparams", "Dbparams", &ret.DBContextBase)			
			ret.Department =              (new(orm.DBTable[m.Department])).Constr("departm", "Department", &ret.DBContextBase)			
			ret.Employee1 =               (new(orm.DBTable[m.Employee1])).Constr("employee1", "Employee1", &ret.DBContextBase)			
			ret.Employee2 =               (new(orm.DBTable[m.Employee2])).Constr("employee2", "Employee2", &ret.DBContextBase)			
			ret.Project =                 (new(orm.DBTable[m.Project])).Constr("project", "Project", &ret.DBContextBase)			
			ret.ProjectStatus =           (new(orm.DBTable[m.ProjectStatus])).Constr("projstatus", "ProjectStatus", &ret.DBContextBase)			
			ret.Statusrole =              (new(orm.DBTable[m.Statusrole])).Constr("statusrole", "Statusrole", &ret.DBContextBase)			
			ret.User =                    (new(orm.DBTable[m.User])).Constr("user", "User", &ret.DBContextBase)			
			ret.UserRole =                (new(orm.DBTable[m.UserRole])).Constr("user_role", "UserRole", &ret.DBContextBase)
			
			return ret, err
		}

		
		func (_this *DBContext) Close(){
			
			_this.DBContextBase.Db.Close()
		}

