
/* this class is generated automatically by compiler AtomicSql goscanner.exe exporter*/

package atomicsql_sqldefs
	import (
		orm	 "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	)

func (_this *DBContext) LoadCompiledQuery(){

	_this.DBContextBase.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
		"asdax-W1": 
		{
			//CompiledQuery: {#@money@#}>={@@UserMoney@@} AND ({#@userRole_ID.roleName@#}={@@RoleNameDefault@@} OR {#@userRole_ID@#} IS NULL),
			CompiledQuery:`eyNAbW9uZXlAI30+PXtAQFVzZXJNb25leUBAfSBBTkQgKHsjQHVzZXJSb2xlX0lELnJvbGVOYW1lQCN9PXtAQFJvbGVOYW1lRGVmYXVsdEBAfSBPUiB7I0B1c2VyUm9sZV9JREAjfSBJUyBOVUxMKQ==`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "Money":"{#@money@#}",
             "UserRoleID":"{#@userRole_ID@#}",
             "RoleName":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"UserMoney", VarType:"float64"},
									{VarName:"RoleNameDefault", VarType:"string"},
 
								},
			Tag: 			"asdax",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		8412,
			EndOff:  		8541,
			Hash:  			"2BD856B2CC369E11ECD3F9D24B3C2A29",
			SubQueries: 	null,
			},

		},
		"tst1074-W1": 
		{
			//CompiledQuery: {#@userName@#}={@@UserName@@},
			CompiledQuery:`eyNAdXNlck5hbWVAI309e0BAVXNlck5hbWVAQH0=`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "UserName":"{#@userName@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"UserName", VarType:"string"},
 
								},
			Tag: 			"tst1074",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		8700,
			EndOff:  		8758,
			Hash:  			"3FC12F87D4546693DF4D2809784E04B9",
			SubQueries: 	null,
			},

		},
		"evcy59-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true,
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZQ==`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy59",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		10148,
			EndOff:  		10215,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
			SubQueries: 	null,
			},

		},
		"evcy59-S2": 
		{
			//CompiledQuery: {#@ID@#} AS "ID", {#@userName@#} AS "UserName", {#@UUID@#} AS "UUID", {#@user_psw@#} AS "UserPsw", {#@userRole_ID@#} AS "UserRole_ID", {#@time1@#} AS "Time1", {#@money@#} AS "Money", {#@userRole_ID.roleName@#} AS "UserRole",
			CompiledQuery:`eyNASURAI30gQVMgIklEIiwgeyNAdXNlck5hbWVAI30gQVMgIlVzZXJOYW1lIiwgeyNAVVVJREAjfSBBUyAiVVVJRCIsIHsjQHVzZXJfcHN3QCN9IEFTICJVc2VyUHN3IiwgeyNAdXNlclJvbGVfSURAI30gQVMgIlVzZXJSb2xlX0lEIiwgeyNAdGltZTFAI30gQVMgIlRpbWUxIiwgeyNAbW9uZXlAI30gQVMgIk1vbmV5IiwgeyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI30gQVMgIlVzZXJSb2xlIg==`,			
			SelectSqlFields: map[string]string{
             "ID":"{#@ID@#}",
             "UserName":"{#@userName@#}",
             "UUID":"{#@UUID@#}",
             "UserPsw":"{#@user_psw@#}",
             "UserRole_ID":"{#@userRole_ID@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
             "UserRole":"{#@userRole_ID.roleName@#}",
            },			
			Fields:			map[string]string{
             "UserRole_ID":"{#@userRole_ID@#}",
             "RoleName":"{#@userRole_ID.roleName@#}",
             "ID":"{#@ID@#}",
             "UserName":"{#@userName@#}",
             "UUID":"{#@UUID@#}",
             "UserPsw":"{#@user_psw@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy59",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		10222,
			EndOff:  		10334,
			Hash:  			"84B8D21C767C75A054571BCA27A1F625",
			SubQueries: 	null,
			},

		},
		"evcy58-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true,
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZQ==`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		12046,
			EndOff:  		12115,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
			SubQueries: 	null,
			},

		},
		"evcy58-S2": 
		{
			//CompiledQuery: {#@userRole_ID.roleName@#} AS "UserRoleName", MIN( {#@time1@#} ) AS "MinTime1", SUM( {#@money@#} ) AS "SumMoney",
			CompiledQuery:`eyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI30gQVMgIlVzZXJSb2xlTmFtZSIsIE1JTiggeyNAdGltZTFAI30gKSBBUyAiTWluVGltZTEiLCBTVU0oIHsjQG1vbmV5QCN9ICkgQVMgIlN1bU1vbmV5Ig==`,			
			SelectSqlFields: map[string]string{
             "UserRoleName":"{#@userRole_ID.roleName@#}",
             "MinTime1":"MIN( {#@time1@#} )",
             "SumMoney":"SUM( {#@money@#} )",
            },			
			Fields:			map[string]string{
             "RoleName":"{#@userRole_ID.roleName@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		12128,
			EndOff:  		12342,
			Hash:  			"8B5687E0DE678F62832408F1E304EFCC",
			SubQueries: 	null,
			},

		},
		"tst253-V1": 
		{
			//CompiledQuery: {#@userRole_ID.roleName@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI30=`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "RoleName":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst253",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		13555,
			EndOff:  		13614,
			Hash:  			"C9B74D6F70CBFC75B718CFFD7208EBE3",
			SubQueries: 	null,
			},

		},
		"tst254-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst254",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		14645,
			EndOff:  		14704,
			Hash:  			"2FB1357DACB09276B8F66871D9084958",
			SubQueries: 	null,
			},

		},
		"tst254-X2": 
		{
			//CompiledQuery: {#@userRole_ID.ID@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuSURAI30=`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "ID":"{#@userRole_ID.ID@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst254",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		14724,
			EndOff:  		14785,
			Hash:  			"BACD3298080469A33137FE4B42BF083B",
			SubQueries: 	null,
			},

		},"tst1_066-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true,
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZQ==`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst1_066",
			File: 			"src\\test_where\\test1_rtm.go",
			StartOff: 		2116,
			EndOff:  		2191,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
			SubQueries: 	null,
			},

		},
		"tst1_066-S2": 
		{
			//CompiledQuery: {#@userRole_ID.roleName@#} AS "UserRoleName", MIN( {#@time1@#} ) AS "MinTime1", SUM( {#@money@#} ) AS "SumMoney",
			CompiledQuery:`eyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI30gQVMgIlVzZXJSb2xlTmFtZSIsIE1JTiggeyNAdGltZTFAI30gKSBBUyAiTWluVGltZTEiLCBTVU0oIHsjQG1vbmV5QCN9ICkgQVMgIlN1bU1vbmV5Ig==`,			
			SelectSqlFields: map[string]string{
             "UserRoleName":"{#@userRole_ID.roleName@#}",
             "MinTime1":"MIN( {#@time1@#} )",
             "SumMoney":"SUM( {#@money@#} )",
            },			
			Fields:			map[string]string{
             "RoleName":"{#@userRole_ID.roleName@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst1_066",
			File: 			"src\\test_where\\test1_rtm.go",
			StartOff: 		2210,
			EndOff:  		2440,
			Hash:  			"4AAC82C4C328EB8EB9877CEBA29E506F",
			SubQueries: 	null,
			},

		},"tst143-W1": 
		{
			//CompiledQuery: (CASE WHEN ({#@userRole_ID@#} IS NOT NULL) THEN {#@userRole_ID.isActive@#} ELSE false END),
			CompiledQuery:`KENBU0UgV0hFTiAoeyNAdXNlclJvbGVfSURAI30gSVMgTk9UIE5VTEwpIFRIRU4geyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30gRUxTRSBmYWxzZSBFTkQp`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "UserRoleID.IsActive":"{#@userRole_ID@#}",
             "IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst143",
			File: 			"src\\test_where\\test_misc.go",
			StartOff: 		976,
			EndOff:  		1093,
			Hash:  			"74372D433428301E95B05E49BBD6EED1",
			SubQueries: 	null,
			},

		},"tst1340-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true AND {#@userRole_ID.roleName@#}='admin' AND {#@UUID@#}={@@uuid@@},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZSBBTkQgeyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI309J2FkbWluJyBBTkQgeyNAVVVJREAjfT17QEB1dWlkQEB9`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "IsActive":"{#@userRole_ID.isActive@#}",
             "RoleName":"{#@userRole_ID.roleName@#}",
             "UUID":"{#@UUID@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"uuid", VarType:"string"},
 
								},
			Tag: 			"tst1340",
			File: 			"src\\test_crud\\Crud.go",
			StartOff: 		12088,
			EndOff:  		12219,
			Hash:  			"B8F7275000F6AB0B7895B045C38B915A",
			SubQueries: 	null,
			},

		},"tsql082-W2": 
		{
			//CompiledQuery: ,
			CompiledQuery:``,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"ctx", VarType:"DBContext"},
 
								},
			Tag: 			"tsql082",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		2322,
			EndOff:  		2624,
			Hash:  			"930DA7905000D1DAB21553B9F6C9A079",
			SubQueries: 	null,
			},

		},
		"-W1": 
		{
			//CompiledQuery: {#@roleName@#}={@@RoleNameDefault@@} AND {#@ID@#}={#@userRole_ID@#},
			CompiledQuery:`eyNAcm9sZU5hbWVAI309e0BAUm9sZU5hbWVEZWZhdWx0QEB9IEFORCB7I0BJREAjfT17I0B1c2VyUm9sZV9JREAjfQ==`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "RoleName":"{#@roleName@#}",
             "ID":"{#@ID@#}",
             "UserRole_ID":"{#@userRole_ID@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"RoleNameDefault", VarType:"string"},
 
								},
			Tag: 			"",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		2387,
			EndOff:  		2489,
			Hash:  			"4AA4236F4A7AF1C9B044210E8966919D",
			SubQueries: 	null,
			},

		},
		"tsql082-W1": 
		{
			//CompiledQuery: {#@money@#}>={@@UserMoney@@} AND <error 107>,
			CompiledQuery:`eyNAbW9uZXlAI30+PXtAQFVzZXJNb25leUBAfSBBTkQgPGVycm9yIDEwNz4=`,			
			SelectSqlFields: nil,			
			Fields:			map[string]string{
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"ctx", VarType:"DBContext"},
									{VarName:"UserMoney", VarType:"float64"},
 
								},
			Tag: 			"tsql082",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		4085,
			EndOff:  		4252,
			Hash:  			"C006A1F43AF479A864ECF0DE19CCEA6F",
			SubQueries: 	null,
			},

		},
	}
	
	_this.DBContextBase.ProcessCompiledQuery( &_this.DBContextBase.CompiledSqlQueries, false );
}

func (_this *DBContext) CompiledQueries_GetDialectName() string { return "postgres"; }


