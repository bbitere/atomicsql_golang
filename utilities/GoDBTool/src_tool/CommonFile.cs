using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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
        public String sqlType = "";
        public DbTable ForeignKey = null;
        public DbColumn initSql(string sqlName, string sqlType, bool bIsNullable)
        {
            this.sqlName = sqlName;
            this.sqlType = sqlType;
            this.bIsIdentity = false;
            this.bIsNullable = bIsNullable;
            return this;
        }

        public DbColumn initSqlPrimary(string sqlName)
        {
            this.sqlName = sqlName;
            this.sqlType = "";
            this.bIsIdentity = true;
            this.bIsNullable = false;
            return this;
        }
        public DbColumn initLangSql(
            string langName, string langName2, string langType,
            string sqlName, string sqlType,
            bool bIsNullable, bool bIsIdentity, 
            DbTable ForeignKey )
        {
            this.langName = langName;
            this.langName2 = langName2;
            this.langType = langType;
            this.bIsIdentity = bIsIdentity;
                    
            this.sqlName = sqlName;
            this.sqlType = sqlType;
            this.bIsNullable = bIsNullable;
            this.ForeignKey = ForeignKey;
            return this;
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

        public string getPluralTableNameModel(){ return LangTableNameModel;}

        public DbTable initSql(string sqlName, DbColumn primaryKey)
        {
            this.SqlTableNameModel = sqlName;
            this.LangTableNameModel = sqlName;
            this.PrimaryColumn = primaryKey;
            return this;
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
