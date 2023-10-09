using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.IO;
using Npgsql;
using Npgsql.Logging;

namespace src_tool
{
class Utils
{
    public static void Nop(){ }
    
    public static string getFileInfoName( FileInfo f0Info)
    {
        if( f0Info.Extension == "")
            return f0Info.Name;
        else
            return f0Info.Name.Replace(f0Info.Extension, "");
    }
    public static string getFileInfoName( string f0)
    {
        var f0Info = new FileInfo(f0);
        if( f0Info.Extension == "")
            return f0Info.Name;
        else
            return f0Info.Name.Replace(f0Info.Extension, "");
    }
    public static string normalizePath(string path)
    {
        if( path == null )
            return path;

        return path.Replace("/", "\\");
    }

    public static string normalizeName(string name )
    {
        int intVal = 0;
        var outName = name.Replace("-", "_");
        if( outName == "in" 
            || outName == "for"
            || outName == "if"
            || outName == "while"
            || outName == "do"
            || outName == "switch"
            || int.TryParse( outName, out intVal ))
        {
            outName = "_"+outName;
        }
        return outName;
    }

    public static List<T> getListFromDict<T>(Dictionary<string, T> dict)
    {
        var lst = new List<T>();
        if( dict == null )
            return null;

        foreach( var it in dict)
        {
            lst.Add(it.Value);
        }
        return lst;
    }

    public static Dictionary<string, T> getDictFromList<T>(List<T> list, Func<T, string> fn)
    {
        var dict = new Dictionary<string,T>();
        if( list == null )
            return null;

        foreach( var it in list)
        {
            var key = fn( it );
            dict[key] = it;
        }
        return dict;
    }

    public static Dictionary<T, U> getDictionaryDifference<T, U>(
                                    Dictionary<T, U> dicB,
                                    Dictionary<T, U> dicA)
    {
        var dict = new Dictionary<T, U>();
        // Walk A, and if any of the entries are not
        // in B, add them to the result dictionary.

        foreach (var kvp in dicB)
        {
            if (!dicA.ContainsKey(kvp.Key))
            {
                dict[kvp.Key] = kvp.Value;
            }
        }
        return dict;
    }
    public static Dictionary<T, Tuple<U,U> > getDictionaryUnion<T, U>(
                                    Dictionary<T, U> dicB,
                                    Dictionary<T, U> dicA)
    {
        var dict = new Dictionary<T, Tuple<U,U> >();
        // Walk A, and if any of the entries are not
        // in B, add them to the result dictionary.

        foreach (var kvp in dicA)
        {
            if (dicB.ContainsKey(kvp.Key))
            {
                dict[kvp.Key] = new Tuple<U,U>( dicB[kvp.Key], dicA[kvp.Key]);
            }
        }
        return dict;
    }

}
}


                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              