package main

import (
	"fmt"
	///importer "go/importer"

	test1 "github.com/bbitere/atomicsql_golang.git/tests/test1/test_where"
)

func main(){
	
	var counter = 0;

	
	Exec_test( test1.Test1_01, &counter );
	Exec_test( test1.Test1_02N, &counter );
	Exec_test( test1.Test1_02, &counter );
	Exec_test( test1.Test1_03, &counter );
	//orm.Arr_Append( &arrTests, test1.Test1_04 );
	Exec_test( test1.Test1_05, &counter );
	//orm.Arr_Append( &arrTests, test1.Test1_06 );
	//orm.Arr_Append( &arrTests, test1.Test1_07 );
	Exec_test( test1.Test1_08, &counter );
	Exec_test( test1.Test1_09, &counter );
	Exec_test( test1.Test1_10, &counter );

	/*
	for iTest := 0; iTest < len(arrTests); iTest++ {
				
		var fnTest = arrTests[iTest];
		ret, err, msg = fnTest( iTest, false );
		
		//if( ret == -1 ){return;}
		printTest(ret, msg, err, &counter);
	}*/
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