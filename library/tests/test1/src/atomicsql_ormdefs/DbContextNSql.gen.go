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
		
		}		
		func (_this *DBContextNSql) Constr(dbBase orm.DBContextBaseNoSql)  (*DBContextNSql,error) {
		
			_this.DBContextBaseNoSql = dbBase;
			_this.DBContextBaseNoSql.AllTables = make( map[string]*orm.DBTableNoSql[atomicsql.IGeneric_MODEL])
			
			var sqlSchemaDef atomicsql.TSchemaDef = atomicsql.TSchemaDef{
				
	"Statusrole": orm.TDefTable{
			SchemaTable:           "",
			SqlTableName:          "statusRole",
			PrimaryColumnLangName: "NoSqlID",
			PrimaryColumnSqlName:  "_id",
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
			PrimaryColumnLangName: "NoSqlID",
			PrimaryColumnSqlName:  "_id",
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
					LangName:   "UserRoleID",
					SqlName:    "userRole_ID",
					SqlType:    "TEXT",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
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
			PrimaryColumnLangName: "NoSqlID",
			PrimaryColumnSqlName:  "_id",
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
					LangName:   "RoleStatusID",
					SqlName:    "Statusrole",
					SqlType:    "TEXT",
					LangType:   "string",
					Flags:      "",
					IsPrimary:  false,
					IsNullable: false,
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

			}
			
			_this.Statusrole_ = m.T_Statusrole {
	        StatusName:          "status_name",
			}

		
			_this.User_ = m.T_User {
	        UserName:            "userName",
            UUID:                "UUID",
            UserPsw:             "user_psw",
            UserRoleID:          m.T_UserRole {
                        IsActive:            "userRole_ID.isActive",
                        RoleName:            "userRole_ID.roleName",
                        RoleStatusID:        m.T_Statusrole {
                        StatusName:          "userRole_ID.Statusrole.status_name", 
                        },
                        MyActive:            "userRole_ID.my_active", 
                        },
            Time1:               "time1",
            Money:               "money",
			}

		
			_this.UserRole_ = m.T_UserRole {
	        IsActive:            "isActive",
            RoleName:            "roleName",
            RoleStatusID:        m.T_Statusrole {
                        StatusName:          "Statusrole.status_name", 
                        },
            MyActive:            "my_active",
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
			
			return ret, err
		}		
		func (_this *DBContextNSql) Close(){
			
			/*_this.DBContextBaseNoSql.Db.Close()*/
		}
		func (_this *DBContextNSql) GetSqlName() string{
			return "mongodb";
		}

