﻿//******************************************************************************************************
//  Converter_FunctionLit.cs - Gbtc
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

using Antlr4.Runtime.Misc;
using goscanner.Metadata;
using System;

namespace goscanner.ConvSql;

public partial class SqlConvert
{
    public const string FunctionLiteralParametersMarker = ">>MARKER:FUNCTIONLIT_PARAMETERS<<";

    public override void EnterFunctionLit(GoParser.FunctionLitContext context)
    {
        var block = PushBlock();
        if (EMITTING_CODE)
        { 
            m_targetOutputFile.Append($"{FunctionLiteralParametersMarker} =>");

            if (Options.UseAnsiBraceStyle)
                m_targetOutputFile.AppendLine();
        }
        this.Lambda_enterFuncLit(context);

        var srcPath = context.Start.InputStream.SourceName;
        /*if( srcPath.EndsWith("test_misc.go") )    
        { 
            Utils.Nop();
        }*/
    }

    public override void ExitFunctionLit(GoParser.FunctionLitContext context)
    {
        if( m_LambdaCode != null )
        {
            m_LambdaCode.endSource(context);
        }
        // functionLit
        //     : 'func' function
        if (EMITTING_CODE)
        { 
            var srcPath = context.Start.InputStream.SourceName;
            if( srcPath.EndsWith("Test.go") )
                Utils.Nop();

            string parametersSignature = "()";

            if (Signatures.TryGetValue(context?.signature(), out Signature signature))
            {
                parametersSignature = signature.GenerateParameterNameList();

                if (signature.Parameters.Length != 1)
                    parametersSignature = $"({parametersSignature})";
            }
            else
            {
                AddWarning(context, $"Failed to find signature for function literal inside \"{CurrentFunctionName}\" function");
            }

            // Replace marker for function literal
            m_targetOutputFile.Replace(FunctionLiteralParametersMarker, parametersSignature);

            // operand
            //     : literal
            //     | operandName
            //     | methodExpr
            //     | '(' expression ')'

            // literal
            //     : basicLit
            //     | compositeLit
            //     | functionLit

            if (context?.Parent.Parent is not GoParser.OperandContext operandContext)
            {
                AddWarning(context, $"Could not derive parent operand context from function literal inside \"{CurrentFunctionName}\" function: \"{context?.GetText()}\"");
                PopBlock();
                return;
            }

            var lambdaExpression = PopBlock(false);

            // Simplify lambda expressions that consist of a single return statement
            if (m_firstStatementIsReturn)
            {
                int index = lambdaExpression.IndexOf("=>", StringComparison.Ordinal);

                if (index > -1)
                {
                    string startBlock = $"{{{Environment.NewLine}";
                    
                    index = lambdaExpression.IndexOf(startBlock, index, StringComparison.Ordinal);

                    if (index > -1)
                    {
                        string parameters = lambdaExpression[..index].Trim();

                        lambdaExpression = lambdaExpression[(index + startBlock.Length)..].Trim();

                        if (lambdaExpression.StartsWith("return ", StringComparison.Ordinal))
                            lambdaExpression = lambdaExpression[7..].Trim();

                        if (lambdaExpression.EndsWith("}", StringComparison.Ordinal))
                            lambdaExpression = lambdaExpression[..^1].Trim();

                        if (lambdaExpression.EndsWith(";", StringComparison.Ordinal))
                            lambdaExpression = lambdaExpression[..^1].Trim();

                        lambdaExpression = $"{parameters} {lambdaExpression}";
                    }
                }
            }

            RequiredUsings.Add("System");

            string typeList = signature.GenerateParameterTypeList();
            string resultSignature = signature.GenerateResultSignature();
            string typeName, fullTypeName;

            if (resultSignature == "void")
            {
                if (string.IsNullOrEmpty(typeList))
                {
                    typeName = "Action";
                    fullTypeName = "System.Action";
                }
                else
                {
                    typeName = $"Action<{typeList}>";
                    fullTypeName = $"System.Action<{typeList}>";
                }
            }
            else
            {
                if (!string.IsNullOrEmpty(typeList))
                    typeList = $"{typeList}, ";

                typeName = $"Func<{typeList}{resultSignature}>";
                fullTypeName = $"System.Func<{typeList}{resultSignature}>";
            }

            // Update expression operand (managed in ScannerBase_Expression.cs)
            Operands[operandContext] = new ExpressionInfo
            {
                LastToken = "",
                Text = lambdaExpression,
                Type = new TypeInfo
                {
                    Name = parametersSignature,
                    TypeName = typeName,
                    FullTypeName = fullTypeName,
                    TypeClass = TypeClass.Function
                }
            };
        }
    }
}
