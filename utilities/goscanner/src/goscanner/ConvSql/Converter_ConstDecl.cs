//******************************************************************************************************
//  Converter_ConstDecl.cs - Gbtc
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

using Antlr4.Runtime.Misc;
using goscanner.Metadata;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using static goscanner.Common;

namespace goscanner.ConvSql;

public partial class SqlConvert
{
    private int m_iota;
    private int m_constIdentifierCount;
    private bool m_constMultipleDeclaration;

    public override void EnterConstDecl(GoParser.ConstDeclContext context)
    {
        // constDecl
        //     : 'const' ( constSpec | '(' ( constSpec eos )* ')' )

        m_constIdentifierCount = 0;
        m_constMultipleDeclaration = context.children.Count > 2;
        m_iota = 0;
    }

    public override void ExitConstDecl(GoParser.ConstDeclContext context)
    {
        // constDecl
        //     : 'const' ( constSpec | '(' ( constSpec eos )* ')' )

        if (EMITTING_CODE)
        {
            if (m_constMultipleDeclaration && EndsWithLineFeed(m_targetOutputFile.ToString()))
            {
                string removedLineFeed = RemoveLastLineFeed(m_targetOutputFile.ToString());
                m_targetOutputFile.Clear();
                m_targetOutputFile.Append(removedLineFeed);
            }

            m_targetOutputFile.Append(CheckForCommentsRight(context));
        }
    }

    public override void ExitConstSpec(GoParser.ConstSpecContext context)
    {
        // constSpec
        //     : identifierList ( type ? '=' expressionList ) ?

        if (EMITTING_CODE)
        {
	        var firstVar       = context.identifierList()?.IDENTIFIER(0);

	        var identifierList = context.identifierList();
		
            if (m_constIdentifierCount == 0 && m_constMultipleDeclaration)
                m_targetOutputFile.Append(RemoveFirstLineFeed(CheckForCommentsLeft(context)));

            if (!Identifiers.TryGetValue(context.identifierList(), out string[] identifiers))
            {
                AddWarning(context, $"No identifiers specified in constant expression: {context.GetText()}");
                return;
            }

            ExpressionLists.TryGetValue(context.expressionList(), out ExpressionInfo[] expressions);

            if (expressions is not null && identifiers.Length != expressions.Length)
            {
                AddWarning(context, $"Encountered identifier to expression count mismatch in constant expression: {context.GetText()}");
                return;
            }

			for (int i = 0; i < identifiers.Length; i++)
	        {
	            if (expressions is null || expressions.Length > i)
	                m_variableIdentifiers.Add(identifierList.IDENTIFIER(i), GetUniqueIdentifier(m_variableIdentifiers, identifiers[i]));
	        }
            Types.TryGetValue(context.type_(), out TypeInfo typeInfo);

            string type = typeInfo?.TypeName;
            int length = Math.Min(identifiers.Length, expressions?.Length ?? int.MaxValue);

            for (int i = 0; i < identifiers.Length; i++)
            {
                string identifier = identifiers[i];
                //string expression = expressions?[i].Text ?? $"{m_iota++}";
                //string typeName = type ?? expressions?[i].Type.TypeName ?? "var";
                //string castAs = string.Empty;
				
	            var _typeInfo = typeInfo;
	            if( context.type_() == null && context.expressionList() != null )
	            {
	                _typeInfo = expressions[i].Type; 
	            }

	            string variableName = firstVar != null? firstVar.GetText() : null;
	            bool isInitialDeclaration = true;
	            VariableInfo variable = null;
	            //bool heapAllocated = false;
	            //bool defaultInit = false;

	            m_targetOutputFile.Append($"{Spacing()}");

	            if (!InFunction)
	                m_targetOutputFile.Append($"{(char.IsUpper(identifier[0]) ? "public" : "private")} static ");

	            // Determine if this is the initial declaration
	            if (InFunction && m_variableIdentifiers.TryGetValue(identifierList.IDENTIFIER(i), out variableName))
	                isInitialDeclaration = !variableName.Contains("@@");

	            if (CurrentFunction != null && isInitialDeclaration && !string.IsNullOrWhiteSpace(variableName))
	            {
	                CurrentFunction.Variables.TryGetValue(variableName, out variable);
	                if( typeInfo == null && variable != null )
	                    typeInfo = variable.Type;

	                if( typeInfo == null)
	                    Debugger.Break();
	            }
	            if (isInitialDeclaration && _typeInfo!= null)
	            {
	                addVariable( variableName, _typeInfo, false);
	                if( !this.InFunction )
	                {
	                    var varInfo = new VariableInfo( variableName, _typeInfo);
                        if( LastDictElement != null)
	                        varInfo.setInitStructExpr(LastDictElement);
                        else
                            varInfo.setConstExpr( expressions[i] );
                    
	                    m_globalVariables.Add(variableName, varInfo);
	                }
	            }

                // Since multiple specifications can be on one line, only check for comments after last specification
                //if (i < length - 1)
                 //   m_targetOutputFile.AppendLine();
                //else
                //    m_targetOutputFile.Append(CheckForCommentsRight(context));
            }
	        m_varIdentifierCount++;
        }

        m_constIdentifierCount++;
    }
}
