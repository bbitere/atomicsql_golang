
using System;
using System.Collections.Generic;
using System.Data.SqlTypes;
using System.Dynamic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using Gemstone.CharExtensions;


namespace goscanner.ConvCommon
{
    internal static class Utils1
    {
        public static V Get<T, V>( this Dictionary<T,V> dict, T key ) 
        {
            if( dict.ContainsKey( key ) )
            { 
                return dict[key];
            }
            return default(V); 
        }
        public static void  Nop()
        {
        }

        public static string NormaliseDirPath(string path)
        {
            if( path == null )
                return null;
            return path.Replace("/", Path.DirectorySeparatorChar.ToString() );
        }
            

        
        public static string getFileName( string file1)
        {
            var fi = new FileInfo( file1 );
            return fi.Name;
        }
        public static string getAbsolutePath( string file1)
        {
            Uri file = new Uri(file1);            
            return file.AbsolutePath;
        }
        public static string corectPathSeparator(string path)
        {
            if( path == null)   
                return null;
            path = path.Replace('/', Path.DirectorySeparatorChar);
            return path.Replace('\\', Path.DirectorySeparatorChar);
        }
        public static string getAbsoluteDirPath( string directoryPath, string sep = "")
        {
            if( sep != "" )
            {
                var parts = directoryPath.Split( sep);
                var arr = new List<string>();

                foreach( var part_ in parts ) {

                    var part = part_.Trim();    
                    Uri file = new Uri(part);            
                    var dir = file.AbsolutePath;
                    if( dir.EndsWith("/") )
                        dir = dir.Substring(0, dir.Length - 1);
                    dir = corectPathSeparator(dir);
                    arr.Add(dir);
                }
            
                return string.Join( sep, arr.ToArray() );
            }else
            {
                Uri file = new Uri(directoryPath);            
                var dir = file.AbsolutePath;
                if( dir.EndsWith("/") )
                    dir = dir.Substring(0, dir.Length - 1);
                dir = corectPathSeparator(dir);
                return dir;
            }
        }
        public static string getDirPathOfFile( string file1)
        {
            var file = new FileInfo(file1);            
            var dir = file.DirectoryName;
            if( dir.EndsWith("/") )
                dir = dir.Substring(0, dir.Length - 1);
            
            return dir;
        }
        

        public static string getRelativePath(string folder1, string file1)
        {
            Uri file = new Uri(file1);            
            Uri folder = new Uri(folder1);

            var relativePath = Uri.UnescapeDataString(
                folder.MakeRelativeUri(file).ToString() );

            return relativePath;
        }
        public static string CreateMD5(string input)
        {
            // Use input string to calculate MD5 hash
            using (System.Security.Cryptography.MD5 md5 = System.Security.Cryptography.MD5.Create())
            {
                byte[] inputBytes = System.Text.Encoding.ASCII.GetBytes(input);
                byte[] hashBytes = md5.ComputeHash(inputBytes);

                return Convert.ToHexString(hashBytes); // .NET 5 +

                // Convert the byte array to hexadecimal string prior to .NET 5
                // StringBuilder sb = new System.Text.StringBuilder();
                // for (int i = 0; i < hashBytes.Length; i++)
                // {
                //     sb.Append(hashBytes[i].ToString("X2"));
                // }
                // return sb.ToString();
            }
        }

        public static string Base64Encode(string plainText) 
        {
            var plainTextBytes = Encoding.UTF8.GetBytes(plainText);
            return Convert.ToBase64String(plainTextBytes);
        }
        public static string Base64Decode( string base64EncodedData ) 
        {
            var base64EncodedBytes = Convert.FromBase64String(base64EncodedData);
            return Encoding.UTF8.GetString(base64EncodedBytes);
        }

        public static string GetFirstPart(string path, string separator)
        {
            var parts = path.Split(separator);
            return parts[0];    
        }
        public static string GetLastPart(string path, string separator)
        {
            var parts = path.Split(separator);
            return parts[ parts.Length -1 ];    
        }

        public static bool isPublicIdentGoLang( string identifName)
        {
            if( identifName == null )
                return false;

            if( identifName[0] == identifName[0].ToUpper())
                return true;
            return false;
        }

        public static string readFile(string filename )
        {
            if( filename.StartsWith("http://")
             || filename.StartsWith("https://"))
            {
                using (HttpClient httpClient = new HttpClient())
                {
                    try
                    {
                        // Trimiteți cererea HTTP GET și obțineți răspunsul
                        var response = httpClient.GetAsync(filename);
                        response.Wait(10000);
                        response.Result.EnsureSuccessStatusCode();

                        // Asigurați-vă că cererea a avut succes
                        //response.EnsureSuccessStatusCode();

                        // Citirea conținutului răspunsului
                        var content = response.Result.Content.ReadAsStringAsync();
                        content.Wait(20000);

                        // Afișați conținutul
                        return content.Result;
                    }
                    catch (HttpRequestException ex)
                    {
                        Console.WriteLine($"Eroare HTTP: {ex.Message}");
                    }
                }
            }else
            {
                string content = File.ReadAllText(filename);
                return content;
            }
            return null;
        }
    }
}
