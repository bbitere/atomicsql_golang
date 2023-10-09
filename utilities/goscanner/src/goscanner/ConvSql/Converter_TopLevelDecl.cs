//******************************************************************************************************
//  Converter_TopLevelDecl.cs - Gbtc
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
//  05/03/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using System;
using System.Text;

namespace goscanner.ConvSql;

public partial class SqlConvert
{
    private bool m_firstTopLevelDeclaration = true;

    // TopLevelDecl is visited once per each encountered Declaration, FunctionDecl or MethodDecl
    public override void EnterTopLevelDecl(GoParser.TopLevelDeclContext context)
    {
        if (EMITTING_CODE)
        {
            if (m_firstTopLevelDeclaration)
            {
                StringBuilder lineBreaks = new();

                if (!EndsWithDuplicateLineFeed(m_targetOutputFile.ToString()))
                {
                    if (EndsWithLineFeed(m_targetOutputFile.ToString()))
                        lineBreaks.AppendLine();
                    else
                        lineBreaks.AppendLine(Environment.NewLine);
                }

                // Mark end of using statements so that other usings and type aliases can be added later
                m_targetOutputFile.Append(UsingsMarker);
            
                m_targetOutputFile.Append(lineBreaks);

                // Begin class
                m_targetOutputFile.Append($"public static {UnsafeMarker}partial class {Package}{ClassSuffix} {{");

                // End class and namespace braces occur as a last step in Convert() method

                // Check for comments before initial declaration
                string initialDeclComments = CheckForCommentsLeft(context);

                // Write any initial declaration comments post any final EOL comments in Converter_ImportDecl visit 
                if (!string.IsNullOrEmpty(initialDeclComments) && initialDeclComments != m_lastImportDeclComment)
                {
                    m_targetOutputFile.Append(initialDeclComments);
                }
                else if (!EndsWithDuplicateLineFeed(m_targetOutputFile.ToString()))
                {
                    if (EndsWithLineFeed(m_targetOutputFile.ToString()))
                        m_targetOutputFile.AppendLine();
                    else
                        m_targetOutputFile.AppendLine(Environment.NewLine);
                }
            }
        }
    }

    public override void ExitTopLevelDecl(GoParser.TopLevelDeclContext context)
    {
        if (EMITTING_CODE)
        {

            // There can be only one... first top level declaration
            if (m_firstTopLevelDeclaration)
                m_firstTopLevelDeclaration = false;
        }
    }
}
