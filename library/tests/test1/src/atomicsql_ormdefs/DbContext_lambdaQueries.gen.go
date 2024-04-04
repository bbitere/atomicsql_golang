
/* this class is generated automatically by compiler AtomicSql goscanner.exe exporter*/

package atomicsql_sqldefs
	import (
		orm	 "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	)

func (_this *DBContext) LoadCompiledQuery(){

	_this.DBContextBase.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
		"asdax-W1": 
		{
			//CompiledQuery: {#@money\"         bson:\"money@#}>={@@UserMoney@@} AND ({#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}={@@RoleNameDefault@@} OR {#@-\"             bson:\"userRole_ID@#} IS NULL),
			CompiledQuery:`eyNAbW9uZXlcIiAgICAgICAgIGJzb246XCJtb25leUAjfT49e0BAVXNlck1vbmV5QEB9IEFORCAoeyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5yb2xlTmFtZVwiICAgICAgYnNvbjpcInJvbGVOYW1lQCN9PXtAQFJvbGVOYW1lRGVmYXVsdEBAfSBPUiB7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lEQCN9IElTIE5VTEwp`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "Money", "UserRoleID", "UserRoleID.RoleName" },
			Fields:			map[string]string{
             "Money":"{#@money\"         bson:\"money@#}",
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
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
			NosqlQuery: 	nil,
		},
		"tst1074-W1": 
		{
			//CompiledQuery: {#@userName\"      bson:\"userName@#}={@@UserName@@},
			CompiledQuery:`eyNAdXNlck5hbWVcIiAgICAgIGJzb246XCJ1c2VyTmFtZUAjfT17QEBVc2VyTmFtZUBAfQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserName" },
			Fields:			map[string]string{
             "UserName":"{#@userName\"      bson:\"userName@#}",
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
			NosqlQuery: 	nil,
		},
		"evcy59-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},
		"evcy59-S2": 
		{
			//CompiledQuery: {#@-\"             bson:\"_id@#} AS "NoSqlID", {#@ID\"            bson:\"-@#} AS "ID", {#@userName\"      bson:\"userName@#} AS "UserName", {#@UUID\"          bson:\"UUID@#} AS "UUID", {#@user_psw\"      bson:\"user_psw@#} AS "UserPsw", {#@userRole_ID\"   bson:\"-@#} AS "UserRole_ID", {#@time1\"         bson:\"time1@#} AS "Time1", {#@money\"         bson:\"money@#} AS "Money", {#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#} AS "UserRole",
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJfaWRAI30gQVMgIk5vU3FsSUQiLCB7I0BJRFwiICAgICAgICAgICAgYnNvbjpcIi1AI30gQVMgIklEIiwgeyNAdXNlck5hbWVcIiAgICAgIGJzb246XCJ1c2VyTmFtZUAjfSBBUyAiVXNlck5hbWUiLCB7I0BVVUlEXCIgICAgICAgICAgYnNvbjpcIlVVSURAI30gQVMgIlVVSUQiLCB7I0B1c2VyX3Bzd1wiICAgICAgYnNvbjpcInVzZXJfcHN3QCN9IEFTICJVc2VyUHN3IiwgeyNAdXNlclJvbGVfSURcIiAgIGJzb246XCItQCN9IEFTICJVc2VyUm9sZV9JRCIsIHsjQHRpbWUxXCIgICAgICAgICBic29uOlwidGltZTFAI30gQVMgIlRpbWUxIiwgeyNAbW9uZXlcIiAgICAgICAgIGJzb246XCJtb25leUAjfSBBUyAiTW9uZXkiLCB7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELnJvbGVOYW1lXCIgICAgICBic29uOlwicm9sZU5hbWVAI30gQVMgIlVzZXJSb2xlIg==`,			
			SelectSqlFields: map[string]string{
             "NoSqlID":"{#@-\"             bson:\"_id@#}",
             "ID":"{#@ID\"            bson:\"-@#}",
             "UserName":"{#@userName\"      bson:\"userName@#}",
             "UUID":"{#@UUID\"          bson:\"UUID@#}",
             "UserPsw":"{#@user_psw\"      bson:\"user_psw@#}",
             "UserRole_ID":"{#@userRole_ID\"   bson:\"-@#}",
             "Time1":"{#@time1\"         bson:\"time1@#}",
             "Money":"{#@money\"         bson:\"money@#}",
             "UserRole":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
            },			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName", "NoSqlID", "ID", "UserName", "UUID", "UserPsw", "UserRole_ID", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "NoSqlID":"{#@-\"             bson:\"_id@#}",
             "ID":"{#@ID\"            bson:\"-@#}",
             "UserName":"{#@userName\"      bson:\"userName@#}",
             "UUID":"{#@UUID\"          bson:\"UUID@#}",
             "UserPsw":"{#@user_psw\"      bson:\"user_psw@#}",
             "UserRole_ID":"{#@userRole_ID\"   bson:\"-@#}",
             "Time1":"{#@time1\"         bson:\"time1@#}",
             "Money":"{#@money\"         bson:\"money@#}",
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
			NosqlQuery: 	nil,
		},
		"evcy58-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},
		"evcy58-S2": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#} AS "UserRoleName", MIN( {#@time1\"         bson:\"time1@#} ) AS "MinTime1", SUM( {#@money\"         bson:\"money@#} ) AS "SumMoney",
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5yb2xlTmFtZVwiICAgICAgYnNvbjpcInJvbGVOYW1lQCN9IEFTICJVc2VyUm9sZU5hbWUiLCBNSU4oIHsjQHRpbWUxXCIgICAgICAgICBic29uOlwidGltZTFAI30gKSBBUyAiTWluVGltZTEiLCBTVU0oIHsjQG1vbmV5XCIgICAgICAgICBic29uOlwibW9uZXlAI30gKSBBUyAiU3VtTW9uZXki`,			
			SelectSqlFields: map[string]string{
             "UserRoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "MinTime1":"MIN( {#@time1\"         bson:\"time1@#} )",
             "SumMoney":"SUM( {#@money\"         bson:\"money@#} )",
            },			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "Time1":"{#@time1\"         bson:\"time1@#}",
             "Money":"{#@money\"         bson:\"money@#}",
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
			NosqlQuery: 	nil,
		},
		"tst253-V1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5yb2xlTmFtZVwiICAgICAgYnNvbjpcInJvbGVOYW1lQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
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
			NosqlQuery: 	nil,
		},
		"tst254-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},
		"tst254-X2": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.ID\"            bson:\"-@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5JRFwiICAgICAgICAgICAgYnNvbjpcIi1AI30=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.ID" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.ID":"{#@-\"             bson:\"userRole_ID.ID\"            bson:\"-@#}",
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
			NosqlQuery: 	nil,
		},
		"tst665-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},"tst1_066-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},
		"tst1_066-S2": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#} AS "UserRoleName", MIN( {#@time1\"         bson:\"time1@#} ) AS "MinTime1", SUM( {#@money\"         bson:\"money@#} ) AS "SumMoney",
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5yb2xlTmFtZVwiICAgICAgYnNvbjpcInJvbGVOYW1lQCN9IEFTICJVc2VyUm9sZU5hbWUiLCBNSU4oIHsjQHRpbWUxXCIgICAgICAgICBic29uOlwidGltZTFAI30gKSBBUyAiTWluVGltZTEiLCBTVU0oIHsjQG1vbmV5XCIgICAgICAgICBic29uOlwibW9uZXlAI30gKSBBUyAiU3VtTW9uZXki`,			
			SelectSqlFields: map[string]string{
             "UserRoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "MinTime1":"MIN( {#@time1\"         bson:\"time1@#} )",
             "SumMoney":"SUM( {#@money\"         bson:\"money@#} )",
            },			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleName", "Time1", "Money" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "Time1":"{#@time1\"         bson:\"time1@#}",
             "Money":"{#@money\"         bson:\"money@#}",
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
			NosqlQuery: 	nil,
		},"tst143-W1": 
		{
			//CompiledQuery: (CASE WHEN ({#@-\"             bson:\"userRole_ID@#} IS NOT NULL) THEN {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#} ELSE false END),
			CompiledQuery:`KENBU0UgV0hFTiAoeyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JREAjfSBJUyBOT1QgTlVMTCkgVEhFTiB7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELmlzQWN0aXZlXCIgICAgICBic29uOlwiaXNBY3RpdmVAI30gRUxTRSBmYWxzZSBFTkQp`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
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
			NosqlQuery: 	nil,
		},"tst1340-W1": 
		{
			//CompiledQuery: {#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}=true AND {#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}='admin' AND {#@UUID\"          bson:\"UUID@#}={@@uuid@@},
			CompiledQuery:`eyNALVwiICAgICAgICAgICAgIGJzb246XCJ1c2VyUm9sZV9JRC5pc0FjdGl2ZVwiICAgICAgYnNvbjpcImlzQWN0aXZlQCN9PXRydWUgQU5EIHsjQC1cIiAgICAgICAgICAgICBic29uOlwidXNlclJvbGVfSUQucm9sZU5hbWVcIiAgICAgIGJzb246XCJyb2xlTmFtZUAjfT0nYWRtaW4nIEFORCB7I0BVVUlEXCIgICAgICAgICAgYnNvbjpcIlVVSURAI309e0BAdXVpZEBAfQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive", "UserRoleID.RoleName", "UUID" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.IsActive":"{#@-\"             bson:\"userRole_ID.isActive\"      bson:\"isActive@#}",
             "UserRoleID.RoleName":"{#@-\"             bson:\"userRole_ID.roleName\"      bson:\"roleName@#}",
             "UUID":"{#@UUID\"          bson:\"UUID@#}",
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
			NosqlQuery: 	nil,
		},"tsql082a-W1": 
		{
			//CompiledQuery: (({#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole@#} IS NULL) OR {#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole.ID@#} > 0) AND {#@money\"         bson:\"money@#}>={@@UserMoney@@} AND {#@ID\"            bson:\"-@#}<>123 AND {#@userName\"      bson:\"userName@#}<>E'a''"\n mm',
			CompiledQuery:`KCh7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELi1cIiAgICAgICAgICAgICBic29uOlwiU3RhdHVzcm9sZUAjfSBJUyBOVUxMKSBPUiB7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELi1cIiAgICAgICAgICAgICBic29uOlwiU3RhdHVzcm9sZS5JREAjfSA+IDApIEFORCB7I0Btb25leVwiICAgICAgICAgYnNvbjpcIm1vbmV5QCN9Pj17QEBVc2VyTW9uZXlAQH0gQU5EIHsjQElEXCIgICAgICAgICAgICBic29uOlwiLUAjfTw+MTIzIEFORCB7I0B1c2VyTmFtZVwiICAgICAgYnNvbjpcInVzZXJOYW1lQCN9PD5FJ2EnJyJcbiBtbSc=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleStatusID", "UserRoleID.RoleStatusID.ID", "Money", "ID", "UserName" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleStatusID":"{#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole@#}",
             "UserRoleID.RoleStatusID.ID":"{#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole.ID@#}",
             "Money":"{#@money\"         bson:\"money@#}",
             "ID":"{#@ID\"            bson:\"-@#}",
             "UserName":"{#@userName\"      bson:\"userName@#}",
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
			NosqlQuery: 	nil,
		},
		"tsql082-Q1": 
		{
			//CompiledQuery: (({#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole@#} IS NULL) OR {#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole.ID@#} > 0) AND {#@money\"         bson:\"money@#}>={@@UserMoney@@} AND {#@userRole_ID\"   bson:\"-@#} IN ( {@$ids$@} ) AND {@@bActive@@},
			CompiledQuery:`KCh7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELi1cIiAgICAgICAgICAgICBic29uOlwiU3RhdHVzcm9sZUAjfSBJUyBOVUxMKSBPUiB7I0AtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELi1cIiAgICAgICAgICAgICBic29uOlwiU3RhdHVzcm9sZS5JREAjfSA+IDApIEFORCB7I0Btb25leVwiICAgICAgICAgYnNvbjpcIm1vbmV5QCN9Pj17QEBVc2VyTW9uZXlAQH0gQU5EIHsjQHVzZXJSb2xlX0lEXCIgICBic29uOlwiLUAjfSBJTiAoIHtAJGlkcyRAfSApIEFORCB7QEBiQWN0aXZlQEB9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.RoleStatusID", "UserRoleID.RoleStatusID.ID", "Money", "UserRole_ID" },
			Fields:			map[string]string{
             "UserRoleID":"{#@-\"             bson:\"userRole_ID@#}",
             "UserRoleID.RoleStatusID":"{#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole@#}",
             "UserRoleID.RoleStatusID.ID":"{#@-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole.ID@#}",
             "Money":"{#@money\"         bson:\"money@#}",
             "UserRole_ID":"{#@userRole_ID\"   bson:\"-@#}",
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
			NosqlQuery: 	nil,
		},
		"tsql082.ids-W1": 
		{
			//CompiledQuery: {#@roleName\"      bson:\"roleName@#}={@@RoleNameDefault@@} AND ((NOT (({#@statusrole_ID\" bson:\"-@#} IS NOT NULL))) OR {#@statusrole_ID\" bson:\"-@#}={#$-\"             bson:\"userRole_ID.-\"             bson:\"Statusrole.ID$#}),
			CompiledQuery:`eyNAcm9sZU5hbWVcIiAgICAgIGJzb246XCJyb2xlTmFtZUAjfT17QEBSb2xlTmFtZURlZmF1bHRAQH0gQU5EICgoTk9UICgoeyNAc3RhdHVzcm9sZV9JRFwiIGJzb246XCItQCN9IElTIE5PVCBOVUxMKSkpIE9SIHsjQHN0YXR1c3JvbGVfSURcIiBic29uOlwiLUAjfT17IyQtXCIgICAgICAgICAgICAgYnNvbjpcInVzZXJSb2xlX0lELi1cIiAgICAgICAgICAgICBic29uOlwiU3RhdHVzcm9sZS5JRCQjfSk=`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "RoleName", "Role_status_ID" },
			Fields:			map[string]string{
             "RoleName":"{#@roleName\"      bson:\"roleName@#}",
             "Role_status_ID":"{#@statusrole_ID\" bson:\"-@#}",
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
			NosqlQuery: 	nil,
		},
		"tsql147-Q1": 
		{
			//CompiledQuery: {#@money\"         bson:\"money@#}>={@@UserMoney@@} AND {@$cnt$@} > 0 AND {@@bActive@@},
			CompiledQuery:`eyNAbW9uZXlcIiAgICAgICAgIGJzb246XCJtb25leUAjfT49e0BAVXNlck1vbmV5QEB9IEFORCB7QCRjbnQkQH0gPiAwIEFORCB7QEBiQWN0aXZlQEB9`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "UserRole_ID", "Money" },
			Fields:			map[string]string{
             "UserRole_ID":"{#@userRole_ID\"   bson:\"-@#}",
             "Money":"{#@money\"         bson:\"money@#}",
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
			NosqlQuery: 	nil,
		},
		"tsql147.cnt-W1": 
		{
			//CompiledQuery: {#@ID\"            bson:\"-@#}={#$userRole_ID\"   bson:\"-$#},
			CompiledQuery:`eyNASURcIiAgICAgICAgICAgIGJzb246XCItQCN9PXsjJHVzZXJSb2xlX0lEXCIgICBic29uOlwiLSQjfQ==`,			
			SelectSqlFields: nil,			
			OrderedFields:	[]string{ "ID" },
			Fields:			map[string]string{
             "ID":"{#@ID\"            bson:\"-@#}",
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
			NosqlQuery: 	nil,
		},
	}
	
	_this.DBContextBase.ProcessCompiledQuery( &_this.DBContextBase.CompiledSqlQueries, false );
}

func (_this *DBContext) CompiledQueries_GetDialectName() string { return "postgres"; }


