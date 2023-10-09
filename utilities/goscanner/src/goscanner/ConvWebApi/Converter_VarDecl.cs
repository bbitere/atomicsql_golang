//******************************************************************************************************
//  Converter_VarDecl.cs - Gbtc
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
//  07/12/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using goscanner.Metadata;
using System;
using System.Collections.Generic;
using System.Linq;
using static goscanner.Common;

namespace goscanner.ConvWebApi;

public partial class ConvWebApi
{
    private readonly ParseTreeValues<string> m_variableIdentifiers = new();
    private readonly Dictionary<string, TypeInfo> m_variableTypes = new();
    private int m_varIdentifierCount;
    private bool m_varMultipleDeclaration;

    private string GetUniqueIdentifier(ParseTreeValues<string> source, string identifier)
    {
        int count = 0;
        string uniqueIdentifier = identifier;

        while (source.ContainsValue(uniqueIdentifier))
            uniqueIdentifier = $"{identifier}@@{++count}";

        return uniqueIdentifier;
    }

    public override void EnterVarDecl(GoParser.VarDeclContext context)
    {
        // varDecl
        //     : 'var' ( varSpec | '(' ( varSpec eos )* ')' )

        m_varIdentifierCount = 0;
        m_varMultipleDeclaration = context.children.Count > 2;
    }

    public override void ExitVarDecl(GoParser.VarDeclContext context)
    {
        // varDecl
        //     : 'var' ( varSpec | '(' ( varSpec eos )* ')' )

        if (m_varMultipleDeclaration && EndsWithLineFeed(m_targetOutputFile.ToString()))
        {
            string removedLineFeed = RemoveLastLineFeed(m_targetOutputFile.ToString());
            m_targetOutputFile.Clear();
            m_targetOutputFile.Append(removedLineFeed);
        }

        m_targetOutputFile.Append(CheckForCommentsRight(context));
    }

    /// <remarks>
    /// See related operations:
    /// <see cref="ExitShortVarDecl(GoParser.ShortVarDeclContext)"/>
    /// <see cref="ExitAssignment(GoParser.AssignmentContext)"/>
    /// </remarks>
    public override void ExitVarSpec(GoParser.VarSpecContext context)
    {
        // varSpec
        //     : identifierList ( type ( '=' expressionList ) ? | '=' expressionList )

        if (m_varIdentifierCount == 0 && m_varMultipleDeclaration)
            m_targetOutputFile.Append(RemoveFirstLineFeed(CheckForCommentsLeft(context)));

        GoParser.IdentifierListContext identifierList = context.identifierList();

        if (!Identifiers.TryGetValue(identifierList, out string[] identifiers))
        {
            AddWarning(context, $"No identifiers specified in var specification expression: {context.GetText()}");
            return;
        }

        ExpressionLists.TryGetValue(context.expressionList(), out ExpressionInfo[] expressions);

        if (expressions is not null && identifiers.Length != expressions.Length)
        {
            AddWarning(context, $"Encountered count mismatch for identifiers and expressions in var specification expression: {context.GetText()}");
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
        bool isPointer = typeInfo is PointerTypeInfo;

        for (int i = 0; i < length; i++)
        {
            string identifier = SanitizedIdentifier(identifiers[i]);
            string expression = expressions?[i].Text;
            string typeName = type ?? expressions?[i].Type.TypeName ?? "var";
            string variableName = null;
            bool isInitialDeclaration = true;
            VariableInfo variable = null;
            bool heapAllocated = false;
            bool defaultInit = false;

            if( this.Options.ConvertWebApi.ExportEnums.Contains( identifier ) )
            {
                var exprInit = expressions?[i];
                if( exprInit != null ) 
                {
                    /*
                    var typeName1 = exprInit.Type.TypeName;
                    var structType = getStructInfoByName( typeName1 );
                    if( structType != null )
                    {
                        var fields = structType.GetLocalFields();
                        var fieldsList = fields.ToList();
                        //bbitere
                        var dict =  new Dictionary<string, string>();
                        var offset = exprInit.SubExpressions.Count-1 - fieldsList.Count;
                        for( var iExpr = 0; iExpr < exprInit.SubExpressions.Count; iExpr++ ) 
                        {
                            if( iExpr >= fieldsList.Count)
                                break;
                            var key = fieldsList[ iExpr ].Name;
                            dict[key] = exprInit.SubExpressions[offset + iExpr];
                        }
                        this.Options.ConvertWebApi.AddEnumExpression( identifier, dict );
                    }*/
                    this.Options.ConvertWebApi.AddEnumExpression( identifier, LastDictElement );
                }
            }

            m_targetOutputFile.Append($"{Spacing()}");

            if (!InFunction)
                m_targetOutputFile.Append($"{(char.IsUpper(identifier[0]) ? "public" : "private")} static ");

            // Determine if this is the initial declaration
            if (InFunction && m_variableIdentifiers.TryGetValue(identifierList.IDENTIFIER(i), out variableName))
                isInitialDeclaration = !variableName.Contains("@@");

            if (isInitialDeclaration && !string.IsNullOrWhiteSpace(variableName))
            {
                m_variableTypes[variableName] = typeInfo;
                CurrentFunction.Variables.TryGetValue(variableName, out variable);
            }

            if (isInitialDeclaration)
            {
                if ((variable?.HeapAllocated ?? false) && !isPointer)
                {
                    heapAllocated = true;
                    m_targetOutputFile.Append($"ref {typeName} {identifier} = ref heap(");
                }
                else
                {
                    m_targetOutputFile.Append($"{typeName} {identifier}");

                    if (!isPointer)
                        m_targetOutputFile.Append(" = ");
                }
            }
            else
            {
                m_targetOutputFile.Append($"{identifier} = ");
            }

            if (expression is not null)
            {
                if (typeInfo?.TypeClass == TypeClass.Interface)
                    m_targetOutputFile.Append($"{typeInfo.TypeName}.As({expression})!");
                else
                    m_targetOutputFile.Append(expression);
            }
            else if (typeInfo?.TypeClass == TypeClass.Array && typeInfo is ArrayTypeInfo arrayTypeInfo)
            {
                m_targetOutputFile.Append($"new {typeName}({arrayTypeInfo.Length.Text})");
            }
            else
            {
                defaultInit = true;

                if (!heapAllocated && !isPointer)
                    m_targetOutputFile.Append($"default{((typeInfo ?? expressions?[i].Type)?.TypeClass == TypeClass.Interface ? "!" : string.Empty)}");
            }

            if (heapAllocated)
            {
                if (!defaultInit)
                    m_targetOutputFile.Append(", ");

                m_targetOutputFile.Append($"out ptr<{typeName}> {AddressPrefix}{identifier})");
            }

            // Since multiple specifications can be on one line, only check for comments after last specification
            if (i < length - 1 || length == 1)
                m_targetOutputFile.Append(";");
            else
                m_targetOutputFile.Append($";{CheckForCommentsRight(context)}");
        }

        m_varIdentifierCount++;
    }
}
