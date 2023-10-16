@echo off


set GoScanPath=..\..\tools\atomicsql_compiler
set GoServerToolPath=..\..\tools\GoServerTool

set PATH=%PATH%;%GoScanPath%
set PATH=%PATH%;%GoServerToolPath%

set GoScan=%GoScanPath%\goscanner.exe
set GoServerTool=%GoServerToolPath%\DBTool.exe

set SRCDIR_PATH=%CD%\..\..\..\library
set GITHUB_URL=github.com/bbitere/atomicsql_golang.git


rem set CONNECTION_STRING=Host=localhost;Username=??;Password=??;Database=test1
rem set SQL_LANG=postgres

set CONNECTION_STRING="server=localhost;user id=??;password=??;persistsecurityinfo=True;database=test1"
set SQL_LANG=mysql




