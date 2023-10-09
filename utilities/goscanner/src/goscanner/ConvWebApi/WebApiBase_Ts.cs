using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Diagnostics.Metrics;
using System.IO;
using System.Linq;
using System.Xml.Linq;
using Antlr4.Runtime.Misc;
using goscanner.AST;
using goscanner.Metadata;

namespace goscanner.ConvWebApi
{

    
internal class WebApiBase_Ts : WebApiBase_Go
{
public const string OMITEMIT = "omitempty";
public static string SvcObj_MODEL = "SvcObj_MODEL";

public static string SVCM = "SVCM";
public static string    Arg0 = "arg0";
//public static string VERSION_NUM = "102";


private static TypeInfo getReturnType( ConvWebApi scanner, FunctionInfo m )
{
    HashSet<string> inheritedTypeNames = new(StringComparer.Ordinal);
    goscanner.Metadata.ParameterInfo[] returnList = m.Signature.Signature.Result;

    if( returnList.Length == 1)
    {
        if( (returnList[0].Type.Name == "" || returnList[0].Type.Name == null) )
        { 
            Console.WriteLine("Method of SVC should have 1 return paramenter. this should be a struct type having a name");
            return null;
        }
        if( scanner.isInheritedBy( returnList[0].Type, CLASS_OutGeneralData, inheritedTypeNames ) )
        {
            return returnList[0].Type;
        }
                    
    }else
    if( returnList.Length == 2 )
    { 
        if( returnList[1].Type.Name != "error" )
        {
            if( scanner.isInheritedBy( returnList[0].Type, CLASS_OutGeneralData, inheritedTypeNames ) )
            {
                Console.WriteLine("Method of SVC should return : (Struct, error). ");
                return null;
            }
        }
        if( (returnList[0].Type.Name == "" || returnList[0].Type.Name == null) )
        { 
            Console.WriteLine("Method of SVC should have 1 return paramenter. this should be a struct type having a name");
            return null;
        }
        if( scanner.isInheritedBy( returnList[0].Type, CLASS_OutGeneralData, inheritedTypeNames ) )
        {
            return returnList[0].Type;
        }
    }else
    {
        //Console.WriteLine("Method of SVC should return : Struct or (Struct, error). ");
    }
    return null;
}
private static TypeInfo getArg0( ConvWebApi scanner, FunctionInfo m)
{
    HashSet<string> inheritedTypeNames = new(StringComparer.Ordinal);

    ParameterInfo[] returnList = m.Signature.Signature.Parameters;
    if( returnList.Length == 1)
    {
        if( (returnList[0].Type.Name == "" || returnList[0].Type.Name == null) )
        { 
            Console.WriteLine("Method of SVC should have 1 input paramenter. this should be a struct type having a name");
            return null;
        }
        if( scanner.isInheritedBy( returnList[0].Type, CLASS_InGeneralData, inheritedTypeNames ) )
        {
            return returnList[0].Type;
        }
    }
    {
        //return null;
        //Console.WriteLine("Method of SVC should return : Struct or (Struct, error). ");
    }
    return null;
}
public static string exportSvc_AllMethodsTs( ConvWebApi scanner, List<FunctionInfo> methods, string svcModelFiles)
{
    var meths = "";
    foreach (var m in methods)
    {
        var returnType = getReturnType( scanner, m );
        var arg0Type   = getArg0( scanner, m );
        if( returnType != null && arg0Type != null )
        {
            var methTxt = Export_Method(m.Signature.Name, cleanNameGoStruct( returnType.Name), arg0Type );
            meths += methTxt;
        }
    }           

    return meths;
    //return Export_Service("Svc", "Svc",  meths, svcModelFiles);
}

public static string exportSvcModels_AllMethodsTs(
    ConvWebApi scanner,
    List<StructInfo> models__, 
    List<FunctionInfo> methods,
    Dictionary<string, string>  dictCollectionFields,
    Dictionary<string, string>  dictExportedModel )
{
    var meths = "";
    var NameSpacePrefix = "";//SvcM
    var involvedTypes = new List<StructInfo>();
    var models        = new List<StructInfo> ();
    //se exporta prima data tipurile In si Out de la metodele intalnite in 
    // 
    foreach (var m in methods)
    {
        var returnType = getReturnType( scanner, m );
        if( returnType != null )
        {
            if( !dictExportedModel.ContainsKey(returnType.Name) )
            {
                dictExportedModel[ returnType.Name ] = returnType.Name;

                scanner.inspectModelsInvolvedInType( false, returnType, involvedTypes, models );

                var nameStruct = cleanNameGoStruct( returnType.Name );
                meths += Export_ArgType( scanner, nameStruct, NameSpacePrefix );
            }
        }else
            continue;

        var arg0Type = getArg0( scanner, m );
        if( arg0Type != null )
        {
            if( !dictExportedModel.ContainsKey(arg0Type.Name) )
            {
                dictExportedModel[ arg0Type.Name ] = arg0Type.Name;

                scanner.inspectModelsInvolvedInType( false, arg0Type, involvedTypes, models );
                var nameStruct1 = cleanNameGoStruct( arg0Type.Name );
                meths += Export_ArgType( scanner, nameStruct1, NameSpacePrefix );
            }
        }
    }

    //exporta structurile ne modele
    foreach (var m in involvedTypes)
    {
        if( dictExportedModel.ContainsKey(m.Name) )
            continue;
        dictExportedModel[ m.Name ] = m.Name;

        var nameStruct1 = cleanNameGoStruct( m.Name );
        meths += Export_ArgType( scanner, nameStruct1, NameSpacePrefix );
    }
            

    //apoi modelele implicate in export
    foreach (var m in models)
    {
        //evita sa export de 2 ori acelasi model
        if( dictExportedModel.ContainsKey(m.Name) )
            continue;
        dictExportedModel[ m.Name ] = m.Name;

        var readData = "";
        var fields = "";
        TypeInfo unusedBase = null;

        generateModelsFieldsAndRead( scanner, NameSpacePrefix, m.Fields, ref fields, ref readData, ref unusedBase, 
                        cleanNameGoStruct(m.Name), dictCollectionFields);
                
        var methTxt =  Export_Models(m.Name, fields, readData);

        meths += methTxt;
    }

    return meths;

    //return Export_SvcModels( "", "", meths);
}

private bool isInheritClass(TypeInfo typeclass, string nameClassBase)
{
    var structBase = ConvWebApi.getStructInfoByName( typeclass.Name );
    if( structBase == null ) 
        return false;

    return isInheritClass(structBase, nameClassBase);
}
private bool isInheritClass(StructInfo typeClass, string nameClassBase)
{
    var infoFields = typeClass.Fields;

    foreach ( var fld in infoFields)
    {
        if( fld.Description != null && fld.Description.Contains( OMITEMIT ) )
        {
            //it is included in the definition
            continue;
        }else
        if( fld.Name == FLD_ID )
        {
            //it is included in the definition
            continue;
        }else
        {
            var baseClass = fld.Type;
            if( !( baseClass.TypeClass == TypeClass.Simple
                || baseClass.TypeClass == TypeClass.Interface) )
            { 
                return false;
            }
            if( fld.IsPromoted ) 
            {
                var structBase = ConvWebApi.getStructInfoByName( fld.Type.Name );
                if( structBase != null )
                {
                    return isInheritClass(structBase, nameClassBase);
                }
            }
                    
            if( (fld.Name == nameClassBase && fld.IsPromoted) )
            {
                return true;
            }
        }
    }
    return false;
}


        
        

private static void generateModelsFieldsAndRead(            
    ConvWebApi scanner,
    string NameSpacePrefix,
    FieldInfo[] infoFields, 
    ref string fields, ref string readData, 
    ref TypeInfo baseClass, 
    string modelTypeName, 
    Dictionary<string, string>  dictCollectionFields)
{
    //TypeInfo baseClass = null;
    foreach (var fld in infoFields)
    {
        if( fld.Description != null && fld.Description.Contains( OMITEMIT ) )
        {
            //it is included in the definition
            continue;
        }else
        if( fld.Name == FLD_ID )
        {
            //it is included in the definition
            continue;
        }else
        if( fld.Name == BASE_NAME_FLD
            || fld.Name == BASE_NAME_FLD1 
            || (fld.Name == CLASS_OutGeneralData && fld.IsPromoted) 
            || (fld.Name == CLASS_InGeneralData && fld.IsPromoted)
            || (fld.Name == Class_Generic_MODEL && fld.IsPromoted)
            )
        {
            baseClass = fld.Type;
            if( !( baseClass.TypeClass == TypeClass.Simple
                || baseClass.TypeClass == TypeClass.Interface
                || baseClass.TypeClass == TypeClass.Struct) )
            { 
                        
                Console.WriteLine($"{fld.Location}: Base class for input/ouput arguments should be simple. not pointers");
            }
                    
            if(  (fld.Name == CLASS_OutGeneralData && fld.IsPromoted) 
              || (fld.Name == CLASS_InGeneralData && fld.IsPromoted))
            {
                //ii permit sa expuna si VersX si Res in modelul de In Out pt fiecare clasa
                var structBase = ConvWebApi.getStructInfoByName( baseClass.Name );
                if( structBase != null )
                {
                    TypeInfo unusedBase = null;
                    generateModelsFieldsAndRead( scanner, NameSpacePrefix, structBase.Fields, ref fields, ref readData, ref unusedBase, modelTypeName, dictCollectionFields);
                }
            }else
            if( baseClass.Name == CLASS_OutGeneralData 
                || baseClass.Name == CLASS_InGeneralData )
            {
            }else
            {
                //check if is derivated from CLASS_OutGeneralData
                
                var structBase = ConvWebApi.getStructInfoByName( baseClass.Name );
                if( structBase != null )
                {
                    TypeInfo unusedBase = null;
                    generateModelsFieldsAndRead( scanner, NameSpacePrefix, structBase.Fields, ref fields, ref readData, ref unusedBase, modelTypeName, dictCollectionFields);
                }else
                if( baseClass.TypeClass != TypeClass.Interface )
                {
                    var genType = baseClass.getNormalizedType();
                    if( scanner.searchStructInImports( genType.Name, genType.PackageInfo, out structBase))
                    {
                        TypeInfo unusedBase = null;
                        generateModelsFieldsAndRead( scanner, NameSpacePrefix, structBase.Fields, ref fields, ref readData, ref unusedBase, modelTypeName, dictCollectionFields);
                    }else
                    {
                        if( genType.Name == Class_Generic_MODEL)
                        {
                        }else
                        {
                            Console.WriteLine( $"internal error : not found struct {baseClass.Name}. ");
                        }
                    }
                }
            }
        }else
        if( fld.IsPromoted ) 
        {
            var structInfo = ConvWebApi.getStructInfoByName( fld.Type.Name );
            if( structInfo != null ) 
            {
                var infoFields1 = structInfo.Fields;
                generateModelsFieldsAndRead( scanner,           
                    NameSpacePrefix, infoFields1, 
                    ref  fields, ref  readData, ref baseClass, 
                    modelTypeName, dictCollectionFields);
            }else
            {
                Console.WriteLine($"Promoted Field should be only struct");
            }
        }else
        {
            if( dictCollectionFields != null)
            {
                //dictCollectionFields[$"{modelTypeName}_{fld.Name}"] = $"{modelTypeName}.{fld.Name}";
                // it must be the same as in class generation
                dictCollectionFields[$"{modelTypeName}_{fld.Name}"] = $"{fld.Name}";
            }
            var fldTxt = Export_Model_Fields( NameSpacePrefix, fld.Name, getJsonName(fld), fld.Type );
            checkText(fldTxt);
            fields += fldTxt;

            var readTxt = Export_Model_Read( NameSpacePrefix, fld.Name, fld.Type);
            checkText(readTxt);
            readData += readTxt;
        }
    }
}

private static string getJsonName(FieldInfo fld)
{
    if( fld.Description != null )
    {
        //`json:"id,omitempty"`
        var prefixJson  = "json:\"";
        if( fld.Description.StartsWith(prefixJson ) )
        {
            var name1 = fld.Description.Substring( prefixJson.Length );
            var parts = name1.Split(",");
            return parts[0];
        }
    }
    return fld.Name;
}

private static void generateAnonymousFields(FieldInfo[] infoFields, ref List<string> fields, ref List<string> listCheckingArg0Members)
{
    TypeInfo baseClass = null;
    var bFirst = true;

    foreach (var fld in infoFields)
    {
        if( fld.Name == BASE_NAME_FLD || (fld.Name == CLASS_OutGeneralData && fld.IsPromoted) 
            || fld.Name == BASE_NAME_FLD1|| (fld.Name == CLASS_InGeneralData && fld.IsPromoted)  )
        {
            baseClass = fld.Type;
            var baseClassName = baseClass.Name;

            if( fld.IsPromoted )
            {
                if(  (fld.Name == CLASS_OutGeneralData && fld.IsPromoted) 
                    || (fld.Name == CLASS_InGeneralData  && fld.IsPromoted)  ){
                    baseClassName = fld.Name;
                } 
            }
            if( !(baseClass.TypeClass == TypeClass.Simple || baseClass.TypeClass == TypeClass.Struct) )
            { 
                Console.WriteLine($"{fld.Location}: Base class for input/ouput arguments should be simple. not pointers");
            }

            if( baseClassName == CLASS_InGeneralData && InlineAllFields_ClassInGeneralData )
            {
                //I ignore the fields of InGeneralData because they are implicit arguments 
            }else
            if( baseClassName == CLASS_OutGeneralData 
                || baseClassName == CLASS_InGeneralData )
            {
                var structBase = ConvWebApi.getStructInfoByName( baseClassName );
                generateAnonymousFields( structBase.Fields, ref fields, ref listCheckingArg0Members);
            }else
            {
                //check if is derivated from CLASS_OutGeneralData
                var structBase = ConvWebApi.getStructInfoByName( baseClassName );
                generateAnonymousFields( structBase.Fields, ref fields, ref listCheckingArg0Members);
            }
        }else
        {
            var fldTxt = $"{fld.Name}: { ConvertTypeToTs(SVCM, fld.Type) }";
            if( bFirst)
            {
                fields.Add( fldTxt );
                bFirst = false;
            } else
            {
                fields.Add( fldTxt );
            }

            addCodeCheckingArg0( fld, ref listCheckingArg0Members );
        }
    }
}

private static void addCodeCheckingArg0(FieldInfo fld, ref List<string> listCheckingArg0Members)
{
    if( fld.Name == "IndicesOfSelectItems" )
        Utils.Nop();

    var prefixArr = "";
    if( fld.Type.TypeName.StartsWith("slice<")
     || fld.Type.TypeName.StartsWith("array<"))
    {
        prefixArr = "Ajax.Arr_";
    }

            //checking arg0
    if ( fld.Type.Name == "string"    ||fld.Type.Name == "NullString"
        ||fld.Type.Name == "uuid.UUID" || fld.Type.Name == "UUID" 
        ||fld.Type.Name == "time"      ||fld.Type.Name == "sql.NullTime")
    {
        listCheckingArg0Members.Add($"{Arg0}.{fld.Name} = {prefixArr}S({Arg0}.{fld.Name})");
    }else
    if ( fld.Type.Name == "bool"
        ||fld.Type.Name == "sql.NullBool" )
    {
        listCheckingArg0Members.Add($"{Arg0}.{fld.Name} = {prefixArr}B({Arg0}.{fld.Name})");
    }else
    if (fld.Type.Name == "int"
        || fld.Type.Name == "byte"            
        || fld.Type.Name == "uint8"            
        || fld.Type.Name == "int16"
        || fld.Type.Name == "int32"
        || fld.Type.Name == "int64"
    || fld.Type.Name == "sql.NullInt"
    || fld.Type.Name == "sql.NullByte"            
    || fld.Type.Name == "sql.NullInt16"
    || fld.Type.Name == "sql.NullInt32"
    || fld.Type.Name == "sql.NullInt64" )
    {
        listCheckingArg0Members.Add($"{Arg0}.{fld.Name} = {prefixArr}I({Arg0}.{fld.Name})");
    }else
    if (   fld.Type.Name == "float32"
        || fld.Type.Name == "float64"
        || fld.Type.Name == "sql.NullFloat32"
        || fld.Type.Name == "sql.NullFloat64"
        )
    {
        listCheckingArg0Members.Add($"{Arg0}.{fld.Name} = {prefixArr}F({Arg0}.{fld.Name})");
    }
}


private static string Export_ArgType( 
    ConvWebApi convWebApi,
    string nameStruct, string NameSpacePrefix )
{
    var meths = "";
    var typeStruct =  ConvWebApi.getStructInfoByName(nameStruct);
    if( typeStruct != null )
    {
        var readData = "";
        var fields = "";
        TypeInfo baseClass = null;

        generateModelsFieldsAndRead( convWebApi, NameSpacePrefix, typeStruct.Fields, ref fields, ref readData, ref baseClass, "", null );

        if( baseClass != null )
        {
            var methTxt = Export_ArgModels(typeStruct.Name, fields, readData, baseClass);
            meths += checkText(methTxt);
        }else
        {
            var methTxt = Export_ArgModels1(typeStruct.Name, fields, readData, SvcObj_MODEL );
            meths += checkText(methTxt);
                    
            //Console.WriteLine($"Base struct for '{typeStruct.Name}'  not exist. must declare `Base StructBase;`");
        }
    }
    return meths;
}

private static TypeInfo getTypeDef(TypeInfo type1 )
{
    if( ConvWebApi.s_AllTypeDefs.ContainsKey(type1.Name) )
    {
        var type = ConvWebApi.s_AllTypeDefs[type1.Name];
        var t1 = getTypeDef( type );
        return t1;
    }
    return type1;
}

        

public static string ConvertTypeToTs(string prefixSvc,goscanner.Metadata.TypeInfo typeGen)
{
    //if (typeGen.TypeClass == TypeClass.Simple)
    { 
        typeGen = getTypeDef( typeGen );
    }
    return convertTypeToTs( prefixSvc, typeGen);
}

public static string convertTypeToTs(string prefixSvc,goscanner.Metadata.TypeInfo typeGen)
{
    if (typeGen.TypeClass == TypeClass.Simple)
    { 
        if ( typeGen.Name == "string")
            return "string";

        if ( typeGen.Name == "uuid.UUID" || typeGen.Name == "UUID")
            return "string";

        if (typeGen.Name == "sql.NullBool")
            return "boolean";
        if (typeGen.Name == "bool")
            return "boolean";

        if (typeGen.Name == "sql.NullTime")
            return "Date";
        if (typeGen.Name == "time")
            return "Date";

        if( typeGen.Name == "sql.NullInt32")
            Utils.Nop();

        if( typeGen.Name == "uint8")
            return "number";

        if (typeGen.Name == "int"
            || typeGen.Name == "byte"            
            || typeGen.Name == "uint8"            
            || typeGen.Name == "int16"
            || typeGen.Name == "int32"
            || typeGen.Name == "int64"
            || typeGen.Name == "float32"
            || typeGen.Name == "float64"
            || typeGen.Name == "sql.NullInt"
            || typeGen.Name == "sql.NullByte"            
            || typeGen.Name == "sql.NullInt16"
            || typeGen.Name == "sql.NullInt32"
            || typeGen.Name == "sql.NullInt64"
            || typeGen.Name == "sql.NullFloat32"
            || typeGen.Name == "sql.NullFloat64"
            )
            return "number";

        if (typeGen.Name == "sql.NullString")
            return "string";
        if (typeGen.Name == "string")
            return "string";
    }

    if (typeGen.TypeClass == TypeClass.Array)
    {
        var arrType = typeGen as ArrayTypeInfo;
        var elemType = ConvertTypeToTs(prefixSvc, arrType.TargetTypeInfo);
        return $"Array<{elemType}>";
    }
    if( typeGen.TypeClass == TypeClass.Slice)
    {
        var typeElem = typeGen.Clone();
        typeElem.TypeClass = TypeClass.Simple;
        typeElem.TypeName  = typeElem.Name;

        var elemType = ConvertTypeToTs(prefixSvc, typeElem);
        return $"Array<{elemType}>";
    }

    if (typeGen.TypeClass == TypeClass.Map)
    {
        var arrType = typeGen as MapTypeInfo;

        if(arrType.KeyTypeInfo.Name != "string")
            Console.WriteLine("Map type should be only [string] Struct");

        var elemType = ConvertTypeToTs(prefixSvc, arrType.ElementTypeInfo);
        return $"DictT<{elemType}>";
    }
    if (typeGen.TypeClass == TypeClass.Interface
     || typeGen.TypeClass == TypeClass.Simple
     || typeGen.TypeClass == TypeClass.Struct)
    {
        return (prefixSvc!="" ? prefixSvc + "." : "") + cleanNameGoStruct(typeGen.Name);
    }
    if (typeGen.TypeClass == TypeClass.Function)
    {
        Console.WriteLine( $"not allowed Function type in svc export {typeGen.Name}");
        return cleanNameGoStruct(typeGen.Name);
    }

    Console.WriteLine($"cannot translate type '{typeGen.Name}'");
    return "string";
}


public static string Export_Service( Options options,
            string ajax_service, string nameClass,
            string methods,
            string SvcModelFile )
{

    var fiSvcModelFile = new FileInfo(SvcModelFile);
    var SvcModelFileName = fiSvcModelFile.Name.Replace(fiSvcModelFile.Extension, "");
    //SvcModelFile = SvcModels.gen

var txt = $@"

    import * as Ajax  from  '../services/AjaxScanTraffic'
    import * as {SVCM}  from  './{SvcModelFileName}'
    type TElement=any
    
    
    /** this file is generated automatically by go_tool.exe */
    export type SVC_FnError = ( err: string ) => void;
    
    export class {nameClass}
    {{
        public static readonly ExportVersion:number = 102;
        public static readonly ExportVersionName:string = 'VersX';
        public static readonly CmdID:string = 'Msg';// it is the field from service that make the dispatch
        public static readonly SessionUID:string = 'SessionUID';
        
        constructor()
        {{
        }}

        {methods}
    }}
    export var SVCInst = new Svc();

    ";

    txt = UseTemplate( txt, options.ConvertWebApi.Templ_TsExportService,
                new Dictionary<string,string>()
                { 
                    { "SVCM", SVCM},
                    { "SvcModelFileName", SvcModelFileName},
                    { "nameClass", nameClass},
                    //{ "VERSION_NUM", VERSION_NUM},
                    { "methods", methods},
                }
                );

     return checkText( txt );
}

private static string computeAnonymousDef(goscanner.Metadata.TypeInfo type, ref List<string> listCheckingArg0Members )
{
    var fields = new List<string>();
    var structType = ConvWebApi.getStructInfoByName( type.Name );
    if( structType == null && type.TypeClass != TypeClass.Interface)
    {
        Console.WriteLine( $"internal error for struct Type {type.Name}" );
        return "";
    }

    generateAnonymousFields( structType.Fields, ref fields, ref listCheckingArg0Members);
    return String.Join(", ", fields.ToArray());
}
/// <summary>
/// Mines_getGame
/// </summary>
public static string Export_Method(string methodName, string methodOutType, TypeInfo methodInputType )
{
    var listCheckingArg0Members = new List<string>();
    var anonymousDef = computeAnonymousDef( methodInputType, ref listCheckingArg0Members );

    var checking_arg0 =  String.Join(  
            @"
            ", listCheckingArg0Members);

    var text = $@"
        //----------------------------------------------------------------------------------
        {methodName}( {Arg0}:{{ {anonymousDef} }},
            fnSuccess:(res:SVCM.{methodOutType})=>void, fnError:SVC_FnError )
        {{
            {checking_arg0}

            Ajax.AjaxSoap_Send( {{
                methodName:'{methodName}',                
                instance:new SVCM.{methodOutType}(),                
                args: {Arg0}, 
                fnSuccess: function(ret:any){{ fnSuccess( ret as SVCM.{methodOutType}); }}, 
                fnError: fnError
                }}
            )
        }}
    ";

    return checkText( text );
}


private static string generateEnumFromDict(Dictionary<string, string> dictCollectionFields) 
{ 
    var txt = "";
    foreach( var it in dictCollectionFields)
    {
        txt += @$"{it.Key}='{it.Value}',
                ";
    }
    return checkText( txt );
}

public static string Export_SvcModels( Options options, string models, Dictionary<string, string> dictCollectionFields)
{

    var export_enumsLst = new List<string>();
    foreach( var it in options.ConvertWebApi.DictExportEnums ) 
    {
        var enumCode = generateEnumFromDict( it.Value );
        var enumName = it.Key;

        var enumDecl = $@"

        export enum {it.Key} {{
            {enumCode}
        }}
        ";
        export_enumsLst.Add( enumDecl );
    }
    var export_enums =  String.Join( "", export_enumsLst );
    //SvcModelFile = SvcModels.gen

    var _CollectionFields = generateEnumFromDict( dictCollectionFields);
    var txt = $@"

    
    /** this file is generated automatically by go_tool.exe */
    
    import * as Ajax from  '../services/AjaxScanTraffic'
    import * as {SVCM}  from  './AjaxExport'
    import {{UtilsBase,S,F,I}} from '../../_framework/utils/UtilsBase'
    import {{{SvcObj_MODEL}, {Class_Generic_MODEL} }} from '../../_framework/common/Generic_MODEL'
    
    type TElement=any
              
    export type clsGenericItem = {Class_Generic_MODEL};//dtoGenericItem;//replace with something
        
    export class dtoGenericItem  extends {SvcObj_MODEL}
    {{
        public constructor( data: any=null)
        {{
            super(data);
        }}
    }}

    export class AJX_classifList  extends {SvcObj_MODEL}
    {{
        public Contents(): HashMap< Array< clsGenericItem> > {{ return null; }};
        public Titles():   HashMap<String> {{ return null; }};
    
        public constructor( data: any=null)
        {{
            super(data);
        }}
    }}

    export class {CLASS_InGeneralData}  extends {SvcObj_MODEL}
    {{
        public constructor( data: any=null)
        {{
            super(data);
        }}
    }}
    
    export class {CLASS_OutGeneralData}  extends {SvcObj_MODEL}
    {{
        public constructor( data: any=null)
        {{
            super(data);
        }}
    }}
    
    
    export class AJX_ComboDynamic extends {CLASS_OutGeneralData}
    {{
        get options_db(): Array<{SVCM}.dtoGenericItem>  {{ return this.__data['options_db'] as Array<{SVCM}.dtoGenericItem>;}}
        set options_db(v: Array<{SVCM}.dtoGenericItem>) {{        this.__data['options_db'] = v;}} 
        get field_id(): string  {{ return this.__data['field_id'] as string;}}
        set field_id(v: string) {{        this.__data['field_id'] = v;}} 
        get field_title(): Array<string>  {{ return this.__data['field_title'] as Array<string>;}}
        set field_title(v: Array<string>) {{        this.__data['field_title'] = v;}} 
    
        readData(name:string, el:TElement):void
        {{
            el = Ajax._parseFindRoot(name, el);
            
            this.options_db = Ajax._parseArray<{SVCM}.dtoGenericItem >( 'options_db', el, function(el:TElement){{ return ( function( el: Element ) {{ var i:{SVCM}.dtoGenericItem = new {SVCM}.dtoGenericItem(); i.readData('options_db', el ); return i; }})(el); }} );
            this.field_id = Ajax._parseString( 'field_id', el );
            this.field_title = Ajax._parseArray<string>( 'field_title', el, function(el:TElement){{ return S(el) ); }} );
        }}
    }}
    
    {models}
        
    export enum COL_ID {{

        {_CollectionFields}
    }}

    {export_enums}
    ";


    txt = UseTemplate( txt, options.ConvertWebApi.Templ_TsExportModels,
                new Dictionary<string,string>()
                { 
                    { "Generic_MODEL", Class_Generic_MODEL},
                    { "SvcObj_MODEL", SvcObj_MODEL},
                    { "CLASS_InGeneralData", CLASS_InGeneralData},
                    { "CLASS_OutGeneralData", CLASS_OutGeneralData},
                    { "SVCM", SVCM},
                    { "models", models},
                    { "_CollectionFields", _CollectionFields},
                    { "export_enums", export_enums},
                    
                }
                );

    return checkText( txt );
}

public static string ConvertFunc(string typeTsName)
{
    switch( typeTsName )
    {
        case "string": return "S";
        case "boolean": return "B";
        case "number": return "F";
        case "Date":   return "";
    }
    return "";
}
public static string Export_Model_Fields(string prefixSVC, string fieldName, string fieldJsonName, TypeInfo type)
{
    if( ConvWebApi.s_currentScanner.SourceFileName.Contains("WheelGame.go"))
    {
        if( fieldName == "Items")
            Utils.Nop();
    }

    var typeTsName = ConvertTypeToTs(prefixSVC, type);

    var convertFunc = ConvertFunc(typeTsName);

    var text = $@"

        get {fieldName}(): {typeTsName}  {{ return {convertFunc}(this.__data['{fieldJsonName}']) as {typeTsName};}}
        set {fieldName}(v: {typeTsName}) {{        this.__data['{fieldJsonName}'] = v;}} 
    ";
    return checkText( text );
}

public static bool isPrimaryType(goscanner.Metadata.TypeInfo typeGen)
{
    if (typeGen.Name == "string")
        return true;

    if (typeGen.Name == "bool")
        return true;

    if (typeGen.Name == "time")
        return true;

    if (typeGen.Name == "int"
     || typeGen.Name == "byte"
     || typeGen.Name == "uint8"        
     || typeGen.Name == "int16"
     || typeGen.Name == "int32"
     || typeGen.Name == "int64"
     || typeGen.Name == "float32"
     || typeGen.Name == "float64")
         return true;

    return false;
}

public static string Export_Model_Read_ArrayElem(string prefixSVC, string fieldName, TypeInfo typeGen)
{ 
    if( (typeGen.TypeClass == TypeClass.Interface
     || typeGen.TypeClass == TypeClass.Simple
     || typeGen.TypeClass == TypeClass.Struct
     )
     && !isPrimaryType(typeGen) )
    {
        string elemTypeName = cleanNameGoStruct( typeGen.Name );

        return $"(function(el: Element) {{ var i:{elemTypeName} = new {elemTypeName}(); i.readData('{fieldName}', el); return i;}})(el)";
    }else
    {
        return Export_Model_Read1( prefixSVC, null, typeGen);
    }
}
public static string Export_Model_Read(string prefixSVC, string fieldName, goscanner.Metadata.TypeInfo typeGen)
{
    var txt = Export_Model_Read1( prefixSVC, fieldName, typeGen);
    return $@"this.{fieldName} = {txt};
            ";
}
public static string Export_Model_Read1(string prefixSVC, string fieldName, TypeInfo typeGen)
{
    if(typeGen.Name == "TMoney")
        Utils.Nop();

    typeGen = getTypeDef( typeGen );

    var strfieldName = fieldName != null? $"'{fieldName}'" : "null";

    var typeTsName = ConvertTypeToTs( prefixSVC, typeGen);
    if(typeTsName == "User"
       ||typeTsName == "UserRole" )
        Utils.Nop();

    if (typeGen.TypeClass == goscanner.Metadata.TypeClass.Simple)
    {
        if (typeGen.Name == "string"
         || typeGen.Name == "sql.NullString"
         )
            return $"Ajax._parseString( {strfieldName}, el )";

        if (typeGen.Name == "bool"         
         || typeGen.Name == "sql.NullBool")
            return $"Ajax._parseBool( {strfieldName}, el )";

        if (typeGen.Name == "uuid.UUID" || typeGen.Name == "UUID")
            return $"Ajax._parseString( {strfieldName}, el )";

        if (typeGen.Name == "time"
         || typeGen.Name == "time.Time"
         
         || typeGen.Name == "sql.NullTime")
            return $"Ajax._parseDate( {strfieldName}, el )";

        if ( typeGen.Name == "uint8")
            return $"Ajax._parseNumber( {strfieldName}, el )";

        if (   typeGen.Name == "int"
            || typeGen.Name == "byte"            
            || typeGen.Name == "uint8"
            || typeGen.Name == "int16"
            || typeGen.Name == "int32"
            || typeGen.Name == "int64"
            || typeGen.Name == "sql.NullInt"
            || typeGen.Name == "sql.NullByte"            
            || typeGen.Name == "sql.NullInt16"
            || typeGen.Name == "sql.NullInt32"
            || typeGen.Name == "sql.NullInt64"            
            )
            return $"Ajax._parseNumber( {strfieldName}, el )";

        if (   typeGen.Name == "float32"
            || typeGen.Name == "float64"
            || typeGen.Name == "sql.NullFloat32"
            || typeGen.Name == "sql.NullFloat64"
            )
            return $"Ajax._parseFloat( {strfieldName}, el )";
    }

    if (typeGen.TypeClass == TypeClass.Array)
    {
        var arrType = typeGen as ArrayTypeInfo;
        var elemType = ConvertTypeToTs(prefixSVC, arrType.TargetTypeInfo);
        var elemFunc =  Export_Model_Read_ArrayElem( prefixSVC, fieldName, arrType.TargetTypeInfo);

        return $"Ajax._parseArray<{elemType}>('{fieldName}', el, function(el: TElement){{ return {elemFunc}; }} )";
    }
    if (typeGen.TypeClass == TypeClass.Slice)
    {
        var typeElem = typeGen.Clone();
        typeElem.TypeClass = TypeClass.Simple;
        typeElem.TypeName  = typeElem.Name;
        var elemType = ConvertTypeToTs(prefixSVC, typeElem);
        var elemFunc =  Export_Model_Read_ArrayElem( prefixSVC, fieldName, typeElem);

        return $"Ajax._parseArray<{elemType}>('{fieldName}', el, function(el: TElement){{ return {elemFunc}; }} )";
    }
    if (typeGen.TypeClass == TypeClass.Map)
    {
        var arrType = typeGen as MapTypeInfo;
        if (arrType.KeyTypeInfo.Name != "string")
            Console.WriteLine("Map type should be only [string] Struct");

        var elemType = ConvertTypeToTs(prefixSVC, arrType.ElementTypeInfo);
        var elemFunc =  Export_Model_Read_ArrayElem( prefixSVC, fieldName, arrType.ElementTypeInfo);

        return $"Ajax._parseDict<{elemType}>('{fieldName}', el, function(el: TElement){{ return {elemFunc}; }} )";
    }

    
    if( typeGen.TypeClass == TypeClass.Interface
     || typeGen.TypeClass == TypeClass.Simple
     || typeGen.TypeClass == TypeClass.Struct)
    {
        return $@"new {typeTsName}();
            this.{fieldName}.readData( name, el);
            ";
          //          $"Ajax._parseDict<{elemType}>('{fieldName}', el, function(el: TElement){{ return {elemFunc}; }} )";
        //return typeTsName;
    }
    
    return "/* error conversion*/";
}
/*
        /// <summary>
        /// Mines_getGame
        /// </summary>
public static string Export_MethodArgs(string methodName, string methodOutType, string fields, string readData )
{
    var text = $@"
        //----------------------------------------------------------------------------------
    export class {methodOutType} extends {CLASS_OutGeneralData}
    {{
        {fields}
    
        readData(name:string, el:TElement):void
        {{
            el = Ajax._parseFindRoot(name, el);
            
            {readData}
        }}
    }}
    ";

    return text;
}
*/

public static string Export_ArgModels(string ModelName, string fields, string readData, TypeInfo baseClass)
{
    var baseClassName = cleanNameGoStruct( baseClass.Name );
    return Export_ArgModels1( ModelName, fields, readData, baseClassName);
}
public static string Export_ArgModels1(string ModelName, string fields, string readData, string baseClassName)
{

    var text = $@"
        //----------------------------------------------------------------------------------
    export class {ModelName} extends {baseClassName}
    {{
        public constructor( data: any = null ){{super(data);}}
    
        {fields}
    
        readData(name:string, el:TElement):void
        {{
            el = Ajax._parseFindRoot(name, el);
            
            {readData}
        }}
    }}
    ";

    return text;
}

public static string Export_Models(string ModelName, string fields, string readData)
{
    var text = $@"
        //----------------------------------------------------------------------------------
    export class {ModelName} extends {Class_Generic_MODEL}
    {{
        public constructor( data: any = null ){{super(data);}}
    
        public  get ID(): string {{return ''+(this.__data['ID']  as String); }};
        public  set ID(s: string) {{this.__data['ID'] = s; }};


        {fields}
    
        readData(name:string, el:TElement):void
        {{
            el = Ajax._parseFindRoot(name, el);
            
            {readData}
        }}
    }}
    ";

    return checkText(text);
    }


    public static string checkText(string s)
    {
        if (s.Contains("this.Time1.readData( name, el);"))
            Debugger.Break();

        return s;
    }







    }
}

