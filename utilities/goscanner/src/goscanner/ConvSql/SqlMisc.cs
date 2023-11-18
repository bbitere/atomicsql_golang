using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime;
using goscanner.ConvCommon;

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

        public TField(string sqlName, string langName, string type )
        {
            if( sqlName == "RoleNameDefault")
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
    public class TLambdaCode
    {        
        public bool IsAllowArray = false;

        //ctx.Qry("tag name").Where()...
        public string Tag{ get; set; }

        //this lamda is after select. Select().Where(x=>lambda)  <- here is positioned
        public bool IsQueuedSelect{ get; set; }

        //public string PrevSubTag;
        public string SubTag{ get; set; }
        public ParserRuleContext TagCtx;
        public string SqlCode{ get; set; } = "";
        public Dictionary<string,string> Select_SqlFields = null;

        public Dictionary<string, TField> Fields = new ();

        public Dictionary<string, TField> ExternVar = new ();

        public string SrcFile;
        public int  SrcLine;
        public int  SrcCol;
        
        public long SrcStartOffset;
        public long SrcEndOffset;
        public string Hash;

        public TLambdaCode ParentLambda { get; set; }

        public bool  returnClause = false;
        public ParserRuleContext ctxTag;

        public TLambdaCode( 
            SqlConvert inst, string lambdaTag, SqlConvert.SubTag lambdaSubTag, 
            TLambdaCode parentLambda, 
            bool bIsQueuedSelect,
            ParserRuleContext ctxTag,             
            ParserRuleContext context) 
        {
            this.ctxTag = ctxTag;
            Tag = lambdaTag;
            this.IsQueuedSelect = bIsQueuedSelect;

            var This = this;
            SubTag = lambdaSubTag.Tag;
            lambdaSubTag.LambdaCode = this;

            TagCtx = ctxTag;
            ParentLambda = parentLambda;

            SrcFile = context.Start.InputStream.SourceName;
            SrcLine = context.Start.Line;
            SrcCol  = context.Start.Column;

            if(lambdaTag == null && SrcLine == 66)
            {
                Utils1.Nop();
            }

            var dirRoot = inst.Options.ConvertSql.SourcePathDir;
            if( !dirRoot.EndsWith( Path.PathSeparator))
                dirRoot += Path.DirectorySeparatorChar;

            SrcFile = SrcFile.Replace( dirRoot, "" );
            SrcStartOffset = context.Start.StartIndex;
        }

        public void UpdateSubTag(string subTag)
        {
            this.SubTag = subTag;
        }

        public void endSource(ParserRuleContext ctx)
        {
            SrcEndOffset = ctx.Stop.StartIndex;
                                                  
            var intv = new Antlr4.Runtime.Misc.Interval( (int)SrcStartOffset, (int)SrcEndOffset ); 
            var srcData = ctx.Start.InputStream.GetText(intv);
            srcData = srcData.Replace( "\n", " ");
            srcData = srcData.Replace( "\r", " ");
            srcData = srcData.Replace( "\t", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Replace( "  ", " ");
            srcData = srcData.Trim();

            Hash = Utils1.CreateMD5(srcData);
        }
    }





}
