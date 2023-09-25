rem cd ..

@echo off
set DIR_TOOL=..\..\..\server-tools\atomicsql_compiler_tool\go2cs\go2cs-master\src\go2cs\bin\SqlConv\net7.0
set path=%path%;%CD%\%DIR_TOOL%;

SET SVCTOOL=%DIR_TOOL%\go2cs.exe

rem merge dar doar full path,pt ca directorul curent il modifica in C:\users
%SVCTOOL% -o -i -h -q=.\sql-config.cfg C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers 
pause
