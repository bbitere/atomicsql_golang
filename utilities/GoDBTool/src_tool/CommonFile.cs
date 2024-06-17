using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Diagnostics;

namespace src_tool
{
    public static class EOuputLang
    { 
        public const string GoLang= "go";        
    }
    public class DbColumn
    {
        public string langName = "";//UserID
        public string langName2 = "";//pt foreignkeys, tre sa memoreze si numele original al fkey. User_ID
        public string langType = "";
        
        public string sqlName = "";        
        public bool bIsIdentity = false;        
        public bool bIsNullable = false;
        public string sqlType = "";
        public DbTable ForeignKey = null;
        public string descriptionTag;
        public DbColumn initSql(string sqlName, string sqlType, bool bIsNullable)
        {
            this.sqlName = sqlName;
            checkedSqlName(this.sqlName);
            this.sqlType = sqlType;
            this.bIsIdentity = false;
            this.bIsNullable = bIsNullable;
            return this;
        }

        public DbColumn initSqlPrimary(string sqlName)
        {
            this.sqlName = sqlName;
            checkedSqlName(this.sqlName);
            this.sqlType = "";
            this.bIsIdentity = true;
            this.bIsNullable = false;
            return this;
        }
        public DbColumn initLangSql(
            string langName, string langName2, string langType,
            string sqlName, string sqlType,
            bool bIsNullable, bool bIsIdentity, 
            DbTable ForeignKey,
            string tags)
        {
            this.langName = langName;
            this.langName2 = langName2;
            this.langType = langType;
            this.bIsIdentity = bIsIdentity;
                    
            this.sqlName = sqlName;
            checkedSqlName(this.sqlName);
            this.sqlType = sqlType;
            this.bIsNullable = bIsNullable;
            this.ForeignKey = ForeignKey;
            this.descriptionTag = tags;
            return this;
        }
        void checkedSqlName(string sqlName)
        {
            if( sqlName.IndexOf("\"") >= 0
             || sqlName.IndexOf(":") >= 0)
                Debugger.Break();
        }
        public string LangTypeIsPointer()
        {
            if( this.langType != null)
            {
                if( this.langType.StartsWith("*") )
                {
                    return this.langType.Replace("*", "");
                }
            }

            return "";
        }
    }
    public class DbTable
    {
        public string Schema;// go name
        public string LangTableNameModel;// go name
        public string SqlTableNameModel;//sqlname
        public List<DbColumn> columns = new List<DbColumn>();
        public DbColumn PrimaryColumn = null;

        public string json;//used for private purpose

        public string getPluralTableNameModel()
        {
            return LangTableNameModel;
        }

        public DbTable initSql( string sqlName, DbColumn primaryKey )
        {
            this.SqlTableNameModel = sqlName;
            this.LangTableNameModel = sqlName;
            this.PrimaryColumn = primaryKey;
            return this;
        }


        /**
         * 
         * `json:"ID,omitempty"`
         * `json:"UUID"`
         *  `json:"-"`
         */ 
        public static string parseSqlName(string tags, GenericDialect dialect)
        {
            if( tags == "")
                return "";

            if( tags == GoModelTemplate.FIELD_IS_OMITTED_INTEGRAL)
                return "";

            tags = tags.Trim();
            if( tags.StartsWith("\"") && tags.EndsWith("\""))
                tags = tags.Substring(1, tags.Length-2);


            var bFoundDescr = false;
            var arrTags = tags.Split(new string[]{ "  ", }, StringSplitOptions.RemoveEmptyEntries );
            foreach( var tag1 in arrTags)
            {
                var tag = tag1.Trim();
                if( tag.StartsWith("bson:\\\"") )
                {
                    bFoundDescr = true;
                    if( dialect.isNoSql() )
                    {
                        var tags1 = tag.Replace("bson:\\\"", "");
                        tags1 = tags1.Replace("\\\"\"", "");
                        tags1 = tags1.Replace("\\\"", "");
                        
                        var p = tags1.Split(',');
                        return p[0].Trim();
                    }
                }else
                if( tag.StartsWith("json:\\\"") )
                {
                    bFoundDescr = true;

                    if( !dialect.isNoSql() )
                    {
                        var tags1 = tag.Replace("json:\\\"", "");
                        tags1 = tags1.Replace("\\\"\"", "");
                        tags1 = tags1.Replace("\\\"", "");
                    
                        var p = tags1.Split(',');
                        return p[0].Trim();
                    }
                } else
                if( tag.StartsWith("atmsql:\\\"") )
                {
                    bFoundDescr = true;

                    if( !dialect.isNoSql() )
                    {
                        var tags1 = tag.Replace("atmsql:\\\"", "");
                        tags1 = tags1.Replace("\\\"\"", "");
                        tags1 = tags1.Replace("\\\"", "");
                    
                        var p = tags1.Split(',');
                        return p[0].Trim();
                    }
                }
            }
            
            if( !bFoundDescr )
            {
                Console.WriteLine($"json definition of description field is incomplete :{tags}");
            }
            return "";
        }
        public bool HasSerializableAttribute(GenericDialect dialect)
        {
            foreach(var col in columns)
            {
                var sqlName = parseSqlName(col.descriptionTag, dialect);
                if( sqlName != "")
                    return true;
            }
            return false;
        }
    }

    public class FKRootTgt
    {
        public DbTable tableRoot;
        public DbColumn columnRoot;
        public DbTable tableTgt;
        public FKRootTgt(

            DbTable _tableRoot,
            DbColumn _columnRoot,
            DbTable _tableTgt
            )
        {
            this.tableRoot = _tableRoot;
            this.columnRoot = _columnRoot;
            this.tableTgt = _tableTgt;
        }
    }
}
