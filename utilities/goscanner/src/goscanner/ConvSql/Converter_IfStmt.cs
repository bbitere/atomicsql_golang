//******************************************************************************************************
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

using goscanner.Metadata;
using System;

namespace goscanner.ConvSql;

public partial class SqlConvert
{
    public const string IfElseMarker = ">>MARKER:IFELSE_LEVEL_{0}<<";
    public const string IfElseBreakMarker = ">>MARKER:IFELSEBREAK_LEVEL_{0}<<";
    public const string IfExpressionMarker = ">>MARKER:IFEXPR_LEVEL_{0}<<";
    public const string IfStatementMarker = ">>MARKER:IFSTATEMENT_LEVEL_{0}<<";

    private int m_ifExpressionLevel;

    public override void EnterIfStmt(GoParser.IfStmtContext context)
    {
        // ifStmt
        //     : 'if '(simpleStmt ';') ? expression block ( 'else' ( ifStmt | block ) ) ?

        m_ifExpressionLevel++;
        if (EMITTING_CODE)
        { 
            if (context.simpleStmt() is not null && context.simpleStmt().emptyStmt() is null)
            {
                // Any declared variable will be scoped to if statement, so create a sub-block for it
                if (context.simpleStmt().shortVarDecl() is not null)
                {
                    m_targetOutputFile.AppendLine($"{Spacing()}{{");
                    IndentLevel++;

                    // Handle storing of current values of any redeclared variables
                    m_targetOutputFile.Append(OpenRedeclaredVariableBlock(context.simpleStmt().shortVarDecl().identifierList(), m_ifExpressionLevel));
                }

                m_targetOutputFile.Append(string.Format(IfStatementMarker, m_ifExpressionLevel));
            }

            m_targetOutputFile.Append($"{string.Format(IfElseBreakMarker, m_ifExpressionLevel)}{Spacing()}{string.Format(IfElseMarker, m_ifExpressionLevel)}if ({string.Format(IfExpressionMarker, m_ifExpressionLevel)})");

            if (Options.UseAnsiBraceStyle)
                m_targetOutputFile.AppendLine();

            if (context.block().Length == 2)
            {
                PushOuterBlockSuffix(null); // For current block
                PushOuterBlockSuffix($"{Environment.NewLine}{Spacing()}else{(LineTerminatorAhead(context.block(0)) ? string.Empty : Environment.NewLine)}");
            }
            else
            {
                PushOuterBlockSuffix(null); // For current block
            }
        }
    }

    public override void ExitIfStmt(GoParser.IfStmtContext context)
    {
        // ifStmt
        //     : 'if '(simpleStmt ';') ? expression block ( 'else' ( ifStmt | block ) ) ?

        if (EMITTING_CODE)
        { 
            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                bool isElseIf = context.Parent is GoParser.IfStmtContext;

                // Replace if markers
                m_targetOutputFile.Replace(string.Format(IfExpressionMarker, m_ifExpressionLevel), expression.Text);
                m_targetOutputFile.Replace(string.Format(IfElseBreakMarker, m_ifExpressionLevel), isElseIf ? Environment.NewLine : string.Empty);
                m_targetOutputFile.Replace(string.Format(IfElseMarker, m_ifExpressionLevel), isElseIf ? "else " : string.Empty);
            }
            else
            {
                AddWarning(context, $"Failed to find expression for if statement: {context.GetText()}");
            }

            if (context.simpleStmt() is not null && context.simpleStmt().emptyStmt() is null)
            {
                if (m_simpleStatements.TryGetValue(context.simpleStmt(), out string statement))
                    m_targetOutputFile.Replace(string.Format(IfStatementMarker, m_ifExpressionLevel), statement + Environment.NewLine);
                else
                    AddWarning(context, $"Failed to find simple statement for if statement: {context.simpleStmt().GetText()}");
                
                // Close any locally scoped declared variable sub-block
                if (context.simpleStmt().shortVarDecl() is not null)
                {
                    // Handle restoration of previous values of any redeclared variables
                    m_targetOutputFile.Append(CloseRedeclaredVariableBlock(context.simpleStmt().shortVarDecl().identifierList(), m_ifExpressionLevel));

                    IndentLevel--;
                    m_targetOutputFile.AppendLine();
                    m_targetOutputFile.Append($"{Spacing()}}}{CheckForCommentsRight(context)}");
                }
            }

            if (!EndsWithLineFeed(m_targetOutputFile.ToString()))
                m_targetOutputFile.AppendLine();

        }
        m_ifExpressionLevel--;
    }
}
