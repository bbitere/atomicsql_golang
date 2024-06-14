package test1_where

import (
	//"database/sql"
	"fmt"
	Str "strings"

	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicNSql"
	//atmsql_func "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/src/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
)

const NULL_ID string = "000000000000000000000000"
/*
type TConstr struct {
	PathImg string
	EID int32
}
var EConst = TConstr{
	PathImg:"asada",
	EID: 342,
}
*/
type TestFunc func(step int, bCheckName bool) (int, string, error)

/*
package test1_where
import (atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql")

// it is not commited intentionatly on github.
// please write your connection string.
func Test1_GetConnectionString() atmsql.TConnectionString{

	var connString = atmsql.TConnectionString{
		Host:     "localhost",
		Port:     5432,
		User:     "",
		Password: "",
		DbName:   "",
	}
	return connString;
}
*/

func Test1_init() (*orm.DBContextNSql, string, error) {

	var connString = Test1_GetConnectionString()
	ctxBase, err := atmsql.OpenDB_NoSql(connString, 0, 10)
	if ctxBase == nil {
		return nil, "initTest", err
	}

	ctx, err := orm.New_DBContextNSql(*ctxBase)
	if err != nil {
		return nil, "initTest1", err
	}
	if ctx.GetSqlName() != Str.ToLower(string(connString.SqlLang)) {
		return nil, "initTest1", fmt.Errorf("context DB is not exported as %s", connString.SqlLang)
	}

	Test_cleanUp(ctx)
	return ctx, "initTest1", nil
}

func Test_cleanUp(ctx *orm.DBContextNSql) {

	ctx.User.Qry("").DeleteModels()
	ctx.UserRole.Qry("").DeleteModels()
	ctx.Statusrole.Qry("").DeleteModels()

	//return errcode, nameTest, err;
}

func Test1_00(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: CheckIntegrity()"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var bResult = ctx.DBContextBaseNoSql.CheckIntegrity("..\\..\\")
	if bResult != "" {
		return 0, nameTest, err
	}
	return 1, nameTest, err
}

//---------------------------------------------------------

func Test1_01(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: DeleteRecords()"
	var RoleNameDefault = "default"

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

	var err1 = ctx.UserRole.Qry("").DeleteModels()
	if err1 != nil {
		return 0, nameTest, err
	}

	var count, err2 = ctx.UserRole.Qry("").GetCount()
	if count != 0 {
		if err != nil {
			return 0, nameTest, err2
		}
	}

	return 1, nameTest, nil
}

// ---------------------------------------------------------
func Test1_02N(step int, bCheckName bool) (int, string, error) {

	//ORM: DeleteAllRecords;; Delete all records
	var nameTest = "ORM: N!Test GetCount() + GetDistinctCount()"
	var RoleNameDefault = "default"

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

	_, err = ctx.UserRole.Qry("").InsertModel(&userRole)
	if err == nil {
		return 0, nameTest, err
	}
	//the insert should faild because the user has the same id as before

	return 1, nameTest, nil
}
func Test1_02(step int, bCheckName bool) (int, string, error) {

	//ORM: DeleteAllRecords;; Delete all records
	var nameTest = "ORM: Test GetCount() + GetDistinctCount()"
	var RoleNameDefault = "default"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var userRole1 = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err := ctx.UserRole.Qry("").InsertModel(&userRole1)
	if err != nil || userRole_id == 0 {
		return 0, nameTest, err
	}

	var userRole2 = m.UserRole{RoleName: RoleNameDefault}
	userRole_id, err = ctx.UserRole.Qry("").InsertModel(&userRole2)
	if err != nil || userRole_id == 0 {
		return 0, nameTest, err
	}

	var count1, err2 = ctx.UserRole.Qry("").GetCount()
	var count2, err3 = ctx.UserRole.Qry("").GetDistinct1Count(ctx.UserRole_.RoleName)
	if err2 != nil {
		return 0, nameTest, err
	}
	if err3 != nil {
		return 0, nameTest, err
	}

	if count2 != 1 {
		return 0, nameTest, err
	}
	if count1 != 2 {
		return 0, nameTest, err
	}

	return 1, nameTest, nil
}

// ---------------------------------------------------------
func Test1_03(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: Test GetFirstModelRel + InsertModel Multiple FK()"
	//var RoleNameDefault = "default";
	var RoleNameAdmin = "admin"
	var UserMoney float64 = 100
	//var UserName 	string =  "a";
	var UserName2 string = "b"
	var StatusNameActive string = "active"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	//insert user.fk.fk
	var user2 = m.User{UserName: UserName2, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameAdmin, IsActive: true,
			RoleStatusID: &m.Statusrole{StatusName: atmsql.Null_String(StatusNameActive)}}}
	_, err = ctx.User.Qry("").InsertModel(&user2)
	if err != nil || user2.UID == NULL_ID || user2.UID == ""||
		user2.UserRoleID== nil ||
		user2.UserRoleID.RoleStatusID == nil {
		return 0, nameTest, err
	}

	var usrM2, err4 = ctx.User.Qry("").GetFirstModelRel(ctx.User_.UserRoleID.Def(),
		ctx.User_.UserRoleID.RoleStatusID.Def())
	if err4 != nil || usrM2 == nil ||
		usrM2.UserRoleID == nil ||
		usrM2.UserRoleID.RoleStatusID == nil {
		return 0, nameTest, err
	}

	if usrM2.UserRoleID.RoleName != RoleNameAdmin {
		return 0, nameTest, err
	}
	if usrM2.UserRoleID.RoleStatusID.StatusName.String != StatusNameActive {
		return 0, nameTest, err
	}

	return 1, nameTest, nil
}

func Test1_04(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: UpdateModel() and test the new UserRole inserted"

	var newname = "Vasile"
	//var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	//var UserName2	string =  "b";

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var user = m.User{
		UserName: newname,
		Money:    UserMoney,
		UserRoleID: nil,
	}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	user = m.User{
		UserName: UserName,
		Money:    UserMoney,
		UserRoleID: nil,
	}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var usrT *m.User
	usrT, err = ctx.User.Qry("").WhereEq( ctx.User_.UserName, UserName ).
								GetFirstModelRel(ctx.User_.UserRoleID.Def())
	if usrT == nil || err != nil {
		return 0, nameTest, err
	}

	if usrT.UserName != UserName {
		//field is not updated
		return 0, nameTest, err
	}
	return 1, nameTest, nil
}


func Test1_05(step int, bCheckName bool) (int, string, error) {

	var nameTest = "ORM: UpdateModel() and test the new UserRole inserted"

	var newname = "Vasile"
	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName string = "a"
	//var UserName2	string =  "b";

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var user = m.User{
		UserName: UserName,
		Money:    UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault,
			IsActive: true},
	}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var usrT *m.User
	usrT, err = ctx.User.Qry("").WhereEq(ctx.User_.UserRoleID.IsActive, true).
		GetFirstModelRel(ctx.User_.UserRoleID.Def())
	if usrT == nil || err != nil {
		return 0, nameTest, err
	}
	
	usrT.UserName = newname
	usrT.UserRoleID.ID = 0 // sa il inserez din nou in tablea

	err = ctx.User.Qry("").UpdateModel(usrT)
	if err != nil {
		return 0, nameTest, err
	}

	usrT, err = ctx.User.Qry("").
		WhereEq(ctx.User_.UserRoleID.IsActive, true).
		WhereEq(ctx.User_.UserName, newname).
		GetFirstModel()
	if usrT == nil || err != nil {
		return 0, nameTest, err
	}

	if usrT.UserName != newname {
		//field is not updated
		return 0, nameTest, err
	}
	return 1, nameTest, nil
}

// ---------------------------------------------------------
func Test1_08(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Test WhereEq + Relation inside Where"

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

	//do the query using FK relation
	users1, err := ctx.User.Qry("ns-asdax").Where(func(x *m.User) bool {
		return x.Money >= UserMoney &&
			(x.UserRoleID == nil || x.UserRoleID.RoleName == RoleNameDefault )
	}).GetModelsRel(ctx.User_.UserRoleID.RoleStatusID.Def())
	if err != nil {
		return 0, nameTest, err
	}
	if( len(users1) != 2 ){
		return 0, nameTest, err
	}

	usersCount, err := ctx.User.Qry("ns-tst1074").Where(func(x *m.User) bool {
		return x.UserName == UserName
	}).GetCount()
	if err != nil {
		return 0, nameTest, err
	}

	//it must be 1 row
	if len(users1) != 2 || usersCount != 1 {
		return 0, nameTest, err
	}

	var u = users1[0]
	if u.Money < UserMoney && u.UserRoleID.RoleName == RoleNameDefault {
		return 0, "error: not passed", nil
	}

	return 1, nameTest, nil
}

// ---------------------------------------------------------
func Test1_09(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Select( Where )"

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

	var user = m.User{UserName: UserName, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: false}}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var user1 = m.User{UserName: UserName2, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user1)
	if err != nil {
		return 0, nameTest, err
	}

	//---------------------------

	type vUser1 struct {
		m.User   `atomicsql:"copy-model"`
		UserRole string
	}

	//Nopp();

	usrs4, err := atmsql.SelectN(ctx.User.Qry("evcy59").
		Where(func(x *m.User) bool {
			return x.UserRoleID.IsActive
		}),
		func(x *m.User) *vUser1 {
			return &vUser1{
				User:     *x,
				UserRole: x.UserRoleID.RoleName,
			}
		}).GetModels()

	if err != nil || len(usrs4) != 1 {
		return 0, nameTest, nil
	}

	return 1, nameTest, nil
}

// ---------------------------------------------------------
/*
func Test1_10(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: Select( Aggregate( Where() ) "

	var RoleNameDefault = "default"
	var RoleNameAdmin = "Admin"
	var UserMoney float64 = 100
	var UserName1 string = "a"
	var UserName2 string = "b"
	var UserName3 string = "c"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var user = m.User{UserName: UserName1, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: false}}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var user1 = m.User{UserName: UserName2, Money: 2 * UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user1)
	if err != nil {
		return 0, nameTest, err
	}

	//second has different user role
	var user2 = m.User{UserName: UserName3, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameAdmin, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user2)
	if err != nil {
		return 0, nameTest, err
	}

	//---------------------------

	type TUserAggr struct {
		atmsql.Generic_MODEL
		UserRoleID  *m.UserRole
		UserRole_ID sql.NullInt32
		Time1       []sql.NullTime
		Money       []float64
	}
	type TUserView struct {
		atmsql.Generic_MODEL
		UserRoleName string
		MinTime1     sql.NullTime
		SumMoney     float64
	}

	usrs4, err := atmsql.Select(
		atmsql.Aggregate[m.User, TUserAggr](ctx.User.Qry("evcy58").
			Where(func(x *m.User) bool {
				return x.UserRoleID.IsActive
			}),
		),
		func(x *TUserAggr) *TUserView {

			return &TUserView{
				UserRoleName: x.UserRoleID.RoleName,
				MinTime1:     atmsql_func.Sql_MinDateN(x.Time1),
				SumMoney:     atmsql_func.Sql_SumF64(x.Money),
			}
		}).OrderAsc("UserRoleName").GetModels()

	if err != nil {
		return 0, nameTest, err
	}
	if len(usrs4) != 2 {
		return 0, nameTest, err
	}
	if usrs4[0].SumMoney != UserMoney {
		return 0, nameTest, err
	}
	if usrs4[1].SumMoney != 2*UserMoney {
		return 0, nameTest, err
	}
	return 1, nameTest, nil
}
*/

func Test1_11(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: GetValueString"

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

	var user = m.User{UserName: UserName, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: false}}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var user1 = m.User{UserName: UserName2, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user1)
	if err != nil {
		return 0, nameTest, err
	}

	//---------------------------

	var usrName, err1 = ctx.User.Qry("tst253").GetValueString(func(x *m.User) string {
		return x.UserRoleID.RoleName
	})

	if err1 != nil || usrName != RoleNameDefault {
		return 0, nameTest, nil
	}

	return 1, nameTest, nil
}

func Test1_12(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: GetValuesString( Where )"

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

	var user = m.User{UserName: UserName, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: false}}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var user1 = m.User{UserName: UserName2, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user1)
	if err != nil {
		return 0, nameTest, err
	}

	//---------------------------	
	var userRoleIDs, err1 = ctx.User.Qry("tst254").
		Where(func(x *m.User) bool {
			return x.UserRoleID.IsActive
		}).
		GetValuesInt(func(x *m.User) int64 {
			return int64(x.UserRoleID.ID)
		})

	if err1 != nil || len(userRoleIDs) == 2 {
		return 0, nameTest, nil
	}

	return 1, nameTest, nil
}


func Test1_13(step int, bCheckName bool) (int, string, error) {

	//insert 2 users, 1 userrole.test where( FK. )
	var nameTest = "ORM: GetValuesString( Where )"

	var RoleNameDefault = "default"
	var UserMoney float64 = 100
	var UserName1 string = "a"
	var UserName2 string = "b"
	var UserName3 string = "c"

	ctx, _, err := Test1_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if err != nil {
		return 0, nameTest, err
	}

	var user = m.User{UserName: UserName1, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: false}}
	_, err = ctx.User.Qry("").InsertModel(&user)
	if err != nil {
		return 0, nameTest, err
	}

	var user2 = m.User{UserName: UserName2, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user2)
	if err != nil {
		return 0, nameTest, err
	}
	var user3 = m.User{UserName: UserName3, Money: UserMoney,
		UserRoleID: &m.UserRole{RoleName: RoleNameDefault, IsActive: true}}
	_, err = ctx.User.Qry("").InsertModel(&user3)
	if err != nil {
		return 0, nameTest, err
	}

	//---------------------------
	//Nopp()
	var newQry, err1 = ctx.User.Qry("tst665").
		Where(func(x *m.User) bool {
			return x.UserRoleID.IsActive
		}).CloneQry()
		
	var cnt,   errCnt  = newQry.GetCount()
	var users, errRows = newQry.OrderAsc( ctx.User_.UserName).SetLimit( 0, 1).GetModels()

	if( errCnt != nil){
		return 0, nameTest, errCnt
	}
	if( errRows != nil){
		return 0, nameTest, errRows
	}

	if err1 != nil || cnt == 3 || len(users) != 1 || cnt != 2{
		return 0, nameTest, nil
	}

	return 1, nameTest, nil
}
