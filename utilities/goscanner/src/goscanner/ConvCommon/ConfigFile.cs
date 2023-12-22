using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using Antlr4.Runtime;
using System.Reflection.Metadata;
using System.Diagnostics;
using System.Text.RegularExpressions;

namespace goscanner.ConvCommon
{
    
    public class ConfigFile
    { 
        public int NumErrors { get; set; }

        public bool UseMetadata = true;//ca se busest importurile la MetaDataimport
        public bool bCollectMethods = false;
        //public List<string> ExportEnums = new List<string>();

        public Dictionary<string, Dictionary<string, string> > 
            DictExportEnums = new Dictionary<string, Dictionary<string, string> >();
        
        public bool   DisableWarnings = true;
        public string TypeLang;
        public string SourcePathDir;
        public string SourcePathDir2;
        //public List<string> SourceFiles = new List<string>();
        public string Class_Base_MODEL;

        public string RootProject;
        public string GitProject;


        

        public void ParseConfigFile( string pathFile)
        {
            try
            {
                var text = File.ReadAllText(pathFile);
                var lines = text.Split( '\n' );
                for( var iLine = 0; iLine < lines.Length; iLine++ ) 
                {
                    var line1 = lines[ iLine ].Trim();
                    if( line1.StartsWith("#") )
                        continue;

                    var token = getToken( ref line1 );
                    if( token == "$" )
                    {
                        var token1 = getToken( ref line1 );
                        var token2 = getToken( ref line1 );
                        if( token2 == "=" )
                        {
                            var propertyValue = getPropertyValue( line1.Trim());
                            setupSLineProperty( token1, propertyValue );
                        }
                    }else
                    if( token == "@" )
                    {
                        var content = "";
                        iLine++;
                        for( ; iLine < lines.Length; iLine++ ) 
                        {
                            var line2 = lines[ iLine ].Trim();
                            if( line2 == "@#@" )
                                break;
                            content += lines[ iLine ];
                            content += "\r\n";
                        }

                        var token1 = getToken( ref line1 );

                        setupMLineProperty( token1, content );
                    }

                }

            }catch( Exception e )
            {
                Console.WriteLine(e.Message);
            }

            updateProperties();
        }
        private string getPropertyValue(string val)
        {
            var retValue = "";
            var prevIdx = 0;
            for( var idx = 0; idx < val.Length; )
            {
                idx = val.IndexOf("%", prevIdx);
                if( idx >=0 )
                {
                    retValue += val.Substring( prevIdx, idx- prevIdx);

                    idx = idx+1;
                    var idx2 = val.IndexOf("%", idx);
                    var strProp = val.Substring( idx, idx2-idx);

                    var v = Environment.GetEnvironmentVariable(strProp);
                    if( v != "" && v != null)
                    {
                        retValue += v;
                    }else
                    {
                        retValue += $"%{strProp}%";
                    }
                    prevIdx = idx2+1;
                    idx = prevIdx;
                }else
                    break;
            }

            retValue += val.Substring( prevIdx);
            return retValue;
        }
        protected virtual void setupSLineProperty(string token1, string content)
        {
            Debugger.Break();
        }
        protected virtual void updateProperties()
        {
            Debugger.Break();
        }
        protected virtual void setupMLineProperty(string token1, string content)
        {
            Debugger.Break();
        }

        private string cleanStringValue(string s )
        {
            if( s == null )
                return s;
            if( s.Length >= 2 && s.StartsWith("\"" ) && s.EndsWith("\"u8" ) )
                return s.Substring( 1, s.Length-1 -"\"u8".Length ); 
            if( s.Length >= 2 && s.StartsWith("\"" ) && s.EndsWith("\"" ) )
                return s.Substring( 1, s.Length-1 -1 ); 
            if( s.Length >= 2 && s.StartsWith("'" ) && s.EndsWith("'" ) )
                return s.Substring( 1, s.Length-1 -1 ); 
            if( s.Length >= 2 && s.StartsWith("`" ) && s.EndsWith("`" ) )
                return s.Substring( 1, s.Length-1 -1 ); 
            return s;
        }
        
        public void AddEnumExpression( string identifier, Dictionary<string, string> dictValues )
        {
            foreach( var it in dictValues.Keys ) 
            {
                dictValues[it] = cleanStringValue( dictValues[it] );
            }
            this.DictExportEnums[ identifier ] = dictValues;
        }
        protected List<string> processExportListItems(string content)
        {
            var result = new List<string>();

            var parts = content.Split( "\n", StringSplitOptions.RemoveEmptyEntries);
            foreach( var part in parts ) 
            {
                var line = part.Trim();
                if( line !="" )
                {
                    var line1 = getPropertyValue( line );
                    result.Add( line1 );
                }
            }
            return result;
        }

        private static string getToken( ref string line)
        {
            var i = 0;
            for( ; i < line.Length; i++ ) 
            {
                var ch = line[ i ];
                if( (ch == '\t'
                 || ch == '\n'
                 || ch == '\r'
                 || ch == ' '))
                {
                }
                else
                    break;
            }

            var content = "";
            for( ; i < line.Length; i++ ) 
            {
                var ch = line[ i ];
                if( ch == '\t'
                 || ch == '\n'
                 || ch == '\r'
                 || ch == ' ')
                {
                    line = line.Substring( i );
                    return content;
                }
                content += ch;
            }
            line = line.Substring( i );
            return content;
        }
    }
    public class TemplateItem
    {
        public string Name { get; set; }
        public string Text { get; set; }

        public TemplateItem( string name, string text) 
        { 
            Name = name;
            Text = text.Replace("\r\n", "\n");
        }

        public string ConvertTemplate( Dictionary<string, string>dict)
        {
            var text = this.Text;
            var templateName = this.Name;

            foreach( var it in dict )
            {
                var val = $"@@{{{it.Key}}}";

                text = text.Replace( val, it.Value );
            }
            var idx = text.IndexOf("@@{");
            if( idx >= 0)
            {
                var idx2 = text.IndexOf( "}", idx );
                var item = text.Substring( idx, idx2+1-idx);

                Console.WriteLine($"Error : tag {item} is still present in {templateName} ");
            }
            return text;
        }

    }

}
