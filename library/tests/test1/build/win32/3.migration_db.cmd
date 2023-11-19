call 0.common_config.cmd

rem @echo VAR: %GoServerTool%
@echo ------------------------------------------------------
%GoServerTool%  -migration_db  -sql_lang=%SQL_LANG% -input_dir="%SRCDIR_PATH%\_db_migration"^
 -out_dir="" -type_out_file=go^
 -connectionString=%CONNECTION_STRING%
 pause
 
 @echo ------------------------------------------------------
 rem create the models from Database
 %GoServerTool% -x -export_db -config_file="%CD%\cfg_exportdb.cfg" -type_out_file=go
 pause
 