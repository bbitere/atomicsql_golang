using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using src_tool;

namespace src_tool.templates
{
public class GoDBContext:GoModelTemplate
{
	public const string pkgModels = "m";
	//public const string pkgOrm = "orm";
	public const string _DBTable = "DBTable";
	//public const string Template_V = "V";
	//public const string Template_Def = pkgOrm+".IGeneric_MODEL";
		
		
	
    public static string GetTemplateDBContext_def1( DbTable table )
    {
		if(table.LangTableNameModel == "Employee1")
			Utils.Nop();

		var tableInstName = table.getPluralTableNameModel();
        var tableName = table.LangTableNameModel;
		var space1 = new String( ' ', Math.Max(1, 20 -$"{tableInstName}".Length) );
		var space2 = new String( ' ', Math.Max(1, 20 -$"{tableName}".Length) );

        var text = $@"
			{tableInstName}{space1}*{pkgOrm}.{_DBTable}[{pkgModels}.{tableName}]
			{tableName}{PREF_VAR}{space2} {pkgModels}.{PREF_DEF}{tableName}
		";
        return text;
    }

	public static string GetTemplateDBContext_tableInit( DbTable table )
	{
		var tableName = table.LangTableNameModel;
		var sqlTableNameModel = table.SqlTableNameModel;
		var tableInstName = table.getPluralTableNameModel();

		var space1 = new String( ' ', Math.Max(1, 30 -$"ret.{tableInstName} =".Length) );

		var text = $@"			
			ret.{tableInstName} ={space1}(new({pkgOrm}.{_DBTable}[{pkgModels}.{tableName}])).Constr(""{sqlTableNameModel}"", ""{tableName}"", &ret.DBContextBase)";
		return text;
	}

	public static string GetTemplateDBContext_tableInit_templateV( DbTable table )
	{
		var tableName = table.LangTableNameModel;
		var sqlTableNameModel = table.SqlTableNameModel;
		var tableInstName = table.getPluralTableNameModel();

		var space1 = new String( ' ', Math.Max(1, 30 -$"ret.{tableInstName} =".Length) );

		var text = $@"			
			ret.{tableInstName} ={space1}(new({pkgOrm}.{_DBTable}[{pkgModels}.{tableName}])).Constr(""{sqlTableNameModel}"", ""{tableName}"", &ret.DBContextBase)";
		return text;
	}

	public static string GetTemplateVardef(string nameModel, string txtDefVarModel)
	{
			/*
		var text = $@"
			var {PREF_VAR}{nameModel} = {pkgModels}.{PREF_DEF}{nameModel} {{
	        
	        {txtDefVarModel},

		}}";
			*/

		var text = $@"
			_this.{nameModel}{PREF_VAR} = {pkgModels}.{PREF_DEF}{nameModel} {{
	        {txtDefVarModel},
			}}

		";
		return text;
	}

	private static string GetTemplateForeignKey(
		ConfigFile config,
		string fk, DbTable tableRoot, DbColumn colRoot, DbTable tableTgt )
	{
		var tgtTable_sqlName = tableTgt.SqlTableNameModel;
		var tgtFldID_sqlName = tableTgt.PrimaryColumn.sqlName;

		var rootTable_sqlName = tableRoot.SqlTableNameModel;
		var rootFldFk_sqlName = colRoot.sqlName;
		var rootFldFk_langName = colRoot.langName;
		var rootFldFk_lang2Name = colRoot.langName2;

		var str = 
			$@" ""{fk}"":
			{{  
				TgtTable_sqlName:   ""{tgtTable_sqlName}"", 
				TgtFldID_sqlName:   ""{tgtFldID_sqlName}"", 
					
				RootTable_sqlName:  ""{rootTable_sqlName}"", 

				RootFldFk_sqlName:  ""{rootFldFk_sqlName}"", 	
				RootFldFk_langName: ""{rootFldFk_langName}"", 
				RootFldFk_lang2Name:""{rootFldFk_lang2Name}"", 
			}},
			";	

		str = config.UseTemplate( str, config.Templ_GoLang_ForeignKey,
				new Dictionary<string,string>()
				{ 
					{ "tgtTable_sqlName", tgtTable_sqlName},
					{ "tgtFldID_sqlName", tgtFldID_sqlName},
					{ "rootTable_sqlName", rootTable_sqlName},
					{ "rootFldFk_sqlName", rootFldFk_sqlName},
					{ "rootFldFk_langName", rootFldFk_langName},
					{ "rootFldFk_lang2Name", rootFldFk_lang2Name},
					{ "fk", fk},					
				});
				str = str.Replace("\r\n", "\n");

		return str;
	}

    public static string GetTemplateDBContext(
		GenericDialect dialect,
		ConfigFile config,
		string package, string import_package_models,
		string import_package_orm,
		string tables_DefVar,
		Dictionary<string,FKRootTgt> ForeignKeys,
		List<DbTable> tables )
    {

		var tables_def1_list = new List<string>();
		tables.ForEach( x=> { tables_def1_list.Add( GetTemplateDBContext_def1(x) );} );
		string tables_def1 = String.Join("", tables_def1_list);

		var tables_init1_list = new List<string>();
		tables.ForEach( x=> { tables_init1_list.Add( GetTemplateDBContext_tableInit(x) );} );
		string init_tables = String.Join("", tables_init1_list);

		var list_tables_TemplateV = new List<string>();
		tables.ForEach( x=> { list_tables_TemplateV.Add( GetTemplateDBContext_tableInit_templateV(x) );} );
		string init_tables_TemplateV = String.Join("", list_tables_TemplateV);

		var listSchemaDefItem = new List<string>();
		var unusedPackageImports = new Dictionary<string, string>();

		var txtForeignKeys = "";
		foreach( var it in ForeignKeys)
		{
			txtForeignKeys += 
					GetTemplateForeignKey( config, it.Key, 
					it.Value.tableRoot, it.Value.columnRoot, it.Value.tableTgt );

			//$@" ""{it.Key}"":[]string{{ {it.Value} }},
			//";
		}

		foreach( var table in tables )
		{
			var listSchemaDefItem_cols = new List<string>();
			foreach( var col in table.columns )
			{
				var col_LangName = col.langName;
				if( col.ForeignKey != null )
				{
					col_LangName = col.langName2; 
				}
				var col_SqlName  = col.sqlName;
				var col_SqlType	 = col.sqlType;
				if( col_SqlName == "role_status_ID")
					Utils.Nop();
				var col_LangType = dialect.getGoLangType( col, ref unusedPackageImports );
					
				var col_IsPrimary	 = col.bIsIdentity?"true":"false";
				var col_IsNullable	 = col.bIsNullable?"true":"false";
				var col_foreignKeyLangName = col.ForeignKey != null? $@"[]string{{ ""{col.langName}"", ""{col.ForeignKey.LangTableNameModel}"" }}":"nil";

				var colTemplate = $@"
				{{
					LangName:   ""{col_LangName}"",
					SqlName:    ""{col_SqlName}"",
					SqlType:    ""{col_SqlType}"",
					LangType:   ""{col_LangType}"",
					Flags:      """",
					IsPrimary:  {col_IsPrimary},
					IsNullable: {col_IsNullable},
					ForeignKeyLangName:{col_foreignKeyLangName},
				}},
				";

				colTemplate = config.UseTemplate( colTemplate, config.Templ_GoLang_SchemaDefItem_Col,
				new Dictionary<string,string>()
				{ 
					{ "col_LangName", col_LangName},
					{ "col_SqlName", col_SqlName},
					{ "col_LangType", col_LangType},
					{ "col_SqlType", col_SqlType},
					{ "col_IsPrimary", col_IsPrimary},
					{ "col_IsNullable", col_IsNullable},
					{ "col_foreignKeyLangName", col_foreignKeyLangName},
				});
				colTemplate = colTemplate.Replace("\r\n", "\n");

				listSchemaDefItem_cols.Add( colTemplate );
			}

			var table_langName = table.LangTableNameModel;
			var table_sqlName = table.SqlTableNameModel;
			var table_schema = table.Schema;
			var table_primaryCol_langName = table.PrimaryColumn.langName;
			var table_primaryCol_sqlName  = table.PrimaryColumn.sqlName;
			var schemaDefItem_cols = String.Join("", listSchemaDefItem_cols.ToArray() );

			var schemaDefItem = $@"
				 ""{table_langName}"": {pkgOrm}.TDefTable{{
				SchemaTable:           ""{table_schema}"",
				SqlTableName:          ""{table_sqlName}"",
				PrimaryColumnLangName: ""{table_primaryCol_langName}"",
				PrimaryColumnSqlName:  ""{table_primaryCol_sqlName}"",
				Columns: []TSqlColumnDef{{
				
					{schemaDefItem_cols}
				}},
			}},
			";
			schemaDefItem = config.UseTemplate( schemaDefItem, config.Templ_GoLang_SchemaDefItem,
				new Dictionary<string,string>()
				{ 
					{ "table_langName", table_langName},
					{ "table_sqlName", table_sqlName},					
					{ "table_schema", table_schema},
					{ "table_primaryCol_langName", table_primaryCol_langName},
					{ "table_primaryCol_sqlName", table_primaryCol_sqlName},
					{ "schemaDefItem_cols", schemaDefItem_cols},
				});
			schemaDefItem = schemaDefItem.Replace("\r\n", "\n");
			listSchemaDefItem.Add(schemaDefItem);
		}

		var schemaDef = String.Join("", listSchemaDefItem.ToArray());
		var sqlName   = dialect.getSql();

		
        var text = $@"

		/* this class is generated automatically by DB_Tool.exe exporter*/
		package {package}

		import (
			{pkgModels} ""{import_package_models}""
			{pkgOrm}	""{import_package_orm}""
		)

		type Vvalue m.IGeneric_MODEL

		//------------------------------------------------------

		type DBContext struct {{
			{pkgOrm}.DBContextBase

			{tables_def1}
			
		}}

		func (_this *DBContext) Constr() (*DBContext,error) {{

			_this.DBContextBase = dbBase;
			_this.AllTables = make( map[string]*{pkgOrm}.{_DBTable}[m.IGeneric_MODEL])

			var sqlSchemaDef {pkgOrm}.TSchemaDef = {pkgOrm}.TSchemaDef{{
				{schemaDef}
			}}
			{tables_DefVar}

			_this.FOREIGN_KEYS = map[string]{pkgOrm}.TForeignKey{{
				{txtForeignKeys}
			}} 
			
			_, err := _this.DBContextBase.Constr(dbBase.Dialect, sqlSchemaDef)
			return _this, err
		}}


		func New_DBContext(dbBase DBContextBase) (*DBContext, error) {{

			ret, err := (new(DBContext)).Constr()

			{init_tables}
			
			return ret, err
		}}
		func (_this *DBContext) Close(){{
			
			_this.DBContextBase.Db.Close()
		}}
		func (_this *DBContext) GetSqlName(){{
			return ""{sqlName}"";
		}}

        ";

		text = config.UseTemplate( text, config.Templ_GoLangOrmFile,
				new Dictionary<string,string>()
				{ 
					{ "package", package},
					{ "pkgModels", pkgModels},
					{ "pkgOrm", pkgOrm},
					{ "import_package_orm", import_package_orm},
					{ "import_package_models", import_package_models},
					{ "tables_def1", tables_def1},
					{ "tables_DefVar", tables_DefVar},
					{ "init_tables", init_tables},
					{ "init_tables_TemplateV", init_tables_TemplateV},
					//{ "Template_V", Template_V},
					//{ "Template_Def", Template_Def},
					{ "schemaDef", schemaDef},	
					{ "txtForeignKeys", txtForeignKeys},
					{ "sqlName", sqlName},
				}
			);
		return text;
        }
    }
}
