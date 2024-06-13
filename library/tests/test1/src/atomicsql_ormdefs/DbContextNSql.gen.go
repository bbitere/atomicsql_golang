/* this class is generated automatically by DB_Tool.exe exporter*/

		package atomicsql_sqldefs

		import (
					
			m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
			orm	  "github.com/bbitere/atomicsql_golang.git/src/atomicNSql"
			atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
		)
//------------------------------------------------------

		type DBContextNSql struct {
			orm.DBContextBaseNoSql

			
			Statusrole          *orm.DBTableNoSql[m.Statusrole]
			Statusrole_           m.T_Statusrole
		
			User                *orm.DBTableNoSql[m.User]
			User_                 m.T_User
		
			UserRole            *orm.DBTableNoSql[m.UserRole]
			UserRole_             m.T_UserRole
		
			Department          *orm.DBTableNoSql[m.Department]
			Department_           m.T_Department
		
			Employee1           *orm.DBTableNoSql[m.Employee1]
			Employee1_            m.T_Employee1
		
			Employee2           *orm.DBTableNoSql[m.Employee2]
			Employee2_            m.T_Employee2
		
			ProjectStatus       *orm.DBTableNoSql[m.ProjectStatus]
			ProjectStatus_        m.T_ProjectStatus
		
			Project             *orm.DBTableNoSql[m.Project]
			Project_              m.T_Project
		
		}		
		func (_this *DBContextNSql) Constr(dbBase orm.DBContextBaseNoSql)  (*DBContextNSql,error) {
		
			_this.DBContextBaseNoSql = dbBase;
			_this.DBContextBaseNoSql.AllTables = make( map[string]*orm.DBTableNoSql[atomicsql.IGeneric_MODEL])
			
			var sqlSchemaDef atomicsql.TSchemaDef = atomicsql.TSchemaDef{
				
	"Statusrole": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "statusRole",
			PrimaryColumnLangName: "ID",
			PrimaryColumnSqlName:  "ID",
			Columns: []orm.TSqlColumnDef{
								{
					LangName:   "NoSqlID",
					SqlName:    "_id",
					SqlType:    "ObjectID",
					LangType:   "primitive.ObjectID",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
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
					LangName:   "NoSqlID",
					SqlName:    "_id",
					SqlType:    "ObjectID",
					LangType:   "primitive.ObjectID",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
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
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
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
					LangName:   "NoSqlID",
					SqlName:    "_id",
					SqlType:    "ObjectID",
					LangType:   "primitive.ObjectID",
					Flags:      "",
					IsPrimary:  true,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
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
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
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
					LangType:   "int32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Department_ID",
					SqlName:    "departm_ID",
					SqlType:    "INT",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
				},
				{
					LangName:   "Department2_ID",
					SqlName:    "departm2_ID",
					SqlType:    "INT",
					LangType:   "sql.NullInt32",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: true,
					ForeignKeyLangName:nil,
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
            Time1:               "time1",
            Money:               "money",
			}

		
			_this.UserRole_ = m.T_UserRole {
	        ID:                  "ID",
            IsActive:            "isActive",
            RoleName:            "roleName",
            Role_status_ID:      "statusrole_ID",
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
            Department_ID:       "departm_ID",
            Department2_ID:      "departm2_ID",
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

		
			/*
			_this.DBContextBaseNoSql.FOREIGN_KEYS = map[string]orm.TForeignKey{
				{txtForeignKeys}
			}
			/*
			_this.DBContextBaseNoSql.DictTablesIncludeRelDefs	= map[string]*orm.TDefIncludeRelation{
				{txtIncludeRelation}
			}
			*/
			
			_, err := _this.DBContextBaseNoSql.Constr( dbBase.Dialect, sqlSchemaDef, _this)
			
			_this.LoadCompiledQuery();
			
			return _this, err
		}
		func New_DBContextNSql(dbBase orm.DBContextBaseNoSql) (*DBContextNSql, error) {
			
			ret, err := (new(DBContextNSql)).Constr(dbBase)
			
						
			ret.Statusrole =              (new(orm.DBTableNoSql[m.Statusrole])).Constr("statusRole", "Statusrole", &ret.DBContextBaseNoSql)			
			ret.User =                    (new(orm.DBTableNoSql[m.User])).Constr("user", "User", &ret.DBContextBaseNoSql)			
			ret.UserRole =                (new(orm.DBTableNoSql[m.UserRole])).Constr("user_role", "UserRole", &ret.DBContextBaseNoSql)			
			ret.Department =              (new(orm.DBTableNoSql[m.Department])).Constr("departm", "Department", &ret.DBContextBaseNoSql)			
			ret.Employee1 =               (new(orm.DBTableNoSql[m.Employee1])).Constr("employee1", "Employee1", &ret.DBContextBaseNoSql)			
			ret.Employee2 =               (new(orm.DBTableNoSql[m.Employee2])).Constr("employee2", "Employee2", &ret.DBContextBaseNoSql)			
			ret.ProjectStatus =           (new(orm.DBTableNoSql[m.ProjectStatus])).Constr("projstatus", "ProjectStatus", &ret.DBContextBaseNoSql)			
			ret.Project =                 (new(orm.DBTableNoSql[m.Project])).Constr("project", "Project", &ret.DBContextBaseNoSql)
			
			return ret, err
		}		
		func (_this *DBContextNSql) Close(){
			
			/*_this.DBContextBaseNoSql.Db.Close()*/
		}
		func (_this *DBContextNSql) GetSqlName() string{
			return "mongodb";
		}

