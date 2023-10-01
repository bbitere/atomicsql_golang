package main

import (
	"fmt"
	///importer "go/importer"

	test1_crud "github.com/bbitere/atomicsql_golang.git/tests/test1/test_crud"
	test1 "github.com/bbitere/atomicsql_golang.git/tests/test1/test_where"
)

func main(){
	
	var counter = 0;

	
	Exec_test( test1.Test1_01, &counter );
	Exec_test( test1.Test1_02N, &counter );
	Exec_test( test1.Test1_02, &counter );
	Exec_test( test1.Test1_03, &counter );
	//Exec_test( test1.Test1_0, &counter );
	Exec_test( test1.Test1_05, &counter );
	//orm.Arr_Append( &arrTests, test1.Test1_06 );
	//orm.Arr_Append( &arrTests, test1.Test1_07 );
	Exec_test( test1.Test1_08, &counter );
	Exec_test( test1.Test1_09, &counter );
	Exec_test( test1.Test1_10, &counter );

	Exec_test( test1.Test1_10, &counter );

	Exec_test( test1_crud.Tst_Example_CreateUser, &counter );
	Exec_test( test1_crud.Tst_Example_Create2Users, &counter );
	Exec_test( test1_crud.Tst_Example_RetrieveUser, &counter );
	Exec_test( test1_crud.Tst_Example_RetrieveUsers, &counter );
	
}

func Exec_test(fnTest test1.TestFunc, pIndex *int){

	(*pIndex)++;
	var ret = 0;
	var err error;
	var msg = "";
	ret, err, msg = fnTest( *pIndex, false );
			
	printTest(ret, msg, err, *pIndex);
}

func printTest(codeSucceded int, testName string, err error, idx int){
		

	if( codeSucceded == -1 ){
		fmt.Printf("Test %d crash -> %s", idx, testName)
		fmt.Println("");		
	}else
	if( codeSucceded == 0 ){
		fmt.Printf("Test %d failed -> %s", idx, testName)
		fmt.Println("");		
	}else {
		fmt.Printf("Test %d passed -> %s", idx, testName)
		fmt.Println("");
	}
}