call 0.common_config.cmd

@echo ------------------------------------------------------
@echo "1. parse golang models -> generate_jsons"
%GoScan% -o -i -h -e=\.gen\.go -z=%CD%\asqlmodels-config.cfg   %SRCDIR_PATH%\src\mymodels


@echo ------------------------------------------------------
@echo "2. parse generated_jsons -> generate sql_scripts"
%GoServerTool% -asql_migration  -sql_lang=%SQL_LANG% -execute_scripts=n -delimeter=@@@@@@@@######@@@@@@^
		-json_dir=%SRCDIR_PATH%\_db_jsons^
		-sql_dir=%SRCDIR_PATH%\_db_migration^
	    -connection_string=%CONNECTION_STRING%

@echo ------------------------------------------------------
@echo "3. apply sql_scripts in database, iterative and only once."
%GoServerTool% -migration_db  -sql_lang=%SQL_LANG% -input_dir="%SRCDIR_PATH%\_db_migration"^
 -out_dir="" -type_out_file=go^
 -connectionString=%CONNECTION_STRING%
  
@echo ------------------------------------------------------ 
@echo "4. create the additional golang models information from database, useful for atomicsql engine"
%GoServerTool% -export_db -config_file="%CD%\cfg_exportdb.cfg" -type_out_file=go
pause

@echo ------------------------------------------------------ 
@echo "5. generate the sql seqs for lambda expression : where() & select()"
%GoScan% -o -i -h -q=.\sql-config.cfg %SRCDIR_PATH%\src
pause
 

