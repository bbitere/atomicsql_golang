package test1_crud

import (
	"fmt"

	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	//atmf "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/src/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/src/mymodels"
	test1_where "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
)

func Example_init() (*orm.DBContext, error, string) {

	var connString = test1_where.Test1_GetConnectionString()
	ctxBase, err := atmsql.OpenDB(connString, 10, 10)
	if ctxBase == nil {
		return nil, err, "initTest"
	}

	ctx, err := orm.New_DBContext(*ctxBase)
	if err != nil {
		return nil, err, "initTest1"
	}

	ctx.User.Qry("").DeleteModels()
	ctx.UserRole.Qry("").DeleteModels()
	ctx.Statusrole.Qry("").DeleteModels()

	return ctx, err, "initTest1"
}

// ---------------------------------------------------------
func Tst_Example_CreateUser(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_CreateUser"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var user, err = Example_CreateUser(ctx, "aa", "24234-5252315-25234")
	if user == nil {
		return 0, err, nameTest
	}

	return atmsql.IFF(err == nil, 1, 0), err, nameTest

}
func Example_CreateUser(ctx *orm.DBContext, name string, uuid string) (*m.User, error) {

	var user = m.User{UserName: name, UUID: uuid}
	user_id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user_id == 0 {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("insert not working")
	}

	return &user, nil
}

func Tst_Example_Create2Users(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_Create2Users"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var users, err = Example_Create2Users(ctx, "aa", "24234-5252315-25234", "bb", "24234-5252315-2523124")
	if len(users) != 2 {
		return 0, err, nameTest
	}

	return atmsql.IFF(err == nil, 1, 0), err, nameTest
}
func Example_Create2Users(ctx *orm.DBContext,
	name string, uuid string,
	name2 string, uuid2 string) ([]*m.User, error) {

	var users = []*m.User{}

	var user1 = m.User{UserName: name, UUID: uuid}
	var user2 = m.User{UserName: name2, UUID: uuid2}

	users = []*m.User{&user1, &user2}
	var err = ctx.User.Qry("").InsertOrUpdateModels(users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func Tst_Example_RetrieveUser(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_RetrieveUser"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var uuid = "24234-5252315-25234"
	var user1, err1 = Example_CreateUser(ctx, "aa", uuid)
	if err1 != nil || user1 == nil {
		return 0, err1, nameTest
	}

	var user2, err2 = Example_RetrieveUser(ctx, uuid)
	if err2 != nil || user2 == nil {
		return 0, err2, nameTest
	}

	if user1.ID != user2.ID {
		return 0, err2, nameTest
	}

	return 1, nil, nameTest
}
func Example_RetrieveUser(ctx *orm.DBContext, uuid string) (*m.User, error) {

	var model, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).GetFirstModel()
	if err1 != nil {
		return nil, err1
	}

	return model, nil
}

func Example_RetrieveUserByName(ctx *orm.DBContext, name string) (*m.User, error) {

	var model, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UserName, name).GetFirstModel()
	if err1 != nil {
		return nil, err1
	}

	return model, nil
}

func Tst_Example_RetrieveUsers(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_RetrieveUsers"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var uuid = "24234-5252315-25234"
	var users1, err1 = Example_Create2Users(ctx, "aa", uuid, "bb", uuid)
	if err1 != nil || len(users1) == 0 {
		return 0, err1, nameTest
	}

	var users2, err2 = Example_RetrieveUsers(ctx, uuid)
	if err2 != nil || len(users2) == 0 {
		return 0, err2, nameTest
	}

	if len(users1) != len(users2) {
		return 0, err2, nameTest
	}

	return 1, nil, nameTest
}
func Example_RetrieveUsers(ctx *orm.DBContext, uuid string) ([]*m.User, error) {

	var models, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).GetModels()
	if err1 != nil {
		return nil, err1
	}

	return models, nil
}

// -----------------------------------------------------------------------
func Tst_Example_DeleteUser(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_DeleteUser"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var uuid = "24234-5252315-25234"
	var users, err = Example_Create2Users(ctx, "aa", uuid, "bb", uuid)
	if len(users) != 2 {
		return 0, err, nameTest
	}

	var err2 = Example_DeleteUser(ctx, uuid)
	if err2 != nil {
		return 0, err, nameTest
	}

	var users2, err3 = Example_RetrieveUsers(ctx, uuid)
	if err3 != nil {
		return 0, err, nameTest
	}
	if len(users2) != 1 {
		return 0, err, nameTest
	}

	return 1, nil, nameTest
}
func Example_DeleteUser(ctx *orm.DBContext, uuid string) error {

	var model, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).GetFirstModel()
	if err1 != nil {
		return err1
	}

	var err = ctx.User.Qry("").DeleteModel(model)
	if err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------
func Tst_Example_DeleteUsers(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_DeleteUsers"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var uuid = "24234-5252315-25234"
	var users, err = Example_Create2Users(ctx, "aa", uuid, "bb", uuid)
	if len(users) != 2 {
		return 0, err, nameTest
	}

	var err2 = Example_DeleteUsers(ctx, uuid)
	if err2 != nil {
		return 0, err, nameTest
	}

	var user2, err3 = Example_RetrieveUser(ctx, uuid)
	if err3 != nil {
		return 0, err, nameTest
	}
	if user2 != nil {
		return 0, err, nameTest
	}

	return 1, nil, nameTest
}

func Example_DeleteUsers(ctx *orm.DBContext, uuid string) error {

	var err = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).DeleteModels()
	if err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------
func Tst_Example_UpdateUser(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_UpdateUser"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var userName = "cc"
	var uuid = "24234-5252315-25234"
	var users, err = Example_Create2Users(ctx, "aa", uuid, "bb", uuid)
	if len(users) != 2 {
		return 0, err, nameTest
	}

	var err3 = Example_UpdateUser(ctx, uuid, userName)
	if err3 != nil {
		return 0, err, nameTest
	}

	var user2, err4 = Example_RetrieveUserByName(ctx, userName)
	if err4 != nil {
		return 0, err, nameTest
	}
	if user2 == nil {
		return 0, err, nameTest
	}

	return 1, nil, nameTest
}

func Example_UpdateUser(ctx *orm.DBContext, uuid string, newName string) error {

	var model, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).GetFirstModel()
	if err1 != nil {
		return err1
	}

	//update same fields
	model.UserName = newName

	var err = ctx.User.Qry("").UpdateModel(model)
	if err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------
func Tst_Example_UpdateUsers(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_UpdateUsers"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var userName = "cc"
	var uuid = "24234-5252315-25234"
	var users, err = Example_Create2Users(ctx, "aa", uuid, "bb", uuid)
	if len(users) != 2 {
		return 0, err, nameTest
	}

	var err3 = Example_UpdateUsers(ctx, uuid, userName)
	if err3 != nil {
		return 0, err, nameTest
	}

	var users2, err4 = Example_RetrieveUsers(ctx, uuid)
	if err4 != nil {
		return 0, err, nameTest
	}
	if len(users2) != 2 {
		return 0, err, nameTest
	}

	return 1, nil, nameTest
}
func Example_UpdateUsers(ctx *orm.DBContext, uuid string, newName string) error {

	var models, err1 = ctx.User.Qry("").WhereEq(ctx.User_.UUID, uuid).GetModels()
	if err1 != nil {
		return err1
	}

	//update same fields
	for _, m1 := range models {
		m1.UserName = newName
	}

	var err = ctx.User.Qry("").UpdateModels(&models)
	if err != nil {
		return err
	}

	return nil
}

//---------------------------------------------------------

// -----------------------------------------------------------------------
func Tst_Example_CreateUserRelation(step int, bCheckName bool) (int, error, string) {

	var nameTest = "ORM: Example_CreateUserRelation"

	ctx, errCtx, _ := Example_init() // (orm.DBContextBase, error, string){
	if ctx != nil {
		defer ctx.Close()
	}
	if errCtx != nil {
		return 0, errCtx, nameTest
	}

	var uuid = "24234-5252315-25234"
	var RoleAdmin = "admin"

	var err = Example_CreateUserRelation(ctx, "aa", uuid, RoleAdmin)
	if err != nil {
		return 0, errCtx, nameTest
	}

	var user, err2 = Example_RetrieveUserRelation1(ctx, uuid)
	if user == nil || err2 != nil {
		return 0, errCtx, nameTest
	}
	if !(user != nil && user.UserRoleID != nil && user.UserRoleID.RoleName == RoleAdmin) {
		return 0, errCtx, nameTest
	}

	user, err2 = Example_RetrieveUserRelation2(ctx, uuid)
	if user == nil || err2 != nil {
		return 0, errCtx, nameTest
	}
	if !(user != nil && user.UserRoleID != nil && user.UserRoleID.RoleName == RoleAdmin) {
		return 0, errCtx, nameTest
	}

	return 1, nil, nameTest
}
func Example_CreateUserRelation(ctx *orm.DBContext, nameUser string, uuid string, userRole string) error {

	var user = m.User{UserName: nameUser, UUID: uuid,
		UserRoleID: &m.UserRole{RoleName: userRole, IsActive: true}}
	user_id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user_id == 0 {
		return err
	}

	return nil
}

func Example_CreateUserRelationCheck(ctx *orm.DBContext,
	nameUser string, uuid string, userRoleName string) error {

	var userRole, err1 = ctx.UserRole.Qry("").
		WhereEq(ctx.UserRole_.RoleName, userRoleName).
		GetFirstModel()
	if err1 != nil {
		return err1
	}
	if userRole == nil {
		userRole = &m.UserRole{RoleName: userRoleName, IsActive: true}
	}

	var user = m.User{UserName: nameUser, UUID: uuid, UserRoleID: userRole}
	user_id, err := ctx.User.Qry("").InsertModel(&user)
	if err != nil || user_id == 0 {
		return err
	}

	return nil
}

func Example_RetrieveUserRelation1(ctx *orm.DBContext, uuid string) (*m.User, error) {

	var model, err1 = ctx.User.Qry("").
		WhereEq(ctx.User_.UserRoleID.IsActive, true).
		WhereEq(ctx.User_.UserRoleID.RoleName, "admin").
		WhereEq(ctx.User_.UUID, uuid).
		GetFirstModelRel(ctx.User_.UserRoleID.Def())
	if err1 != nil {
		return nil, err1
	}

	if model != nil && model.UserRoleID != nil && model.UserRoleID.RoleName == "admin" {

	}

	return model, nil
}

func Example_RetrieveUserRelation2(ctx *orm.DBContext, uuid string) (*m.User, error) {

	//Nopp();
	var model, err1 = ctx.User.Qry("tst1340").
		Where(func(x *m.User) bool {
			return x.UserRoleID.IsActive == true &&
				x.UserRoleID.RoleName == "admin" &&
				x.UUID == uuid
		}).
		GetFirstModelRel(ctx.User_.UserRoleID.Def())
	if err1 != nil {
		return nil, err1
	}

	if model != nil && model.UserRoleID != nil && model.UserRoleID.RoleName == "admin" {

	}

	return model, nil
}

