call 0.common_config.cmd

set GoScanPath=..\..\..\..\..\utilities\goscanner\src\goscanner\bin\SqlConv\net7.0
set PATH=%PATH%;%GoScanPath%
set GoScan=%GoScanPath%\goscanner.exe

del /s /q /f ..\..\..\..\..\*.metadata

%GoScan%   -o -i -h -q=.\sql-config.cfg %SRCDIR_PATH%\src
pause
