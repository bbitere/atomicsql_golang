package atomicsql

import (
	"reflect"
	time "time"
)

//type mUser models.User
//type Vvalue models.IGeneric_MODEL

type DBTable[T IGeneric_MODEL] struct {
	//rows      []*T
	m_sqlName     string
	m_langName    string
	m_schemaTable string
	m_ctx         *DBContextBase
	//m_ctxGeneric  any /*  *DBContext */
}

// Constructor of DBTable
func (_this *DBTable[T]) Constr(sqlName string, langName string, ctx *DBContextBase) *DBTable[T] {

	_this.m_sqlName = sqlName
	_this.m_langName = langName
	tableDB, has := ctx.SCHEMA_SQL[langName]
	if has {
		_this.m_schemaTable = tableDB.SchemaTable
	}
	_this.m_ctx = ctx

	var model = _this.CloneGenericModel()
	ctx.AllTables[sqlName] = model
	return _this
}

func (_this *DBTable[T]) cloneTable_GenModel() *DBTable[IGeneric_MODEL] {

	var table = (new(DBTable[IGeneric_MODEL])).Constr(_this.m_sqlName, _this.m_langName, _this.m_ctx)
	table.m_schemaTable = _this.m_schemaTable

	return table
}

// Detache a model from the ORM. From this point you can insert it again in the table.
func (_this *DBTable[T]) DeatachModel(model *T) {

	var table, has = _this.m_ctx.SCHEMA_SQL[_this.m_langName]
	if has {

		var fld = reflect.ValueOf(model).Elem().FieldByName(table.PrimaryColumnLangName)
		fld.SetInt(0)
	}
	//model.SetID( 0 )
}

// Do a clone of a DBTable
func (_this *DBTable[T]) CloneGenericModel() *DBTable[IGeneric_MODEL] {

	var newInst = new(DBTable[IGeneric_MODEL])

	newInst.m_sqlName = _this.m_sqlName
	newInst.m_langName = _this.m_langName
	newInst.m_schemaTable = _this.m_schemaTable
	newInst.m_ctx = _this.m_ctx
	//newInst.m_ctxGeneric = _this.m_ctxGeneric

	
	return newInst
}

// Qry( tag ) - each DBTable can start a query sequence (DBQuery) usign Qry() method,
//
// The query sequence has implemented all methods to generate the sql query and it should end in one of :
//
// GetModels(), GetFirstModel(),
// GetModelsRel(), GetFirstModelRel(),
// GetRecords(), GetFirstRecord(),
// GetCount(), GetDistinctCount(),
// GetSingleDataString(), GetSingleDataInt(),
// GetRowsAsFieldString(), GetRowsAsFieldInt()
// GetDistinctRecords()
// DeleteModels(), DeleteModel()
//
// the `tag` argument should be non empty and unique per app only in this 3 examples
//
//	context.Table.Qry("tag1").Where( func(x *Table) bool{ ... }).
//
// or
//
//	atmsql.Select( context.Table.Qry("tag2"), func(x *Table) *TView{ ... }).
//
// or
//
//	atmsql.Select( orm.Aggregate[ Table, TableAggr ]( context.Table.Qry("tag2")), func(x *TableAggr) *TView{ ... }).
//
// Each of this statemets will be translated by the scanner tool (using script 4.scan_queries.cmd),
// and the inner content will be translated in a SQL Query.
// So this tag, help to do a correct identification of the precompiled sql query.
//
// Usign this Qry(tag) method to generate sql query, is the main diference between atomicSql library and linq (C#) or jinq (java)
func (_this *DBTable[T]) Qry(tagID string) *DBQuery[T] {
	return _this.QryS(tagID, nil)
}

// QryS( tag, query ) - identical as Qry(tag), but is using only in lit func of WhereSubQ() and SelectSubQ()
//
func (_this *DBTable[T]) QryS(tagID string, containerQuery IDBQuery) *DBQuery[T] {

	
	_this.m_ctx.currOperationDTime = time.Now()
	_this.m_ctx.resetSubTag()

	//var this2 any = any(_this)
	//var this1 = _this.(*DBTable[models.IGeneric_MODEL, models.IGeneric_MODEL])
	query := (new(DBQuery[T])).Constr(_this)

	if containerQuery != nil {		
		query.myTag = containerQuery.GetTagID() + "."+ tagID;
		query.currentSubQueryID    = tagID;
		query.parentContainerQuery = containerQuery;
	}else{
		query.myTag = tagID
	}
	

	//query.rows = _this.getRows()
	return query
}

func (_this *DBQuery[T]) finishSubQuery(sqlQuery string) bool {

	if( _this.parentContainerQuery != nil){
		
		_this.clearCachedSyntax();
		_this.parentContainerQuery.SetSubQueryString( _this.currentSubQueryID, sqlQuery);
		return true;
	}
	return false;
}

func (_this *DBTable[T]) getSchemaTable() string {
	return _this.m_schemaTable
}

func (_this *DBTable[T]) getSqlName() string {
	return _this.m_sqlName
}
func (_this *DBTable[T]) getLangName() string {
	return _this.m_langName
}

/*#PHPARG=[ T ];*/
func (_this *DBTable[T]) getModel2( /*#HashMap<Object>*/ args map[string]any) *T {
	return new(T) //new_CLASS1( $this->m_PhpMODEL, $args);
}

/*
func Tag2[T models.IGeneric_MODEL, V Vvalue](tagID string, table IDBTable[T]) IDBQuery[T, V] {

	query := new_DBQuery2[T, V](table)
	return query
}*/
