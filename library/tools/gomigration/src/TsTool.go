package main

import (
	"fmt"
	"os"

	//"path/filepath"
	//"regexp"

	"strings"
	"time"
)

const (
	VERSION = "1.0.1"
)




func main() {

    var args = os.Args;
    GoTool_Main(args);
}

func GoTool_Main(args []string ){

    
	if len(args) < 2 {
		fmt.Println("No arguments: AppType Path flags")
		return
	}
	fmt.Printf("GoLangDBTool: v%s\n\n", VERSION)

	if args[0] == "-x" {
		GoTool_WaitDebugger()
		args = args[1:]
	}

	if args[0] == "-migration_db" {
		// ...
	} else {
		fmt.Println("Not Implemented parameter")
	}
}

func GoTool_WaitDebugger() {

	for i := 0; i < 100000000; i++ {
		time.Sleep(100 * time.Millisecond)
	}
}

func GoTool_parseArguments(args []string, argNames []string, descArgs []string) map[string]string {
	arg0 := ""
	if len(args) > 0 {
		arg0 = args[0]
	}

	dictRet := make(map[string]string)
	for iDef := 0; iDef < len(argNames); iDef++ {
		argName := argNames[iDef]
		bOptional := false

		if strings.HasPrefix(argName, "?") {
			argName = argName[1:]
			bOptional = true
		}

		var value string
		for iArg := 0; iArg < len(args); iArg++ {
			prefixArg := "-" + argName + "="
			if strings.HasPrefix(args[iArg], prefixArg) {
				value = strings.Replace(args[iArg], prefixArg, "", 1)
				break
			}
		}

		if value != "" {
			dictRet[argName] = value
		} else {
			if !bOptional {
				GoTool_printUsage(argNames, descArgs, arg0, argName)
				return nil
			}
		}
	}
	return dictRet
}

func GoTool_printUsage(argNames []string, descArgs []string, cmd string, argWithError string) {
	fmt.Println("\n\n")

	if argWithError != "" {
		fmt.Printf("argument '%s' is with error \n", argWithError)
		fmt.Println("")
	}

	s := fmt.Sprintf("Usage: %s ", cmd)
	descrTxt := "Description: \n\r"

	for iArg := 0; iArg < len(argNames); iArg++ {
		desc := descArgs[iArg]
		parts := strings.Split(desc, ";")
		valueInfo := parts[0]
		descrInfo := parts[1]

		s += fmt.Sprintf("-%s=<%s> ", argNames[iArg], valueInfo)

		descrTxt += fmt.Sprintf("-%s: %s\r\n", argNames[iArg], descrInfo)
	}
	fmt.Println(s)
	fmt.Println("")
	fmt.Println(descrTxt)
}

func normalizePath(path string) string {
	return Utils_NormalizePath(path)
}

