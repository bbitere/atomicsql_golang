call 0.common_config.cmd

@echo ------------------------------------------------------
%GoScan% -o -i -h -e=\.gen\.go -z=%CD%\asqlmodels-config.cfg   %SRCDIR_PATH%\tests\test1\mymodels


@echo ------------------------------------------------------
%GoServerTool% -asql_migration  -sql_lang=postgres -execute_scripts=n -json_dir=%SRCDIR_PATH%\_db_migration -delimeter=@@@@@@@@######@@@@@@^
	   -connection_string=%CONNECTION_STRING%

@echo ------------------------------------------------------
%GoServerTool% -migration_db  -sql_lang=postgres -input_dir="%SRCDIR_PATH%\_db_migration"^
 -out_dir="" -type_out_file=go^
 -connectionString=%CONNECTION_STRING%
  
@echo ------------------------------------------------------ 
 rem create the models from Database
 %GoServerTool% -export_db -config_file="%CD%\cfg_exportdb.cfg" -type_out_file=go
 pause
 

