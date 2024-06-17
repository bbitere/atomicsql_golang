call 0.common_config.cmd

@echo ------------------------------------------------------
@echo "1. parse golang models -> generate_jsons"
%GoScan% -o -i -h -e=\.gen\.go -z=%CD%\asqlmodels-config.cfg   %SRCDIR_PATH%\src\mymodels

  
@echo ------------------------------------------------------ 
@echo "4. create the additional golang models information from database, useful for atomicsql engine"
%GoServerTool%   -export_db -config_file="%CD%\cfg_exportdb_ModelFirst_NSql.cfg" -type_out_file=go
pause

@echo ------------------------------------------------------ 
@echo "5. generate the sql seqs for lambda expression : where() & select()"
%GoScan%   -o -i -h -q=.\sql-config_NoSql.cfg %SRCDIR_PATH%\src
pause
 

