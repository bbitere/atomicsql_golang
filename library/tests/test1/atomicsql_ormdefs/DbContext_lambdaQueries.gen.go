
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
             "userRole_ID":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
																	{VarName:"UserMoney", VarType:"float64"},
									{VarName:"RoleNameDefault", VarType:"string"},
 
								},
			Tag: 			"asdax",
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		8446,
			EndOff:  		8584,
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
			StartOff: 		8744,
			EndOff:  		8803,
			Hash:  			"7B5C0D4477A52364A382166ECCC60180",
		},
		"evcy59-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true,
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZQ==`,
			
			SelectSqlFields: nil,
			
			Fields:			map[string]string{
             "userRole_ID":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy59",
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		10226,
			EndOff:  		10299,
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
             "userRole_ID":"{#@userRole_ID.roleName@#}",
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
			StartOff: 		10309,
			EndOff:  		10438,
			Hash:  			"F128F38C0147DBEC1F3DC30C2BA3FD24",
		},
		"evcy58-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true,
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZQ==`,
			
			SelectSqlFields: nil,
			
			Fields:			map[string]string{
             "userRole_ID":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		12220,
			EndOff:  		12295,
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
             "userRole_ID":"{#@userRole_ID.roleName@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"library\\tests\\test1\\test_where\\test1.go",
			StartOff: 		12314,
			EndOff:  		12534,
			Hash:  			"B08DE8ED853C271C00FC1FFD127E71BC",
		},
	}
	
	_this.DBContextBase.ProcessCompiledQuery( &_this.CompiledSqlQueries, false );
}


