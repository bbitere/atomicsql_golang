using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using Antlr4.Runtime;
using System.Diagnostics;
using goscanner.Metadata;
using Antlr4.Runtime.Misc;
using goscanner.ConvCommon;

namespace goscanner.ConvSql;

public partial class SqlConvert
{

    public class SubTag
    {
        public string      Tag;
        public TLambdaCode LambdaCode; 
        public SubTag(string tag)
        {
            Tag = tag;
            LambdaCode = null;
        }
    }
    List<SubTag> m_LambdaFunc_SubTags = new List<SubTag>();
    int m_LambdaFunc_SubTagCounter = 0;
    bool m_LambdaCode_IsQueuedSelect = false;
    static Dictionary<string, GoParser.PrimaryExprContext> dictLambdaTag = new ();
    //int m_LambdaFunc_Select = 0;

    internal List<TLambdaCode> m_exportedLambdas = new List<TLambdaCode>();

    internal TLambdaCode m_LambdaCode = null;
    internal String      m_LambdaTag = null;
    //internal String      m_LambdaSubTag = null;
    
    internal ParserRuleContext m_LambdaTagCtx = null;

    //------------------------------------------------------------------

    protected void Lambda_endSequence()
    {
        if( this.m_LambdaCode == null 
            && this.m_LambdaFunc_SubTags.Count == 0 )
        {
            //cleear the tag when the 
            m_LambdaTag = "";
            m_LambdaFunc_SubTagCounter = 0;
            m_LambdaCode_IsQueuedSelect = false;
            //m_LambdaFunc_Select = 0;
        }  
    }

    //when is calling Select() or Where()
    protected void Lambda_createSubtagID(string subTagName)
    {
        m_LambdaFunc_SubTags.Add( new SubTag($"{subTagName}{m_LambdaFunc_SubTagCounter}" )); //OrmDef.SubTag_Where
    }
    
    //cand se apeleaza Qry("tag1")
    protected void Lambda_callQryMethod(GoParser.PrimaryExprContext context, string argumentList)
    {
        m_LambdaTag = _getQueryTag(argumentList);
        if( context.Start.Line ==65 
         || m_LambdaTag == null
         || context.Start.TokenSource.SourceName.Contains("test1_rtm"))
        {
            Utils.Nop();
        }
        m_LambdaTagCtx = context;
        m_LambdaFunc_SubTagCounter = 0;

        if( m_LambdaTag != "" )
        {
            if( dictLambdaTag.ContainsKey( m_LambdaTag ))
            {
                if( dictLambdaTag[ m_LambdaTag] != context)
                {
                    this.Log_Error( context, $"{OrmDef.Func_DBTable_Qry}() method should receive a valid tag unique per app. See: {OrmDef.Func_DBTable_Qry}(\"{m_LambdaTag}\") ");
                }
            }
            dictLambdaTag[ m_LambdaTag] = context;
        }
    }

    //cand se apeleaza Where( func(x *Model) bool {...})
    protected void Lambda_callWhereMethod(ParserRuleContext context, string subTagName)        
    {
        var ttagLamda = m_LambdaFunc_SubTags[ m_LambdaFunc_SubTags.Count -1 ];
        if( ttagLamda.LambdaCode != null )
        {   //din cauza ca in apelul din orm, incrementarea counterului se face la apelul functiei, nu la EnterCall()
            //tre sa reupdatam corect
            m_LambdaFunc_SubTagCounter++;
            ttagLamda.LambdaCode.UpdateSubTag( $"{subTagName}{m_LambdaFunc_SubTagCounter}" );
        }
        m_LambdaFunc_SubTags.RemoveAt( m_LambdaFunc_SubTags.Count -1);

        if( m_LambdaCode != null ) 
        {   //put back the old lambda
            m_LambdaCode = m_LambdaCode.ParentLambda;
        }
    }
    //cand se apeleaza Select( ctx.Qry("").Where( func(x *Model) bool {...})
    // dar el se apeleaza la ExitRule(), deci vine dupa Where()
    protected void Lambda_callSelectMethod( ParserRuleContext context  )
    {
        this.Lambda_callWhereMethod( context, OrmDef.SubTag_Select );
        m_LambdaCode_IsQueuedSelect = true;
    }

    protected void Lambda_Return(ParserRuleContext context, ExpressionInfo[] expressions)
    {
        if( m_LambdaCode != null )
        {
            if (expressions.Length != 1)
                getTextSQLError("return suport only 1 expression", context);

            m_LambdaCode.SqlCode = expressions[0].SQLText;
            if( m_LambdaCode.SqlCode == null)
                Debugger.Break();
        }
    }

    protected string Lambda_getSQLIdentif(string identifExpr, string sqlIdentif, string sqlField, TypeInfo type, string fldName, ParserRuleContext ctx)
    {
        if( m_LambdaCode == null )
            return null;
        
        var origIdentif = sqlIdentif;
        sqlIdentif = sqlIdentif.Replace( PREFIX_FIELD, "" );
        sqlIdentif = sqlIdentif.Replace( POSTFIX_FIELD, "" );

        if( type.Name == "var")
        {
            return getTextSQLError( $"internal error 187. not found {identifExpr}", ctx);
        }

        if( sqlIdentif == "x")
            Utils.Nop();

        var genType = type.getNormalizedType();
        if( genType == null )
        {
            return getTextSQLError( $"internal error 194. Not found {identifExpr}", ctx);
        }
        if( sqlField == "" )
        {
            //m_LambdaCode.Fields[identif] = new TField(field, genType.Name );
            // I dont want to add it because it is the table
            //return $"{PREFIX_FIELD}{identif}{POSTFIX_FIELD}";
            return $"{PREFIX_FIELD}{POSTFIX_FIELD}"; //the engine add the table (instance this ) automatically
        }
        if( sqlIdentif == "" )
        {
            m_LambdaCode.Fields[$"{sqlField}"] = new TField(sqlField, fldName, genType.Name );
            return $"{PREFIX_FIELD}{sqlField}{POSTFIX_FIELD}";
        }
        else
        {
            if( sqlIdentif.Contains( PREFIX_VAR ) 
                && sqlIdentif.Contains( POSTFIX_VAR ) )
            { 
                Log_Error( ctx, $"Not support,yet, external variable having type a struct inside Lambda function, {identifExpr??""}" );
            }
            if( !m_LambdaCode.Fields.ContainsKey(sqlIdentif))
            {
                m_LambdaCode.Fields[sqlIdentif] = new TField(sqlField, fldName, genType.Name );
            }else
            {
                m_LambdaCode.Fields[sqlIdentif].setFK(true);
                m_LambdaCode.Fields[sqlIdentif].addLangName(fldName);
            }
            m_LambdaCode.Fields[$"{sqlIdentif}.{sqlField}"] = new TField(sqlField, fldName, genType.Name );
            return $"{PREFIX_FIELD}{sqlIdentif}.{sqlField}{POSTFIX_FIELD}";
        }
    }
    protected string Lambda_getSQLVarIdentif(string identif, TypeInfo type, ParserRuleContext ctx)
    {
        if( m_LambdaCode == null )
            return null;
        
        identif = identif.Replace( PREFIX_VAR, "" );
        identif = identif.Replace( POSTFIX_VAR, "" );

        if( type.Name == "var")
        {
            return getTextSQLError( $"internal error 202. Not found {identif}", ctx);
        }
        if( identif == "x")
            Utils.Nop();

        var genType = type.getNormalizedType();
        if( genType != null )
        {
            m_LambdaCode.ExternVar[identif] = new TField(identif, identif, genType.Name );
            return $"{PREFIX_VAR}{identif}{POSTFIX_VAR}";
        }else
        {
            return getTextSQLError("internal error 245", ctx);
        }
    }
    void Lambda_enterFuncLit(ParserRuleContext context)
    {
        if( m_LambdaFunc_SubTags.Count > 0 )
        {
            m_LambdaCode = new TLambdaCode( this, m_LambdaTag, 
                                m_LambdaFunc_SubTags[m_LambdaFunc_SubTags.Count-1], 
                                m_LambdaCode, m_LambdaCode_IsQueuedSelect,
                                m_LambdaTagCtx, context);
            AddLambda( m_LambdaCode );
        }
    }

    // var v = atm.Select( ctx.table.Qry() ...,
    //          func(x *T) *V{
    //              return &V{ field1: val,}
    //          } )
    // se reitereaza fieldurile de inistalizare a lui V din func lit
    protected string Lambda_SelectFields( ParserRuleContext context, 
                GoParser.KeyedElementContext[] keyedElements, TypeInfo typeInfo1 )
    {
        var arrSqlSelectFld = new List<string>();
        var bError = false;
        var textSqlSelect = "";

        m_LambdaCode.Select_SqlFields = new Dictionary<string, string>();

        foreach (GoParser.KeyedElementContext keyedElement in keyedElements)
        {
            var key         = keyedElement.key().IDENTIFIER().Symbol.Text;
            var expression  = keyedElement.element().expression();

            if( !Utils1.isPublicIdentGoLang( key ))
            {
                Log_Error( context, $"The Name of field should start with capital in order to be public: {key}" );
            }

            var (isValid, sqlItem) = Select_generateItem( m_LambdaCode.Select_SqlFields,
                                        key, expression, context, typeInfo1);
            if( isValid ) 
            {
                arrSqlSelectFld.Add(sqlItem);   
            }else
            {
                bError = true;
                textSqlSelect = sqlItem;
                break;
            }
        }
        if( !bError)
        {
            textSqlSelect = $"{string.Join(", ", arrSqlSelectFld )}";
        }
        return textSqlSelect;
    }
}

