package main

import (
	//"fmt"
	///importer "go/importer"

	orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
)

func main(){
	
	orm.Str_SubString("", 0, 0); 

	/*
	pkg, err := importer.Default().Import("github.com/bbitere/atomicsql_golang");
	if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return
    }

    for _, declName := range pkg.Scope().Names() {
        fmt.Println(declName)

		var obj = pkg.Scope().Lookup(declName);
		fmt.Println(obj.Name());
		var type1 = obj.Type();

		fmt.Println( type1.String() );
		
    }*/
}