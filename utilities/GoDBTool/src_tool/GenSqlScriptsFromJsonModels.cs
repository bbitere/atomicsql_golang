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

//-export_db -config_file="_cmd\cfg_exportdb.cfg" -type_out_file=go

namespace src_tool
{
    public class GenSqlScriptsFromJsonModels
    {
        //this are defined in GenerateJsonModels, go2cs
        const string TAG_Type ="Type";
        const string TAG_tags ="tags";
        const string TABLE_SQLNAME ="_tblsql";

        const string EXTENSION_JSON = ".json";
        const string EXTENSION_SQL= ".sql";
        const string PATTERN_NAME_SQL = ".models-";//all files have this in their name
        

        //public const string GO_EXTENSION = ".gen.go";
        public const string CONCAT_DOT = ".";

        public Dictionary<string, FKRootTgt> dictForeignKeys = new Dictionary<string, FKRootTgt>();
        public Dictionary<string, DbTable> tables = new Dictionary<string,DbTable>();

        public class DialectArg: GenericDialectArg
        {
            public string dir_jsons;
            public string dir_sql;
            public string delimeter;
            public bool   bApplyScripts;

            public DialectArg(
                string strConnection,
                string dir_jsons,
                string dir_sql,
                string delimeter,
                bool bApplyScripts):base(strConnection)
            {
                this.dir_jsons = dir_jsons;
                this.dir_sql = dir_sql;
                this.delimeter = delimeter;
                this.bApplyScripts = bApplyScripts;
            }
        }

        public class TDbTable
        { 
            public int counterOrder;
            public string script;
            public string nameDbTable;
            public DbTable tableInst;
            

            //public Dictionary<string, string> dependency = new Dictionary<string, string>();
            public TDbTable(string nameDbTable, string script)
            { 
                this.nameDbTable = nameDbTable;
                this.script = script;
                this.tableInst = null;
            }
            public TDbTable(DbTable tableInst, string script)
            { 
                this.nameDbTable = tableInst.LangTableNameModel;
                this.script = script;
                this.tableInst = tableInst;
            }
        }
        public class TUpdColumn
        { 
            public string script;
            public DbTable tableInst;
            public TUpdColumn(DbTable tableInst, string script)
            { 
                this.tableInst = tableInst;
                this.script = script;
            }
        }
        public class PreGenerateScript
        {
            public List<TDbTable>   listAddTables   = new List<TDbTable>();
            public List<TDbTable>   listDropTables  = new List<TDbTable>();
            public List<TUpdColumn> listUpdTables   = new List<TUpdColumn>();

            public string generateScript(GenericDialect dialect)
            {
                var separator = dialect.SqlSeparator();
                var s = "";
                var arr = new List<string>();
                arr.AddRange( listDropTables.Select(x=>x.script));
                
                if( listAddTables.Count > 0 )
                {
                    var listOrder = new List<TDbTable>();
                    var stack     = new List<DbTable>();
                    var dictAllTDbtables = new Dictionary<string, TDbTable>();
                    foreach( var it in listAddTables)
                    {
                        if( it.tableInst != null )
                            dictAllTDbtables[it.tableInst.LangTableNameModel] = it;
                        else
                            it.counterOrder = 10000;
                    }
                    for( var i = 0; i < 10; i++)
                    {
                        //order by hierarchy dependency
                        foreach( var it in listAddTables)
                        {
                            if( it.tableInst != null )
                            {
                                if( !orderByDependency( it, dictAllTDbtables, stack, 0 ))
                                    return "-- errors. please fix the errors.";
                            }
                        }
                    }
                    listAddTables = listAddTables.OrderByDescending(x=>x.counterOrder).ToList();

                    arr.AddRange( listAddTables.Select(x=>x.script));
                }
                arr.AddRange( listUpdTables.Select(x=>x.script));

                s += string.Join( separator, arr);
                
                return s;
            }

            bool orderByDependency(TDbTable table, Dictionary<string, TDbTable> dictAllTDbtables,
                List<DbTable> stack, int counter )
            {
                if( stack.Contains(table.tableInst))
                {
                    Console.WriteLine( $"cyclic dependency for table {table.tableInst.LangTableNameModel}" );
                    return false;
                }
                var idxStack = stack.Count;
                stack.Add(table.tableInst);

                table.counterOrder = Math.Max( table.counterOrder, counter+1);
                
                foreach( var col in table.tableInst.columns)
                {
                    if( col.ForeignKey != null)
                    {
                        if( !dictAllTDbtables.ContainsKey( col.ForeignKey.LangTableNameModel ))
                        {
                            Console.WriteLine( $"Not found table {col.ForeignKey.LangTableNameModel}" );
                            return false;
                        }
                        var tDbTable = dictAllTDbtables[ col.ForeignKey.LangTableNameModel ];
                        
                        if( !orderByDependency( tDbTable, dictAllTDbtables, stack, 
                                    Math.Max( table.counterOrder, counter+1) ))
                            return false;
                    }
                }

                stack.RemoveAt(idxStack);
                return true;
            }
        }

        public void GenerateScripts( 
                string sqlLang,
                string dir_jsons,
                string dir_sql,
                string connection_string,
                string delimeter,
                bool bApplyScripts
            )
        {
            var arg = new DialectArg( connection_string, dir_jsons, dir_sql, delimeter, bApplyScripts);
            var dialect = GenericDialect.GetDialectByName(sqlLang);
            if( dialect == null)
                return;

            dialect.fnProcessData = ProcessData;
            dialect.startConnection( arg );
        }
        void ProcessData(GenericDialect dialect, GenericDialectArg arg1 )
        {
            var arg = arg1 as DialectArg;
            var inst = new MigrationDB();

            string[] files= null;
            try
            {
                var sqlfiles = Directory.GetFiles( arg.dir_sql, $"*{EXTENSION_SQL}", SearchOption.AllDirectories);
                foreach( var sqlFile in sqlfiles)
                {
                    if( !sqlFile.Contains(PATTERN_NAME_SQL) )
                        continue;
                    dialect.setCurrentFile(sqlFile);
                    File.Delete(sqlFile);
                }

                files = Directory.GetFiles( arg.dir_jsons, $"*{EXTENSION_JSON}", SearchOption.AllDirectories);
            }catch( Exception e)
            {
                dialect.printError(e.Message);
            }
            files = files.OrderBy( x => x).ToArray();

            string prevFile = null;
            for( var iFile = 0; iFile < files.Length; iFile++)
            {
                var file = files[iFile];
                if(file.Contains( "2023.09.25.125837.models.json"))
                    Utils.Nop();

                var bMustDeleteJson = false;
                var tagFile = "";
                var content = generateScript( prevFile, file, arg.delimeter, dialect, ref tagFile, ref bMustDeleteJson);
                if( bMustDeleteJson)
                {
                    File.Delete(file);
                    continue;
                }
                if( content == "")
                    continue;


                if( arg.bApplyScripts)
                {
                    if( content.Trim() != "" )
                    {
                        inst.ApplyImport( dialect, file, content );
                    }
                }else
                {
                    var etichet  = tagFile != ""? $"-{tagFile}":"";
                    var fileName = Utils.getFileInfoName( file );                    
                    var sqlFile  = $"{arg.dir_sql}\\{fileName}{etichet}{EXTENSION_SQL}";
                        //file.Replace(EXTENSION_JSON, $"{etichet}{EXTENSION_SQL}");
                    try
                    {
                        File.WriteAllText( sqlFile, content);
                        var name = Utils.getFileInfoName(sqlFile);
                        Console.WriteLine($"{iFile.ToString("000")}. Generate file: {name}");
                    }catch( Exception e)
                    {
                        dialect.printError( e.Message );
                    }
                }

                prevFile = file;
            }
        }

        private string generateScript( string prevFile, string file, string delimeter, GenericDialect dialect, ref string tagFile, ref bool bMustDeleteJson )
        {
            var dbTables0 = ReadCustomJson(prevFile, delimeter, dialect, ref bMustDeleteJson );
            var dbTables1 = ReadCustomJson(file,     delimeter, dialect, ref bMustDeleteJson);

            if( bMustDeleteJson)
            {
                return "";
            }

            PreGenerateScript preGenerate = null;
            if( dbTables0 != null )
            {
                preGenerate = generatePartialSql(dbTables1, dbTables0, dialect, ref tagFile);
            }else
            {
                preGenerate = generateNewSqlScript(dbTables1, dialect, ref tagFile);
            }
            return preGenerate.generateScript(dialect);
        }
        private PreGenerateScript generateNewSqlScript(List<DbTable>dbTables, GenericDialect dialect, ref string tagFile)
        {
            var ret = new PreGenerateScript();
            var tableMigration = MigrationDB.createMigrationTable( dialect );

            var scriptTable = dialect.addTable(tableMigration);
            ret.listAddTables.Add( new TDbTable( MigrationDB.TABLE_MIGRATION, scriptTable));

            foreach(var table in dbTables)
            { 
                var txt = dialect.addTable( table);
                if( txt != "")
                {
                    ret.listAddTables.Add( new TDbTable( table, txt));
                }
            }
            tagFile = $"initDb";
            //var s = string.Join( dialect.SqlSeparator(), arr);
            return ret;
        }
        private PreGenerateScript generatePartialSql(
            List<DbTable>dbTables1, List<DbTable>dbTables0, 
            GenericDialect dialect, ref string tagFile )
        {
            var ret = new PreGenerateScript();
            //SQL_SEPARAT
            var s = "";
            var dict1 = Utils.getDictFromList( dbTables1, x=> x.LangTableNameModel);
            var dict0 = Utils.getDictFromList( dbTables0, x=> x.LangTableNameModel);

            var dict1Sql = Utils.getDictFromList( dbTables1, x=> x.SqlTableNameModel);
            var dict0Sql = Utils.getDictFromList( dbTables0, x=> x.SqlTableNameModel);

            
            {
                var _tagFile = "";
                var _tagFile2 = "";
                var diffTables = Utils.getDictionaryDifference(dict0,    dict1);
                var diffTblSql = Utils.getDictionaryDifference(dict0Sql, dict1Sql);

                //var arr = new List<string>();
                foreach(var table in diffTables )
                { 
                    if( diffTables.ContainsKey(table.Value.LangTableNameModel)
                     && diffTblSql.ContainsKey(table.Value.SqlTableNameModel))
                    {
                        var txt = dialect.dropTable( table.Value);
                        ret.listDropTables.Add( new TDbTable( table.Value, txt ) );
                        _tagFile += table.Value.LangTableNameModel;
                    }else
                    {
                        _tagFile2 += table.Value.LangTableNameModel;
                        //the name in golang changed
                    }
                }
                var s1 = string.Join( dialect.SqlSeparator(), ret.listDropTables.Select(x=>x.script).ToArray() );
                s += s1;
                if( _tagFile2 != "" )
                    tagFile = $"renTable{_tagFile2}";
                if( _tagFile != "" )
                    tagFile = $"dropTable{_tagFile}";
            }

            {
                var diffTables = Utils.getDictionaryDifference(dict1,    dict0);
                var diffTblSql = Utils.getDictionaryDifference(dict1Sql, dict0Sql);
                var _tagFile = "";
                var _tagFile2 = "";

                var arr = new List<string>();
                foreach(var table in diffTables )
                { 
                    if(  diffTables.ContainsKey(table.Value.LangTableNameModel)
                      && diffTblSql.ContainsKey(table.Value.SqlTableNameModel) )
                    {
                        var txt = dialect.addTable( table.Value);
                        if( txt != "")
                        {
                            ret.listDropTables.Add( new TDbTable( table.Value, txt));
                            _tagFile += table.Value.LangTableNameModel;
                        }
                    }else
                    {
                        _tagFile2 += table.Value.LangTableNameModel;
                        //the name in golang changed
                    }
                }
                var s1 = string.Join( dialect.SqlSeparator(), arr);
                s += s1;

                if( _tagFile2 != "" )
                    tagFile = $"renTable{_tagFile2}";
                if( _tagFile != "" )
                    tagFile = $"addTable{_tagFile}";
            }

            {
                var _tagFile = "";
                var _tagFile2 = "";

                var commonTables = Utils.getDictionaryUnion( dict1, dict0);
                var arr = new List<string>();
                var tagUpdate = "";
                foreach(var pair in commonTables )
                { 
                    var txt = generateDiffInsideTable( pair.Value.Item1, pair.Value.Item2, dialect, ref tagUpdate );
                    if( txt != "")
                    {
                        _tagFile += pair.Value.Item1.LangTableNameModel;
                    }

                    var sql = dialect.updateTable( pair.Value.Item1, pair.Value.Item2);
                    if(sql != "")
                    {
                        if( txt != "")
                            txt += dialect.SqlSeparator()+sql;
                        else
                            txt += sql;

                        _tagFile2 += pair.Value.Item1.LangTableNameModel;
                    }
                    if( txt != "")
                    {
                        ret.listUpdTables.Add( new TUpdColumn( pair.Value.Item1, txt));
                    }
                }
                var s1 = string.Join( dialect.SqlSeparator(), arr);
                s += s1;

                if( _tagFile2 != "" )
                    tagFile = $"renTable{_tagFile2}";

                if( _tagFile != "" )
                {
                    if( tagUpdate.Length < 50)
                        tagFile = $"upd{_tagFile}{tagUpdate}";
                    else
                        tagFile = $"updTable{_tagFile}_multiple";                    
                }
            }

            return ret;
        }

        /**
         * 
         * generate the script diferences between 2 tables
         */ 
        private string generateDiffInsideTable(DbTable table1, DbTable table2, GenericDialect dialect, ref string tagUpdate)
        {
            if( table1.LangTableNameModel != table2.LangTableNameModel)
                return dialect.printError("not the same name");

            if( table1.LangTableNameModel == "Project" )
                Utils.Nop();

            var s = "";
            var dict1 = Utils.getDictFromList(table1.columns, x=> x.langName);
            var dict0 = Utils.getDictFromList(table2.columns, x=> x.langName);

            var dict1Sql = Utils.getDictFromList(table1.columns, x=> x.sqlName);
            var dict0Sql = Utils.getDictFromList(table2.columns, x=> x.sqlName);

            var commonColumns    = Utils.getDictionaryUnion( dict1, dict0);
            var commonColumnsSql = Utils.getDictionaryUnion( dict1Sql, dict0Sql);

            var arr = new List<string>();
            {
                var diffColumns = Utils.getDictionaryDifference( dict1,     dict0);
                var diffColSql  = Utils.getDictionaryDifference( dict1Sql,  dict0Sql);

                var _tagAddUpdate = "";
                var _tagAddConstraint = "";
                foreach(var column in diffColumns )
                { 
                    if( (diffColumns.ContainsKey( column.Value.langName) )
                      && diffColSql.ContainsKey( column.Value.sqlName) )
                    {
                        var txt = dialect.addColumn( table1, column.Value );
                        arr.Add(txt);
                        _tagAddUpdate += column.Value.langName;
                    }else
                    {
                        var fkey_field = commonColumnsSql.ContainsKey(column.Value.sqlName)?
                                         commonColumnsSql[ column.Value.sqlName ]:null;
                        if( fkey_field != null )
                        { 
                            if( fkey_field.Item1.langName2 == fkey_field.Item2.langName
                             && fkey_field.Item1.langName != "" 
                             && fkey_field.Item1.ForeignKey != null
                             && fkey_field.Item2.ForeignKey == null )
                            {
                                var txt = dialect.addFKConstrictor( table1, fkey_field.Item1 );
                                if( txt != "")
                                {
                                    arr.Add( txt );
                                    _tagAddConstraint += fkey_field.Item1.langName;
                                }
                            }
                        }
                        //do nothig, just rename the structs in golang
                        //var txt = dialect.renameColumn( table1, column.Value );
                        //arr.Add(txt);
                    }
                }
                
                if( _tagAddUpdate != "" )
                    tagUpdate += $"_addCol{_tagAddUpdate}";
                if( _tagAddConstraint != "" )
                    tagUpdate += $"_addFk{_tagAddConstraint}";
            }

            {
                var _tagDropUpdate = "";
                var _tagDropConstraint = "";

                var diffColumns = Utils.getDictionaryDifference( dict0,    dict1);
                var diffColSql  = Utils.getDictionaryDifference( dict0Sql, dict1Sql);
                foreach(var column in diffColumns )
                { 
                    if( diffColumns.ContainsKey( column.Value.langName)
                     && diffColSql.ContainsKey( column.Value.sqlName) )
                    {
                        var txt = dialect.dropColumn( table1, column.Value );
                        if( txt != "")
                        {
                            arr.Add(txt);
                            _tagDropUpdate += column.Value.langName;
                        }
                    }else
                    {
                        //just the name in golang is changed
                        var fkey_field = commonColumnsSql.ContainsKey(column.Value.sqlName)?
                                         commonColumnsSql[ column.Value.sqlName ]:null;
                        if( fkey_field != null )
                        { 
                            if( fkey_field.Item2.langName2 == fkey_field.Item1.langName
                             && fkey_field.Item2.langName != "" 
                             && fkey_field.Item2.ForeignKey != null
                             && fkey_field.Item1.ForeignKey == null )
                            {
                                var txt = dialect.dropFKConstrictor( table1, fkey_field.Item2 );
                                if( txt != "")
                                {
                                    arr.Add( txt );
                                    _tagDropUpdate += fkey_field.Item2.langName;
                                }
                            }
                        }
                    }
                }
                if( _tagDropUpdate != "" )
                    tagUpdate += $"_delCol{_tagDropUpdate}";
                if( _tagDropConstraint != "" )
                    tagUpdate += $"_delFk{_tagDropConstraint}";
            }

            {   //commons
                var _tagUpdColUpdate = "";
                foreach(var pair in commonColumns )
                { 
                    var txt = dialect.updateColumn( table1, pair.Value.Item1, pair.Value.Item2 );
                    if( txt != "")
                    {
                        arr.Add(txt);
                        _tagUpdColUpdate += pair.Value.Item1.langName;
                    }
                }
                if( _tagUpdColUpdate != "" )
                    tagUpdate += $"_updCol{_tagUpdColUpdate}";
            }

            var s1 = string.Join( dialect.SqlSeparator(), arr);
            s += s1;

            return s;
        }
        /**
         * read the json hwre is the definition of structs/tables.
         */ 
        public static List<DbTable> ReadCustomJson(
            string file, string delimeter, 
            GenericDialect dialect, ref bool bMustDeleteJsonFile)
        {
            if(file == null)
                return null;

            var dbTables = new List<DbTable>();

            var content = File.ReadAllText(file);
            var parts = content.Split( new String[]{delimeter}, StringSplitOptions.RemoveEmptyEntries );
            
            foreach( var part in parts)
            {
                var text = part.Trim();
                if( text == "")
                    continue;

                var idx = text.IndexOf("={");
                if( idx >= 0)
                {
                    var tableName = text.Substring(0, idx);
                    var json      = text.Substring( idx+1);

                    var dbTable = new DbTable();
                    dbTable.initSql( tableName, null);
                    dbTable.json = json; //store here, and use it after
                    dbTables.Add( dbTable );
                }else
                {
                    Console.WriteLine("json definition wrong");
                }
            }

            var dictSqlTables = new Dictionary<string, DbTable>();

            foreach( var dbTable in dbTables)
            {
                if( dbTable.LangTableNameModel == "User")
                    Utils.Nop();

                var nextShouldBeFK = "";
                var nextShouldBeFK_pointerType= "";

                var obj = Newtonsoft.Json.JsonConvert.DeserializeObject<Dictionary<string, Dictionary<string, string>>>(dbTable.json);

                var sqlTableName = dbTable.LangTableNameModel;
                if( obj.ContainsKey(TABLE_SQLNAME))
                    sqlTableName = decodeSqlTableName( obj[TABLE_SQLNAME][TAG_tags] );
                if( sqlTableName != "" )
                {
                    dbTable.SqlTableNameModel = sqlTableName;
                    if( !dictSqlTables.ContainsKey(dbTable.SqlTableNameModel) )
                    {
                        dictSqlTables[ dbTable.SqlTableNameModel ] = dbTable;
                    }else
                    {   bMustDeleteJsonFile = true;
                        dialect.printError( $"duplicate table at sql Name: '{dbTable.SqlTableNameModel}'. check if is duplicate or you miss to do in 2 steps the renaming of a table");
                    }
                }

                var dictSqlColumns = new Dictionary<string, DbColumn>();

                foreach( var it in obj)
                {
                    if( it.Key == TABLE_SQLNAME)
                        continue;
                    var column = new DbColumn();
                    var type = it.Value[TAG_Type];
                    var tags = it.Value[TAG_tags];
                    var langName2 = "";
                    var langName  = it.Key;
                    DbTable tblPointer = null;

                    var sqlName = getSqlName(tags, dialect );
                    if( sqlName == "")
                    {
                        continue;
                    }
                    if( !dialect.isNoSql() && it.Key == GoModelTemplate.NoSqlID)
                        continue;
                    
                    if( sqlName == GoModelTemplate.FIELD_IS_OMITTED_INTEGRAL)
                    {
                        nextShouldBeFK = it.Key;
                        nextShouldBeFK_pointerType = type;
                        continue;
                    }else
                    if( nextShouldBeFK != "" && !dialect.isNoSql() )
                    {
                        var compressedName = GoModelTemplate.ConvertToIdent_GoLang( it.Key);
                        if( nextShouldBeFK != compressedName)
                        {
                            bMustDeleteJsonFile = true;
                            dialect.printError($"names dont match {it.Key} - {nextShouldBeFK}");
                            return null;
                        }
                        nextShouldBeFK = "";
                        var tableName = GenericDialect.CleanNameGoStruct( nextShouldBeFK_pointerType );
                        tblPointer = dbTables.Where(x=>x.LangTableNameModel == tableName).FirstOrDefault();
                        if(tblPointer == null )
                        {
                            bMustDeleteJsonFile = true;
                            dialect.printError($"not found table {tableName}");
                            return null;
                        }
                        var __ID = GoModelTemplate.POSTFIX_FOREIGNKEY;
                        if( !sqlName.EndsWith(__ID)
                         || !langName.EndsWith(__ID) )
                        {
                            bMustDeleteJsonFile = true;
                            dialect.printError($"in table '{dbTable.LangTableNameModel}' name of foreign key field ({langName} -> {sqlName}) should end in '{__ID}'.");
                            return null;
                        }
                        langName2 = langName;
                        langName  = compressedName;
                        //column.ForeignKey = tblPointer;
                    }

                    //var langName = it.Key;
                    var langType = type;
                    var bIsIdentity = tags.Contains(GoModelTemplate.MARK_PRIMARY_KEY);
                    if( bIsIdentity)
                        dbTable.PrimaryColumn = column;

                    //column.sqlName     = getSqlName(tags);//`json:"ID'
                    //column.sqlType     = dialect.getSqlType( column.langType, ref column.bIsNullable);
                    var bIsNullable = false;
                    var sqlType = dialect.getSqlType( langType, ref bIsNullable, sqlName);

                    column.initLangSql(
                        langName: langName,
                        langName2: langName2,
                        langType: type,
                        sqlName: sqlName,//`json:"ID'
                        sqlType: sqlType,
                        bIsNullable: bIsNullable,
                        bIsIdentity: bIsIdentity,
                        ForeignKey: tblPointer,
                        tags
                        );

                    if( !dictSqlColumns.ContainsKey(column.sqlName) )
                    {
                        dictSqlColumns[ column.sqlName ] = column;
                    }else
                    {   
                        bMustDeleteJsonFile = true;
                        dialect.printError( $"duplicate column aat sql Name: '{column.sqlName}'. check if is duplicate or you miss to do in 2 steps the renaming of a column");
                    }
                    dbTable.columns.Add( column);
                }
                //dbTables.Add(dbTable);
            }

            return dbTables;
        }
        /**
         * 
         * `json:"ID,omitempty"`
         * `json:"UUID"`
         *  `json:"-"`
         */ 
        private static string getSqlName(string tags, GenericDialect dialect)
        {
            if( tags == "")
                return "";

            if( tags == GoModelTemplate.FIELD_IS_OMITTED_INTEGRAL)
                return "";

            var bFoundDescr = false;
            var arrTags = tags.Split(new string[]{ "  ", }, StringSplitOptions.RemoveEmptyEntries );
            foreach( var tag1 in arrTags)
            {
                var tag = tag1.Trim();
                if( tag.StartsWith("\"bson:\\\"") )
                {
                    bFoundDescr = true;
                    if( dialect.isNoSql() )
                    {
                        var tags1 = tag.Replace("\"bson:\\\"", "");
                        tags1 = tags1.Replace("\\\"\"", "");
                        tags1 = tags1.Replace("\\\"", "");
                        
                        var p = tags1.Split(',');
                        return p[0].Trim();
                    }
                }else
                if( tag.StartsWith("\"json:\\\"") )
                {
                    bFoundDescr = true;

                    var tags1 = tag.Replace("\"json:\\\"", "");
                    tags1 = tags1.Replace("\\\"\"", "");
                    tags1 = tags1.Replace("\\\"", "");
                    
                    var p = tags1.Split(',');
                    return p[0].Trim();
                } else
                if( tag.StartsWith("\"atmsql:\\\"") )
                {
                    bFoundDescr = true;

                    var tags1 = tag.Replace("\"atmsql:\\\"", "");
                    tags1 = tags1.Replace("\\\"\"", "");
                    tags1 = tags1.Replace("\\\"", "");
                    
                    var p = tags1.Split(',');
                    return p[0].Trim();
                }
            }
            
            if( !bFoundDescr )
            {
                Console.WriteLine($"json definition of description field is incomplete :{tags}");
            }
            return "";
        }
        /**
         * 
         * atomicsql-table:"user_role"
         */ 
        private static string decodeSqlTableName(string tags)
        {
            if( tags == "")
                return "";

            if( tags.StartsWith("/*atomicsql-table:\"") )
            {
                var tags1 = tags.Replace("/*atomicsql-table:\"", "");
                //tags1 = tags1.Substring( 0, tags1.Length-1);//remove ending "
                tags1 = tags1.Replace("\"*/", "");
                return tags1;
            } else
            {
                Console.WriteLine($"json definition of atomicsql-table incorrect :{tags}");
            }
            return "";
        }
        
        
        bool isYes(string s)
        {
            return s != null && s == "YES";
        }
    }
}
