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
    public partial class MongodbDialect : GenericDialect
    {
        public override string getSql()
        { 
            return "mongodb";
        }
        public override bool isNoSql()
        {
            return true;
        }
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
            return $"{tokenizIdentif(table.SqlTableNameModel)}";
        }
        public string tokenizTableName(string tableName)
        {
            return $"{tokenizIdentif(tableName)}";
        }
        public override string dropTable(DbTable table)
        {
            var s = $@"";
            return s;
        }

        public override string addTable(DbTable table)
        {
            return "";
        }

        public override string dropColumn(DbTable table,DbColumn column)
        {
            var s = $@"";
            return s;
        }
        public override string addColumn(DbTable table, DbColumn column)
        { 
            var s = $@"";
            return s;
        }
        public string _addColumn(DbColumn column)
        {
            return "";
        }
        public override string updateTable(DbTable table, DbTable tablePrev)
        {
            return "";
        }


        //----------------------------------------------------




        public override string updateColumn(DbTable table, DbColumn column, DbColumn columnPrev)
        {
            
            return "";
        }
        
        public override string addFKConstrictor(DbTable table,DbColumn column)
        {
            return "";
        }
        public override string dropFKConstrictor(DbTable table,DbColumn column)
        {
            return "";
        }

        public override string getDefaultSchema()
        {
            return null;
        }
        
        public override string getSqlType( string langType, ref bool bIsNullable, string nameOfColumn)
        {
            var originalLangType = langType;
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
                case "ObjectID" : return "ObjectID";
            }
            if( nameOfColumn.EndsWith("_ID") 
             || originalLangType.StartsWith("*"))
            {   //for foreignkey.
                return "TEXT";
            }
            return printError($"not supported type {langType}");
        }
        public override string getProperty(  string propName, 
            string tableName, string colName, string colValue)
        {
            
            return null;
        }
        public override void updateProperty( 
            string propName, string value, 
            string tableName, string colName, string colValue)
        {
            
        }
        public override void insertProperty(
            string propName, string value, 
            string tableName, string colName, string colValue)
        {
            
        }
        public override void execScript( string scriptTxt)
        {
            
        }

        public override bool startConnection(GenericDialectArg arg)
        {
            try
            {
                if( true)
                {
                    this.fnProcessData( this, arg);
                    return true;
                }
            }catch(Exception e )
            {
                Console.WriteLine(e.Message);
            }finally
            {
            }
            return false;
        }
    }
}
