package test1_where

import (
	atmsql_func "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
	//orm "github.com/bbitere/atomicsql_golang.git/tests/test1/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
)

func Nopp(){

}

func TestMisc_01(step int, bCheckName bool) ( int, string, error) {
	
	var nameTest = "ORM: test isNull(x.ForeignKey)";	
	var RoleNameDefault = "default";
	var nameUser = "aa";
	var uuid 	 = "12312-2145314-12314124";
	
	ctx, _, err := Test1_init();// (orm.DBContextBase, string, error){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return 0, nameTest, err;}

	

	var user = m.User{ UserName: nameUser, UUID:uuid,
	UserRoleID: &m.UserRole{ RoleName: RoleNameDefault, IsActive: true},}
	user_id, err := ctx.User.Qry("").InsertModel( &user );
	if( err != nil || user_id == 0 ){return  0, nameTest, err;}
	
	var countActive, _ = ctx.User.Qry("tst143").Where( func(x *m.User) bool{
							return atmsql_func.Sql_IIF( x.UserRoleID != nil, x.UserRoleID.IsActive, false);
						}).GetCount();

	if( countActive == 0){
		return  0, nameTest, err;
	}

	return 1, nameTest, nil;
}