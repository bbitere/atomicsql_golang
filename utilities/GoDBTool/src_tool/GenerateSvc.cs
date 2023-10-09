using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Reflection;
using System.IO;
using System.Diagnostics;

namespace src_tool
{
    class GenerateSvc
    {

        const string el = "el";//the name of element inf function _parse(el:Element)
        Dictionary< string, Type> collectClasses= new Dictionary<string, Type>();
        Dictionary< string, Type> collectEnum   = new Dictionary<string, Type>();

        public GenerateSvc()
        {

        }

        public static string assemblydirectory = "";
        public static Assembly CurrentDomaine_assemblyResolve( object sender, ResolveEventArgs args)
        {
            var tokens = args.Name.Split(",".ToCharArray() );
            var ret = Assembly.LoadFile( Path.Combine(new string[]{ assemblydirectory, tokens[0]+".dll" } ) );
            if( ret != null)
                return ret;

            ret = Assembly.LoadFile( Path.Combine(new string[]{ assemblydirectory, tokens[0]+".lib" } ) );
            return ret;
        }

        public void _GenerateSvc( String dllPath, String svcFile, String svcModelsFile )
        {
            var f = new FileInfo(dllPath);
            
            Type[] types  = null;
            var dir = f.Directory;
                
            //var dir1 = Directory.GetCurrentDirectory();
            //Directory.SetCurrentDirectory( dir.FullName );
            try
            {
                AppDomain.CurrentDomain.AssemblyResolve += CurrentDomaine_assemblyResolve;
                assemblydirectory = Path.GetDirectoryName( f.FullName );

                var assembly = Assembly.LoadFile(f.FullName);
                types = assembly.GetExportedTypes();

            }catch( Exception e )
            {
                Console.WriteLine( e.Message );
                return;
            }
            //Directory.SetCurrentDirectory( dir1 );
            

            var methodsTxt = "";

            foreach( var type in types)
            {
                var attrs = type.GetCustomAttributes(true);
                foreach( var attr in attrs )
                {
                    if( attr.GetType().FullName == "System.ServiceModel.ServiceContractAttribute" )
                    {
                        var methods = type.GetMethods();
                        foreach( var m in methods )
                        {
                            var b = 0;
                            var attrsM = m.GetCustomAttributes(true);
                            foreach( var attr1 in attrsM )
                            {
                                if( attr1.GetType().FullName == "System.ServiceModel.OperationContractAttribute")
                                {
                                    b |= 1;                                    
                                }else                                
                                if( attr1.GetType().Name == "TsAttribute")
                                {
                                    b |= 2;                                    
                                }
                            }
                            if( b == 3 )
                            {
                                methodsTxt += printMethod(m);
                            }
                        }
                    }
                }
            }

            var classesCode = "";

            var onlyCls = collectClasses.Where(x=>x.Value.IsValueType).ToList();

            //check all sub classes
            var list = collectClasses.ToList();
            for( var i= 0; i < list.Count; i++)
            {
                printClass( list[i].Value);
                list = collectClasses.ToList();
            }

            var list1 = collectClasses.ToList();
            foreach( var it in list1)
            {
                if( it.Value.IsValueType )
                {
                    //structurile sunt folosite doar pt return si deci nu intra in export
                }else
                {
                    classesCode += printClass( it.Value );
                }
            }


            var SvcTxt = @"
    /** this file is generated automatically by TsTool.exe */
    type SVC_FnError = ( string ) => void;

    class Svc
    {
        constructor()
        {
        }
        "+methodsTxt+@"
    }
    var SVC = new Svc();
            ";


            var SvcModelsTxt = $@"
        /** this file is generated automatically by TsTool.exe */

        type clsGenericItem = dtoGenericItem;//replace with something

        abstract class SvcObj_MODEL
        {{
            protected __data:any;
            public constructor( data: any=null)
            {{
                this.__data = data === null ? {{}}: data;
            }}
            public setData(d:any):void {{this.__data=d;}};
            public readData(name:string, el:Element){{}};
        }}
        abstract class Generic_MODEL extends SvcObj_MODEL
        {{
            public abstract get ID(): string;

            public getFldData(fldName:string):string{{ return S(this.__data[fldName]); }};
            public setFldData( fldName: string, val:string ) {{ return this.__data[fldName] = val}};

            public constructor( data: any=null)
            {{
                super(data);
            }}
        }}
        class AJX_classifList  extends SvcObj_MODEL
        {{
            public Contents(): HashMap< Array< clsGenericItem> > {{ return null; }};
            public Titles():   HashMap<string> {{ return null; }};

            public constructor( data: any=null)
            {{
                super(data);
            }}
        }}

        {classesCode}
        ";

            File.WriteAllText(svcFile, SvcTxt);
            File.WriteAllText(svcModelsFile, SvcModelsTxt);
        }


        string getClassAnonyomus(Type t, string dataName, 
            out string argsSuccess, 
            out string instance_code )
        {
            instance_code = "{}";
            argsSuccess = "";
            var bFirst = false;
            var s = "";
            var flds = t.GetFields().Where( x=>x.GetCustomAttributes(false).Where(y=>y.GetType().Name == "DataMemberAttribute").Count() > 0).ToList();
            foreach( var fld in flds)
            {
                if( !bFirst)
                    bFirst = true;
                else
                {
                    s+= ", ";
                    argsSuccess += ", ";
                }

                var typeN = getTsTypeName( fld.FieldType );
                s += $"{fld.Name}: {typeN}";

                
                if( typeN == "number")
                {
                    argsSuccess += $"{dataName}['{fld.Name}']";
                }else
                if( collectClasses.ContainsKey( fld.FieldType.Name ))
                {
                    var cls = CheckClassName( fld.FieldType.Name );
                    argsSuccess += $"new {cls}( {dataName}['{fld.Name}'] )";
                }else
                {
                    argsSuccess += $"{dataName}['{fld.Name}']";
                }
            }
            
            return s;
        }

        string CheckClassName(string name)
        {
            return name;
        }

        string getTsTypeName(Type t)
        {
            if( t.IsArray)
            {
                //Debugger.Break();
                var typeArg = getTsTypeName(t.GetElementType() );                
                return $"Array<{typeArg}>";
            }else
            if( t.Name == "List`1" )
            {
                var typeArg = getTsTypeName(t.GetGenericArguments()[0]);                
                return $"Array<{typeArg}>";
            }else
            if( t.Name == "Nullable`1" )
            {
                var typeArg = getTsTypeName(t.GetGenericArguments()[0]);                
                return $"{typeArg}";
            }else
            if( t.Name == "Dictionary`2" )
            {
                var typeArg = getTsTypeName(t.GetGenericArguments()[1]);
                return $"HashMap<{typeArg}>";
            }else
            if( t.Name == typeof(int).Name 
             || t.Name == typeof(byte).Name
             || t.Name == typeof(short).Name 
             || t.Name == typeof(decimal).Name )
            {
                return "number";
            }else
            if( t.Name == typeof(void).Name)
            {
                return "";
            }else
            if( t.Name == typeof(bool).Name)
            {
                return "boolean";
            }else
            if( t.Name == typeof(DateTime).Name)
            {
                return "Date";
            }else
            if( t.Name == typeof(string).Name )
            {
                return "string";
            }else
            if( t.Name == typeof(float).Name 
             || t.Name == typeof(double).Name )
            {
                return "FLOAT";
            }else
            if( t.IsEnum)
            {
                collectEnum[ t.Name ] = t;
                return t.Name;
            }else
            {
                if( !t.IsValueType)
                    collectClasses[t.Name] = t;

                return t.Name;
            }
        }


        string getParseMethBy_TypeName( string propName, Type t )
        {
            if( t.IsArray)
            {
                //Debugger.Break();
                var codeArg = getParseMethBy_TypeName(propName, t.GetElementType() );                
                var typeArg = getTsTypeName(t.GetElementType());

                return $"_parseArray<{typeArg} >( '{propName}', {el}, function({el}:Element){{ return {codeArg}; }} )";
            }else
            if( t.Name == "List`1" )
            {
                var codeArg = getParseMethBy_TypeName(propName, t.GetGenericArguments()[0]);                
                var typeArg = getTsTypeName(t.GetGenericArguments()[0]);

                return $"_parseArray<{typeArg} >( '{propName}', {el}, function({el}:Element){{ return {codeArg}; }} )";
            }else
            if( t.Name == "Nullable`1" )
            {
                var typeArg = getParseMethBy_TypeName(propName, t.GetGenericArguments()[0]);                
                return $"{typeArg}";
            }else
            if( t.Name == "Dictionary`2" )
            {
                if( t.GetGenericArguments()[1].Name == "String")
                { 
                    var typeArg = getParseMethBy_TypeName(propName, t.GetGenericArguments()[1]);
                    return $"_parseDictS( '{propName}', {el} )";
                }else
                {
                    var codeArg = getParseMethBy_TypeName(propName, t.GetGenericArguments()[1]);                
                    var typeArg = getTsTypeName(t.GetGenericArguments()[1]);

                    return $"_parseDict<{typeArg} >( '{propName}', {el}, function({el}:Element){{ return {codeArg}; }} )";                    
                }
                //return $"HashMap<{typeArg}>";
            }else
            if( t.Name == typeof(int).Name 
             || t.Name == typeof(byte).Name
             || t.Name == typeof(short).Name 
             || t.Name == typeof(decimal).Name )
            {
                return $"_parseNumber( '{propName}', {el} )";
            }else
            if( t.Name == typeof(void).Name)
            {
                return "";
            }else                    
            if( t.Name == typeof(bool).Name)
            {
                return $"_parseBool( '{propName}', {el} )";
            }else
            if( t.Name == typeof(DateTime).Name)
            {
                return $"_parseDate( '{propName}', {el} )";
            }else
            if( t.Name == typeof(string).Name )
            {
                return $"_parseString( '{propName}', {el} )";
            }else
            if( t.Name == typeof(float).Name 
             || t.Name == typeof(double).Name )
            {
                return $"_parseFloat( {propName}, {el} )";
            }else
            if( t.IsEnum)
            {
                collectEnum[ t.Name ] = t;
                return "/*t.Name*/";
            }else
            {
                if( !t.IsValueType)
                    collectClasses[t.Name] = t;
                
                return $"( function( {el}: Element ) {{ var i = new {t.Name}(); i.readData('{propName}', {el} ); return i; }})({el})";
                //return t.Name;
            }
        }
    
        void Nop(){ }
        string printMethod( MethodInfo m)
        {
            var paramsTs = "";
            var methParams = m.GetParameters();

            if( m.Name == "Reqs_getClassification")
            { 
                Nop();
            }

            var argsAjax = "";
            var bFirst = false;
            foreach( var arg in methParams )
            {
                var t = arg.ParameterType;
                /*
                if( !collectClasses.ContainsKey( t.FullName ) )
                {
                    if( !t.IsValueType)
                        collectClasses[ t.FullName ] = t;
                }*/

                if( !bFirst)
                {
                    bFirst = true;
                }else
                {
                    paramsTs += ", ";
                    argsAjax += @",
                ";
                }

                var nameType = getTsTypeName(t);
                paramsTs += $"{arg.Name}:{nameType}"; 

                if( t.Name == "List`1" )
                {
                    argsAjax += $"'{arg.Name}':SoapArgArr({arg.Name})"; 
                }else
                if( t.Name == "Nullable`1" )
                {
                    argsAjax += $"'{arg.Name}':SoapArgNull({arg.Name})"; 
                }else
                if( t.Name == "Dictionary`2" )
                {
                    argsAjax += $"'{arg.Name}':SoapArgDict({arg.Name})"; 
                }else
                {
                    argsAjax += $"'{arg.Name}':{arg.Name}"; 
                }
            }
            
            var retSucces  = "ret";
            var argSucces  = "";
            var instance_code = "";
            var TypefnSuccess = getTsTypeName(m.ReturnType);
            if( m.ReturnType.GetCustomAttributes(false).Where(x=>x.GetType().Name == "TsAttribute").Count() > 0 )
            {
                if( m.ReturnType.IsValueType )//doar pt strcturi se extinde
                {
                    TypefnSuccess = getClassAnonyomus( m.ReturnType, retSucces, out argSucces, out instance_code );
                }else
                if( collectClasses.ContainsKey( m.ReturnType.Name ))
                {
                    var cls = CheckClassName( m.ReturnType.Name );
                    //argSucces = $"new {cls}( {retSucces} )";
                    argSucces = $"{retSucces} as {cls}";
                    instance_code = $"new {cls}()";
                }
            }else
            {
                if( collectClasses.ContainsKey( m.ReturnType.Name ))
                {
                    var cls = CheckClassName( m.ReturnType.Name );
                    //argSucces = $"new {cls}( {retSucces} )";
                    argSucces = $"{retSucces} as {cls}";
                    instance_code = $"new {cls}()";
                }
            }
            if( instance_code == "")
                instance_code = "null";

        var s = $@"
        {m.Name}( {paramsTs}, 
            fnSuccess:({TypefnSuccess})=>void, fnError:SVC_FnError )    
        {{
            AjaxSoap_Send( {{
                methodName:'{m.Name}', 
                instance:{instance_code}, 
                args:{{ {argsAjax} }}, 
                fnSuccess: function({retSucces}:any){{ fnSuccess( {argSucces}); }}, 
                fnError: fnError
                }}
            )
        }}
        ";

            return s;
        }


        string scanFields( Type cls, bool bGenerateSet, bool bEntityObject, List<string> excludeFields, ref string readData_Code )
        {
            var s = "";
            if( bGenerateSet )
            {
                var flds = cls.GetFields().Where( x=>x.GetCustomAttributes(false).Where(y=>y.GetType().Name == "DataMemberAttribute").Count() > 0).ToList();
                var props = cls.GetProperties().Where( x=>x.GetCustomAttributes(false).Where(y=>y.GetType().Name == "DataMemberAttribute").Count() > 0).ToList();

                if( bEntityObject ||( flds.Count == 0 && props.Count > 0 ))
                {
                    foreach( var fld in props)
                    {
                        var propName = fld.Name;
                        if( propName == "EntityKey")
                            continue;
                        //if( propName == "ID")
                        //    continue;
                        if( excludeFields != null && excludeFields.IndexOf(propName) >= 0 )
                            continue;

                        var propType = getTsTypeName( fld.PropertyType );
                        var parseCodeByType = getParseMethBy_TypeName( propName, fld.PropertyType );

            
        if( propName == "ID")
        {
            readData_Code += $@"
            this.{"ID"} = _parseString( '{"ID"}', {el} );";
        }else
        {
            readData_Code += $@"
            this.{propName} = {parseCodeByType};";

            s += $@"
            get {propName}(): {propType}  {{ return this.__data['{propName}'] as {propType};}}
            set {propName}(v: {propType}) {{        this.__data['{propName}'] = v;}} ";
                        }
                    }   
                }else
                {
                    foreach( var fld in flds)
                    {
                        var propName = fld.Name;
                        if( propName == "EntityKey")
                            continue;
                        if( propName == "ID")
                            continue;
                        if( excludeFields != null && excludeFields.IndexOf(propName) >= 0 )
                            continue;

                        var propType = getTsTypeName( fld.FieldType );
                        var parseCodeByType = getParseMethBy_TypeName( propName, fld.FieldType );

        if( propName == "ID")
        {
            readData_Code += $@"
            this.{"ID"} = _parseString( '{"ID"}', {el} );";
        }else
        {
            readData_Code += $@"
            this.{propName} = {parseCodeByType};";

            s += $@"
            get {propName}(): {propType}  {{ return this.__data['{propName}'] as {propType};}}
            set {propName}(v: {propType}) {{        this.__data['{propName}'] = v;}} ";
                        }
                    }
                }

            }else
            {
                var flds = cls.GetFields().Where( x=>x.GetCustomAttributes(false).Where(y=>y.GetType().Name == "DataMemberAttribute").Count() > 0).ToList();
                
                foreach( var fld in flds)
                {
                    var propName = fld.Name;
                    if( excludeFields != null && excludeFields.IndexOf(propName) >= 0 )
                        continue;

                    var propType = getTsTypeName( fld.FieldType );
                    var parseCodeByType = getParseMethBy_TypeName( propName, fld.FieldType );

            readData_Code += $@"
            this.{propName} = {parseCodeByType};";

            s += $@"
            get {propName}(): {propType}  {{ return this.__data['{propName}'] as {propType};}}
            set {propName}(v: {propType}) {{        this.__data['{propName}'] = v;}} ";
                }
            }
            return s;
        }
        string printClass( Type cls )
        {
            var s = "";

            var readData_Code = "";
            if( cls.BaseType.Name == "AJX_classifList")
            {
                var excl = new String[]{ "Contents", "Titles" };
                var fields = scanFields(cls, false, false, excl.ToList(), ref readData_Code );

        s = $@"
        class {cls.Name} extends AJX_classifList
        {{
            public Contents(): HashMap<Array< clsGenericItem> > {{ return this.content;}};
            public Titles():   HashMap<string> {{ return this.__data['titles']   as HashMap<string>;}};

            {fields}

            readData(name:string, {el}:Element):void
            {{
                {el} = _parseFindRoot(name, {el});
                {readData_Code}                
            }}
        }}
                ";

            }else
            if( cls.BaseType.Name == "EntityObject"
             || cls.BaseType.Name == "dtoBase"
             || cls.BaseType.Name == "Object" )
            {
                var fields = scanFields(cls, true, cls.BaseType.Name == "EntityObject", null, ref readData_Code );

        s = $@"
        class {cls.Name} extends Generic_MODEL
        {{
            public constructor( data: any = null ){{super(data);}}

            public  get ID(): string {{return ''+(this.__data['ID']  as string); }};
            public  set ID(s: string) {{this.__data['ID'] = s; }};

            {fields}

            readData( name:string, {el}:Element):void
            {{
                {el} = _parseFindRoot(name, {el});
                {readData_Code}                
            }}
        }}
    

                ";
            }else
            {
                Debugger.Break();
            }
            return s;
        }
    }
}
