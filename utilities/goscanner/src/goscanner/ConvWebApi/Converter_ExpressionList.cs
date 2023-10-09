//******************************************************************************************************
//  ScannerBase_ExpressionList.cs - Gbtc
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
//  05/04/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using System.Collections.Generic;
using System.Diagnostics;
using goscanner.Metadata;

namespace goscanner.ConvWebApi;

public partial class ConvWebApi
{
    // Stack handlers:
    //  constSpec (optional)
    //  varSpec (optional)
    //  assignment (required)
    //  shortValDecl (required)
    //  returnStmt (optional)
    //  exprSwitchCase (optional)
    //  recvStmt (optional)
    //  rangeClause (optional)
    //  arguments (optional)
    protected readonly new  ParseTreeValues<ExpressionInfo[]> ExpressionLists = new();

    public override void ExitExpressionList(GoParser.ExpressionListContext context)
    {
        List<ExpressionInfo> expressions = new();

        for (int i = 0; i < context.expression().Length; i++)
        {
            if (Expressions.TryGetValue(context.expression(i), out ExpressionInfo expression))
                expressions.Add(expression);
        }

        ExpressionLists[context] = expressions.ToArray();
    }


    List< Dictionary<string, string> > DictElementList = new ();
    Dictionary<string, string> LastDictElement = null;
    public override void EnterElementList( GoParser.ElementListContext context) 
    { 
        //Debugger.Break();
        if(DictElementList.Count== 0)
            DictElementList.Add( new Dictionary<string, string>() );
        else
            DictElementList.Insert( 0, new Dictionary<string, string>() );
    }
    public override void ExitElementList( GoParser.ElementListContext context) 
    { 
        //Debugger.Break();
        LastDictElement = DictElementList[0];
        DictElementList.RemoveAt( 0 );
    }
	public override void ExitKeyedElement( GoParser.KeyedElementContext context) 
    {
        //Debugger.Break();
        
        if( context.key() != null && context.key().IDENTIFIER() != null && DictElementList.Count > 0) 
        {
            var dict = DictElementList[0];

            var fldName = context.key().IDENTIFIER().GetText();
            var fldValExprCtx = context.element().expression();
            if( Expressions.TryGetValue( fldValExprCtx,out ExpressionInfo expr) )
            {
                var fldVal = expr.Text;
                dict[ fldName ] =  fldVal;
            }
        }
    }


}
