call 0.common_config.cmd

rem del /s /q /f ..\..\..\..\..\*.metadata

%GoScan%  -o -i -h -q=.\sql-config_metadata.cfg %SRCDIR_PATH%\src
pause
