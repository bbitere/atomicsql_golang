package test1_subquery

import (
	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	atmf "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/src/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
	test_where "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
)

func Test1_init() (*orm.DBContext, error, string) {

	var connString = test_where.Test1_GetConnectionString()
	ctxBase, err := atmsql.OpenDB(connString, 10, 10)
	if ctxBase == nil {
		return nil, err, "initTest"
	}

	ctx, err := orm.New_DBContext(*ctxBase)
	if err != nil {
		return nil, err, "initTest1"
	}

	Test_cleanUp(ctx)
	return ctx, err, "initTest1"
}

func Test_cleanUp(ctx *orm.DBContext) {

	ctx.User.Qry("").DeleteModels()
	ctx.UserRole.Qry("").DeleteModels()
	ctx.Statusrole.Qry("").DeleteModels()

	//return errcode, err, nameTest;
}

//---------------------------------------------------------

func Test1_01(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Test Subquery + var := "

	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	var UserName2 string = "b"

	ctx, err, _ := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, err, nameTest
	}

	var userRole = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err := ctx.UserRole.Qry("").InsertModel(&userRole)
	if err != nil || userRole_id == 0 {
		return 0, err, nameTest
	}

	//do the FK relation
	var user = m.User{UserName: UserName, Money: UserMoney}
	user.UserRoleID, err = ctx.UserRole.Qry("").
		WhereEq(ctx.UserRole_.RoleName, RoleNameDefault).GetFirstModel()
	if err != nil {
		return 0, err, nameTest
	}
	user1_Id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user1_Id == 0 {
		return 0, err, nameTest
	}

	var user2 = m.User{UserName: UserName2, Money: UserMoney}
	user2_Id, err := ctx.User.Qry("").InsertModel(&user2)
	if err != nil || user2_Id == 0 {
		return 0, err, nameTest
	}

	var bActive = true
	//do the query using subquery
	usersCnt, err := ctx.User.Qry("tsql082").WhereSubQ(func(x *m.User, q atmsql.IDBQuery) bool {

		ids, _ := ctx.UserRole.QryS("ids", q).Where(func(y *m.UserRole) bool {
			return y.RoleName == RoleNameDefault && y.Role_status_ID.Int32 == x.UserRoleID.RoleStatusID.ID
		}).GetRowsAsFieldInt(ctx.UserRole_.ID)

		return x.Money >= UserMoney &&
			atmf.Sql_ArrayContain(ids, int64(x.UserRole_ID.Int32)) &&
			x.UserRoleID.RoleStatusID.ID > 0 &&
			bActive
	}).GetCount()

	if err != nil {
		return 0, err, nameTest
	}
	if usersCnt == 1 {
		return 1, nil, nameTest
	}

	return 0, nil, nameTest
}

//---------------------------------------------------------

func Test1_02(step int, bCheckName bool) (int, error, string) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Test inline Subquery"

	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	var UserName2 string = "b"

	ctx, err, _ := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, err, nameTest
	}

	var userRole = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err := ctx.UserRole.Qry("").InsertModel(&userRole)
	if err != nil || userRole_id == 0 {
		return 0, err, nameTest
	}

	//do the FK relation
	var user = m.User{UserName: UserName, Money: UserMoney}
	user.UserRoleID, err = ctx.UserRole.Qry("").WhereEq(ctx.UserRole_.RoleName, RoleNameDefault).GetFirstModel()
	if err != nil {
		return 0, err, nameTest
	}
	user1_Id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user1_Id == 0 {
		return 0, err, nameTest
	}

	var user2 = m.User{UserName: UserName2, Money: UserMoney}
	user2_Id, err := ctx.User.Qry("").InsertModel(&user2)
	if err != nil || user2_Id == 0 {
		return 0, err, nameTest
	}

	var bActive = true
	//do the query using subquery
	usersCnt, err := ctx.User.Qry("tsql147").WhereSubQ(func(x *m.User, q atmsql.IDBQuery) bool {

		//var cnt, _ = ctx.UserRole.QryS("cnt", q).WhereEq( ctx.UserRole_.ID, x.UserRole_ID.Int32 ).GetCount()
		var cnt, _ = ctx.UserRole.QryS("cnt", q).
			Where(func(y *m.UserRole) bool {
				return y.ID == x.UserRole_ID.Int32
			}).GetCount()
		return x.Money >= UserMoney && cnt > 0 && bActive
	}).GetCount()

	if err != nil {
		return 0, err, nameTest
	}
	if usersCnt == 1 {
		return 1, nil, nameTest
	}

	return 0, nil, nameTest
}
