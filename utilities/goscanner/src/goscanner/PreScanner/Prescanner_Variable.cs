using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;
using goscanner.Metadata;

namespace goscanner
{
    public partial class PreScanner
    {
        List< Dictionary<string, string> > DictElementList = new ();
        Dictionary<string, string> LastDictElement = null;
        public override void EnterElementList( GoParser.ElementListContext context) 
        { 
            if( !InFunction)
            {
                //Debugger.Break();
                if( DictElementList.Count == 0)
                    DictElementList.Add( new Dictionary<string, string>() );
                else
                    DictElementList.Insert( 0, new Dictionary<string, string>() );
            }
        }
        public override void ExitElementList( GoParser.ElementListContext context) 
        { 
            if( !InFunction)
            {
                //Debugger.Break();
                LastDictElement = DictElementList[0];
                DictElementList.RemoveAt( 0 );
            }
        }
	    public override void ExitKeyedElement( GoParser.KeyedElementContext context) 
        {
            //Debugger.Break();
            if( !InFunction)
            {
                if( context.key()?.IDENTIFIER() != null && DictElementList.Count > 0) 
                {
                    var dict = DictElementList[0];

                    var fldName = context.key().IDENTIFIER().Symbol.Text;
                    var fldValLit = context.element().expression()?
                                        .primaryExpr()?.operand()?.literal()?.basicLit();

                    var str = "";
                    if(fldValLit != null)
                    {
                        if( fldValLit.string_() != null )
                        {
                            if( fldValLit.string_().RAW_STRING_LIT() != null)
                                str = fldValLit.string_().RAW_STRING_LIT().Symbol.Text;

                            if( fldValLit.string_().INTERPRETED_STRING_LIT() != null )
                                str = fldValLit.string_().INTERPRETED_STRING_LIT()?.Symbol.Text;
                        }
                        if( fldValLit.integer() != null )
                        {
                            if( fldValLit.integer().HEX_LIT() != null )
                                str = fldValLit.integer().HEX_LIT().Symbol.Text;

                            if( fldValLit.integer().RUNE_LIT() != null )
                                str = fldValLit.integer().RUNE_LIT().Symbol.Text;
                        }
                    }

                    if( fldName != null && str != null )
                    {
                        dict[ fldName ] =  str;
                    }
                }
            }
        }

        /*

        public override void ExitVarDecl(GoParser.VarDeclContext context)
            {

                var varspec = context.varSpec();
                if( varspec.Length == 1 )
                { 
                    var identif =  varspec[0].identifierList().IDENTIFIER();
                    if( identif.Length == 1 ) 
                    {
                        Types.TryGetValue( varspec[0].type_(), out TypeInfo typeInfo );

                        var expression = varspec[0].expressionList().expression()[0];
                        Expressions.TryGetValue( expression, out ExpressionInfo exprInfo );

                        if( exprInfo != null ) 
                        {
                            var listexpr = new List<string>();
                            if( exprInfo.SubExpressions != null )
                            foreach( var expr in exprInfo.SubExpressions)
                            {
                                listexpr.Add( expr );
                            }
                            var varInfo = new VariableInfo
                            {
                                Name = identif[0].GetText(),
                                Type = typeInfo,
                                InitExpr = listexpr.ToArray(),
                            };

                            m_globalVariables.Add( varInfo.Name, varInfo);
                        }
                    }
                }
            }
        */
        
    }
}
