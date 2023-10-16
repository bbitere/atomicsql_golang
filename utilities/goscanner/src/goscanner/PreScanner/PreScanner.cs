﻿//******************************************************************************************************
//  PreScanner.cs - Gbtc
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

using Antlr4.Runtime;
using goscanner.ConvCommon;
using goscanner.Metadata;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.Json;

#pragma warning disable SCS0018 // Path traversal

namespace goscanner;

/// <summary>
/// Represents a pre-scanner used to create meta-data dictionaries of Go source code.
/// </summary>
/// <remarks>
/// A full pre-scan of source code is needed in order to properly handle promotions within
/// interfaces and structures.
/// </remarks>
public partial class PreScanner : ScannerBase
{
    private readonly Dictionary<string, (string targetImport, string targetUsing)> m_importAliases = new(StringComparer.Ordinal);
    private readonly Dictionary<string, InterfaceInfo> m_interfaces = new(StringComparer.Ordinal);
    private readonly Dictionary<string, StructInfo> m_structs = new(StringComparer.Ordinal);
    private readonly Dictionary<string, FunctionInfo> m_functions = new(StringComparer.Ordinal);
    private readonly Dictionary<string, VariableInfo> m_variables = new(StringComparer.Ordinal);
    private readonly Dictionary<string, VariableInfo> m_globalVariables = new(StringComparer.Ordinal);
    

    
    public string FolderMetadataFileName { get; }

    public PreScanner(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName) : base(tokenStream, parser, options, fileName)
    {
        FolderMetadataFileName = GetFolderMetadataFileName(options, null, fileName);
    }

    public override (bool, string) Scan(bool _)
    {
        FolderMetadata folderMetadata = GetFolderMetadata(Options, null, SourceFileName) ?? new FolderMetadata();
        FileMetadata fileMetadata = folderMetadata.Files.GetOrAdd(SourceFileName, new FileMetadata());
        TypeInfo.s_CurrentMetaData = fileMetadata;

        // Base class walks parse tree
        (bool success, string result) = base.Scan(false);

        if (!success)
            return (false, result);        

        if( folderMetadata == null)
            Debugger.Break();

        fileMetadata.setFolderMetadata( folderMetadata);
        fileMetadata.Package = Package;
        fileMetadata.PackageImport = PackageImport;
        fileMetadata.SourceFileName = SourceFileName;
        fileMetadata.TargetFileName = TargetFileName;
        fileMetadata.ImportAliases = m_importAliases;
        fileMetadata.Interfaces = m_interfaces;
        fileMetadata.Structs = m_structs;
        fileMetadata.Functions = m_functions;
        fileMetadata.GlobalVariables = m_globalVariables;
        fileMetadata.LastUpdate = DateTime.UtcNow;

    #if !DEBUG
        try
        {
    #endif
        
        var useMetadata = true;
        if( Options.ConvertSql != null)
            useMetadata = Options.ConvertSql.UseMetadata;

        if( useMetadata )
        {
            var exist = File.Exists( FolderMetadataFileName );
            string directory = Path.GetDirectoryName(FolderMetadataFileName);

            if (!Directory.Exists(directory))
                Directory.CreateDirectory(directory);

            string serializedData = JsonSerializer.Serialize(folderMetadata, GetSerializationOptions());
            File.WriteAllText(FolderMetadataFileName, serializedData);
            if(!exist)
            {
                Console.WriteLine("");
                Console.WriteLine($"scanner - create metadata file: {Utils1.getFileName(FolderMetadataFileName)}");
                Console.WriteLine($"At: '{FolderMetadataFileName}'");
                
            }                
        }
    #if !DEBUG
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Failed to write metadata file \"{FolderMetadataFileName}\": {ex.Message}");
        }
    #endif

        Console.Write(".");
        return (true, null);
    }

    protected override void BeforeScan()
    {
        //Console.WriteLine($"Starting pre-scan of \"{SourceFileName}\"...");
    }

    protected override void AfterScan()
    {
       //Console.WriteLine("    Finished.");
    }

    protected override void SkippingImport(string import)
    {
        //Console.WriteLine($"Skipping pre-scan for Go standard library import package \"{import}\".");
        //Console.WriteLine();
    }

    public static int TotalUpToDateMetadata { get; private set; }

    public static int TotalMetadataUpdates { get; private set; }

    public static void Scan(Options options)
    {
        Console.WriteLine($"GoScan.exe: Version {options.GetVersion() }");
        Console.WriteLine("Starting Go code pre-scan to update metadata...");
        Console.WriteLine();

        ResetScanner();
        Scan(false,options, false, CreateNewPreScanner, MetadataOutOfDate, HandleSkippedScan);

        TotalUpToDateMetadata = TotalProcessedFiles - TotalMetadataUpdates;
    }

    private static ScannerBase CreateNewPreScanner(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName)
    {
        return new PreScanner(tokenStream, parser, options, fileName);
    }

    private static bool MetadataOutOfDate(Options options, string fileName, out string message)
    {
        FolderMetadata  infoMetadata = null;
        return MetadataOutOfDate1( options,  fileName, out message, ref infoMetadata);
    }
    public static bool MetadataOutOfDate1(Options options, string fileName, out string message, ref FolderMetadata infoMetadata)
    {
        message = null;
        TotalMetadataUpdates++;

        if (options.ForceMetadataUpdate)
            return true;

        FolderMetadata folderMetadata = GetFolderMetadata(options, null, fileName);

        if (folderMetadata is null || !folderMetadata.Files.TryGetValue(fileName, out FileMetadata fileMetadata))
            return true;

        infoMetadata = folderMetadata;

        if (File.GetLastWriteTimeUtc(fileName) > fileMetadata.LastUpdate)
            return true;

        message = "";//$"Metadata for \"{fileName}\" is up to date.{Environment.NewLine}";
        TotalMetadataUpdates--;

        return false;
    }

    private static void HandleSkippedScan(Options options, string fileName, bool showParseTree)
    {
        if (options.LocalConvertOnly)
            return;

        // Check meta-data status for imports
        GetFilePaths(options, null, fileName, out string sourceFileName, out _, out _, out string targetFilePath);
        FolderMetadata folderMetadata = GetFolderMetadata(options, sourceFileName, targetFilePath);

        if (folderMetadata is not null && folderMetadata.Files.TryGetValue(sourceFileName, out FileMetadata fileMetadata))
        {
            ImportQueue.UnionWith(fileMetadata.ImportAliases.Select(import => import.Value.targetImport));
            ScanImports(CreateNewPreScanner(null, null, options, fileName), showParseTree, CreateNewPreScanner, MetadataOutOfDate, HandleSkippedScan);
        }
    }
}