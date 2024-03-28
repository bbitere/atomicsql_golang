package main

import (
	"fmt"
	"os"

	//"sync"
	//"time"

	///importer "go/importer"

	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
	test1_crud "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_crud"
	test1_subquery "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_subquery"

	//test1 "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"
	test1_where "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where"

	test1_where_nosql "github.com/bbitere/atomicsql_golang.git/tests/test1/src/test_where_nosql"

	log "log"
	"runtime/debug"
	"strings"
)

type TestsResult struct {
	Succeded int
	Totals   int
	Crashed  int
	Failed   int
}

var testsResult TestsResult

func main() {

	test1M()
	var counter = 0

	defer logPanic();

	Exec_test(test1_where_nosql.Test1_00, &counter)
	Exec_test(test1_where_nosql.Test1_01, &counter)
	Exec_test(test1_where_nosql.Test1_02, &counter)
	Exec_test(test1_where_nosql.Test1_04, &counter)	
	Exec_test(test1_where_nosql.Test1_05, &counter)
	Exec_test(test1_where_nosql.Test1_03, &counter)

	Exec_test(test1_where_nosql.Test1_08, &counter)

	//--------------------------------------------------

	Exec_test(test1_where.Test1_00, &counter)
	Exec_test(test1_where.Test1_01, &counter)
	Exec_test(test1_where.Test1_02N, &counter)
	Exec_test(test1_where.Test1_02, &counter)
	Exec_test(test1_where.Test1_03, &counter)
	//Exec_test( test1.Test1_0, &counter );
	Exec_test(test1_where.Test1_05, &counter)
	//orm.Arr_Append( &arrTests, test1.Test1_06 );
	//orm.Arr_Append( &arrTests, test1.Test1_07 );
	Exec_test(test1_where.Test1_08, &counter)
	Exec_test(test1_where.Test1_09, &counter)
	Exec_test(test1_where.Test1_10, &counter)

	Exec_test(test1_crud.Tst_Example_CreateUser, &counter)
	Exec_test(test1_crud.Tst_Example_Create2Users, &counter)
	Exec_test(test1_crud.Tst_Example_RetrieveUser, &counter)
	Exec_test(test1_crud.Tst_Example_RetrieveUsers, &counter)

	Exec_test(test1_crud.Tst_Example_DeleteUser, &counter)
	Exec_test(test1_crud.Tst_Example_DeleteUsers, &counter)

	Exec_test(test1_crud.Tst_Example_UpdateUser, &counter)
	Exec_test(test1_crud.Tst_Example_UpdateUsers, &counter)

	Exec_test(test1_crud.Tst_Example_CreateUserRelation, &counter)
	Exec_test(test1_where.TestMisc_01, &counter)
	Exec_test(test1_where.Test1Rtm_10, &counter)
	Exec_test(test1_where.Test1_11, &counter)
	Exec_test(test1_where.Test1_12, &counter)
	Exec_test(test1_where.Test1_13, &counter)

	//printResults();
	Exec_test( test1_subquery.Test1_01, &counter );
	Exec_test( test1_subquery.Test1_02, &counter );

	printResults();
	
}

func Exec_test(fnTest test1_where.TestFunc, pIndex *int) {

	(*pIndex)++
	var ret = 0
	var err error
	var msg = ""
	ret, msg, err = fnTest(*pIndex, false)

	printTest(ret, msg, err, *pIndex)
}

func printResults() {

	fmt.Println("**************************")
	fmt.Printf("Tests Results: %d succeded/ %d totals", testsResult.Succeded, testsResult.Totals)
	fmt.Println("")
}

func printTest(codeSucceded int, testName string, err error, idx int) {

	if(err != nil){

	}
	testsResult.Totals++

	//colorRed := "\033[31m"
	//colorGreen := "\033[32m"
	//colorYellow := "\033[33m"
	//colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	//colorWhite := "\033[37m"

	if codeSucceded == -1 {
		testsResult.Crashed++
		fmt.Println(string(colorCyan), fmt.Errorf("test %d crash -> %s", idx, testName))
		fmt.Println("")
	} else if codeSucceded == 0 {

		testsResult.Failed++
		fmt.Println(string(colorPurple), fmt.Errorf("test %d failed -> %s", idx, testName))
		fmt.Println("")
	} else {
		testsResult.Succeded++
		fmt.Printf("Test %d passed -> %s", idx, testName)
		fmt.Println("")
	}
}

func logPanic() {

	file, err1 := os.OpenFile("logfile.txt", os.O_APPEND, os.ModeAppend)
	if err1 != nil {
		log.Fatal("Cannot create log file: ", err1)
	}
	defer file.Close()
	log.SetOutput(file)

	if err := recover(); err != nil {
		if err != nil {

			var linesText = string(debug.Stack())
			var lines = strings.Split(linesText, "\n")

			for i := 0; i < len(lines) && i <= 6; i++ {
				atmsql.Arr_RemoveAt(&lines, 0)
			}
			var linesCleaned = strings.Join(lines, "\n")
			log.Printf("panic occurred: %v %s", err, linesCleaned)
		}
		log.SetOutput(os.Stdout)
		log.Printf("panic occurred: write error in logfile.txt")
	}
}

/*
	type ChannelData struct{
		data string
	}

	type Channel struct{
		ch chan ChannelData
	}

func new_ChannelAcquire(permits int) *Channel{

		return &Channel{
			ch : make( chan ChannelData, permits),
		}
	}

func (This *Channel) Acquire(){

		This.ch <- ChannelData{ data: "data1"}
	}

func (This *Channel) Release(){

		<- This.ch
	}

func (This *Channel) Send(msg string){

		This.ch <- ChannelData{ data: msg}
	}

func WgDone(wg* sync.WaitGroup, id int){

		var msg = fmt.Sprintf("routine %d - decrease counter ", id );
		fmt.Println(msg);
		wg.Done();
	}

var wg sync.WaitGroup;
var mutex1 sync.Mutex

func taskRoutine( id int, channelComm *Channel) {

		defer WgDone( &wg, id );
		var data = <- channelComm.ch
		if( data.data != "" ){

			var msg1 = fmt.Sprintf("routine %s has msg", data.data );
			fmt.Println(msg1);
		}
		mutex1.Lock()
		{
			var msg1 = fmt.Sprintf("routine %d has the permission", id );
			fmt.Println(msg1);

			time.Sleep( time.Second );
			var msg2 = fmt.Sprintf("routine %d realease the permission", id );
			fmt.Println(msg2);
		}
		mutex1.Unlock()
	}

func test1M(){

		test2M();
		test3M();
	}

func test3M(){

	var channelsComm = []*Channel{};
	for i := 0; i < 3; i++{

		wg.Add( i )
		var ch = new_ChannelAcquire(0);
		atmsql.Arr_Append( &channelsComm, ch );

		go taskRoutine( i, ch );

	}
	time.Sleep( 2*time.Second );
	for i := 0; i < 3; i++{

		channelsComm[i].Send("data1");
	}

	//semaphore.Release()
	wg.Wait();

}

func test2M(){

		var channelComm = new_ChannelAcquire(0)

		go func(){
			time.Sleep( 2*time.Second);
			close( channelComm.ch )
		}()

		fmt.Println("wait");
		<- channelComm.ch
		fmt.Println("close");
	}

/
*/
func test1M() {}

//*/
