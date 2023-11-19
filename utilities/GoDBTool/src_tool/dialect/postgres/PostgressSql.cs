using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Data;
using System.Data.SqlClient;
using System.Drawing;
using Npgsql;
using Npgsql.Logging;
using System.IO;
using src_tool.templates;

namespace src_tool
{
    public partial class PostgressDialect : GenericDialect
    {
        public NpgsqlConnection connection;

        public override string SqlSeparator()
        { 
            return ";";
        }
        
        public string tokenizIdentif(string identif)
        {
            return $"\"{identif}\"";
        }
        public string tokenizTable(DbTable table)
        {
            return $"public.{tokenizIdentif(table.SqlTableNameModel)}";
        }
        public string tokenizTableName(string tableName)
        {
            return $"public.{tokenizIdentif(tableName)}";
        }
        public override string dropTable(DbTable table)
        {
            var s = $@"DROP TABLE {tokenizTable(table)}";
            return s;
        }

        public override string addTable(DbTable table)
        {
            var NL = 
                @"
                ";
            var sqlTableName = table.SqlTableNameModel;
            var columnsArr = new List<string>();
            foreach( var col in table.columns)
            {
                columnsArr.Add( this._addColumn(col) );
            }
            var columnsDefs = string.Join( ","+NL, columnsArr );
            var colID_Name = table.PrimaryColumn.sqlName;

            var tableAdd = $@"        
            -------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS {tokenizTable(table)}
            (
                { columnsDefs}
            )
            TABLESPACE pg_default;
                ";

                return tableAdd;

            //,CONSTRAINT {tokenizIdentif(sqlTableName+"_pkey")} PRIMARY KEY ({tokenizIdentif(colID_Name)})
        }



        public override string dropColumn(DbTable table,DbColumn column)
        {
            var s = $@"
            ALTER TABLE {tokenizTable(table)}
            DROP COLUMN {tokenizIdentif(column.sqlName)}";
            return s;
        }
        public override string addColumn(DbTable table, DbColumn column)
        { 
            var colData = _addColumn( column);
            var s = $@"
            ALTER TABLE {tokenizTable(table)}
            ADD COLUMN {colData}";
            return s;
        }
        public string _addColumn(DbColumn column)
        {
            var colName = column.sqlName;
            if( column.bIsIdentity )
            {
                if( column.langType == "int64")
                {
                    return $"{tokenizIdentif(colName)} bigserial NOT NULL AUTO_INCREMENT";
                }else
                    return $"{tokenizIdentif(colName)} serial PRIMARY KEY NOT NULL";
            }else
            {
                var sqlType = column.sqlType;
                if( column.ForeignKey != null)
                {
                    //var targetTableSqlName = column.ForeignKey.SqlTableNameModel;
                    var targetTable_ID  = column.ForeignKey.PrimaryColumn.sqlName;
                    var fk_fkName       = tokenizIdentif($"fk_{column.sqlName}");
                    var fkName          = tokenizIdentif(colName);
                    
                    //return $"{tokenizIdentif(colName)} integer NULL REFERENCES {tokenizTable(column.ForeignKey)} ({tokenizIdentif(targetTable_ID)})";
                    return $"{fkName} integer NULL CONSTRAINT {fk_fkName} REFERENCES {tokenizTable(column.ForeignKey)} ({tokenizIdentif(targetTable_ID)})";
                }else
                {
                    if( column.bIsNullable )
                    {
                        return $"{tokenizIdentif(colName)} {sqlType} NULL";
                    }else
                    {
                        return $"{tokenizIdentif(colName)} {sqlType} NOT NULL";
                        //return $"{tokenizIdentif(colName)} {sqlType} NOT NULL DEFAULT ''::character varying";                        
                    }
                }
            }
        }
        public override string updateTable(DbTable table, DbTable tablePrev)
        {
            if( table.SqlTableNameModel != tablePrev.SqlTableNameModel)
            {
                var s = $@"
                ALTER TABLE public.{tokenizIdentif(tablePrev.SqlTableNameModel)}
                RENAME TO {tokenizIdentif(table.SqlTableNameModel)}";
                return s;
            }
            return "";
        }


        //----------------------------------------------------




        public override string updateColumn(DbTable table, DbColumn column, DbColumn columnPrev)
        {
            if( columnPrev.langName != column.langName)
                return printError("internal error");

            if( column.sqlType != columnPrev.sqlType )
            {
                var r = $@"
                ALTER TABLE {tokenizTable(table)} 
                ALTER COLUMN {tokenizIdentif(column.sqlName)} TYPE {column.sqlType}";
                return r;
            }

            if( column.sqlName != columnPrev.sqlName)
            {
                var s = $@"
                ALTER TABLE {tokenizTable(table)}
                RENAME COLUMN {tokenizIdentif(columnPrev.sqlName)} TO {tokenizIdentif(column.sqlName)}";
                return s;
            }

            return "";
        }
        
        public override string addFKConstrictor(DbTable table,DbColumn column)
        {
            var tablename   = tokenizTable(column.ForeignKey);
            var tableID     = tokenizIdentif(column.ForeignKey.PrimaryColumn.sqlName);

            var myTableName = tokenizTable(table);
            var fkName      = tokenizIdentif(column.sqlName);
            var fk_fkName   = tokenizIdentif($"fk_{column.sqlName}");
            
            var s = $@"
                ALTER TABLE {myTableName}
                ADD CONSTRAINT {fk_fkName} FOREIGN KEY ({fkName}) REFERENCES {tablename} ({tableID});";
            return s;
        }
        public override string dropFKConstrictor(DbTable table,DbColumn column)
        {
            //var tablename   = tokenizIdentif(column.ForeignKey.SqlTableNameModel);
            //var tableID     = tokenizIdentif(column.ForeignKey.PrimaryColumn.sqlName);

            var myTableName = tokenizTable(table);
            //var fkName      = tokenizIdentif(column.sqlName);
            var fk_fkName   = tokenizIdentif($"fk_{column.sqlName}");

            var s = $@"
                ALTER TABLE {myTableName}
                DROP CONSTRAINT {fk_fkName};";
            return s;
        }

        
        public override string getSqlType( string langType, ref bool bIsNullable, string nameOfColumn)
        {
            langType = cleanNameGoStruct(langType);
            if(langType.StartsWith("[]"))
            {
                return printError($"type {langType} not supported!");
            }

            switch( langType )
            { 
                case "NullString":      bIsNullable= true; return isLongType(nameOfColumn)?"TEXT":"VARCHAR";
                case "NullBool":        bIsNullable= true; return "BOOLEAN";

                case "NullByte":        bIsNullable= true; return "SMALLINT";
                case "NullInt16":       bIsNullable= true; return "SMALLINT";
                case "NullInt":         bIsNullable= true; return "INT";
                case "NullInt32":       bIsNullable= true; return "INT";
                case "NullInt64":       bIsNullable= true; return "BIGINT";
                case "NullFloat64":         bIsNullable= true; return "double precision";
                case "time.NullTime":   bIsNullable= true; return "TIMESTAMP";
                case "NullTime":        bIsNullable= true; return "TIMESTAMP";

                case "string":      return isLongType(nameOfColumn)?"TEXT":"VARCHAR";
                case "char":        return "CHAR";
                case "bool":        return "BOOLEAN";

                case "byte":        return "SMALLINT";
                case "int16":       return "SMALLINT";
                case "int":         return "INT";
                case "int32":       return "INT";
                case "int64":       return "BIGINT";

                //case "time.Time":   return "TIMESTAMP";
                case "Time":        return "TIMESTAMP";

                case "UUID": return "UUID";
                case "uuid.UUID": return "UUID";

                case "float32": return "real";
                case "float64": return "double precision";
            }
            return printError($"not supported type {langType}");
        }
        public override string getProperty(  string propName, 
            string tableName, string colName, string colValue)
        {
            try
            {
                var requestText = $@"
                SELECT {tokenizIdentif(colValue)} FROM {tokenizTableName(tableName)}
                 WHERE {tokenizIdentif(colName)} = '{propName}'
            ";

            using( var cmd = new NpgsqlCommand( requestText, this.connection))
            {
                var obj = cmd.ExecuteScalar();
                if( obj != null)
                    return obj.ToString();
            }
            }catch(Exception e )
            {
                var s = e.Message;
            }
            return null;
        }
        public override void updateProperty( 
            string propName, string value, 
            string tableName, string colName, string colValue)
        {
            var requestText = $@"

            UPDATE {tokenizIdentif(tableName)}
            SET   {tokenizIdentif(colValue)}  = '{value}'
            WHERE {tokenizIdentif(colName)}   = '{propName}'
            ";

            using( var cmd = new NpgsqlCommand( requestText, this.connection))
            {
                var reader = cmd.ExecuteNonQuery();
            }
        }
        public override void insertProperty(
            string propName, string value, 
            string tableName, string colName, string colValue)
        {
            var requestText = $@"
            INSERT INTO {tokenizTableName(tableName)} ( {tokenizIdentif("ID")}, {tokenizIdentif(colName)}, {tokenizIdentif(colValue)}) VALUES (  default, '{propName}', '{value}' ) ";

            using( var cmd = new NpgsqlCommand( requestText, this.connection))
            {
                var reader = cmd.ExecuteNonQuery();
            }
        }
        public override void execScript( string scriptTxt)
        {
            var parts = scriptTxt.Split(';');
            var iPart = 0;
            foreach( var part in parts)
            {
                var script1 = part.Trim();
                if( script1.Length >  0 )
                {
                    using( var cmd = new NpgsqlCommand( script1, this.connection))
                    {
                        try
                        {
                            var reader = cmd.ExecuteNonQuery();
                        }catch(Exception e)
                        {
                            var msg = e.InnerException != null ? e.InnerException.Message : e.Message;
                            Console.WriteLine($"Error exec script part {iPart}: {msg}");
                            throw e;
                        }
                    }
                }
                iPart ++;
            }
        }

        public override bool startConnection(GenericDialectArg arg)
        {
            NpgsqlConnection cnn = null;
            try
            {
                var connectionString = arg.connection_string;
                cnn = new NpgsqlConnection(connectionString);
                cnn.Open();
                this.connection = cnn;

                if( true)
                {
                    this.fnProcessData( this, arg);
                    return true;
                }
                else
                {
                    using(var transaction = cnn.BeginTransaction())
                    {
                        try
                        {
                            this.fnProcessData( this, arg);
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
            }finally
            {
                if( cnn != null)
                    cnn.Close();
            }
            return false;
        }
    }
}
