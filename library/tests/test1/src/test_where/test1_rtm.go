package test1_where

import (
	"database/sql"

	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	atmsql_func "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"

	//orm "github.com/bbitere/atomicsql_golang.git/tests/test1/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
	//test1_where "github.com/bbitere/atomicsql_golang.git/tests/test1/test_where"
)

func Test1Rtm_10( step int, bCheckName bool) ( int, error, string) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Select( Aggregate( Where() ) "
	
	var RoleNameDefault = "default";
	var RoleNameAdmin = "Admin";
	var UserMoney 	float64 =  100;
	var UserName1 	string =  "a";
	var UserName2	string =  "b";
	var UserName3	string =  "c";

	ctx, err, _ := Test1_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return 0, err, nameTest;}

	var user = m.User{UserName: UserName1, Money: UserMoney,
			UserRoleID: &m.UserRole{ RoleName: RoleNameDefault, IsActive: false},};
	_, err = ctx.User.Qry("").InsertModel(&user);

	var user1 = m.User{UserName: UserName2, Money: 2*UserMoney,
	UserRoleID: &m.UserRole{ RoleName: RoleNameDefault, IsActive: true},};
	_, err = ctx.User.Qry("").InsertModel(&user1);

	//second has different user role
	var user2 = m.User{UserName: UserName3, Money: UserMoney,
		UserRoleID: &m.UserRole{ RoleName: RoleNameAdmin, IsActive: true},};
	_, err = ctx.User.Qry("").InsertModel(&user2);

	//---------------------------
	type TUserAggr struct {

		atmsql.Generic_MODEL
		UserRoleID          *m.UserRole
		UserRole_ID         sql.NullInt32
		Time1               []sql.NullTime
		Money               []float64 
	}
	type TUserView struct{

		atmsql.Generic_MODEL
		UserRoleName        string
		MinTime1           	sql.NullTime
		SumMoney            float64 
	}
	Nopp();
	
	usrs4, err := atmsql.Select(
					atmsql.Aggregate[ m.User, TUserAggr ]( 
						ctx.User.Qry("tst1_066").ToRTM(true, ctx.User_.UserRoleID.Def() ).
						Where(func(x *m.User) bool {
							return x.UserRoleID.IsActive == true
						}),
					),
					func (x *TUserAggr ) *TUserView {

						return &TUserView{
							UserRoleName: x.UserRoleID.RoleName,
							MinTime1: atmsql_func.Sql_MinDateN( x.Time1 ),
							SumMoney: atmsql_func.Sql_SumF64( x.Money ),
						}
					}).OrderAsc( "UserRoleName" ).GetModels();

	if( err != nil){
		return 0, err, nameTest
	}
	if(len(usrs4) != 2){
		return 0, err, nameTest
	}
	if( usrs4[0].SumMoney != UserMoney){
		return 0, err, nameTest
	}
	if( usrs4[1].SumMoney != 2*UserMoney){
		return 0, err, nameTest
	}
	return 1, nil, nameTest;
}