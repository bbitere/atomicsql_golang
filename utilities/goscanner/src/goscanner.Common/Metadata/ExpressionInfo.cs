//******************************************************************************************************
//  ExpressionInfo.cs - Gbtc
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
using System.Diagnostics;
using System.Linq;

namespace goscanner.Metadata
{
    // expression is fundamentally one of:
    //     unaryExpr
    //     expression operator expression
    //     primaryExpr

    // For unaryExpr, result type matches sub-expression

    // For expression operator expression
    //     if operator is comparison, result type is boolean
    //     otherwise, result type matches sub-expression

    // primaryExpr is more complex, it breaks down to one of:
    //     operand
    //     conversion
    //     primaryExpr operation

    // When primaryExpr is operand, result types are as follows:
    //     literal, result is literal type
    //     operandName, result is operand type
    //     methodExpr, result is a method expression
    //     (expression), result type matched sub-expression

    // When primaryExpr is conversion,
    //     result type is specified target type

    // When primary is expression has an operation, options are:
    //     DOT IDENIFIER, result type matches identifier
    //     index, result type matches sub-primaryExpr
    //     slice, result type matches sub-primaryExpr
    //     typeAssersion, result type matches specified target type
    //     arguments, result type matches sub-primaryExpr

    public enum EOperandKind
    {
        Simple,
        Operator,
        Function,
        THIS,
    }

    public class TNoSqlCode
    {
        public static TNoSqlCode TRUE = new TNoSqlCode("true");
        public static int s_UID = 1;
        public int UID = s_UID++;
        public string operatorName = "";
        public List< TNoSqlCode > operands = new List< TNoSqlCode >();

        public TNoSqlCode( string opName, params TNoSqlCode[] operands )
        {
            this.operatorName = opName;
            this.operands     = operands.ToList();
            if( this.UID == 622)
                UID = UID;
        }
        public TNoSqlCode( string opName, List<TNoSqlCode> operands )
        {
            this.operatorName = opName;
            this.operands     = operands;
            if( this.UID == 622)
                UID = UID;
        }
        private string generateTabs(int indentTab)
        {
            var tabs = "";
            for(var i = 0; i < indentTab; i++) 
            {
                tabs += "\t";
            }
            return tabs;
        }
        public virtual string getNoSqlCode(int indentTab)
        {
            var tabs = generateTabs(indentTab);
            if( this.operands.Count == 0 )
            {
                return $"\"{this.operatorName}\"";
            }
            if( this.operands.Count == 1 )
            {
                if( this.operands[0] != null)
                {
                    var op0 = this.operands[0].getNoSqlCode(indentTab+1);
                    return $@"
                            {tabs}[]any{{
					        {tabs}    ""{this.operatorName}"", {op0} 
                            {tabs}}}";
                }
                return "/*error 1 operands */";
            }
            if( this.operands.Count == 2 )
            {
                if( this.operands[0] != null && this.operands[1] != null)
                {
                    var op0 = this.operands[0].getNoSqlCode(indentTab+1);
                    var op1 = this.operands[1].getNoSqlCode(indentTab+1);
                    return $@"
                            {tabs}[]any{{
					        {tabs}    ""{this.operatorName}"", {op0}, {op1},
                            {tabs}}}";
                }
                return "/*error 2 operands. one is missing*/";
            }
            //if( this.operands.Count == 0 )
            {
                return $@" []any{{ ""{this.operatorName}"", }}";
            }
        }
    }
    
    public class TNoSqlIdentifier : TNoSqlCode
    {
        public TNoSqlIdentifier( string fieldName)
            :base( fieldName)
        {
        }
        public override string getNoSqlCode(int indentTab)
        {
            return $"\"{this.operatorName}\"";
        }
    }
    public class TNoSqlString : TNoSqlCode
    {
        public TNoSqlString( string fieldName)
            :base( fieldName)
        {
        }
        public override string getNoSqlCode(int indentTab)
        {
            return $"\"{this.operatorName}\"";
        }
    }
    public class TNoSqlNumber : TNoSqlCode
    {
        public TNoSqlNumber( string fieldName)
            :base( fieldName)
        {
        }
        public override string getNoSqlCode(int indentTab)
        {
            return $"{this.operatorName}";
        }
    }

    public class TNoSqlSelectField : TNoSqlCode
    {
        public string fieldName = "";
        public string  expression;

        public TNoSqlSelectField( string fieldName, string expression )
            :base( "-")
        {
            this.fieldName = fieldName;
            this.expression = expression;
        }
        public override string getNoSqlCode(int indentTab)
        {
            return $"\"{this.operatorName}\"";
        }
    }


    [Serializable]
    public class ExpressionInfo
    {

        public static int s_UID = 1;
        public int UID = ++s_UID;

        public string Text;
        public string LastToken;
        public TypeInfo Type;
        public ParameterInfo[] Types;
        //public List<string> SubExpressions = null;
        public string SQLText;
        public TNoSqlCode NoSQLCode = null;
        public bool       bIsNoSql = false;
        public EOperandKind  OperandKind = EOperandKind.Simple;
        public ExpressionInfo() {
            if(    false
                /*|| UID == 8843 
                || UID == 1245
                || UID == 1244  
                || UID == 1241                  
                || UID == 1240
                
                || UID == 1626
                || UID == 1627                
                || UID == 1628 
                || UID == 1625 
                || UID == 1624
                */

                
                || UID == 27186
                || UID == 27187
                || UID == 27188
                || UID == 27189
                || UID == 27190
                || UID == 27191
                || UID == 27192
                || UID == 27549
                
                )
            {
                UID = UID;
            }
        }
        public override string ToString() => Text;
    }
}
