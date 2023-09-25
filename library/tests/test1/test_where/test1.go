package test1_where

import (

	//atsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	orm "github.com/bbitere/atomicsql_golang.git/tests/test1/atomicsql_sqldefs"
	//m     "github.com/bbitere/atomicsql_golang.git/tests/test1/mymodels"
)



func Test1_init() (orm.DBContextBase, error, string){

	var connString = orm.TConnectionString{
		Host:     "localhost",
		Port:     5432,
		User:     "dev_original",
		Password: "XCZ12345678",
		DbName:   "test1",
	}
	ctxBase, err := orm.OpenDB(connString, orm.ESqlDialect.Postgress, 10, 10)
	if ctxBase == nil {
		return nil, err, "initTest"
	}

	ctx,err := orm.New_DBContext(*ctxBase)
	if( err != nil){
		return nil, err, "initTest1";
	}
	return ctx, err, "initTest1";
}

func Test1_01(step int, bCheckName bool) ( int, error, string) {

	var dbData = Test1_init();// (orm.DBContextBase, error, string){
	var ctx = dbData.ctx;
	if( ctx != nil ){
	defer ctx.Close()
	}
	//ctx.

	return 0, nil, "";
}