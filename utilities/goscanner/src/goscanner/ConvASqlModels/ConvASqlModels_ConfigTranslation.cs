using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using Antlr4.Runtime;
using System.Reflection.Metadata;
using System.Text.RegularExpressions;

namespace goscanner.ConvASqlModels
{
    
    public class ConvASqlModel_ConfigTranslation : goscanner.ConvCommon.ConfigFile
    { 
        // -o -i -h -w=.\export_svc_config.txt C:\Work\SlotMonitor\SlotGames\server-main\NewtonGamesGo\pkg\controllers 
        // C:\Work\SlotMonitor\SlotGames\game-originals-main\tools

        public const string Version = "1.0.0";

        public string  CLASS_InGeneralData;
        public string  CLASS_OutGeneralData;
        public string  FLD_ID;
        public string  DELIMETER;

        public string JsonOutputFile;

        
        protected override void setupSLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "JsonOutputFile":  
                { 
                    this.JsonOutputFile = content; 
                    if( this.JsonOutputFile.Contains("{date}") )
                    {
                        var date = DateTime.Now.ToString("yyyy.MM.dd-HH.mm.ss");
                        this.JsonOutputFile = this.JsonOutputFile.Replace("{date}", date);
                    }
                    if( this.JsonOutputFile.Contains("{date1}") )
                    {
                        //2023.05.12.122301.name.sql
                        //const string RECOGNIZE_FILE_PATTERN=@"^\d{4}\.\d{2}\.\d{2}\.(\d{2}|\d{3}|\d{4}|\d{6})\.(.+)$";
                        var date = DateTime.Now.ToString("yyyy.MM.dd.HHmmss");
                        this.JsonOutputFile = this.JsonOutputFile.Replace("{date1}", date);
                    }
                    break;
                }
                case "SourcePathDir":   this.SourcePathDir = content; break;
                case "SourcePathDir2":  this.SourcePathDir2 = content; break;

                case "Class_Base_MODEL":this.Class_Base_MODEL = content; break;
                case "DELIMETER":       this.DELIMETER = content; break;
                    
                
                default: 
                    Console.WriteLine( $"Not identified token {token1} in single line property ");
                break;
            }
        }
        /*
        protected override void modifyExcludeFiles(ref Regex m_excludeExpression)
        {
            //System.Text.RegularExpressions.Regex
            m_excludeExpression = new Regex( "\\.gen\\.go");
        }
        */

        protected override void updateProperties()
        {

           /*
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
            */

            if( this.Class_Base_MODEL != null )
            {
                GenerateJsonModels.Class_Generic_MODEL = this.Class_Base_MODEL;
                GenerateJsonModels.BASE_NAME_FLD = this.Class_Base_MODEL;
                GenerateJsonModels.BASE_NAME_FLD1 = this.Class_Base_MODEL;
            }
            if( this.DELIMETER != null )
                GenerateJsonModels.DELIMETER = this.DELIMETER;
            

        }
        protected override void setupMLineProperty(string token1, string content)
        {
            
        }
    }

}
