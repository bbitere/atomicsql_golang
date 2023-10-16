call 0.common_config.cmd

@echo ------------------------------------------------------
%GoScan% -o -i -h -e=\.gen\.go -z=%CD%\asqlmodels-config.cfg   %SRCDIR_PATH%\src\mymodels
pause


@echo ------------------------------------------------------
%GoServerTool% -asql_migration  -sql_lang=%SQL_LANG% -execute_scripts=n -delimeter=@@@@@@@@######@@@@@@^
		-json_dir=%SRCDIR_PATH%\_db_jsons^
		-sql_dir=%SRCDIR_PATH%\_db_migration^
	    -connection_string=%CONNECTION_STRING%
pause

