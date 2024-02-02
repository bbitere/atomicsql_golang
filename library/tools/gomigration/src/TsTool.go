package main

import (
	"fmt"
	"os"

	//"path/filepath"
	//"regexp"

	"strings"
	"time"

	utils "github.com/bbitere/atomicsql_golang.git/tools/gomigration/src/utils"
)

const (
	VERSION = "1.0.1"
)


func main() {

    var args = os.Args;

	utils.Arr_RemoveAt( &args, 0);
	
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

	if args[0] == "-asql_migration" {

		dictArgs := GoTool_parseArguments(args, []string{
			"sql_lang",
			"json_dir",
			"sql_dir",
			"connection_string",
			"delimeter",
			"execute_scripts",
		}, []string{
			"sql_lang; the name of sql dialect:postgres, mysql",
			"json Directory;the directory where all json defs are stored",
			"sql_dir Directory;the directory where all sql scritps definitions for migration are stored",
			"connection string; connection string of db",
			"delimeter; string that separates the defs of table in json file",
			"execute the scripts; if execute_scripts=y,n => execute scripts, else generate the files",
		})
		if dictArgs == nil {
			return
		}
	
		inst := &GenSqlScriptsFromJsonModels{}
		inst.GenerateScripts(
			dictArgs["sql_lang"],
			normalizePath(dictArgs["json_dir"]),
			normalizePath(dictArgs["sql_dir"]),
			dictArgs["connection_string"],
			dictArgs["delimeter"],
			dictArgs["execute_scripts"] == "y",
		)
	}else
	if args[0] == "-migration_db" {

		// -migration_db -input_dir="C:\Work\SlotMonitor\SlotMonitor\GoServer\Database\scripts" -out_models="" -type_out_file=go -connectionString=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=newton_original
		dictArgs := GoTool_parseArguments(args, []string{
			"sql_lang",
			"connectionString",
			"input_dir",
			"out_dir",
			"type_out_file",
		}, []string{
			"sql_lang; the name of sql dialect:postgres, mysql",
			"connection string; connection string of db",
			"sql scripts directory;input directory of sql scripts",
			"output directory;the output dir for generate same definitions of migration",
			"type of out file; type_of_file = go | ts",
		})
		if dictArgs == nil {
			return
		}

		inputDir := dictArgs["input_dir"]

		fmt.Println("")
		fmt.Printf("Do migration for directory %s\n", inputDir)

		migration := &MigrationDB{}

		if migration.DoMigration(
			dictArgs["sql_lang"],
			normalizePath(dictArgs["connectionString"]),
			normalizePath(dictArgs["input_dir"]),
			normalizePath(dictArgs["out_dir"]),
			dictArgs["type_out_file"],
		) {
			fmt.Printf("Completed migration for directory %s\n", inputDir)
		}


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

	fmt.Println("")
	fmt.Println("")

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
	return utils.Utils_NormalizePath(path)
}

