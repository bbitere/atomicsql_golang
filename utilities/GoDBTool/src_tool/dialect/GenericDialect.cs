using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace src_tool
{
    
    
    public abstract class GenericDialect
    {
        string currentFile = "";
        Dictionary<string, string> dictError = new Dictionary<string, string>();


        public delegate void TProcessData(GenericDialect dialect, GenericDialectArg arg1 );
        public TProcessData fnProcessData;

        
        public abstract string getSql();
        public abstract bool isNoSql();

        
        public abstract string SqlSeparator();
        public abstract string dropColumn(DbTable table,DbColumn column);
        public abstract string addColumn(DbTable table,DbColumn column);
        public abstract string updateColumn(DbTable table, DbColumn column, DbColumn columnPrev);

        

        public abstract string dropTable(DbTable table);
        public abstract string addTable(DbTable table);
        public abstract string updateTable(DbTable table, DbTable tablePrev);
        
        public abstract string dropFKConstrictor(DbTable table,DbColumn column);
        public abstract string addFKConstrictor(DbTable table,DbColumn column);

        public abstract string getSqlType( string langType, ref bool bIsNullable, string nameOfColumn="");

        public abstract string getDefaultSchema();

        public abstract bool startConnection(GenericDialectArg arg);

        public abstract string getProperty(string name, string tableName, string colName, string colValue);
        public abstract void insertProperty(
            string propName, string value, 
            string tableName, string colName, string colValue);
        public abstract void updateProperty( 
            string propName, string value, 
            string tableName, string colName, string colValue);

        public abstract void execScript( string scriptTxt);
        public abstract bool readConstraintors( Dictionary<string, DbTable> tables  );
        public abstract Dictionary<string, DbTable> readTables( string SqlLang );

        public abstract string getGoLangTypeIntFk( DbColumn column, ref Dictionary<string, string> packageImports );
        public abstract string getGoLangType( DbColumn column, ref Dictionary<string, string> importPackage );


        //------------------------------------------------------------
        public void setCurrentFile(string file)
        {
            this.currentFile = file;
        }
        
        public string printError(string err)
        {
            if( !dictError.ContainsKey(err))
            {
                dictError[err] =err;
                Console.WriteLine(err);

                if( this.currentFile != "")                
                    Console.WriteLine($"Error In File: {this.currentFile}");

                Console.WriteLine("");
            }
            return "";
        }
        public string cleanNameGoStruct(string name)
        { 
            return GenericDialect.CleanNameGoStruct(name);
        }
        public static string CleanNameGoStruct(string name)
        {
            //I renmove the pointer name
            var parts = name.Split('.');
            if( parts.Length > 0 )
            {
                return parts[ parts.Length -1].Replace("*","");
            }
            return name.Replace("*","");
        }

        public static GenericDialect GetDialectByName(string SqlLang)
        {
            if( SqlLang == ELangSql.MongoDB )
            {
                return new MongodbDialect();
            }else
            if( SqlLang == ELangSql.PostgresSql )
            {
                return new PostgressDialect();
            }else
            if( SqlLang == ELangSql.MySql )
            {
                return new MySqlDialect();
            }else
            {
                Console.WriteLine( $"Error: sql dialect '{SqlLang}' not implemented");
                return null;
            }
        }

        public bool isLongType(string sqlColumnName)
        {
            if (sqlColumnName != null && sqlColumnName != "")
            {
                if( sqlColumnName.StartsWith("_"))
                    return true;
            }
            return false;
        }
        
    }
}
