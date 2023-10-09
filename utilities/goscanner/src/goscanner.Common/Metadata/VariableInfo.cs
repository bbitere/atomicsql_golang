//******************************************************************************************************
//  VariableInfo.cs - Gbtc
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
//  06/18/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using System;
using System.Collections.Generic;

namespace goscanner.Metadata
{
    [Serializable]
    public class LocalVariableInfo: Identif
    {
        public static int s_uid = 0;
        public int UID = s_uid++;

        public TypeInfo Type;
        public bool IsParameter;

        public LocalVariableInfo()
        {
            if( UID == 148)
                UID = UID;
                

        }
    }
    [Serializable]
    public class VariableInfo: Identif
    {
        public TypeInfo Type;
        
        public bool HeapAllocated;
        public bool Redeclared;

        public String[] InitExprFld;
        public String[] InitExprVal;

        private Dictionary<string,string> initExpr;
        public Dictionary<string,string> getInitExpr()
        {
            if( initExpr != null)
                return initExpr;

            initExpr = new Dictionary<string,string>();
            for( var i = 0; i < InitExprFld.Length; i++ )
            {
                initExpr[ InitExprFld[i] ] =  InitExprVal[i];
            }
            return initExpr;
        }
        public void setInitExpr(Dictionary<string,string> dict)
        { 
            initExpr = dict;
            var _InitExprFld = new List<string>();
            var _InitExprVal = new List<string>();
            foreach( var k in initExpr )
            {
                _InitExprFld.Add(k.Key);
                _InitExprVal.Add(k.Value);
            }
            this.InitExprFld = _InitExprFld.ToArray();
            this.InitExprVal = _InitExprVal.ToArray();
        }

    }
}
