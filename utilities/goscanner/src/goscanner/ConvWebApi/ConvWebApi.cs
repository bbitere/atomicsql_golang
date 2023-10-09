//******************************************************************************************************
//  Converter.cs - Gbtc
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
//  05/01/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using Antlr4.Runtime;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using goscanner.Metadata;
//using goscanner.Templates;
using static goscanner.Common;
using System.Security.AccessControl;
using Gemstone.StringExtensions;
using Antlr4.Runtime.Atn;
using System.Xml.Linq;

#pragma warning disable SCS0018 // Path traversal

namespace goscanner.ConvWebApi;

/// <summary>
/// Represents a converter used to convert Go source code to C#.
/// </summary>
/// 

public partial class ConvWebApi : goscanner.ConvCommon.ConvCommon
{
    // -o -i -h -w=.\export_svc_config.txt C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers 
    // C:\Work\SlotMonitor\SlotGames\game-originals-main\tools
    public ConvWebApi(BufferedTokenStream tokenStream, 
        GoParser parser, Options options, string currentFileName) 
        : base(tokenStream, parser, options, currentFileName)
    {
    }

    public class DataTranslation: DataBaseTranslation
    {
        public string textGoSvcFile = "";
        public string textTsSvcFile = "";
        public string textTsSvcModelsFile = "";
    }
    private static List<ConvWebApi> s_scanners= new List<ConvWebApi> ();
    protected override void collectScannerForTranslation()
    {
        s_scanners.Add(this);
    }
    
    public static ConvWebApi s_currentScanner = null;
    public static DataTranslation DoTranslation()
    {
        var data = new DataTranslation();

        try
        {
            foreach( var scanner in s_scanners ) 
            {
                s_currentScanner = scanner;

                scanner.scannerTranslation(data);
            }
        }
        catch(Exception e)
        {
            Console.WriteLine( e.Message );
        }
        return data;
    }
    private void scannerTranslation(DataTranslation data)
    {
        var txtGoSvc = WebApiBase_Go.exportSvc_AllMethods( this,
                        Options, m_exportedMethods, ref s_textPostGetRoute );
        //File.AppendAllText(this.Options.ConvertWebApi.GoSvcFile, txtGoSvc);
        data.textGoSvcFile += txtGoSvc;

        var txtTsSvc = WebApiBase_Ts.exportSvc_AllMethodsTs( this,
                        m_exportedMethods, this.Options.ConvertWebApi.TsSvcModelsFile);
        //File.AppendAllText(this.Options.ConvertWebApi.TsSvcFile, txtTsSvc);
        data.textTsSvcFile += txtTsSvc;

        var txtTsSvcModels = WebApiBase_Ts.exportSvcModels_AllMethodsTs(  this,
                    m_models, m_exportedMethods, s_dictCollectionFields, s_dictExportedModel );
        //File.AppendAllText(this.Options.ConvertWebApi.TsSvcModelsFile, txtTsSvcModels);
        data.textTsSvcModelsFile += txtTsSvcModels;

             
        /*
        var txtGoSvc = goscanner.ConvWebApi.WebApiBase_Go.exportSvc_AllMethods( m_methods, this.Options.ExportWebApi_GitHubControllers);
        File.AppendAllText(this.Options.ExportWebApi_GoSvcFile, txtGoSvc);

        var txtTsSvc = goscanner.ConvWebApi.WebApiBase_Ts.exportSvc_AllMethodsTs( m_methods, this.Options.ExportWebApi_TsSvcModelsFile);
        File.AppendAllText(this.Options.ExportWebApi_TsSvcFile, txtTsSvc);

        var txtTsSvcModels = goscanner.ConvWebApi.WebApiBase_Ts.exportSvcModels_AllMethodsTs( m_models, m_methods);
        File.AppendAllText(this.Options.ExportWebApi_TsSvcModelsFile, txtTsSvc);
        */
    }

    public static void Convert(Options options)
    {
        if (options.OnlyUpdateMetadata)
            return;

        if(options.ExportWebApi_ConfigFile == "" )
        { 
            Console.WriteLine("file config is not defined");
            return;
        }else
        {
            //moved up before prescan()
            //options.ConvertWebApi = new WebApi_ConfigTranslation();
            //options.ConvertWebApi.ParseConfigFile(options.ExportWebApi_ConfigFile);

            //File.WriteAllText(options.ConvertWebApi.TsSvcFile,"" );
            //File.WriteAllText(options.ConvertWebApi.TsSvcModelsFile, "");
        }

        if (options.ConvertWebApi.GoSvcFile == "")
        { 
            Console.WriteLine("file output GoSvc is node defined");
            return;
        }
        else
        {
            //File.WriteAllText(options.ConvertWebApi.GoSvcFile, "");
        }


        ResetScanner();
        var data = Scan( true, options, options.ShowParseTree, CreateNewConverter);
        //WriteProjectFiles(options);

        writeSvcFiles( options, data );        
    }
    private static void writeSvcFiles(Options options, DataBaseTranslation dataBase)
    {
        var data = dataBase as DataTranslation;
        if( data == null 
        || (data.textGoSvcFile == ""
         || data.textTsSvcFile == "" 
         || data.textTsSvcModelsFile == "" ) )
            return;

        //if(! s_exist_CLASS_GenericModel )
        //{
        //    Console.WriteLine("Generic_MODEL struct not defined. if you export models, you need to have it!");
        //}
        //var allMeths = File.ReadAllText(options.ConvertWebApi.GoSvcFile);

        var allMeths = data.textGoSvcFile;
        var txtGoSvc = WebApiBase_Go.ExportGoFile_Service( options, "ajax_service", "AjaxSvc", 
                            "",//options.ConvertWebApi.GitHubControllers,
                            allMeths, s_textPostGetRoute);
        File.WriteAllText( options.ConvertWebApi.GoSvcFile, txtGoSvc);
        Console.WriteLine($"Generate file: {options.ConvertWebApi.GoSvcFile}");

        //var allTsMeths = File.ReadAllText(options.ConvertWebApi.TsSvcFile);
        var allTsMeths = data.textTsSvcFile;
        var txtTsSvc = WebApiBase_Ts.Export_Service(options, "Ajax", "Svc", allTsMeths, options.ConvertWebApi.TsSvcModelsFile);
        File.WriteAllText(options.ConvertWebApi.TsSvcFile, txtTsSvc);
        Console.WriteLine($"Generate file: {options.ConvertWebApi.TsSvcFile}");

        //var allTsModels = File.ReadAllText(options.ConvertWebApi.TsSvcModelsFile);
        var allTsModels = data.textTsSvcModelsFile;
        var txtTsModels = WebApiBase_Ts.Export_SvcModels( options, allTsModels, s_dictCollectionFields );
        File.WriteAllText(options.ConvertWebApi.TsSvcModelsFile, txtTsModels);
        Console.WriteLine($"Generate file: {options.ConvertWebApi.TsSvcModelsFile}");
    }

    private static ScannerBase CreateNewConverter(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName)
    {
        return new ConvWebApi(tokenStream, parser, options, fileName);
    }

    private static void WriteProjectFiles(Options options)
    {
    #if !DEBUG
        try
        {
    #endif
        // Map of package names to list of package path and file names
        Dictionary<string, List<(string path, string[] fileNames)>> groupedPackageData;

        // Process import packages - these become shared projects
        groupedPackageData = CreateGroupedPackageData();

        // Process packages with "main" functions - these become standard projects
        ProcessMainProjectPackages(options);

        if (options.ConvertStandardLibrary && options.RecurseSubdirectories && AddPathSuffix(options.SourcePath).Equals(GoPath))
            ProcessStandardLibraryPackages(options, groupedPackageData);
    #if !DEBUG
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Failed to write project files: {ex.Message}");
        }
    #endif
    }
}

    
