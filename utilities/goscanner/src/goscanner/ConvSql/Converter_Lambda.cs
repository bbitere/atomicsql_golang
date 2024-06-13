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
using static GoParser;

namespace goscanner.ConvSql;

public partial class SqlConvert
{

    public class QueryTag
    {
        public string       Tag;
        //public string       Subtag;
        public int          SubTagCounter;
        public TLambdaCode LambdaCodeContainer; 
        public ParserRuleContext Context;
        public QueryTag(string tag, TLambdaCode lambdaCode, ParserRuleContext ctx)
        {
            Tag    = tag;
            //Subtag = subtag;
            LambdaCodeContainer = lambdaCode;
            Context = ctx;
        }
        public int UpdateCounterSubTag()
        {
            ++SubTagCounter;
            return SubTagCounter;
            //return $"{Subtag}{SubTagCounter}";
        }
    }
    //ctx.Table1.Qry("tag1").Where(func(x Table1)bool{ ... code1 }).Where(func(x Table1)bool{ ... code2 })
    //a subTag is generated at Qry("tag1") and continue until end the chain.
    //
    // if is entering inside a Where() it should reset and also put back after exit
    List<QueryTag> m_LambdaFunc_SubTags = new List<QueryTag>();
    //int m_LambdaFunc_SubTagCounter = 0;
    bool m_LambdaCode_IsQueuedSelect = false;
    static Dictionary<string, GoParser.PrimaryExprContext> dictLambdaTag = new ();
    //int m_LambdaFunc_Select = 0;

    internal List<TLambdaCode> m_exportedLambdas = new List<TLambdaCode>();

    internal List<SubQuery>    m_listLambdaSubqueries = new List<SubQuery>();
    internal TLambdaCode m_LambdaCode = null;
    //internal String      m_LambdaTag = null;
    //internal String      m_LambdaSubTag = null;
    
    //internal ParserRuleContext m_LambdaTagCtx = null;

    //------------------------------------------------------------------

    /*
    protected string GenerateSqlSubquery( TLambdaCode lambdaCode, string varName )
    {
        if( lambdaCode != null && lambdaCode.SubQueries != null )
        {
            foreach( var subquery in lambdaCode.SubQueries)
            {
                if( subquery.VariableStorageName == varName )
                { 
                    /*
                    var arguments = new List<string>();
                    foreach(var arg in subquery.StaticsVarNames)
                    {
                        arguments.Add( arg.Name );
                    }
                    var txtArguments = string.Join( ",", arguments.ToArray() );
                    * /
                    return $" {START_SUBQUERY} {varName} {END_SUBQUERY} ";
                }
            }
        }
        return null;
    }*/
    protected SubQuery AddSubquery(
        TLambdaCode lambdaCode, ParserRuleContext context,
        string varAssign, string methName, string golangText )
    {
        var q = new SubQuery();
        q.VariableStorageName = varAssign;
        /*
        golangText = golangText.Trim();
        golangText = golangText.Replace( $"{OrmDef.Func_DBTable_Qry}(\"\")",
                                         $"{OrmDef.Func_DBTable_Qry}(__tagQuery)");
        golangText = golangText.Replace( $"{methName}(",  $"Sqlquery_{methName}(");
        golangText = golangText.Replace( $"{methName} (", $"Sqlquery_{methName}(");
        q.GolangCode = golangText;
        q.context    = context;

        ++SqlConvert.s_SubQueries_uniqueID;
        //q.UniqueID = $"@_SUBQUERY_{SqlConvert.s_SubQueries_uniqueID}";
        */
        if( lambdaCode.SubQueries != null )
        {
            lambdaCode.SubQueries.Add( q );
            m_listLambdaSubqueries.Add( q );
        }
        return q;
    }
    protected SubQuery PopSubquery(ParserRuleContext context, string varNameSubQuery )
    {
        if( m_listLambdaSubqueries.Count > 0 )
        {
            var top =  m_listLambdaSubqueries[ m_listLambdaSubqueries.Count-1];
            if( top.VariableStorageName == varNameSubQuery ) 
            {
                m_listLambdaSubqueries.RemoveAt( m_listLambdaSubqueries.Count-1 );
                return top;
            }
        }
        return null;
    }
    protected SubQuery GetTopSubquery()
    {
        if( m_listLambdaSubqueries.Count > 0 )
        {
            var top =  m_listLambdaSubqueries[ m_listLambdaSubqueries.Count-1];
            return top;
        }
        return null;
    }
    protected void Lambda_endChainOfQuery()
    {
        if( this.m_LambdaCode == null 
            && this.m_LambdaFunc_SubTags.Count == 0 )
        {
            //cleear the tag when the 
            //m_LambdaTag = "";
            //m_LambdaFunc_SubTagCounter = 0;
            m_LambdaCode_IsQueuedSelect = false;
            //m_LambdaFunc_Select = 0;
        } 
        this.Lambda_popQueryTag(null);
    }

    //end where chain of query end
    protected QueryTag Lambda_popQueryTag(TLambdaCode lamdbaCode)
    {
        if(m_LambdaFunc_SubTags.Count == 0)
            Debugger.Break();

        var top = Lambda_getTopQueryTag();
        if( top != null )
        {
            if( lamdbaCode == null )
            {
                m_LambdaFunc_SubTags.RemoveAt( m_LambdaFunc_SubTags.Count -1);
            }else
            if(top.LambdaCodeContainer != null && top.LambdaCodeContainer != lamdbaCode)
            {
                m_LambdaFunc_SubTags.RemoveAt( m_LambdaFunc_SubTags.Count -1);
            }
        }
        return top;
    }
    protected QueryTag Lambda_getTopQueryTag()
    {
        if( m_LambdaFunc_SubTags.Count == 0 )
            return null;
        return m_LambdaFunc_SubTags[ m_LambdaFunc_SubTags.Count -1 ];
    }

    
    protected void Lambda_createQueryTag_MethodQry(string Tag, ParserRuleContext ctx)
    {
        if( Tag == "tst1074")
            Utils.Nop();

        var top = Lambda_getTopQueryTag();
        if( top != null && top.LambdaCodeContainer == m_LambdaCode)
        {
            Log_Error(top.Context, $"Last call of {OrmDef.Func_DBTable_Qry}('{top.Tag}') is not ended in a right manner!");
            //Log_Error(ctx, $"Last call of {OrmDef.Func_DBTable_Qry}('{top.Tag}') is not ended in a right manner!");
        }
        m_LambdaFunc_SubTags.Add( new QueryTag(Tag, m_LambdaCode, ctx ) );
        
    }
    
    
    //cand se apeleaza Qry("tag1")
    protected void Lambda_callQryMethod(GoParser.PrimaryExprContext context, string argument0, string parentLambdaTag="")
    {
        var lambdaTag = parentLambdaTag + _getQueryTag(argument0);
        Lambda_createQueryTag_MethodQry( lambdaTag, context);
        
        //if( m_LambdaTag == "tsql082.\"ids\"u8, q")
        //      Utils1.Nop();

        if( context.Start.Line ==65 
         || lambdaTag == null
         || context.Start.TokenSource.SourceName.Contains("test1_rtm"))
        {
            Utils.Nop();
        }
        //Lambda_resetSubTag( null, context);
        //m_LambdaTagCtx = context;
        //m_LambdaFunc_SubTagCounter = 0;

        if( lambdaTag != "" )
        {
            if( dictLambdaTag.ContainsKey( lambdaTag ))
            {
                if( dictLambdaTag[ lambdaTag] != context)
                {
                    this.Log_Error( context, $"{OrmDef.Func_DBTable_Qry}() method should receive a valid tag unique per app. See: {OrmDef.Func_DBTable_Qry}(\"{lambdaTag}\") ");
                }
            }
            //Debug_Console($"Func_DBTable_Qry1: {m_LambdaTag}");
            dictLambdaTag[ lambdaTag] = context;
        }
    }

    
    //cand se apeleaza Where( func(x *Model) bool {...})
    protected void Lambda_callWhereMethod(ParserRuleContext context, string subTagName)        
    {
        /*
        var ttagLamda = Lamda_getTopSubTag();
        if( ttagLamda.CurrentLambdaCode != null )
        {   //din cauza ca in apelul din orm, incrementarea counterului se face la apelul functiei, nu la EnterCall()
            //tre sa reupdatam corect
            m_LambdaFunc_SubTagCounter = ttagLamda.CurrentLambdaCode.SubTagCounter;
            m_LambdaFunc_SubTagCounter++;
            ttagLamda.CurrentLambdaCode.UpdateSubTag( $"{subTagName}{m_LambdaFunc_SubTagCounter}" );

            if( m_LambdaCode !=  null && ttagLamda.CurrentLambdaCode != m_LambdaCode) 
                Utils.Nop();
        }

        if( m_LambdaCode != null ) 
        {   //put back the old lambda
            //setLambdaCode( m_LambdaCode.ParentLambda );
            //if( m_LambdaCode != null )
            //    Lambda_resetSubTag( m_LambdaCode, context);
        }*/
    }
    //cand se apeleaza Select( ctx.Qry("").Where( func(x *Model) bool {...})
    // dar el se apeleaza la ExitRule(), deci vine dupa Where()
    protected void Lambda_callSelectMethod( ParserRuleContext context, string SubTag_Select  )
    {
        this.Lambda_callWhereMethod( context, SubTag_Select );
        //m_LambdaCode_IsQueuedSelect = true;
    }

    protected void Lambda_Return(ParserRuleContext context, ExpressionInfo[] expressions)
    {
        if( m_LambdaCode != null )
        {
            if (expressions.Length != 1)
                getTextSQLError("return suport only 1 expression", context);

            if(m_LambdaCode.SubQueries != null)
                Utils.Nop();

            m_LambdaCode.SqlCode   = expressions[0].SQLText;
            if( m_LambdaCode.SqlCode == null)
                Debugger.Break();

            m_LambdaCode.NoSqlCode = "";
            if( expressions[0].NoSQLCode != null)
            {
                var noSQLCode = expressions[0].NoSQLCode;
                if( expressions[0].Type.IsBool() )
                {
                    noSQLCode = NSqlConvertToBool( noSQLCode, expressions[0].OperandKind );
                }
                m_LambdaCode.NoSqlCode =  noSQLCode.getNoSqlCode(0);
            }
            
        }
    }

    protected string Lambda_getSQLIdentif(string identifExpr, string sqlIdentif, string sqlField, TypeInfo type, string fldName, ParserRuleContext ctx, string operandDOT)
    {
        if( m_LambdaCode == null )
            return null;

        if( m_LambdaCode.Tag == "tsql082a")
            Utils.Nop();
        if( identifExpr == "x.UserRoleID.RoleStatusID.ID")
            Utils.Nop();


        var subQuery = this.GetTopSubquery();
        if( m_LambdaCode.ParentLambda == null)
        {
            if( subQuery != null)
            {
                // nothing
                return $"{sqlField}%%%";
            }            
        }else
        {
            if( subQuery != null)
                Utils.Nop();
        }
        
        var origIdentif = sqlIdentif;
        sqlIdentif = sqlIdentif.Replace( PREFIX_FIELD, "" );
        sqlIdentif = sqlIdentif.Replace( POSTFIX_FIELD, "" );

        sqlIdentif = sqlIdentif.Replace( SUBQ_PREFIX_FIELD, "" );
        sqlIdentif = sqlIdentif.Replace( SUBQ_POSTFIX_FIELD, "" );

        

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
        
        /*
        var subQuerySql  = GenerateSqlSubquery( m_LambdaCode, fldName);
        if( subQuerySql != null)
            return subQuerySql;
        */
        if( sqlField == "" )
        {
            //m_LambdaCode.Fields[identif] = new TField(field, genType.Name );
            // I dont want to add it because it is the table
            //return $"{PREFIX_FIELD}{identif}{POSTFIX_FIELD}";
            return $"{PREFIX_FIELD}{POSTFIX_FIELD}"; //the engine add the table (instance this ) automatically
        }

        if( subQuery != null)
        if( sqlField == "userRole_ID")
            Utils.Nop();
        
        var lambdaName = operandDOT.Split(".")[0];

        if( sqlIdentif == "" )
        {
            var parentLambda = m_LambdaCode.ParentLambda;
            if( parentLambda != null
                && (  parentLambda.Fields.ContainsKey( sqlField )
                   || parentLambda.IsLambdaVariable( lambdaName ))
                )
            {
                //Here the subquery has an Where() or Select() and the code using a parent field
                // ctx.Tabl1.Qry("label1").WhereSubQ( func(x Table1, q IDBQuery){
                // ids, _ := ctx.Table2.QryS("ids",q).Where( func(y Table2) { reutrn y.ID == x.FK_ID })
                // })
                if( !parentLambda.Fields.ContainsKey(sqlField))
                {
                    parentLambda.Fields[$"{sqlField}"] = new TField(sqlField, fldName, genType.Name );
                }

                return $"{SUBQ_PREFIX_FIELD}{sqlField}{SUBQ_POSTFIX_FIELD}";
            }else
            {
                m_LambdaCode.Fields[$"{sqlField}"] = new TField(sqlField, fldName, genType.Name );
                return $"{PREFIX_FIELD}{sqlField}{POSTFIX_FIELD}";
            }
        }
        else
        {
            if( sqlIdentif.Contains( PREFIX_VAR ) 
             && sqlIdentif.Contains( POSTFIX_VAR ) )
            { 
                if( m_LambdaCode != null && m_LambdaCode.SubQueries == null )
                    Log_Error( ctx, $"Not support,yet, external variable having type a struct inside Lambda function, {identifExpr??""}" );
            }

            var parentLambda = m_LambdaCode.ParentLambda;
            if( parentLambda != null
                && (  parentLambda.Fields.ContainsKey( sqlField )
                   || parentLambda.IsLambdaVariable( lambdaName ))
                )
            {
                //Here the subquery has an Where() or Select() and the code using a parent field
                // ctx.Tabl1.Qry("label1").WhereSubQ( func(x Table1, q IDBQuery){
                // ids, _ := ctx.Table2.QryS("ids",q).Where( func(y Table2) { reutrn y.ID == x.FK_ID })
                // })

                var langName = fldName;
                if( !parentLambda.Fields.ContainsKey(sqlIdentif))
                {
                    parentLambda.Fields[sqlIdentif] = new TField(sqlField, fldName, genType.Name );
                }else
                {
                    parentLambda.Fields[sqlIdentif].setFK(true);
                    langName = parentLambda.Fields[sqlIdentif].addLangName(fldName);
                }
                parentLambda.Fields[$"{sqlIdentif}.{sqlField}"] = new TField(sqlField, langName, genType.Name );
                return $"{SUBQ_PREFIX_FIELD}{sqlIdentif}.{sqlField}{SUBQ_POSTFIX_FIELD}";
            }else
            {
                var langName = fldName;
                if( !m_LambdaCode.Fields.ContainsKey(sqlIdentif))
                {
                    m_LambdaCode.Fields[sqlIdentif] = new TField(sqlField, fldName, genType.Name );
                }else
                {
                    m_LambdaCode.Fields[sqlIdentif].setFK(true);
                    langName = m_LambdaCode.Fields[sqlIdentif].addLangName(fldName);

                }
                m_LambdaCode.Fields[$"{sqlIdentif}.{sqlField}"] = new TField(sqlField, langName, genType.Name );
                return $"{PREFIX_FIELD}{sqlIdentif}.{sqlField}{POSTFIX_FIELD}";
            }
        }
    }
    protected string Lambda_getSQLVarIdentif(string identif, TypeInfo type, ParserRuleContext ctx)
    {
        if( m_LambdaCode == null )
            return null;
        
        identif = identif.Replace( PREFIX_VAR, "" );
        identif = identif.Replace( POSTFIX_VAR, "" );

        if( m_LambdaCode.getSubQueryByVarName(identif) != null )
        {
            return $"{START_SUBQUERY}{identif}{END_SUBQUERY}";
        }

        if( type.Name == "var")
        {
            return getTextSQLError( $"internal error 202. Not found {identif}", ctx);
        }
        if( identif == "x")
            Utils.Nop();

        var genType = type.getNormalizedType();
        if( genType != null )
        {
            if( genType.Name == OrmDef.Class_IDBQuery)
            {
                return $"/*{identif} - noused */";
            }else
            {
                if(identif == "ids")
                    Utils.Nop();
                m_LambdaCode.ExternVar[identif] = new TField(identif, identif, genType.Name );
                return $"{PREFIX_VAR}{identif}{POSTFIX_VAR}";
            }
        }else
        {
            return getTextSQLError("internal error 245", ctx);
        }
    }
    GoParser.PrimaryExprContext getPrimaryContext(ParserRuleContext context)
    {
        var ctxPrimaryExpr = _getArgumentContext(context);
        if( ctxPrimaryExpr != null ) 
        {
            var ctx0 = ctxPrimaryExpr.Parent as GoParser.PrimaryExprContext;
            if( ctx0 != null ) 
                return ctx0;
        }
        return null;
    }
    GoParser.ArgumentsContext _getArgumentContext(ParserRuleContext context)
    {
        for( var ctx = context.Parent; ctx != null; ctx = ctx.Parent )
        { 
            var ctx0 = ctx as GoParser.ArgumentsContext;
            if( ctx0 != null ) 
                return ctx0;
        }

        return null;
    }
    void Lambda_enterFuncLit(GoParser.FunctionLitContext context)
    {
        var topQryTag = Lambda_getTopQueryTag();
        if( topQryTag != null )
        {
            var paramNameCtx = context.signature()?.parameters()?.parameterDecl(0)?.identifierList()?.IDENTIFIER(0);
            var paramName   = paramNameCtx != null? paramNameCtx.Symbol.Text : "";
            //Debug_Console($"Lambda_enterFuncLit: {m_LambdaTag} -> {topSubTag.SubTag}");

            var ctxPrimaryExpr = getPrimaryContext(context);
            if( ctxPrimaryExpr != null )
            {
                var subTagName  = OrmDef.GetSubTabByFuncName( ctxPrimaryExpr.m_funcMethodName, ctxPrimaryExpr.m_bIsNoSql );
                if( subTagName != null )
                {
                    if( subTagName != "" )
                    {
                        setLambdaCode( new TLambdaCode( this, topQryTag, subTagName, m_LambdaCode, 
                                                        ctxPrimaryExpr, context, paramName, ctxPrimaryExpr.m_bIsNoSql) );
                        AddLambda( m_LambdaCode );
                    }else
                    {
                        Log_Error(context, "It is Not Allowed that lit func to be exposed not in a call method or Where() of Select() ");
                    }
                }
            }else
            { 
                Log_Error(context, "It is Not Allowed that lit func to be exposed not in a call method of Where() or Select() ");
            }
            
        }
    }

    // var v = atm.Select( ctx.table.Qry() ...,
    //          func(x *T) *V{
    //              return &V{ field1: val,}
    //          } )
    // se reitereaza fieldurile de inistalizare a lui V din func lit
    protected (string, TNoSqlCode) Lambda_SelectFields( ParserRuleContext context, 
                GoParser.KeyedElementContext[] keyedElements, TypeInfo typeInfo1 )
    {
        var arrSqlSelectFld = new List<string>();
        var arrNoSqlSelectFld = new List<TNoSqlCode>();
        
        var bError = false;
        var textSqlSelect = "";
        TNoSqlCode noSqlSelect = null;

        m_LambdaCode.Select_SqlFields = new Dictionary<string, string>();

        foreach (GoParser.KeyedElementContext keyedElement in keyedElements)
        {
            var key         = keyedElement.key().IDENTIFIER().Symbol.Text;
            var expression  = keyedElement.element().expression();

            if( !Utils1.isPublicIdentGoLang( key ))
            {
                Log_Error( context, $"The Name of field should start with capital in order to be public: {key}" );
            }

            var (isValid, sqlItem, noSqlCode) = Select_generateItem( m_LambdaCode.Select_SqlFields,
                                        key, expression, context, typeInfo1);
            if( isValid ) 
            {
                arrSqlSelectFld.Add(sqlItem);   
                arrNoSqlSelectFld.Add(noSqlCode);
            }else
            {
                bError = true;
                textSqlSelect = sqlItem;
                noSqlSelect = noSqlCode;
                break;
            }
        }
        if( !bError)
        {
            textSqlSelect = $"{string.Join(", ", arrSqlSelectFld )}";
            noSqlSelect = new TNoSqlCode("select", arrNoSqlSelectFld);
        }
        return (textSqlSelect, noSqlSelect);
    }
}

