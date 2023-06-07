package atomicsql

import (
	//"reflect"

	//"unsafe"

	//sql "database/sql"
	"fmt"
	"reflect"
	"time"

	//"time"

	//reflect
	Str "strings"
	//"reflect"
	//"sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/common/arrays"
	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
	//fmt "fmt"
)

/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _quote( /*#Object*/ data any, columnTable *TSqlColumnDef) string {
	
	ctx := _this.tableInst.m_ctx

	if data == nil {
		return ctx.LangDB.VALUE_NULL
	}
	
	switch data.(type) {
		case time.Time: return data.(time.Time).Format(time.DateTime)
		
		case int: return fmt.Sprintf("%d", data.(int))
		case int32: return fmt.Sprintf("%d", data.(int32))
		case int64: return fmt.Sprintf("%d", data.(int64))
		case int16: return fmt.Sprintf("%d", data.(int16))
		case byte: return fmt.Sprintf("%d", data.(byte))
		case float32: return fmt.Sprintf("%f", data.(float32))
		case float64: return fmt.Sprintf("%f", data.(float64))

		case bool:  return IFF( data.(bool), ctx.LangDB.VALUE_TRUE, 
											 ctx.LangDB.VALUE_FALSE)
		case string: 
		{
			data = Str.ReplaceAll( data.(string), "'", "''")
			return fmt.Sprintf("'%s'", data)
		}
		case []uint8:
		{	
			if( columnTable != nil){
				var slice = data.([]uint8)
				var strData = string( slice );
				if( strData[0] == '{' && strData[len(strData)-1] == '}'){

					strData = Str_SubString(strData, 1, len(strData)-2 );
				}

				var typeSqlElement = Str.ReplaceAll( columnTable.SqlType,"[]", "" )
					
				var valSql = fmt.Sprintf( "ARRAY[%s]::%s[]", strData, typeSqlElement)
				return valSql;
			}
			return "ARRAY[]::text[]"				
		}								
		default:
		{
			dataType:= reflect.ValueOf(data);
			ctx.Log_Print( fmt.Sprintf("type %s is not processed", dataType ));
			return ""
		}
	}

	//return "'data'";

	//ret := mysqli_real_escape_string(_this.g_DB.connection, S(data))
	//ret = mysqli_real_escape_string( data);
	
}

/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _quote1( /*#String*/ data string) string {
	return data
}

/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) CLEAN__quoteTable( /*#String*/ tableName string) string {

	if _this.tableInst.m_ctx.Dialect == ESqlDialect.Postgress {
		tableName = str_replace(`"`, "", tableName)
	} else if _this.tableInst.m_ctx.Dialect == ESqlDialect.MsSql {

		tableName = str_replace("[", "", tableName)
		tableName = str_replace("]", "", tableName)
	} else {
		//if( _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql){

		tableName = str_replace("`", "", tableName)
	}

	return tableName
}

/**
* remove all space in order to avoid SQL injection
* @param type data
* @return type
 */
/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _quoteTable( /*#String*/ tableName string) string {

	var ctx = _this.tableInst.m_ctx;

	if _this.tableInst.m_ctx.Dialect == ESqlDialect.Postgress {
		
		var table, has = ctx.SCHEMA_SQL_BySqlName[tableName];
		if( has ){

			tableName = str_replace(`"`, "", tableName)
			return fmt.Sprintf(`"%s"%s"%s"`, table.SchemaTable, CONCAT_FIELDS, tableName)
		}else
		{
			tableName = str_replace(`"`, "", tableName)
			return fmt.Sprintf(`"%s"`, tableName)
		}

	} else if _this.tableInst.m_ctx.Dialect == ESqlDialect.MsSql {

		var table, has = ctx.SCHEMA_SQL_BySqlName[tableName];
		if( has ){

			tableName = str_replace(`"`, "", tableName)
			return fmt.Sprintf(`[%s]%s[%s]`, table.SchemaTable, CONCAT_FIELDS, tableName)
		}else{
		
			tableName = str_replace("[", "", tableName)
			tableName = str_replace("]", "", tableName)
			return fmt.Sprintf("[%s]", tableName)
		}
	} else {
		//if( _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql)

		//data = str_replace(" ", "_", data)

		var table, has = ctx.SCHEMA_SQL_BySqlName[tableName];
		if( has ){

			tableName = str_replace("`", "", tableName)
			return fmt.Sprintf("`%s`%s`%s`", table.SchemaTable, CONCAT_FIELDS, tableName)
		}else{
			tableName = str_replace("`", "", tableName)
			return fmt.Sprintf("`%s`", tableName)
		}
	}
}

/**
* remove all space in order to avoid SQL injection
* @param type data
* @return type
 */
/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _quoteField( /*#String*/ data string) string {
	if data == "" {
		return data
	}
	if _this.tableInst.m_ctx.Dialect == ESqlDialect.Postgress {
		data = str_replace(`"`, "", data)
		return fmt.Sprintf(`"%s"`, data)

	} else if _this.tableInst.m_ctx.Dialect == ESqlDialect.MsSql {
		data = str_replace("[", "", data)
		data = str_replace("]", "", data)
		return fmt.Sprintf("[%s]", data)

	} else {
		//if( _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql)
		//data = str_replace(" ", "_", data)
		data = str_replace("`", "", data)
		return fmt.Sprintf("`%s`", data)
	}
}

/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _quoteTableField(
	/*#String*/ table_field string,
	/*#BOOL*/ bAlias bool,
	/*#DBSqlJoinCollection*/ joinsCollection *DBSqlJoinCollection) string {

	return _this._quoteTableField1(table_field, bAlias, joinsCollection, false)
}

func (_this *DBQuery[T]) _quoteTableField1(
	/*#String*/ table_field string,
	/*#BOOL*/ bAlias bool,
	joinsCollection *DBSqlJoinCollection,
	/*#BOOL*/ bAddFieldView bool) string {

	if table_field == "" {
		return table_field
	}

	var items = Str.Split(table_field, ".") // table_field

	var itemFK = _this.get_SQL_ITEM_DEF()
	var FK = ""
	var ret = ""
	var count = 0

	tableDBName := _this.tableName

	for iItem := 0; iItem < len(items); iItem++ {

		item := items[iItem]
		ret = _this._quoteField(item)

		if count > 0 {

			var itemsFK = Arr_Slice(&items, 0, count)
			var FK_id = Str.Join(itemsFK, CONCAT_FIELDS)
			FK    = fmt.Sprintf("%s.%s", tableDBName, items[count-1])

			var item1,err = joinsCollection.addJoin( _this.tableInst.m_ctx, _this,
						// _this.pivotProvider?_this.pivotProvider.tableNameOrig:
						item, FK_id, FK, itemFK, &tableDBName)
			if( err != nil){
				_this.errorRet = err;
				return "";
			}
			ret = fmt.Sprintf("%s.%s", item1, ret)
			//itemFK = item;//aici era o eroare
			itemFK = item1
		} else {
			if joinsCollection != nil && joinsCollection._joins != nil {
				ret = fmt.Sprintf("%s.%s", _this.get_SQL_ITEM_DEF(), ret)
			} else {
				//ret = SQL_ITEM_DEF.".ret";
			}
		}
		if count > 0 {
			FK += "."
		}
		FK += item

		count++
	}
	if len(items) > 0 && bAlias {

		if bAddFieldView && len(items) == 1 {
			//in apelul de mai sus: sqlQuery = "SELECT ITEM.*, sqlQueryFields FROM table ITEM";, acest field este inexistent. este un syntentized field in view
			return fmt.Sprintf("'' AS %s", _this._quoteField(table_field))
		} else {
			return fmt.Sprintf("%s AS %s", ret, _this._quoteField(table_field))
		}
	} else {
		return ret
	}
}
