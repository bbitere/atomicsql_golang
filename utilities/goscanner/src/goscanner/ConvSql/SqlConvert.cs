//******************************************************************************************************
//  Converter.cs - Gbtc
//
//  Copyright © 2018, Grid Protection Alliance.  All Rights Reserved.
//
//  Licensed to the Grid Protection Alliance (GPA) under one or more contributor license agreements. See
//  the NOTICE file distributed with this work for additional information regarding copyright ownership.
//  The GPA licenses this file to you under the MIT License (MIT), the "License"; you may not use this
//  file except in compliance with the License. You may obtain a copy of the License at:
//
//      http://opensource.org/licenses/MIT
//
//  Unless agreed to in writing, the subject software distributed under the License is distributed on an
//  "AS-IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. Refer to the
//  License for the specific language governing permissions and limitations.
//
//  Code Modification History:
//  ----------------------------------------------------------------------------------------------------
//  05/01/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using Antlr4.Runtime;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using goscanner.Metadata;
//using goscanner.Templates;
using static goscanner.Common;
using Gemstone.StringExtensions;
using System.Diagnostics;
using goscanner.ConvCommon;
using Antlr4.Runtime.Misc;
using Gemstone.Collections.CollectionExtensions;

#pragma warning disable SCS0018 // Path traversal

namespace goscanner.ConvSql;

/// <summary>
/// Represents a converter used to convert Go source code to C#.
/// </summary>
/// 

public class ESqlOutputType
{
    public const string Postgres = "postgres";
    public const string Mysql = "mysql";
    public const string Mssql = "mssql";

    public const string MongoDB = "mongodb";
}

//this are used in def of SQL_Method()
class SqlFuncDef
{
    public const string Postgres = "Postgres";
    public const string Mysql = "Mysql";
    public const string Mssql = "Mssql";
}
    

public delegate void FnTagUpdate( string tagName);

public partial class SqlConvert : goscanner.ConvCommon.ConvCommon
{
    public const string sql_NULL = "NULL";
    public static int  s_SubQueries_uniqueID = 0;
    
    
    
    public SqlConvert(BufferedTokenStream tokenStream, 
        GoParser parser, Options options, string currentFileName) 
        : base(tokenStream, parser, options, currentFileName)
    {

    }

    protected void setLambdaCode(TLambdaCode lamdbaCode )
    {
        if( m_LambdaCode != lamdbaCode )
        {
            if( lamdbaCode != null && lamdbaCode.Tag == "tsql082")
                Utils.Nop();

            if( m_LambdaCode != null && m_LambdaCode.ParentLambda != null
                && m_LambdaCode.ParentLambda == lamdbaCode)
            {
                var top = this.Lambda_getTopQueryTag();
                if( top != null && top.LambdaCodeContainer == lamdbaCode)
                {   //exit from a function, discard current subtag
                    Lambda_popQueryTag(lamdbaCode);
                }
            }
            m_LambdaCode = lamdbaCode;
        }
    }
    public bool AddLambda(TLambdaCode lambda)
    {
        if( lambda.ctxTag != null ) 
        {
            foreach( var it in m_exportedLambdas)
            {
                if( it.ctxTag == lambda.ctxTag && it.SubTag == lambda.SubTag )
                    return false;
            }
        }
        m_exportedLambdas.Add(lambda);  
        return true;
    }

        public void Log_Error( ParserRuleContext ctx, string message)
    {
        Log_Error( ctx.Start, message);
    }
    public void Log_Error( IToken loc, string message)
    {        
        var location = $"{loc.InputStream.SourceName}:{loc.Line}:{loc.Column}: ";
        Console.WriteLine(location + message);

        Options.ConvertSql.NumErrors++;
    }
    
    public class DataTranslation: DataBaseTranslation
    {
        public string textSqlQuery = "";
    }
    private static List<SqlConvert> s_scanners= new List<SqlConvert> ();
    protected override void collectScannerForTranslation()
    {
        s_scanners.Add(this);
    }
    
    
    public static SqlConvert s_currentScanner = null;
    public static DataTranslation DoTranslation()
    {
        var data = new DataTranslation();

        try
        {
            foreach( var scanner in s_scanners ) 
            {
                s_currentScanner = scanner;

                scanner.scannerTranslation(data);
            }
        }
        catch(Exception e)
        {
            Console.WriteLine( e.Message );
        }
        return data;
    }

    private void scannerTranslation(DataTranslation data)
    {
        if( m_exportedLambdas.Count > 0 ) 
        {
            var fileName = Utils1.getFileName(SourceFileName);
            Console.WriteLine($"{fileName}: Found {m_exportedLambdas.Count } lambda expressions");
            Console.WriteLine($"At '{this.SourceFileName}'");
            Console.WriteLine("");

            var txtSqlQuery = this.GoLang_ExportAllQueries( m_exportedLambdas, Options.ConvertSql);
            //File.AppendAllText(this.Options.ConvertWebApi.TsSvcModelsFile, txtTsSvcModels);
            data.textSqlQuery += txtSqlQuery;
        }else
        {
            //Console.WriteLine($"No lambda expression found!");
        }
        Console.WriteLine("");

             
        /*
        var txtGoSvc = goscanner.ConvWebApi.WebApiBase_Go.exportSvc_AllMethods( m_methods, this.Options.ExportWebApi_GitHubControllers);
        File.AppendAllText(this.Options.ExportWebApi_GoSvcFile, txtGoSvc);

        var txtTsSvc = goscanner.ConvWebApi.WebApiBase_Ts.exportSvc_AllMethodsTs( m_methods, this.Options.ExportWebApi_TsSvcModelsFile);
        File.AppendAllText(this.Options.ExportWebApi_TsSvcFile, txtTsSvc);

        var txtTsSvcModels = goscanner.ConvWebApi.WebApiBase_Ts.exportSvcModels_AllMethodsTs( m_models, m_methods);
        File.AppendAllText(this.Options.ExportWebApi_TsSvcModelsFile, txtTsSvc);
        */
        
    }

    public static void Convert(Options options)
    {
        if (options.OnlyUpdateMetadata)
            return;

        if(options.ExportWebApi_ConfigFile != "" )
        { 
            if (options.ConvertWebApi.GoSvcFile == "")
            { 
                Console.WriteLine("file output GoSvc is node defined");
                return;
            }
        }else
        if(options.ExportSql_ConfigFile != "" )
        { 
            if (options.ConvertSql.SqlOutputFile == "")
            { 
                Console.WriteLine("file output is node defined");
                return;
            }
        }else
        {
            Console.WriteLine("file config is not defined");
            return;
        }

        ResetScanner();
        var data = Scan(true, options, options.ShowParseTree, CreateNewConverter);
        //WriteProjectFiles(options);

        writeOutputFiles( options, data );        
    }
    private static void writeOutputFiles(Options options, DataBaseTranslation dataBase)
    {
        var data = dataBase as DataTranslation;
        if( data == null 
        || (data.textSqlQuery == "") )
            return;

        //if(! s_exist_CLASS_GenericModel )
        //{
        //    Console.WriteLine("Generic_MODEL struct not defined. if you export models, you need to have it!");
        //}
        //var allMeths = File.ReadAllText(options.ConvertWebApi.GoSvcFile);

        var allQueries = data.textSqlQuery;
        var txtGoSql = SqlConvert.GoLang_ExportSqlFile( options, allQueries);
        File.WriteAllText( options.ConvertSql.SqlOutputFile, txtGoSql);
        Console.WriteLine($"Generate file: {options.ConvertSql.SqlOutputFile}");
        Console.WriteLine("");
        
    }

    private static ScannerBase CreateNewConverter(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName)
    {
        return new SqlConvert(tokenStream, parser, options, fileName);
    }

    private static void WriteProjectFiles(Options options)
    {
    #if !DEBUG
        try
        {
    #endif
        // Map of package names to list of package path and file names
        Dictionary<string, List<(string path, string[] fileNames)>> groupedPackageData;

        // Process import packages - these become shared projects
        groupedPackageData = CreateGroupedPackageData();

        // Process packages with "main" functions - these become standard projects
        ProcessMainProjectPackages(options);

        if (options.ConvertStandardLibrary && options.RecurseSubdirectories && AddPathSuffix(options.SourcePath).Equals(GoPath))
            ProcessStandardLibraryPackages(options, groupedPackageData);
    #if !DEBUG
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Failed to write project files: {ex.Message}");
        }
    #endif
    }

    protected ParameterInfo[] new_ParameterInfo_TypeInfo(TypeInfo type)
    {
        return new ParameterInfo[]{ new ParameterInfo(){
            Name = "p0",
            Type = type
        }
        };
    }

    int errorNumb = 100;
    public string getTextSQLError(string errorMsg, ParserRuleContext ctx)
    {
        if( m_LambdaCode != null && !m_LambdaCode.IsNoSql)
        {
            Log_Error( ctx.Start, errorMsg);

            errorNumb++;
            return $"<error {errorNumb}>";
        }
        return "";
    }
    public string getTextNoSQLError(string errorMsg, ParserRuleContext ctx)
    {
        if( m_LambdaCode != null && m_LambdaCode.IsNoSql)
        {
            //if( false)
            {
                Log_Error( ctx.Start, errorMsg);

                errorNumb++;
                return $"<error {errorNumb}>";
            }
        }
        return "";
    }
    protected string convertGolangStringToSqlString( string str)
    {
        return ConvertGolangStringToSqlString( str, this.Options.ConvertSql.SqlLang);
    }

    protected static string ConvertGolangStringToSqlString( string str, string SqlLang)
    {
        var prefix = false;
        var s = "";
        var s2 = "";
        if( str.StartsWith("\"")
         && str.EndsWith("\"") )
        {   
            str = str.Substring(1, str.Length - 2*1);
        }

        for(var i = 0; i < str.Length; i++ )             
        {   
            var ch = str[i];
            if( i + 1 < str.Length)
            {
                var ch1 = str[i + 1];   
                if( ch == '\\' && ch1 == '\\')
                {
                    s += "\\";
                    s2 += "\\";
                    i++;
                }else
                if( ch == '\\' && ch1 == '"')
                {
                    s += "\"";
                    s2 += "\"";
                    i++;
                }else
                if( ch == '\\' && ch1 == 'r')
                {
                    prefix = true;
                    s += "\\r";
                    s2 += "'+CHAR(0x0A)+'";
                    i++;
                }else
                if( ch == '\\' && ch1 == 'n')
                {
                    prefix = true;
                    s += "\\n";
                    s2 += "'+CHAR(0x0D)+'";
                    i++;
                }else
                {
                    s += ch;
                    s2 += ch;
                    if( ch == '\'' )
                    {
                        //duplicate
                        s += "'";
                        s2 += "'";
                    }
                }
            }else
            {
                s += ch;
                s2 += ch;
                if( ch == '\'' )
                {
                    //duplicate
                    s += "'";
                    s2 += "'";
                }
            }
        }
        
        if( prefix )
        {
            if( SqlLang == ESqlOutputType.Postgres ) 
                return $"E'{s}'";

            if( SqlLang == ESqlOutputType.Mysql ) 
                return $"'{s}'";

            if( SqlLang == ESqlOutputType.Mssql )
                return $"'{s2}'";
            else
                return $"'{s}'";
        }
        else
        {
            return $"'{s}'";
        }
    }
    bool unused = SqlConvert.test_convertGolangStringToSqlString();
    static bool test_convertGolangStringToSqlString()
    {
        var s = SqlConvert.ConvertGolangStringToSqlString("a'\"\\n mm", ESqlOutputType.Postgres);
        if( s != "E'a''\"\\n mm'")
        {
            Debugger.Break();
            return false;
        }
        var s1 = SqlConvert.ConvertGolangStringToSqlString("a'\"\\n mm", ESqlOutputType.Mysql);
        if( s1 != "'a''\"\\n mm'")
        {
            Debugger.Break();
            return false;
        }
        return true;
    }

    public class TTypeConv
    {

        public string stringCollation;
    }

    protected string convertToGolangType( string typeName, ParserRuleContext ctx
        //, TTypeConv details
        )
    {
        var typeName1 = typeName.Replace("@", ""); 
        typeName1 = typeName1.Replace("*", ""); 


        return typeName1;
    }

    

    protected class TTextSql
    {
        public string defaultSql{ get; set;} 
        public string postgresSqlText{ get; set;}
        public string mySqlText{ get; set;}
        public string msSqlText{ get; set;}
    }

    protected bool findSqlSubqueryMethod( string funcName, string packageName, TypeInfo normalizedTypeReceiver )
    {
        if( OrmDef.SubQueryTransform.ContainsKey(funcName) )
        {
            return true;
        }
        return false;
    }
    
    protected (TTextSql,FunctionInfo) findSqlNativeMethod( string funcName, string packageName, 
                    TypeInfo normalizedTypeReceiver )
    {
        
        if( funcName.StartsWith(OrmDef.FuncSql_PrefixMeth) )
        {
            if( ImportMetadata.TryGetValue( packageName, out FolderMetadata metaImport))
            {
                foreach( var itemMetadata in metaImport.Files)
                {
                    //if( itemMetadata.Value.Functions.TryGetValue($"{functionName}()", out functionInfo))
                    if( itemMetadata.Value.Functions.TryGetValue(funcName, out FunctionInfo functionInfo))
                    {
                        var varName = funcName.Replace( OrmDef.FuncSql_PrefixMeth, OrmDef.FuncSql_PrefixVAR);
                        if( itemMetadata.Value.GlobalVariables.TryGetValue( varName, out VariableInfo varInfo))
                        {
                            var sqlFunc = new TTextSql();
                            var initExpr = varInfo.getInitExpr();
                            sqlFunc.postgresSqlText = extractString( initExpr[SqlFuncDef.Postgres] );
                            sqlFunc.mySqlText       = extractString( initExpr[SqlFuncDef.Mysql] );
                            sqlFunc.msSqlText       = extractString( initExpr[SqlFuncDef.Mssql] );
                    
                            return (sqlFunc, functionInfo);
                        }
                        return (null,null);
                    }
                }
            }
        }

        if( normalizedTypeReceiver != null && normalizedTypeReceiver.Name == OrmDef.myTTime )
        {
            if( ImportMetadata.TryGetValue( packageName, out FolderMetadata metaImport))
            {
                foreach( var itemMetadata in metaImport.Files)
                {
                    if( itemMetadata.Value.Functions.TryGetValue(funcName, out FunctionInfo functionInfo))
                    {
                        var varName = $"{OrmDef.FuncSql_PrefixVAR}{funcName}";
                        if( itemMetadata.Value.GlobalVariables.TryGetValue( varName, out VariableInfo varInfo))
                        {
                            var sqlFunc = new TTextSql();
                            var initExpr = varInfo.getInitExpr();
                            sqlFunc.postgresSqlText = extractString( initExpr["Postgres"] );
                            sqlFunc.mySqlText = extractString( initExpr["Mysql"] );
                            sqlFunc.msSqlText = extractString( initExpr["Mssql"] );
                    
                            return (sqlFunc, functionInfo);
                        }
                    }
                }
            }
        }

        return (null, null);
    }

    protected string getTextSQL( ParserRuleContext ctx, TTextSql arg )
    {
        switch( Options.ConvertSql.SqlLang )
        {
            case ESqlOutputType.Postgres:
                return arg.postgresSqlText != null?arg.postgresSqlText:arg.defaultSql;

            case ESqlOutputType.Mysql:
                return arg.mySqlText != null? arg.mySqlText: arg.defaultSql;

            case ESqlOutputType.Mssql:
                return arg.msSqlText != null? arg.msSqlText: arg.defaultSql;

            case ESqlOutputType.MongoDB:

                if( m_LambdaCode != null)
                {
                    Log_Error( ctx, "calling methods is not allowed in nosql query");
                }
                return "";
                
            default:
            {
                Debugger.Break();
                return "";
            }
        }
    }

    //`json:"-"`
    //`json:"userrole_id"`

    protected string getSqlField_ExtractFromJsonAnnotation( Sql_Dialect dialect, FieldInfo fld )
    {
        var descriptionsTxt = fld.Description;
        if( descriptionsTxt != null ) 
        {
            if( descriptionsTxt.StartsWith("\"")
             && descriptionsTxt.EndsWith("\"") )
            {
                descriptionsTxt = descriptionsTxt.Substring(1, descriptionsTxt.Length -2);
            }
            var descriptions = descriptionsTxt.Split( new String[]{ "  " }, StringSplitOptions.RemoveEmptyEntries );
            foreach( var description1 in descriptions )
            {
                var description = description1.Trim();
                if( !dialect.isNoSql() && description.StartsWith(OrmDef.Atomicsql_atmsql_fld) )
                {
                    var fieldJson = description.Replace(OrmDef.Atomicsql_atmsql_fld, "");
                    if( fieldJson.StartsWith("\"")
                     && fieldJson.EndsWith("\"") )
                    {
                        fieldJson = fieldJson.Substring(1, fieldJson.Length -2);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                    if( fieldJson.StartsWith("\\\"")
                     && fieldJson.EndsWith("\\\"") )
                    {
                        fieldJson = fieldJson.Substring(2, fieldJson.Length -4);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                }
                if( !dialect.isNoSql() && description.StartsWith(OrmDef.Atomicsql_json_fld) )
                {
                    var fieldJson = description.Replace(OrmDef.Atomicsql_json_fld, "");
                    if( fieldJson.StartsWith("\"")
                     && fieldJson.EndsWith("\"") )
                    {
                        fieldJson = fieldJson.Substring(1, fieldJson.Length -2);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                    if( fieldJson.StartsWith("\\\"")
                     && fieldJson.EndsWith("\\\"") )
                    {
                        fieldJson = fieldJson.Substring(2, fieldJson.Length -4);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                }

                if( dialect.isNoSql() && description.StartsWith(OrmDef.MongoDB_bson_fld) )
                {
                    var fieldJson = description.Replace(OrmDef.MongoDB_bson_fld, "");
                    if( fieldJson.StartsWith("\"")
                     && fieldJson.EndsWith("\"") )
                    {
                        fieldJson = fieldJson.Substring(1, fieldJson.Length -2);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                    if( fieldJson.StartsWith("\\\"")
                     && fieldJson.EndsWith("\\\"") )
                    {
                        fieldJson = fieldJson.Substring(2, fieldJson.Length -4);
                        if( fieldJson == "-")
                            return null;

                        var parts = fieldJson.Split(',');
                        return parts[0];
                    }
                }
            }
        }
        return fld.Name;
    }

    protected string getSqlTextForField( Sql_Dialect dialect, FieldInfo field, StructInfo structInfo, ref TypeInfo typeInfo, bool bAllowPointer)
    {
        var sqlField = "";
        if( field.Type.IsPointer() )
        {
            if( !bAllowPointer)
                return null;

            var iFld = Array.IndexOf( structInfo.Fields, field );
            if( iFld >= 0 )
            {
                typeInfo = field.Type;

                sqlField = getSqlField_ExtractFromJsonAnnotation( dialect, field );
                if( sqlField == null //daca este ForeignKey
                    && field.Type.Name.StartsWith("*")
                    && iFld+1 < structInfo.Fields.Length )
                {
                    sqlField = getSqlField_ExtractFromJsonAnnotation( dialect, structInfo.Fields[iFld+1] );
                }
            }else
                Debugger.Break();
            
        }else
        {
            sqlField = getSqlField_ExtractFromJsonAnnotation(  dialect, field );
        }
        //if( sqlField == null)
        //    Debugger.Break();
        return sqlField;
    }

    //$"{identif1}.{sqlField}"
    protected string getTextSQLIdentif(string identifExpr, string sqlIdentif, string sqlField, TypeInfo type, string fldName, ParserRuleContext ctx, string operandDOT)
    {
        var str = this.Lambda_getSQLIdentif( identifExpr,  sqlIdentif, sqlField, type, fldName, ctx, operandDOT);
        if( str != null ) 
        {
            return str;
        }
        return $"{PREFIX_FIELD}{sqlIdentif}.{sqlField}{POSTFIX_FIELD}";
    }
    protected TNoSqlCode getTextNoSQLIdentif(string identifExpr, string sqlIdentif, string sqlField, TypeInfo type, string fldName, ParserRuleContext ctx, string operandDOT)
    {
        var str = this.Lambda_getSQLIdentif( identifExpr,  sqlIdentif, sqlField, type, fldName, ctx, operandDOT);
        if( str != null ) 
        {
            return new TNoSqlCode( str );
        }
        return new TNoSqlCode($"{PREFIX_FIELD}{sqlIdentif}.{sqlField}{POSTFIX_FIELD}" );
    }

    protected string getTextSQLVarIdentif(string identif, TypeInfo type, ParserRuleContext ctx)
    {
        var topSubQuery = this.GetTopSubquery();
        if( topSubQuery != null && topSubQuery.VariableStorageName == identif)
        {
            return $"{START_SUBQUERY}{identif}{END_SUBQUERY}";
        }

        var txtSql = this.Lambda_getSQLVarIdentif( identif, type, ctx);
        if( txtSql != null ) {
            return txtSql;
        }
        if(identif == "ids")
             Utils.Nop();
        return $"{PREFIX_VAR}{identif}{POSTFIX_VAR}";
    }

    private TNoSqlCode NSqlConvertToBool( TNoSqlCode arg, EOperandKind bIsOperator)
    {
        if( bIsOperator != EOperandKind.Operator )
            return new TNoSqlCode("==", arg, TNoSqlCode.TRUE );
                                                             
        return arg;
    }
    private string convertToBool( string arg, EOperandKind bIsOperator )
    {
        if( this.Options.ConvertSql.SqlLang == ESqlOutputType.Mysql)
        {
            if( bIsOperator != EOperandKind.Operator )
                return $"(arg)=1";

            return arg;
        }
        return arg;
    }

    private string getSqlTableNameFromStructInfo( StructInfo structInfo)
    {
        if( structInfo.Comment != null && structInfo.Comment != "" ) 
        { 
            var comment = structInfo.Comment;
            if( comment.StartsWith(OrmDef.Atomicsql_table) )
            {
                comment = comment.Substring(OrmDef.Atomicsql_table.Length);
            }
            if( comment.StartsWith("\"")
             && comment.EndsWith("\"") )
            {
                comment = comment.Substring( 1, comment.Length - 2*1 );
            }
            return comment;
        }
        return "";
    }

    private string extractString(string s)
    {
        if( s.StartsWith("\"")
            && s.EndsWith("\"") ) 
        {
            s = s.Substring( 1, s.Length - 2*1 );
        }
        return s;
    }

    private string _quoteFld(string fieldName)
    {
        
        return this.Options.ConvertSql.SqlDialect.QuoteField(fieldName);;
    }

    private ( bool, string, TNoSqlCode) Select_generateItem( 
        Dictionary<string,string> select_Fields,
        string selectFieldName,  
        ParserRuleContext expression, 
        ParserRuleContext context, TypeInfo selectType )
    {
        var sqlText = "";
        TNoSqlCode noSQLText = null;
        var bIsNoSql = false;
        TypeInfo typeExpr = null;
        var operandDOT = "";

        if( Expressions.TryGetValue( expression, out var exprElem) ) 
        {
            sqlText = exprElem.SQLText;
            noSQLText = exprElem.NoSQLCode;
            bIsNoSql = exprElem.bIsNoSql;
            operandDOT =  exprElem.Text;
            typeExpr = exprElem.Type;
        }else
        if( Operands.TryGetValue( expression, out var operand) ) 
        {
            sqlText = operand.SQLText;
            noSQLText = exprElem.NoSQLCode;
            bIsNoSql = exprElem.bIsNoSql;
            operandDOT =  operand.Text;
            typeExpr = operand.Type;
        }
        else
        if( PrimaryExpressions.TryGetValue( expression, out var primaryExpr) ) 
        {
            sqlText = primaryExpr.SQLText;
            noSQLText = exprElem.NoSQLCode;
            bIsNoSql = exprElem.bIsNoSql;
            operandDOT =  primaryExpr.Text;
            typeExpr = primaryExpr.Type;
        } 
        else
        if( UnaryExpressions.TryGetValue( expression, out var unaryExpr) ) 
        {
            sqlText = unaryExpr.SQLText;
            noSQLText = exprElem.NoSQLCode;
            bIsNoSql = exprElem.bIsNoSql;
            operandDOT =  unaryExpr.Text;
            typeExpr = unaryExpr.Type;
        }

        if( sqlText != "" )
        {
            //var ITM = OrmDef.SELECT_STRCT_ITM;
            var selectStructInfo = getStructInfoByType( selectType );
            if( selectStructInfo != null ) 
            {
                var fld = selectStructInfo.Fields.Where(x=>x.Name == selectFieldName).FirstOrDefault();
                if( fld != null )
                {
                    var genType2 =  fld.Type.getNormalizedType();
                    var fldStructInfo = getStructInfoByType( fld.Type );

                    if( fldStructInfo != null )
                    { 
                        if( fld.IsPromoted && fld.Description.Contains( OrmDef.Atomicsql_CopyModel1) )
                        {
                            var arr = new List<string>();   
                            var arrNoSql = new List<TNoSqlCode>();   

                            foreach( var f1 in fldStructInfo.Fields) 
                            {
                                if( f1.Name == OrmDef.Generic_MODEL_Name)
                                    continue;
                                if( f1.Name == OrmDef.NoSqlID_Name)
                                    continue;

                                TypeInfo fldType2= fld.Type;
                                var sqlField = getSqlTextForField( Options.ConvertSql.SqlDialect, f1, fldStructInfo, ref fldType2, false );//{f1.Name}
                                if( sqlField == null )
                                    continue;// do not export pointer field. only the foreignbkey_id

                                var sqlFieldIdentif = getTextSQLIdentif(null, "", sqlField, fldType2, f1.Name, context, operandDOT);

                                //arr.Add( $" {ITM}.{sqlField} as {f1.Name}" );
                                select_Fields[f1.Name] = sqlFieldIdentif;
                                arr.Add( $"{sqlFieldIdentif} AS {_quoteFld(f1.Name)}" );

                                arrNoSql.Add( new TNoSqlSelectField( f1.Name, sqlFieldIdentif ) );
                            }
                            return (true, String.Join(", ", arr ), new TNoSqlCode("select", arrNoSql.ToList() ) );
                        }else
                        {
                            return (false, getTextSQLError($"copy entire struct not allowed {genType2.Name})", context), null);
                        }
                    }else
                    {
                        select_Fields[selectFieldName] = sqlText;
                        return (true, $"{sqlText} AS {_quoteFld(selectFieldName)}", noSQLText);
                    }
                }else
                {
                    return (false, getTextSQLError($"internal error 1254 )", context), null );
                }
            }else
            {
                select_Fields[selectFieldName] = sqlText;
                return (true, $" {sqlText} AS {_quoteFld(selectFieldName)}", noSQLText );
                //return (false, getTextSQLError($"internal error 347 (not found struct {genType.Name})", context));
            }
        }else
        {
            return (false, getTextSQLError("internal error 345 (expressio not found)", context), null );
        }
    }

    private string Aggregate_checkArgumentsCall(string typeT, string typeV)
    {
        var structT = getStructInfoByName2(typeT);
        var structV = getStructInfoByName2(typeV);
        if( structT != null && structV != null ) 
        {
            //I updaet the comment because here is store the table sql name. for the aggrate function, the both models point to the same table
            structV.Comment = structT.Comment;//copy the flag: OrmDef.Atomicsql_table 

            var dictFld = new Dictionary<string, FieldInfo >();

	        for( var i= 0; i < structT.Fields.Length; i++)
            {
		        var fld = structT.Fields[i];
		        dictFld[ fld.Name ] = fld;
            }
	
	        var selectFields  = new List<string>();
	        var groupByFields = new List<string>();

            for( var i= 0; i < structV.Fields.Length; i++)
            {
		        var fldV = structV.Fields[ i ];
		        
		        if( !dictFld.ContainsKey( fldV.Name ) )
                {
			        return $"Aggregate() incompatibile struct: {structT.Name} - {structV.Name}. Field not found {fldV.Name}" ;
		        }
                var fldT = dictFld[ fldV.Name ];

                fldV.Description = fldT.Description;
		
		        selectFields.Add( fldV.Name );

                var type1 = fldV.Type.getNormalizedType();
                var type2 = fldT.Type.getNormalizedType();

		        if( type1.Name == type2.Name )
                {
			        groupByFields.Add(fldV.Name  );
		        }else
		        if( type1.Name == ArrayTypeInfo.ArrName                      
                    && (type1 as ArrayTypeInfo).TargetTypeInfo?.Name == type2.Name )
                {
			        //here are the fields that will be aggregate
		        }else
                {
			        return $"Aggregate() incompatibile struct: {structT.Name} - {structV.Name}. For field {fldV.Name} the types are incompatible ";
                }
		    }
	        if( selectFields.Count == groupByFields.Count )
            {
		        return $"Aggregate() incompatibile struct for aggregate: {structT.Name} - {structV.Name}. No fields to aggregate. you must have minimum 1 field with same name and different type ([]Type) in both structs ";
	        }
	        if( groupByFields.Count == 0  )
            {
		        return $"Aggregate() incompatibile struct for aggregate: {structT.Name} - {structV.Name}. No fields to group by. you must have minimum 1 field with same name and type in both structs ";
            }
        }
        return "";
    }


    protected bool variableTypesTryGetValue( string nameVar, out LocalVariableInfo typeVar)
    {
        if( m_globalVariables.TryGetValue( nameVar, out VariableInfo varGlobal))
        {
            typeVar = new LocalVariableInfo( varGlobal.Type, false);
            return true;
        }

        if( m_variableTypes.TryGetValue( nameVar, out typeVar))
        {
            return true;
        }
        foreach( var it in m_variableTypes ) 
        {
            var varsName = it.Key.Trim();
            if( varsName.StartsWith("(") ) 
            {
                varsName = varsName.Substring(1);
                varsName = varsName.Substring(0,varsName.IndexOf(")"));
                var parts = varsName.Split(",");
                foreach( var part_ in parts) 
                {
                    var part = part_.Trim();
                    if( part == nameVar)
                    {
                        typeVar = new LocalVariableInfo();
                        typeVar.IsParameter = it.Value.IsParameter;
                        typeVar.Type = it.Value.Type;
                        return true;
                    }
                }
            }
        }
        return false;
    }
       

}
