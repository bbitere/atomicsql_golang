﻿//******************************************************************************************************
//  Converter_Statement.cs - Gbtc
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
using System.Linq;
using System.Text;
using static goscanner.Common;

namespace goscanner.ConvASqlModels;

public partial class ConvASqlModels
{
    private readonly ParseTreeValues<string> m_simpleStatements = new();
    private readonly ParseTreeValues<TypeInfo> m_elementTypes = new();
    private readonly Dictionary<string, bool> m_labels = new(StringComparer.Ordinal);
    private readonly Stack<HashSet<string>> m_blockLabeledContinues = new();
    private readonly Stack<HashSet<string>> m_blockLabeledBreaks = new();
        
    public override void ExitStatement(GoParser.StatementContext context)
    {
        if (EMITTING_CODE)
        {
            if (context.simpleStmt() is not null && context.simpleStmt().emptyStmt() is null)
            {
                if (m_simpleStatements.TryGetValue(context.simpleStmt(), out string statement))
                    m_targetOutputFile.Append($"{statement}{(LineTerminatorAhead(context.simpleStmt()) ? string.Empty : Environment.NewLine)}");
                else
                    AddWarning(context, $"Failed to find simple statement for: {context.simpleStmt().GetText()}");
            }
        }
    }

    public override void EnterLabeledStmt(GoParser.LabeledStmtContext context)
    {
        // labeledStmt
        //     : IDENTIFIER ':' statement

        PushBlock();
        if (EMITTING_CODE)
        { 

            string label = SanitizedIdentifier(context.IDENTIFIER()?.GetText());

            if (!string.IsNullOrEmpty(label) && !m_labels.ContainsKey(label))
                m_labels.Add(label, false);

            // Check labeled continue in for loop
            // Check labeled break in for loop, select and switch
        }
    }

    public override void ExitLabeledStmt(GoParser.LabeledStmtContext context)
    {
        // labeledStmt
        //     : IDENTIFIER ':' statement
        if (EMITTING_CODE)
        { 
            string label = SanitizedIdentifier(context.IDENTIFIER()?.GetText());
            string statement = PopBlock(false);

            if (!string.IsNullOrEmpty(label))
                m_targetOutputFile.Append($"{label}:");

            m_targetOutputFile.Append($"{CheckForCommentsRight(context)}");

            if (!WroteLineFeed)
                m_targetOutputFile.AppendLine();

            m_targetOutputFile.Append(statement);
        }
    }

    public override void ExitSendStmt(GoParser.SendStmtContext context)
    {
        // sendStmt
        //     : expression '<-' expression

        if (EMITTING_CODE)
        { 
            if (Expressions.TryGetValue(context.expression(0), out ExpressionInfo channel) && Expressions.TryGetValue(context.expression(1), out ExpressionInfo value))
            {
                StringBuilder statement = new();

                statement.Append($"{Spacing()}{channel}.Send({value});{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    statement.AppendLine();

                m_simpleStatements[context.Parent] = statement.ToString();
            }
            else
            {
                AddWarning(context, $"Failed to find both channel and value expression for send statement: {context.GetText()}");
            }
        }
    }

    public override void ExitExpressionStmt(GoParser.ExpressionStmtContext context)
    {
        // expressionStmt
        //     : expression

        if (EMITTING_CODE)
        { 
            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                StringBuilder statement = new();

                statement.Append($"{Spacing()}{expression};{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    statement.AppendLine();

                m_simpleStatements[context.Parent] = statement.ToString();
            }
            else
            {
                AddWarning(context, $"Failed to find expression for expression statement: {context.GetText()}");
            }
        }
    }

    public override void ExitIncDecStmt(GoParser.IncDecStmtContext context)
    {
        // incDecStmt
        //     : expression('++' | '--')
        if (EMITTING_CODE)
        { 
            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                StringBuilder statement = new();

                statement.Append($"{Spacing()}{expression}{context.children[1].GetText()};{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    statement.AppendLine();

                m_simpleStatements[context.Parent] = statement.ToString();
            }
            else
            {
                AddWarning(context, $"Failed to find expression for inc/dec statement: {context.GetText()}");
            }
        }
    }

    /// <remarks>
    /// See related operations:
    /// <see cref="ExitShortVarDecl(GoParser.ShortVarDeclContext)"/>
    /// <see cref="ExitVarSpec(GoParser.VarSpecContext)"/>
    /// </remarks>
    public override void ExitAssignment(GoParser.AssignmentContext context)
    {
        if (EMITTING_CODE)
        { 
            // assignment
            //     : expressionList assign_op expressionList
                if (ExpressionLists.TryGetValue(context.expressionList(0), out ExpressionInfo[] leftOperands) && ExpressionLists.TryGetValue(context.expressionList(1), out ExpressionInfo[] rightOperands))
            {
                if (leftOperands.Length != rightOperands.Length)
                {
                    if (leftOperands.Length > rightOperands.Length && rightOperands.Length == 1)
                    {
                        leftOperands = new[]
                        {
                            new ExpressionInfo 
                            { 
                                Text = string.Join(", ", leftOperands.Select(expr => expr.Text)),
                                Type = leftOperands[0].Type
                            }
                        };
                    }
                    else
                    {
                        AddWarning(context, $"Encountered count mismatch for left and right operand expressions in assignment statement: {context.GetText()}");
                    }
                }

                int length = Math.Min(leftOperands.Length, rightOperands.Length);
                string assignOP = context.assign_op().GetText();

                if (assignOP.Equals("=") && length > 1 && leftOperands.Length == rightOperands.Length)
                {
                    // TODO: Check for edge cases related to this tuple assignment operation, e.g., pointers, does this need "_addr_" prefix?
                    // This handles assignment order of operations, see https://goscanner.net/ConversionStrategies.html#inline-assignment-order-of-operations
                    m_simpleStatements[context.Parent] = $"{Spacing()}({string.Join(", ", leftOperands.Select(op => op.Text))}) {assignOP} ({string.Join(", ", rightOperands.Select(op => op.Text))});";
                }
                else
                {
                    StringBuilder statement = new();
                    
                    for (int i = 0; i < length; i++)
                    {
                        string leftOperandText = leftOperands[i].Text;
                        string rightOperandText = rightOperands[i].Text;

                        if (assignOP.Equals("<<=") || assignOP.Equals(">>="))
                        {
                            // TODO: Need expression evaluation - cast not needed for int expressions
                            // FYI, .NET 11 supports other integral types for shift operator, test with nint
                            if (!int.TryParse(rightOperands[i].Text, out int _))
                                rightOperands[i].Text = $"(int)({rightOperands[i]})";
                        }
                        else if (assignOP.Equals("&^="))
                        {
                            assignOP = " &= ";
                            rightOperands[i].Text = $"~({rightOperands[i].Text})";
                        }
                        else
                        {
                            if (!m_variableTypes.TryGetValue(leftOperandText, out TypeInfo leftOperandType))
                                leftOperandType = leftOperands[i].Type;

                            if (assignOP == "=" && leftOperandType?.TypeClass == TypeClass.Interface)
                                rightOperandText = $"{leftOperandType.TypeName}.As({rightOperandText})!";

                            if (assignOP == "=" && leftOperandType is not PointerTypeInfo && rightOperandText.StartsWith(AddressPrefix, StringComparison.Ordinal))
                            {
                                string targetVariable = rightOperandText.Replace(AddressPrefix, string.Empty);

                                if (m_variableTypes.TryGetValue(targetVariable, out TypeInfo rightOperandType) && rightOperandType is not PointerTypeInfo)
                                {
                                    rightOperandText = $"{rightOperandText};{Environment.NewLine}{Spacing()}{leftOperandText} = ref {AddressPrefix}{leftOperandText}.val";
                                    leftOperandText = $"{AddressPrefix}{leftOperandText}";
                                }
                            }

                            if (assignOP == "=" && leftOperandType is PointerTypeInfo)
                            {
                                string targetVariable = rightOperandText.Replace(AddressPrefix, string.Empty);

                                if (m_variableTypes.TryGetValue(targetVariable, out TypeInfo rightOperandType) && rightOperandType is PointerTypeInfo)
                                    rightOperandText = $"addr({targetVariable})";
                            }

                            assignOP = $" {assignOP} ";
                        }

                        statement.Append($"{Spacing()}{leftOperandText}{assignOP}{rightOperandText};");

                        // Since multiple assignments can be on one line, only check for comments after last assignment
                        if (i < length - 1)
                        {
                            statement.AppendLine();
                        }
                        else
                        {
                            statement.Append(CheckForCommentsRight(context));

                            if (!WroteLineFeed)
                                statement.AppendLine();
                        }

                        m_simpleStatements[context.Parent] = statement.ToString();
                    }
                }
            }
            else
            {
                AddWarning(context, $"Failed to find both left and right operand expressions for assignment statement: {context.GetText()}");
            }
        }
    }

    public override void ExitElementType(GoParser.ElementTypeContext context)
    {
        if (EMITTING_CODE)
        { 
            if (Types.TryGetValue(context.type_(), out TypeInfo typeInfo))
                m_elementTypes[context] = typeInfo;
            else
                AddWarning(context, $"Failed to find type info for: {context.GetText()}");
        }
    }

    //public override void ExitArrayType(GoParser.ArrayTypeContext context)
    //{
    //    base.ExitArrayType(context);
    //}

    //public override void ExitLiteralType(GoParser.LiteralTypeContext context)
    //{
    //    base.ExitLiteralType(context);
    //}

    //public override void ExitLiteralValue(GoParser.LiteralValueContext context)
    //{
    //    base.ExitLiteralValue(context);
    //}

    public override void ExitGoStmt(GoParser.GoStmtContext context)
    {
        // goStmt
        //     : 'go' expression
        if (EMITTING_CODE)
        { 
            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                RequiredUsings.Add("System.Threading");
                m_targetOutputFile.Append($"{Spacing()}go_(() => {expression});{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    m_targetOutputFile.AppendLine();
            }
            else
            {
                AddWarning(context, $"Failed to find expression for go statement: {context.GetText()}");
            }
        }
    }

    public override void EnterReturnStmt([NotNull] GoParser.ReturnStmtContext context)
    {
        if (EMITTING_CODE)
        { 
            base.EnterReturnStmt(context); 
        }
    }

    public override void ExitReturnStmt(GoParser.ReturnStmtContext context)
    {
        // returnStmt
        //     : 'return' expressionList?

        if (EMITTING_CODE)
        { 
            m_targetOutputFile.Append($"{Spacing()}return ");

            if (ExpressionLists.TryGetValue(context.expressionList(), out ExpressionInfo[] expressions))
            {
                ParameterInfo[] resultParameters = CurrentFunction?.Signature.Signature.Result;

                if (expressions.Length > 1)
                    m_targetOutputFile.Append('(');

                for (int i = 0; i < expressions.Length; i++)
                {
                    if (i > 0)
                        m_targetOutputFile.Append(", ");

                    TypeInfo resultType = resultParameters?.Length > i ? resultParameters[i].Type : TypeInfo.ObjectType;

                    if (resultType?.TypeClass == TypeClass.Interface)
                        m_targetOutputFile.Append($"{resultType.TypeName}.As({expressions[i].ToString().Trim()}{(expressions[i].ToString().Trim().Equals("null") || expressions[i].Type is PointerTypeInfo || expressions[i].Type.TypeClass == TypeClass.Interface ? "!" : string.Empty)})!");
                    else if (resultType is PointerTypeInfo && expressions[i].Type is not PointerTypeInfo)
                        m_targetOutputFile.Append($"{AddressPrefix}{expressions[i].ToString().Trim()}!");
                    else
                        m_targetOutputFile.Append($"{expressions[i].ToString().Trim()}");
                }
                
                if (expressions.Length > 1)
                    m_targetOutputFile.Append(')');
            }

            m_targetOutputFile.Append($";{CheckForCommentsRight(context)}");

            if (!WroteLineFeed)
                m_targetOutputFile.AppendLine();
        }
    }

    public override void ExitBreakStmt(GoParser.BreakStmtContext context)
    {
        // breakStmt
        //     : 'break' IDENTIFIER ?
        if (EMITTING_CODE)
        { 
            bool breakHandled = false;

            if (context.IDENTIFIER() is not null)
            {
                string label = SanitizedIdentifier(context.IDENTIFIER().GetText());

                if (m_labels.ContainsKey(label))
                {
                    breakHandled = true;

                    foreach (HashSet<string> blockBreaks in m_blockLabeledBreaks)
                        blockBreaks.Add(label);

                    m_targetOutputFile.Append($"{Spacing()}_break{label} = true;{CheckForCommentsRight(context)}");

                    if (!WroteLineFeed)
                        m_targetOutputFile.AppendLine();

                    m_targetOutputFile.AppendLine($"{Spacing()}break;");
                }
            }

            if (!breakHandled)
            {
                m_targetOutputFile.Append($"{Spacing()}break;{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    m_targetOutputFile.AppendLine();
            }
        }
    }

    public override void ExitContinueStmt(GoParser.ContinueStmtContext context)
    {
        // continueStmt
        //     : 'continue' IDENTIFIER ?

        if (EMITTING_CODE)
        {

            bool continueHandled = false;

            if (context.IDENTIFIER() is not null)
            {
                string label = SanitizedIdentifier(context.IDENTIFIER().GetText());

                if (m_labels.ContainsKey(label))
                {
                    continueHandled = true;

                    foreach (HashSet<string> blockContinues in m_blockLabeledContinues)
                        blockContinues.Add(label);

                    m_targetOutputFile.Append($"{Spacing()}_continue{label} = true;{CheckForCommentsRight(context)}");

                    if (!WroteLineFeed)
                        m_targetOutputFile.AppendLine();

                    m_targetOutputFile.AppendLine($"{Spacing()}break;");
                }
            }

            if (!continueHandled)
            {
                m_targetOutputFile.Append($"{Spacing()}continue;{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    m_targetOutputFile.AppendLine();
            }
        }
    }

    public override void ExitGotoStmt(GoParser.GotoStmtContext context)
    {
        // gotoStmt
        //     : 'goto' IDENTIFIER

        if (EMITTING_CODE)
        {
            m_targetOutputFile.Append($"{Spacing()}goto {SanitizedIdentifier(context.IDENTIFIER().GetText())};{CheckForCommentsRight(context)}");

            if (!WroteLineFeed)
                m_targetOutputFile.AppendLine();
        }
    }

    public override void ExitSelectStmt(GoParser.SelectStmtContext context)
    {
        // selectStmt
        //     : 'select' '{' commClause * '}'
    }

    public override void ExitDeferStmt(GoParser.DeferStmtContext context)
    {
        // deferStmt
        //     : 'defer' expression

        if (EMITTING_CODE)
        {

            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                m_targetOutputFile.Append($"{Spacing()}defer({expression});{CheckForCommentsRight(context)}");

                if (!WroteLineFeed)
                    m_targetOutputFile.AppendLine();
            }
            else
            {
                AddWarning(context, $"Failed to find expression for defer statement: {context.GetText()}");
            }
        }
    }
}
