using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using goscanner.ConvSql;
using goscanner.Metadata;

namespace goscanner.ConvSql
{
    internal static class OrmDef
    {
        public const string OrmPackageName1      = "atomicsql_func";

        public const string Generic_MODEL      = "Generic_MODEL";



        //public const string SELECT_STRCT_ITM    = "##slItm##";

        public const string Class_DBTable       = "DBTable";
        public const string Func_DBTable_Qry    = "Qry";

        public const string Class_DBContext     = "DBContext";
        public const string Func_New_DBContext  = "New_DBContext";
        public const string Func_Select         = "Select";    //Select[T IGeneric_MODEL, V IGeneric_MODEL]
        public const string Func_Aggregate      = "Aggregate";  //Aggregate[T IGeneric_MODEL, V IGeneric_MODEL]

        public const string Class_DBQuery       = "DBQuery";
        public const string Func_DBQuery_Where  = "Where";

        public const string SubTag_Where = "W";
        public const string SubTag_Select = "S";
        public const string SubTag_GetValue  = "V";
        public const string SubTag_GetValues = "X";

        public static readonly string[] Func_DBQuery_GetValueArr = new string[]
        {
            "GetValueString", "GetValueInt", "GetValueFloat","GetValueBool", "GetValueTime"
        };
        public static readonly string[] Func_DBQuery_GetValuesArr = new string[]
        {
            "GetValuesString", "GetValuesInt", "GetValuesFloat","GetValuesBool", "GetValuesTime"
        };
        // each Query should end in one of this function
        public static readonly string[] Func_DBQuery_End = new string[]
        { 
            "GetModels","GetFirstModel",
            "GetModelsRel", "GetFirstModelRel",
            "GetRecords", "GetFirstRecord", 
            "GetCount", "GetDistinctCount", "GetDistinct1Count",
            "GetSingleDataS", "GetSingleDataI", "GetSingleFieldRows",//deprecated
            "GetSingleDataString", "GetSingleDataInt",
            "GetRowsAsFieldString", "GetRowsAsFieldInt",
            "GetDistinctModels",//deprecated
            "GetDistinctRecords",
            "DeleteModels", "DeleteModel",
            "GetValueString", "GetValueInt", "GetValueFloat","GetValueBool", "GetValueTime",
            "GetValuesString", "GetValuesInt", "GetValuesFloat","GetValuesBool", "GetValuesTime",
            "ToRTM"
        };
        

        public const string Atomicsql_table     = "atomicsql-table:";
        public const string Atomicsql_json_fld  = "json:";
        public const string Atomicsql_CopyModel = "atomicsql:\"copy-model\"";
        public const string Atomicsql_CopyModel1 = @"""atomicsql:\""copy-model\""";


        public const string Sql_NullTime    = "NullTime";
        public const string Sql_NullByte    = "NullByte";
        public const string Sql_NullInt16   = "NullInt16";
        public const string Sql_NullInt32   = "NullInt32";
        public const string Sql_NullInt64   = "NullInt64";
        public const string Sql_NullFloat64 = "NullFloat64";
        public const string Sql_NullString  = "NullString";
        public const string Sql_NullBool    = "NullBool";

        public const string Sql_Null_Valid  = "Valid";

        public const string Sql_NullTime_Val    = "Time";
        public const string Sql_NullByte_Val    = "Byte";
        public const string Sql_NullInt16_Val   = "Int16";
        public const string Sql_NullInt32_Val   = "Int32";
        public const string Sql_NullInt64_Val   = "Int64";
        public const string Sql_NullFloat64_Val = "Float64";
        public const string Sql_NullString_Val  = "String";
        public const string Sql_NullBool_Val    = "Bool";



        public const string FuncSql_PrefixMeth = "Sql_";
        public const string FuncSql_PrefixVAR  = "SQL_";

        public const string myTTime             = "TTime";
        public const string myTTime_Day         = "Day";
        public const string myTTime_Month       = "Month";
        public const string myTTime_Year        = "Year";
        public const string myTTime_YearDay     = "YearDay";
        public const string myTTime_Hour        = "Hour";
        public const string myTTime_Minute      = "Minute";
        public const string myTTime_Second      = "Second";
        public const string myTTime_Nanosecond  = "Nanosecond";
        public const string myTTime_UnixMilli   = "UnixMilli";
        public const string myTTime_Compare     = "Compare";
        public const string myTTime_Format      = "Format";
        

        public static readonly TypeInfo myTTimeType =
            new()
            {
                Name = $"{OrmDef.OrmPackageName1}.{myTTime}",
                TypeName = $"{OrmDef.OrmPackageName1}.{myTTime}",
                FullTypeName = $"{OrmDef.OrmPackageName1}.{myTTime}",
                TypeClass = TypeClass.Simple
            };
        
        public static string GetSubTabByFuncName(string funcName)
        {
            if( OrmDef.Func_DBQuery_GetValueArr.Contains(funcName) )
                return SubTag_GetValue;

            if( OrmDef.Func_DBQuery_GetValuesArr.Contains(funcName) )
                return OrmDef.SubTag_GetValues;

            if( Func_DBQuery_Where == funcName)
                return OrmDef.SubTag_Where;

            if( Func_Select == funcName)
                return OrmDef.SubTag_Select;
            return "";
        }
        
    }
}
