call 0.common_config.cmd

@echo ------------------------------------------------------
%GoScan% -o -i -h -e=\.gen\.go -z=%CD%\asqlmodels-config.cfg   %SRCDIR_PATH%\tests\test1\mymodels
pause


@echo ------------------------------------------------------
%GoServerTool% -asql_migration  -sql_lang=%SQL_LANG% -execute_scripts=n -json_dir=%SRCDIR_PATH%\_db_migration -delimeter=@@@@@@@@######@@@@@@^
	   -connection_string=%CONNECTION_STRING%
pause

