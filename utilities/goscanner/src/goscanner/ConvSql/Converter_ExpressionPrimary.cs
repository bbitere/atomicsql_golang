using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime.Misc;
using Gemstone.Units;
using goscanner.Metadata;
using System.Diagnostics;
using Antlr4.Runtime;
using System.Xml;
using System.Linq.Expressions;

namespace goscanner.ConvSql;

public partial class SqlConvert
{
    public override void AfterFirstChildEveryRule(ParserRuleContext _context)
    { 
        var context = _context as GoParser.PrimaryExprContext;
        if( context != null )
        {
            //if( context.UID == 23011)
            //    Debugger.Break();

            PrimaryExpressions.TryGetValue(context.primaryExpr(), out ExpressionInfo primaryExpression);

        
            if (context.arguments() is not null)
            {
                // arguments
                //     : '('((expressionList | type(',' expressionList) ? ) '...' ? ',' ? ) ? ')'
                ExpressionCall_AfterFirstChild( context, primaryExpression);
            }
        }
    }

    public override void ExitPrimaryExpr(GoParser.PrimaryExprContext context)
    {
        string packageImport = $"{PackageImport.Replace('/', '.')}";

        // primaryExpr
        //     : operand
        //     | conversion
        //     | primaryExpr selector
        //     | primaryExpr index
        //     | primaryExpr slice
        //     | primaryExpr typeAssertion
        //     | primaryExpr arguments

        PrimaryExpressions.TryGetValue(context.primaryExpr(), out ExpressionInfo primaryExpression);

        //if (!(primaryExpression is null) && !string.IsNullOrEmpty(primaryExpression.Text))
        //    primaryExpression.Text = SanitizedIdentifier(primaryExpression.Text);

        
        if (Operands.TryGetValue(context.operand(), out ExpressionInfo operand))
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = operand.LastToken,
                Text = operand.Text,
                Type = operand.Type,
                Types = operand.Types,
                SQLText = operand.SQLText,
                OperandKind = operand.OperandKind,
                //SubExpressions = SubExpressions,
            };
        }
        else if (context.conversion() is not null)
        {
            // conversion
            //     : type '(' expression ',' ? ')'

            if (Types.TryGetValue(context.conversion().type_(), out TypeInfo typeInfo) 
                && Expressions.TryGetValue(context.conversion().expression(), out ExpressionInfo expression))
            {
                // TODO: Complex pointer expression may need special handling - could opt for unsafe implementation
                //if (typeInfo.TypeName.StartsWith("*(*"))
                //{
                //    PrimaryExpressions[context] = new ExpressionInfo
                //    {
                //        Text = $"{expression}.Value",
                //        Type = typeInfo
                //    };
                //}
                //else
                    
                if (typeInfo is PointerTypeInfo)
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = typeInfo.TypeName,
                        Text = $"new ptr<{typeInfo.TypeName}>({expression})",
                        Type = typeInfo,
                        SQLText = expression.SQLText,
                        OperandKind = expression.OperandKind,
                    };
                }
                else if (typeInfo.IsDerefPointer)
                {
                    string functionName = typeInfo.TypeName;
                    FunctionInfo functionInfo = null;
                    Metadata?.Functions.TryGetValue($"{functionName}()", out functionInfo);

                    if (functionInfo is not null)
                    {
                        typeInfo = functionInfo.Signature.Signature.Result[0].Type;

                        if (typeInfo is PointerTypeInfo pointerTypeInfo)
                            typeInfo = pointerTypeInfo.TargetTypeInfo;
                    }
                    else
                    {
                        typeInfo = TypeInfo.ObjectType;
                    }

                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = functionName,
                        Text = $"{functionName}({expression}).val",
                        Type = typeInfo,
                        SQLText = expression.SQLText,
                        OperandKind = operand.OperandKind,
                    };
                }
                else if (typeInfo.TypeClass == TypeClass.Struct)
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = typeInfo.TypeName,
                        Text = $"{typeInfo.TypeName}_cast({expression})",
                        Type = typeInfo,
                        SQLText = getTextSQLError("cast to struct not supported in sql translation", context),
                    };
                }
                else
                {
                    var typeConv1 = Options.ConvertSql.SqlDialect.convertGolangTypeToSqlType( typeInfo.TypeName, context, this);
                    var typeConv2 = Options.ConvertSql.SqlDialect.convertGolangTypeToSqlType( expression.Type.TypeName, context, this);
                    var SQLText =  expression.SQLText;
                    var OperandKind = expression.OperandKind;
                    if( typeConv1 != typeConv2)
                    {
                        SQLText =  $"CAST({expression.SQLText} AS {_quoteFld(typeConv1)})";
                        OperandKind = EOperandKind.Operator;
                    }

                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = typeInfo.TypeName,
                        Text = $"({typeInfo.TypeName}){expression}",
                        Type = typeInfo,
                        SQLText =  SQLText,
                        OperandKind = OperandKind,
                    };
                }
            }
            else
            {
                AddWarning(context, $"Failed to find type or sub-expression for the conversion expression in \"{context.GetText()}\"");
            }
        }
        else if (context.DOT() is not null)
        {
            // selector
            //     : '.' IDENTIFIER

            var identif1    = Common.SanitizedIdentifier( primaryExpression?.Text ??string.Empty);
            var identifSQL1 = Common.SanitizedIdentifier( primaryExpression?.SQLText ??string.Empty);
            var field1      = Common.SanitizedIdentifier( context.IDENTIFIER().GetText());

            TypeInfo typeInfo = null;

            // TODO: Will need to lookup IDENTIFIER type in metadata to determine type
            if (primaryExpression is null || primaryExpression.Type.FullTypeName == "System.Object")
                typeInfo = TypeInfo.VarType;

            typeInfo = typeInfo ?? primaryExpression?.Type;

            //FileMetadata structFileMetadata;
            StructInfo structInfo;

            var genericType = typeInfo.getNormalizedType();
            if( genericType == null )
            {
                if( searchFunctionInImports( field1, new TPackageInfo( identif1), out FunctionInfo functionInfo))
                {
                    //if( typeInfo != null && functionInfo != null )
                    
                    //var type  = TypeInfo.setPackage( functionInfo.Signature.Signature.Result[0].Type, "", functionInfo.getMetaFile() );
                    //var types = TypeInfo.setPackages( functionInfo.Signature.Signature.Result, "", functionInfo.getMetaFile()));

                    var expr1 = new ExpressionInfo
                    {
                        LastToken = field1,
                        Text = $"{identif1}.{field1}",
                        SQLText = identifSQL1,
                        OperandKind = EOperandKind.Simple,
                        Type = typeInfo ?? primaryExpression?.Type,
                        Types = typeInfo != null ? new_ParameterInfo_TypeInfo(typeInfo ) : primaryExpression?.Types
                    };
                    //expr1.Type.PackageInfo_setName(identif1);
                    PrimaryExpressions[context] = expr1;
                    return;
                }
            }

            var sqlField = "";

            if( genericType != null && this.Metadata != null && this.Metadata.getFolderMetadata() != null )
            {
                if( searchStructInImports( genericType.Name, genericType.PackageInfo,out structInfo ) )
                {
                    var field = structInfo.Fields.Where(x=>x.Name == field1 ).FirstOrDefault();
                    if( field != null ) 
                    {   
                        sqlField = getSqlTextForField( field, structInfo, ref typeInfo, true );

                        var expr2 = new ExpressionInfo
                        {
                            LastToken = field1,
                            Text = $"{identif1}.{field1}",
                            SQLText = getTextSQLIdentif( $"{identif1}.{field1}", identifSQL1, sqlField, field.Type, field.Name, context ),
                            OperandKind = EOperandKind.Simple,
                            Type = field.Type,
                            Types = new_ParameterInfo_TypeInfo(field.Type ),
                        };
                        PrimaryExpressions[context] = expr2;
                        //expr2.Type.PackageInfo_setName(identif1);
                        return;
                    }
                }

                if( searchFunctionInImports( genericType.Name, genericType.PackageInfo, out FunctionInfo funct) )
                {
                    //var type  = TypeInfo.setPackage( functionInfo.Signature.Signature.Result[0].Type, "", functionInfo.getMetaFile() );
                    //var types = TypeInfo.setPackages( functionInfo.Signature.Signature.Result, "", functionInfo.getMetaFile()));

                    var expr3 = new ExpressionInfo
                    {
                        LastToken = field1,
                        Text = $"{identif1}.{field1}",
                        SQLText = identifSQL1,
                        OperandKind = EOperandKind.Simple,
                        Type = typeInfo ?? primaryExpression?.Type,
                        Types = typeInfo != null ? new_ParameterInfo_TypeInfo(typeInfo ) : primaryExpression?.Types
                    };
                    PrimaryExpressions[context] = expr3;
                    //expr3.Type.PackageInfo_setName(identif1);
                    return;
                }
            }

            if( genericType != null )
            {
                if( genericType.PackageName == "sql"
                 &&( false
                 || genericType.Name == OrmDef.Sql_NullTime
                 || genericType.Name == OrmDef.Sql_NullInt32
                 || genericType.Name == OrmDef.Sql_NullInt16
                 || genericType.Name == OrmDef.Sql_NullByte 
                 || genericType.Name == OrmDef.Sql_NullInt64 
                 || genericType.Name == OrmDef.Sql_NullFloat64
                 || genericType.Name == OrmDef.Sql_NullString
                 || genericType.Name == OrmDef.Sql_NullBool ) )
                {
                    generatePrimaryExpression_SqlNull( identif1, identifSQL1, field1, context );
                    return;
                }

                if( genericType.PackageName == "time"    //time.Time
                 && genericType.Name == "Time")
                {
                    if( field1 == OrmDef.myTTime_Day 
                     || field1 == OrmDef.myTTime_Year
                     || field1 == OrmDef.myTTime_Month
                     || field1 == OrmDef.myTTime_YearDay
                     || field1 == OrmDef.myTTime_Hour        
                     || field1 == OrmDef.myTTime_Minute      
                     || field1 == OrmDef.myTTime_Second      
                     || field1 == OrmDef.myTTime_Nanosecond  
                     || field1 == OrmDef.myTTime_UnixMilli   
                     || field1 == OrmDef.myTTime_Compare     
                     || field1 == OrmDef.myTTime_Format      
                     )
                    { 
                        var expr4 = new ExpressionInfo
                        {
                            LastToken = field1,
                            Text = $"{identif1}.{field1}",
                            SQLText = $"{identifSQL1}",
                            OperandKind = EOperandKind.Simple,
                            Type  = OrmDef.myTTimeType,
                            Types = new_ParameterInfo_TypeInfo( OrmDef.myTTimeType ),
                        };
                        //expr4.Type.PackageInfo_setName(identif1);
                        PrimaryExpressions[context] = expr4;
                        return;
                    }
                }
            }
            /*
            if( structInfo != null )
            {
                if( searchFunctionWithThisParamInImports( structInfo, field1, genericType.PackageInfo, out  FunctionInfo funct  ) )
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = field1,
                        Text = $"{identif1}.{field1}",
                        SQLText = identifSQL1,
                        OperandKind = EOperandKind.Simple,
                        Type = typeInfo ?? primaryExpression?.Type,
                        Types = typeInfo != null ? new_ParameterInfo_TypeInfo(typeInfo ) : primaryExpression?.Types
                    };
                }
            }
            */

            string selectionExpression = $"{identif1}.{field1}";
            if( sqlField == "" )
            {   
                if( genericType != null && genericType.Name !="")
                { 
                    sqlField = getTextSQLError($"type {genericType.Name}: not found. Check the $ GitProject settings in sql-config.cfg", context);
                }else
                {
                    sqlField = getTextSQLError("incorect declaration : not supported in sql translation", context);
                }
            }

            var expr5 = new ExpressionInfo
            {
                Text = selectionExpression,
                LastToken = field1,
                SQLText = getTextSQLIdentif( selectionExpression, identifSQL1, sqlField, typeInfo ?? primaryExpression?.Type, field1, context ),
                OperandKind = EOperandKind.Simple,
                Type = typeInfo ?? primaryExpression?.Type,
                Types = typeInfo != null ? new_ParameterInfo_TypeInfo(typeInfo ) : primaryExpression?.Types
            };
            //expr5.Type.PackageInfo_setName(identif1);
            PrimaryExpressions[context] = expr5;
        }
        else if (context.index() is not null)
        {
            // index
            //     : '[' expression ']'
            // TODO: Will need to lookup IDENTIFIER type in metadata to determine type
            if (Expressions.TryGetValue(context.index().expression(), out ExpressionInfo expression))
            {
                PrimaryExpressions[context] = new ExpressionInfo
                {
                    LastToken = primaryExpression.LastToken,
                    Text = $"{primaryExpression}[{expression}]",
                    SQLText = getTextSQLError("array index not supported in sql translation", context),
                    Type = primaryExpression.Type,
                    Types = primaryExpression.Types
                };
            }
            else
            {
                AddWarning(context, $"Failed to find index expression for \"{context.GetText()}\"");
            }
        }
        else if (context.slice_() is not null)
        {
            // slice
            //     : '['((expression ? ':' expression ? ) | (expression ? ':' expression ':' expression)) ']'

            GoParser.Slice_Context sliceContext = context.slice_();

            if (sliceContext.children.Count == 3)
            {
                // primaryExpr[:]
                PrimaryExpressions[context] = new ExpressionInfo
                {
                    LastToken = primaryExpression.LastToken,
                    Text = $"{primaryExpression}[..]",
                    SQLText = getTextSQLError("slice not supported in sql translation", context),
                    Type = primaryExpression.Type,
                    Types = primaryExpression.Types
                };
            }
            else if (sliceContext.children.Count == 4)
            {
                bool expressionIsLeft = sliceContext.children[1] is GoParser.ExpressionContext;

                // primaryExpr[low:] or primaryExpr[:high]
                if (Expressions.TryGetValue(sliceContext.expression(0), out ExpressionInfo expression))
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = primaryExpression.LastToken,
                        Text = $"{primaryExpression}[{(expressionIsLeft
                        ? $"{(expression.Type?.TypeName == "int" ? string.Empty : "(int)")}{expression}.."
                        : $"..{(expression.Type?.TypeName == "int" ? string.Empty : "(int)")}{expression}")}]",
                        SQLText = getTextSQLError("slice not supported in sql translation", context),
                        Type = primaryExpression.Type,
                        Types = primaryExpression.Types
                    };
                }
                else
                {
                    AddWarning(context, $"Failed to find slice expression for \"{context.GetText()}\"");
                }
            }
            else if (sliceContext.children.Count == 5)
            {
                if (sliceContext.children[1] is GoParser.ExpressionContext && sliceContext.children[3] is GoParser.ExpressionContext)
                {
                    // primaryExpr[low:high]
                    if (Expressions.TryGetValue(sliceContext.expression(0), out ExpressionInfo lowExpression)
                        && Expressions.TryGetValue(sliceContext.expression(1), out ExpressionInfo highExpression))
                    {
                        PrimaryExpressions[context] = new ExpressionInfo
                        {
                            LastToken = primaryExpression.LastToken,
                            Text = $"{primaryExpression}[{(lowExpression.Type?.TypeName == "int" ? string.Empty : "(int)")}{lowExpression}..{(highExpression.Type?.TypeName == "int" ? string.Empty : "(int)")}{highExpression}]",
                            SQLText = getTextSQLError("slice not supported in sql translation", context),
                            Type = primaryExpression.Type,
                            Types = primaryExpression.Types
                        };
                    }
                    else
                    {
                        AddWarning(context, $"Failed to find one of the slice expressions for \"{context.GetText()}\"");
                    }
                }
                else
                {
                    AddWarning(context, $"Failed to find slice expression for \"{context.GetText()}\"");
                }
            }
            else if (sliceContext.children.Count == 6)
            {
                // primaryExpr[:high:max]
                if (Expressions.TryGetValue(sliceContext.expression(0), out ExpressionInfo highExpression)
                    && Expressions.TryGetValue(sliceContext.expression(1), out ExpressionInfo maxExpression))
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = primaryExpression.LastToken,
                        Text = $"{primaryExpression}.slice(-1, {highExpression}, {maxExpression})",
                        SQLText = getTextSQLError("slice not supported in sql translation", context),
                        Type = primaryExpression.Type,
                        Types = primaryExpression.Types
                    };
                }
                else
                {
                    AddWarning(context, $"Failed to find one of the slice expressions for \"{context.GetText()}\"");
                }
            }
            else if (sliceContext.children.Count == 7)
            {
                // primaryExpr[low:high:max]
                if (Expressions.TryGetValue(sliceContext.expression(0), out ExpressionInfo lowExpression) 
                    && Expressions.TryGetValue(sliceContext.expression(1), out ExpressionInfo highExpression)
                    && Expressions.TryGetValue(sliceContext.expression(2), out ExpressionInfo maxExpression))
                {
                    PrimaryExpressions[context] = new ExpressionInfo
                    {
                        Text = $"{primaryExpression}.slice({lowExpression}, {highExpression}, {maxExpression})",
                        SQLText = getTextSQLError("slice not supported in sql translation", context),
                        Type = primaryExpression.Type,
                        Types = primaryExpression.Types,
                        
                    };
                }
                else
                {
                    AddWarning(context, $"Failed to find one of the slice expressions for \"{context.GetText()}\"");
                }
            }
        }
        else if (context.typeAssertion() is not null)
        {
            // typeAssertion
            //     : '.' '(' type ')'

            if (Types.TryGetValue(context.typeAssertion().type_(), out TypeInfo typeInfo))
            {
                PrimaryExpressions[context] = new ExpressionInfo
                {
                    LastToken = primaryExpression.LastToken,
                    Text = $"{primaryExpression}._<{typeInfo.TypeName}>()",
                    Type = typeInfo,
                    SQLText = primaryExpression.SQLText,
                    OperandKind = primaryExpression.OperandKind,
                };
            }
            else
            {
                AddWarning(context, $"Failed to find type for the type assertion expression in \"{context.GetText()}\"");
            }
        }
        //---------------------------------------------------------------------------------------------------

        //---------------------------------------------------------------------------------------------------
        else if (context.arguments() is not null)
        {
            ExpressionCall( context, primaryExpression, packageImport );
            
        }
        else
        {
            AddWarning(context, $"Unexpected primary expression \"{context.GetText()}\"");
        }
    }

    Signature findMethodInPackage( FunctionSignature Signature )
    {
        return Signature.Signature;
    }

    void generatePrimaryExpression_SqlNull(string identif1, string identifSQL1, string field1, GoParser.PrimaryExprContext context )
    {
        if( field1 == OrmDef.Sql_Null_Valid)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"({identifSQL1} IS NOT NULL)",
                OperandKind = EOperandKind.Operator,
                Type  = TypeInfo.BoolType,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.BoolType ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullBool_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.BoolType,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.BoolType ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullByte_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.ByteType,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.ByteType ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullInt32_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.Int32Type,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.Int32Type ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullInt16_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.Int16Type,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.Int16Type ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullInt64_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.Int64Type,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.Int64Type ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullFloat64_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.Float64Type,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.Float64Type ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullString_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.StringType,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.StringType ),
            };
            return;
        }else
        if( field1 == OrmDef.Sql_NullTime_Val)
        {
            PrimaryExpressions[context] = new ExpressionInfo
            {
                LastToken = field1,
                Text = $"{identif1}.{field1}",
                SQLText = $"{identifSQL1}",
                OperandKind = EOperandKind.Simple,
                Type  = TypeInfo.TimeType,
                Types = new_ParameterInfo_TypeInfo( TypeInfo.TimeType ),
            };
            return;
        }
    }

        
}

