using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;

namespace src_tool
{
    public static class ELangSql
    {
        public static readonly string PostgresSql = "postgres";
        public static readonly string MySql = "mysql";
        public static readonly string MsSql = "mssql";
        public static readonly string Oracle = "oracle";
    }

    public class ConfigFile
    {
        public TemplateItem Templ_GoLangOrmFile;
        public TemplateItem Templ_GoLangModelFile;

        public TemplateItem Templ_GoLang_SchemaDefItem;
        public TemplateItem Templ_GoLang_SchemaDefItem_Col;
        public TemplateItem Templ_GoLang_ForeignKey;
        
        public string OutputDBContextFile;
        public string Models_Extension;
        public string ModelsOutputDir;
        public string BaseModelName;
        public string ImportPackageModels;
        public string ImportPackageOrm;
        
        public string PackageGenSql;
        public string ConnectionString;
        public string SqlLang = ELangSql.PostgresSql;
        public string Delimeter;
        public string DirJsons;

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
            this.checkProps();
        }

        private string getPropertyValue(string val)
        {
            var retValue = "";
            var prevIdx = 0;
            for( var idx = 0; idx < val.Length; )
            {
                idx = val.IndexOf("%", idx);
                if( idx >=0 )
                {
                    retValue += val.Substring( prevIdx, idx);

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

        private void setupSLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "OutputDBContextFile": this.OutputDBContextFile = content; break;
                case "Models_Extension": this.Models_Extension = content; break;
                    
                case "ModelsOutputDir":     this.ModelsOutputDir = content; break;
                case "BaseModelName":       this.BaseModelName = content; break;
                case "FullName_PackageModels": this.ImportPackageModels = content; break;
                case "FullName_PackageOrm":    this.ImportPackageOrm = content; break;
                case "FullName_PackageGenSql":       this.PackageGenSql = content; break;
                    
                case "ConnectionString":    this.ConnectionString = content; break;
                case "SqlLang":             this.SqlLang = content; break;
                case "DELIMETER":           this.Delimeter = content; break;
                case "DirJsons":           this.DirJsons = content; break;
                    
                
                    
                default: 
                    Console.WriteLine( $"Not identified token {token1} in single line property ");
                break;
            }
        }
        
        private void setupMLineProperty(string token1, string content)
        {
            switch( token1 ) 
            {
                case "Templ_GoLangOrmFile":             this.Templ_GoLangOrmFile     = new TemplateItem( token1, content ); break;
                case "Templ_GoLangModelFile":           this.Templ_GoLangModelFile   = new TemplateItem( token1, content ); break;
                case "Templ_GoLang_SchemaDefItem":      this.Templ_GoLang_SchemaDefItem     = new TemplateItem( token1, content ); break;
                case "Templ_GoLang_SchemaDefItem_Col":  this.Templ_GoLang_SchemaDefItem_Col = new TemplateItem( token1, content ); break;
                case "Templ_GoLang_ForeignKey":         this.Templ_GoLang_ForeignKey = new TemplateItem( token1, content ); break;

                default: 
                    Console.WriteLine( $"Not identified token {token1} in multiline declaration ");
                break;
            }
        }

        void checkProps()
        {
            var retError = checkAllProps();
            if( retError != "" )
            {
                Console.WriteLine( retError );
            }
        }
        string checkAllProps()
        {
            if( this.OutputDBContextFile == null )
                return $"Missing prop OutputDBContextFile";

            if( this.Models_Extension == null)
                return $"Missing prop Models_Extension";

            if( this.ModelsOutputDir == null)
                return $"Missing prop ModelsOutputDir";

            if( this.BaseModelName == null)
                return $"Missing prop BaseModelName";

            if( this.ImportPackageModels == null)
                return $"Missing prop ImportPackageModels";

            if( this.ImportPackageOrm == null)
                return $"Missing prop ImportPackageOrm";

            if( this.PackageGenSql == null)
                return $"Missing prop PackageGenSql";

            if( this.ConnectionString == null)
                return $"Missing prop ConnectionString";

            if( this.SqlLang == null)
                return $"Missing prop SqlLang";

            if( this.Delimeter == null)
                return $"Missing prop Delimeter";

            if( this.DirJsons == null)
                return $"Missing prop DirJsons";

            if( this.Templ_GoLangOrmFile == null)
                return $"Missing prop Templ_GoLangOrmFile";

            if( this.Templ_GoLangModelFile == null)
                return $"Missing prop Templ_GoLangModelFile";

            if( this.Templ_GoLang_SchemaDefItem == null)
                return $"Missing prop Templ_GoLang_SchemaDefItem";

            if( this.Templ_GoLang_SchemaDefItem_Col == null)
                return $"Missing prop Templ_GoLang_SchemaDefItem_Col";

            if( this.Templ_GoLang_ForeignKey == null)
                return $"Missing prop Templ_GoLang_ForeignKey";

            return "";
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
        public string UseTemplate(string original, TemplateItem template, Dictionary<string, string>dict)
        { 
            if( template == null) {
                return original;
            }
            return template.ConvertTemplate( dict);
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
