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
using MySql.Data.MySqlClient;


/**
pt mssql: se dezactiveaza indecsii cand stergem chestii mari
  ALTER INDEX ALL ON schema_name.table_name DISABLE;

ALTER INDEX ALL ON schema_name.table_name REBUILD WITH (ONLINE = ON);

pt mysql

ALTER TABLE table_name DISABLE KEYS;
ALTER TABLE table_name ENABLE KEYS;
SET FOREIGN_KEY_CHECKS=0;
SET FOREIGN_KEY_CHECKS=1;

pentru a investiga paramentrii procedurilor stocate
SELECT * from INFORMATION_SCHEMA.PARAMETERS




*/
namespace src_tool
{
    public partial class MySqlDialect : GenericDialect
    {
        public MySqlConnection connection;

        public override string getSql()
        { 
            return "mysql";
        }
    
        public override string SqlSeparator()
        { 
            return ";";
        }
        public string tokenizIdentif(string identif)
        {
            return $"`{identif}`";
        }
        public string tokenizTable(DbTable table)
        {
            return tokenizIdentif(table.SqlTableNameModel);
        }
        public override string dropTable(DbTable table)
        {
            var s = $@"DROP TABLE IF EXISTS {tokenizIdentif(table.SqlTableNameModel)}";
            return s;
        }
        public override string addTable(DbTable table)
        {
            var NL = 
                @"
                ";
            var constraintorList = new List<string>();
            var sqlTableName = table.SqlTableNameModel;

            var columnsArr = new List<string>();
            foreach( var col in table.columns)
            {
                columnsArr.Add( this._addColumn(col, ref constraintorList) );
            }
            var columnsDefs = string.Join( ","+NL, columnsArr );
            var colID_Name = table.PrimaryColumn.sqlName;

            var constraintors = string.Join( ","+NL, constraintorList );
            if( constraintors != "")
                constraintors = $",{NL}{constraintors}";

            var tableAdd = $@"        
            #-------------------------------------------------------------------	
            CREATE TABLE IF NOT EXISTS {tokenizIdentif(sqlTableName)}
            (
                { columnsDefs}

                ,PRIMARY KEY ({tokenizIdentif(colID_Name)})
                {constraintors}
            ) ENGINE=InnoDB";

                return tableAdd;
        }
        public override string dropColumn(DbTable table,DbColumn column)
        {
            var s = $@"
            ALTER TABLE {tokenizIdentif(table.SqlTableNameModel)}
            DROP COLUMN {tokenizIdentif(column.sqlName)}";
            return s;
        }
        public override string addColumn(DbTable table, DbColumn column)
        { 
            var listConstraintors = new  List<string>();
            var colData = _addColumn( column, ref listConstraintors);
            var s = $@"
            ALTER TABLE {tokenizIdentif(table.SqlTableNameModel)}
            ADD COLUMN {colData}";
            return s;
        }
        public string _addColumn(DbColumn column, ref List<string> listConstraintors)
        {
            var colName = column.sqlName;
            if( column.bIsIdentity )
            {
                if( column.langType == "int64")
                {
                    return $"{tokenizIdentif(colName)} BIGINT NOT NULL AUTO_INCREMENT";
                }else
                {
                    return $"{tokenizIdentif(colName)} INT NOT NULL AUTO_INCREMENT";
                }
            }else
            {
                var sqlType = column.sqlType;
                if( column.ForeignKey != null)
                {
                    var targetTableSqlName = tokenizTable(column.ForeignKey);
                    var targetTable_ID     = tokenizIdentif(column.ForeignKey.PrimaryColumn.sqlName);

                    var fkName      = tokenizIdentif(column.sqlName);
                    var fk_fkName   = tokenizIdentif($"fk_{column.sqlName}");

                    listConstraintors.Add($"CONSTRAINT {fk_fkName} FOREIGN KEY ({fkName}) REFERENCES {targetTableSqlName}({targetTable_ID})");
                    //listConstraintors.Add($"CONSTRAINT {fk_fkName} FOREIGN KEY ({fkName}) REFERENCES {targetTableSqlName}({targetTable_ID}) ON DELETE RESTRICT");
                    
                    return $"{fkName} INT NULL";
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
                RENAME TABLE {tokenizIdentif(tablePrev.SqlTableNameModel)}
                TO {tokenizIdentif(table.SqlTableNameModel)}";
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
                ALTER TABLE {tokenizIdentif(table.SqlTableNameModel)} 
                MODIFY {tokenizIdentif(column.sqlName)} {column.sqlType}";
                return r;
            }

            if( column.sqlName != columnPrev.sqlName)
            {
                var s = $@"
                ALTER TABLE {tokenizIdentif(table.SqlTableNameModel)}
                CHANGE COLUMN {tokenizIdentif(columnPrev.sqlName)} {tokenizIdentif(column.sqlName)} {column.sqlType}";
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
                DROP FOREIGN KEY {fk_fkName};";
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
                case "NullString":      bIsNullable= true; return isLongType(nameOfColumn)?"TEXT":"VARCHAR(255)";
                case "NullBool":        bIsNullable= true; return "BOOLEAN";

                case "NullByte":        bIsNullable= true; return "SMALLINT";
                case "NullInt16":       bIsNullable= true; return "SMALLINT";
                case "NullInt":         bIsNullable= true; return "INT";
                case "NullInt32":       bIsNullable= true; return "INT";
                case "NullInt64":       bIsNullable= true; return "BIGINT";
                case "NullFloat64":         bIsNullable= true; return "double precision";
                case "time.NullTime":   bIsNullable= true; return "TIMESTAMP";
                case "NullTime":        bIsNullable= true; return "TIMESTAMP";

                case "string":      return isLongType(nameOfColumn)?"TEXT":"VARCHAR(255)";
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

                case "float32": return "FLOAT";
                case "float64": return "DOUBLE PRECISION";
            }
            return printError($"not supported type {langType}");
        }
        public override string getProperty(  string propName, 
            string tableName, string colName, string colValue)
        {
            try
            {
                var requestText = $@"
                SELECT {tokenizIdentif(colValue)} FROM {tokenizIdentif(tableName)}
                 WHERE {tokenizIdentif(colName)} = '{propName}'
            ";

            using( var cmd = new MySqlCommand( requestText, this.connection))
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

            using( var cmd = new MySqlCommand( requestText, this.connection))
            {
                var reader = cmd.ExecuteNonQuery();
            }
        }
        public override void insertProperty(
            string propName, string value, 
            string tableName, string colName, string colValue)
        {
            var requestText = $@"
            INSERT INTO {tokenizIdentif(tableName)} ( {tokenizIdentif("ID")}, {tokenizIdentif(colName)}, {tokenizIdentif(colValue)}) 
            VALUES (  default, '{propName}', '{value}' ) ";

            using( var cmd = new MySqlCommand ( requestText, this.connection))
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
                    using( var cmd = new MySqlCommand( script1, this.connection))
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
            MySqlConnection cnn = null;
            try
            {
                var connectionString = arg.connection_string;
                cnn = new MySqlConnection(connectionString);
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
