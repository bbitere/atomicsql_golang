//******************************************************************************************************
//  GoTypeInfo.cs - Gbtc
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
//  06/21/2018 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Runtime.CompilerServices;
using System.Text.Json.Serialization;
using System.Xml.Linq;
using Dahomey.Json.Attributes;

namespace goscanner.Metadata
{
    [Serializable]
    public enum TypeClass
    {
        Simple,
        Array,
        Slice,
        Map,        // Dictionary<kT, vT>
        Channel,
        Struct,
        Function,   // Func<..., T> / Action<...>
        Interface
    }

    public enum DerivedTypeInfo
    {
        Pointer,
        Array,
        Map
    }
    public struct TPackageInfo
    { 
        public string Name;
        public FileMetadata Metafile;
        public TPackageInfo( string name, FileMetadata metafile)
        {
            this.Name = name;
            this.Metafile = metafile;
        }
        public TPackageInfo( string name)
        {
            this.Name = name;
            this.Metafile = null; //the declaration is done in the current file
        }
    }

    [Serializable]
    public class TypeInfo
    {
        private static int s_UID = 1;
        public int UID = s_UID++;

        private string _Name;
        public string Name{ 
                get{return _Name; }
                set{_Name = value; 
                //if( value == "m.User")
               //     Debugger.Break();
                } 
            }
        public string PackageName = "";
        public string TypeName;
        public string FullTypeName;
        public TypeClass TypeClass;
        public bool IsDerefPointer;
        public bool IsByRefPointer;
        public bool IsConst;

        [JsonIgnore]
        private FileMetadata fileDeclaredLocation;// i need this because the package name can be corrently identiied only when metafile is correct

        [JsonIgnore]
        public TPackageInfo PackageInfo 
        { 
            get{ return new TPackageInfo(PackageName, fileDeclaredLocation); } 
            set
            { 
                if( !(_Name == "var"    || _Name == "dynamic" 
                   || _Name == "object" || _Name == "string" 
                   || _Name == "void"    || _Name == "bool" ))
                {
                    PackageName             = value.Name;
                    fileDeclaredLocation    = value.Metafile;
                }else
                {
                    var ii = 0;
                }
            } 
        }
        public void setFileMetadata(FileMetadata f){ this.fileDeclaredLocation = f;}
        public void PackageInfo_setName(string name)
        { 
            if( !(_Name == "var"    || _Name == "dynamic" 
                || _Name == "object" || _Name == "string" 
                || _Name == "void"    || _Name == "bool" ))
            {
                this.PackageName = name;
                //fileDeclaredLocation    = value.Metafile;
            }
            
        }

        public static FileMetadata s_CurrentMetaData;
        public TypeInfo()
        {
            if( UID == 7090 )
                UID = UID;
            this.fileDeclaredLocation = s_CurrentMetaData;
        }
        public bool IsPointer()
        {
            if( Name.StartsWith("*") )
                return true;

            return false;
        }

        public bool IsBool()
        {
            return this.Name == "bool"
                || this.Name == "@bool";
        }
        public bool IsString ()
        {
            return this.Name == "string"
                || this.Name == "@string";
        }
        public static bool isPrimaryType(string idenif)
        {
            if( idenif == "int32"
            ||  idenif == "int64"
            ||  idenif == "int16"
            ||  idenif == "byte"
            ||  idenif == "int"
            ||  idenif == "int8"
            ||  idenif == "uint32"
            ||  idenif == "uint64"
            ||  idenif == "uint16"
            ||  idenif == "byte"
            ||  idenif == "uint8"
            ||  idenif == "float32"
            ||  idenif == "float64"
            ||  idenif == "bool"
            ||  idenif == "string")
                return true;

            return false;

        }

        public TypeInfo getNormalizedType( TGenTypeDetails details=null)
        {
            var type = this;
            if (type.Name == "var")
                return null;

            if (type.Name.StartsWith( "map") )
                type = type;


            if( type.TypeName != null )
            {
                var s = type.TypeName;
                var ret = getPrimaryTypeName1(s, type.PackageName, this.fileDeclaredLocation);
                ret.fileDeclaredLocation = this.fileDeclaredLocation;
                return ret;
            }else
            {
                var s = type.Name;
                var ret =  getPrimaryTypeName1(s, type.PackageName, this.fileDeclaredLocation);
                ret.fileDeclaredLocation = this.fileDeclaredLocation;
                return ret;
            }
        }
        public static TypeInfo getPrimaryTypeName1(string s, string PackageName, FileMetadata fileDeclaredLocation)
        {
            s = s .Replace("[", " [ ");
            s = s .Replace("]", " ] ");
            s = s .Replace("<", " [ ");
            s = s .Replace(">", " ] ");
            s = s .Replace(",", " , ");
            s = s .Replace("*", "");
            s = s .Replace("@", "");

            var parts = s.Split(" ", StringSplitOptions.RemoveEmptyEntries);

            var offset = 0;
            var ret = _parsePrimaryTypeName( parts, ref offset, PackageName );
            var ret1 = ret as TGenType;

            if( ret1 != null )
            {
                if( ret1.Name == "ptr" && ret1.genericArgs.Length == 1)
                {
                    return ret1.genericArgs[0];
                }else
                if( ret.Name == "slice" && ret1.genericArgs.Length == 1)
                {
                    ret = new ArrayTypeInfo
                            {
                                Name = ArrayTypeInfo.ArrName,
                                FullTypeName = ArrayTypeInfo.ArrName,
                                TargetTypeInfo = ret1.genericArgs[0],
                                Length = null,
                            };
                }else
                if( ret.Name == "map" && ret1.genericArgs.Length == 1)
                {
                    ret = new MapTypeInfo
                            {
                                Name = "map",
                                FullTypeName = "map",
                                KeyTypeInfo = ret1.genericArgs[0],
                                ElementTypeInfo= ret1.genericArgs[0]
                            };
                }else
                if( ret.Name == "map" && ret1.genericArgs.Length == 2)
                {
                    ret = new MapTypeInfo
                            {
                                Name = "map",
                                FullTypeName = "map",
                                KeyTypeInfo = ret1.genericArgs[0],
                                ElementTypeInfo= ret1.genericArgs[1]
                            };
                }
            }
            
            return ret;
        }

        protected static TypeInfo _parsePrimaryTypeName(string[] parts, ref int offset, string packageName)
        {
            var typeName = parts[offset];
            offset++;

            if( offset  < parts.Length 
             && parts[ offset] == "[" )
            {
                offset++;
                var listArgs = new List<TypeInfo>();    
                for( ; offset < parts.Length; offset++)
                {
                    var arg = _parsePrimaryTypeName( parts, ref offset, packageName );
                    listArgs.Add(arg);

                    //offset++;
                    if( parts[ offset ] == "]")
                    {
                        offset++;
                        break;
                    }
                    else
                    if( parts[ offset ] == ",")
                        continue;
                    else
                    {
                        Debugger.Break();//error
                        break;
                    }
                }
                var namePArts = typeName.Split( '.' );
                var myPackaqe  = "";
                var myname     = "";
                if( namePArts.Length > 1)
                {
                    myPackaqe  = namePArts[0];
                    myname     = namePArts[namePArts.Length-1];
                }else
                {
                    myPackaqe  = "";
                    myname     = namePArts[0];
                }

                if( myname.StartsWith( "map" ) )
                {
                    if( offset < parts.Length && parts[offset] != "]") 
                    {
                        var arg1 = _parsePrimaryTypeName( parts, ref offset, packageName );
                        listArgs.Add(arg1);
                    }
                }
                 
                {
                    return setPackage( new TGenType{ 
                        Name = myname, 
                        FullTypeName = typeName,
                        genericArgs = setPackages( listArgs.ToArray(), myPackaqe != ""? myPackaqe: packageName, null ),
                        
                    },  myPackaqe != ""? myPackaqe: packageName, null );
                }
            }

            var nameParts = typeName.Split( '.' );
            if( nameParts.Length > 1 )
            {
                return new TypeInfo{ 
                    Name = nameParts[nameParts.Length-1], 
                    FullTypeName =typeName, 
                    TypeClass = TypeClass.Simple,
                    PackageName = nameParts[0]
                    };
            }else
            {
                return setPackage( new TypeInfo{ 
                                Name = typeName,
                                FullTypeName =typeName, 
                                TypeClass = TypeClass.Simple 
                                }, packageName, null );
            }
        }

        static TypeInfo test = getPrimaryTypeName1( "DD[ string, DD[string, int] ]","", null);
        static TypeInfo testMap = getPrimaryTypeName1( "map[ string] DD[string, int]","", null);
        static TypeInfo testMap2 = getPrimaryTypeName1( "map<@string, map<@string, @string>>","", null);
        

        private static (string, string) getPacket_Type(string name)
        { 
            //I renmove the pointer name
            var name1 = name.Replace("*","");
            var parts = name1.Split( "." );
            if( parts.Length == 2)
                return (parts[0], parts[1]);
            return ("", parts[0]);
        }

        public static TypeInfo setPackage( TypeInfo type, string packageName, FileMetadata fileMetadata)
        {
            if( type.PackageName == "" )
                type.PackageInfo = new TPackageInfo( packageName, fileMetadata);
            return type;
        }

        

        public static ParameterInfo[] setPackages( ParameterInfo[] types, string packageName,  FileMetadata fileMetadata)
        {
            foreach( var param1 in types ) 
            {
                if( param1.Type.PackageName == "" )
                {
                    var (package, nameType) = getPacket_Type( param1.Type.Name);
                    if( package != "" && fileMetadata != null )
                    {
                        param1.Type.Name = nameType;
                        param1.Type.PackageInfo = new TPackageInfo( package, fileMetadata );
                    }else
                    {
                        param1.Type.PackageInfo = new TPackageInfo( packageName );
                    }
                }
            }
            return types;
        }
        public static TypeInfo[] setPackages( TypeInfo[] types, string packageName,  FileMetadata fileMetadata)
        {
            foreach( var param1 in types ) 
            {
                if( param1.PackageName == "" )
                {
                    var (package, nameType) = getPacket_Type( param1.Name);
                    if( package != "" && fileMetadata != null )
                    {
                        param1.Name = nameType;
                        param1.PackageInfo = new TPackageInfo( package, fileMetadata );
                    }else
                    {
                        param1.PackageInfo = new TPackageInfo( packageName );
                    }
                }
            }

            return types;
        }
    


        public virtual TypeInfo Clone() =>
            new()
            {
                Name = Name,
                TypeName = TypeName,
                FullTypeName = FullTypeName,
                TypeClass = TypeClass,
                IsDerefPointer = IsDerefPointer,
                IsByRefPointer = IsByRefPointer,
                IsConst = IsConst
            };

        public static readonly TypeInfo ObjectType =
            new()
            {
                Name = "object",
                TypeName = "object",
                FullTypeName = "System.Object",
                TypeClass = TypeClass.Simple
            };

        public static readonly TypeInfo VoidType =
            new()
            {
                Name = "void",
                TypeName = "void",
                FullTypeName = "void",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo BoolType =
            new()
            {
                Name = "bool",
                TypeName = "bool",
                FullTypeName = "bool",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo Int32Type =
            new()
            {
                Name = "int32",
                TypeName = "int32",
                FullTypeName = "int32",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo Int64Type =
            new()
            {
                Name = "int64",
                TypeName = "int64",
                FullTypeName = "int64",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo Int16Type =
            new()
            {
                Name = "int16",
                TypeName = "int16",
                FullTypeName = "int16",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo ByteType =
            new()
            {
                Name = "byte",
                TypeName = "byte",
                FullTypeName = "byte",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo Float64Type =
            new()
            {
                Name = "float64",
                TypeName = "float64",
                FullTypeName = "float64",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo StringType =
            new()
            {
                Name = "string",
                TypeName = "string",
                FullTypeName = "string",
                TypeClass = TypeClass.Simple
            };
        public static readonly TypeInfo TimeType =
            new()
            {
                Name = "time.Time",
                TypeName = "time.Time",
                FullTypeName = "time.Time",
                TypeClass = TypeClass.Simple
            };
        

        public static readonly TypeInfo VarType =
            new()
            {
                Name = "var",
                TypeName = "var",
                FullTypeName = "var",
                TypeClass = TypeClass.Simple
            };

        public static readonly TypeInfo DynamicType =
            new()
            {
                Name = "dynamic",
                TypeName = "dynamic",
                FullTypeName = "System.Dynamic.DynamicObject",
                TypeClass = TypeClass.Simple
            };

        public static readonly TypeInfo EmptyInterfaceType =
            new()
            {
                Name = "object",
                TypeName = "object",
                FullTypeName = "System.Object",
                TypeClass = TypeClass.Simple
            };
    }

    [Serializable]
    [JsonDiscriminator(DerivedTypeInfo.Pointer)]
    public class PointerTypeInfo : TypeInfo
    {
        public TypeInfo TargetTypeInfo;

        public override TypeInfo Clone() =>
            new PointerTypeInfo
            {
                Name = Name,
                TypeName = $"ptr<{TypeName}>",
                FullTypeName = $"go.ptr<{FullTypeName}>",
                TypeClass = TypeClass,
                IsDerefPointer = IsDerefPointer,
                IsByRefPointer = IsByRefPointer,
                IsConst = IsConst,
                TargetTypeInfo = TargetTypeInfo
            };
    }

    [Serializable]
    [JsonDiscriminator(DerivedTypeInfo.Array)]
    public class ArrayTypeInfo : TypeInfo
    {
        public const string ArrName = "array";
        public TypeInfo TargetTypeInfo;
        public ExpressionInfo Length;

        public override TypeInfo Clone() =>
            new ArrayTypeInfo
            {
                Name = Name,
                TypeName = $"{ArrName}<{TypeName}>",
                FullTypeName = $"go.{ArrName}<{FullTypeName}>",
                TypeClass = TypeClass,
                IsDerefPointer = IsDerefPointer,
                IsByRefPointer = IsByRefPointer,
                IsConst = IsConst,
                TargetTypeInfo = TargetTypeInfo,
                Length = Length
            };

        public static ArrayTypeInfo new_ArrayTypeInfo(TypeInfo TargetTypeInfo, ExpressionInfo Length )
        {
            return new ArrayTypeInfo
            {
                Name = TargetTypeInfo.Name,
                TypeName = $"{ArrName}<{TargetTypeInfo.Name}>",
                FullTypeName = $"go.{ArrName}<{TargetTypeInfo.Name}>",
                TypeClass = TypeClass.Slice,
                IsDerefPointer = false,
                IsByRefPointer = false,
                IsConst = false,
                TargetTypeInfo = TargetTypeInfo,
                Length = Length
            };

        }
        

    }

    [Serializable]
    [JsonDiscriminator(DerivedTypeInfo.Map)]
    public class MapTypeInfo : TypeInfo
    {
        public TypeInfo KeyTypeInfo;
        public TypeInfo ElementTypeInfo;

        public override TypeInfo Clone() =>
            new MapTypeInfo
            {
                Name = Name,
                TypeName = TypeName,
                FullTypeName = FullTypeName,
                TypeClass = TypeClass,
                IsDerefPointer = IsDerefPointer,
                IsByRefPointer = IsByRefPointer,
                IsConst = IsConst,
                KeyTypeInfo = KeyTypeInfo,
                ElementTypeInfo = ElementTypeInfo
            };
    

    }


    public class TGenType : TypeInfo
    { 
        public TypeInfo[] genericArgs;
    };
    public class TGenTypeDetails 
    {
        public string fullPackage;
        public string last_Item_Package;
    }




    
    
}
