﻿//******************************************************************************************************
//  Converter_Statement.cs - Gbtc
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
//  05/04/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using goscanner.Metadata;
using System;
using System.Text;
using static goscanner.Common;

namespace goscanner.ConvWebApi;

public partial class ConvWebApi
{
    public override void EnterStructType( GoParser.StructTypeContext context)
    {
        //var model = new StructInfo();
        //this.m_models.Add();
    }
    /// <summary>
    /// Exit a parse tree produced by <see cref="GoParser.structType"/>.
    /// </summary>
    /// <param name="context">The parse tree.</param>
    public override void ExitStructType( GoParser.StructTypeContext context)
    {

    }
    
}
