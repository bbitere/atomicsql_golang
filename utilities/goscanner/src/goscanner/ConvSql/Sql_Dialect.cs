 using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime;
using goscanner.AST;
using goscanner.ConvSql;

namespace goscanner.ConvSql
{
    public abstract class Sql_Dialect
    {
        public string TRUE = "true";
        public string FALSE = "false";
        public string NULL = "null";
        public string THIS = "@@THIS@@";

        public Sql_Dialect() { }

        public virtual string convertGolangTypeToSqlType( 
            string typeName, ParserRuleContext ctx, SqlConvert convertInst)
        { 
            return typeName;
        }
        public virtual string convertFormatTime_GolangToSql( string format)
        {
            return format;
        }
        public virtual string QuoteField(string fieldName){ return fieldName;}

        public abstract string GetDialectName();//{ return "";}

        public abstract bool isNoSql();//{ return "";}
        
    }

    public class PostgresSql_Dialect :Sql_Dialect
    {
        public PostgresSql_Dialect()
        {
            TRUE = "true";
            FALSE = "false";
            NULL = "null";
        }

        public override bool isNoSql(){ return false;}
        public override string GetDialectName(){ return "postgres";}
        public override string convertGolangTypeToSqlType( 
            string typeName, ParserRuleContext ctx, SqlConvert convertInst)
        {
            typeName = typeName.Replace("@", ""); 

            switch(typeName)
            { 
                case "string": return "VARCHAR(MAX)";
                case "char": return "CHAR";
                case "bool": return "BOOLEAN";

                case "byte": return "SMALLINT";
                case "int16": return "SMALLINT";
                case "int": return "INT";
                case "int32": return "INT";
                case "int64": return "BIGINT";

                case "time.Time": return "TIMESTAMP";
                case "Time": return "TIMESTAMP";

                case "UUID": return "UUID";
                case "uuid.UUID": return "UUID";

                case "float32": return "real";
                case "float64": return "double precision";
            }
            return convertInst.getTextSQLError("type cannot be converted", ctx);
        }

        public override string convertFormatTime_GolangToSql( string format1)
        {
            var format = format1;
            format = format.Replace("2006","YYYY");
            format = format.Replace("15","HH24");
            format = format.Replace("01","MM");
            format = format.Replace("02","DD");
            format = format.Replace("04","MI");
            format = format.Replace("05","SS");

            format = format.Replace(".000000","FF6");
            format = format.Replace(".000","FF3");

            format = format.Replace("Jan","Mon");
            format = format.Replace("January","Month");
            format = format.Replace("Mon","Day");
            
            return format;
        }
        public override string QuoteField(string fieldName)
        { 
            return $"\"{fieldName}\"";
        }
    }
    public class MySql_Dialect :Sql_Dialect
    {
        public MySql_Dialect()
        {
            TRUE = "true";
            FALSE = "false";
            NULL = "null";
        }

        public override bool isNoSql(){ return false;}
        public override string GetDialectName(){ return "mysql";}
        public override string convertGolangTypeToSqlType( 
            string typeName, ParserRuleContext ctx, SqlConvert convertInst)
        {
            typeName = typeName.Replace("@", ""); 
            switch(typeName)
            { 
                case "string": return "NVARCHAR(MAX)";
                case "char": return "CHAR";
                case "bool": return "BIT";

                case "byte": return "TINYINT";
                case "int16": return "SMALLINT";
                case "int": return "INT";
                case "int32": return "INT";
                case "int64": return "BIGINT";

                case "time.Time": return "TIMESTAMP";
                case "Time": return "datetime2";

                case "UUID": return "VARCHAR(MAX)";
                case "uuid.UUID": return "VARCHAR(MAX)";

                case "float32": return "FLOAT";
                case "float64": return "REAL";

                //case "float64": return "DECIMAL";
            }
            return convertInst.getTextSQLError("type cannot be converted", ctx);
        }

        public override string convertFormatTime_GolangToSql( string format1)
        {
            var format = format1;
            format = format.Replace("2006","%Y");
            
            format = format.Replace("01","%m");
            format = format.Replace("02","%d");

            format = format.Replace("15","%H");
            format = format.Replace("04","%i");
            format = format.Replace("05","%s");

            format = format.Replace(".000000",".%f");
            format = format.Replace(".000",".%f");

            format = format.Replace("Jan","%b");
            format = format.Replace("January","%M");
            format = format.Replace("Mon","%d");
            
            
            return format;
        }

        public override string QuoteField(string fieldName)
        { 
            return $"`{fieldName}`";
        }
    }
    public class MsSql_Dialect :Sql_Dialect
    {
        public MsSql_Dialect()
        {
            TRUE = "1";
            FALSE = "0";
            NULL = "null";
        }

        public override bool isNoSql(){ return false;}
        public override string GetDialectName(){ return "mssql";}
        public override string convertGolangTypeToSqlType( 
            string typeName, ParserRuleContext ctx, SqlConvert convertInst)
        {
            typeName = typeName.Replace("@", ""); 
            switch(typeName)
            { 
                case "string": return "NVARCHAR(MAX)";
                case "char": return "CHAR";
                case "bool": return "BIT";

                case "byte": return "TINYINT";
                case "int16": return "SMALLINT";
                case "int": return "INT";
                case "int32": return "INT";
                case "int64": return "BIGINT";

                case "time.Time": return "TIMESTAMP";
                case "Time": return "datetime2";

                case "UUID": return "VARCHAR(MAX)";
                case "uuid.UUID": return "VARCHAR(MAX)";

                case "float32": return "FLOAT";
                case "float64": return "REAL";
                //case "float64": return "DECIMAL";
            }
            return convertInst.getTextSQLError("type cannot be converted", ctx);
        }

        public override string convertFormatTime_GolangToSql( string format1)
        {
            var format = format1;
            format = format.Replace("2006","yyyy");
            format = format.Replace("01","mm");
            format = format.Replace("02","dd");

            format = format.Replace("15","HH");
            format = format.Replace("04","mm");
            format = format.Replace("05","ss");

            format = format.Replace(".000000","US");
            format = format.Replace(".000","MS");

            format = format.Replace("Jan","Mon");
            format = format.Replace("January","Month");
            format = format.Replace("Mon","Day");
            
            return format;
        }

        public override string QuoteField(string fieldName)
        { 
            return $"[{fieldName}]";
        }
    }



    public class MongoDBNoSql_Dialect :Sql_Dialect
    {
        public MongoDBNoSql_Dialect()
        {
            TRUE = "true";
            FALSE = "false";
            NULL = "null";
        }

        public override bool isNoSql(){ return true;}
        public override string GetDialectName(){ return "mongodb";}
        public override string convertGolangTypeToSqlType( 
            string typeName, ParserRuleContext ctx, SqlConvert convertInst)
        {
            typeName = typeName.Replace("@", ""); 

            switch(typeName)
            { 
                case "string": return "String";
                case "char": return "String";
                case "bool": return "Boolean";

                case "byte": return "Int32";
                case "int16": return "Int32";
                case "int": return "Int32";
                case "int32": return "Int32";
                case "int64": return "Integer";

                case "time.Time": return "Timestamp";
                case "Time": return "Timestamp";

                case "UUID": return "String";
                case "uuid.UUID": return "String";

                case "float32": return "Decimal128";
                case "float64": return "Decimal128";
            }
            return convertInst.getTextSQLError("type cannot be converted", ctx);
        }

        public override string convertFormatTime_GolangToSql( string format1)
        {
            var format = format1;
            format = format.Replace("2006","%Y");
            format = format.Replace("15","%H");
            format = format.Replace("01","%m");
            format = format.Replace("02","%d");
            format = format.Replace("04","%M");
            format = format.Replace("05","%S");

            format = format.Replace(".000000","FF6");
            format = format.Replace(".000","FF3");

            format = format.Replace("Jan","%B");
            format = format.Replace("January","%B");
            format = format.Replace("Mon","%w");
            
            return format;
        }
        public override string QuoteField(string fieldName)
        { 
            return $"\"{fieldName}\"";
        }
    }
}
