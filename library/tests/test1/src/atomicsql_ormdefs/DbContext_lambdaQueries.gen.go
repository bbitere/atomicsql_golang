
/* this class is generated automatically by compiler AtomicSql goscanner.exe exporter*/

package atomicsql_sqldefs
	import (
		orm	 "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	)

func (_this *DBContext) LoadCompiledQuery(){

	_this.DBContextBase.CompiledSqlQueries	= map[string]orm.TCompiledSqlQuery{
		"evcy59-S0": 
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
			StartOff: 		10144,
			EndOff:  		10211,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
		},
		"evcy58-S0": 
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
			StartOff: 		12042,
			EndOff:  		12111,
			Hash:  			"6CFA253FCAB447A6E1C2AA01869B10E9",
		},"tst1_066-S0": 
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
		},
	}
	
	_this.DBContextBase.ProcessCompiledQuery( &_this.DBContextBase.CompiledSqlQueries, false );
}

func (_this *DBContext) CompiledQueries_GetDialectName() string { return "postgres"; }


