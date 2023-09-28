@echo off

set GoScanPath=C:\Work\SlotMonitor\SlotGames\server-tools\atomicsql_compiler_tool\go2cs\go2cs-master\src\go2cs\bin\Sqlconv\net7.0\
set GoServerToolPath=C:\Work\SlotMonitor\SlotGames\server-tools\GoServerTool\bin\Debug
C:\Work\SlotMonitor\SlotGames\server-tools\GoServerTool\bin\Debug
set PATH=%PATH%;%GoScanPath%
set PATH=%PATH%;%GoServerToolPath%

set GoScan=%GoScanPath%\go2cs.exe
set GoServerTool=%GoServerToolPath%\DBTool.exe

set SRCDIR_PATH=%CD%\..\..\..\library
set CONNECTION_STRING=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=test1
set GITHUB_URL=github.com/bbitere/atomicsql_golang.git
rem set GITHUB_URL="github.com/bbitere/atomicsql_golang.git/src/atomicsql"
