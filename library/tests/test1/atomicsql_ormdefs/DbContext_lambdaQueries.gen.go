
/* this class is generated automatically by compiler AtomicSql.exe exporter*/

package atomicsql_sqldefs
	import (
		orm	 "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	)

func (_this *DBContext) LoadCompiledQuery(){

	_this.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		8577,
			EndOff:  		8715,
			Hash:  			"BA78C254E0BC9C69C4D363F967614D10",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		8875,
			EndOff:  		8934,
			Hash:  			"7B5C0D4477A52364A382166ECCC60180",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		10359,
			EndOff:  		10432,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		10442,
			EndOff:  		10571,
			Hash:  			"F128F38C0147DBEC1F3DC30C2BA3FD24",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		12377,
			EndOff:  		12452,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		12471,
			EndOff:  		12701,
			Hash:  			"4AAC82C4C328EB8EB9877CEBA29E506F",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		13945,
			EndOff:  		14009,
			Hash:  			"AD142E307B794F8F795DBC178403F9A0",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		15073,
			EndOff:  		15141,
			Hash:  			"22BAFB9DF4BEC951F91D82A3D6DD75F6",
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
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		15166,
			EndOff:  		15230,
			Hash:  			"D7206B40862AE8D64139BE634B6424BA",
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
			File: 			"library\\tests\\test1\\test_where\\test1_rtm.go",
			StartOff: 		2104,
			EndOff:  		2179,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
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
			File: 			"library\\tests\\test1\\test_where\\test1_rtm.go",
			StartOff: 		2198,
			EndOff:  		2428,
			Hash:  			"4AAC82C4C328EB8EB9877CEBA29E506F",
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
			File: 			"library\\tests\\test1\\test_where\\test_misc.go",
			StartOff: 		972,
			EndOff:  		1089,
			Hash:  			"74372D433428301E95B05E49BBD6EED1",
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
			File: 			"library\\tests\\test1\\test_crud\\Crud.go",
			StartOff: 		12179,
			EndOff:  		12324,
			Hash:  			"059EAF721518868C8166B9607E3B28A4",
		},
	}
	
	_this.DBContextBase.ProcessCompiledQuery( &_this.CompiledSqlQueries, false );
}


