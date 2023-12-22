using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime;
using goscanner.ConvCommon;
using goscanner.Metadata;

namespace goscanner.ConvSql
{
    internal class TBlock
    {
        internal TLambdaCode lambda;

        internal StringBuilder genCode;

        public TBlock( StringBuilder genCode, TLambdaCode lambda )
        {
            this.lambda = lambda;
            this.genCode = genCode;
        }   
    }
    public class  TField 
    {
        public string SqlName { get; set; }

        public string LangName { get; set; }
        public string TypeLang { get; set; }
        public bool IsFK;
        public int UID;
        private static int s_UID= 0;

        public TField(string sqlName, string langName, string type )
        {
            if( sqlName == "RoleNameDefault")
                Utils1.Nop();

            UID= s_UID++;
            if( UID == 47 || UID == 48)
                Utils1.Nop();

            SqlName = sqlName;
            TypeLang = type;
            LangName = langName;//Utils1.GetLastPart( identif, "." );
            this.IsFK = false;
        }
        public void setFK(bool bFK)
        {
            this.IsFK = bFK;    
        }
        public void addLangName(string langName )
        {
            this.LangName = this.LangName + "." + langName;
        }
            
    }

    public class SubQueryVar
    {
        public string Name { get; set; }

        //golang type name def
        public string Type { get; set; }   
    }

    public class SubQuery
    {
        public ParserRuleContext context;
        //public List<SubQueryVar> StaticsVarNames = new List<SubQueryVar>();
        public string VariableStorageName =""; // var ids = ctx.Users.where()..
        //public string GolangCode = "";   //ctx.Users.where(x=> UserName== userName).sqlQuery_GetCount();
        //public List<TLambdaCode> Queries = new List<TLambdaCode>(); //UserName== userName
        //public string UniqueID = "";
    }
    

    public class TLambdaCode
    {        
        public bool IsAllowArray = false;

        //ctx.Qry("tag name").Where()...
        public string Tag{ get; set; }

        //this lamda is after select. Select().Where(x=>lambda)  <- here is positioned
        public bool IsQueuedSelect{ get; set; }

        //public string PrevSubTag;
        string _SubTag;
        public string SubTag{ get{return _SubTag; }
            set{ _SubTag = value; 

                if( $"{Tag}-{_SubTag}" == "evcy59-S1")
                    Utils1.Nop();

                if( $"{Tag}-{_SubTag}" == "evcy59-W1")
                    Utils1.Nop();
                    
                } 
            }
        public ParserRuleContext TagCtx;
        public string lambdaVariable;
        public string SqlCode{ get; set; } = "";
        public Dictionary<string,string> Select_SqlFields = null;

        public Dictionary<string, TField> Fields = new ();

        public Dictionary<string, TField> ExternVar = new ();

        public string SrcFile;
        public int  SrcLine;
        public int  SrcCol;
        
        public long SrcQryOffset;
        public long SrcStartOffset;
        public long SrcEndOffset;
        public string Hash;
        public bool  isQryS=false;

        public TLambdaCode ParentLambda { get; set; }

        public bool  returnClause = false;
        public ParserRuleContext ctxTag;
        public int SubTagCounter;//save the Counter fo subTag

        public List<SubQuery> SubQueries { get; set; } = null;
        


        public TLambdaCode( 
            SqlConvert inst, SqlConvert.QueryTag lambdaQryTag, string subTag,
            TLambdaCode parentLambda, 
            //bool bIsQueuedSelect,
            ParserRuleContext ctxTag,             
            ParserRuleContext context,
            string lambdaVariable ) 
        {
            this.lambdaVariable = lambdaVariable;
            this.ctxTag = ctxTag;
            Tag    = lambdaQryTag.Tag;
            SubTag = subTag + lambdaQryTag.UpdateCounterSubTag();
            //if( Tag == "tsql082.\"ids\"u8, q")
            //    Utils1.Nop();

            this.IsQueuedSelect = SubTag == OrmDef.Func_Select 
                               || SubTag == OrmDef.Func_SelectSubQ;

            //var This = this;
            
            //lambdaSubTag.CurrentLambdaCode = this;
            if( SubTag.StartsWith( OrmDef.SubTag_SelectSubQ) 
             || SubTag.StartsWith( OrmDef.SubTag_WhereSubQ ) )
            {
                this.SubQueries = new List<SubQuery>();
            }

            TagCtx = ctxTag;
            ParentLambda = parentLambda;

            SrcFile = context.Start.InputStream.SourceName;
            SrcLine = context.Start.Line;
            SrcCol  = context.Start.Column;


            var dirRoot = inst.Options.ConvertSql.SourcePathDir;
            if( !dirRoot.EndsWith( Path.PathSeparator))
                dirRoot += Path.DirectorySeparatorChar;

            SrcFile = SrcFile.Replace( dirRoot, "" );

            
            SrcQryOffset    = lambdaQryTag.Context.Start.StartIndex;
            SrcStartOffset  = context.Start.StartIndex;
        }

        public void UpdateSubTag(string subTag)
        {
            this.SubTag = subTag;
        }
        public bool IsLambdaVariable(string lambdaVar)
        {
            return this.lambdaVariable == lambdaVar;
        }

        public void endSource(ParserRuleContext ctx)
        {
            SrcEndOffset = ctx.Stop.StartIndex;
                                                  
            var intv = new Antlr4.Runtime.Misc.Interval( (int)SrcStartOffset, (int)SrcEndOffset-1 ); 
            var srcData = ctx.Start.InputStream.GetText(intv);

            /*
            srcData = srcData.Replace( "\n", " ");
            srcData = srcData.Replace( "\r", " ");
            srcData = srcData.Replace( "\t", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Trim();

            Hash = Utils1.CreateMD5(srcData);
            */
            if( this.Tag == "tst1074")
                Utils1.Nop();


            var intv2 = new Antlr4.Runtime.Misc.Interval( (int)SrcQryOffset, (int)SrcStartOffset-1 ); 
            var callQry = ctx.Start.InputStream.GetText(intv2);

            var idx1 = callQry.IndexOf($".{OrmDef.Func_DBTable_Qry}(");
            if( idx1 >= 0 )
            {
                SrcQryOffset +=  idx1;
            }else
            {
                idx1 = callQry.IndexOf($".{OrmDef.Func_DBTable_QryS}(");
                if( idx1 >= 0)
                {
                    SrcQryOffset +=  idx1;
                    this.isQryS = true;
                }
            }

            SrcStartOffset -= SrcQryOffset;
            SrcEndOffset   -= SrcQryOffset;
            this.Hash = Utils1.Base64Encode(srcData);
            
        }

        public SubQuery getSubQueryByVarName(string variableName)
        {
            if( this.SubQueries != null ) 
            {
                foreach( var subQuery in this.SubQueries )
                { 
                    if( subQuery.VariableStorageName == variableName ) 
                        return subQuery;
                }
            }
            return null;
        }
    }





}
