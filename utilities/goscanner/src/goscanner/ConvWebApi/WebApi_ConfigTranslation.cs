using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using Antlr4.Runtime;
using System.Reflection.Metadata;

namespace goscanner.ConvWebApi
{
    
    public class ConvWebApi_ConfigTranslation : goscanner.ConvCommon.ConfigFile
    { 
        // -o -i -h -w=.\export_svc_config.txt C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers 
        // C:\Work\SlotMonitor\SlotGames\game-originals-main\tools

        public const string Version = "1.0.1";


        public ConvCommon.TemplateItem Templ_TsExportService;
        public ConvCommon.TemplateItem Templ_TsExportModels;

        public ConvCommon.TemplateItem Templ_GoLang;
        public ConvCommon.TemplateItem Templ_GoLangMethod;
        public List<string> ExportEnums = new List<string>();
        
        
        public string  CLASS_InGeneralData;
        public string  CLASS_OutGeneralData;
        public string  CLASS_SvcObj_MODEL;
        public string  SVCM;
        public string  FLD_ID;


        public bool   InlineAllFields_ClassInGeneralData = true;
        public string TsSvcFile;
        public string TsSvcModelsFile;
        public string GoSvcFile;
        


        
        protected override void setupSLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "SvcFile":         this.TsSvcFile = content; break;
                case "SvcModelsFile":   this.TsSvcModelsFile = content; break;
                case "TypeLang":        this.TypeLang = content; break;
                case "OutSvcFile":      this.GoSvcFile = content; break;
                case "SourcePathDir":   this.SourcePathDir = content; break;
                case "SourcePathDir2":  this.SourcePathDir2 = content; break;
                case "RootProject":     this.RootProject = content; break;
                case "GitProject":      this.GitProject = content; break;
                    
                case "SVCM":            this.SVCM = content; break;
                case "FLD_ID":          this.FLD_ID = content; break;

                case "CLASS_Base_MODEL":      this.Class_Base_MODEL = content; break;
                case "CLASS_InGeneralData":   this.CLASS_InGeneralData = content; break;
                case "CLASS_OutGeneralData":  this.CLASS_OutGeneralData = content; break;
                case "CLASS_SvcObj_MODEL":    this.CLASS_SvcObj_MODEL = content; break;
                case "InlineAllFields_ClassInGeneralData":    this.InlineAllFields_ClassInGeneralData = content.Trim() == "true"; break;
                    
                    
                
                default: 
                    Console.WriteLine( $"Not identified token {token1} in single line property ");
                break;
            }
        }

        protected override void updateProperties()
        {
            if( ExportEnums == null )
                Console.WriteLine("Export_Enums - not set");

            if( Templ_GoLangMethod == null )
                Console.WriteLine("Templ_GoLangMethod - not set");

            if( Templ_GoLang == null )
                Console.WriteLine("Templ_GoLang - not set");

            if( Templ_TsExportService == null )
                Console.WriteLine("Templ_TsExportService - not set");

            if( Templ_TsExportModels == null )
                Console.WriteLine("Templ_TsExportModels - not set");


            if( this.Class_Base_MODEL != null )
                WebApiBase_Go.Class_Generic_MODEL = this.Class_Base_MODEL;

            //if( this.InlineAllFields_ClassInGeneralData )
                WebApiBase_Go.InlineAllFields_ClassInGeneralData = this.InlineAllFields_ClassInGeneralData;

            if( this.CLASS_InGeneralData != null )
                WebApiBase_Go.CLASS_InGeneralData = this.CLASS_InGeneralData;

            if( this.CLASS_OutGeneralData != null )
                WebApiBase_Go.CLASS_OutGeneralData = this.CLASS_OutGeneralData;

            if( this.FLD_ID != null )
                WebApiBase_Go.FLD_ID =  this.FLD_ID;

            if( this.SVCM != null )
                WebApiBase_Ts.SVCM =  this.SVCM;

            if( this.CLASS_SvcObj_MODEL != null )
                WebApiBase_Ts.SvcObj_MODEL = this.CLASS_SvcObj_MODEL;

        }
        protected override void setupMLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "Export_Enums":            this.ExportEnums            = processExportListItems( content ); break;
                case "Templ_GoLangMethod":      this.Templ_GoLangMethod     = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_GoLang":            this.Templ_GoLang           = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_TsExportService":   this.Templ_TsExportService  = new ConvCommon.TemplateItem( token1, content ); break;
                case "Templ_TsExportModels":    this.Templ_TsExportModels   = new ConvCommon.TemplateItem( token1, content ); break;
                default: 
                    Console.WriteLine( $"Not identified token {token1} in multiline declaration ");
                break;
            }
        }
    }

}
