using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using goscanner.ConvCommon;
using goscanner.Metadata;
using static System.Net.Mime.MediaTypeNames;

namespace goscanner.ConvSql;

/// <summary>
/// Represents a converter used to convert Go source code to C#.
/// </summary>
/// 

public partial class SqlConvert 
{
    const string PREFIX_FIELD = "{#@";
    const string POSTFIX_FIELD = "@#}";

    const string PREFIX_VAR = "{@@";
    const string POSTFIX_VAR = "@@}";

    static Dictionary<string, TLambdaCode> s_dictLambdaTagPerApp = new();

internal string GoLang_externVar( TField vardef)
{
    var varName = vardef.SqlName;
    var varType = vardef.TypeLang;

    var text = $@"
                    {{varName:""{varName}"", varType:""{varType}""}},";

    text = UseTemplate( text, Options.ConvertSql.Templ_GoSqlVarDef, 
        new Dictionary<string,string>()
        { 
            { "varName", varName},
            { "varType", varType},
        }
        );

    return text;
}
internal string GoLang_ExportAllQueries( 
    List<TLambdaCode> listLambdas, Sql_ConfigTranslation optionsConfig)
{
    
    foreach ( var lambda in listLambdas )
    { 
        if(lambda.Tag == null){

            Console.WriteLine( $"Lambda export: internal error, Tag = null");
            Console.WriteLine( $"{lambda.SrcFile}:{lambda.SrcFile}:1" );
            continue;
        }
        var queryTag = _getQueryTag( lambda.Tag ) + lambda.SubTag;
        if( queryTag == "")
        {
            this.Log_Error( lambda.TagCtx, $"{OrmDef.Func_DBTable_Qry}() method should receive a valid tag unique per app");
        }
        if( queryTag == "tst1340-W1")
            Utils1.Nop();

        if( !s_dictLambdaTagPerApp.ContainsKey( queryTag ))
        {
            s_dictLambdaTagPerApp[ queryTag ] = lambda;
        }else
        {
            var tag = getDefTagNameOfLambda( lambda.SrcFile, lambda.SrcLine );

            this.Log_Error( lambda.TagCtx, $"{OrmDef.Func_DBTable_Qry}() method should receive a valid tag and unique per app. We suggest to set this tag = \"{tag}\" ");
        }
    }

    var text = "";
    foreach ( var lambda in listLambdas )
    {
        if(lambda.Tag == null){
            continue;
        }
        text += GoLang_ExportQuery( lambda, optionsConfig );
    }

    return text.Trim();
}

private string convertTo1000( int SrcLine )
{
    var val = SrcLine % 1000;

    if( val < 10 )
        return "00"+val;

    if( val < 100 )
        return "0"+val;

    return ""+val;
}
private string getDefTagNameOfLambda( string SrcFile, int SrcLine )
{
    var parts = Utils1.getFileName( SrcFile).Split('.');
    var name = parts[0].ToLower();
    var ret = ""+name[0];

    for( var i = 1; i < name.Length; i++ )
    {
        var ch = name[i];
        if( !(ch == 'a' || ch == 'o' || ch == 'e' || ch == 'u' || ch == 'i'|| ch == 'y' ) )
            ret += name[i];

        if( ret.Length > 4 )
            break;
    }
    return $"{ret}{ convertTo1000(SrcLine) }";
}

private string _getQueryTag(string tag)
{
    if( tag.EndsWith("u8"))
        tag = tag.Substring(0, tag.Length - "u8".Length );

    if( tag.StartsWith("\"")
     && tag.EndsWith("\"") )
        tag = tag.Substring(1, tag.Length - 2*1 );

    return tag;
}

/// <summary>
/// generate fields in order to be filtered by AGGREGATE Group by.
/// </summary>
/// <param name="dict"></param>
/// <returns></returns>
private string generateMap( Dictionary<string,string>dict)
{
    if( dict == null)
        return "nil";

    var txt = @"map[string]string{
            ";
    foreach( var it in dict )
    {
        txt += $@" ""{it.Key}"":""{it.Value}"",
            ";
    }
    txt += "}";
    return txt;
}
internal string GoLang_ExportQuery( TLambdaCode lambda, Sql_ConfigTranslation optionsConfig )
{
    var queryVars = "";
    //var queryFields = "";
    

    var exportedDictFields = new Dictionary<string, string>();
    foreach( var field in lambda.Fields ) 
    {
        var fldSignature = $@"{PREFIX_FIELD}{field.Key}{POSTFIX_FIELD}";
        if( lambda.SqlCode.Contains(fldSignature) )
        {
            //queryFields += $@" ""{fldName}"",";
            var identifKey = field.Value.LangName;
            if( identifKey == "")
                identifKey = field.Value.SqlName;
            exportedDictFields[ identifKey ] = fldSignature;
        }
    }
    foreach( var externVar in lambda.ExternVar ) 
    {
        queryVars += GoLang_externVar( externVar.Value );
    }

    
    var querySql = lambda.SqlCode;
    var querySelectSqlFields = generateMap( lambda.Select_SqlFields);
    var queryFields = generateMap( exportedDictFields);
    var querySqlBase64 = Utils1.Base64Encode( querySql );
    //var queryFields = lambda.Fields;
    //var queryVars = lambda.ExternVar;

    var relativePath = Utils1.getRelativePath( optionsConfig.RootProject, lambda.SrcFile );

    relativePath = relativePath.Replace('/', Path.DirectorySeparatorChar);
    relativePath = relativePath.Replace(@"\", @"\\");

    var queryFile = relativePath;//lambda.SrcFile;
    var queryFileStartOffset = lambda.SrcStartOffset;
    var queryFileEndOffset = lambda.SrcEndOffset;
    var queryHash = lambda.Hash;

    var queryTag1 = _getQueryTag( lambda.Tag );
    var queryTag  = $"{queryTag1}-{lambda.SubTag}";

    var text = $@"
    ""{queryTag}"": 
	{{
        //compiledQuery: {querySql},
		CompiledQuery:`{querySqlBase64}`,
        
        SelectSqlFields: {querySelectSqlFields},
		
		Fields:			{queryFields},
		ExternVar:		[]TExternVar{{
								{queryVars.Trim()} 
							}},
        Tag: 			""{queryTag1}"",
		File: 			""{queryFile}"",
		StartOff: 		{queryFileStartOffset},
		EndOff:  		{queryFileEndOffset},
		Hash:  			""{queryHash}"",
	}},
        ";
        text = UseTemplate( text, this.Options.ConvertSql.Templ_GoSqlCompiledQuery, 
            new Dictionary<string,string>()
            { 
                { "queryTag", queryTag},
                { "queryTag1", queryTag1},
                { "querySql", querySql},
                { "querySqlBase64", querySqlBase64},
                { "querySelectSqlFields", querySelectSqlFields},
                
                { "queryFields", queryFields},
                { "queryVars", queryVars},
                { "queryFile", queryFile},
                { "queryFileStartOffset", ""+queryFileStartOffset},
                { "queryFileEndOffset", ""+queryFileEndOffset},
                { "queryHash", queryHash}
            }
            );

    return text;
}
internal static string GoLang_ExportSqlFile( Options options, string allQueries)
{
    var dialectName = options.ConvertSql.SqlDialect.GetDialectName();

    var text = $@"
    /* this class is generated automatically by compiler AtomicSql.exe exporter*/

    package newton_orm

    func (_this *DBContext) LoadCompiledQuery(){{

        _this.CompiledSqlQueries	= map[string]TCompiledSqlQuery{{
		    
        {allQueries}
        }}

        _this.DBContextBase.ProcessCompiledQuery( &_this.CompiledSqlQueries, false );
    }}
    
    func (_this *DBContext) CompiledQueries_GetDialectName() string {{ return ""{dialectName}""; }}
        ";
        text = UseTemplate( text, options.ConvertSql.Templ_GoSqlCompiledFile, 
            new Dictionary<string,string>()
            { 
                { "allQueries", allQueries},
                { "dialectName", dialectName },
            }
            );


    return text;
}

public static string UseTemplate(string original, ConvCommon.TemplateItem template, Dictionary<string, string>dict)
    { 
        if( template == null) {
            return original;
        }
        return template.ConvertTemplate( dict);
    }

}

