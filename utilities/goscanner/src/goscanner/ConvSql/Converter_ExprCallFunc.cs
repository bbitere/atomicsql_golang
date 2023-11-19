using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime.Misc;
using Gemstone.Units;
using goscanner.Metadata;

namespace goscanner.ConvSql
{
    public partial class SqlConvert
    {
        void ExpressionCall_AfterFirstChild( 
            GoParser.PrimaryExprContext context, 
            ExpressionInfo primaryExpression)
        {
            if( primaryExpression == null)
                return;
            if( context.Start.InputStream.SourceName.EndsWith("Test.go") )    
                Utils.Nop();

            string functionName = primaryExpression.Text;
          
            //FunctionInfo functionInfo = null;
            
            var funcName = getIdentOfFunction( functionName );
            if( funcName == OrmDef.Func_New_DBContext )
                Utils.Nop();

            var detailsType = new TGenTypeDetails();
            var normalizedType1 = primaryExpression.Type.getNormalizedType();

            if( funcName == OrmDef.Func_DBQuery_Where )
                Utils.Nop();
            if( funcName == OrmDef.Func_Select )                
                Utils.Nop();

            if( context.Start.InputStream.SourceName.EndsWith("test1.go") 
              &&  funcName == OrmDef.Func_DBQuery_Where )
            {
                if( normalizedType1 != null
                 && normalizedType1.Name == OrmDef.Class_DBQuery )
                { 
                    Debug_Console("normalizedType1: OrmDef.Class_DBQuery");
                }else
                if( normalizedType1 != null
                  )
                { 
                    Debug_Console($"normalizedType1: {normalizedType1.Name}");
                }else
                {
                    Debug_Console("normalizedType1: null");
                }
            }

            if( normalizedType1 != null
             && normalizedType1.Name == OrmDef.Class_DBQuery )
            {
                var subTagName= OrmDef.GetSubTabByFuncName( funcName );
                if( subTagName != "" )
                {
                    this.Lambda_createSubtagID(subTagName);
                }
                /*
                if( OrmDef.Func_DBQuery_End.Contains(funcName) )
                {
                    this.Lambda_endSequence();
                }*/
            }

            var packageName = getPackageNameOfFunction( functionName );
            if( packageName == Options.ConvertSql.OrmDir_Atomicsql_Git
             || ( normalizedType1 != null && normalizedType1.Name == OrmDef.Class_DBContext ) )
            {
                if( funcName == OrmDef.Func_Select )
                {
                    //m_LambdaFunc_SubTagCounter++;
                    //if( m_LambdaCode != null )
                    //    m_LambdaCode.PrevSubTag = m_LambdaSubTag;
                    //m_LambdaFunc_SubTags.Add( new SubTag($"{OrmDef.SubTag_Select}{m_LambdaFunc_SubTagCounter}") );
                    this.Lambda_createSubtagID(OrmDef.SubTag_Select);
                }
                if( funcName == OrmDef.Func_Aggregate )
                {
                    if( context.arguments()?.typeNameGenericList()?.Length == 1 )
                    {
                        var lst  = context.arguments()?.typeNameGenericList()?[0];
                        var arg0 = lst.typeName()[0];
                        var arg1 = lst.typeName()[1];

                        var type1 = arg0.GetText();
                        var type2 = arg1.GetText();

                        var err = Aggregate_checkArgumentsCall( type1, type2 );
                        if( err != "" )
                        {
                            Log_Error(context, err); 
                        }
                    }
                }
            }
        }
        void ExpressionCall( GoParser.PrimaryExprContext context, 
            ExpressionInfo primaryExpression, string packageImport )
        {
            // arguments
            //     : '('((expressionList | type(',' expressionList) ? ) '...' ? ',' ? ) ? ')'

            GoParser.ArgumentsContext argumentsContext = context.arguments();
            List<string> arguments = new();
            List<string> arguments_SQL = new();
            List<EOperandKind> arguments_OperandKind = new();

            Types.TryGetValue(argumentsContext.type_(), out TypeInfo typeInfo);

            // Attempt to lookup expression with arguments as a "function"
            ParameterInfo[] parameters = null;
            string functionName = primaryExpression.Text;
            FunctionInfo functionInfo = null;

            var funcName  = getIdentOfFunction( functionName );
            if( funcName == "Nopp" )
                Utils.Nop();

            if( TypeInfo.isPrimaryType(funcName) )
            {
                //cast: int64(expression)
                if( ExpressionLists.TryGetValue(argumentsContext.expressionList(),
                    out ExpressionInfo[] castExpressions))
                {
                    if( castExpressions.Length == 1 ) 
                    {
                        PrimaryExpressions[context] = new ExpressionInfo
                            {
                                LastToken = primaryExpression.LastToken,
                                Text = $"({funcName})({castExpressions[0].Text})",
                                Type  = castExpressions[0].Type,
                                Types = castExpressions[0].Types,
                                SQLText = castExpressions[0].SQLText,
                                OperandKind = castExpressions[0].OperandKind,
                            };
                        return;
                    }
                }
            }

            var detailsType = new TGenTypeDetails();

            var packageName = "";
            var normalizedType1 = primaryExpression.Type.getNormalizedType();
            if( normalizedType1 != null)
            {
                if( normalizedType1.Name == OrmDef.Class_DBQuery 
                  && OrmDef.Func_DBQuery_End.Contains(funcName) )
                {
                    this.Lambda_endSequence();
                }

                packageName = normalizedType1.PackageName;
                if( searchFunctionInImports( funcName, normalizedType1.PackageInfo, out functionInfo, 
                        USE_SHOW_DEBUGGER && funcName == "Qry" ))
                {
                }
                
            }else
            {
                packageName  = getPackageAliasOfFunction( functionName );
                if( packageName == "" )
                    packageName = primaryExpression.Type.PackageInfo.Name;
                var packageInfo = new TPackageInfo(packageName, primaryExpression.Type.PackageInfo.Metafile);

                // TODO: Need to lookup functions from imported libraries as well
                if( !searchFunctionInImports( funcName, packageInfo, out functionInfo ) )
                {
                    
                }
            }

            if (functionInfo is not null)
                parameters = functionInfo.Signature.Signature.Parameters;

            if (ExpressionLists.TryGetValue(argumentsContext.expressionList(), out ExpressionInfo[] expressions))
            {
                if (InFunction && CurrentFunction is not null)
                {
                    for (int i = 0; i < expressions.Length; i++)
                    {
                        ExpressionInfo expression = expressions[i];
                        ParameterInfo parameter = parameters?.Length > i ? parameters[i] : null;
                        CurrentFunction.Variables.TryGetValue(expression.Text, out VariableInfo variable);

                        if (parameter?.Type is PointerTypeInfo && expression.Type is not PointerTypeInfo
                            && variable?.Type is not PointerTypeInfo && !expression.Text.StartsWith(AddressPrefix, StringComparison.Ordinal))
                        {
                            arguments.Add($"{AddressPrefix}{expression}");
                        }else
                        {
                            arguments.Add(expression.Text);
                        }
                        arguments_SQL.Add( expression.SQLText );
                        arguments_OperandKind.Add( expression.OperandKind );
                    }
                }
                else
                {
                    arguments.AddRange(expressions.Select(expression => expression.Text));
                }
            }

            string argumentList = string.Join(", ", arguments);
            //string argumentSQLList = string.Join(", ", arguments_SQL);
            if( funcName == OrmDef.Func_DBTable_Qry && argumentList == "\"asdax\"u8")
            {
                Utils.Nop();
            }
            if( normalizedType1 != null 
             && funcName == OrmDef.Func_DBTable_Qry 
             && normalizedType1.Name == OrmDef.Class_DBTable 
             && normalizedType1.PackageName!= "" )
            {
                if(argumentList != "\"\"u8")
                    Debug_Console($"Func_DBTable_Qry: {argumentList}");
                this.Lambda_callQryMethod( context, argumentList);
            }else
            {
                var subTagName= OrmDef.GetSubTabByFuncName( funcName );
                if( normalizedType1 != null
                 && (subTagName != "" && normalizedType1.Name == OrmDef.Class_DBQuery ) )
                {
                    this.Lambda_callWhereMethod(context, subTagName );
                }else
                if( funcName == OrmDef.Func_Select )
                {
                    var packageTgtName = getPackageNameOfFunction( functionName );
                    if( packageTgtName == Options.ConvertSql.OrmDir_Atomicsql_Git ) 
                    {
                        this.Lambda_callSelectMethod( context );
                    }
                }
            }

            var SQLText = "";
            var OperandKind = EOperandKind.Function;

            if (primaryExpression.Text == "new")
            {
                if (typeInfo is null)
                {
                    string typeName = expressions?[0].Text;

                    SQLText = expressions?[0].SQLText;
                    TypeInfo argType = null;

                    foreach (TypeInfo typeInfoValue in Types.Values)
                    {
                        if (typeInfoValue.TypeName.Equals(typeName))
                        {
                            argType = typeInfoValue.Clone();
                            break;
                        }

                        if (typeInfoValue.TypeName.Equals($"{packageImport}.{typeName}"))
                        {
                            argType = typeInfoValue.Clone();
                            break;
                        }

                        foreach (string import in Imports)
                        {
                            if (typeInfoValue.TypeName.Equals($"{import}.{typeName}"))
                            {
                                argType = typeInfoValue.Clone();
                                break;
                            }
                        }

                        if (argType is not null)
                            break;
                            
                        if (typeInfoValue.Name.Equals(typeName))
                        {
                            argType = typeInfoValue.Clone();
                            break;
                        }
                    }

                    if (argType == null)
                        argType = TypeInfo.ObjectType.Clone();

                    argType = new PointerTypeInfo
                    {
                        Name = $"ptr<{argType.Name}>",
                        TypeName = $"ptr<{argType.Name}>",
                        FullTypeName = $"go.ptr<{argType.FullTypeName}>",
                        TypeClass = argType.TypeClass,
                        IsDerefPointer = argType.IsDerefPointer,
                        IsByRefPointer = argType.IsByRefPointer,
                        IsConst = argType.IsConst,
                        TargetTypeInfo = argType
                    };

                    PrimaryExpressions[context] = new ExpressionInfo { Text = $"@new<{typeName}>()", Type = argType };
                }
                else
                {
                    TypeInfo argType = expressions?[0].Type;

                    if (argType is not PointerTypeInfo)
                    {
                        argType = expressions?[0].Type.Clone() ?? TypeInfo.VarType.Clone();

                        SQLText = getTextSQLError("new operator not supported in sql translation", context);

                        argType = new PointerTypeInfo
                        {
                            Name = $"ptr<{argType.Name}>",
                            TypeName = $"ptr<{argType.Name}>",
                            FullTypeName = $"go.ptr<{argType.FullTypeName}>",
                            TypeClass = argType.TypeClass,
                            IsDerefPointer = argType.IsDerefPointer,
                            IsByRefPointer = argType.IsByRefPointer,
                            IsConst = argType.IsConst,
                            TargetTypeInfo = argType
                        };
                    }

                    PrimaryExpressions[context] = new ExpressionInfo 
                    { 
                        Text = $"@new<{typeInfo.Name}>({argumentList})", 
                        Type = argType,
                        LastToken = typeInfo.Name,
                        };
                }
            }
            else if (primaryExpression.Text == "make" && typeInfo is not null)
            {
                SQLText = getTextSQLError("make operator not supported in sql translation", context);

                PrimaryExpressions[context] = typeInfo.TypeClass switch
                {
                    TypeClass.Slice => new ExpressionInfo
                    {
                        Text = $"make_slice<{Common.RemoveSurrounding(typeInfo.TypeName, "slice<", ">")}>({argumentList})",
                        Type = ArrayTypeInfo.new_ArrayTypeInfo( primaryExpression.Type, null ), 
                    },
                    TypeClass.Map => new ExpressionInfo
                    {
                        Text = $"make_map<{Common.RemoveSurrounding(typeInfo.TypeName, "map<", ">")}>({argumentList})",
                        Type = primaryExpression.Type
                    },
                    TypeClass.Channel => new ExpressionInfo
                    {
                        Text = $"make_channel<{Common.RemoveSurrounding(typeInfo.TypeName, "channel<", ">")}>({argumentList})",
                        Type = primaryExpression.Type
                    },
                    _ => new ExpressionInfo
                    {
                        Text = $"{primaryExpression}<{typeInfo.TypeName}>({argumentList})",
                        Type = primaryExpression.Type
                    }
                };
            }
            else
            {
                var (sqlMethod, funcInfo1) = findSqlNativeMethod( funcName, packageName, normalizedType1 );
                if( sqlMethod != null && funcInfo1 != null ) 
                {
                    if( functionInfo == null )
                        functionInfo = funcInfo1;

                    if( parameters == null)
                        parameters = funcInfo1.Signature.Signature.Parameters;

                    SQLText = getTextSQL( sqlMethod );
                    if( this.m_LambdaCode != null ) 
                    {
                        Utils.Nop();
                    }
                    if( arguments.Count != parameters.Length)
                    {
                        var iParamLast = parameters.Length-1;
                        if( parameters.Length > 0 
                         && parameters.Length < arguments.Count
                         && parameters[ iParamLast ].IsVariadic )
                        {
                            for( var i = 0; i < iParamLast; i ++)
                            {
                                var arg = arguments[i];
                                if( parameters[i].Type.IsBool())
                                {
                                    arguments_SQL[i] = convertToBool( arguments_SQL[i], arguments_OperandKind[i] );
                                }

                                if( parameters[i].IsVariadic)
                                {
                                    SQLText = SQLText.Replace( $"@xx{(i+1)}@", arguments_SQL[i] );
                                }else
                                {
                                    SQLText = SQLText.Replace( $"@x{(i+1)}@", arguments_SQL[i] );
                                }
                            }
                            var argsSql = new List<string>();
                            for( var i = iParamLast; i < arguments.Count; i ++)
                            {
                                var arg = arguments[i];
                                if( parameters[ iParamLast ].Type.IsBool())
                                {
                                    argsSql.Add( convertToBool( arguments_SQL[i], arguments_OperandKind[i] ) );
                                }else
                                {
                                    argsSql.Add( arguments_SQL[i] );
                                }
                            }
                            var txtArgsSql = String.Join( ", ", argsSql.ToArray() );
                            SQLText = SQLText.Replace( $"@xx{(iParamLast+1)}@", txtArgsSql );
                        }else
                        { 
                            Log_Error( context, "paramanters not match with arguments"); 
                        }
                    }else
                    {
                        if( normalizedType1 != null )
                        {
                            SQLText = SQLText.Replace( $"@THIS@", primaryExpression.SQLText );
                        }

                        if( funcName == OrmDef.myTTime_Format
                         && normalizedType1.Name == OrmDef.myTTime
                         && normalizedType1.PackageName == Options.ConvertSql.OrmDir_AtomicsqlFunc_Git  )
                        {
                            SQLText = Options.ConvertSql.SqlDialect.convertFormatTime_GolangToSql(SQLText);
                        }else
                        {
                            for( var i = 0; i < arguments.Count; i ++)
                            {
                                var arg = arguments[i];
                                if( parameters[i].Type.IsBool())
                                {
                                    arguments_SQL[i] = convertToBool( arguments_SQL[i], arguments_OperandKind[i] );
                                }

                                if( parameters[i].IsVariadic)
                                {
                                    SQLText = SQLText.Replace( $"@xx{(i+1)}@", arguments_SQL[i] );
                                }else
                                {
                                    SQLText = SQLText.Replace( $"@x{(i+1)}@", arguments_SQL[i] );
                                }
                            }
                        }
                    }
                }else
                {
                    string argumentSQLList = string.Join(", ", arguments_SQL);

                    SQLText = getTextSQL( new TTextSql
                            { defaultSql = $"CALL {primaryExpression.SQLText}({argumentSQLList})" 
                            } );
                }
                if (typeInfo is not null)
                    argumentList = $"typeof({typeInfo.TypeName}){(string.IsNullOrEmpty(argumentList) ? string.Empty : $", {argumentList}")}";

                
                if( functionInfo != null)
                {
                    var type  = TypeInfo.setPackage( functionInfo.Signature.Signature.Result[0].Type, packageName, functionInfo.getMetaFile() );
                    var types = TypeInfo.setPackages( functionInfo.Signature.Signature.Result, packageName, functionInfo.getMetaFile());

                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = primaryExpression.LastToken,
                        Text = $"{primaryExpression}({argumentList})",
                        Type  = type,
                        Types = types,
                        SQLText = SQLText,
                        OperandKind = OperandKind,
                    };
                }else
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = primaryExpression.LastToken,
                        Text = $"{primaryExpression}({argumentList})",
                        Type = primaryExpression.Type,
                        SQLText = SQLText,
                        OperandKind = OperandKind,
                    };
                }
            }        
        }
    }
}
