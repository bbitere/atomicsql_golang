/* this class is generated automatically by compiler AtomicSql goscanner.exe exporter*/

package atomicsql_sqldefs

import (
	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

func (_this *DBContextNSql) LoadCompiledQuery(){

	_this.DBContextBaseNoSql.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
		"asdax-W1": 
		{
			//CompiledQuery: {#@money@#}>={@@UserMoney@@} AND ({#@userRole_ID.roleName@#}={@@RoleNameDefault@@} OR {#@userRole_ID@#} IS NULL),
			CompiledQuery:`eyNAbW9uZXlAI30+PXtAQFVzZXJNb25leUBAfSBBTkQgKHsjQHVzZXJSb2xlX0lELnJvbGVOYW1lQCN9PXtAQFJvbGVOYW1lRGVmYXVsdEBAfSBPUiB7I0B1c2VyUm9sZV9JREAjfSBJUyBOVUxMKQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "Money", "UserRoleID", "UserRoleID.RoleName" },
			Fields:			map[string]string{
             "Money":"{#@money@#}",
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"UserMoney", VarType:"float64"},
									{VarName:"RoleNameDefault", VarType:"string"}, 
								},
			Tag: 			"asdax",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		20,
			EndOff:  		149,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Nb25leSA+PSBVc2VyTW9uZXkgJiYNCgkJCSh4LlVzZXJSb2xlSUQuUm9sZU5hbWUgPT0gUm9sZU5hbWVEZWZhdWx0IHx8IHguVXNlclJvbGVJRCA9PSBuaWwpDQoJ",
			IsQryS:			false,
			SubQueries: 	nil,			
		},		 
		"ns-asdax-W1": 
		{
			//CompiledQuery: {#@money@#}>={@@UserMoney@@} AND ({#@userRole_ID.roleName@#}={@@RoleNameDefault@@} OR {#@userRole_ID@#} IS NULL),
			CompiledQuery:`eyNAbW9uZXlAI30+PXtAQFVzZXJNb25leUBAfSBBTkQgKHsjQHVzZXJSb2xlX0lELnJvbGVOYW1lQCN9PXtAQFJvbGVOYW1lRGVmYXVsdEBAfSBPUiB7I0B1c2VyUm9sZV9JREAjfSBJUyBOVUxMKQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "Money", "UserRoleID", "UserRoleID.RoleName" },
			Fields:			map[string]string{
             "Money":"{#@money@#}",
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"UserMoney", VarType:"float64"},
									{VarName:"RoleNameDefault", VarType:"string"}, 
								},
			Tag: 			"asdax",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		20,
			EndOff:  		149,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Nb25leSA+PSBVc2VyTW9uZXkgJiYNCgkJCSh4LlVzZXJSb2xlSUQuUm9sZU5hbWUgPT0gUm9sZU5hbWVEZWZhdWx0IHx8IHguVXNlclJvbGVJRCA9PSBuaWwpDQoJ",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery:		[]any{
				
				[]any{
					">=", "@#Money#@", "@@UserMoney@@",
				},
				[]any{
					"||",
					[]any{
						"==", "@#UserRoleID.RoleName#@", "@@RoleNameDefault@@",
					},
					[]any{
						"==", "@#UserRoleID#@", nil,
					},
				},
			},	
		},
		"tst1074-W1": 
		{
			//CompiledQuery: {#@userName@#}={@@UserName@@},
			CompiledQuery:`eyNAdXNlck5hbWVAI309e0BAVXNlck5hbWVAQH0=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserName" },
			Fields:			map[string]string{
             "UserName":"{#@userName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"UserName", VarType:"string"}, 
								},
			Tag: 			"tst1074",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		22,
			EndOff:  		80,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Vc2VyTmFtZSA9PSBVc2VyTmFtZQ0KCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"ns-tst1074-W1": 
		{			
			CompiledQuery:``,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserName" },
			Fields:			map[string]string{
             "UserName":"{#@userName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"UserName", VarType:"string"}, 
								},
			Tag: 			"ns-tst1074",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		22,
			EndOff:  		80,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Vc2VyTmFtZSA9PSBVc2VyTmFtZQ0KCQ==",
			IsQryS:			false,
			SubQueries: 	nil,	
			NosqlQuery:		[]any{
				
				[]any{
					"==", "@#UserName#@", "@@UserName@@",
				},				
			},		
		},

		
		"evcy59-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy59",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,			
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
			OrderedFields:	[]string{ "UserRole_ID", "UserRoleID.RoleName", "ID", "UserName", "UUID", "UserPsw", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRole_ID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
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
			StartOff: 		91,
			EndOff:  		203,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpICp2VXNlcjEgew0KCQkJcmV0dXJuICZ2VXNlcjF7DQoJCQkJVXNlcjogICAgICp4LA0KCQkJCVVzZXJSb2xlOiB4LlVzZXJSb2xlSUQuUm9sZU5hbWUsDQoJCQl9DQoJCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"evcy58-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		26,
			EndOff:  		87,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJCXJldHVybiB4LlVzZXJSb2xlSUQuSXNBY3RpdmUNCgkJCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
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
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy58",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		100,
			EndOff:  		314,
			Hash:  			"ZnVuYyh4ICpUVXNlckFnZ3IpICpUVXNlclZpZXcgew0KDQoJCQlyZXR1cm4gJlRVc2VyVmlld3sNCgkJCQlVc2VyUm9sZU5hbWU6IHguVXNlclJvbGVJRC5Sb2xlTmFtZSwNCgkJCQlNaW5UaW1lMTogICAgIGF0bXNxbF9mdW5jLlNxbF9NaW5EYXRlTih4LlRpbWUxKSwNCgkJCQlTdW1Nb25leTogICAgIGF0bXNxbF9mdW5jLlNxbF9TdW1GNjQoeC5Nb25leSksDQoJCQl9DQoJCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"tst253-V1": 
		{
			//CompiledQuery: {#@userRole_ID.roleName@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst253",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		30,
			EndOff:  		89,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIHN0cmluZyB7DQoJCXJldHVybiB4LlVzZXJSb2xlSUQuUm9sZU5hbWUNCgk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"tst254-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst254",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"tst254-X2": 
		{
			//CompiledQuery: {#@userRole_ID.ID@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuSURAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.ID" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.ID":"{#@userRole_ID.ID@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst254",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		104,
			EndOff:  		165,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGludDY0IHsNCgkJCXJldHVybiBpbnQ2NCh4LlVzZXJSb2xlSUQuSUQpDQoJCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"tst665-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst665",
			File: 			"src\\test_where\\test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},"tst1_066-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst1_066",
			File: 			"src\\test_where\\test1_rtm.go",
			StartOff: 		72,
			EndOff:  		139,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJCQkJCXJldHVybiB4LlVzZXJSb2xlSUQuSXNBY3RpdmUNCgkJCQkJCQ==",
			IsQryS:			false,
			SubQueries: 	nil,			
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
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
             "Time1":"{#@time1@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst1_066",
			File: 			"src\\test_where\\test1_rtm.go",
			StartOff: 		158,
			EndOff:  		388,
			Hash:  			"ZnVuYyAoeCAqVFVzZXJBZ2dyICkgKlRVc2VyVmlldyB7DQoNCgkJCQkJCXJldHVybiAmVFVzZXJWaWV3ew0KCQkJCQkJCVVzZXJSb2xlTmFtZTogeC5Vc2VyUm9sZUlELlJvbGVOYW1lLA0KCQkJCQkJCU1pblRpbWUxOiBhdG1zcWxfZnVuYy5TcWxfTWluRGF0ZU4oIHguVGltZTEgKSwNCgkJCQkJCQlTdW1Nb25leTogYXRtc3FsX2Z1bmMuU3FsX1N1bUY2NCggeC5Nb25leSApLA0KCQkJCQkJfQ0KCQkJCQk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},"tst143-W1": 
		{
			//CompiledQuery: (CASE WHEN ({#@userRole_ID@#} IS NOT NULL) THEN {#@userRole_ID.isActive@#} ELSE false END),
			CompiledQuery:`KENBU0UgV0hFTiAoeyNAdXNlclJvbGVfSURAI30gSVMgTk9UIE5VTEwpIFRIRU4geyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI30gRUxTRSBmYWxzZSBFTkQp`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst143",
			File: 			"src\\test_where\\test_misc.go",
			StartOff: 		22,
			EndOff:  		139,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2x7DQoJCQkJCQkJcmV0dXJuIGF0bXNxbF9mdW5jLlNxbF9JSUYoIHguVXNlclJvbGVJRCAhPSBuaWwsIHguVXNlclJvbGVJRC5Jc0FjdGl2ZSwgZmFsc2UpOw0KCQkJCQkJ",
			IsQryS:			false,
			SubQueries: 	nil,			
		},"tst1340-W1": 
		{
			//CompiledQuery: {#@userRole_ID.isActive@#}=true AND {#@userRole_ID.roleName@#}='admin' AND {#@UUID@#}={@@uuid@@},
			CompiledQuery:`eyNAdXNlclJvbGVfSUQuaXNBY3RpdmVAI309dHJ1ZSBBTkQgeyNAdXNlclJvbGVfSUQucm9sZU5hbWVAI309J2FkbWluJyBBTkQgeyNAVVVJREAjfT17QEB1dWlkQEB9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive", "UserRoleID.RoleName", "UUID" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
             "UserRoleID.RoleName":"{#@userRole_ID.roleName@#}",
             "UUID":"{#@UUID@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"uuid", VarType:"string"}, 
								},
			Tag: 			"tst1340",
			File: 			"src\\test_crud\\Crud.go",
			StartOff: 		26,
			EndOff:  		157,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZSA9PSB0cnVlICYmDQoJCQkJeC5Vc2VyUm9sZUlELlJvbGVOYW1lID09ICJhZG1pbiIgJiYNCgkJCQl4LlVVSUQgPT0gdXVpZA0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},"tsql082a-W1": 
		{
			//CompiledQuery: (({#@userRole_ID.role_status_ID@#} IS NULL) OR {#@userRole_ID.role_status_ID.ID@#} > 0) AND {#@money@#}>={@@UserMoney@@} AND {#@ID@#}<>123 AND {#@userName@#}<>E'a''"\n mm',
			CompiledQuery:`KCh7I0B1c2VyUm9sZV9JRC5yb2xlX3N0YXR1c19JREAjfSBJUyBOVUxMKSBPUiB7I0B1c2VyUm9sZV9JRC5yb2xlX3N0YXR1c19JRC5JREAjfSA+IDApIEFORCB7I0Btb25leUAjfT49e0BAVXNlck1vbmV5QEB9IEFORCB7I0BJREAjfTw+MTIzIEFORCB7I0B1c2VyTmFtZUAjfTw+RSdhJyciXG4gbW0n`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleStatusID", "UserRoleID.RoleStatusID.ID", "Money", "ID", "UserName" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.RoleStatusID":"{#@userRole_ID.role_status_ID@#}",
             "UserRoleID.RoleStatusID.ID":"{#@userRole_ID.role_status_ID.ID@#}",
             "Money":"{#@money@#}",
             "ID":"{#@ID@#}",
             "UserName":"{#@userName@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"UserMoney", VarType:"float64"}, 
								},
			Tag: 			"tsql082a",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		23,
			EndOff:  		226,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KDQoJCXJldHVybiAoKHguVXNlclJvbGVJRC5Sb2xlU3RhdHVzSUQgPT0gbmlsKSB8fCB4LlVzZXJSb2xlSUQuUm9sZVN0YXR1c0lELklEID4gMCkgJiYNCgkJCXguTW9uZXkgPj0gVXNlck1vbmV5ICYmDQoJCQl4LklEICE9IEVDb25zdF9FSUQgJiYgDQoJCQl4LlVzZXJOYW1lICE9IEVDb25zdF9OYW1lCQkJCSANCgk=",
			IsQryS:			false,
			SubQueries: 	nil,			
		},
		"tsql082-Q1": 
		{
			//CompiledQuery: (({#@userRole_ID.role_status_ID@#} IS NULL) OR {#@userRole_ID.role_status_ID.ID@#} > 0) AND {#@money@#}>={@@UserMoney@@} AND {#@userRole_ID@#} IN ( {@$ids$@} ) AND {@@bActive@@},
			CompiledQuery:`KCh7I0B1c2VyUm9sZV9JRC5yb2xlX3N0YXR1c19JREAjfSBJUyBOVUxMKSBPUiB7I0B1c2VyUm9sZV9JRC5yb2xlX3N0YXR1c19JRC5JREAjfSA+IDApIEFORCB7I0Btb25leUAjfT49e0BAVXNlck1vbmV5QEB9IEFORCB7I0B1c2VyUm9sZV9JREAjfSBJTiAoIHtAJGlkcyRAfSApIEFORCB7QEBiQWN0aXZlQEB9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRole_ID", "UserRoleID.RoleStatusID", "UserRoleID.RoleStatusID.ID", "Money" },
			Fields:			map[string]string{
             "UserRole_ID":"{#@userRole_ID@#}",
             "UserRoleID.RoleStatusID":"{#@userRole_ID.role_status_ID@#}",
             "UserRoleID.RoleStatusID.ID":"{#@userRole_ID.role_status_ID.ID@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"ctx", VarType:"DBContext"},
									{VarName:"UserMoney", VarType:"float64"},
									{VarName:"bActive", VarType:"bool"}, 
								},
			Tag: 			"tsql082",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		26,
			EndOff:  		585,
			Hash:  			"ZnVuYyh4ICptLlVzZXIsIHEgYXRtc3FsLklEQlF1ZXJ5KSBib29sIHsNCg0KCQlpZHMsIF8gOj0gY3R4LlVzZXJSb2xlLlFyeVMoImlkcyIsIHEpLldoZXJlKGZ1bmMoeSAqbS5Vc2VyUm9sZSkgYm9vbCB7DQoJCQlyZXR1cm4geS5Sb2xlTmFtZSA9PSBSb2xlTmFtZURlZmF1bHQgJiYNCgkJCQkgICAvL3kuSUQgIT0gRUNvbnN0RUlEMSAmJg0KCQkJCSAgIC8veS5JRCAhPSBFQ29uc3RfRUlEICYmDQoJCQkJKCgheS5Sb2xlX3N0YXR1c19JRC5WYWxpZCkgfHwNCgkJCQkJeS5Sb2xlX3N0YXR1c19JRC5JbnQzMiA9PSB4LlVzZXJSb2xlSUQuUm9sZVN0YXR1c0lELklEKQ0KCQl9KS5HZXRSb3dzQXNGaWVsZEludChjdHguVXNlclJvbGVfLklEKQ0KDQoJCXJldHVybiAoKHguVXNlclJvbGVJRC5Sb2xlU3RhdHVzSUQgPT0gbmlsKSB8fCB4LlVzZXJSb2xlSUQuUm9sZVN0YXR1c0lELklEID4gMCkgJiYNCgkJCXguTW9uZXkgPj0gVXNlck1vbmV5ICYmDQoJCQlhdG1mLlNxbF9BcnJheUNvbnRhaW4oaWRzLCBpbnQ2NCh4LlVzZXJSb2xlX0lELkludDMyKSkgJiYNCg0KCQkJYkFjdGl2ZQ0KCQ==",
			IsQryS:			false,
			SubQueries: 	[]orm.TSubQuery{
								{ 
					VariableName : "ids",
				},

			},			
		},
		"tsql082.ids-W1": 
		{
			//CompiledQuery: {#@roleName@#}={@@RoleNameDefault@@} AND ((NOT (({#@role_status_ID@#} IS NOT NULL))) OR {#@role_status_ID@#}={#$userRole_ID.role_status_ID.ID$#}),
			CompiledQuery:`eyNAcm9sZU5hbWVAI309e0BAUm9sZU5hbWVEZWZhdWx0QEB9IEFORCAoKE5PVCAoKHsjQHJvbGVfc3RhdHVzX0lEQCN9IElTIE5PVCBOVUxMKSkpIE9SIHsjQHJvbGVfc3RhdHVzX0lEQCN9PXsjJHVzZXJSb2xlX0lELnJvbGVfc3RhdHVzX0lELklEJCN9KQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "RoleName", "Role_status_ID" },
			Fields:			map[string]string{
             "RoleName":"{#@roleName@#}",
             "Role_status_ID":"{#@role_status_ID@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"RoleNameDefault", VarType:"string"}, 
								},
			Tag: 			"tsql082.ids",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		22,
			EndOff:  		257,
			Hash:  			"ZnVuYyh5ICptLlVzZXJSb2xlKSBib29sIHsNCgkJCXJldHVybiB5LlJvbGVOYW1lID09IFJvbGVOYW1lRGVmYXVsdCAmJg0KCQkJCSAgIC8veS5JRCAhPSBFQ29uc3RFSUQxICYmDQoJCQkJICAgLy95LklEICE9IEVDb25zdF9FSUQgJiYNCgkJCQkoKCF5LlJvbGVfc3RhdHVzX0lELlZhbGlkKSB8fA0KCQkJCQl5LlJvbGVfc3RhdHVzX0lELkludDMyID09IHguVXNlclJvbGVJRC5Sb2xlU3RhdHVzSUQuSUQpDQoJCQ==",
			IsQryS:			true,
			SubQueries: 	nil,			
		},
		"tsql147-Q1": 
		{
			//CompiledQuery: {#@money@#}>={@@UserMoney@@} AND {@$cnt$@} > 0 AND {@@bActive@@},
			CompiledQuery:`eyNAbW9uZXlAI30+PXtAQFVzZXJNb25leUBAfSBBTkQge0AkY250JEB9ID4gMCBBTkQge0BAYkFjdGl2ZUBAfQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRole_ID", "Money" },
			Fields:			map[string]string{
             "UserRole_ID":"{#@userRole_ID@#}",
             "Money":"{#@money@#}",
            },
			ExternVar:			[]orm.TExternVar{
								{VarName:"ctx", VarType:"DBContext"},
									{VarName:"UserMoney", VarType:"float64"},
									{VarName:"bActive", VarType:"bool"}, 
								},
			Tag: 			"tsql147",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		26,
			EndOff:  		371,
			Hash:  			"ZnVuYyh4ICptLlVzZXIsIHEgYXRtc3FsLklEQlF1ZXJ5KSBib29sIHsNCg0KCQkvL3ZhciBjbnQsIF8gPSBjdHguVXNlclJvbGUuUXJ5UygiY250IiwgcSkuV2hlcmVFcSggY3R4LlVzZXJSb2xlXy5JRCwgeC5Vc2VyUm9sZV9JRC5JbnQzMiApLkdldENvdW50KCkNCgkJdmFyIGNudCwgXyA9IGN0eC5Vc2VyUm9sZS5RcnlTKCJjbnQiLCBxKS4NCgkJCVdoZXJlKGZ1bmMoeSAqbS5Vc2VyUm9sZSkgYm9vbCB7DQoJCQkJcmV0dXJuIHkuSUQgPT0geC5Vc2VyUm9sZV9JRC5JbnQzMg0KCQkJfSkuR2V0Q291bnQoKQ0KCQlyZXR1cm4geC5Nb25leSA+PSBVc2VyTW9uZXkgJiYgY250ID4gMCAmJiBiQWN0aXZlDQoJ",
			IsQryS:			false,
			SubQueries: 	[]orm.TSubQuery{
								{ 
					VariableName : "cnt",
				},

			},			
		},
		"tsql147.cnt-W1": 
		{
			//CompiledQuery: {#@ID@#}={#$userRole_ID$#},
			CompiledQuery:`eyNASURAI309eyMkdXNlclJvbGVfSUQkI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "ID" },
			Fields:			map[string]string{
             "ID":"{#@ID@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tsql147.cnt",
			File: 			"src\\test_subquery\\test1.go",
			StartOff: 		27,
			EndOff:  		98,
			Hash:  			"ZnVuYyh5ICptLlVzZXJSb2xlKSBib29sIHsNCgkJCQlyZXR1cm4geS5JRCA9PSB4LlVzZXJSb2xlX0lELkludDMyDQoJCQk=",
			IsQryS:			true,
			SubQueries: 	nil,			
		},
	}
	
	_this.DBContextBaseNoSql.ProcessCompiledQuery( &_this.DBContextBaseNoSql.CompiledSqlQueries, false );
}

func (_this *DBContextNSql) CompiledQueries_GetDialectName() string { return "postgres"; }


