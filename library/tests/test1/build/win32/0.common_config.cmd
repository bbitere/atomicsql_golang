@echo off


set GoScanPath=C:\Work\SlotMonitor\atomicsql_golang\utilities\goscanner\src\goscanner\bin\Debug\net7.0
set GoServerToolPath=C:\Work\SlotMonitor\atomicsql_golang\utilities\GoDBTool\bin\Debug

set PATH=%PATH%;%GoScanPath%
set PATH=%PATH%;%GoServerToolPath%

set GoScan=%GoScanPath%\goscanner.exe
set GoServerTool=%GoServerToolPath%\DBTool.exe

set SRCDIR_PATH_LIBRARY=%CD%\..\..\..\..\..\library
set SRCDIR_PATH=%CD%\..\..

rem set SQL_LANG=mysql
rem set CONNECTION_STRING=Host=localhost;Username=??;Password=??;Database=test1
set SQL_LANG=postgres
set CONNECTION_STRING=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=test1
set GITHUB_URL=github.com/bbitere/atomicsql_golang.git
rem set GITHUB_URL="github.com/bbitere/atomicsql_golang.git/src/atomicsql"
echo SRCDIR_PATH_LIBRARY = %SRCDIR_PATH_LIBRARY%