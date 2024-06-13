using System;
using System.Collections.Generic;
using System.Linq;
using Npgsql;
using MySql.Data.MySqlClient;

namespace src_tool
{
    partial class MySqlDialect
    {
        public override Dictionary<string, DbTable> readTables( string SqlLang )
        {
            var tables = new Dictionary<string, DbTable>();
            var cnn = this.connection;
            var packageImports = new Dictionary<string,string>();

            var myDatabaseName = this.connection.Database;
            var requestText = $@"
                SELECT 
                    table_name,          /*0*/
                    table_schema,        /*1*/
                    column_name,         /*2*/ 
                    data_type,           /*3*/
                    table_catalog,       /*4*/
                    is_nullable,         /*5*/
                    extra,               /*6*/
                    column_default,      /*7*/
                    ordinal_position
                FROM information_schema.columns
                WHERE table_schema = DATABASE()
                order by table_schema, table_name, ordinal_position
			;
            ";

            //var cmd = cnn.CreateCommand();
            using( var cmd = new MySqlCommand( requestText, cnn))
            {
                //cmd.CommandText = requestText;
                //cmd.ExecuteReader();requestText);
                var reader = cmd.ExecuteReader();

                if (reader.HasRows)
                {
                    DbTable table = null;
                    while (reader.Read())
                    {
                        var tableName   = getString(reader, 0);
                        var tableSchema = getString(reader, 1);

                        table = tables.ContainsKey(tableName)
                              ? tables[tableName]:new DbTable();

                        table.LangTableNameModel = GoModelTemplate.ConvertToIdent_GoLang( tableName );
                        table.Schema = tableSchema;
                        table.SqlTableNameModel = tableName;
                        if( tableName == "user_role")
                            Utils.Nop();
                        tables[ table.SqlTableNameModel ] = table;
                        
                        var columnName      = getString(reader, 2);
                        if( columnName.Contains( "Opened" ))
                            Utils.Nop();
                        if( columnName == "\"OpenedItems\"" )
                            Utils.Nop();
                        var colType         = getString(reader, 3);
                        var colType2        = getString(reader, 4);
                        var colIsNullable   = getString(reader, 5);
                        var colIsIdentity   = getString(reader, 6);
                        var colDefalut      = getString(reader, 7);

                        var column  = new DbColumn();
                        column.langName  = GoModelTemplate.ConvertToIdent_GoLang( columnName, true );
                        column.langName2 = GoModelTemplate.ConvertToIdent_GoLang( columnName, false );
                                            
                        column.sqlName = columnName;
                        column.sqlType = colType;
                        column.langType =  this.getGoLangType( column, ref packageImports );
                        column.bIsIdentity = colIsIdentity != null && colIsIdentity.Contains("auto_increment");
                        column.bIsNullable = isYes(colIsNullable);
                        table.columns.Add( column );

                        if( column.bIsIdentity )
                        {
                            if( table.PrimaryColumn != null ) 
                            {
                                Console.WriteLine($"{table} has already a primary key");
                                return null;
                            }
                            table.PrimaryColumn = column;
                        }

                        //var column          = reader.GetString(6);
                        //Console.WriteLine("{0}\t{1}", reader.GetInt32(0),reader.GetString(1));
                    }
                    if( table == null )
                    {
                        printError($"table no loaded");
                        return null;
                    }
                    if( table.PrimaryColumn == null )
                    {
                        printError($"table {table.LangTableNameModel} doesnt have a primary key");
                        return null;
                    }
                    reader.Close();
                
                    return tables;
                }else
                {
                    Console.WriteLine("No rows found. Maybe, The user has no rights to read DB defs");
                }
            }
            return null;
        }

        public override bool readConstraintors( Dictionary<string, DbTable> tables  )
        {
            var cnn = this.connection;
            var requestText = $@"
            SELECT DISTINCT(TC.table_name) AS table_name,   /*0*/
                TC.table_schema AS table_schema,            /*1*/
                KC.column_name as column_name,              /*2*/  
                KC.REFERENCED_TABLE_NAME,                   /*3*/
                KC.REFERENCED_COLUMN_NAME,                  /*4*/
                KC.CONSTRAINT_NAME,
                TC. CONSTRAINT_TYPE AS ConstraintType
                FROM 
                information_schema.table_constraints TC,  
                information_schema.key_column_usage KC  
                WHERE
                TC.CONSTRAINT_TYPE IN ('FOREIGN KEY', 'PRIMARY KEY') 
                and KC.table_name = TC.table_name and KC.TABLE_SCHEMA = TC.TABLE_SCHEMA
                and KC.CONSTRAINT_NAME = TC.CONSTRAINT_NAME
                and KC.table_schema = DATABASE()
                order by table_name asc, column_name;
            ";

            //var cmd = cnn.CreateCommand();
            using( var cmd = new MySqlCommand( requestText, cnn))
            {
                //cmd.CommandText = requestText;
                //cmd.ExecuteReader();requestText);
                var reader = cmd.ExecuteReader();

                if (reader.HasRows)
                {
                    while (reader.Read())
                    {
                        var tableName          = getString(reader, 0);
                        var schemaName         = getString(reader, 1);
                        

                        var columnName      = getString(reader, 2);
                        var colRefTable     = getString(reader, 3);
                        var colRefColumn    = getString(reader, 4);
                        //var colIsNullable   = reader.GetString(4);
                        //var colIsIdentity   = reader.GetString(5);

                        if( colRefTable != null && colRefColumn != null )
                        if( tables.ContainsKey(tableName) )
                        {
                            var table = tables[tableName];
                            var column = table.columns.Where(x=>x.sqlName == columnName).FirstOrDefault();
                            if(column != null )
                            {
                                if( tables.ContainsKey(colRefTable) )
                                {
                                    var primary = tables[colRefTable].PrimaryColumn;
                                    if( primary != null && primary.langName == colRefColumn)
                                    {
                                        column.ForeignKey = tables[colRefTable];
                                    }else
                                    {
                                        printError($"Foreign key from table {tableName} to {colRefTable} is not pointed to primary key. Add keyword {GoModelTemplate.MARK_PRIMARY_KEY} in column tags");
                                        return false;
                                    }
                                }else
                                {
                                    printError($"not found table {colRefTable}");
                                    return false;
                                }
                            }
                        }else
                        {
                            printError($"not found table {tableName}");
                            return false;
                        }

                        //var column          = reader.GetString(6);
                        //Console.WriteLine("{0}\t{1}", reader.GetInt32(0),reader.GetString(1));
                    }
                    reader.Close();
                    return true;
                }else
                {
                    printError("No rows found. Maybe, The user has no rights to read DB defs");
                }
            }
            return false;
        }

        private string getString(MySqlDataReader reader, int index)
        {
            if( !reader.IsDBNull(index) )
                return reader.GetString( index );
            return null;
        }
        

        bool isYes(string s)
        {
            return s != null && s == "YES";
        }
    }
}
