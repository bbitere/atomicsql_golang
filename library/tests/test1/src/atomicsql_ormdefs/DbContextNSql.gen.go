/* this class is generated automatically by DB_Tool.exe exporter*/

package atomicsql_sqldefs

import (
	orm "github.com/bbitere/atomicsql_golang.git/src/atomicNSql"
	atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
)

//------------------------------------------------------

		type DBContextNSql struct {
			orm.DBContextBaseNoSql

			
			Bar                 *orm.DBTableNoSql[m.Bar]
			Bar_                  m.T_Bar
		
			Bar1                *orm.DBTableNoSql[m.Bar1]
			Bar1_                 m.T_Bar1
		
			Bar2                *orm.DBTableNoSql[m.Bar2]
			Bar2_                 m.T_Bar2
		
			Dbparams            *orm.DBTableNoSql[m.Dbparams]
			Dbparams_             m.T_Dbparams
		
			Department          *orm.DBTableNoSql[m.Department]
			Department_           m.T_Department
		
			Employee1           *orm.DBTableNoSql[m.Employee1]
			Employee1_            m.T_Employee1
		
			Employee2           *orm.DBTableNoSql[m.Employee2]
			Employee2_            m.T_Employee2
		
			Project             *orm.DBTableNoSql[m.Project]
			Project_              m.T_Project
		
			ProjectStatus       *orm.DBTableNoSql[m.ProjectStatus]
			ProjectStatus_        m.T_ProjectStatus
		
			Statusrole          *orm.DBTableNoSql[m.Statusrole]
			Statusrole_           m.T_Statusrole
		
			User                *orm.DBTableNoSql[m.User]
			User_                 m.T_User
		
			UserRole            *orm.DBTableNoSql[m.UserRole]
			UserRole_             m.T_UserRole
		
		}		
		func (_this *DBContextNSql) Constr(dbBase orm.DBContextBaseNoSql)  (*DBContextNSql,error) {
		
			_this.DBContextBaseNoSql = dbBase;
			_this.DBContextBaseNoSql.AllTables = make( map[string]*orm.DBTableNoSql[atomicsql.IGeneric_MODEL])
			
			var sqlSchemaDef atomicsql.TSchemaDef = atomicsql.TSchemaDef{
				
	"Bar": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "bar",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "integer",
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
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Bar1": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "bar1",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "integer",
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
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Bar2": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "bar2",
			PrimaryColumnLangName: "Id",
			PrimaryColumnSqlName:  "Id",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "Id",
					SqlName:    "Id",
					SqlType:    "integer",
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
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NonuniqueInt",
					SqlName:    "nonunique_int",
					SqlType:    "integer",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Dbparams": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "dbparams",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "ParamName",
					SqlName:    "ParamName",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "ParamValue",
					SqlName:    "ParamValue",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Department": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "departm",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "boolean",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "NameDep",
					SqlName:    "name_dep",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Employee1": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "employee1",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "boolean",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "User_ID",
					SqlName:    "user_ID",
					SqlType:    "integer",
					LangType:   "*User",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "UserID", "User" },
				},
				{
					LangName:   "Department_ID",
					SqlName:    "departm_ID",
					SqlType:    "integer",
					LangType:   "*Department",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "DepartmentID", "Department" },
				},
				{
					LangName:   "Department2_ID",
					SqlName:    "departm2_ID",
					SqlType:    "integer",
					LangType:   "*Department",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "Department2ID", "Department" },
				},

			},
		},

	"Employee2": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "employee2",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Project": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "project",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"ProjectStatus": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "projstatus",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Name",
					SqlName:    "name",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},

			},
		},

	"Statusrole": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "statusRole",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "StatusName",
					SqlName:    "status_name",
					SqlType:    "character varying",
					LangType:   "sql.NullString",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},

			},
		},

	"User": orm.TDefTable{
			SchemaTable:           "public",
			SqlTableName:          "user",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserName",
					SqlName:    "userName",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UUID",
					SqlName:    "UUID",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserPsw",
					SqlName:    "user_psw",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "UserRole_ID",
					SqlName:    "userRole_ID",
					SqlType:    "integer",
					LangType:   "*UserRole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "UserRoleID", "UserRole" },
				},
				{
					LangName:   "Time1",
					SqlName:    "time1",
					SqlType:    "timestamp without time zone",
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
			SchemaTable:           "public",
			SqlTableName:          "user_role",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "ID",
					SqlName:    "ID",
					SqlType:    "integer",
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "IsActive",
					SqlName:    "isActive",
					SqlType:    "boolean",
					LangType:   "bool",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "RoleName",
					SqlName:    "roleName",
					SqlType:    "character varying",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Role_status_ID",
					SqlName:    "role_status_ID",
					SqlType:    "integer",
					LangType:   "*Statusrole",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:[]string{ "RoleStatusID", "Statusrole" },
				},
				{
					LangName:   "MyActive",
					SqlName:    "my_active",
					SqlType:    "boolean",
					LangType:   "bool",
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

		
			/*
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
					TgtTable_sqlName:"statusRole", 
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
			*/
			
			_this.DBContextBaseNoSql.DictTablesIncludeRelDefs	= map[string]*orm.TDefIncludeRelation{
				"bar": _this.Bar_.Def(),
					"bar1": _this.Bar1_.Def(),
					"bar2": _this.Bar2_.Def(),
					"dbparams": _this.Dbparams_.Def(),
					"departm": _this.Department_.Def(),
					"employee1": _this.Employee1_.Def(),
					"employee2": _this.Employee2_.Def(),
					"project": _this.Project_.Def(),
					"projstatus": _this.ProjectStatus_.Def(),
					"statusRole": _this.Statusrole_.Def(),
					"user": _this.User_.Def(),
					"user_role": _this.UserRole_.Def(),
					
			}
			
			_, err := _this.DBContextBaseNoSql.Constr( dbBase.Dialect, sqlSchemaDef, _this)
			
			//_this.LoadCompiledQuery();
			
			return _this, err
		}
		func New_DBContextNSql(dbBase orm.DBContextBaseNoSql) (*DBContextNSql, error) {
			
			ret, err := (new(DBContextNSql)).Constr(dbBase)
			
						
			ret.Bar =                     (new(orm.DBTableNoSql[m.Bar])).Constr("bar", "Bar", &ret.DBContextBaseNoSql)			
			ret.Bar1 =                    (new(orm.DBTableNoSql[m.Bar1])).Constr("bar1", "Bar1", &ret.DBContextBaseNoSql)			
			ret.Bar2 =                    (new(orm.DBTableNoSql[m.Bar2])).Constr("bar2", "Bar2", &ret.DBContextBaseNoSql)			
			ret.Dbparams =                (new(orm.DBTableNoSql[m.Dbparams])).Constr("dbparams", "Dbparams", &ret.DBContextBaseNoSql)			
			ret.Department =              (new(orm.DBTableNoSql[m.Department])).Constr("departm", "Department", &ret.DBContextBaseNoSql)			
			ret.Employee1 =               (new(orm.DBTableNoSql[m.Employee1])).Constr("employee1", "Employee1", &ret.DBContextBaseNoSql)			
			ret.Employee2 =               (new(orm.DBTableNoSql[m.Employee2])).Constr("employee2", "Employee2", &ret.DBContextBaseNoSql)			
			ret.Project =                 (new(orm.DBTableNoSql[m.Project])).Constr("project", "Project", &ret.DBContextBaseNoSql)			
			ret.ProjectStatus =           (new(orm.DBTableNoSql[m.ProjectStatus])).Constr("projstatus", "ProjectStatus", &ret.DBContextBaseNoSql)			
			ret.Statusrole =              (new(orm.DBTableNoSql[m.Statusrole])).Constr("statusRole", "Statusrole", &ret.DBContextBaseNoSql)			
			ret.User =                    (new(orm.DBTableNoSql[m.User])).Constr("user", "User", &ret.DBContextBaseNoSql)			
			ret.UserRole =                (new(orm.DBTableNoSql[m.UserRole])).Constr("user_role", "UserRole", &ret.DBContextBaseNoSql)
			
			return ret, err
		}		
		func (_this *DBContextNSql) Close(){
			
			//_this.DBContextBase.Db.Close()
		}
		func (_this *DBContextNSql) GetSqlName() string{
			return "postgres";
		}
		
		

			
		

