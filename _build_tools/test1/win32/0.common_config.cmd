@echo off

set PATH=%PATH%;C:\Work\SlotMonitor\SlotGames\server-tools\atomicsql_compiler_tool\go2cs\go2cs-master\src\go2cs\bin\Debug\net7.0\
set PATH=%PATH%;C:\Work\SlotMonitor\SlotGames\server-tools\GoServerTool\bin\Debug\

set GoScan=C:\Work\SlotMonitor\SlotGames\server-tools\atomicsql_compiler_tool\go2cs\go2cs-master\src\go2cs\bin\Debug\net7.0\go2cs.exe
set GoServerTool=C:\Work\SlotMonitor\SlotGames\server-tools\GoServerTool\bin\Debug\DBTool.exe

set SRCDIR_PATH=%CD%\..\..\..\library
set CONNECTION_STRING=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=test1
set GITHUB_URL=github.com/bbitere/atomicsql_golang.git
rem set GITHUB_URL="github.com/bbitere/atomicsql_golang.git/src/atomicsql"
