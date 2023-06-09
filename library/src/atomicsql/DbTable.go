package atomicsql

import (
	"reflect"
	time "time"
)

//type mUser models.User

//type Vvalue models.IGeneric_MODEL



type DBTable[T IGeneric_MODEL] struct {
	//rows      []*T
	m_sqlName  		string
	m_langName 		string
	m_schemaTable 	string 
	m_ctx      		*DBContextBase
}

func (_this *DBTable[T]) Constr(sqlName string, langName string, ctx *DBContextBase) *DBTable[T] {

	_this.m_sqlName = sqlName
	_this.m_langName = langName
	tableDB, has := ctx.SCHEMA_SQL[ langName ]
	if( has ){
		_this.m_schemaTable = tableDB.SchemaTable;
	}
	_this.m_ctx = ctx

	var model = _this.CloneGenericModel();
	ctx.AllTables[ sqlName ] = model
	return _this
}

func (_this *DBTable[T]) cloneTable_GenModel() *DBTable[IGeneric_MODEL] {

	var table = (new (DBTable[IGeneric_MODEL])).Constr(_this.m_sqlName, _this.m_langName, _this.m_ctx);
	table.m_schemaTable = _this.m_schemaTable;

	return table;
}

func (_this *DBTable[T]) DeatachModel( model *T ){

	var table, has = _this.m_ctx.SCHEMA_SQL[ _this.m_langName ]
	if( has ){

		var fld = reflect.ValueOf( model ).Elem().FieldByName( table.PrimaryColumnLangName )
		fld.SetInt( 0 );		
	}
	//model.SetID( 0 )
}

func (_this *DBTable[T]) CloneGenericModel() *DBTable[IGeneric_MODEL] {

	var newInst = new( DBTable[ IGeneric_MODEL])

	newInst.m_sqlName = _this.m_sqlName
	newInst.m_langName = _this.m_langName
	newInst.m_schemaTable = _this.m_schemaTable
	newInst.m_ctx = _this.m_ctx	
	return newInst
}

//func (_this *DBTable[T]) getRows() []*T { return _this.rows }

func (_this *DBTable[T]) Qry(tagID string) *DBQuery[T] {

	_this.m_ctx.currOperationDTime = time.Now()
	_this.m_ctx.resetSubTag();
	
	//var this2 any = any(_this)
	//var this1 = _this.(*DBTable[models.IGeneric_MODEL, models.IGeneric_MODEL])
	query := (new(DBQuery[T])).Constr(_this)
	query.myTag = tagID;

	//query.rows = _this.getRows()
	return query
}


func (_this *DBTable[T]) getSchemaTable() string {
	return _this.m_schemaTable;
}

func (_this *DBTable[T]) getSqlName() string {
	return _this.m_sqlName
}
func (_this *DBTable[T]) getLangName() string {
	return _this.m_langName
}

/*#PHPARG=[ T ];*/
func (_this *DBTable[T]) getModel( /*#HashMap<Object>*/ args map[string]any) *T {
	return new(T) //new_CLASS1( $this->m_PhpMODEL, $args);
}

/*
func Tag2[T models.IGeneric_MODEL, V Vvalue](tagID string, table IDBTable[T]) IDBQuery[T, V] {

	query := new_DBQuery2[T, V](table)
	return query
}*/
