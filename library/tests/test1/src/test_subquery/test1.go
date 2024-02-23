package test1_subquery

import (
	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	atmf "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/src/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
	test_where "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
)

const(
	EConst_EID = 123
	EConst_Name = "a'\"\n mm"
)
type TConstr struct {
	PathImg string
	EID int32
}
var EConst = TConstr{
	PathImg:"asada",
	EID: 342,
}

func Test1_init() (*orm.DBContext, string, error) {

	var connString = test_where.Test1_GetConnectionString()
	ctxBase, err := atmsql.OpenDB(connString, 10, 10)
	if err != nil {
		return nil, "initTest1", err
	}
	if ctxBase == nil {
		return nil, "initTest1", err
	}
	ctx, err := orm.New_DBContext(*ctxBase)
	if err != nil {
		return nil, "initTest1", err
	}

	Test_cleanUp(ctx)
	return ctx, "initTest1", err
}

func Test_cleanUp(ctx *orm.DBContext) {

	ctx.User.Qry("").DeleteModels()
	ctx.UserRole.Qry("").DeleteModels()
	ctx.Statusrole.Qry("").DeleteModels()

	//return errcode, nameTest, err;
}

//---------------------------------------------------------

func Test1_01(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: Test Subquery + var := "

	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	var UserName2 string = "b"
	

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var userRole = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err := ctx.UserRole.Qry("").InsertModel(&userRole)
	if err != nil || userRole_id == 0 {
		return 0, nameTest, err
	}

	//do the FK relation
	var user = m.User{UserName: UserName, Money: UserMoney}
	user.UserRoleID, err = ctx.UserRole.Qry("").
		WhereEq(ctx.UserRole_.RoleName, RoleNameDefault).GetFirstModel()
	if err != nil {
		return 0, nameTest, err
	}
	user1_Id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user1_Id == 0 {
		return 0, nameTest, err
	}

	var user2 = m.User{UserName: UserName2, Money: UserMoney}
	user2_Id, err := ctx.User.Qry("").InsertModel(&user2)
	if err != nil || user2_Id == 0 {
		return 0, nameTest, err
	}

	//var EConstEID1 = EConst.EID;
	//EConstEID1 = EConst.EID;
	//var EConst_EID1 int32= int32(EConst_EID);

	usersCnt1, err1 := ctx.User.Qry("tsql082a").Where(func(x *m.User) bool {

		return ((x.UserRoleID.RoleStatusID == nil) || x.UserRoleID.RoleStatusID.ID > 0) &&
			x.Money >= UserMoney &&
			x.ID != EConst_EID && 
			x.UserName != EConst_Name				 
	}).GetCount()

	if err1 != nil {
		return 0, nameTest, err
	}
	if usersCnt1 == 1000 {
		return 0, nameTest, nil
	}


	var bActive = true
	//do the query using subquery
	usersCnt, err := ctx.User.Qry("tsql082").WhereSubQ(func(x *m.User, q atmsql.IDBQuery) bool {

		ids, _ := ctx.UserRole.QryS("ids", q).Where(func(y *m.UserRole) bool {
			return y.RoleName == RoleNameDefault &&
				   //y.ID != EConstEID1 &&
				   //y.ID != EConst_EID &&
				((!y.Role_status_ID.Valid) ||
					y.Role_status_ID.Int32 == x.UserRoleID.RoleStatusID.ID)
		}).GetRowsAsFieldInt(ctx.UserRole_.ID)

		return ((x.UserRoleID.RoleStatusID == nil) || x.UserRoleID.RoleStatusID.ID > 0) &&
			x.Money >= UserMoney &&
			atmf.Sql_ArrayContain(ids, int64(x.UserRole_ID.Int32)) &&

			bActive
	}).GetCount()

	if err != nil {
		return 0, nameTest, err
	}
	if usersCnt == 1 {
		return 1, nameTest, nil
	}

	return 0, nameTest, nil
}

//---------------------------------------------------------

func Test1_02(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Test inline Subquery"

	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	var UserName2 string = "b"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var userRole = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err := ctx.UserRole.Qry("").InsertModel(&userRole)
	if err != nil || userRole_id == 0 {
		return 0, nameTest, err
	}

	//do the FK relation
	var user = m.User{UserName: UserName, Money: UserMoney}
	user.UserRoleID, err = ctx.UserRole.Qry("").WhereEq(ctx.UserRole_.RoleName, RoleNameDefault).GetFirstModel()
	if err != nil {
		return 0, nameTest, err
	}
	user1_Id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user1_Id == 0 {
		return 0, nameTest, err
	}

	var user2 = m.User{UserName: UserName2, Money: UserMoney}
	user2_Id, err := ctx.User.Qry("").InsertModel(&user2)
	if err != nil || user2_Id == 0 {
		return 0, nameTest, err
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
		return 0, nameTest, err
	}
	if usersCnt == 1 {
		return 1, nameTest, nil
	}

	return 0, nameTest, nil
}
