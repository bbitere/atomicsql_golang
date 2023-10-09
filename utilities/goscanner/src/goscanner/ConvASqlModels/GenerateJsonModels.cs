using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Diagnostics.Metrics;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Xml.Linq;
using Antlr4.Runtime.Misc;
using Dahomey.Json;
using goscanner.AST;
using goscanner.ConvSql;
using goscanner.Metadata;

namespace goscanner.ConvASqlModels
{

    
internal class GenerateJsonModels 
{

const string TAG_Type ="Type";
const string TAG_tags ="tags";
const string TABLE_SQLNAME ="_tblsql";
        

public const string OMITEMIT = "omitempty";
public static string SvcObj_MODEL = "SvcObj_MODEL";

public static string SVCM = "SVCM";
public static string    Arg0 = "arg0";
//public static string VERSION_NUM = "102";


public static string FLD_ID = "ID";
public static string BASE_NAME_FLD="Generic_MODEL";
public static string BASE_NAME_FLD1 ="Generic_MODEL";
public static string Class_Generic_MODEL ="Generic_MODEL";
public static string  DELIMETER="@@@@@@@@######@@@@@@";
        

public static string cleanNameGoStruct(string name)
{ 
    //I renmove the pointer name
    var parts = name.Split(".");
    if( parts.Length > 0 )
    {
        return parts[ parts.Length -1].Replace("*","");
    }
    return name.Replace("*","");
}

public static string exportJsonModels(
    ConvASqlModels scanner,
    List<StructInfo> models, 
    Dictionary<string, string>  dictCollectionFields,
    Dictionary<string, string>  dictExportedModel )
{
    var jsonOut = "";
          

    //apoi modelele implicate in export
    foreach (var model in models)
    {
        //evita sa export de 2 ori acelasi model
        if( dictExportedModel.ContainsKey(model.Name) )
            continue;
        dictExportedModel[ model.Name ] = model.Name;
                
        if( !model.Comment.Trim().StartsWith($"/*{OrmDef.Atomicsql_table}\""))
        {
            Console.WriteLine($"Ignore to export struct {model.Name}. not found marker atomicsql-table ");
            continue;
        }

        var dict = new Dictionary<string, Dictionary<string, string> >();
        dict[TABLE_SQLNAME] = new Dictionary<string, string>
                {
                    { TAG_Type, "" },
                    { TAG_tags, model.Comment},
                };
        foreach( var field in model.Fields )
        { 
            if( cleanNameGoStruct( field.Type.Name) == Class_Generic_MODEL)
                continue;

            dict[field.Name] = new Dictionary<string, string>
                {
                    { TAG_Type, field.Type.Name },
                    { TAG_tags, field.Description},
                };
        }
        JsonSerializerOptions options = new JsonSerializerOptions();
        options.SetupExtensions();
        var str = JsonSerializer.Serialize(dict, options);

        var txt = $@"

        {model.Name}={str}
        {DELIMETER}
        ";
        jsonOut += txt;
        
    }

    return jsonOut;

    //return Export_SvcModels( "", "", meths);
}

private static void generateModelsFieldsAndRead(            
    ConvASqlModels scanner,
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
            //|| (fld.Name == CLASS_OutGeneralData && fld.IsPromoted) 
            //|| (fld.Name == CLASS_InGeneralData && fld.IsPromoted)
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
                    
            
            {
                //check if is derivated from CLASS_OutGeneralData
                
                
                if( baseClass.TypeClass != TypeClass.Interface )
                {
                    var genType = baseClass.getNormalizedType();
                    StructInfo structBase = null;
                    if( scanner.searchStructInImports( genType.Name, genType.PackageInfo, out structBase))
                    {
                        TypeInfo unusedBase = null;
                        generateModelsFieldsAndRead( scanner, NameSpacePrefix, structBase.Fields,
                                ref fields, ref readData, ref unusedBase, modelTypeName, dictCollectionFields);
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
        if( fld.IsPromoted && fld.Name == Class_Generic_MODEL) 
        {
            
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


private bool isInheritClass(TypeInfo typeclass, string nameClassBase)
{
    var structBase = ConvASqlModels.getStructInfoByName( typeclass.Name );
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
                var structBase = ConvASqlModels.getStructInfoByName( fld.Type.Name );
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




private static TypeInfo getTypeDef(TypeInfo type1 )
{
    if( ConvASqlModels.s_AllTypeDefs.ContainsKey(type1.Name) )
    {
        var type = ConvASqlModels.s_AllTypeDefs[type1.Name];
        var t1 = getTypeDef( type );
        return t1;
    }
    return type1;
}

        

public static string ConvertTypeToTs(string prefixSvc, goscanner.Metadata.TypeInfo typeGen)
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

public static string Export_SvcModels( Options options, 
    string models, 
    Dictionary<string, string> dictCollectionFields)
{

    return checkText( "" );
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
    if( ConvASqlModels.s_currentScanner.SourceFileName.Contains("WheelGame.go"))
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

