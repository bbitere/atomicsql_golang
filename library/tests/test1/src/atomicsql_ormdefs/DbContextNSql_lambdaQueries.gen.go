
/* this class is generated automatically by compiler AtomicSql goscanner.exe exporter*/

package atomicsql_sqldefs
	import (
		orm	 "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	)

func (_this *DBContextNSql) LoadCompiledQuery(){

	_this.DBContextBaseNoSql.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
		"ns-asdax-W1": 
		{
			//CompiledQuery: ,
			CompiledQuery:``,			
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
			Tag: 			"ns-asdax",
			File: 			"test1.go",
			StartOff: 		23,
			EndOff:  		152,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Nb25leSA+PSBVc2VyTW9uZXkgJiYNCgkJCSh4LlVzZXJSb2xlSUQuUm9sZU5hbWUgPT0gUm9sZU5hbWVEZWZhdWx0IHx8IHguVXNlclJvbGVJRCA9PSBuaWwpDQoJ",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery: 	
                            []any{
					            "&&", 
                            	[]any{
					        	    ">=", "{#@money@#}", "{@@UserMoney@@}",
                            	}, 
                            	[]any{
					        	    "||", 
                            		[]any{
					        		    "==", "{#@userRole_ID.roleName@#}", "{@@RoleNameDefault@@}",
                            		}, 
                            		[]any{
					        		    "==", "{#@userRole_ID@#}", "nil",
                            		},
                            	},
                            },
		},
		"ns-tst1074-W1": 
		{
			//CompiledQuery: ,
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
			File: 			"test1.go",
			StartOff: 		25,
			EndOff:  		83,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQlyZXR1cm4geC5Vc2VyTmFtZSA9PSBVc2VyTmFtZQ0KCQ==",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery: 	
                            []any{
					            "==", "{#@userName@#}", "{@@UserName@@}",
                            },
		},
		"evcy59-W1": 
		{
			//CompiledQuery: ,
			CompiledQuery:``,			
			SelectSqlFields: nil,
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"evcy59",
			File: 			"test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery: 	
                            []any{
					            "==", "{#@userRole_ID.isActive@#}", "true",
                            },
		},
		"tst254-W1": 
		{
			//CompiledQuery: ,
			CompiledQuery:``,			
			SelectSqlFields: nil,
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst254",
			File: 			"test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery: 	
                            []any{
					            "==", "{#@userRole_ID.isActive@#}", "true",
                            },
		},
		"tst665-W1": 
		{
			//CompiledQuery: ,
			CompiledQuery:``,			
			SelectSqlFields: nil,
			OrderedFields:	[]string{ "UserRoleID", "UserRoleID.IsActive" },
			Fields:			map[string]string{
             "UserRoleID":"{#@userRole_ID@#}",
             "UserRoleID.IsActive":"{#@userRole_ID.isActive@#}",
            },
			ExternVar:			[]orm.TExternVar{
								 
								},
			Tag: 			"tst665",
			File: 			"test1.go",
			StartOff: 		25,
			EndOff:  		84,
			Hash:  			"ZnVuYyh4ICptLlVzZXIpIGJvb2wgew0KCQkJcmV0dXJuIHguVXNlclJvbGVJRC5Jc0FjdGl2ZQ0KCQk=",
			IsQryS:			false,
			SubQueries: 	nil,
			NosqlQuery: 	
                            []any{
					            "==", "{#@userRole_ID.isActive@#}", "true",
                            },
		},
	}
	
	_this.DBContextBaseNoSql.ProcessCompiledQuery( &_this.DBContextBaseNoSql.CompiledSqlQueries, false );
}

func (_this *DBContextNSql) CompiledQueries_GetDialectName() string { return "mongodb"; }


