using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Data;
using System.Data.SqlClient;
using System.Drawing;
using System.Text.RegularExpressions;

using Npgsql;
using Npgsql.Logging;
using System.IO;

namespace src_tool
{
    public class MigrationDB
    {
        //2023.05.12.122301.name.sql
        public const string RECOGNIZE_FILE_PATTERN=@"^\d{4}\.\d{2}\.\d{2}\.(\d{2}|\d{3}|\d{4}|\d{6})\.(.+)$";
        public const string TABLE_MIGRATION = "DBParams";

        public const string ParamName = "ParamName";
        public const string ParamValue = "ParamValue";
        //const string ParamType = "ParamType";

        public const string SQL_FILE_APPLIED = "SQL_APPL";
        public const string PROP_HISTORY = "version";


        public class DialectArg: GenericDialectArg
        {
            public string input_dir;
            public string out_dir;
            public string type_out_file;

            public DialectArg(
                string strConnection,
                string input_dir, 
                string out_dir, 
                string type_out_file):base(strConnection)
            {
                this.input_dir = input_dir;
                this.out_dir = out_dir;
                this.type_out_file = type_out_file;
            }
        }

        public bool DoMigration( 
                string SqlLang,
                string connection_string,
                string input_dir,
                string out_dir,
                string type_out_file
                )
        {
            var arg = new DialectArg( connection_string, input_dir, out_dir, type_out_file);
            GenericDialect dialect = GenericDialect.GetDialectByName(SqlLang);
            if( dialect == null)
                return false;

            dialect.fnProcessData = ProcessData;
            return dialect.startConnection( arg );
        }
        /*
        void ProcessData(GenericDialect dialect, GenericDialectArg arg1 )

            NpgsqlConnection cnn = null;
            try
            {
                var connectionString = "Host=localhost;Username=dev_user;Password=Xz1234567890.;Database=ntw_original";
                if ( connString == "" || connString == null)
                    connString = connectionString;
                cnn = new NpgsqlConnection(connString);

                cnn.Open();
                
                Console.WriteLine("Begin Transaction");
                if(!false)
                {
                    doInternalMigration(cnn, input_dir);
                    return true;
                }else
                {
                    using(var transaction = cnn.BeginTransaction())
                    {
                        try
                        {
                            doInternalMigration(cnn, input_dir);
                            transaction.Commit();
                            return true;
                        }catch( Exception e )
                        {
                            transaction.Rollback();

                            Console.WriteLine(e.InnerException != null? e.InnerException.Message: e.Message);
                            Console.WriteLine("Transaction is Canceled!!");

                            if( cnn != null)
                                cnn.Close();
                            return false;
                        }
                    }                
                }
            }catch(Exception e )
            {
                Console.WriteLine(e.Message);
                return false;
            }finally
            {
                if( cnn != null)
                    cnn.Close();
            }
        }
        */
        void ProcessData(GenericDialect dialect, GenericDialectArg arg1 )
        //private void doInternalMigration(NpgsqlConnection cnn, string input_dir)
        {
            var arg = arg1 as DialectArg;
            var history = dialect.getProperty( PROP_HISTORY, TABLE_MIGRATION, ParamName, ParamValue );
            if( history == null )
            {
                setupMigration( dialect );
            }

            var files = scanFiles(arg.input_dir);
            var sorted_files = files.OrderBy(x=>x).ToList();
            applyImports( dialect, sorted_files.ToArray() );
            
            Console.WriteLine("End Transaction");
        }

        private List<String> scanFiles(string dir)
        {
            var listFiles = new List<String>();
            
            var regex = new Regex(RECOGNIZE_FILE_PATTERN);
            
            var allFiles = Directory.GetFiles( dir, "*" );
            foreach( var file in allFiles )
            {
                var fileInfo = new FileInfo(file);
                if( fileInfo.Extension == ".sql" 
                 && regex.IsMatch(fileInfo.Name) )
                {
                    listFiles.Add( file );
                }
            }
            

            string[] subfolders = Directory.GetDirectories(dir);
            foreach( var subdir in subfolders )
            {
                var files = scanFiles(subdir);
                listFiles.AddRange( files );
            }
            return listFiles;
        }

        private void applyImports( GenericDialect dialect, string[] files )
        {
            foreach( var file in files)
            {
                var content = File.ReadAllText( file );
                ApplyImport( dialect, file, content );
            }
        }

        public void ApplyImport( GenericDialect dialect, string file, string content )
        {
            //foreach( var file in files)
            {
                var fileInfo = new FileInfo(file);
                var property = fileInfo.Name.Replace( fileInfo.Extension, "");

                var fileApplied = dialect.getProperty( property, TABLE_MIGRATION, ParamName, ParamValue);
                if( fileApplied == null )
                {
                    Console.WriteLine( $"Start script: {property}");
                    {
                        dialect.execScript( content);

                        dialect.insertProperty( property, SQL_FILE_APPLIED, TABLE_MIGRATION, ParamName, ParamValue );

                        var historyPropValue = getHistoryName(property);
                        dialect.updateProperty( PROP_HISTORY, historyPropValue, TABLE_MIGRATION, ParamName, ParamValue );
                    }
                    Console.WriteLine( $"End   script: {property}");
                }
            }
        }

        private string getHistoryName(string property)
        {
            var parts = property.Split('.').ToList();
            parts.RemoveAt( parts.Count - 1 );
            var prop = String.Join(".", parts);
            return prop;
        }

        public static DbTable createMigrationTable(GenericDialect dialect)
        {
            var isNullable = false;
            var tableMigration = new DbTable();
            tableMigration.columns.Add( new DbColumn().initSqlPrimary("ID") );
            tableMigration.columns.Add( new DbColumn().initSql(ParamName,  dialect.getSqlType("string", ref isNullable,""), isNullable) );
            tableMigration.columns.Add( new DbColumn().initSql(ParamValue, dialect.getSqlType("string", ref isNullable,""), isNullable) );
            //tableMigration.columns.Add( new DbColumn().initSql(ParamType,  dialect.getSqlType("string", ref isNullable), isNullable) );
            tableMigration.initSql( TABLE_MIGRATION, tableMigration.columns[0] );

            return tableMigration;
        }
        

        private void setupMigration( GenericDialect dialect )
        {
            /*
            var requestText = $@"

            CREATE TABLE IF NOT EXISTS public.""{TABLE_MIGRATION}""
            (
                ""ID"" SERIAL NOT NULL,
                ""{ParamName}"" character   varying(255)  COLLATE pg_catalog.""default"" NOT NULL DEFAULT ''::character varying,
                ""{ParamValue}"" character  varying(1024) COLLATE pg_catalog.""default"" NOT NULL DEFAULT ''::character varying,
                ""{ParamType}"" character   varying(255)  COLLATE pg_catalog.""default"" NOT NULL DEFAULT ''::character varying,
                CONSTRAINT ""{TABLE_MIGRATION}_pkey"" PRIMARY KEY (""ID"")
            )
            TABLESPACE pg_default;

            ALTER TABLE IF EXISTS public.""{TABLE_MIGRATION}"" OWNER to postgres;
            ";
            dialect.execScript( requestText);
            */
            var tableMigration = createMigrationTable( dialect );

            var scriptTable = dialect.addTable(tableMigration);
            dialect.execScript( scriptTable );

            dialect.insertProperty( PROP_HISTORY, "0000.00.00", TABLE_MIGRATION, ParamName, ParamValue );
        }
        
    }
}
