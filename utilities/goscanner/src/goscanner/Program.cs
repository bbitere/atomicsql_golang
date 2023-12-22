//******************************************************************************************************
//  Program.cs - Gbtc
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

using System;
using System.Diagnostics;

namespace goscanner;

internal class Program
{
    private static int Main(string[] args)
    {
        DateTime startTime = DateTime.UtcNow;
        Arguments arguments = Arguments.Parse(args);
        Options options = arguments.ParsedOptions;
        options.CleanUpData();
        int exitCode = 1;

        //bbitere
        //if( options.ExportWebApi_LangType != null)
        //    options.ExportWebApi_LangType = options.ExportWebApi_LangType.Replace("=", "");
        //if (options.ExportSql_LangType != null)
        //    options.ExportSql_LangType = options.ExportSql_LangType.Replace("=", "");

        if (arguments.ParseSuccess)
            exitCode = RunConversion(options);
            
        if (exitCode == 0)
        {
            //Console.WriteLine($"Conversion complete for: go2cs {string.Join(" ", args)}");
            //Console.WriteLine();
            if( options.ConvertSql != null && options.ConvertSql.NumErrors > 0 )
            {
                Console.WriteLine($"!!!!!----------------!!!!!!!");
                Console.WriteLine($"Num of Errors in translation = {options.ConvertSql.NumErrors} ");
            }else
            {
                if( options.IsDefaultBehaviour() )
                {
                    Console.WriteLine($"Updated {PreScanner.TotalMetadataUpdates:N0} metadata files, {PreScanner.TotalUpToDateMetadata:N0} already up-to-date");
                }
            }

            if (!options.OnlyUpdateMetadata)
            {
                if(  options.IsDefaultBehaviour() )
                {
                    Console.WriteLine($"Converted {ScannerBase.TotalProcessedFiles:N0} Go files to C# with {ScannerBase.TotalWarnings:N0} total warnings");

                    if ((!options.OverwriteExistingFiles || !options.OverwriteExistingPackages) && ScannerBase.TotalSkippedFiles > 0)
                        Console.WriteLine($"Skipped {ScannerBase.TotalSkippedFiles:N0} already converted files (-o or -i option not set)");

                    if (!options.ConvertStandardLibrary && ScannerBase.TotalSkippedPackages > 0)
                        Console.WriteLine($"Skipped conversion of {ScannerBase.TotalSkippedPackages:N0} standard library packages (-s option not set)");
                }
            }
            if( options.IsDefaultBehaviour())
            {
                Console.WriteLine($"Processing time: {DateTime.UtcNow - startTime}");
            }
        }

        return exitCode;
    }

    private static int RunConversion(Options options)
    {
    #if !DEBUG
        try
        {
    #endif

        Common.RestoreResources(options.TargetGoSrcPath);

        if(options.ExportASqlModel_ConfigFile != "" )
        {
            options.ConvertAsqlModels = new goscanner.ConvASqlModels.ConvASqlModel_ConfigTranslation();
            options.ConvertAsqlModels.ParseConfigFile(options.ExportASqlModel_ConfigFile);
        }else
        if(options.ExportWebApi_ConfigFile != "" )
        {
            options.ConvertWebApi = new goscanner.ConvWebApi.ConvWebApi_ConfigTranslation();
            options.ConvertWebApi.ParseConfigFile(options.ExportWebApi_ConfigFile);
        }else
        if(options.ExportSql_ConfigFile != "" )
        {
            options.ConvertSql = new goscanner.ConvSql.Sql_ConfigTranslation();
            options.ConvertSql.ParseConfigFile(options.ExportSql_ConfigFile);
        }

        PreScanner.Scan(options);

        if( options.ExportASqlModel_ConfigFile != "")
        {
            goscanner.ConvASqlModels.ConvASqlModels.Convert(options);
        }
        else
        if( options.ExportWebApi_ConfigFile != "")
        {
            goscanner.ConvWebApi.ConvWebApi.Convert(options);
        }
        else
        if (options.ExportSql_ConfigFile != "")
        {
            goscanner.ConvSql.SqlConvert.Convert(options);
        }
        else
        { 
            //Converter.Convert(options);
            Debugger.Break();
        }

    #if !DEBUG
        }
        catch (TypeInitializationException ex)
        {
            Console.Error.WriteLine($"{Environment.NewLine}Error: {ex.InnerException?.Message ?? ex.Message}");
            return 2;
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"{Environment.NewLine}Error: {ex.Message}");
            return 3;
        }
    #endif

        return 0;
    }
}
