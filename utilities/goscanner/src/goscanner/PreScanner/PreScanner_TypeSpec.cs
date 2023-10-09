//******************************************************************************************************
//  PreScanner_TypeSpec.cs - Gbtc
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
//  06/21/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using goscanner.Metadata;
using System.Collections.Generic;
using System.Diagnostics;
using System.Xml.Linq;

namespace goscanner;

public partial class PreScanner
{
    public override void ExitTypeSpecGeneric( GoParser.TypeSpecGenericContext context)
    {
        string identifier = context.IDENTIFIER().GetText();
        if( identifier.Contains( "User"))
            Debugger.Break();

        var types = new List<ParameterInfo>();


        if( context.typeSpecGenericList() != null)
        {
            foreach( var t in context.typeSpecGenericList().typeSpecGenericItem() )
            {
                if (Types.TryGetValue(t.type_(), out TypeInfo typeInfo))
                    types.Add( new_ParameterInfo( t.IDENTIFIER().GetText(), typeInfo ) );
            }
        }

        if (m_interfaceMethods.TryGetValue(context.type_()?.typeLit()?.interfaceType(), out List<FunctionSignature> methods))
        {
            m_interfaces.Add(GetUniqueIdentifier(m_interfaces, identifier), new InterfaceInfo
            {
                Name = identifier,
                Methods = methods.ToArray(),
                //GenericTypes = types.ToArray()
            });
        }
        else if (m_structFields.TryGetValue(context.type_()?.typeLit()?.structType(), out List<FieldInfo> fields))
        {
            var comment = "";
            //m_structComments.TryGetValue(context.type_()?.typeLit()?.structType(), out comment);

            m_structs.Add( GetUniqueIdentifier(m_structs, identifier), new StructInfo
            {
                Name = identifier,
                Fields = fields.ToArray(),
                GenericTypes = types.ToArray(),
                Comment = comment,

            });
        }
    }
    public override void ExitTypeSpec(GoParser.TypeSpecContext context)
    {
        string identifier = context.IDENTIFIER().GetText();
        //if( identifier.Contains( "User"))
         //   Debugger.Break();

        

        if (m_interfaceMethods.TryGetValue(context.type_()?.typeLit()?.interfaceType(), out List<FunctionSignature> methods))
        {
            m_interfaces.Add(GetUniqueIdentifier(m_interfaces, identifier), new InterfaceInfo
            {
                Name = identifier,
                Methods = methods.ToArray()
            });
            return;
        }

        var structTypeCtx = context.type_()?.typeLit()?.structType();
        if (m_structFields.TryGetValue(structTypeCtx, out List<FieldInfo> fields))
        {
            var comment =  CheckForCommentsLeft( structTypeCtx.L_CURLY() );

            //var comment1 =  CheckForCommentsInRange( structTypeCtx.STRUCT(), 
            //                                        structTypeCtx.L_CURLY() );

            //m_structComments.TryGetValue(context.type_()?.typeLit()?.structType(), out comment);

            m_structs.Add(GetUniqueIdentifier(m_structs, identifier), new StructInfo
            {
                Name = identifier,
                Fields = fields.ToArray(),
                Comment = cleanComment(comment),
            });
            return ;
        }
    }


    private string cleanComment(string comment1)
    {
        if( comment1 == null || comment1 == "" )
            return "";

        comment1 = comment1.Trim();
        if( comment1.StartsWith("/*")
             && comment1.EndsWith("* /") )
        {
            return comment1.Substring("/*".Length, comment1.Length - ("/*".Length + "* /".Length) );
        }else
        if( comment1.StartsWith("//") )
        {
            return comment1.Substring("//".Length, comment1.Length - ("//".Length) );
        }
        return comment1;
    }
}
