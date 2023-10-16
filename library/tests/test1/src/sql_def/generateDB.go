package main

import (
	"fmt"
	///importer "go/importer"

	test1_crud "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_crud"
	test1 "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
	test1_where "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
)

type TestsResult struct{

	Succeded 	int
	Totals	 	int
	Crashed 	int 
	Failed 		int 
}
var testsResult TestsResult

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
	

	Exec_test( test1_crud.Tst_Example_CreateUser, &counter );
	Exec_test( test1_crud.Tst_Example_Create2Users, &counter );
	Exec_test( test1_crud.Tst_Example_RetrieveUser, &counter );
	Exec_test( test1_crud.Tst_Example_RetrieveUsers, &counter );

	Exec_test( test1_crud.Tst_Example_DeleteUser, &counter );
	Exec_test( test1_crud.Tst_Example_DeleteUsers, &counter );

	Exec_test( test1_crud.Tst_Example_UpdateUser, &counter );
	Exec_test( test1_crud.Tst_Example_UpdateUsers, &counter );

	Exec_test( test1_crud.Tst_Example_CreateUserRelation, &counter );


	Exec_test( test1_where.TestMisc_01, &counter );

	Exec_test( test1_where.Test1Rtm_10, &counter );

	Exec_test( test1_where.Test1_11, &counter );
	Exec_test( test1_where.Test1_12, &counter );
	

	

	printResults();
	
}

func Exec_test(fnTest test1.TestFunc, pIndex *int){

	(*pIndex)++;
	var ret = 0;
	var err error;
	var msg = "";
	ret, err, msg = fnTest( *pIndex, false );
			
	printTest(ret, msg, err, *pIndex);
}

func printResults(){

	fmt.Println("**************************");
	fmt.Printf("Tests Results: %d succeded/ %d totals", testsResult.Succeded, testsResult.Totals)
	fmt.Println("");
}

func printTest(codeSucceded int, testName string, err error, idx int){
		
	testsResult.Totals++;
	
	//colorRed := "\033[31m"
    //colorGreen := "\033[32m"
    //colorYellow := "\033[33m"
    //colorBlue := "\033[34m"
    colorPurple := "\033[35m"
    colorCyan := "\033[36m"
    //colorWhite := "\033[37m"
	
	if( codeSucceded == -1 ){
		testsResult.Crashed++;		
		fmt.Println( string( colorCyan), fmt.Errorf("Test %d crash -> %s", idx, testName));
		fmt.Println("");		
	}else
	if( codeSucceded == 0 ){

		testsResult.Failed++;
		fmt.Println( string( colorPurple ), fmt.Errorf("Test %d failed -> %s", idx, testName));
		fmt.Println("");		
	}else {
		testsResult.Succeded++;
		fmt.Printf("Test %d passed -> %s", idx, testName)
		fmt.Println("");
	}
}