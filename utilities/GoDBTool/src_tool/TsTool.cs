using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Data.SqlClient;
using System.IO;
using System.Diagnostics;
using System.Threading;


using System.Text.RegularExpressions;

namespace src_tool
{
    public class EOutputFile
    {
        public const string TS = "ts";
    }
    class TsTool
    {
        static string VERSION = "1.0.1";
        public static void Main(String[] args )
        {
            if( args.Length < 2 )
            {
                 Console.WriteLine( "No arguments: AppType Path flags" );
                return;
            }
            Console.WriteLine( $"GoLangDBTool: v{VERSION}" );
            Console.WriteLine( $"" );

            if( args[ 0 ] == "-x")
            {
                WaitDebugger();
                var lst = args.ToList();
                lst.RemoveAt(0);
                args = lst.ToArray();
            }
            
            if( args[ 0 ] == "-asql_migration")
            {   
                // -asql_migration  -sql_lang=postgres -execute_scripts=n -json_dir=C:\Work\SlotMonitor\atomicsql_golang\library\_db_migration  -delimeter=@@@@@@@@######@@@@@@  -connection_string=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=newton_original
                var dictArgs = parseArguments( args, 
                    new String[]{   
                                    "sql_lang",
                                    "json_dir",
                                    "sql_dir",
                                    "connection_string",
                                    "delimeter",
                                    "execute_scripts",
                                  },
                    new String[]{ 
                        "sql_lang; the name of sql dialect:postgres, mysql",
                        "json Directory;the directory where all json defs are stored",
                        "sql_dir Directory;the directory where all sql scritps definitions for migration are stored",
                        "connection string; connection string of db",
                        "delimeter; string that separate the defs of table in json file",
                        "execute the scripts; if execute_scripts=y,n  => execute scripts, else generate the files"
                        });;
                if( dictArgs == null )
                    return;

                var inst = new GenSqlScriptsFromJsonModels();
                inst.GenerateScripts( 
                                dictArgs["sql_lang"],
                                normalizePath(dictArgs[ "json_dir" ]),
                                normalizePath(dictArgs[ "sql_dir" ]),                                
                                dictArgs["connection_string"],
                                dictArgs["delimeter"],
                                dictArgs["execute_scripts"] =="y"
                                );

                //Console.WriteLine( $"Completed  Generate Models for package {package}");
            }else
            if( args[ 0 ] == "-export_db")
            {   
                //WaitDebugger();
                //-export_db  -config_file="..\..\config.txt" -type_out_file=go
                //-export_db -config_file="_cmd\cfg_exportdb.cfg" -type_out_file=go
                var dictArgs = parseArguments( args, 
                    new String[]{ 
                        "config_file", 
                        "type_out_file",
                        "?databasefirst",
                        },
                    new String[]{ 
                        "config file;the config file containing all defs for export DB", 
                        "type of out file; type_of_file = go | ts",
                        "database first; databasefirst = yes|no. if is yes, generate code from reading table of database",
                        });
                if( dictArgs == null )
                    return;

                var inst = new ExtractModelsFromSqlDatabase();
                inst.GenDBModels( 
                                normalizePath(dictArgs[ "config_file" ]),
                                dictArgs[ "type_out_file" ],
                                getVal(dictArgs, "databasefirst") == "yes"
                                );

                //Console.WriteLine( $"Completed  Generate Models for package {package}");
            }else  
            if( args[ 0 ] == "-migration_db")
            {                
                //-migration_db  -input_dir="C:\Work\SlotMonitor\SlotMonitor\GoServer\Database\scripts" -out_models=""  -type_out_file=go -connectionString=Host=localhost;Username=dev_original;Password=XCZ12345678;Database=newton_original
                var dictArgs = parseArguments( args, 
                        new String[]{ 
                            "sql_lang",
                            "connectionString", 
                            "input_dir", 
                            "out_dir", 
                            "type_out_file"},
                        new String[]{ 
                            "sql_lang; the name of sql dialect:postgres, mysql",
                            "connection string; connection string of db",
                            "sql scripts directory;input directory of sql scripts", 
                            "output directory;the output dir for generate same definitions of migration",
                            "type of out file; type_of_file = go | ts"
                            });
                if( dictArgs == null )
                    return;

                var inputDir = dictArgs[ "input_dir" ];
                Console.WriteLine( $"");
                Console.WriteLine( $"Do migration for directory {inputDir}");

                var migration = new MigrationDB();
                if( migration.DoMigration( 
                                dictArgs[ "sql_lang" ],
                                normalizePath(dictArgs[ "connectionString" ]), 
                                normalizePath(dictArgs[ "input_dir" ]),
                                normalizePath(dictArgs[ "out_dir" ]),
                                dictArgs[ "type_out_file" ]
                                ) )
                {
                    Console.WriteLine( $"Completed migration for directory {inputDir}");
                }
            }else
            if( args[ 0 ] == "-read_svc")
            {
                // -read_svc -dlls_path="./OEFWEbSvc/bin"  -out_svc_dll="app\services\svc.gen.ts"  -out_file_models="SvcModels.gen.ts"
                Console.WriteLine( $"Start Generate SVC from {args[ 1 ]} dll");
                var dictArgs = parseArguments( args, 
                                new String[]{ 
                                    "dlls_path", 
                                    "out_svc_dll", 
                                    "out_file_models"
                                    },
                                new String[]{ 
                                    "path of dll; the path with many dlls that must be scaned and exported in svc",
                                    "svc dll;the file with the description of svc", 
                                    "models file; file with the description of all Models used in svc"
                                    });
                if( dictArgs == null )
                    return;

                var inst = new GenerateSvc( );
                inst._GenerateSvc( 
                                normalizePath(dictArgs[ "dlls_path" ]), 
                                normalizePath(dictArgs[ "out_svc_dll" ]),
                                normalizePath(dictArgs[ "out_file_models" ])                                
                                );

                Console.WriteLine( $"Completed Generate SVC");
            }
            else            
            {
                Console.WriteLine("Not Implemented paramenter");
            }

            //Debugger.Break();
        }

        private static string getVal( Dictionary<string,string> dict, string key)
        {
            var val = "";
            if( dict.TryGetValue(key, out val))
                return val;

            return null;
        }
        private static string normalizePath(string path)
        {
            return Utils.normalizePath(path);
        }

        private static Dictionary<string, string> parseArguments(String[]args, String[] arg_names, String[] desc_args)
        {
            var arg0 = args.Length > 0? args[0] : "";

            var dictRet = new Dictionary<string, string>();
            for( var iDef = 0; iDef < arg_names.Length; iDef++)
            {
                var arg_name = arg_names[iDef];
                var bOptional = false;

                if( arg_name.StartsWith("?") ) 
                {
                    arg_name = arg_name.Substring(1);
                    bOptional = true;
                }

                string value = null;
                for( var iArg = 0; iArg < args.Length; iArg++)
                {
                    var prefixArg = $"-{arg_name}=";
                    if( args[iArg].StartsWith( prefixArg ) )
                    { 
                        value = args[iArg].Replace( prefixArg, "" );
                        break;
                    }
                }

                if( value != null ) 
                {
                     dictRet[arg_name] = value;
                }else
                {
                    if( !bOptional)
                    {
                        printUsage(arg_names, desc_args, arg0, arg_name);
                        return null;
                    }
                }
            }
            return dictRet;
        }
        private static void printUsage(String[] arg_names, String[] desc_args, string cmd, string arg_with_error )
        { 
            Console.WriteLine("");
            Console.WriteLine("");

            if( arg_with_error != null )
            {
                Console.WriteLine($"argument '{arg_with_error}' is with error ");
                Console.WriteLine("");
            }


            var s = $"Usage: {cmd} ";
            var descr_txt = "Description: \n\r";

            for( var iArg = 0; iArg < arg_names.Length; iArg++)
            {
                var desc = desc_args[iArg];
                var parts = desc.Split(';');
                var value_info = parts[0];
                var descr_info = parts[1];

                s += $"-{arg_names[iArg]}=<{value_info}> ";

                descr_txt +=  $"-{arg_names[iArg]}: {descr_info}\r\n";
            }
            Console.WriteLine(s);
            Console.WriteLine("");
            Console.WriteLine(descr_txt);
        }

        public static void WaitDebugger()
        {
            for( var i = 0; i < 100000000; i++)
            {
                Thread.Sleep(100);
            }
        }
    }
}
