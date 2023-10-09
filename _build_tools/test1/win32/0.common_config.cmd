@echo off

set GoScanPath=..\..\tools\atomicsql_compiler
set GoServerToolPath=..\..\tools\GoServerTool

set PATH=%PATH%;%GoScanPath%
set PATH=%PATH%;%GoServerToolPath%

set GoScan=%GoScanPath%\go2cs.exe
set GoServerTool=%GoServerToolPath%\DBTool.exe

set SRCDIR_PATH=%CD%\..\..\..\library
rem set CONNECTION_STRING=Host=localhost;Username=?;Password=?;Database=test1
rem set SQL_LANG=postgres
set CONNECTION_STRING=Host=localhost;Username=rbitere;Password=xz021621;Database=test1
set SQL_LANG=mysql

set GITHUB_URL=github.com/bbitere/atomicsql_golang.git
rem set GITHUB_URL="github.com/bbitere/atomicsql_golang.git/src/atomicsql"