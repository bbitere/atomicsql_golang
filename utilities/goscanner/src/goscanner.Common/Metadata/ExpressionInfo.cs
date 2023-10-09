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
                || UID == 346
                || UID == 347
                || UID == 348
                || UID == 2861
                || UID == 2863
                || UID == 2864
                || UID == 2862
                || UID == 2789
                
                )
            {
                UID = UID;
            }
        }
        public override string ToString() => Text;
    }
}
