//******************************************************************************************************
//  PreScanner_Type.cs - Gbtc
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

using Antlr4.Runtime;
using goscanner.Metadata;
using System.Collections.Generic;
using static goscanner.Common;

namespace goscanner;

public partial class PreScanner
{
    private readonly ParseTreeValues<List<FieldInfo>> m_structFields = new();
    //private readonly ParseTreeValues<string> m_structComments = new();
    private readonly ParseTreeValues<List<FunctionSignature>> m_interfaceMethods = new();

    public override void ExitStructType(GoParser.StructTypeContext context)
    {
        var comment = "";
        List<FieldInfo> fields = new();

        /*
        var comment1 =  "";//CheckForCommentsInRange(context.lis, context.L_CURLY());
        if( context.COMMENT() != null ) 
        {
            comment = context.COMMENT().Symbol.Text;
        }else
        if( comment1 != null && comment1.Trim() != "" ) 
        { 
            comment1 = comment1.Trim();
            if( comment1.StartsWith("/*")
             && comment1.EndsWith("* /") )
            {
                comment1 = comment1.Substring("/*".Length, comment1.Length - ("/*".Length + "* /".Length) );
            }else
            if( comment1.StartsWith("//") )
            {
                comment1 = comment1.Substring("//".Length, comment1.Length - ("//".Length) );
            }

            comment = comment1;
        }*/


        for (int i = 0; i < context.fieldDecl().Length; i++)
        {
            GoParser.FieldDeclContext fieldDecl = context.fieldDecl(i);
            string description = ToStringLiteral(fieldDecl.string_()?.GetText());

            if (Identifiers.TryGetValue(fieldDecl.identifierList(), out string[] identifiers) && Types.TryGetValue(fieldDecl.type_(), out TypeInfo typeInfo))
            {
                foreach (string identifier in identifiers)
                {
                    fields.Add(new FieldInfo
                    {
                        Name = SanitizedIdentifier(identifier),
                        Type = ConvertByRefToBasicPointer(typeInfo),
                        Description = description,
                        Comments = CheckForCommentsRight(fieldDecl),
                        IsPromoted = false,
                        Location = Getlocation(context)
                    });
                }
            }
            else
            {
                GoParser.AnonymousFieldContext anonymousField = fieldDecl.anonymousField();

                if (Types.TryGetValue(anonymousField, out typeInfo))
                {
                    if (anonymousField.ChildCount > 1 && anonymousField.children[0].GetText() == "*")
                    {
                        typeInfo = new PointerTypeInfo
                        {
                            Name = $"ptr<{typeInfo.TypeName}>",
                            TypeName = $"ptr<{typeInfo.TypeName}>",
                            FullTypeName = $"go.ptr<{typeInfo.FullTypeName}>",
                            TypeClass = typeInfo.TypeClass,
                            IsDerefPointer = typeInfo.IsDerefPointer,
                            IsByRefPointer = typeInfo.IsByRefPointer,
                            IsConst = typeInfo.IsConst,
                            TargetTypeInfo = typeInfo
                        };
                    }

                    fields.Add(new FieldInfo
                    {
                        Name = GetValidIdentifierName(typeInfo.TypeName),
                        Type = typeInfo,
                        Description = description,
                        Comments = CheckForCommentsRight(fieldDecl),
                        IsPromoted = true,
                        Location = Getlocation(context)
                    });
                }
            }
        }

        //m_structComments[context] = comment;
        m_structFields[context] = fields;
    }

    public override void EnterInterfaceType(GoParser.InterfaceTypeContext context)
    {
        base.EnterInterfaceType(context);
        
        Result = new List<ParameterInfo>(new[] { new ParameterInfo
        {
            Name = string.Empty,
            Type = TypeInfo.VoidType,
            IsVariadic = false
        }});
    }

    public override void ExitInterfaceType(GoParser.InterfaceTypeContext context)
    {
        List<FunctionSignature> methods = new();

        if (InterfaceTypes.TryGetValue(context, out ParseTreeValues<TypeInfo> types))
        {
            for (int i = 0; i < context.typeName().Length; i++)
            {
                GoParser.TypeNameContext typeName = context.typeName(i);
            
                if (types.TryGetValue(typeName, out TypeInfo typeInfo))
                {
                    methods.Add(new FunctionSignature
                    {
                        Name = GetValidIdentifierName(typeInfo.TypeName),
                        Signature = new Signature
                        {
                            Parameters = System.Array.Empty<ParameterInfo>(),
                            Result = new[]
                            {
                                new ParameterInfo
                                {
                                    Name = string.Empty,
                                    Type = typeInfo,
                                    IsVariadic = false
                                }
                            }
                        },
                        Comments = CheckForCommentsRight(typeName),
                        IsPromoted = true
                    });
                }
            }
        }

        for (int i = 0; i < context.methodSpec().Length; i++)
        {
            GoParser.MethodSpecContext methodSpec = context.methodSpec(i);

            string identifier = methodSpec.IDENTIFIER()?.GetText();

            Parameters.TryGetValue(methodSpec.parameters(), out List<ParameterInfo> parameters);

            methods.Add(new FunctionSignature
            {
                Name = identifier,
                Signature = Signatures[context] = new Signature
                {
                    Parameters = parameters?.ToArray() ?? System.Array.Empty<ParameterInfo>(),
                    Result = Result?.ToArray() ?? System.Array.Empty<ParameterInfo>()
                },
                Comments = CheckForCommentsRight(methodSpec),
                IsPromoted = false
            });
        }

        m_interfaceMethods[context] = methods;
    }
}
