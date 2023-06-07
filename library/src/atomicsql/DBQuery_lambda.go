package atomicsql

//import abi "internal/abi"

/*
func FuncPCABI0(f any) uintptr

// FuncPCABIInternal returns the entry PC of the function f. If f is a
// direct reference of a function, it must be defined as ABIInternal.
// Otherwise it is a compile-time error. If f is not a direct reference
// of a defined function, it assumes that f is a func value. Otherwise
// the behavior is undefined.
//
// Implemented as a compile intrinsic.
func FuncPCABIInternal(f any) uintptr

*/
func test(){

	Where1( func(x int) bool{
		return x > 10
	 } )

}
func Where1( fn func(x int) bool ){


	//var ret = abi.FuncPCABI0( fn );
	//if( ret != 0){
	//	
	//}
	
}