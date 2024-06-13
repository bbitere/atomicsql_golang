using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Data;
using System.Data.SqlClient;
using System.Drawing;
using Npgsql;
using Npgsql.Logging;
using System.IO;
using src_tool.templates;
using System.Text.RegularExpressions;

namespace src_tool
{
    public class ExtractModelsFromSqlDatabase
    {
        //public const string GO_EXTENSION = ".gen.go";
        public const string CONCAT_DOT = ".";

        public Dictionary<string, FKRootTgt> dictForeignKeys = new Dictionary<string, FKRootTgt>();
        //public Dictionary<string, DbTable> tables = new Dictionary<string,DbTable>();
        
        public class DialectArg: GenericDialectArg
        {
            public ConfigFile config = new ConfigFile();
            public string type_out_file = "";
            public bool bDatabaseFirst = false;

            public DialectArg(
                ConfigFile config,
                string type_out_file ):base(config.ConnectionString)
            {
                this.config = config;
                this.type_out_file = type_out_file;
            }
        }

        public void GenDBModels( 
                string cfg_file,
                string type_out_file,
                bool bDatabaseFirst
            )
        {
            var config = new ConfigFile();
            config.ParseConfigFile(cfg_file);
        
            var arg = new DialectArg( config, type_out_file);
            GenericDialect dialect = GenericDialect.GetDialectByName(config.SqlLang);
            if( dialect == null)
                return;

            arg.bDatabaseFirst = bDatabaseFirst;

            if( arg.bDatabaseFirst)            
            {
                dialect.fnProcessData = ProcessData;
                dialect.startConnection( arg );
            }else
            {
                dialect.fnProcessData = ProcessData;
                ProcessData( dialect, arg );
            }
        }
        void ProcessData(GenericDialect dialect, GenericDialectArg arg1 )
        {
            var bOk = false;
            var dictTables = new Dictionary<string, DbTable>();
            var arg = arg1 as DialectArg;
            if( arg.bDatabaseFirst)
            {
                dictTables = dialect.readTables( arg.config.SqlLang );
                if( dictTables != null )
                {
                    updateGoLangNameInTablesAndColumns( ref dictTables, arg.config.Delimeter, arg.config.DirJsons, dialect);

                    if( dialect.readConstraintors( dictTables  ))
                    {
                        bOk = true;
                    }
                }
            }else
            {
                updateGoLangNameInTablesAndColumns( ref dictTables, arg.config.Delimeter, arg.config.DirJsons, dialect);
                bOk = true;
            }
            if( bOk ) 
            {
                Console.WriteLine("Generate DB defs");
                if( arg.type_out_file == EOuputLang.GoLang)
                {
                    var config = arg.config;
                    var out_models = config.ModelsOutputDir;
                    var out_file_dbcontext = config.OutputDBContextFile;
                    var import_package_models = config.ImportPackageModels.Trim();
                    var import_package_orm    = config.ImportPackageOrm.Trim();
                    var packageNameNsql1       = config.PackageGenSql.Trim();
            
                    string packageNameModels    = getPackageByImport(import_package_models);
                    //string packageNameOrm       = getPackageByImport(import_package_orm);
                    string packageNameNsql      = getPackageByImport(packageNameNsql1);

                    string tables_DefVar = "";
                    var listTables = Utils.getListFromDict(dictTables);

                    Golang_writeModelsDefs( dialect, listTables, arg.config, out_models, arg.config.Models_Extension,
                        packageNameModels, arg.config.BaseModelName, ref tables_DefVar);

                    Golang_writeDBContext( dialect, listTables, arg.config, out_file_dbcontext,
                        packageNameNsql, import_package_models, 
                        import_package_orm,
                        tables_DefVar );
                }
            }
        }

        private string foundLatestJson(string dir)
        {
            if( dir == "" || dir == null)
                return null;
            var regex = new Regex(MigrationDB.RECOGNIZE_FILE_PATTERN);
            
            var listFiles = new List<string>();
            var allFiles = Directory.GetFiles( dir, "*.json" );
            foreach( var file in allFiles )
            {
                var fileInfo = new FileInfo(file);
                if( fileInfo.Extension == ".json" 
                 && regex.IsMatch(fileInfo.Name) )
                {
                    listFiles.Add( file );
                }
            }
            if( listFiles.Count > 0 )
            {
                listFiles = listFiles.OrderByDescending(x=> x ).ToList();
                return listFiles[0];
            }
            return null;

        }
        private void updateGoLangNameInTablesAndColumns(
            ref Dictionary<string, DbTable> tables, 
            string delimeter,
            string directoryJsons,
            GenericDialect dialect)
        { 
            string file = foundLatestJson( directoryJsons );
            if( file != null )
            {
                var bMustDeleteJsonFile = false;
                var list = GenSqlScriptsFromJsonModels.ReadCustomJson(
                    file, delimeter, dialect, ref bMustDeleteJsonFile );

                //DbTable tableDBParam = null;
                var dictTables = Utils.getDictFromList(list, x=>x.SqlTableNameModel );

                foreach( var it in dictTables )
                {
                    //if( dict.ContainsKey( it.Key ))
                    {
                        var table = it.Value;
                        tables[ it.Key ] = table;

                        table.LangTableNameModel = dictTables[it.Key].LangTableNameModel; 

                        var dictCol = Utils.getDictFromList(dictTables[it.Key].columns, x=>x.sqlName );

                        DbColumn mainColumn = null;
                        foreach( var col in table.columns)
                        { 
                            if( dialect.isNoSql() )
                            {
                                if( col.descriptionTag != null 
                                    
                                 && col.descriptionTag.Contains(GoModelTemplate.MARK_PRIMARY_KEY) )
                                {
                                    mainColumn = col;
                                }
                            }

                            if( dictCol.ContainsKey( col.sqlName))
                            {
                                col.langName  = dictCol[col.sqlName].langName; 
                                col.langName2 = dictCol[col.sqlName].langName2; 
                            }
                        }
                        if( dialect.isNoSql() && table.PrimaryColumn == null && mainColumn != null )
                        {
                            table.PrimaryColumn = mainColumn;
                        }
                    }
                }
                
            }
        }
        /*
        public void GenDBModels1( 
                string cfg_file,
                string type_out_file
            )
        {
            config.ParseConfigFile(cfg_file);

            var import_package_models = config.ImportPackageModels.Trim();
            var import_package_orm    = config.ImportPackageOrm.Trim();
            var packageNameNsql1       = config.PackageGenSql.Trim();
            
            string packageNameModels    = getPackageByImport(import_package_models);
            string packageNameOrm       = getPackageByImport(import_package_orm);
            string packageNameNsql      = getPackageByImport(packageNameNsql1);

            NpgsqlConnection cnn = null;
            try
            {
                var out_models = config.ModelsOutputDir;
                var out_file_dbcontext = config.OutputDBContextFile;
                var connString = config.ConnectionString;
                var connectionString = "Host=localhost;Username=dev_user;Password=Xz1234567890.;Database=ntw_original";
                if ( connString == "" || connString == null)
                    connString = connectionString;
                cnn = new NpgsqlConnection(connString);

                cnn.Open();
                if( readTables( cnn ) 
                 && readConstraintors( cnn  ))
                {
                    Console.WriteLine("Generate DB defs");
                    if( type_out_file == EOuputLang.GoLang)
                    {
                        string tables_DefVar = "";
                        Golang_writeModelsDefs(config, out_models, config.Models_Extension,
                            packageNameModels, config.BaseModelName, ref tables_DefVar);

                        Golang_writeDBContext( config, out_file_dbcontext,
                            packageNameNsql, import_package_models, 
                            import_package_orm,
                            tables_DefVar );
                    }
                }
            }catch(Exception e )
            {
                Console.WriteLine(e.Message);
            }finally
            {
                if( cnn != null)
                    cnn.Close();
            }
        }
        */
        public string getPackageByImport(string imp_package)
        { 
            var parts = imp_package.Split('/');

            return parts[parts.Length-1];
        }


        private List<string> Golang_getModelCol( 
                GenericDialect dialect, ConfigFile config, DbTable table, 
                ref Dictionary<string, string>packageImports)
        {
            var defsModel = new List<string>();
            foreach( var column in table.columns )
            {
                var type = dialect.getGoLangType( column, ref packageImports );

                if( column.ForeignKey != null )
                {
                    var type2 = dialect.getGoLangTypeIntFk( column, ref packageImports);
                    defsModel.Add( GoModelTemplate.GetModelColumn(column.langName,  column.sqlName, type, column.bIsIdentity, true ) );
                    defsModel.Add( GoModelTemplate.GetModelColumn(column.langName2, column.sqlName, type2, column.bIsIdentity, false ) );
                }else
                {
                    defsModel.Add( GoModelTemplate.GetModelColumn(column.langName, column.sqlName, type, column.bIsIdentity, false ) );
                }
            }
            return defsModel;
        }

        string getGoLangType_DefVar( DbColumn column, string fieldName, 
            List< DbTable> listRecursiveStack, string prefix_member)
        {
            if( listRecursiveStack != null )
            {
                var Model_DefVar = Golang_getModelCol_DefVar(column.ForeignKey, $"{prefix_member}{fieldName}{CONCAT_DOT}", listRecursiveStack );

                var txtDefVarModel = String.Join( 
                        @",
                        ",Model_DefVar );

                return $@"{GoDBContext.pkgModels}.{GoModelTemplate.PREF_DEF}{column.ForeignKey.LangTableNameModel} {{
                        {txtDefVarModel}, 
                        }}";
            }else
            {
                return $"\"{prefix_member}{fieldName}\"";
            }
        }
        private List<string> Golang_getModelCol_DefVar(DbTable table, string prefix_member, 
                    List< DbTable> listRecursiveStack)
        {   
            if( listRecursiveStack.Contains(table) )
            {
                return new List<string>();
            }
            //dictRecursiveStack[ table.SqlTableNameModel ] = table;
            listRecursiveStack.Add( table );

            var varDefModel = new List<string>();
            foreach( var column in table.columns )
            {
                if( column.langName == GoModelTemplate.NoSqlID)
                    continue;

                if( column.ForeignKey != null)
                {
                    //define the UserRole_ID
                    var type2 = getGoLangType_DefVar( column, column.sqlName, null, prefix_member);
                    varDefModel.Add( GoModelTemplate.GetModelColumn_DefVar(column.langName2, type2, column.bIsIdentity ) );

                    //define the UserRoleID : UserRole{ ... }
                    //var type = getGoLangType_DefVar( column, column.langName, true, prefix_member);
                    //define the UserRoleID : UserRole{ "UserRole_ID.Member" } //here I put the sql fields

                    //collect ForeignKeys
                    var keyFK = $"{table.SqlTableNameModel}{CONCAT_DOT}{column.sqlName}";
                    //var defFK = $@" ""{column.sqlName}"", ""{column.ForeignKey.sqlTableNameModel}"", ""{column.ForeignKey.PrimaryColumn.sqlName }""";


                    dictForeignKeys[ keyFK ] = new FKRootTgt( table, column, column.ForeignKey );

                    var type = getGoLangType_DefVar( column, column.sqlName, listRecursiveStack, prefix_member);
                    varDefModel.Add( GoModelTemplate.GetModelColumn_DefVar(column.langName, type, column.bIsIdentity ) );
                }else
                {
                    var type = getGoLangType_DefVar( column, column.sqlName, null, prefix_member);
                    varDefModel.Add( GoModelTemplate.GetModelColumn_DefVar(column.langName, type, column.bIsIdentity ) );
                }
            }
            listRecursiveStack.Remove( table );
            return varDefModel;
        }
        private List<string> Golang_getModelCol_Def(DbTable table)
        {
            var TDefModel = new List<string>();
            foreach( var column in table.columns )
            {
                if( column.ForeignKey != null )
                {
                    var type = GoModelTemplate.getGoLangType_Def( column, true );
                    TDefModel.Add( GoModelTemplate.GetModelColumn_Def(column.langName, type, column.bIsIdentity ) );

                    var type2 = GoModelTemplate.getGoLangType_Def( column, false );
                    TDefModel.Add( GoModelTemplate.GetModelColumn_Def(column.langName2, type2, column.bIsIdentity ) );
                }else
                {
                    var type = GoModelTemplate.getGoLangType_Def( column, false );
                    TDefModel.Add( GoModelTemplate.GetModelColumn_Def(column.langName, type, column.bIsIdentity ) );
                }
            }
            return TDefModel;
        }
        private List<string> Golang_getModelInitilized(DbTable table)
        {
            var TDefModel = new List<string>();
            foreach( var column in table.columns )
            {
                if( column.ForeignKey != null )
                {
                    //var defUserRole = T_UserRole{}
				    //model.UserRoleID = (defUserRole.Def().FnNewInst(bFull)).(*UserRole)

                    var Table = column.ForeignKey.LangTableNameModel;
                    var instName = column.langName;
                    var FK_ID = column.langName;
                    var def =$@"
                        var def{instName} = {GoModelTemplate.PREF_DEF}{Table}{{}}
				        model.{FK_ID} = (def{instName}.Def().FnNewInst(bFull)).(*{Table})";

                    TDefModel.Add(def);
                }
            }
            return TDefModel;
        }
        

        private bool Golang_writeModelsDefs(
                GenericDialect dialect,
                List<DbTable> tables,
                ConfigFile config,
                string out_models,
                string extension_file,
                string packageName, 
                string baseModel,
                ref string tables_DefVar)
        {
            foreach( var table in tables)
            {
                if( table.PrimaryColumn == null)
                {
                    dialect.printError($"Primary column from table {table.LangTableNameModel}. Add keyword {GoModelTemplate.MARK_PRIMARY_KEY} in primary column tags");
                }

                var packageImports = new Dictionary<string, string>();
                //var table = iTable;

                var tableLangName = table.LangTableNameModel;
                var tableSqlName = table.SqlTableNameModel;

                if( tableLangName == "Employee1")
                    Utils.Nop();

                var defsModel = Golang_getModelCol( dialect, config, table, ref packageImports);

                var listRecursiveStack = new List< DbTable>();
                var TDefModel   = Golang_getModelCol_Def(table);
                var varDefModel = Golang_getModelCol_DefVar(table, "", listRecursiveStack);
                var ModelInitialize_Def = Golang_getModelInitilized(table);
                
                
                if( table.PrimaryColumn == null )
                {
                    Console.WriteLine($"table {tableSqlName} doesnt have a primary key");
                }else
                {
                    var sep = @",
            ";
                    var defVarTxt  = String.Join( sep, varDefModel.ToArray());
                    tables_DefVar += GoDBContext.GetTemplateVardef(tableLangName, defVarTxt );

                    var text = GoModelTemplate.GetModel(
                        config,
                        packageName, baseModel,
                        tableLangName, tableSqlName,
                        table.PrimaryColumn.langName, 
                        table.PrimaryColumn.langType, 
                        Utils.getListFromDict( packageImports ).ToArray(), 
                        defsModel.ToArray(),
                        TDefModel.ToArray(),
                        ModelInitialize_Def.ToArray()
                        //varDefModel.ToArray()
                        
                        );

                    var filePath = out_models + Path.DirectorySeparatorChar + tableLangName + extension_file;
                    try
                    {
                        File.WriteAllText( filePath, text);
                        Console.WriteLine( $"export file {filePath}" );
                    }catch( Exception e)
                    {
                        Console.WriteLine( e.Message );
                    }
                }
            }
            return true;
        }

        private bool Golang_writeDBContext( 
            GenericDialect dialect,
            List<DbTable> tables,
            ConfigFile config,
            string output_file,
            string packageOrm, 
            string import_package_models,
            string import_package_orm, 
            string tables_DefVar
            )
        {
            var text = GoDBContext.GetTemplateDBContext(
                            dialect,
                            config, packageOrm, 
                            import_package_models, 
                            import_package_orm, 
                            tables_DefVar, 
                            dictForeignKeys,
                            tables );

            var filePath = output_file;
            try
            {
                File.WriteAllText( filePath, text);
                Console.WriteLine( $"export file {filePath}" );
            }catch( Exception e)
            {
                Console.WriteLine( e.Message );
                return false;
            }
            return true;
        }
    }
}
