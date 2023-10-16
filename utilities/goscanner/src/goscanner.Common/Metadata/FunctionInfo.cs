﻿//******************************************************************************************************
//  FunctionInfo.cs - Gbtc
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
using System.Text.Json.Serialization;

namespace goscanner.Metadata
{
    [Serializable]
    public class FunctionInfo
    {
        public FunctionSignature Signature;
        public Dictionary<string, VariableInfo> Variables;
        public bool HasDefer;
        public bool HasPanic;
        public bool HasRecover;
        public ParameterInfo ThisVar;

        public ParameterInfo[] GenericTypes;
        

        [JsonIgnore]
        private FileMetadata m_fileMetadata;

        
        public FileMetadata getMetaFile(){ return m_fileMetadata;}
        public  void setMetaFile(FileMetadata f ){ m_fileMetadata = f;}
    }
}