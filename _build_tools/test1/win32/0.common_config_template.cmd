@echo off

set GoScanPath=..\..\tools\atomicsql_compiler
set GoServerToolPath=..\..\tools\GoServerTool

set PATH=%PATH%;%GoScanPath%
set PATH=%PATH%;%GoServerToolPath%

set GoScan=%GoScanPath%\go2cs.exe
set GoServerTool=%GoServerToolPath%\DBTool.exe

set SRCDIR_PATH=%CD%\..\..\..\library
set CONNECTION_STRING=Host=localhost;Username=??;Password=??;Database=test1
set GITHUB_URL=github.com/bbitere/atomicsql_golang.git
rem set GITHUB_URL="github.com/bbitere/atomicsql_golang.git/src/atomicsql"