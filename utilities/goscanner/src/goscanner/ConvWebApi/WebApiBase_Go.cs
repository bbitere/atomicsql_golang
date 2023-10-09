using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Reflection.Metadata;
using System.Text;
using System.Threading.Tasks;
using goscanner.Metadata;
using static System.Net.Mime.MediaTypeNames;

namespace goscanner.ConvWebApi
{
    
internal class WebApiBase_Go
{
    
    public static string BASE_NAME_FLD = "@base";
    public static string BASE_NAME_FLD1 = "Base";
    public static  string FLD_ID = "ID";

    //public const string Arg0 = "arg0";

    public static string CLS_ajax_service = "";
        

    public static string CLASS_InGeneralData = "InGeneralData";
    public static string CLASS_OutGeneralData = "OutGeneralData";
    public static string Class_Generic_MODEL = "Generic_MODEL";

    public static bool   InlineAllFields_ClassInGeneralData = true;
        

    public static string UseTemplate(string original, ConvCommon.TemplateItem template, Dictionary<string, string>dict)
    { 
        if( template == null) {
            return original;
        }
        return template.ConvertTemplate( dict);
    }

    public static string exportSvc_AllMethods( ConvWebApi scanner, Options options, List<FunctionInfo> methods, ref string textPostGetRoute  )
    {
            var meths = "";
            foreach( var m in methods)
            {
                goscanner.Metadata.ParameterInfo[] paramList = m.Signature.Signature.Parameters;
                var err = false;
                if( paramList.Length != 1)
                {
                    err = true;
                    if( paramList.Length > 1)
                    { 
                        HashSet<string> inheritedTypeNames = new(StringComparer.Ordinal);
                        if( scanner.isInheritedBy( paramList[0].Type, CLASS_InGeneralData, inheritedTypeNames ) )
                        {
                            Console.WriteLine( $"Method '{m.Signature.Name}' member of SVC should have 1 input paramenter. (That is a sturcture containing all input data)");
                        }
                    }
                    continue;
                }

                if (paramList[0].Type.Name == "" || paramList[0].Type.Name == null)
                {
                    err = true;
                    Console.WriteLine($"Method '{m.Signature.Name}' member of SVC should have 1 input paramenter. This should be a struct type having a name");
                }
                
                if( !err)
                {
                    var methTxt = "";
                    var postText = "";
                    ExportGoFile_Method( options, CLS_ajax_service, m.Signature.Name, cleanNameGoStruct(paramList[0].Type.Name), ref methTxt, ref postText );
                    meths += methTxt;
                    textPostGetRoute += postText;
                }
            }
            return meths;

            //return ExportGoFile_Service( "Svc", "Svc", dir_controllers, meths);
    }
    public static string generateModelsFile(string ExportWebApi_TsSvcFile)
    {
        var fi = new FileInfo(ExportWebApi_TsSvcFile);
        var name = fi.Name.Replace(fi.Extension,"");
        var parts = name.Split('.');

        var separator = Path.DirectorySeparatorChar;

        if (parts.Length == 3)
            return $"{fi.DirectoryName}{separator}{parts[0]}Models.{parts[1]}.{parts[2]}{fi.Extension}";

        if ( parts.Length == 2 )
            return $"{fi.DirectoryName}{separator}{parts[0]}Models.{parts[1]}{fi.Extension}";

        if (parts.Length == 1)
            return $"{fi.DirectoryName}{separator}{parts[0]}Models{fi.Extension}";

        return $"{fi.DirectoryName}{separator}{fi.Name}Models{fi.Extension}";
    }
        


/// <summary>
/// 
/// </summary>
/// <param name="ajax_service"></param>
/// <param name="nameClass"></param>
/// <param name="dir_controllers">
/// dir_controllers = sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/controllers
/// </param>
/// <returns></returns>
public static string ExportGoFile_Service(
            Options options,
            string ajax_service, 
            string nameClass, 
            string dir_controllers,
            string text_method_controllers,
            string text_method_post )
{
    var txt = $@"
    package {ajax_service}

import (
	""net/http""
	echo ""github.com/labstack/echo/v4""
	controllers ""{dir_controllers}""
)
/*
func (_this *{nameClass}) {nameClass}_processAjax(c echo.Context) error {{

	postData := new(controllers.{CLASS_InGeneralData})
	err := c.Bind(postData)
	if err != nil {{
		return err
	}}

	switch postData.Msg {{

	//--------------------------------------------------------------------
		{{text_cases}}
		//--------------------------------------------------------------------
		default: {{
			return c.JSON(http.StatusInternalServerError, ""not_defined_msg"")
		}}
		//--------------------------------------------------------------------
	}}

	//return c.JSON(http.StatusInternalServerError, err.Error())
}}
*/
{text_method_controllers}

func Generate_SetupRouter(svcHandler *{nameClass}, e *echo.Echo) {{

{text_method_post}
}}     


    ";

    txt = UseTemplate( txt, options.ConvertWebApi.Templ_GoLang,
                
        new Dictionary<string,string>()
        { 
            { "dir_controllers", dir_controllers},
            { "text_method_controllers", text_method_controllers},
            { "text_method_post", text_method_post},
            { "nameClass", nameClass},
            
        }
        );

    return txt;


}

    

/// <summary>
/// Mines_getGame
/// </summary>
public static string ExportGoFile_Method(Options options, string nameClass, string methodName, string methodTypeInput, ref string text, ref string postText)
{
    postText = $@"
        e.POST(""/{methodName}"", svcHandler.{methodName})
        e.GET(""/{methodName}"", svcHandler.{methodName})
    ";
    /*
    text = $@"
        //----------------------------------------------------------------------------------
        case ""{methodName}"":
		{{
            methData := new(controllers.{methodTypeInput})
			err := c.Bind(methData)
			if err != nil {{
				return err
			}}

			data, err := _this.CtxControllers.{methodName}(methData)
			if err != nil {{
				return c.JSON(http.StatusInternalServerError, err.Error())
			}}
			return c.JSON(http.StatusCreated, data)
		}}
    ";*/

    text = $@"
        //----------------------------------------------------------------------------------
func (_this *AjaxSvc) {methodName}(c echo.Context) error {{

	postData := new(controllers.{methodTypeInput})
	err := c.Bind(postData)
	if err != nil {{
		return err
	}}
    data, err := _this.CtxControllers.{methodName}(postData)
		if err != nil {{
			return c.JSON(http.StatusInternalServerError, err.Error())
		}}
		return c.JSON(http.StatusCreated, data)
}}
    
            ";


    text = UseTemplate( text, options.ConvertWebApi.Templ_GoLangMethod,
                
        new Dictionary<string,string>()
        { 
            { "methodName", methodName},
            { "methodTypeInput", methodTypeInput},
        }
        );

    return text;
}



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

    }
}
