package test1_crud

import (
	"fmt"

	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/atomicsql_ormdefs"
	m "github.com/bbitere/atomicsql_golang.git/tests/test1/mymodels"
)


func Example_init() (*orm.DBContext, error, string){

	var connString = atmsql.TConnectionString{
		Host:     "localhost",
		Port:     5432,
		User:     "dev_original",
		Password: "XCZ12345678",
		DbName:   "test1",
	}
	ctxBase, err := atmsql.OpenDB(connString, atmsql.ESqlDialect.Postgress, 10, 10)
	if ctxBase == nil {
		return nil, err, "initTest"
	}

	ctx,err := orm.New_DBContext(*ctxBase)
	if( err != nil){
		return nil, err, "initTest1";
	}

	ctx.User.Qry("").DeleteModels();
	ctx.UserRole.Qry("").DeleteModels();
	ctx.StatusRole.Qry("").DeleteModels();

	return ctx, err, "initTest1";
}

//---------------------------------------------------------
func Tst_Example_CreateUser(step int, bCheckName bool) ( int, error, string) {
	
	var nameTest = "ORM: Example_CreateUser";
	var user, err = Example_CreateUser( "aa", "24234-5252315-25234");
	if( user == nil){return 0, err, nameTest}
	
	return atmsql.IFF(err == nil, 1, 0), err, nameTest;

}
func Example_CreateUser(name string, uuid string) ( *m.User, error) {
	
	ctx, err, _ := Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return nil, err;}

	var user = m.User{ UserName: name, UUID:uuid}
	user_id, err := ctx.User.Qry("").InsertModel(&user);
	if( err != nil || user_id == 0 ){return  nil, err;}

	if( user.ID == 0 ){return  nil, fmt.Errorf("insert not working");}

	return &user, nil
}

func Tst_Example_Create2Users(step int, bCheckName bool) ( int, error, string) {
	
	var nameTest = "ORM: Example_Create2Users";
	var users, err = Example_Create2Users( "aa", "24234-5252315-25234", "bb", "24234-5252315-2523124");
	if( len(users) != 2 ){return 0, err, nameTest}

	return atmsql.IFF(err == nil, 1, 0), err, nameTest;
}
func Example_Create2Users(name string, uuid string, 
						 name2 string, uuid2 string) ( []*m.User, error) {
	
	var users = []*m.User{}
	ctx, err, _ := Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return users, err;}

	var user1 = m.User{ UserName: name, UUID:uuid}
	var user2 = m.User{ UserName: name2, UUID:uuid2}

	users = []*m.User{ &user1, &user2}
	err = ctx.User.Qry("").InsertOrUpdateModels( users );
	if( err != nil ){return  users, err;}

	return users, nil
}

func Tst_Example_RetrieveUser(step int, bCheckName bool) ( int, error, string) {
	
	var nameTest = "ORM: Example_RetrieveUser";

	var uuid = "24234-5252315-25234";
	var user1, err1 = Example_CreateUser( "aa", uuid);
	if( err1 != nil || user1 == nil){ return 0, err1, nameTest}

	var user2, err2 = Example_RetrieveUser( uuid );
	if( err2 != nil || user2 == nil){ return 0, err2, nameTest}

	if( user1.ID != user2.ID ) { return 0, err2, nameTest}
	
	return 1, nil, nameTest
}
func Example_RetrieveUser(uuid string) ( *m.User, error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return nil, err;}

	var model, err1 = ctx.User.Qry("").WhereEq( ctx.User_.UUID, uuid ).GetFirstModel();
	if( err1 != nil ){return  nil, err;}

	return model, nil
}

func Tst_Example_RetrieveUsers(step int, bCheckName bool) ( int, error, string) {
	
	var nameTest = "ORM: Example_RetrieveUsers";

	var uuid = "24234-5252315-25234";
	var users1, err1 = Example_Create2Users( "aa", uuid, "bb", uuid);
	if( err1 != nil || len(users1) == 0){ return 0, err1, nameTest}

	var users2, err2 = Example_RetrieveUsers( uuid );
	if( err2 != nil || len(users2) == 0 ){ return 0, err2, nameTest}

	if( len(users1) != len(users2) ) { return 0, err2, nameTest}
	
	return 1, nil, nameTest
}
func Example_RetrieveUsers(uuid string) ( []*m.User, error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return nil, err;}

	var models, err1 = ctx.User.Qry("").WhereEq( ctx.User_.UUID, uuid ).GetModels();
	if( err1 != nil ){return  nil, err;}

	return models, nil
}

func Example_DeleteUser(uuid string) ( error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}

	var model, err1 = ctx.User.Qry("").WhereEq( ctx.User_.UUID, uuid ).GetFirstModel();
	if( err1 != nil ){return  err;}

	err = ctx.User.Qry("").DeleteModel( model );
	if( err != nil ){return  err;}

	return nil
}
func Example_DeleteUsers(uuid string) ( error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}

	err = ctx.User.Qry("").WhereEq( ctx.User_.UUID, uuid).DeleteModels();
	if( err != nil ){return  err;}

	return nil
}

func Example_UpdateUser(uuid string, newName string) ( error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}

	var model, err1 = ctx.User.Qry("").WhereEq( ctx.User_.UUID, uuid ).GetFirstModel();
	if( err1 != nil ){return  err;}

	//update same fields
	model.UserName = newName;

	err = ctx.User.Qry("").UpdateModel(model);
	if( err != nil ){return  err;}

	return nil
}

func Example_UpdateUsers(models[] *m.User, newName string) ( error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}


	//update same fields
	for _, m1 := range(models){
		m1.UserName = newName;
	}

	err = ctx.User.Qry("").UpdateModels( &models );
	if( err != nil ){return  err;}

	return nil
}



//---------------------------------------------------------
func Example_CreateUserRelation(nameUser string, uuid string, userRole string) ( error) {
	
	ctx, err, _ := Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}

	var user = m.User{ UserName: nameUser, UUID:uuid,
					UserRoleID: &m.UserRole{ RoleName: userRole, IsActive: true},}
	user_id, err := ctx.User.Qry("").InsertModel( &user );
	if( err != nil || user_id == 0 ){return  err;}

	return nil
}

func Example_CreateUserRelationCheck(nameUser string, uuid string, userRoleName string) ( error) {
	
	ctx, err, _ := Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return err;}


	var userRole, err1 = ctx.UserRole.Qry("").
					WhereEq( ctx.UserRole_.RoleName, userRoleName).
					GetFirstModel();
	if( err1 != nil ){return  err;}
	if( userRole == nil ){
		userRole = &m.UserRole{ RoleName: userRoleName, IsActive: true};
	}

	var user = m.User{ UserName: nameUser, UUID:uuid, UserRoleID: userRole,}
	user_id, err := ctx.User.Qry("").InsertModel( &user );
	if( err != nil || user_id == 0 ){return  err;}

	return nil
}

func Example_RetrieveUserRelation1(uuid string) ( *m.User, error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return nil, err;}

	var model, err1 = ctx.User.Qry("").
						WhereEq( ctx.User_.UserRoleID.IsActive, true ).
						WhereEq( ctx.User_.UserRoleID.RoleName, "admin" ).
						WhereEq( ctx.User_.UUID, uuid ).
						GetFirstModel();
	if( err1 != nil ){return  nil, err;}

	if( model != nil && model.UserRoleID.RoleName == "admin"){

	}

	return model, nil
}

func Example_RetrieveUserRelation2(uuid string) ( *m.User, error) {
	
	var ctx, err, _ = Example_init();// (orm.DBContextBase, error, string){	
	if( ctx != nil ){
		defer ctx.Close()
	}
	if( err != nil ){return nil, err;}

	var model, err1 = ctx.User.Qry("").
						Where(  func(x *m.User) bool{
							return x.UserRoleID.IsActive == true &&
							x.UserRoleID.RoleName == "admin" &&
							x.UUID == uuid 
						}).
						GetFirstModelRel( ctx.User_.UserRoleID.Def() );
	if( err1 != nil ){return  nil, err;}

	if( model != nil && model.UserRoleID.RoleName == "admin"){
		
	}

	return model, nil
}

