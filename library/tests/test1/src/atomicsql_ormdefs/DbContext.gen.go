/* this class is generated automatically by DB_Tool.exe exporter*/

		package atomicsql_sqldefs

		import (
		
			
			m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
			orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
		)
//------------------------------------------------------

		type DBContext struct {
			orm.DBContextBase

			
			Statusrole          *orm.DBTable[m.Statusrole]
			Statusrole_           m.T_Statusrole
		
			User                *orm.DBTable[m.User]
			User_                 m.T_User
		
			UserRole            *orm.DBTable[m.UserRole]
			UserRole_             m.T_UserRole
		
			Department          *orm.DBTable[m.Department]
			Department_           m.T_Department
		
			Employee1           *orm.DBTable[m.Employee1]
			Employee1_            m.T_Employee1
		
			Employee2           *orm.DBTable[m.Employee2]
			Employee2_            m.T_Employee2
		
			ProjectStatus       *orm.DBTable[m.ProjectStatus]
			ProjectStatus_        m.T_ProjectStatus
		
			Project             *orm.DBTable[m.Project]
			Project_              m.T_Project
		
		}		
		func (_this *DBContext) Constr(dbBase orm.DBContextBase)  (*DBContext,error) {
		
			_this.DBContextBase = dbBase;
			_this.DBContextBase.AllTables = make( map[string]*orm.DBTable[orm.IGeneric_MODEL])
			
			var sqlSchemaDef orm.TSchemaDef = orm.TSchemaDef{
				
	"Statusrole": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "statusRole",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "StatusName",
					SqlName:    "status_name",
					SqlType:    "VARCHAR",
					LangType:   "sql.NullString",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"User": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "user",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserName",
					SqlName:    "userName",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UUID",
					SqlName:    "UUID",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserPsw",
					SqlName:    "user_psw",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserRole_ID",
					SqlName:    "userRole_ID",
					SqlType:    "INT",
					LangType:   "*UserRole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "UserRoleID", "UserRole" },
				},
				{
					LangName:   "Time1",
					SqlName:    "time1",
					SqlType:    "TIMESTAMP",
					LangType:   "sql.NullTime",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Money",
					SqlName:    "money",
					SqlType:    "double precision",
					LangType:   "float64",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"UserRole": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "user_role",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "BOOLEAN",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "RoleName",
					SqlName:    "roleName",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Role_status_ID",
					SqlName:    "statusrole_ID",
					SqlType:    "INT",
					LangType:   "*Statusrole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "RoleStatusID", "Statusrole" },
				},
				{
					LangName:   "MyActive",
					SqlName:    "my_active",
					SqlType:    "BOOLEAN",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Department": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "departm",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "BOOLEAN",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NameDep",
					SqlName:    "name_dep",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Employee1": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "employee1",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "BOOLEAN",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "User_ID",
					SqlName:    "user_ID",
					SqlType:    "INT",
					LangType:   "*User",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:[]string{ "UserID", "User" },
				},
				{
					LangName:   "Department_ID",
					SqlName:    "departm_ID",
					SqlType:    "INT",
					LangType:   "*Department",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "DepartmentID", "Department" },
				},
				{
					LangName:   "Department2_ID",
					SqlName:    "departm2_ID",
					SqlType:    "INT",
					LangType:   "*Department",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "Department2ID", "Department" },
				},

			},
		},

	"Employee2": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "employee2",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"ProjectStatus": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "projstatus",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Project": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "project",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "INT",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "VARCHAR",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

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
                        Role_status_ID:      "userRole_ID.statusrole_ID",
                        RoleStatusID:        m.T_Statusrole {
                        ID:                  "userRole_ID.statusrole_ID.ID",
                        StatusName:          "userRole_ID.statusrole_ID.status_name", 
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
            Role_status_ID:      "statusrole_ID",
            RoleStatusID:        m.T_Statusrole {
                        ID:                  "statusrole_ID.ID",
                        StatusName:          "statusrole_ID.status_name", 
                        },
            MyActive:            "my_active",
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
                        Role_status_ID:      "user_ID.userRole_ID.statusrole_ID",
                        RoleStatusID:        m.T_Statusrole {
                        ID:                  "user_ID.userRole_ID.statusrole_ID.ID",
                        StatusName:          "user_ID.userRole_ID.statusrole_ID.status_name", 
                        },
                        MyActive:            "user_ID.userRole_ID.my_active", 
                        },
                        Time1:               "user_ID.time1",
                        Money:               "user_ID.money", 
                        },
            Department_ID:       "departm_ID",
            DepartmentID:        m.T_Department {
                        ID:                  "departm_ID.ID",
                        IsActive:            "departm_ID.isActive",
                        NameDep:             "departm_ID.name_dep", 
                        },
            Department2_ID:      "departm2_ID",
            Department2ID:       m.T_Department {
                        ID:                  "departm2_ID.ID",
                        IsActive:            "departm2_ID.isActive",
                        NameDep:             "departm2_ID.name_dep", 
                        },
			}

		
			_this.Employee2_ = m.T_Employee2 {
	        ID:                  "ID",
			}

		
			_this.ProjectStatus_ = m.T_ProjectStatus {
	        ID:                  "ID",
            Name:                "name",
			}

		
			_this.Project_ = m.T_Project {
	        ID:                  "ID",
            Name:                "name",
			}

		
			
			_this.DBContextBase.FOREIGN_KEYS = map[string]orm.TForeignKey{
								
				"user.userRole_ID":
				{  
					TgtTable_sqlName:"user_role", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"userRole_ID", 
					RootTable_sqlName:"user", 
					
					RootFldFk_langName:"UserRoleID", 
					RootFldFk_lang2Name:"UserRole_ID", 
				},
				
				
				"user_role.statusrole_ID":
				{  
					TgtTable_sqlName:"statusRole", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"statusrole_ID", 
					RootTable_sqlName:"user_role", 
					
					RootFldFk_langName:"RoleStatusID", 
					RootFldFk_lang2Name:"Role_status_ID", 
				},
				
				
				"employee1.user_ID":
				{  
					TgtTable_sqlName:"user", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"user_ID", 
					RootTable_sqlName:"employee1", 
					
					RootFldFk_langName:"UserID", 
					RootFldFk_lang2Name:"User_ID", 
				},
				
				
				"employee1.departm_ID":
				{  
					TgtTable_sqlName:"departm", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"departm_ID", 
					RootTable_sqlName:"employee1", 
					
					RootFldFk_langName:"DepartmentID", 
					RootFldFk_lang2Name:"Department_ID", 
				},
				
				
				"employee1.departm2_ID":
				{  
					TgtTable_sqlName:"departm", 
					TgtFldID_sqlName:"ID", 
					
					RootFldFk_sqlName:"departm2_ID", 
					RootTable_sqlName:"employee1", 
					
					RootFldFk_langName:"Department2ID", 
					RootFldFk_lang2Name:"Department2_ID", 
				},
				

			}
			
			_this.DBContextBase.DictTablesIncludeRelDefs	= map[string]*orm.TDefIncludeRelation{
				"statusRole": _this.Statusrole_.Def(),
					"user": _this.User_.Def(),
					"user_role": _this.UserRole_.Def(),
					"departm": _this.Department_.Def(),
					"employee1": _this.Employee1_.Def(),
					"employee2": _this.Employee2_.Def(),
					"projstatus": _this.ProjectStatus_.Def(),
					"project": _this.Project_.Def(),
					
			}
			
			_, err := _this.DBContextBase.Constr( dbBase.Dialect, sqlSchemaDef, _this)
			
			_this.LoadCompiledQuery();
			
			return _this, err
		}
		func New_DBContext(dbBase orm.DBContextBase) (*DBContext, error) {
			
			ret, err := (new(DBContext)).Constr(dbBase)
			
						
			ret.Statusrole =              (new(orm.DBTable[m.Statusrole])).Constr("statusRole", "Statusrole", &ret.DBContextBase)			
			ret.User =                    (new(orm.DBTable[m.User])).Constr("user", "User", &ret.DBContextBase)			
			ret.UserRole =                (new(orm.DBTable[m.UserRole])).Constr("user_role", "UserRole", &ret.DBContextBase)			
			ret.Department =              (new(orm.DBTable[m.Department])).Constr("departm", "Department", &ret.DBContextBase)			
			ret.Employee1 =               (new(orm.DBTable[m.Employee1])).Constr("employee1", "Employee1", &ret.DBContextBase)			
			ret.Employee2 =               (new(orm.DBTable[m.Employee2])).Constr("employee2", "Employee2", &ret.DBContextBase)			
			ret.ProjectStatus =           (new(orm.DBTable[m.ProjectStatus])).Constr("projstatus", "ProjectStatus", &ret.DBContextBase)			
			ret.Project =                 (new(orm.DBTable[m.Project])).Constr("project", "Project", &ret.DBContextBase)
			
			return ret, err
		}		
		func (_this *DBContext) Close(){
			
			_this.DBContextBase.Db.Close()
		}
		func (_this *DBContext) GetSqlName() string{
			return "postgres";
		}
		
		

			
		

