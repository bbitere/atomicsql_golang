call 0.common_config.cmd

del /s /q /f ..\..\..\..\..\*.metadata

%GoScan%   -o -i -h -q=.\sql-config.cfg %SRCDIR_PATH%\src
pause
