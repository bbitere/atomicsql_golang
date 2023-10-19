//******************************************************************************************************
//  ScannerBase.cs - Gbtc
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
// ReSharper disable UnusedMember.Global
// ReSharper disable InconsistentNaming

using Antlr4.Runtime;
using Antlr4.Runtime.Tree;
using Dahomey.Json;
using Dahomey.Json.Serialization.Conventions;
using goscanner.Metadata;
using Gemstone.Console;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.RegularExpressions;
using static goscanner.Common;
using System.Diagnostics;
using goscanner.ConvCommon;
using goscanner.ConvSql;

#pragma warning disable SCS0018 // Path traversal
#pragma warning disable SCS0028 // Deserialization

namespace goscanner;

public delegate ScannerBase CreateNewScannerFunction(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName);
public delegate bool FileNeedsScanFunction(Options options, string fileName, out string message);
public delegate void SkippedFileScanFunction(Options options, string fileName, bool showParseTree);


public class DataBaseTranslation { }
/// <summary>
/// Represents common Antlr scanning functionality.
/// </summary>
/// <remarks>
/// A full source code pre-scan is needed in order to be aware of promotions in structures and
/// interfaces. This base class represents the common scanning code that is used by both the
/// <see cref="PreScanner"/> and <see cref="ConvWebApi"/> classes.
/// </remarks>
public abstract partial class ScannerBase : GoParserBaseListener
{
    public const string RootNamespace = "go";
    public const string ClassSuffix = "_package";
    public const string AddressPrefix = "_addr_";

    private readonly List<string> m_warnings = new();

    protected readonly DependencyCounter RequiredUsings = new();

    private static readonly Regex s_buildDirective = new(@"^\s*\/\/\s*\+build", RegexOptions.Singleline | RegexOptions.Compiled);

    //public static FileMetadata  s_CurrentMetaData;

    protected sealed class ParserErrorListener : IAntlrErrorListener<IToken>
    {
        private readonly ScannerBase m_scanner;

        public ParserErrorListener(ScannerBase scanner) => m_scanner = scanner;

        public void SyntaxError(TextWriter output, IRecognizer recognizer, IToken offendingSymbol, int line, int charPositionInLine, string msg, RecognitionException e)
        {
            m_scanner.m_warnings.Add($"{Path.GetFileName(m_scanner.SourceFileName)}:{line}:{charPositionInLine}: {msg}");
        }
    }

    public Options Options { get; }

    public BufferedTokenStream TokenStream { get; }

    public GoParser Parser { get; }

    public FileMetadata Metadata { get; set;} // Only available after pre-scan

    public string SourceFileName { get; }

    public string SourceFilePath { get; }

    public string TargetFileName { get; }

    public string TargetFilePath { get; }

    public string[] Warnings => m_warnings.ToArray();

    protected string CurrentImportPath { get; private set; }

    protected bool UsesUnsafePointers { get; set; }

    protected bool InFunction { get; set; }

    protected FunctionInfo CurrentFunction { get; set; }

    protected string OriginalFunctionName { get; set; }
        
    protected string CurrentFunctionName { get; set; }

    protected ScannerBase(BufferedTokenStream tokenStream, GoParser parser, Options options, string fileName)
    {
        Options = options;

        if (fileName is null)
            throw new ArgumentNullException(nameof(fileName));

        if (!File.Exists(fileName))
            throw new FileNotFoundException($"WARNING: Source file \"{fileName}\" cannot be found.", fileName);

        TokenStream = tokenStream;
        Parser = parser;

        GetFilePaths(options, null, fileName, out string sourceFileName, out string sourceFilePath, out string targetFileName, out string targetFilePath);

        SourceFileName = sourceFileName;
        SourceFilePath = sourceFilePath;
        TargetFileName = targetFileName;
        TargetFilePath = targetFilePath;

        FolderMetadata folderMetadata = GetFolderMetadata(Options, null, SourceFileName, targetFilePath);

        if (folderMetadata is not null && folderMetadata.Files.TryGetValue(fileName, out FileMetadata metadata))
        {
            Metadata = metadata;
            Metadata.setFolderMetadata( folderMetadata);
        }
    }

    public virtual (bool, string) Scan(bool showParseTree)
    {
        IParseTree sourceTree = Parser.sourceFile();

        if (showParseTree)
            Console.WriteLine(sourceTree.ToStringTree(Parser));

        // Walk parsed source tree to start visiting nodes
        ParseTreeWalker walker = new();

        try
        {
            walker.Walk(this, sourceTree);
        }
        catch (CgoTargetException)
        {
            return (true, "Import \"C\" CGO Target");
        }

        return (true, null);
    }

    protected void AddWarning(ParserRuleContext context, string message)
    {
        var disableWarnings = false;
        if( this.Options.ConvertWebApi != null)
            disableWarnings = this.Options.ConvertWebApi.DisableWarnings;
        else
        if( this.Options.ConvertSql != null)
            disableWarnings = this.Options.ConvertSql.DisableWarnings;

        if(!(disableWarnings) && message != "" )
        {
            var loc = Getlocation(context);
            m_warnings.Add($"{loc}: {message}");
        }
    }
    public string Getlocation(ParserRuleContext context)
    {
        return $"{Path.GetFileName(SourceFileName)}:{context.Start.Line}:{context.Start.Column}";
    }

    protected virtual void BeforeScan()
    {
    }

    protected virtual void AfterScan()
    {
    }

    protected virtual void SkippingScan()
    {
    }

    protected virtual void SkippingImport(string import)
    {
    }

    protected static readonly bool IsPosix;
    protected static readonly string GoRoot;
    protected static readonly string GoPath;
    protected static readonly string[] NewLineDelimeters;
    protected static readonly HashSet<string> ImportQueue;
    protected static readonly List<string> Imports;

    private static readonly HashSet<string> s_processedFiles;
    private static readonly HashSet<string> s_processedImports;
    protected static string s_currentFolderMetadataFileName;
    private static FolderMetadata s_currentFolderMetadata;

    static ScannerBase()
    {
        IsPosix = Path.DirectorySeparatorChar == '/';

        GoRoot = GetGoEnv("GOROOT");

        if (string.IsNullOrWhiteSpace(GoRoot))
            GoRoot = Path.GetFullPath($"{Path.DirectorySeparatorChar}Go");

        GoRoot = AddPathSuffix($"{AddPathSuffix(GoRoot)}src");

        if (!Directory.Exists(GoRoot))
            throw new InvalidOperationException($"Unable to resolve GOROOT src directory: \"{GoRoot}\". Validate that Go is properly installed.");

        GoPath = GetGoEnv("GOPATH");

        if (string.IsNullOrWhiteSpace(GoPath))
            GoPath = Environment.ExpandEnvironmentVariables(IsPosix ? "$HOME/go" : "%USERPROFILE%\\go");

        GoPath = AddPathSuffix($"{AddPathSuffix(GoPath)}src");

        if (!Directory.Exists(GoPath))
            Directory.CreateDirectory(GoPath);

        NewLineDelimeters = new[] { "\r\n", "\n" };

        s_processedFiles = new HashSet<string>(StringComparer.OrdinalIgnoreCase);
        s_processedImports = new HashSet<string>(StringComparer.OrdinalIgnoreCase);
        ImportQueue = new HashSet<string>(StringComparer.OrdinalIgnoreCase);
        Imports = new List<string>();
    }

    private static string GetGoEnv(string var)
    {
        string result = Environment.GetEnvironmentVariable(var);

        if (!string.IsNullOrWhiteSpace(result))
            return result;
        
        CommandResponse response = Command.Execute("go", $"env -json {var}");

        if (response.ExitCode != 0)
        {
            Console.Error.WriteLine($"Unable to resolve Go environment variable \"{var}\". Validate that Go is properly installed.");
            return null;
        }

        JsonDocument.Parse(response.StandardOutput).RootElement.TryGetProperty(var, out JsonElement value);
        result = value.GetString();

        try
        {
            Environment.SetEnvironmentVariable(var, result);
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Unable to set Go environment variable \"{var}\" to \"{result}\": {ex.Message}");
        }

        return result;
    }

    public static int TotalProcessedFiles => s_processedFiles.Count - TotalSkippedFiles;

    public static int TotalSkippedFiles { get; private set; }

    public static int TotalSkippedPackages { get; private set; }

    public static int TotalWarnings { get; private set; }

    protected static void ResetScanner()
    {
        s_processedFiles.Clear();
        s_processedImports.Clear();
        ImportQueue.Clear();
        Imports.Clear();
        TotalSkippedFiles = 0;
        TotalSkippedPackages = 0;
        TotalWarnings = 0;
    }

    private static bool ScanOrmMetadataOutOfDate(Options options, string fileName, out string message)
    {
        FolderMetadata folder = null;
        var r = PreScanner.MetadataOutOfDate1( options, fileName, out message, ref folder);

        if( folder != null && options.ConvertSql != null )
        {
            options.ConvertSql.OrmDirAtomicsql_DirMetadata = folder;
        }

        return r;
    }
    private static bool ScanOrmFuncMetadataOutOfDate(Options options, string fileName, out string message)
    {
        FolderMetadata folder = null;
        var r = PreScanner.MetadataOutOfDate1( options, fileName, out message, ref folder);

        if( folder != null && options.ConvertSql != null )
        {
            options.ConvertSql.OrmDirAtomicsqlFunc_DirMetadata = folder;
        }

        return r;
    }
    protected static DataBaseTranslation Scan(bool bTranslate, Options options, bool showParseTree, CreateNewScannerFunction createNewScanner, FileNeedsScanFunction fileNeedsScan = null, SkippedFileScanFunction handleSkippedScan = null)
    {
        if (fileNeedsScan is null)
            fileNeedsScan = DefaultFileNeedsScan;

        string sourcePath = GetAbsolutePath(options.SourcePath);

        if(options.ConvertSql != null )
        {
            if( options.ConvertSql.SourcePathDir != null )
                sourcePath = GetAbsolutePath(options.ConvertSql.SourcePathDir, Sql_ConfigTranslation.CONV_SOURCE_PATH_SEP);
        }else
        if(options.ConvertWebApi != null )
        {
            if( options.ConvertWebApi.SourcePathDir != null )
                sourcePath = GetAbsolutePath(options.ConvertWebApi.SourcePathDir);
        }

        if (File.Exists(sourcePath))
        {
            string filePath = Path.GetDirectoryName(sourcePath) ?? string.Empty;

            if (filePath.StartsWith(GoRoot, StringComparison.OrdinalIgnoreCase))
                ScanFile(Options.Clone(options, options.OverwriteExistingFiles, sourcePath, Path.Combine(options.TargetGoSrcPath, filePath[GoRoot.Length..])), sourcePath, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
            else
                ScanFile(options, sourcePath, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);

            return null;
        }
        else
        {
            if( options.ConvertAsqlModels != null)
            {
                scanDirectory(bTranslate, sourcePath, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                
                //if( options.ConvertWebApi.SourcePathDir2!= null )
                //    scanDirectory(bTranslate, options.ConvertWebApi.SourcePathDir2, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);

                return goscanner.ConvASqlModels.ConvASqlModels.DoTranslation();
            }else
            if( options.ConvertWebApi != null)
            {
                options.ConvertWebApi.bCollectMethods = true;
                scanDirectory(bTranslate, sourcePath, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                options.ConvertWebApi.bCollectMethods = false;
                
                if( options.ConvertWebApi.SourcePathDir2!= null )
                    scanDirectory(bTranslate, options.ConvertWebApi.SourcePathDir2, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);

                return goscanner.ConvWebApi.ConvWebApi.DoTranslation();
            }else
            if( options.ConvertSql != null)
            {
                options.ConvertSql.bCollectMethods = true;
                scanDirectory(bTranslate, sourcePath, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                options.ConvertSql.bCollectMethods = false;
                
                if( !bTranslate )
                {
                    if( options.ConvertSql.SourcePathDir2!= null )
                    {
                        scanDirectory(bTranslate, options.ConvertSql.SourcePathDir2, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                    }
                    if( options.ConvertSql.OrmDir_AtomicSql!= null )
                    {
                        scanDirectory(bTranslate, options.ConvertSql.OrmDir_AtomicSql, 
                                        options, showParseTree, createNewScanner, 
                                        ScanOrmMetadataOutOfDate, handleSkippedScan);
                    }
                    if( options.ConvertSql.OrmDir_AtomicSqlFunc!= null )
                    {
                        scanDirectory(bTranslate, options.ConvertSql.OrmDir_AtomicSqlFunc, 
                                        options, showParseTree, createNewScanner, 
                                        ScanOrmFuncMetadataOutOfDate, handleSkippedScan);
                    }

                    foreach( var file in options.ConvertSql.IncludeFiles  )
                    {
                        ScanFile( options, file, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                    }
                }

                return goscanner.ConvSql.SqlConvert.DoTranslation();

            }else
            {
                Debugger.Break();//nu tre sa ajunga aici
                //scanDirectory(bTranslate, sourcePath, options, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);

                //return goscanner.Converter.DoTranslation();
                return null;
            }
        }
        //else
        {
            //throw new DirectoryNotFoundException($"WARNING: Source path \"{sourcePath}\" cannot be found.");
        }
        //return null;
    }

    protected static void scanDirectory(bool bTranslate, string sourcePath, Options options, 
        bool showParseTree, CreateNewScannerFunction createNewScanner, FileNeedsScanFunction fileNeedsScan, SkippedFileScanFunction handleSkippedScan
        )
    {
        Regex excludeExpression = options.GetExcludeExpression();
        bool scanFileMatch(string fileName) => !excludeExpression.IsMatch(fileName);

        if( options.ConvertSql!= null )
        {
            var sourcePathParts = sourcePath.Split(Sql_ConfigTranslation.CONV_SOURCE_PATH_SEP);
            foreach( var _sourcePath1 in sourcePathParts)
            {
                var _sourcePath = _sourcePath1.Trim();
                if (!Directory.Exists(_sourcePath))
                {
                    new DirectoryNotFoundException($"WARNING: Source path \"{_sourcePath}\" cannot be found.");
                }

                var files = Directory.GetFiles(_sourcePath, "*.go", SearchOption.AllDirectories );
                foreach (string fileName in files)
                {
                    if (scanFileMatch(fileName))
                    {
                        var canonizePathFileName =  Path.GetFullPath(fileName);
                        ScanFile(options, canonizePathFileName, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                    }
                }
            }
        }else
        {
            if (!Directory.Exists(sourcePath))
            {
                new DirectoryNotFoundException($"WARNING: Source path \"{sourcePath}\" cannot be found.");
            }
            foreach (string fileName in Directory.EnumerateFiles(sourcePath, "*.go", SearchOption.TopDirectoryOnly))
            {
                if (scanFileMatch(fileName))
                {
                    var canonizePathFileName =  Path.GetFullPath(fileName);
                    ScanFile(options, canonizePathFileName, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                }
            }

            if (options.RecurseSubdirectories)
            {
                foreach (string subDirectory in Directory.EnumerateDirectories(sourcePath))
                {
                    string targetDirectory = options.TargetPath;

                    if (!string.IsNullOrEmpty(targetDirectory))
                        targetDirectory = Path.Combine(targetDirectory, RemovePathPrefix(subDirectory[sourcePath.Length..]));

                    Scan(bTranslate, Options.Clone(options, options.OverwriteExistingPackages, subDirectory, targetDirectory), showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                }
            }
        }
    }

    protected static void ScanFile(Options options, string fileName, bool showParseTree, CreateNewScannerFunction createNewScanner, FileNeedsScanFunction fileNeedsScan = null, SkippedFileScanFunction handleSkippedScan = null)
    {
        if (fileNeedsScan is null)
            fileNeedsScan = DefaultFileNeedsScan;

        if (!options.ParseGoOSTargets && IsGoOSTarget(fileName, out string suffix))
        {
            Console.WriteLine($"Encountered \"{suffix}\" OS target for \"{fileName}\", skipping scan...{Environment.NewLine}");
            handleSkippedScan?.Invoke(options, fileName, showParseTree);
            return;
        }

        if (!options.ParseGoArchTargets && IsGoArchTarget(fileName, out suffix))
        {
            Console.WriteLine($"Encountered \"{suffix}\" architecture target for \"{fileName}\", skipping scan...{Environment.NewLine}");
            handleSkippedScan?.Invoke(options, fileName, showParseTree);
            return;
        }

        if (!options.SkipBuildIgnoreDirectiveCheck || !options.ParseCgoTargets)
        {
            using StreamReader source = File.OpenText(fileName);
            bool foundIgnoreBuildDirective = false;
            bool foundCgoBuildDirective = false;
            string line;

            while ((line = source.ReadLine()) is not null)
            {
                // Directives must come before package definition
                if (line.TrimStart().StartsWith("package "))
                    break;

                if (!s_buildDirective.IsMatch(line))
                    continue;

                int index = line.IndexOf("+build", StringComparison.Ordinal);

                HashSet<string> directives = new(line[(index + 6)..].Split(' ', StringSplitOptions.RemoveEmptyEntries), StringComparer.Ordinal);

                if (!options.SkipBuildIgnoreDirectiveCheck && directives.Contains("ignore"))
                {
                    foundIgnoreBuildDirective = true;
                    break;
                }

                if (!options.ParseCgoTargets && directives.Any(directive => directive.Contains("cgo")))
                {
                    foundCgoBuildDirective = true;
                    break;
                }
            }

            if (foundIgnoreBuildDirective)
            {
                Console.WriteLine($"Encountered \"+build ignore\" directive for \"{fileName}\", skipping scan...{Environment.NewLine}");
                handleSkippedScan?.Invoke(options, fileName, showParseTree);
                return;
            }

            if (foundCgoBuildDirective)
            {
                Console.WriteLine($"Encountered \"+build cgo\" directive for \"{fileName}\", skipping scan...{Environment.NewLine}");
                handleSkippedScan?.Invoke(options, fileName, showParseTree);
                return;
            }
        }

        if (s_processedFiles.Contains(fileName))
            return;

        s_processedFiles.Add(fileName);

        if (!fileNeedsScan(options, fileName, out string message))
        {
            if( !(message == "" || message == null ))
            {
                Console.WriteLine(message);
            }
            handleSkippedScan?.Invoke(options, fileName, showParseTree);                
            return;
        }

        AntlrInputStream inputStream;

        using (StreamReader reader = File.OpenText(fileName))
            inputStream = new AntlrInputStream(reader);

        GoLexer lexer = new(inputStream);
        CommonTokenStream tokenStream = new(lexer);
        GoParser parser = new(tokenStream);
        ScannerBase scanner;

    #if !DEBUG
        try
        {
    #endif
        scanner = createNewScanner(tokenStream, parser, options, fileName);
    #if !DEBUG
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex.Message);
            TotalSkippedFiles++;
            return;
        }
    #endif

        parser.RemoveErrorListeners();
        parser.AddErrorListener(new ParserErrorListener(scanner));

        if (options.OverwriteExistingFiles || !File.Exists(scanner.TargetFileName))
        {
            scanner.BeforeScan();

            (bool success, string result) = scanner.Scan(showParseTree);

            if (!success)
            {
                Console.WriteLine($"Encountered \"{result}\" for \"{fileName}\", skipping scan...{Environment.NewLine}");
                handleSkippedScan?.Invoke(options, fileName, showParseTree);
                return;
            }

            scanner.AfterScan();            
        }
        else
        {
            scanner.SkippingScan();
            TotalSkippedFiles++;
        }

        string[] warnings = scanner.Warnings;

        if (warnings.Length > 0)
        {
            Console.WriteLine();
            Console.WriteLine("WARNINGS:");

            foreach (string warning in warnings)
                Console.WriteLine($"    {warning}");

            TotalWarnings += warnings.Length;
            Console.WriteLine();
        }
        

        if (!options.LocalConvertOnly)
            ScanImports(scanner, showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
    }

    protected static void ScanImports(ScannerBase scanner, bool showParseTree, CreateNewScannerFunction createNewScanner, FileNeedsScanFunction fileNeedsScan = null, SkippedFileScanFunction handleSkippedScan = null)
    {
        string[] imports = ImportQueue.ToArray();
        Options options = scanner.Options;

        foreach (string import in imports)
        {
            if (s_processedImports.Contains(import))
                continue;

            s_processedImports.Add(import);

            string importPath = AddPathSuffix(import.Replace("/", Path.DirectorySeparatorChar.ToString()));
            string goRootImport = Path.Combine(GoRoot, importPath);
            string goPathImport = Path.Combine(GoPath, importPath);
            string targetPath = null;

            if (Directory.Exists(goRootImport))
            {
                targetPath = Path.Combine(options.TargetGoSrcPath, importPath);

                if (options.ConvertStandardLibrary)
                {
                    Scan(false, Options.Clone(options, options.OverwriteExistingPackages, goRootImport, targetPath), showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
                }
                else
                {
                    // Only count package conversion as skipped when there are no existing converted files
                    if (PathHasFiles(targetPath, "*.cs"))
                        continue;

                    scanner.SkippingImport(import);
                    TotalSkippedPackages++;
                }
            }
            else if (Directory.Exists(goPathImport))
            {
                if (!string.IsNullOrEmpty(options.TargetPath))
                    targetPath = Path.Combine(options.TargetPath, importPath);

                Scan(false, Options.Clone(options, options.OverwriteExistingPackages, goPathImport, targetPath), showParseTree, createNewScanner, fileNeedsScan, handleSkippedScan);
            }
            else
            {
                var disableWarnings = false;

                if( options.ConvertWebApi != null)
                    disableWarnings = options.ConvertWebApi.DisableWarnings;
                else
                if( options.ConvertSql != null)
                    disableWarnings = options.ConvertSql.DisableWarnings;

                if(!(disableWarnings))
                {
                    Console.WriteLine($"WARNING: Failed to locate package \"{import}\" at either:");
                    //Console.WriteLine($"    {goRootImport} (from %GOROOT%)");
                    //Console.WriteLine($"    {goPathImport} (from %GOPATH%)");
                    //Console.WriteLine();
                }
            }
        }
    }

    protected static void GetFilePaths(Options options, string rootSourcePath, string fileName, 
        out string sourceFileName, out string sourceFilePath, 
        out string targetFileName, out string targetFilePath)
    {
        string rootTargetPath = string.IsNullOrEmpty(options.RootTargetPath) ? options.RootTargetPath : Path.GetFullPath(options.RootTargetPath);

        rootTargetPath = string.IsNullOrWhiteSpace(rootTargetPath) ? Path.GetFullPath(options.TargetGoSrcPath) : rootTargetPath;

        if (string.IsNullOrWhiteSpace(rootSourcePath))
            rootSourcePath = GetAbsolutePath(options.RootSourcePath);

        sourceFileName = Path.GetFullPath(fileName);
        sourceFilePath = Path.GetDirectoryName(sourceFileName) ?? string.Empty;
        targetFileName = $"{Path.GetFileNameWithoutExtension(sourceFileName)}.cs";

        if (string.IsNullOrWhiteSpace(options.TargetPath))
            targetFilePath = !options.ConvertStandardLibrary || string.IsNullOrWhiteSpace(options.TargetGoSrcPath) ? sourceFilePath : Path.GetFullPath(options.TargetGoSrcPath);
        else
            targetFilePath = options.TargetPath;

        string targetSubDirectory;

        if (Path.IsPathRooted(targetFilePath) && targetFilePath.StartsWith(rootTargetPath, StringComparison.OrdinalIgnoreCase))
            targetSubDirectory = RemovePathSuffix(RemovePathPrefix(targetFilePath[rootTargetPath.Length..]));
        else
            targetSubDirectory = RemovePathSuffix(RemovePathPrefix(targetFilePath));

        string sourceSubDirectory = "";
        //string sourceSubDirectory = sourceFilePath.StartsWith(rootSourcePath)
        //    ? RemovePathSuffix(RemovePathPrefix(sourceFilePath[rootSourcePath.Length..]))
        //    : string.Empty;

        if (!targetSubDirectory.Equals(sourceSubDirectory))
            targetFilePath = Path.Combine(targetFilePath, sourceSubDirectory!);

        targetFilePath = AddPathSuffix(targetFilePath);
        targetFileName = Path.Combine(targetFilePath, targetFileName);
    }

    protected static string GetFolderMetadataFileName1(Options options, string rootSourcePath, string fileName, string targetFilePath = null)
    {
        var sep = Path.DirectorySeparatorChar;
        var targetFilePath1 = rootSourcePath;
        if( string.IsNullOrWhiteSpace(targetFilePath1))
        {
            targetFilePath1 = Utils1.getDirPathOfFile( fileName );
            string lastDirName = GetLastDirectoryName(fileName);
            return $"{targetFilePath1}{sep}{lastDirName}.metadata";
        }else
        {
            string lastDirName = GetLastDirectoryName(targetFilePath1);
            return $"{targetFilePath1}{sep}{lastDirName}.metadata";
        }
    }
    protected static string GetFolderMetadataFileName(Options options, string rootSourcePath, string fileName, string targetFilePath = null)
    {
        if (string.IsNullOrWhiteSpace(targetFilePath))
            GetFilePaths(options, rootSourcePath, fileName, out _, out _, out _, out targetFilePath);

        //var file = new FileInfo(targetFilePath);
        //targetFilePath = file.Directory.FullName;

        string lastDirName = GetLastDirectoryName(targetFilePath);
        return $"{targetFilePath}{lastDirName}.metadata";
    }

    protected static FolderMetadata GetFolderMetadata(Options options, string rootSourcePath, string fileName, string targetFilePath = null)
    {
        string folderMetadataFileName = GetFolderMetadataFileName(options, rootSourcePath, fileName, targetFilePath);

        if (!File.Exists(folderMetadataFileName))
        {
            folderMetadataFileName = GetFolderMetadataFileName1(options, rootSourcePath, fileName, targetFilePath);
            if (!File.Exists(folderMetadataFileName))
                return null;
        }

        if (folderMetadataFileName.Equals(s_currentFolderMetadataFileName, StringComparison.OrdinalIgnoreCase) && s_currentFolderMetadata is not null)
            return s_currentFolderMetadata;

        FolderMetadata folderMetadata;

    #if !DEBUG
        try
        {
    #endif
        string serializedData = File.ReadAllText(folderMetadataFileName);
        folderMetadata = JsonSerializer.Deserialize<FolderMetadata>(serializedData, GetSerializationOptions());
    #if !DEBUG
        }
        catch (Exception ex)
        {
            folderMetadata = null;

            if (!folderMetadataFileName.Equals(s_currentFolderMetadataFileName, StringComparison.OrdinalIgnoreCase))
                Console.WriteLine($"Failed to load folder metadata from \"{folderMetadataFileName}\": {ex.Message}");
        }
    #endif

        s_currentFolderMetadataFileName = folderMetadataFileName;
        s_currentFolderMetadata = folderMetadata;

        if( folderMetadata != null )
        {
            foreach( var file in folderMetadata.Files )
            {
                if( file.Value.Structs != null )
                foreach( var struct1 in file.Value.Structs )
                {
                    foreach( var fld in struct1.Value.Fields )
                    {
                        fld.Type.setFileMetadata(file.Value);
                    }                
                }
                if( file.Value.Interfaces != null )
                foreach( var struct1 in file.Value.Interfaces )
                {
                    foreach( var func1 in struct1.Value.Methods )
                    {
                        foreach( var fld in func1.Signature.Parameters )
                        {
                            fld.Type.setFileMetadata(file.Value);
                        }     
                        foreach( var fld in func1.Signature.Result )
                        {
                            fld.Type.setFileMetadata(file.Value);
                        }  
                    }               
                }
                if( file.Value.Functions != null )
                foreach( var func1 in file.Value.Functions )
                {
                    foreach( var fld in func1.Value.Signature.Signature.Parameters )
                    {
                        fld.Type.setFileMetadata(file.Value);
                    }     
                    foreach( var fld in func1.Value.Signature.Signature.Result )
                    {
                        fld.Type.setFileMetadata(file.Value);
                    }  
                }
                if( file.Value.GlobalVariables != null )
                foreach( var var1 in file.Value.GlobalVariables )
                {
                    var1.Value.Type.setFileMetadata(file.Value);
                }
            }
        }

        return folderMetadata;
    }

    protected static JsonSerializerOptions GetSerializationOptions()
    {
        JsonSerializerOptions options = new()
        { IncludeFields = true };

        options.SetupExtensions();            
        DiscriminatorConventionRegistry registry = options.GetDiscriminatorConventionRegistry();
            
        registry.ClearConventions();

        registry.RegisterConvention(new DefaultDiscriminatorConvention<DerivedTypeInfo>(options, "DerivedTypeInfo"));
        registry.RegisterType<PointerTypeInfo>();
        registry.RegisterType<ArrayTypeInfo>();
        registry.RegisterType<MapTypeInfo>();

        registry.RegisterConvention(new DefaultDiscriminatorConvention<DerivedFunctionSignature>(options, "DerivedFunctionSignature"));
        registry.RegisterType<MethodSignature>();

        return options;
    }

    protected static FolderMetadata LoadImportMetadata(Options options, string targetImport, out string warning)
    {
        var disableWarnings = false;
        var rootProject = "";
        var gitProject = "";

        if( options.ConvertWebApi != null )
        {
            disableWarnings = options.ConvertWebApi.DisableWarnings;
            rootProject     = options.ConvertWebApi.RootProject;
            gitProject      = options.ConvertWebApi.GitProject;
            if(gitProject == null )
                Console.WriteLine("Warning: GitProject not set!");
            if(rootProject == null )
                Console.WriteLine("Warning: RootProject not set!");
        }

        if( options.ConvertSql != null )
        {
            disableWarnings = options.ConvertSql.DisableWarnings;
            rootProject     = options.ConvertSql.RootProject;
            gitProject      = options.ConvertSql.GitProject;

            if(gitProject == null )
                Console.WriteLine("Warning: GitProject not set!");
            if(rootProject == null )
                Console.WriteLine("Warning: RootProject not set!");
        }


        int lastSlash = targetImport.LastIndexOf('/');
        string packageName = lastSlash > -1 ? targetImport[(lastSlash + 1)..] : targetImport;        
        string importPath = $"{AddPathSuffix(targetImport.Replace("/", Path.DirectorySeparatorChar.ToString()))}{packageName}.go";
        string go2csPath = Path.Combine(GoPath, "go2cs");
        string goRootImport = Path.Combine(GoRoot, importPath);
        string goPathImport = Path.Combine(go2csPath, importPath);
        string targetPath = string.IsNullOrWhiteSpace(options.TargetGoSrcPath) ? go2csPath : options.TargetGoSrcPath;
        FolderMetadata metadata;

        warning = default;
        options = Options.Clone(options, options.OverwriteExistingFiles, GoRoot, targetPath);
        metadata = GetFolderMetadata(options, GoRoot, goRootImport);

        if (metadata is not null)
            return metadata;

        options = Options.Clone(options, options.OverwriteExistingFiles, go2csPath, targetPath);
        metadata = GetFolderMetadata(options, go2csPath, goPathImport);

        if (metadata is not null)
            return metadata;

        if( (rootProject != "" && rootProject != null) 
          && (gitProject != "" && gitProject != null) 
          && targetImport.StartsWith(gitProject) )
        {
            var Sep = Path.DirectorySeparatorChar.ToString();
            var targetImport1 = targetImport.Replace(gitProject, "").Replace("/", Sep );


            //string importPath = $"{AddPathSuffix(targetImport.Replace("/", Path.DirectorySeparatorChar.ToString()))}{packageName}.go";

            string goPathImportFile     = $"{targetImport1}{Sep}{packageName}.go";
            string folderPacketImport   =  $"{rootProject}{targetImport1}";
            string goPathImport1        =  $"{rootProject}{targetImport1}{Sep}{packageName}.go";
            goPathImport1               = Utils1.NormaliseDirPath(goPathImport1);
            folderPacketImport          = Utils1.NormaliseDirPath(folderPacketImport);

            options = Options.Clone(options, options.OverwriteExistingFiles, rootProject, rootProject);
            metadata = GetFolderMetadata(options, folderPacketImport, goPathImport1);
            //metadata = GetFolderMetadata(options, rootProject, goPathImport1);

            if (metadata is not null)
            {
                updateTypesStructFunctionsVars( metadata, packageName );
                return metadata;
            }
            //"github.com\\bbitere\\atomicsql_golang.git\\tests\\test1\\atomicsql_ormdefs\\atomicsql_ormdefs.go"
            if( options.ConvertSql != null)
            {
                if( targetImport == options.ConvertSql.OrmDir_Atomicsql_Git )
                {
                    metadata = options.ConvertSql.OrmDirAtomicsql_DirMetadata;
                    if (metadata is not null)
                    {
                        updateTypesStructFunctionsVars( metadata, packageName );
                        return metadata;
                    }
                }
                if( targetImport == options.ConvertSql.OrmDir_AtomicsqlFunc_Git )
                {
                    metadata = options.ConvertSql.OrmDirAtomicsqlFunc_DirMetadata;
                    if (metadata is not null)
                    {
                        updateTypesStructFunctionsVars( metadata, packageName );
                        return metadata;
                    }
                }
            }
        }


        StringBuilder loadWarning = new();

        

        if(!(disableWarnings))
        {
            if( options.ConvertAsqlModels == null )
            {
                loadWarning.AppendLine($"WARNING: Failed to locate package metadata for \"{targetImport}\" import at either:");
            }
            //loadWarning.AppendLine($"    {GetFolderMetadataFileName(options, GoRoot, goRootImport)} (from -g Go source target path)");
            //loadWarning.AppendLine($"    {GetFolderMetadataFileName(options, go2csPath, goPathImport)} (from %GOPATH%)");
        }

        warning = loadWarning.ToString();
        return null;
    }
    private static void updateTypesStructFunctionsVars( FolderMetadata metadataFolder, string packageName)
    {
        foreach( var file in metadataFolder.Files)
        {
            if( file.Value.Structs != null)
            foreach( var it in file.Value.Structs)
            {
                it.Value.setPackage( packageName );
            }
        }
        foreach( var file in metadataFolder.Files)
        {
            if( file.Value.Structs != null)
            foreach( var it in file.Value.Structs)
            {
                it.Value.setPackage( packageName );

                foreach( var fld in it.Value.Fields )
                {
                    var genType = fld.Type.getNormalizedType() as TGenType;
                    if (genType != null && genType.PackageName != "")
                    {
                        fld.Type.PackageInfo = genType.PackageInfo;
                    }
                    else
                    { 
                        fld.Type.PackageInfo = new TPackageInfo(packageName, file.Value);
                    }

                }
            }
        }
        foreach( var file in metadataFolder.Files)
        {
            if( file.Value.GlobalVariables != null)
            foreach( var it in file.Value.GlobalVariables)
            {
                it.Value.setPackage( packageName );
            }
        }
        foreach( var file in metadataFolder.Files)
        {
            if( file.Value.Interfaces != null )
            foreach( var it in file.Value.Interfaces)
            {
                it.Value.setPackage( packageName );
            }
        }
        
    }

    protected static string GetValidIdentifierName(string identifier)
    {
        int lastDotIndex = identifier.LastIndexOf('.');

        if (lastDotIndex > 0)
            identifier = identifier[(lastDotIndex + 1)..];

        return SanitizedIdentifier(identifier);
    }

    protected string GetUniqueIdentifier<T>(IDictionary<string, T> source, string identifier)
    {
        if (identifier.Equals("_"))
            return identifier;

        int count = 0;
        string uniqueIdentifier = identifier;

        while (source.ContainsKey(uniqueIdentifier))
            uniqueIdentifier = $"{identifier}@@{++count}";

        return uniqueIdentifier;
    }

    protected TypeInfo ConvertByRefToBasicPointer(TypeInfo typeInfo)
    {
        if (!typeInfo.IsByRefPointer)
            return typeInfo;

        string typeName = typeInfo.TypeName;
        string fullTypeName = typeInfo.FullTypeName;

        string[] parts = typeName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

        if (parts.Length == 2)
            typeName = $"ptr<{parts[1]}>";

        parts = fullTypeName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

        if (parts.Length == 2)
            fullTypeName = $"ptr<{parts[1]}>";

        return new PointerTypeInfo
        {
            Name = typeInfo.Name,
            TypeName = typeName,
            FullTypeName = fullTypeName,
            TypeClass = TypeClass.Simple,
            TargetTypeInfo = typeInfo
        };
    }

    protected TypeInfo ConvertByRefToNativePointer(TypeInfo typeInfo)
    {
        if (!typeInfo.IsByRefPointer)
            return typeInfo;

        string typeName = typeInfo.TypeName;
        string fullTypeName = typeInfo.FullTypeName;

        string[] parts = typeName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

        if (parts.Length == 2)
            typeName = $"*{parts[1]}";

        parts = fullTypeName.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

        if (parts.Length == 2)
            fullTypeName = $"*{parts[1]}";

        return new PointerTypeInfo
        {
            Name = typeInfo.Name,
            TypeName = typeName,
            FullTypeName = fullTypeName,
            TypeClass = TypeClass.Simple,
            TargetTypeInfo = typeInfo
        };
    }

    protected static string GetAbsolutePath(string filePath, string sep="")
    {
        if( sep != "" )
        {
            var arr = new List<string>();
            var parts = filePath.Split(sep);

            foreach(var part1 in parts)
            {
                var part = part1.Trim();
                if (!Path.IsPathRooted(part))
                    part = Path.Combine(GoPath, part!);
                arr.Add( Path.GetFullPath(part) );
            }

            return string.Join(sep, arr.ToArray() );
        }else
        {
            if (!Path.IsPathRooted(filePath))
                filePath = Path.Combine(GoPath, filePath!);

            return Path.GetFullPath(filePath);
        }
    }

    private static bool DefaultFileNeedsScan(Options options, string fileName, out string message)
    {
        message = null;
        return true;
    }
}
