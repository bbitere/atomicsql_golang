/* Copyright (c) 2012-2017 The ANTLR Project. All rights reserved.
 * Use of this file is governed by the BSD 3-clause license that
 * can be found in the LICENSE.txt file in the project root.
 */
using System;
using System.Collections.Generic;
using System.Text;

namespace Antlr4.Runtime.Misc
{
    public static class StaticUtils
    {
        public static string ToString<T>(this IEnumerable<T> list)
        {
            return "[" + Utils.Join(", ", list) + "]";
        }
    }

    public class Utils
    {
        public static void Nop(){ }
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
        public static string Join<T>(string separator, IEnumerable<T> items)
        {
            return string.Join(separator, items);
        }

        public static int NumNonnull(object[] data)
        {
            int n = 0;
            if (data == null)
            {
                return n;
            }
            foreach (object o in data)
            {
                if (o != null)
                {
                    n++;
                }
            }
            return n;
        }

        public static void RemoveAllElements<T>(ICollection<T> data, T value)
        {
            if (data == null)
            {
                return;
            }
            while (data.Contains(value))
            {
                data.Remove(value);
            }
        }

        public static string EscapeWhitespace(string s, bool escapeSpaces)
        {
            StringBuilder buf = new StringBuilder();
            foreach (char c in s.ToCharArray())
            {
                if (c == ' ' && escapeSpaces)
                {
                    buf.Append('\u00B7');
                }
                else
                {
                    if (c == '\t')
                    {
                        buf.Append("\\t");
                    }
                    else
                    {
                        if (c == '\n')
                        {
                            buf.Append("\\n");
                        }
                        else
                        {
                            if (c == '\r')
                            {
                                buf.Append("\\r");
                            }
                            else
                            {
                                buf.Append(c);
                            }
                        }
                    }
                }
            }
            return buf.ToString();
        }

        public static void RemoveAll<T>(IList<T> list, Predicate<T> predicate)
        {
            int j = 0;
            for (int i = 0; i < list.Count; i++)
            {
                T item = list[i];
                if (!predicate(item))
                {
                    if (j != i)
                    {
                        list[j] = item;
                    }
                    j++;
                }
            }
            while (j < list.Count)
            {
                list.RemoveAt(list.Count - 1);
            }
        }

        /// <summary>Convert array of strings to string&#x2192;index map.</summary>
        /// <remarks>
        /// Convert array of strings to string&#x2192;index map. Useful for
        /// converting rulenames to name&#x2192;ruleindex map.
        /// </remarks>
        public static IDictionary<string, int> ToMap(string[] keys)
        {
            IDictionary<string, int> m = new Dictionary<string, int>();
            for (int i = 0; i < keys.Length; i++)
            {
                m[keys[i]] = i;
            }
            return m;
        }
    }
}
