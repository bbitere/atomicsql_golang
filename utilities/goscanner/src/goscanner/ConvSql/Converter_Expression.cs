
//******************************************************************************************************
//  ScannerBase_Expression.cs - Gbtc
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

using Antlr4.Runtime.Misc;
using Gemstone.Units;
using goscanner.ConvCommon;
using goscanner.Metadata;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Linq.Expressions;
using System.Xml;
using static goscanner.Common;


namespace goscanner.ConvSql;

public partial class SqlConvert
{
    // Stack handlers:
    //  expressionList (required)
    //  expressionStmt (required)
    //  sendStmt (required)
    //  incDecStmt (required)
    //  defer (required)
    //  ifStmt (required)
    //  exprSwitchStmt (optional)
    //  recvStmt (required)
    //  forClause (optional)
    //  rangeClause (required)
    //  goStmt (required)
    //  arrayLength (required)
    //  operand (optional)
    //  key (optional)
    //  element (optional)
    //  index (required)
    //  slice (optional)
    //  expression (optional)
    //  conversion (required)
    protected readonly new ParseTreeValues<ExpressionInfo> Expressions = new();
    protected readonly new ParseTreeValues<ExpressionInfo> UnaryExpressions = new();
    protected readonly new ParseTreeValues<ExpressionInfo> PrimaryExpressions = new();
    protected readonly new ParseTreeValues<ExpressionInfo> Operands = new();

    private static readonly HashSet<string> s_comparisionOperands = new()
    {
        "==", "!=", "<", "<=", ">=", "&&", "||"
    };

    public override void ExitExpression(GoParser.ExpressionContext context)
    {
        // expression
        //     : primaryExpr
        //     | unaryExpr
        //     | expression('*' | '/' | '%' | '<<' | '>>' | '&' | '&^') expression
        //     | expression('+' | '-' | '|' | '^') expression
        //     | expression('==' | '!=' | '<' | '<=' | '>' | '>=') expression
        //     | expression '&&' expression
        //     | expression '||' expression

        if (context.expression()?.Length == 2)
        {
            string leftOperand = Expressions.TryGetValue(context.expression(0), out ExpressionInfo leftOperandExpression) ? leftOperandExpression.Text : context.expression(0).GetText();
            string rightOperand = Expressions.TryGetValue(context.expression(1), out ExpressionInfo rightOperandExpression) ? rightOperandExpression.Text : context.expression(1).GetText();
            string binaryOP = context.children[1].GetText();

            if (binaryOP.Equals("<<") || binaryOP.Equals(">>"))
            {
                if (!int.TryParse(rightOperand, out int _))
                    rightOperand = $"(int)({rightOperand})";
            }

            binaryOP = binaryOP.Equals("&^") ? " & ~" : $" {binaryOP} ";

            string expression = $"{leftOperand}{binaryOP}{rightOperand}";
            var OperandKind = EOperandKind.Operator;

            var _binaryOP = binaryOP.Trim();
            if (s_comparisionOperands.Contains(binaryOP.Trim() ))
            {
                OperandKind = EOperandKind.Operator;
                //"==", "!=", "<", "<=", ">=", "&&", "||"
                TNoSqlCode nosqlCode = leftOperandExpression.NoSQLCode;
                var sqlText =  leftOperandExpression.SQLText;

                if( _binaryOP == "&&")
                {
                    nosqlCode = new TNoSqlCode("&&", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText} AND {rightOperandExpression.SQLText}";
                } else
                if( _binaryOP == "||")
                {
                    nosqlCode = new TNoSqlCode("||", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText} OR {rightOperandExpression.SQLText}";
                } else
                if( _binaryOP == "==")
                {
                    nosqlCode = new TNoSqlCode("==", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    if( leftOperandExpression.SQLText == sql_NULL &&  rightOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{sql_NULL} = {sql_NULL}";
                    }else
                    if( leftOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{rightOperandExpression.SQLText} IS {sql_NULL}";
                    }else
                    if( rightOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{leftOperandExpression.SQLText} IS {sql_NULL}";
                    }else
                    {
                        sqlText =  $"{leftOperandExpression.SQLText}={rightOperandExpression.SQLText}";
                    }
                }
                else
                if( _binaryOP == "!=")
                {
                    nosqlCode = new TNoSqlCode("!=", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    if( leftOperandExpression.SQLText == sql_NULL &&  rightOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{sql_NULL} <> {sql_NULL}";
                    }else
                    if( leftOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{rightOperandExpression.SQLText} IS NOT {sql_NULL}";
                    }else
                    if( rightOperandExpression.SQLText == sql_NULL)
                    {
                        sqlText =  $"{leftOperandExpression.SQLText} IS NOT {sql_NULL}";
                    }else
                    {
                        sqlText =  $"{leftOperandExpression.SQLText}<>{rightOperandExpression.SQLText}" ;
                    }
                }
                else
                if( _binaryOP == "<")
                {
                    nosqlCode = new TNoSqlCode("<", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText}<{rightOperandExpression.SQLText}";
                } else     
                if( _binaryOP == "<=")
                {
                    nosqlCode = new TNoSqlCode("<=", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText}<={rightOperandExpression.SQLText}";
                } else   
                if( _binaryOP == ">")
                {
                    nosqlCode = new TNoSqlCode(">", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText}>{rightOperandExpression.SQLText}";
                } else  
                if( _binaryOP == ">=")
                {
                    nosqlCode = new TNoSqlCode(">=", leftOperandExpression.NoSQLCode, rightOperandExpression.NoSQLCode);
                    sqlText =  $"{leftOperandExpression.SQLText}>={rightOperandExpression.SQLText}";
                }
                else
                {
                    sqlText = getTextSQLError($"operator {binaryOP} not supproted in sql translation ", context);
                }

                Expressions[context] = new ExpressionInfo
                {
                    LastToken = "",
                    Text = expression,
                    SQLText = sqlText,
                    NoSQLCode = nosqlCode,
                    bIsNoSql = leftOperandExpression!=null?leftOperandExpression.bIsNoSql:false,
                    OperandKind = OperandKind,
                    Type = new TypeInfo
                    {
                        Name = "bool",
                        TypeName = "bool",
                        FullTypeName = "System.Boolean",
                        TypeClass = TypeClass.Simple,
                        IsConst = true
                    }
                };
            }
            else
            {
                OperandKind = EOperandKind.Operator;
                var sqlText =  leftOperandExpression.SQLText;
                var binaryOperator = binaryOP.Trim();

                getTextNoSQLError($"operator {binaryOperator} not supproted in nosql translation ", context);

                if( binaryOperator == "+")
                    sqlText =  $"{leftOperandExpression.SQLText}+{rightOperandExpression.SQLText}";
                else
                if( binaryOperator == "-")
                    sqlText =  $"{leftOperandExpression.SQLText}-{rightOperandExpression.SQLText}";
                else
                if( binaryOperator == "*")
                    sqlText =  $"{leftOperandExpression.SQLText}*{rightOperandExpression.SQLText}";
                else
                if( binaryOperator == "/")
                    sqlText =  $"{leftOperandExpression.SQLText}/{rightOperandExpression.SQLText}";
                else
                if( binaryOperator == ">")
                    sqlText =  $"{leftOperandExpression.SQLText} > {rightOperandExpression.SQLText}";
                else
                if( binaryOperator == ">=")
                    sqlText =  $"{leftOperandExpression.SQLText} >= {rightOperandExpression.SQLText}";
                else
                if( binaryOperator == "<")
                    sqlText =  $"{leftOperandExpression.SQLText} < {rightOperandExpression.SQLText}";
                else
                if( binaryOperator == "<=")
                    sqlText =  $"{leftOperandExpression.SQLText} <= {rightOperandExpression.SQLText}";
                else
                    sqlText = getTextSQLError($"operator {binaryOperator} not supproted in sql translation ", context);

                // TODO: If both operands are integer, expression should be treated as arbitrary-precision numbers until assigned to a variable
                Expressions[context] = new ExpressionInfo
                {
                    LastToken = "",
                    Text = expression,
                    SQLText = sqlText,
                    OperandKind = OperandKind,
                    Type = leftOperandExpression.Type ?? TypeInfo.VarType,
                    
                };
            }
        }
        else
        {
            if (context.primaryExpr() is not null)
            {
                if (PrimaryExpressions.TryGetValue(context.primaryExpr(), out ExpressionInfo primaryExpression))
                    Expressions[context] = primaryExpression;
                else
                    AddWarning(context, $"Failed to find primary expression \"{context.primaryExpr().GetText()}\" in the expression \"{context.GetText()}\"");
            }
            else if (context.unaryExpr() is not null)
            {
                if (UnaryExpressions.TryGetValue(context.unaryExpr(), out ExpressionInfo unaryExpression))
                    Expressions[context] = unaryExpression;
                else
                    AddWarning(context, $"Failed to find unary expression \"{context.unaryExpr().GetText()}\" in the expression \"{context.GetText()}\"");
            }
            else
            {
                AddWarning(context, $"Unexpected expression \"{context.GetText()}\"");
            }
        }
    }

    /*
    public override void ExitUnaryExpr(GoParser.UnaryExprContext context)
    {
        // unaryExpr
        //     : primaryExpr
        //     | ('+' | '-' | '!' | '^' | '*' | '&' | '<-') expression

        if (PrimaryExpressions.TryGetValue(context.primaryExpr(), out ExpressionInfo primaryExpression))
        {
            UnaryExpressions[context] = primaryExpression;
        }
        else if (context.expression() is not null)
        {
            if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
            {
                string unaryOP = context.children[0].GetText();
                string unaryExpression = string.Empty;
                TypeInfo expressionType = expression.Type;

                if (unaryOP.Equals("^", StringComparison.Ordinal))
                {
                    unaryOP = "~";
                }
                else if (unaryOP.Equals("<-", StringComparison.Ordinal))
                {
                    // TODO: Handle channel value access (update when channel class is created):
                    unaryOP = null;
                    unaryExpression = $"{expression}.Receive()";
                }
                else if (unaryOP.Equals("&", StringComparison.Ordinal))
                {
                    unaryOP = null;

                    if (expression.Text.StartsWith("new ", StringComparison.Ordinal))
                    {
                        unaryExpression = $"addr({expression})";
                        expressionType = new PointerTypeInfo
                        {
                            Name = $"ptr<{expressionType.Name}>",
                            TypeName = $"ptr<{expressionType.Name}>",
                            FullTypeName = $"go.ptr<{expressionType.FullTypeName}>",
                            TypeClass = expressionType.TypeClass,
                            IsDerefPointer = expressionType.IsDerefPointer,
                            IsByRefPointer = expressionType.IsByRefPointer,
                            IsConst = expressionType.IsConst,
                            TargetTypeInfo = expressionType
                        };
                    }
                    else
                    {
                        unaryExpression = $"{AddressPrefix}{expression}";
                    }
                }
                else if (unaryOP.Equals("*", StringComparison.Ordinal))
                {
                    unaryOP = null;

                    if (!expression.Text.EndsWith(".val"))
                        unaryExpression = $"{expression}.val";
                    else
                        unaryExpression = expression.Text;
                }

                if (unaryOP is not null)
                    unaryExpression = $"{unaryOP}{expression}";

                UnaryExpressions[context] = new ExpressionInfo
                {
                    Text = unaryExpression,
                    Type = expressionType
                };
            }
            else
            {
                AddWarning(context, $"Unexpected unary expression \"{context.expression().GetText()}\"");
            }
        }
        else if (!UnaryExpressions.ContainsKey(context))
        {
            AddWarning(context, $"Unexpected unary expression \"{context.GetText()}\"");
        }
    }
    */
    public override void ExitUnaryExpr(GoParser.UnaryExprContext context)
    {
        if (EMITTING_CODE)
        { 
            string unaryOP = context.children[0].GetText();
            var sqlText =  "";
            TNoSqlCode noSQLCode = null;
            var bIsNoSql = false;
            ExpressionInfo expression = null;
            var OperandKind = EOperandKind.Simple;

            if (context.expression() is not null)
            {
                if (Expressions.TryGetValue(context.expression(), out expression))
                {
                     OperandKind = expression.OperandKind;
                    string unaryExpression = string.Empty;
                    TypeInfo expressionType = expression.Type;

                    noSQLCode = expression.NoSQLCode;
                    bIsNoSql  = expression.bIsNoSql;
                    sqlText   = expression.SQLText;
                    if( unaryOP == "!")
                    {
                        noSQLCode = new TNoSqlCode( "!", NSqlConvertToBool( expression.NoSQLCode, expression.OperandKind ));

                        expression.SQLText = convertToBool( expression.SQLText, expression.OperandKind );
                        OperandKind = EOperandKind.Operator;
                        sqlText =  $"NOT ({expression.SQLText})";
                    }
                    else
                    if( unaryOP == "-")
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        sqlText =  $"-({expression.SQLText})";
                        OperandKind = EOperandKind.Operator;
                    }
                    else
                    if (unaryOP.Equals("&", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        sqlText = expression.SQLText;
                    }else
                    if (unaryOP.Equals("*", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        sqlText = expression.SQLText;
                    }else
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        sqlText = getTextSQLError($"operator {unaryOP} not supproted in sql translation ", context);
                    }


                    if (unaryOP.Equals("^", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        unaryOP = "~";
                    }
                    else if (unaryOP.Equals("<-", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        // TODO: Handle channel value access (update when channel class is created):
                        unaryOP = null;
                        unaryExpression = $"{expression}.Receive()";
                    }
                    else if (unaryOP.Equals("&", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        unaryOP = null;

                        if (expression.Text.StartsWith("new ", StringComparison.Ordinal))
                        {
                            unaryExpression = $"addr({expression})";
                            expressionType = new PointerTypeInfo
                            {
                                Name = $"ptr<{expressionType.Name}>",
                                TypeName = $"ptr<{expressionType.Name}>",
                                FullTypeName = $"go.ptr<{expressionType.FullTypeName}>",
                                TypeClass = expressionType.TypeClass,
                                IsDerefPointer = expressionType.IsDerefPointer,
                                IsByRefPointer = expressionType.IsByRefPointer,
                                IsConst = expressionType.IsConst,
                                TargetTypeInfo = expressionType
                            };
                        }
                        else
                        {
                            unaryExpression = $"{AddressPrefix}{expression}";
                        }
                    }
                    else if (unaryOP.Equals("*", StringComparison.Ordinal))
                    {
                        getTextNoSQLError( $"operand {unaryOP} not supported in nosql translation", context);
                        unaryOP = null;

                        if (!expression.Text.EndsWith(".val"))
                            unaryExpression = $"{expression}.val";
                        else
                            unaryExpression = expression.Text;
                    }

                    if (unaryOP is not null)
                        unaryExpression = $"{unaryOP}{expression}";

                    UnaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = expression.LastToken,
                        Text = unaryExpression,
                        SQLText = sqlText,
                        NoSQLCode = noSQLCode,
                        bIsNoSql = expression.bIsNoSql,
                        OperandKind = OperandKind,
                        Type = expressionType
                    };
                }
                else
                {
                    AddWarning(context, $"Unexpected unary expression \"{context.expression().GetText()}\"");
                }
            }
            else if (!UnaryExpressions.ContainsKey(context))
            {
                AddWarning(context, $"Unexpected unary expression \"{context.GetText()}\"");
            }

            if( CurrentFunction != null )
            {
                ParameterInfo[] parameters = CurrentFunction.Signature.Signature.Parameters;
                //string unaryOP = context.children[0].GetText();

                if (unaryOP != null && !unaryOP.Equals("*", StringComparison.Ordinal))
                    return;

                ParameterInfo pointerParam = parameters.FirstOrDefault(parameter => parameter.Name.Equals(expression.Text));
                        
                if (pointerParam is not null && pointerParam.Type is PointerTypeInfo pointer)
                {
                    TypeInfo targetType = pointer.TargetTypeInfo.Clone();
                    targetType.IsByRefPointer = true;

                    string derefPointerExpression = expression.Text;

                    // Handle pointer-to-pointer dereferencing
                    int derefs = pointer.Name.Count(chr => chr == '*');

                    if (derefs > 1)
                    {
                        for (int i = 1; i < derefs; i++)
                            derefPointerExpression += ".val";
                    }

                    // Implicitly dereference pointer parameter when dereference operator (*) is used
                    UnaryExpressions[context] = new ExpressionInfo
                    {
                        LastToken = expression.LastToken,
                        Text = derefPointerExpression,
                        SQLText = sqlText,
                        NoSQLCode = noSQLCode,
                        OperandKind = OperandKind,
                        Type = targetType
                    };
                }
            }
        }
    }

    
    public override void ExitOperand(GoParser.OperandContext context)
    {
        // operand
        //     : literal
        //     | operandName
        //     | methodExpr
        //     | '(' expression ')'

        if (Expressions.TryGetValue(context.expression(), out ExpressionInfo expression))
        {
            Operands[context] = new ExpressionInfo
            {
                LastToken = expression.LastToken,
                Text = $"({expression})",
                Type = expression.Type,
                SQLText = $"({expression.SQLText})",
                NoSQLCode = expression.NoSQLCode,
                bIsNoSql = expression.bIsNoSql,
            };
            return;
        }
        if (Expressions.TryGetValue(context.operandName(), out expression))
        {
            Operands[context] = new ExpressionInfo
            {
                LastToken = expression.LastToken,
                Text = $"{expression}",
                Type = expression.Type,
                SQLText = $"{expression.SQLText}",
                NoSQLCode = expression.NoSQLCode,
                bIsNoSql = expression.bIsNoSql,
            };
            return;
        }
        if (Operands.TryGetValue(context, out expression))
        {
            Operands[context] = new ExpressionInfo
            {
                LastToken = expression.LastToken,
                Text = $"{expression}",
                Type = expression.Type,
                SQLText = $"{expression.SQLText}",
                NoSQLCode = expression.NoSQLCode,
                bIsNoSql = expression.bIsNoSql,
            };
            return;
        }

        // Remaining operands contexts handled below...
    }

    public override void ExitBasicLit(GoParser.BasicLitContext context)
    {
        // operand
        //     : literal
        //     | operandName
        //     | methodExpr
        //     | '(' expression ')'

        // literal
        //     : basicLit
        //     | compositeLit
        //     | functionLit

        if (context.Parent.Parent is not GoParser.OperandContext operandContext)
        {
            AddWarning(context, $"Could not derive parent operand context from basic literal: \"{context.GetText()}\"");
            return;
        }

        string basicLiteral;
        TypeInfo typeInfo;

        // basicLit
        //     : INT_LIT
        //     | FLOAT_LIT
        //     | IMAGINARY_LIT
        //     | RUNE_LIT
        //     | STRING_LIT

        var SQLText = "";
        TNoSqlCode NoSQLText = null;

        if (context.IMAGINARY_LIT() is not null)
        {
            SQLText = getTextSQLError("not supported in sql transaltion", context);

            string value = context.IMAGINARY_LIT().GetText();
            bool endsWith_i = value.EndsWith("i");
            value = endsWith_i ? value[..^1] : value;

            if (float.TryParse(value, out _))
            {
                basicLiteral = endsWith_i ? $"i({value}F)" : $"{value}F";

                typeInfo = new TypeInfo
                {
                    Name = "complex64",
                    TypeName = "complex64",
                    FullTypeName = "go.complex64",
                    TypeClass = TypeClass.Simple
                };
            }
            else
            {
                basicLiteral = endsWith_i ? $"i({value}D)" : $"{value}D";

                typeInfo = new TypeInfo
                {
                    Name = "Complex",
                    TypeName = "Complex",
                    FullTypeName = "System.Numerics.Complex",
                    TypeClass = TypeClass.Simple
                };
            }
        }
        else if (context.FLOAT_LIT() is not null)
        {
            basicLiteral = context.GetText();
            SQLText = basicLiteral;
            NoSQLText = new TNoSqlNumber( basicLiteral );

            if (float.TryParse(basicLiteral, out _))
            {
                basicLiteral += "F";

                typeInfo = new TypeInfo
                {
                    Name = "float",
                    TypeName = "float",
                    FullTypeName = "System.Single",
                    TypeClass = TypeClass.Simple,
                    IsConst = true
                };
            }
            else
            {
                basicLiteral += "D";

                typeInfo = new TypeInfo
                {
                    Name = "double",
                    TypeName = "double",
                    FullTypeName = "System.Double",
                    TypeClass = TypeClass.Simple,
                    IsConst = true
                };
            }
        }
        else if (context.integer() is not null)
        {
            NoSQLText = new TNoSqlNumber( ReplaceOctalBytes(context.integer().GetText()) );
            basicLiteral = ReplaceOctalBytes(context.integer().GetText());
            SQLText = basicLiteral;

            if (context.integer().RUNE_LIT() is not null)
            {
                typeInfo = new TypeInfo
                {
                    Name = "char",
                    TypeName = "char",
                    FullTypeName = "System.Char",
                    TypeClass = TypeClass.Simple,
                    IsConst = true
                };
            }
            else
            {
                if (nint.TryParse(basicLiteral, out nint val))
                {
                    if (val > int.MaxValue)
                        basicLiteral = $"(nint){basicLiteral}L";

                    typeInfo = new TypeInfo
                    {
                        Name = "int",
                        TypeName = "int",
                        FullTypeName = "int",
                        TypeClass = TypeClass.Simple,
                        IsConst = true
                    };
                }
                else
                {
                    if (nuint.TryParse(basicLiteral, out nuint uval) && uval > uint.MaxValue)
                        basicLiteral = $"(nuint){basicLiteral}UL";

                    typeInfo = new TypeInfo
                    {
                        Name = "nuint",
                        TypeName = "nuint",
                        FullTypeName = "nuint",
                        TypeClass = TypeClass.Simple,
                        IsConst = true
                    };
                }
            }
        }
        else if (context.RUNE_LIT() is not null)
        {
            getTextNoSQLError("not supported in sql transaltion", context);
            SQLText = getTextSQLError("not supported in sql transaltion", context);
            basicLiteral = ReplaceOctalBytes(context.RUNE_LIT().GetText());

            typeInfo = new TypeInfo
            {
                Name = "char",
                TypeName = "char",
                FullTypeName = "System.Char",
                TypeClass = TypeClass.Simple,
                IsConst = true
            };
        }
        else if (context.string_() is not null)
        {
            basicLiteral = $"{ToStringLiteral(ReplaceOctalBytes(context.string_().GetText()))}u8";
            SQLText = convertGolangStringToSqlString( context.string_().GetText() );
            NoSQLText = new TNoSqlString( context.string_().GetText()  );

            typeInfo = new TypeInfo
            {
                Name = "@string",
                TypeName = "@string",
                FullTypeName = "go.@string",
                TypeClass = TypeClass.Simple,
                IsConst = true
            };
        }
        else if (context.NIL_LIT() is not null)
        {
            SQLText = sql_NULL;
            NoSQLText = (new TNoSqlKeyword( "nil" ));
            basicLiteral = "null";
            typeInfo = TypeInfo.ObjectType;
        }
        else
        {
            AddWarning(context, $"Unexpected basic literal: \"{context.GetText()}\"");
            return;
        }

        Operands[operandContext] = new ExpressionInfo
        {
            LastToken = basicLiteral,
            Text = basicLiteral,
            SQLText = SQLText,
            NoSQLCode = NoSQLText,
            bIsNoSql = false,
            Type = typeInfo,
            OperandKind = EOperandKind.Simple
        };
    }

    public override void ExitCompositeLit(GoParser.CompositeLitContext context)
    {
        // literal
        //     : basicLit
        //     | compositeLit
        //     | functionLit

        // compositeLit
        //    : literalType literalValue

        if (context.Parent.Parent is not GoParser.OperandContext operandContext)
        {
            AddWarning(context, $"Could not derive parent operand context from composite literal: \"{context.GetText()}\"");
            return;
        }

        GoParser.LiteralTypeContext literalType = context.literalType();
        GoParser.LiteralValueContext literalValue = context.literalValue();
        GoParser.KeyedElementContext[] keyedElements = literalValue.elementList()?.keyedElement();
        bool isDynamicSizedArray = literalType.elementType() is not null;
        List<(string key, string element)> elements = new();
        bool hasKeyedElement = false;

        var textSqlSelect = "";
        TNoSqlCode noSqlSelect = null;

        if (keyedElements is not null)
        {
            var bIsStructInit = true;

            foreach (GoParser.KeyedElementContext keyedElement in keyedElements)
            {
                string key = keyedElement.key()?.GetText();
                string element = keyedElement.element().GetText();

                elements.Add((key, element));

                if (key is not null && !hasKeyedElement)
                    hasKeyedElement = true;
                
                if( keyedElement.key() is null ) 
                {
                    bIsStructInit = false;
                }
            }
            Types.TryGetValue( literalType, out var typeInfo1);
            
            if( bIsStructInit && m_LambdaCode != null )
            {
                (textSqlSelect, noSqlSelect) = this.Lambda_SelectFields( context, keyedElements, typeInfo1);
            }
        }

        string expressionText;
        TypeInfo typeInfo;

        if (literalType.structType() is not null)
        {
            // TODO: Need to properly handle in-line struct, see "src\Examples\Manual Tour of Go Conversions\moretypes\slice-literals"
            expressionText = $"/* TODO: Fix this in ScannerBase_Expression::ExitCompositeLit */ {context.GetText()}";
            typeInfo = new TypeInfo
            {
                Name = literalType.GetText(),
                TypeName = literalType.GetText(),
                FullTypeName = literalType.GetText(),
                TypeClass = TypeClass.Struct
            };
        }
        else if (literalType.arrayType() is not null || isDynamicSizedArray)
        {
            if( m_LambdaCode != null && m_LambdaCode.IsAllowArray )
            {
                textSqlSelect = $"{string.Join(", ", elements.Select(kvp => kvp.element))}";
            }else
            {
                textSqlSelect = getTextSQLError( "array are not allowed in sql transalation", context);
            }

            if (Types.TryGetValue(literalType.arrayType()?.elementType() ?? literalType.elementType(), out typeInfo))
            {
                getTextNoSQLError("array index not supported in nosql translation", context);
                if (typeInfo?.TypeClass == TypeClass.Interface)
                {
                    for (int i = 0; i < elements.Count; i++)
                    {
                        (string key, string element) = elements[i];
                        elements[i] = (key, $"{typeInfo.TypeName}.As({element})!");
                    }
                }

                string typeName = typeInfo?.TypeName ?? "object";
                string arrayLength = isDynamicSizedArray ? "-1" : literalType.arrayType().arrayLength().GetText();

                expressionText = hasKeyedElement ? 
                    $"new array<{typeName}>(InitKeyedValues<{typeName}>({(isDynamicSizedArray ? string.Empty: $"{arrayLength}, ")}{string.Join(", ", elements.Select(kvp => kvp.key is null ? kvp.element : $"({kvp.key}, {kvp.element})"))}))" :
                    $"new array<{typeName}>(new {typeName}[] {{ {string.Join(", ", elements.Select(kvp => kvp.element))} }})";

                typeInfo = new ArrayTypeInfo
                {
                    Name = typeName,
                    TypeName = $"array<{typeName}>",
                    FullTypeName = $"go.array<{typeInfo.FullTypeName}>",
                    TargetTypeInfo = typeInfo,
                    Length = new ExpressionInfo
                    {
                        Text = arrayLength,
                        Type = new TypeInfo
                        {
                            Name = "int",
                            TypeName = "int",
                            FullTypeName = "int",
                            TypeClass = TypeClass.Simple,
                            IsConst = true
                        }
                    },
                    TypeClass = TypeClass.Array
                };
            }
            else
            {
                AddWarning(context, $"Failed to find element type for the array type expression in \"{context.GetText()}\"");
                return;
            }
        }
        else if (literalType.sliceType() is not null)
        {
            getTextNoSQLError("array index not supported in nosql translation", context);
            if( m_LambdaCode != null && m_LambdaCode.IsAllowArray )
            {
                textSqlSelect = $"{string.Join(", ", elements.Select(kvp => kvp.element))}";

                var listElems = elements.Select(kvp => (TNoSqlCode)(new TNoSqlSelectField( kvp.key, kvp.element )) ).ToList();
                noSqlSelect  = new TNoSqlCode( "select", listElems );
            }else
            {
                textSqlSelect = getTextSQLError( "array are not allowed in sql transalation", context);
            }

            if (Types.TryGetValue(literalType.sliceType().elementType(), out typeInfo))
            {
                string typeName = typeInfo.TypeName;

                if (typeInfo.TypeClass == TypeClass.Interface)
                {
                    for (int i = 0; i < elements.Count; i++)
                    {
                        (string key, string element) = elements[i];
                        elements[i] = (key, $"{typeInfo.TypeName}.As({element})!");
                    }
                }

                expressionText = hasKeyedElement ?
                    $"InitKeyedValues<{typeName}>({string.Join(", ", elements.Select(kvp => kvp.key is null ? kvp.element : $"({kvp.key}, {kvp.element})"))}).slice()" :
                    $"new {typeName}[] {{ {string.Join(", ", elements.Select(kvp => kvp.element))} }}.slice()";

                typeInfo = new TypeInfo
                {
                    Name = typeName,
                    TypeName = typeName,
                    FullTypeName = typeInfo.FullTypeName,
                    TypeClass = TypeClass.Slice
                };
            }
            else
            {
                AddWarning(context, $"Failed to find element type for the slice type expression in \"{context.GetText()}\"");
                return;
            }
        }
        else if (literalType.mapType() is not null)
        {
            getTextNoSQLError("maps are not allowed in nosql translation", context);
            textSqlSelect = getTextSQLError( "maps are not allowed in sql transalation", context);
            // TODO: Need to properly handle map literals, see "src\Examples\Manual Tour of Go Conversions\moretypes\map-literals-continued"
            if (Types.TryGetValue(literalType.mapType().type_(), out typeInfo) && Types.TryGetValue(literalType.mapType().elementType(), out TypeInfo elementTypeInfo))
            {
                expressionText = $"/* TODO: Fix this in ScannerBase_Expression::ExitCompositeLit */ new map<{typeInfo.TypeName}, {elementTypeInfo.TypeName}>{literalValue.GetText()}";
                    
                typeInfo = new MapTypeInfo
                {
                    Name = "map",
                    TypeName = "map",
                    FullTypeName = "go.map",
                    KeyTypeInfo = typeInfo,
                    ElementTypeInfo = elementTypeInfo,
                    TypeClass = TypeClass.Map
                };
            }
            else
            {
                AddWarning(context, $"Failed to find key and/or value type for the map type expression in \"{context.GetText()}\"");
                return;
            }
        }
        else if (literalType.typeName() is not null)
        {
            getTextNoSQLError("types are not allowed in nosql translation", context);
            // TODO: Need to determine how to properly employ keyed elements here - guess is type aliases to array/slice/map would need to map back to original implementations
            expressionText = $"new {literalType.GetText()}({RemoveSurrounding(literalValue.GetText(), "{", "}")})";
                
            typeInfo = new TypeInfo
            {
                Name = literalType.GetText(),
                TypeName = literalType.GetText(),
                FullTypeName = $"go.{literalType.GetText()}",
                TypeClass = TypeClass.Simple
            };
        }
        else
        {
            AddWarning(context, $"Unexpected literal type \"{context.GetText()}\"");
            return;
        }

        Operands[operandContext] = new ExpressionInfo
        {
            Text = expressionText,
            Type = typeInfo,
            SQLText = textSqlSelect,
            NoSQLCode = noSqlSelect,
            bIsNoSql = false,
        };
    }

    

    public override void ExitOperandName(GoParser.OperandNameContext context)
    {
        // operand
        //     : literal
        //     | operandName
        //     | methodExpr
        //     | '(' expression ')'

        if (context.Parent is not GoParser.OperandContext operandContext)
        {
            AddWarning(context, $"Could not derive parent operand context from operand name: \"{context.GetText()}\"");
            return;
        }

        // operandName
        //     : IDENTIFIER
        //     | qualifiedIdent

        // TODO: var assignment is temporary, to resolve actual type, converter would override to load identifier metadata and recursively resolve identifier based components

        var identifier= context.GetText();
        var identifParts = getIdentifierParts( identifier );
        if( identifParts.Length > 1 )
        { 
            var funcName = identifParts[identifParts.Length-1];
            if( searchFunctionInImports( funcName, new TPackageInfo( identifParts[0]), out FunctionInfo function))
            {
                Operands[operandContext] = new ExpressionInfo
                {
                    Text = identifier,
                    Type = TypeInfo.VarType,
                    SQLText = "", 
                };
                return;
            }

            if( variableTypesTryGetValue( identifParts[0], out var typeVar) )
            {
                Operands[operandContext] = new ExpressionInfo
                {
                    Text = context.GetText(),
                    Type = typeVar.Type,
                    SQLText = getTextSQLError("variable.member is not allowed in sql transalation", context),
                };
                return;
            }
        }else
        if( identifParts.Length == 1 )
        { 
            if( identifParts[0] == "true"
             || identifParts[0] == "false" )
            { 
                Operands[operandContext] = new ExpressionInfo
                {
                    Text = context.GetText(),
                    Type = TypeInfo.BoolType,
                    OperandKind = EOperandKind.Operator,
                    SQLText = identifParts[0] == "true"
                                    ? Options.ConvertSql.SqlDialect.TRUE
                                    : Options.ConvertSql.SqlDialect.FALSE
                };
                return;
            }
            if( identifParts[0] == "nil")
            { 
                Operands[operandContext] = new ExpressionInfo
                {
                    Text = context.GetText(),
                    Type = TypeInfo.VoidType,
                    OperandKind = EOperandKind.Simple,
                    SQLText = Options.ConvertSql.SqlDialect.NULL,
                    NoSQLCode = (new TNoSqlKeyword("nil")),
                    bIsNoSql = false,
                };
                return;
            }
            if( CurrentFunction != null && CurrentFunction.ThisVar != null 
                &&  CurrentFunction.ThisVar.Name == identifParts[0] )
            {
                if( CurrentFunction.ThisVar.Type == null)
                    Debugger.Break();

                Operands[operandContext] = new ExpressionInfo
                {
                    Text = context.GetText(),
                    Type = CurrentFunction.ThisVar.Type,
                    OperandKind = EOperandKind.THIS,
                    SQLText = Options.ConvertSql.SqlDialect.THIS,
                    NoSQLCode = new TNoSqlCode(Options.ConvertSql.SqlDialect.THIS),
                    bIsNoSql = false,
                };
                return;
            }

            if( ImportMetadata != null && ImportMetadata.TryGetValue( identifier, out var metadata ) ) 
            {
                Operands[operandContext] = new ExpressionInfo
                {
                    Text = identifier,
                    Type = TypeInfo.VarType,
                    SQLText = "", 
                };
                return;
            }
            if( ImportAliases.TryGetValue( identifier, out var aliases ) ) {

                Operands[operandContext] = new ExpressionInfo
                {
                    Text = identifier,
                    Type = TypeInfo.VarType,
                    SQLText = "", 
                };
                return;
            }

            //LocalVariableInfo typeVar;
            if( variableTypesTryGetValue( identifParts[0], out var typeVar) )
            {
                if( typeVar == null)
                    Debugger.Break();

                if( typeVar.IsParameter )
                {
                    var SQLText = "";
                    var genType = typeVar.Type.getNormalizedType();
                    if( searchStructInImports(genType.Name, genType.PackageInfo, out StructInfo structInfo))
                    {
                        var tableName = getSqlTableNameFromStructInfo( structInfo);
                        if( tableName != "")
                        {
                            SQLText = getTextSQLIdentif( null, tableName, "", typeVar.Type, "", context, "");
                        }else
                        {
                            if( m_LambdaCode != null && m_LambdaCode.IsQueuedSelect)
                            {
                                SQLText = "";// Select().Where(x->func lambda(){} )
                            }else
                            {
                                SQLText = getTextSQLError( $"struct {typeVar.Type.Name} missing {OrmDef.Atomicsql_table} comment", context );
                            }
                        }
                        Operands[operandContext] = new ExpressionInfo
                        {
                            Text = context.GetText(),
                            Type = typeVar.Type,
                            SQLText = SQLText,
                            OperandKind = EOperandKind.Simple,
                            //NoSQLCode = new TNoSqlIdentifier( getTextSQLVarIdentif( identifier, typeVar.Type, context ) ),
                        };
                    }else
                    {
                        Operands[operandContext] = new ExpressionInfo
                        {
                            Text = context.GetText(),
                            Type = typeVar.Type,
                            SQLText = getTextSQLVarIdentif( identifier, typeVar.Type, context), 
                            NoSQLCode = new TNoSqlIdentifier( getTextSQLVarIdentif( identifier, typeVar.Type, context ) ),
                            OperandKind = EOperandKind.Simple,
                        };
                        //SQLText = getTextSQLError( "internal error 143", context );
                    }
                    
                }else
                {
                    if( m_globalVariables.TryGetValue( identifParts[0], out VariableInfo varGlobal))
                    {
                        var constExpr = varGlobal.getConstantLiteralExpr();
                        if( constExpr != null)
                        {
                            Operands[operandContext] = new ExpressionInfo
                            {
                                Text = context.GetText(),
                                Type = typeVar.Type,
                                SQLText = constExpr.SQLText, 
                                NoSQLCode = constExpr.NoSQLCode,
                                bIsNoSql = constExpr.bIsNoSql,
                                OperandKind = EOperandKind.Simple,
                            };
                            return;
                        }
                    }

                    Operands[operandContext] = new ExpressionInfo
                    {
                        Text = context.GetText(),
                        Type = typeVar.Type,
                        SQLText = getTextSQLVarIdentif( identifier, typeVar.Type, context), 
                        NoSQLCode = new TNoSqlIdentifier( getTextSQLVarIdentif( identifier, typeVar.Type, context ) ),
                        OperandKind = EOperandKind.Simple,
                    
                    };
                }
                return;
            }
            {
                if( TypeInfo.isPrimaryType(identifParts[0]) )
                {
                    Operands[operandContext] = new ExpressionInfo
                    {
                        Text = context.GetText(),
                        Type = TypeInfo.VarType,
                        SQLText = "", 
                        OperandKind = EOperandKind.Simple,
                    };
                    return;
                }
            }
        }
        //this.CurrentFunction.Signature.

        Operands[operandContext] = new ExpressionInfo
        {
            Text = context.GetText(),
            Type = TypeInfo.VarType,
            SQLText = getTextSQLVarIdentif( identifier, TypeInfo.VarType, context ), 
            NoSQLCode = new TNoSqlIdentifier( getTextSQLVarIdentif( identifier, TypeInfo.VarType, context ) ),
            OperandKind = EOperandKind.Simple,
        };
    }

    


    public override void ExitMethodExpr([NotNull] GoParser.MethodExprContext context)
    {
        // operand
        //     : literal
        //     | operandName
        //     | methodExpr
        //     | '(' expression ')'

        if (context.Parent is not GoParser.OperandContext operandContext)
        {
            AddWarning(context, $"Could not derive parent operand context from method expression: \"{context.GetText()}\"");
            return;
        }

        // methodExpr
        //     : receiverType '.' IDENTIFIER

        // receiverType
        //     : typeName
        //     | '(' '*' typeName ')'
        //     | '(' receiverType ')'

        GoParser.ReceiverTypeContext receiverType = context.receiverType();

        // TODO: should this be a delegate to an extension function? Need a use case...
        string receiver;

        if (receiverType?.children.Count == 4)
            receiver = $"ptr<{receiverType.typeName().GetText()}>";
        else
            receiver = context.GetText();

        Operands[operandContext] = new ExpressionInfo
        {
            Text = receiver,
            Type = TypeInfo.ObjectType
        };
    }
}
