call 0.common_config.cmd


  
@echo ------------------------------------------------------ 
@echo "4. create the additional golang models information from database, useful for atomicsql engine"
%GoServerTool%  -export_db -config_file="%CD%\cfg_exportdb_DBFirst.cfg" -type_out_file=go  -databasefirst=yes
pause

@echo ------------------------------------------------------ 
@echo "5. generate the sql seqs for lambda expression : where() & select()"
%GoScan% -o -i -h -q=.\sql-config.cfg %SRCDIR_PATH%\src
pause
 

