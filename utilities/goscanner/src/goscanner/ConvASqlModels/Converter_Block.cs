//******************************************************************************************************
//  Converter_Block.cs - Gbtc
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

using System;
using System.Collections.Generic;
using System.Text;

namespace goscanner.ConvASqlModels;

public partial class ConvASqlModels
{
    private readonly Stack<StringBuilder> m_blocks = new();
    private readonly Stack<string> m_blockInnerPrefixInjection = new();
    private readonly Stack<string> m_blockInnerSuffixInjection = new();
    private readonly Stack<string> m_blockOuterPrefixInjection = new();
    private readonly Stack<string> m_blockOuterSuffixInjection = new();
    private bool m_firstStatementIsReturn;

    private void PushBlock()
    {
        m_blocks.Push(m_targetOutputFile);
        m_targetOutputFile = new StringBuilder();
    }

    private string PopBlock(bool appendToPrevious = true)
    {
        StringBuilder lastTarget = m_blocks.Pop();
        string block = RemoveLastDuplicateLineFeed(m_targetOutputFile.ToString());

        if (appendToPrevious)
            lastTarget.Append(block);

        m_targetOutputFile = lastTarget;

        return block;
    }

    private void PushInnerBlockPrefix(string prefix)
    {
        m_blockInnerPrefixInjection.Push(prefix);
    }

    private void PushInnerBlockSuffix(string suffix)
    {
        m_blockInnerSuffixInjection.Push(suffix);
    }

    private void PushOuterBlockPrefix(string prefix)
    {
        m_blockOuterPrefixInjection.Push(prefix);
    }

    private void PushOuterBlockSuffix(string suffix)
    {
        m_blockOuterSuffixInjection.Push(suffix);
    }

    public override void EnterBlock(GoParser.BlockContext context)
    {
        // block
        //     : '{' statementList '}'

        // statementList
        //     : (statement eos )*
        if( EMITTING_CODE)
        {    
            PushBlock();

            if (m_blockOuterPrefixInjection.Count > 0)
                m_targetOutputFile.Append(m_blockOuterPrefixInjection.Pop());

            if (Options.UseAnsiBraceStyle)
                m_targetOutputFile.Append($"{Spacing()}{{");
            else
                m_targetOutputFile.Append(" {");

            if (m_blockInnerPrefixInjection.Count > 0)
                m_targetOutputFile.Append(m_blockInnerPrefixInjection.Pop());

            if (context.statementList() is null)
                m_targetOutputFile.Append(RemoveFirstDuplicateLineFeed(CheckForCommentsRight(context.children[0], 1)));
            else
                m_targetOutputFile.Append(RemoveFirstDuplicateLineFeed(CheckForCommentsLeft(context.statementList(), 1)));

            if (!WroteLineFeed)
                m_targetOutputFile.AppendLine();
        }

        m_firstStatementIsReturn = false;

        IndentLevel++;
    }

    public override void ExitBlock(GoParser.BlockContext context)
    {
        IndentLevel--;

        GoParser.StatementListContext statementListContext = context.statementList();

        if (EMITTING_CODE)
        {
            if (statementListContext?.statement()?.Length > 0)
                m_firstStatementIsReturn = statementListContext.statement(0).returnStmt() is not null;

            if (m_blockInnerSuffixInjection.Count > 0)
                m_targetOutputFile.Append(m_blockInnerSuffixInjection.Pop());

            if (!EndsWithLineFeed(m_targetOutputFile.ToString()))
                m_targetOutputFile.AppendLine();
            else
                m_targetOutputFile = new StringBuilder(RemoveLastDuplicateLineFeed(m_targetOutputFile.ToString()));

            m_targetOutputFile.Append($"{Spacing()}}}");

            if (m_blockOuterSuffixInjection.Count > 0)
                m_targetOutputFile.Append(m_blockOuterSuffixInjection.Pop());

            if (!m_firstTopLevelDeclaration && IndentLevel > 2)
                m_targetOutputFile.Append(CheckForCommentsRight(context));
        }

        PopBlock();
    }


    private string RemoveFirstDuplicateLineFeed(string line)
    {
        string trimmedLine = line.TrimStart(' ', '\t');

        int index = trimmedLine.IndexOf("\r\n\r\n", StringComparison.Ordinal);

        if (index == 0)
            return trimmedLine[2..];

        index = trimmedLine.IndexOf("\n\n", StringComparison.Ordinal);

        if (index == 0)
            return trimmedLine[1..];

        return line;
    }

    private string RemoveLastDuplicateLineFeed(string line)
    {
        string trimmedLine = line.TrimEnd(' ', '\t');

        int index = trimmedLine.LastIndexOf("\r\n\r\n", StringComparison.Ordinal);

        if (index == trimmedLine.Length - 4)
            return trimmedLine[..^2];

        index = trimmedLine.LastIndexOf("\n\n", StringComparison.Ordinal);

        if (index == trimmedLine.Length - 2)
            return trimmedLine[..^1];

        return line;
    }
}
