using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace src_tool
{
partial class MongodbDialect
{
        
    public override string getGoLangTypeIntFk( 
        DbColumn column, ref Dictionary<string, string> packageImports )
    {
        var nameSqlType = column.sqlType;
        /*
        var nameArr = "";
        if( nameSqlType.Contains("[]") )
        {
            nameArr = "[]";
            nameSqlType = nameSqlType.Replace("[]", "");
        }*/

        if( column.bIsNullable )
        {
            packageImports["sql"] = ("sql \"database/sql\"");
            switch( nameSqlType)
            { 
                case "smallint" :   return "sql.NullInt16";
                case "smallserial": return "sql.NullInt16";

                case "integer" :    return "sql.NullInt32";
                case "serial" :     return "sql.NullInt32";

                case "bigint" :     return "sql.NullInt64";
                case "bigserial":   return "sql.NullInt64";
            
                default: 
                {
                    Console.WriteLine($"{nameSqlType}: *** not implemented ***");
                    return "";
                }
            }
        }else
        {
            switch( nameSqlType)
            { 
                case "smallint" : return "int16";
                case "smallserial" : return "int16";

                case "integer" : return "int32";
                case "serial" : return "int32";

                case "bigint" :return "int64";
                case "bigserial" : return "int64";
            
                default: 
                {
                    Console.WriteLine($"{nameSqlType}: *** not implemented ***");
                    return "";
                }
            }
        }
    }
    
    public override string getGoLangType( DbColumn column, ref Dictionary<string, string> importPackage )
    {
        var nameSqlType = column.sqlType.ToLower();
            
        var nameArr = "";
        if( nameSqlType.Contains("[]") )
        {
            //nameArr = "[]";
            //nameSqlType = nameSqlType.Replace("[]", "");

            if( column.bIsNullable )
            {
                //importPackage["sql"] = ("sql \"database/sql\"");
                //return nameArr+"sql.NullInt16";
                return "[]uint8";
            }
            else
                return "[]uint8";
        }

        if( column.ForeignKey != null )
        {
            return nameArr+"*"+column.ForeignKey.LangTableNameModel;
        }

        switch( nameSqlType)
        { 
            case "smallint" : 
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt16";
                }
                else
                    return nameArr+"int16";
            }
            case "smallserial" : 
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt16";
                }
                else
                    return nameArr+"int16";
            }
            case "int" : 
            case "integer" : 
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt32";
                }
                else
                    return nameArr+"int32";
            }
            case "serial" : 
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt32";
                }
                else
                    return nameArr+"int32";
            }
            case "bigint" :
            case "int4" :
            { 
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt64";
                }
                else
                    return nameArr+"int64";
            }
            case "bigserial" :
            { 
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullInt64";
                }
                else
                    return nameArr+"int64";
            }
            case "uuid" :
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullString";
                }
                else
                    return nameArr+"string";
            }
            /*
            { 
                importPackage.Add("uuid \"github.com/google/uuid\"");
                return nameArr+"uuid.UUID";
            }*/
            case "timestamp" :
            case "timestamp without time zone" :
            case "date" :
            { 
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullTime";
                }
                else
                {
                    importPackage["time"] = ("time \"time\"");
                    return nameArr+"time.Time";
                }
            }            
            case "bit" :
            case "boolean" :
            case "bool" :
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullBool";
                }
                else
                    return nameArr+"bool";
            }            
            case "decimal" ://return nameArr+"decimal";
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullFloat64";
                }
                else
                    return nameArr+"float64";
            }
            case "numeric" ://return nameArr+"decimal";
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullFloat64";
                }
                else
                    return nameArr+"float64";
            }
            case "real" :
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    Console.WriteLine($"sql.NullFloat32 is not exist in golang database/sql. Change the type of field '{column.sqlName}' to double precision");
                    return nameArr+"sql.NullFloat32";
                }
                else
                    return nameArr+"float32";
            }
            case "double precision" :
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullFloat64";
                }
                else
                    return nameArr+"float64";
            }
            case "json" :
            /*
            { 
                importPackage.Add("json \"encoding/json\"");
                return nameArr+"string";
            }*/
            {
                if( column.bIsNullable )
                {
                    //importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"string";
                }
                else
                {
                    //importPackage["json"] = ("json \"encoding/json\"");
                    return nameArr+"string";
                }
            }
            case "array":
            {
                Console.WriteLine($"{nameSqlType}: *** not implemented ***");
                return "[]string";
            }
            case "character": //return nameArr+"string";
            case "char":
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullString";
                }
                else
                    return nameArr+"string";
            }
            case "character varying": //return nameArr+"string";
            case "varchar":
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullString";
                }
                else
                    return nameArr+"string";
            }
            case "text":// return nameArr+"string";
            {
                if( column.bIsNullable )
                {
                    importPackage["sql"] = ("sql \"database/sql\"");
                    return nameArr+"sql.NullString";
                }
                else
                    return nameArr+"string";
            }

            default: 
            {
                if(column.langType != null)
                    return column.langType;
                Console.WriteLine($"Sql Type '{nameSqlType}': *** not implemented ***. Use text or varchar(1024)");
                return "";
            }
        }
    }
}
}