using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using Antlr4.Runtime;
using System.Reflection.Metadata;
using goscanner.Metadata;
using goscanner.ConvCommon;

namespace goscanner.ConvSql
{
    
    // -o -i -h -q="C:\Work\SlotMonitor\SlotGames\server-tools\atomicsql_compiler_tool\go2cs\go2cs-master\src\go2cs\sql-config.cfg" C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers
    //C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers 


    public class Sql_ConfigTranslation : goscanner.ConvCommon.ConfigFile
    { 
        public const string CONV_SOURCE_PATH_SEP =  ";";
        public const string Version = "2.0.3";
        // 2.0.2 - Add GetValue<>() and GetValues<>()

        public Sql_Dialect SqlDialect = new PostgresSql_Dialect();

        public ConvCommon.TemplateItem Templ_GoSqlCompiledFile;
        public ConvCommon.TemplateItem Templ_GoSqlCompiledQuery;
        public ConvCommon.TemplateItem Templ_GoSqlVarDef;
        public ConvCommon.TemplateItem Templ_SubQuery;
        public ConvCommon.TemplateItem Templ_SubQueries;
        //public ConvCommon.TemplateItem Templ_SubQueryVardef;
        //public ConvCommon.TemplateItem Templ_Include_files;

        public List<string> IncludeFiles = new List<string>();
        //public List<string> IncludeTypes = new List<string>();
        //public List<string> IncludeFuncs = new List<string>();
        
        public string SqlOutputFile;
        public string SqlLang;
        public string SCHEMA_DEFS;

        public string OrmDir_AtomicSql;
        //the git of atomicsql
        public string OrmDir_Atomicsql_Git;        

        //file with metadata of atomicsql
        public string OrmDir_Atomicsql_MetadataFile;
        public FolderMetadata OrmDirAtomicsql_DirMetadata;


        public string OrmDir_AtomicSqlFunc;
        //the git of atomicsql
        public string OrmDir_AtomicsqlFunc_Git;        

        //file with metadata of atomicsql
        public string OrmDir_AtomicsqlFunc_MetadataFile;
        public FolderMetadata OrmDirAtomicsqlFunc_DirMetadata;

        
        protected override void setupSLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "SqlOutputFile":   this.SqlOutputFile = content; break;
                case "TypeLang":        this.TypeLang = content; break;
                case "SqlLang":         this.SqlLang = content; break;
                case "SourcePathDir":   this.SourcePathDir = Utils1.getAbsoluteDirPath(content, CONV_SOURCE_PATH_SEP); break;
                case "SourcePathDir2":  this.SourcePathDir2 = Utils1.getAbsoluteDirPath(content, CONV_SOURCE_PATH_SEP); break;

                case "RootProject":     this.RootProject = Utils1.getAbsoluteDirPath(content); break;
                case "GitProject":      this.GitProject = content; break;
                    
                case "SCHEMA_DEFS":     this.SCHEMA_DEFS = content; break;
                case "Class_Base_MODEL":this.Class_Base_MODEL = content; break;

                case "OrmDir_Atomicsql":            this.OrmDir_AtomicSql = Utils1.getAbsoluteDirPath( content ); break;
                case "OrmDir_Atomicsql_Git":        this.OrmDir_Atomicsql_Git = content; break;
                case "OrmDir_Atomicsql_Metadata":   this.OrmDir_Atomicsql_MetadataFile = content; break;

                case "OrmDir_AtomicsqlFunc":        this.OrmDir_AtomicSqlFunc = Utils1.getAbsoluteDirPath( content ); break;
                case "OrmDir_AtomicsqlFunc_Git":    this.OrmDir_AtomicsqlFunc_Git = content; break;
                case "OrmDir_AtomicsqlFunc_Metadata":this.OrmDir_AtomicsqlFunc_MetadataFile = content; break;
                    
                
                default: 
                    Console.WriteLine( $"Not identified token {token1} in single line property ");
                break;
            }

            if( SqlLang == ESqlOutputType.Mssql)
                SqlDialect = new MsSql_Dialect();
            else
            if( SqlLang == ESqlOutputType.Mysql)
                SqlDialect = new MySql_Dialect();
            else
            if( SqlLang == ESqlOutputType.Postgres)
                SqlDialect = new PostgresSql_Dialect();
        }

        protected override void updateProperties()
        {
            if( Templ_GoSqlCompiledFile  == null )
                Console.WriteLine("Templ_GoSqlCompiledFile - not set");

            if( Templ_GoSqlCompiledQuery  == null )
                Console.WriteLine("Templ_GoSqlCompiledQuery - not set");

            if( Templ_GoSqlVarDef  == null )
                Console.WriteLine("Templ_GoSqlVarDef - not set");

            if( Templ_SubQuery  == null )
                Console.WriteLine("Templ_SubQuery - not set");

            if( Templ_SubQueries  == null )
                Console.WriteLine("Templ_SubQueries - not set");

            //if( Templ_SubQueryVardef  == null )
            //    Console.WriteLine("Templ_SubQueryVardef - not set");
            

            if( IncludeFiles  == null )
                Console.WriteLine("Include_files - not set");

            if( OrmDir_AtomicSql  == null )
                Console.WriteLine("OrmDir_Atomicsql - not set");
            else
                Console.WriteLine($"OrmDir_AtomicSql = {OrmDir_AtomicSql}");

            if( OrmDir_Atomicsql_Git  == null )
                Console.WriteLine("OrmDir_Atomicsql_Git - not set");
            else
                Console.WriteLine($"OrmDir_Atomicsql_Git = {OrmDir_Atomicsql_Git}");

            if( OrmDir_AtomicSqlFunc  == null )
                Console.WriteLine("OrmDir_AtomicsqlFunc - not set");

            if( OrmDir_AtomicsqlFunc_Git  == null )
                Console.WriteLine("OrmDir_AtomicsqlFunc_Git - not set");
            
        }
        protected override void setupMLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "Templ_GoSqlCompiledFile":     this.Templ_GoSqlCompiledFile    = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_GoSqlCompiledQuery":    this.Templ_GoSqlCompiledQuery   = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_GoSqlVarDef":           this.Templ_GoSqlVarDef          = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_SubQuery":              this.Templ_SubQuery             = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_SubQueries":            this.Templ_SubQueries           = new ConvCommon.TemplateItem( token1, content ); break;
                //case "Templ_SubQueryVardef":        this.Templ_SubQueryVardef       = new ConvCommon.TemplateItem( token1, content ); break;
                    
                    
                    
                case "Include_files":               this.IncludeFiles               = processExportListItems( content ); break;
                //case "Include_types":               this.IncludeTypes               = processExportListItems( content ); break;
                //case "Include_funcs":               this.IncludeFuncs              = processExportListItems( content ); break;
                    
                default: 
                    Console.WriteLine( $"Not identified token {token1} in multiline declaration ");
                break;
            }
        }


        List<string> getWords(string t1)
        {
            var words = t1.Split(new String[]{" "}, StringSplitOptions.RemoveEmptyEntries );
            for( var i = 0; i < words.Length; i++ ) 
            {   
                words[i] = words[i].Trim();
            }
            return words.ToList();
        }
        /*
        List<TypeInfo> addTypes()
        { 
            
            var types = new List<TypeInfo>();

            foreach( var t in this.IncludeTypes )
            {
                var type = new TypeInfo();

                var t1 = t;
                t1 = t1.Replace("[", " [ ");
                t1 = t1.Replace("]", " ] ");
                var words = getWords(t1);

                if( words.Count == 1 )
                { 
                    type.Name = words[0];
                }else
                {
                    type.Name = words[0];
                }
                types.Add( type );
            }
            return types;
        }
        */
        /*
        List<Signature> addMethods()
        { 
            var meths = new List<Signature>();

            foreach( var meth in this.IncludeFuncs )
            {
                var func = new FunctionInfo();

                var t1 = meth;
                t1 = t1.Replace("(", " ( ");
                t1 = t1.Replace(")", " ) ");
                t1 = t1.Replace(",", " , ");
                
                var words = getWords(t1);
                if( words.Count < 1 )
                {
                    Console.WriteLine($"Section Include_funcs wrong format for '{meth}'");
                    continue;
                }
                var crs = 1;

                func.Signature.Name = words[crs++];
                
                if( words[crs++] != "(") 
                {
                    Console.WriteLine($"Section Include_funcs wrong format for '{meth}'");
                    continue;
                }
                
                var idxEnd = words.IndexOf(")", 2 );
                if( idxEnd == -1 )
                {
                    Console.WriteLine($"Section Include_funcs wrong format for '{meth}'");
                    continue;
                }
                crs = idxEnd++;
                if( crs >= words.Count )
                {
                    Console.WriteLine($"Section Include_funcs wrong format for '{meth}'");
                    continue;
                }

                var returnType = new TypeInfo();

                if( words[crs] == "(")
                { 
                    
                    var list = new List<ParameterInfo>();
                    for(; crs < words.Count; crs ++)
                    {
                        var param = new ParameterInfo { Name = "", Type = returnType };
                        returnType.Name = words[crs++];
                        list.Add(param);

                        if( words[crs++] != ")" )
                            break;

                        if( words[crs++] != "," )
                        {
                            Console.WriteLine($"Section Include_funcs wrong format for '{meth}'");
                            break;
                        }
                    }

                    func.Signature.Signature.Result = list.ToArray();
                }else
                {
                    returnType.Name = words[crs];
                    func.Signature.Signature.Result = new ParameterInfo[]{ new ParameterInfo { Name = "", Type = returnType } };
                }
            }
            return meths;
        }
        */
    }
        

}
