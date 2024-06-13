using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

/**
 type MyNullString struct {
    sql.NullString
}

func (s MyNullString) MarshalJSON() ([]byte, error) {
    if s.Valid {
        return json.Marshal(s.String)
    }
    return []byte(`null`), nil
}

type Company struct {
    ID   int          `json:"id"`              
    Abn  MyNullString `json:"abn,string"`
}
  */
namespace src_tool
{
public class GoModelTemplate
{
    public const string FIELD_IS_OMITTED_INTEGRAL = "-";//sunt in general foreignkeyuri.
    public const string MARK_PRIMARY_KEY = "omitempty";  //this mark the column as primary key

    public const string NoSqlID = "NoSqlID";//ID for NoSql fields.
    public const string POSTFIX_FOREIGNKEY = "_ID";
    
    public const string pkgOrm = "orm";
    public const string PREF_DEF = "T_";
    public const string PREF_VAR = "_";
    public static string GetModel(
        ConfigFile config,
        string packageName, string baseMode, 
        string nameModel, string tableSqlName, string primaryKey, string primaryKeyType,
        string[] packageImport,
        string[] fieldsModel,
        string[] Model_Def,
        string[] ModelInitialize_Def
        //string[] Model_DefVar
        )
    {
        if(nameModel == "Bar")
            Utils.Nop();
        var import_package_orm = config.ImportPackageOrm;
        var pkgImport = "";
        if( packageImport.Length > 0 )
        {
            var import1 = String.Join( 
        @"
            ",packageImport );

            pkgImport = 
        $@"import (
	    {import1}
        )";
        }
        var txtFieldsModel = String.Join( 
        @"
            ",fieldsModel );


        var txtDefModel = String.Join( 
        @"
            ",Model_Def );


        var txtInitializeModel_FK = String.Join(
       @"
            ",ModelInitialize_Def );

        //var txtDefVarModel = "";
        //var txtDefVarModel = String.Join( 
        //@",
        //    ",Model_DefVar );

            var ret = $@"

    /* this class is generated automatically by DB_Tool.exe exporter*/

    package {packageName}
    import (
        ""reflect""
        orm	  ""{import_package_orm}""
    )
        
        {pkgImport}

    type {nameModel} struct /*atomicsql-table:""{tableSqlName}""*/ {{
	        {pkgOrm}.{baseMode}
	        {txtFieldsModel}
    }}

    func (model  {nameModel}) GetID() int32 {{
	    return model.{primaryKey}
    }}
    func (model  {nameModel}) SetID( id int64 ) {{
	    model.{primaryKey} = {primaryKeyType}(id)
    }}

    type {PREF_DEF}{nameModel} struct {{
	        
            {pkgOrm}.Generic_Def
	        {txtDefModel}
    }}

    func (_this *{PREF_DEF}{nameModel}) Def() *{pkgOrm}.TDefIncludeRelation{{
            return &{pkgOrm}.TDefIncludeRelation{{ 
            
                ValueDef: reflect.ValueOf( *_this),
                SqlTable:	""{tableSqlName}"",
                //FnNewInst:	func()any{{ return new ({nameModel}) }},
                FnNewInst:	func(bFull bool)any{{ 
				    var model = new ({nameModel});
				    if( bFull ){{
					    {txtInitializeModel_FK}
				    }}
				    return model;
			    }},
            }}
        }}


            ";

            ret = config.UseTemplate( ret, config.Templ_GoLangModelFile, 
                		new Dictionary<string,string>()
				        { 
					        { "packageName", packageName},
                            { "import_package_orm", import_package_orm},                            
					        { "pkgImport", pkgImport},
					        { "nameModel", nameModel},
                            { "tableSqlName", tableSqlName},                            
					        { "baseMode", baseMode},
					        { "txtFieldsModel", txtFieldsModel},
					        { "PREF_DEF", PREF_DEF},
                            { "primaryKey", primaryKey},
                            { "primaryKeyType", primaryKeyType},
                            { "txtDefModel", txtDefModel},
                            { "txtInitializeModel_FK", txtInitializeModel_FK},
				        }
                );
            return ret;
    }


    public static string ConvertToIdent_GoLang( string ident, bool bRemoveUnderscore=true)
    {
        //var ident1 = ident.Substring(0,1).ToUpper() + ident.Substring(1);
        var ident1 = "";
        var bNextUpper = false;
        var bIsUpper = true;

        if( ident == "user_psw")
            Utils.Nop();

        for( var i = 0; i < ident.Length; i ++)
        {
            var ch  = ident[i].ToString();
            var ch1 = i+1 < ident.Length ?ident[i+1].ToString(): "";

            if( ch == "_" && bRemoveUnderscore 
               && ( (ident.Substring(i) == POSTFIX_FOREIGNKEY)
                 || (ch1 != "_" && ch1.ToUpper() != ch1 && ch1 != "") ) 
              )
            {
                ch = "";
                bNextUpper = true;
            }else
            if( ch == " " )
            {
                ch = "_";
            }

            if( bIsUpper )
                ident1 += ch.ToUpper();
            else
                ident1 += ch;
            bIsUpper = bNextUpper; 

            bNextUpper = false;
        }

        if( ident1 == "User_psw" )
            Utils.Nop();
        return ident1;
    }
    public static string GetModelColumn(string nameCol,string sqlName, string nameType, bool bIdentity, bool bIsForeignKey )
    {
        var jsonName = sqlName;//nameCol.ToLower();
        if(jsonName != "" && bIdentity )
        {
            jsonName = $"{nameCol},{GoModelTemplate.MARK_PRIMARY_KEY}"; //all the time the primary key will be with ID
        }else
        if( bIsForeignKey)
        {
            jsonName = FIELD_IS_OMITTED_INTEGRAL;//omit integral
        }
        var spaces  = new String(' ', Math.Max( 1, 20- nameCol.Length) );
        var spaces2 = new String(' ', Math.Max( 1, 50- ($"{nameCol}{spaces}{nameType}").Length) );
        return $"{nameCol}{spaces}{nameType}{spaces2}`json:\"{jsonName}\"`";
    }

    public static string GetModelColumn_Def(string nameCol, string nameType, bool bIdentity )
    {
        var spaces  = new String(' ', Math.Max( 1, 20- nameCol.Length) );
        return $"{nameCol}{spaces}{nameType}";
    }
    public static string GetModelColumn_DefVar(string nameCol, string nameType, bool bIdentity )
    {
        var spaces  = new String(' ', Math.Max( 1, 20- nameCol.Length) );
        return $"{nameCol}:{spaces}{nameType}";
    }

    public static string getGoLangType_Def( DbColumn column, bool bReference)
    {
        if( bReference )
        {
            return PREF_DEF + column.ForeignKey.LangTableNameModel;
        }
        return "string";
    }

    
    }
}


/**
 * 

select c.attname, pg_catalog.format_type(c.atttypid, NULL) as data_type
from pg_attribute c
  join pg_class t on c.attrelid = t.oid
  join pg_namespace n on t.relnamespace = n.oid
where t.relname = 'table_name'  -- your table name
  and n.nspname = 'public' -- table schema
  and c.attnum >= 0 -- exclude internal columns

 */ 
