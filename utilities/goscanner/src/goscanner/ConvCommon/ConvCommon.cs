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
using System.Diagnostics;
using goscanner.AST;
using Antlr4.Runtime.Misc;

#pragma warning disable SCS0018 // Path traversal

namespace goscanner.ConvCommon;

/// <summary>
/// Represents a converter used to convert Go source code to C#.
/// </summary>
/// 
public enum EOutputFile
{
    Cs,
    Ts,
    GoLang,
}
public partial class ConvCommon : ScannerBase
{
    protected List<FunctionInfo> m_exportedMethods = new List<FunctionInfo>();

    public EOutputFile TypeOfOutputFile = EOutputFile.Ts;

    public const string StandardLibrary = "GoStandardLibrary";
    protected const string UsingsMarker = ">>MARKER:USINGS<<";
    protected const string UnsafeMarker = ">>MARKER:UNSAFE<<";

    public StringBuilder m_targetOutputFile = new StringBuilder();
    protected StringBuilder m_targetGoFile = new StringBuilder();

    
    protected List<StructInfo> m_models= new List<StructInfo>();

    public static Dictionary<string, StructInfo> s_AllStructs = new Dictionary<string, StructInfo>();
    public static Dictionary<string, TypeInfo> s_AllTypeDefs = new Dictionary<string, TypeInfo>();
    
    protected static bool s_exist_CLASS_GenericModel = false;
    protected static Dictionary<string, string>  s_dictCollectionFields = new Dictionary<string, string>();
    protected static Dictionary<string, string>  s_dictExportedModel = new Dictionary<string, string>();
    

    

    protected static string s_textPostGetRoute = "";
    

    public const bool EMITTING_CODE = true;

    public Dictionary<string, (string targetImport, string targetUsing)> ImportAliases { get; }

    public Dictionary<string, FolderMetadata> ImportMetadata { get; }

    public ConvCommon(BufferedTokenStream tokenStream, 
        GoParser parser, Options options, string currentFileName) 
        : base(tokenStream, parser, options, currentFileName)
    {

        if (Metadata is null)
        {
            FolderMetadata folderMetadata = GetFolderMetadata(Options, null, SourceFileName, null);

            if (folderMetadata is not null && folderMetadata.Files.TryGetValue(currentFileName, out FileMetadata metadata))
            {
                Metadata = metadata;
                Metadata.setFolderMetadata( folderMetadata);
            }

            if (Metadata is null)
                throw new InvalidOperationException($"Failed to load metadata for \"{currentFileName}\" - file conversion canceled.");
        }

        //TargetFileName = options.ExportWebApi_Path;
        //TargetFilePath = targetFilePath;


        ImportAliases = Metadata.ImportAliases;
        ImportMetadata = new Dictionary<string, FolderMetadata>(StringComparer.Ordinal);
        
    }
    protected void addImportMetadata(string nameAlias, FolderMetadata metadata)
    {
        ImportMetadata[ nameAlias ] = metadata;
        if( this.SourceFileName.Contains("test1.go"))
            Utils.Nop();

        if( nameAlias == "orm" && this.SourceFileName.Contains("test1.go"))
            Utils.Nop();
    }

    

    public override (bool, string) Scan(bool showParseTree)
    {
        // Base class walks parse tree
        (bool success, string result) = base.Scan(showParseTree);

        if (!success)
            return (false, result);

        //if (!WroteLineFeed)
        //    m_targetOutputFile.AppendLine();

        // Close class and namespaces as begun during Converter_TopLevelDecl visit
        //m_targetOutputFile.AppendLine($"{Environment.NewLine}}} // end {Package}{ClassSuffix}");

        //if (!string.IsNullOrEmpty(m_namespaceFooter))
        //    m_targetOutputFile.AppendLine(m_namespaceFooter);

        //string targetTsFile = m_targetOutputFile.ToString();


        // Find usings marker
        
        collectScannerForTranslation();
        
        //writer.Write(targetFile);

        return (true, null);
    }

    /// <summary>
    /// need to ovveride
    /// </summary>
    protected virtual void collectScannerForTranslation(){ 
        Debugger.Break();
    }

    protected bool isPublicByName(string identName) 
    {
        return identName.Substring(0,1).IsAllUpper();
    }

    protected override void BeforeScan()
    {
        Console.WriteLine($"Scan file {Environment.NewLine}    \"{SourceFileName}\" ...");
    }

    protected override void AfterScan()
    {
        //if (!PackageImport.Equals("main"))
        //    Console.WriteLine($"        import \"{PackageImport}\" ==> using {PackageUsing}");

        if( Options.ConvertAsqlModels == null)
            Console.WriteLine("    Finished.");
    }

    protected override void SkippingScan()
    {
        //Console.WriteLine($"Skipping convert for{Environment.NewLine}    \"{SourceFileName}\", target{Environment.NewLine}    \"{TargetFileName}\" already exists.");
    }

    protected override void SkippingImport(string import)
    {
        // Console.WriteLine($"Skipping convert for Go standard library import package \"{import}\".");
        // Console.WriteLine();
    }

    protected static readonly HashSet<string> s_mainPackageFiles;
    protected static readonly Dictionary<string, Dictionary<string, (string nameSpace, HashSet<string> fileNames)>> s_packageInfo;

    static ConvCommon()
    {
        s_mainPackageFiles = new HashSet<string>(StringComparer.OrdinalIgnoreCase);
        s_packageInfo = new Dictionary<string, Dictionary<string, (string, HashSet<string>)>>(StringComparer.OrdinalIgnoreCase);
    }

    protected static Dictionary<string, List<(string, string[])>> CreateGroupedPackageData()
    {
        Dictionary<string, List<(string path, string[] fileNames)>> groupedPackageData = new(StringComparer.Ordinal);

        foreach (KeyValuePair<string, Dictionary<string, (string nameSpace, HashSet<string> fileNames)>> kvp in s_packageInfo)
        {
            string packagePath = kvp.Key;

            // Depending on the scope of the conversion, the same package name may exist in multiple paths
            foreach (KeyValuePair<string, (string nameSpace, HashSet<string> fileNames)> fileGroup in kvp.Value)
            {
                string package = fileGroup.Key;
                string[] packageFileNames = fileGroup.Value.fileNames.ToArray();

                List<(string, string[])> groupPackageData = groupedPackageData.GetOrAdd(package, _ => new List<(string, string[])>());
                groupPackageData.Add((packagePath, packageFileNames));
            }
        }

        return groupedPackageData;
    }

    protected static void ProcessMainProjectPackages(Options options)
    {
        foreach (string mainPackageFile in s_mainPackageFiles)
        {
            string mainPackageFileName = Path.GetFileName(mainPackageFile) ?? string.Empty;
            string mainPackagePath = Path.GetDirectoryName(mainPackageFile) ?? string.Empty;
            string assemblyName = Path.GetFileNameWithoutExtension(mainPackageFileName);

            FolderMetadata folderMetadata = GetFolderMetadata(options, null, mainPackageFile);
            string sourceFileName = Path.Combine(Path.GetDirectoryName(mainPackageFile) ?? string.Empty, $"{Path.GetFileNameWithoutExtension(mainPackageFile)}.go");

            if (folderMetadata is null || !folderMetadata.Files.TryGetValue(sourceFileName, out FileMetadata metadata))
                throw new InvalidOperationException($"Failed to load metadata for \"{sourceFileName}\" - file conversion canceled.");
            /*
            string mainProjectFile = Path.Combine(mainPackagePath, $"{assemblyName}.csproj");
            string mainProjectFileContent = new MainProjectTemplate
            {
                AssemblyName = assemblyName,
                Imports = metadata.ImportAliases.Select(kvp => kvp.Value.targetImport)
            }.TransformText();

            // Build main project file
            if (File.Exists(mainProjectFile) && GetMD5HashFromFile(mainProjectFile) == GetMD5HashFromString(mainProjectFileContent))
                continue;
            if (EMITTING_CODE)
            {
                using StreamWriter writer = File_CreateText(mainProjectFile);
                writer.Write(mainProjectFileContent);
            }
            */
        }
    }

    protected static void ProcessStandardLibraryPackages(Options options, Dictionary<string, List<(string path, string[] fileNames)>> groupedPackageData)
    {
        foreach (KeyValuePair<string, List<(string path, string[] fileNames)>> packageData in groupedPackageData)
        {
            foreach ((string path, string[] fileNames) rootPackage in packageData.Value.Where(info => info.path.StartsWith(GoRoot)))
            {
                foreach (string fileName in rootPackage.fileNames)
                {
                    if (fileName.EndsWith("_test.go"))
                        continue;

                    string assemblyName = packageData.Key;
                    string libraryProjectFile = Path.Combine(rootPackage.path, $"{assemblyName}.csproj");

                    FolderMetadata folderMetadata = GetFolderMetadata(options, null, fileName);
                    string sourceFileName = Path.Combine(Path.GetDirectoryName(fileName) ?? string.Empty, $"{Path.GetFileNameWithoutExtension(fileName)}.go");

                    if (folderMetadata is null || !folderMetadata.Files.TryGetValue(sourceFileName, out FileMetadata metadata))
                        throw new InvalidOperationException($"Failed to load metadata for \"{sourceFileName}\" - file conversion canceled.");

                    /*
                    string libraryProjectFileContent = new LibraryProjectTemplate
                    {
                        AssemblyName = assemblyName,
                        Imports = metadata.ImportAliases.Select(kvp => kvp.Value.targetImport)
                    }.TransformText();

                    // Build library project file
                    if (File.Exists(libraryProjectFile) && GetMD5HashFromFile(libraryProjectFile) == GetMD5HashFromString(libraryProjectFileContent))
                        continue;

                    if (EMITTING_CODE)
                    { 
                        using StreamWriter writer = File_CreateText(libraryProjectFile);
                        writer.Write(libraryProjectFileContent);
                    }
                    */
                }
            }
        }
    }

    protected static string GetPackageNamespace(string packageImport)
    {
        string[] paths = packageImport.Split('/').Select(SanitizedIdentifier).ToArray();
        return $"{RootNamespace}.{string.Join(".", paths)}{ClassSuffix}";
    }

    protected static void AddFileToPackage(string package, string fileName, string nameSpace)
    {
        // Since the same package name may exist at multiple paths, we track details by path
        Dictionary<string, (string, HashSet<string>)> packageInfo = s_packageInfo.GetOrAdd(Path.GetDirectoryName(fileName), _ => new Dictionary<string, (string, HashSet<string>)>(StringComparer.Ordinal));
        (string, HashSet<string> fileNames) fileGroup = packageInfo.GetOrAdd(package, _ => (nameSpace, new HashSet<string>(StringComparer.OrdinalIgnoreCase)));
        fileGroup.fileNames.Add(fileName);
    }

    /// <summary>
    /// get structInfo by name
    /// </summary>
    /// <param name="baseName"></param>
    /// <returns></returns>
    public static StructInfo getStructInfoByName( string baseName )
    {
        var baseName1 = goscanner.ConvWebApi.WebApiBase_Go.cleanNameGoStruct(baseName);
        if( s_AllStructs.ContainsKey( baseName1 ) )
        {
            var typeStruct = s_AllStructs[ baseName1 ];
            return typeStruct;
        }
        

        return null;
    }
    public StructInfo getStructInfoByName2( string baseName )
    {
        var baseName1 = goscanner.ConvWebApi.WebApiBase_Go.cleanNameGoStruct(baseName);
        if( s_AllStructs.ContainsKey( baseName1 ) )
        {
            var typeStruct = s_AllStructs[ baseName1 ];
            return typeStruct;
        }
        FileMetadata folderMeta = null;//it is null becauase I dont have any ideea how to set it 
        var typeName    = Utils.GetLastPart( baseName, "." );
        var packageName = Utils.GetFirstPart( baseName, "." );
        var genType     = TypeInfo.getPrimaryTypeName1( typeName, packageName, folderMeta);

        if( searchStructInImports(genType.Name, new TPackageInfo( genType.PackageName, folderMeta), out StructInfo structInfo))
        {
            return structInfo;
        }

        return null;
    }

    public StructInfo getStructInfoByType( TypeInfo type )
    {
        var genType = type.getNormalizedType();
        if( searchStructInImports(genType.Name, genType.PackageInfo, out StructInfo structInfo))
        {
            return structInfo;
        }
        return null;
    }

    

    public bool isInheritedBy(TypeInfo typeInfo, 
        string modelClassName, HashSet<string> list1 )
    {
        var structInfo = getStructInfoByName(typeInfo.Name);
        if( structInfo != null )    
        {
            return isInheritedBy( structInfo, modelClassName, list1 );
        }
        return false;
    }

    public bool isInheritedBy(StructInfo structInfo, 
        string modelClassName, HashSet<string> list1 )
    { 
        if( structInfo.Fields.Length > 0 )
        {
            TypeInfo baseType = null;
             if (structInfo.Fields[0].Name == goscanner.ConvWebApi.WebApiBase_Go.BASE_NAME_FLD
              || structInfo.Fields[0].Name == goscanner.ConvWebApi.WebApiBase_Go.BASE_NAME_FLD1 )
            {
                baseType = structInfo.Fields[0].Type;
            }else
            if (structInfo.Fields[0].IsPromoted)
            {
                baseType = structInfo.Fields[0].Type;
            }
            if( baseType != null)
            {
                if( baseType.Name == modelClassName )
                {
                    return true;
                }

                if( baseType.TypeClass == TypeClass.Struct
                 || baseType.TypeClass == TypeClass.Simple )
                {
                    var baseStruct = getStructInfoByName( baseType.Name );
                    if( baseStruct != null )
                    {
                        isInheritedBy( baseStruct, modelClassName, list1 );
                    }
                }
            }
        }
        return list1.Contains(modelClassName);
    }

    public void inspectModelsInvolvedInType( 
            bool bAdd,
            TypeInfo arg0Type, 
            List<StructInfo> involvedTypes, 
            List<StructInfo> models )
    {
        if(   arg0Type.Name == goscanner.ConvWebApi.WebApiBase_Go.CLASS_InGeneralData
           || arg0Type.Name == goscanner.ConvWebApi.WebApiBase_Go.CLASS_OutGeneralData )
           return;

        var structType = getStructInfoByName( arg0Type.Name);
        if( structType == null)
            return;
        

        foreach( var field in structType.Fields )
        {
            var genType = field.Type.getNormalizedType();
            if( field.IsPromoted && genType.Name == Options.ConvertWebApi.Class_Base_MODEL ) 
            {
                //var structTypeModel = getStructInfoByName( field.Type.Name);
                //if( structTypeModel != null )
                {
                    models.Add( structType );
                    bAdd  = false;//do not add in involvedType list
                }
            }else
            if( field.Type.TypeClass == TypeClass.Struct
             || field.Type.TypeClass == TypeClass.Simple)
            {
                inspectModelsInvolvedInType( true, field.Type, involvedTypes, models );
                

            }else
            if( field.Type.TypeClass == TypeClass.Slice)
            {
                var typeElem = field.Type.Clone();
                typeElem.TypeClass = TypeClass.Simple;
                typeElem.TypeName  = typeElem.Name;

                inspectModelsInvolvedInType( true, typeElem, involvedTypes, models );
            }else
            if( field.Type.TypeClass == TypeClass.Array )             
            {
                var arrType = field.Type as ArrayTypeInfo;
                if( arrType != null )
                {
                    inspectModelsInvolvedInType( true, arrType.TargetTypeInfo, involvedTypes, models );
                }

            }else
            if( field.Type.TypeClass == TypeClass.Map)
            {
                var arrType = field.Type as MapTypeInfo;
                if( arrType != null )
                {
                    if( arrType.KeyTypeInfo.Name != "string")
                    {
                        Console.WriteLine("Warning: Dictionary should be map[string]type ");
                    }

                    inspectModelsInvolvedInType( true, arrType.ElementTypeInfo, involvedTypes, models );
                }
            }
            if( field.Type.IsDerefPointer )
            {
                Console.WriteLine($"Warning: {field.Location}: Pointer is not allowed in json conversion");
            }

            CheckPrivateField( field );
        }

        if( bAdd)
        {
            involvedTypes.Add( structType );
        }
    }
    public static void CheckPrivateField(FieldInfo field)
    {
        if( field.Name.Substring(0,1).IsAllLower() )
        {
            Console.WriteLine($"Warning: {field.Location}: field '{field.Name}' is protected. It wont be exported in json conversion");
        }
    }

    static protected StreamWriter File_CreateText(string file)
    {
        //return File.CreateText( file);
        return new StreamWriter( new MemoryStream() );
    }

    protected string getIdentOfFunction(string fullname)
    {
        var parts = fullname.Split('.');
        return parts[ parts.Length - 1 ];
    }
    protected string getPackageAliasOfFunction(string fullname)
    {
        var parts = fullname.Split('.');
        if( parts.Length == 1 ) 
            return "";

        return parts[ 0];
    }
    protected string getPackageNameOfFunction(string fullname)
    {
        var packageAlias = getPackageAliasOfFunction( fullname);

        if( Metadata.ImportAliases.TryGetValue( packageAlias, out ( string, string) target ))
        {
            var packageUsing = target.Item1;
            return packageUsing;
            //var parts = packageUsing.Split("/");
            //return parts[ parts.Length - 1];
        }
        return null;
    }
    protected (string, string) getPackage_Function(string fullname)
    {
        var packageAlias = getPackageAliasOfFunction( fullname);

        if( Metadata.ImportAliases.TryGetValue( packageAlias, out ( string, string) target ))
        {
            var packageUsing = target.Item1;
            var parts = packageUsing.Split("/");
            return (parts[ parts.Length - 1], target.Item2);
        }
        return (null, null);
    }

    protected bool searchFunctionInImports( 
        string functionName, TPackageInfo packageInfo, 
        out FunctionInfo functionInfo, bool bShowDebug=false        
        )
    {
        functionInfo = null;

        //if( Metadata?.Functions.TryGetValue($"{functionName}()", out functionInfo) ?? false)
        if( Metadata?.Functions.TryGetValue($"{functionName}", out functionInfo) ?? false)
        {
            functionInfo.setMetaFile( Metadata );
            return true;
        }
        
        FolderMetadata metaImport;
        if( packageInfo.Metafile != null)
        {
            if( packageInfo.Metafile.Functions.TryGetValue($"{functionName}", out functionInfo))
            {
                functionInfo.setMetaFile(  packageInfo.Metafile );
                return true;
            }

            if( packageInfo.Metafile.ImportAliases.TryGetValue( packageInfo.Name, out (string targetImport, string targetUsing) alias2))
            {
                // here is the case when a type have the same package name like in the current, but the import is different
                // orm atomicsql
                // Users orm.DBTable
                // orm atomic_defs
                // ctx.Users.Qry() => aici era o eroare
                var warning = "";
                var metadata = LoadImportMetadata(Options, alias2.targetImport, out warning, bShowDebug);                
                //if( bShowDebug )
                //    Debug_Console(metadata != null? $"LoadMetadata: {alias2.targetImport} != null ": $"LoadMetadata: {alias2.targetImport} == null");
                if(metadata != null ) {

                    foreach( var itemMetadata in metadata.Files)
                    {
                        if( itemMetadata.Value.Functions.TryGetValue($"{functionName}", out functionInfo))
                        {
                            functionInfo.setMetaFile( itemMetadata.Value );
                            return true;
                        }
                    }
                }
            }
        }else        
        if( ImportMetadata.TryGetValue( packageInfo.Name, out metaImport))
        {
            foreach( var itemMetadata in metaImport.Files)
            {
                //if( itemMetadata.Value.Functions.TryGetValue($"{functionName}()", out functionInfo))
                if( itemMetadata.Value.Functions.TryGetValue($"{functionName}", out functionInfo))
                {
                    functionInfo.setMetaFile( itemMetadata.Value );
                    return true;
                }
            }
        }
        if( ImportAliases.TryGetValue( packageInfo.Name, out (string targetImport, string targetUsing) alias1))
        {
            // here is the case when a type have the same package name like in the current, but the import is different
            // orm atomicsql
            // Users orm.DBTable
            // orm atomic_defs
            // ctx.Users.Qry() => aici era o eroare
            var warning = "";
            var metadata = LoadImportMetadata(Options, alias1.targetImport, out warning);
            if(metadata != null ) {

                foreach( var itemMetadata in metadata.Files)
                {
                    if( itemMetadata.Value.Functions.TryGetValue($"{functionName}", out functionInfo))
                    {
                        functionInfo.setMetaFile( itemMetadata.Value );
                        return true;
                    }
                }
            }
        }
        

        if( Metadata != null )
        {
            foreach( var itemMetadata in Metadata.getFolderMetadata().Files)
            {
                if( itemMetadata.Value.Functions.TryGetValue($"{functionName}", out functionInfo))
                {
                    functionInfo.setMetaFile( itemMetadata.Value );
                    return true;
                }
            }
        }
        //functionInfo.setMetaFile( itemMetadata.Value );
        return false;
    }

    public bool searchStructInImports( string structName, TPackageInfo packageInfo, out StructInfo structInfo)
    {
        structInfo = null;

        //if( Metadata?.Functions.TryGetValue($"{functionName}()", out functionInfo) ?? false)
        if( Metadata?.Structs.TryGetValue(structName, out structInfo) ?? false)
            return true;

        FolderMetadata metaImport;
        if( packageInfo.Metafile != null)
        {
            var warning = "";
            var folder = GetFolderMetadata(Options, null, packageInfo.Metafile.SourceFileName, null);
            if( folder != null )
            {
                foreach( var itemMetadata in folder.Files)
                {
                    if( itemMetadata.Value.Structs.TryGetValue(structName, out structInfo))
                        return true;
                }
            }else
            {
                if( packageInfo.Metafile.Structs.TryGetValue(structName, out structInfo))
                    return true;
            }

            if( packageInfo.Metafile.ImportAliases.TryGetValue( packageInfo.Name, out (string targetImport, string targetUsing) alias1))
            {
                // here is the case when a type have the same package name like in the current, but the import is different
                // orm atomicsql
                // Users orm.DBTable
                // orm atomic_defs
                // ctx.Users.Qry() => aici era o eroare
                
                var metadata = LoadImportMetadata(Options, alias1.targetImport, out warning);
                if(metadata != null ) {

                    foreach( var itemMetadata in metadata.Files)
                    {
                        if( itemMetadata.Value.Structs.TryGetValue(structName, out structInfo))
                            return true;
                    }
                }
            }
        } 

        if( ImportMetadata.TryGetValue( packageInfo.Name, out metaImport))
        {
            foreach( var itemMetadata in metaImport.Files)
            {
                //if( itemMetadata.Value.Functions.TryGetValue($"{functionName}()", out functionInfo))
                if( itemMetadata.Value.Structs.TryGetValue(structName, out structInfo))
                    return true;
            }
        }
        
        if( Metadata != null )
        {
            foreach( var itemMetadata in Metadata.getFolderMetadata().Files)
            {
                if( itemMetadata.Value.Structs.TryGetValue(structName, out structInfo))
                    return true;
            }
        }

        if( ImportAliases.TryGetValue( packageInfo.Name, out (string targetImport, string targetUsing) alias))
        {
            var v = alias.targetImport;
        }
        return false;
    }

    protected string[] getIdentifierParts(string identifier)
    {
        return identifier.Split(".");
    }
    protected void Debug_Console(string msg){

        Console.WriteLine( msg );
    } 


    

    

}
