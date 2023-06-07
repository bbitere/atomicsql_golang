package atomicsql

import (
	//"reflect"

	//"unsafe"

	sql "database/sql"
	fmt "fmt"
	reflect "reflect"
	time "time"
	//"sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/common/arrays"
	//"sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/common/arrays"
	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
	//fmt "fmt"
)

//type mUser m.User

//type Vvalue m.IGeneric_MODEL

// ------------------------------------------------------

const NULL_FK_ID = 0


type IQueryBase interface {

	//return 'val' or val or null
	_quote(val any, colTable *TSqlColumnDef) string;

	//return `field`
	_quoteField(field string) string;

	//return `schema`.`table``
	_quoteTable(field string) string;	
}
type Vvalue IGeneric_MODEL


type DBSqlJoin struct {
	nameItem string
	sql      string //left join Table item on item.ID = item2.FK

}

/*#PHPARG=[ String ];*/
func (_this *DBSqlJoin) getSqlTxt( itm string ) string {
	return _this.sql
}

type DBSqlJoinCollection struct {

	/*#PHPARG=[ Array< DBSqlJoin >];*/
	_joins map[string]*DBSqlJoin
}

func (_this *DBSqlJoinCollection) Constr() *DBSqlJoinCollection{
	_this._joins = nil;//make(map[string]*DBSqlJoin)
	return _this;
}

/*#PHPARG=[ String ];*/
/**
 *
 * @global type DB_FOREIGN_KEYS
 * @param type pivotTableName
 * @param type itemTxt
 * @param type FK_id: _user._mrole._role  (signature)
 * @param type FK    mrole._role  (table + fk)
 * @param type itemFK
 * @return string
 */

type DBSqlQuery[T IGeneric_MODEL] struct {
	text string

	m_op       string
	m_field1   string
	m_field2   string
	m_operand1 string
	m_operand2 any
	/*#PHPARG=[ Delegate1< BOOL, T> ];*/
	fnWhere func(x *T) bool

	/*#PHPARG=[ Array<DBSqlQuery> ];*/
	m_listOperands []*DBSqlQuery[T]

	/*#PHPARG=[ Array<String> ];*/
	m_listOperandsStr []any
}

func (_this *DBSqlQuery[T]) Constr( /*#String*/ text string) *DBSqlQuery[T] {

	_this.text = text
	return _this
}

func (_this *DBSqlQuery[T]) cloneSqlQuery_GenModel()*DBSqlQuery[IGeneric_MODEL] {

	var newQ = (new (DBSqlQuery[IGeneric_MODEL])).Constr(_this.text);
	newQ.text = _this.text;

	newQ.m_op = _this.m_op;
	newQ.m_field1  = _this.m_field1;
	newQ.m_field2  = _this.m_field2 ;
	newQ.m_operand1 = _this.m_operand1;
	newQ.m_operand2 = _this.m_operand2;
	newQ.m_listOperandsStr = _this.m_listOperandsStr;

	//fnWhere 
	//m_listOperands 
	return newQ;
}

/*#PHPARG=[ String ];*/
func (_this *DBSqlQuery[T]) getText() string {
	return _this.text
}

type  IDBQuery interface{

	_generateSelectSql(selectFields string, ITEM string, bLimit bool, select_sqlFields []string) string;
}

type DBQuery[T IGeneric_MODEL] struct {
	//DBSqlQuery[T]
	//rows []*T

	//tableName string;
	/*#PHPARG=[ DBSqlProvider<T> ];*/
	parentQuery IDBQuery;//*DBQuery[T]

	/*#PHPARG=[ DBSqlQuery<T> ];*/
	//querySelectNewRecord any;//*DBSqlQuery[T]
	querySelectNewRecord_Text string;//*DBSqlQuery[T]
	querySelectNewRecord_isAgregator bool

	/*#PHPARG=[ String ];*/
	//lamdaSelectNewRecord string

	//what fields need to be excluded from GroupBy
	excludeLangFieldsFromGroupBy  []string
	//what fields need to be in SELECT in clause from GroupBy
	selectSqlFields				[]string

	/*#PHPARG=[ HashMap<String> ];*/
	withForeignKeys map[string]string

	tableInst *DBTable[T]

	///*#PHPARG=[ String ];*/
	//tablePhpModelName string

	tableName            string //sqlName
	schemaTable          string
	tableNameOrig        string //langName
	tableNameOrig_nonTmp string

	whereTxt       	string
	limit          	string
	orderBy        	string
	last_error     	string //= "";
	m_SQL_ITEM_DEF 	string
	myTag 			string // memorize the tag to be used with Where() and Select()
	subTag			string

	/*#PHPARG=[ DBSqlQuery ];*/
	m_queryAND *DBSqlQuery[T]

	/*#PHPARG=[ DBSqlJoinCollection ];*/
	joins *DBSqlJoinCollection
	errorRet error
}

func (_this *DBQuery[T]) Constr(tableInst *DBTable[T]) *DBQuery[T] {

	//_this.rows = []*T{}

	_this.last_error = ""
	_this.m_SQL_ITEM_DEF = SQL_ITEM_DEF
	_this.tableInst = tableInst
	_this.tableName = tableInst.getSqlName()
	_this.schemaTable = tableInst.getSchemaTable();
	_this.tableNameOrig = tableInst.getLangName()
	//_this.lamdaSelectNewRecord = ""

	_this.joins = (new(DBSqlJoinCollection)).Constr()
	//_this.joins._joins = make(map[string]*DBSqlJoin)

	_this.m_SQL_ITEM_DEF = SQL_ITEM_DEF
	_this.m_queryAND = nil
	_this.whereTxt = ""
	_this.limit = ""
	_this.orderBy = ""
	_this.withForeignKeys = nil

	return _this
}

func (_this *DBQuery[T]) cloneQuery() *DBQuery[T] {

	newQuery := (new(DBQuery[T])).Constr(_this.tableInst) //_this.tableInst.m_DBSqlProvider );

	newQuery.tableNameOrig = _this.tableInst.m_sqlName
	newQuery.tableNameOrig_nonTmp = ""
	newQuery.tableInst = _this.tableInst

	newQuery.tableName = _this.tableName
	newQuery.getJoins()._joins = _this.getJoins()._joins

	newQuery.m_queryAND = _this.m_queryAND
	newQuery.whereTxt = _this.whereTxt
	newQuery.limit = _this.limit
	newQuery.orderBy = _this.orderBy
	newQuery.withForeignKeys = _this.withForeignKeys

	newQuery.parentQuery = _this.parentQuery //.cloneQuery();
	newQuery.querySelectNewRecord_Text = _this.querySelectNewRecord_Text
	newQuery.querySelectNewRecord_isAgregator = _this.querySelectNewRecord_isAgregator
	
	//newQuery.lamdaSelectNewRecord = _this.lamdaSelectNewRecord;

	return newQuery
}

func (_this *DBQuery[T]) cloneQuery_GenModel() *DBQuery[IGeneric_MODEL] {

	var newTable = _this.tableInst.cloneTable_GenModel();
	var newQuery = (new(DBQuery[IGeneric_MODEL])).Constr(newTable) //_this.tableInst.m_DBSqlProvider );

	newQuery.tableNameOrig = _this.tableInst.m_sqlName
	newQuery.tableNameOrig_nonTmp = ""
	newQuery.tableInst = newTable

	newQuery.tableName = _this.tableName
	newQuery.getJoins()._joins = _this.getJoins()._joins

	if( _this.m_queryAND != nil){
		newQuery.m_queryAND = _this.m_queryAND.cloneSqlQuery_GenModel()
	}
	newQuery.whereTxt = _this.whereTxt
	newQuery.limit = _this.limit
	newQuery.orderBy = _this.orderBy
	newQuery.withForeignKeys = _this.withForeignKeys

	newQuery.parentQuery = _this.parentQuery //.cloneQuery();
	newQuery.querySelectNewRecord_Text = _this.querySelectNewRecord_Text
	newQuery.querySelectNewRecord_isAgregator = _this.querySelectNewRecord_isAgregator
	//newQuery.lamdaSelectNewRecord = _this.lamdaSelectNewRecord;
	newQuery.excludeLangFieldsFromGroupBy = _this.excludeLangFieldsFromGroupBy;

	return newQuery
}

/*
func (_this *DBQuery[T]) Select[V any](fnSelect func(x *T) *V) *DBQuery[V] {

	return _selectNewRecord[T]( _this, fnSelect );
}
//*/
func Select[T IGeneric_MODEL, V IGeneric_MODEL]( 
	_this *DBQuery[T],  
	fnSelect func(x *T) *V,
	)  *DBQuery[V] {

	_this.subTag = "S"+_this.tableInst.m_ctx.getSubTag();
	return _Select_query( _this, fnSelect );
}

func Aggregate[T IGeneric_MODEL, V IGeneric_MODEL]( 
	_this *DBQuery[T],
	 )  *DBQuery[V] {

	var ctx = _this.tableInst.m_ctx
	//var safe_SQL_ITEM_DEF = _this.m_SQL_ITEM_DEF;

	var tbl1 = (new(DBTable[V])).Constr(
	_this.tableInst.m_sqlName,
	_this.tableInst.m_langName,
	_this.tableInst.m_ctx)

	var query = (new(DBQuery[V])).Constr(tbl1);

	query.myTag = _this.myTag;
	//_this.excludeLangFieldsFromGroupBy	= excludeFromGroupBy
	query.parentQuery = _this;//.cloneQuery_GenModel();
	//query.querySelectNewRecord = (new (DBSqlQuery[V])).Constr(sql);
	
	query.querySelectNewRecord_isAgregator = true;
	query.m_SQL_ITEM_DEF 		= ctx.newSQL_ITEM( SQL_ITEM_DEF_Aggr );
	//query.lamdaSelectNewRecord 	= _this.m_SQL_ITEM_DEF;

	var sql, excludeFromGroupBy = _Aggregate_generateSql[ T, V]( _this, _this.m_SQL_ITEM_DEF );
	query.querySelectNewRecord_Text = sql;
	query.excludeLangFieldsFromGroupBy	= excludeFromGroupBy		
	//query.tablePhpModelName    = tablePhpModelName;
	
	//_this.m_SQL_ITEM_DEF = safe_SQL_ITEM_DEF;
	return query
}


func (_this *DBQuery[T]) WhereEq(field string, operand any) *DBQuery[T] {

	return _this._whereEq(field, operand, "", false)
}
func (_this *DBQuery[T]) WhereNotEq(field string, operand any) *DBQuery[T] {

	return _this._whereEq(field, operand, "", true)
}
func (_this *DBQuery[T]) WhereEqFld(field string, field2 string) *DBQuery[T] {

	return _this._whereEq(field, nil, field2, false)
}
func (_this *DBQuery[T]) WhereNotEqFld(field string, field2 string) *DBQuery[T] {

	return _this._whereEq(field, nil, field2, true)
}

func (_this *DBQuery[T]) WhereIn( field string,  operandsIn []any) *DBQuery[T]{        

	var query = _this.InFO( field, operandsIn, false, false, "");
	if( _this.whereTxt != "" ){        
		_this.whereTxt += " AND ";
	}
	_this.whereTxt += query.getText();        
	return _this;
}

/*#PHPARG=[ DBSqlProvider< T > ];*/
/**  @return DBSqlProvider  */
func (_this *DBQuery[T]) whereNotIn( field string, operandsIn []any)*DBQuery[T]{        

	var query = _this.InFO( field, operandsIn, true, false, "");
	if( _this.whereTxt != "" ){        
		_this.whereTxt += " AND ";
	}
	_this.whereTxt += query.getText();        
	return _this;
}



func (_this *DBQuery[T]) Where( fnWhere func(x *T) bool) *DBQuery[T] {

	_this.subTag = "W"+_this.tableInst.m_ctx.getSubTag();
	var querySql = _this._whereGeneric( fnWhere );//"($opText1) AND ($opText2)" );

	if( _this.whereTxt != "" ) {
		_this.whereTxt += " AND ";
	}

	if( querySql == nil ){
		_this.whereTxt += "( 1 = 0 )";
	}else{
		_this.whereTxt += fmt.Sprintf( "(%s)", querySql.getText()  );
	}

	return _this
}




func (_this *DBQuery[T]) GetModels() ([]*T, error) {

	return _this.GetRows(nil)
}

func (_this *DBQuery[T]) GetRows(fields []string) ([]*T, error) {

	sqlQuery := _this._getRows(false, fields, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if err == nil {
		_this.clearCachedSyntax()
		var ret1, err1 = _this._arrayRecords(dbResult)
		_this.tableInst.m_ctx.updateDeltaTime();
		return ret1, err1;
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}


func (_this *DBQuery[T]) GetDistinctModel() ([]*T, error) {

	sqlQuery := _this._getRows(true, nil, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if err == nil {
		_this.clearCachedSyntax()
		return _this._arrayRecords(dbResult)
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}

func (_this *DBQuery[T]) GetDistinctRows(fields []string) ([]*T, error) {

	sqlQuery := _this._getRows(true, fields, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if err == nil {
		_this.clearCachedSyntax()
		return _this._arrayRecords(dbResult)
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}

func (_this *DBQuery[T]) GetFirstModelRel( structDefs ... *TDefIncludeRelation ) (*T, error) {

	_this.setLimit(0, 1)
	arrAny, err := _this._getModelRelations(structDefs, nil) 

	var arr = convertToTemplateT[T](arrAny);
	if( len(arr) > 0 ){
		return arr[0], nil
	}
	
	return nil, err
}

func (_this *DBQuery[T]) GetModelsRel( structDefs ... *TDefIncludeRelation ) ([]*T, error) {
	
	arrAny, err := _this._getModelRelations(structDefs, nil) 

	var arr = convertToTemplateT[T](arrAny);
	
	return arr, err
}


func (_this *DBQuery[T]) GetFirstModel() (*T, error) {

	_this.setLimit(0, 1)

	sqlQuery := _this._getRows(false, nil, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if dbResult != nil && err == nil {
		_this.clearCachedSyntax()
		return _this._oneRecord(dbResult)
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}

func (_this *DBQuery[T]) GetFirstRow(fields []string) (*T, error) {

	_this.setLimit(0, 1)

	sqlQuery := _this._getRows(false, fields, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if dbResult != nil && err == nil {
		_this.clearCachedSyntax()
		return _this._oneRecord(dbResult)
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}

func (_this *DBQuery[T]) SingleDataS(sqlResult *sql.Rows, fieldName string) (string, error) {

	_this.clearCachedSyntax()
	model, err := _this._oneRecord(sqlResult)

	if model != nil && err == nil{
		val := reflect.ValueOf(model).Elem().FieldByName(fieldName)
		return val.String(), nil
	}
	return "", err
}

func (_this *DBQuery[T]) SingleDataInt(sqlResult *sql.Rows, fieldName string) (int64, error) {

	_this.clearCachedSyntax()
	model, err := _this._oneRecord(sqlResult)

	if model != nil && err == nil{
		val := reflect.ValueOf(model).Elem().FieldByName(fieldName)
		return val.Int(), nil
	}
	return 0, err
}

func (_this *DBQuery[T]) GetSingleFieldRows(field string) []string {

	sqlQuery := _this._getRows(false, []string{field}, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if dbResult != nil && err == nil {

		_this.clearCachedSyntax()
		return _this._arrayOfSingleField(dbResult, field)
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil
}

func (_this *DBQuery[T]) OrderByFields(orderFields *DataOrderByFields) *DBQuery[T] {

	if orderFields == nil || orderFields.data == nil {
		return _this
	}

	orderBy := ""

	fields := orderFields.data
	for nameField, val := range fields {

		if orderBy != "" {
			orderBy += ", "
		}

		if val == "asc" {
			orderBy += fmt.Sprintf(" %s ASC", _this._quoteField(nameField))
		} else if val == "desc" {
			orderBy += fmt.Sprintf(" %s DESC", _this._quoteField(nameField))
		}
	}

	if orderBy != "" {
		_this.orderBy = fmt.Sprintf(" ORDER BY %s ", orderBy)
	} else {
		_this.orderBy = fmt.Sprintf(" ORDER BY %s, %s", DEF_TABLE_ROW_IDX, DEF_TABLE_COLUMN_ID)
	}
	return _this
}

/*#PHPARG=[ DBSqlProvider< T > ];*/
/**  @return DBSqlProvider  */
func (_this *DBQuery[T]) OrderAsc(field string) *DBQuery[T] {

	if _this.orderBy == "" {
		_this.orderBy = fmt.Sprintf("ORDER BY %s ASC", _this._quoteField(field))
	} else {
		_this.orderBy += fmt.Sprintf(", %s ASC", _this._quoteField(field))
	}
	return _this
}

/*#PHPARG=[ DBSqlProvider< T > ];*/
/**  @return DBSqlProvider  */
func (_this *DBQuery[T]) OrderDesc(field string) *DBQuery[T] {

	if _this.orderBy == "" {
		_this.orderBy = fmt.Sprintf("ORDER BY %s DESC", _this._quoteField(field))
	} else {
		_this.orderBy += fmt.Sprintf(", %s DESC", _this._quoteField(field))
	}
	return _this
}

func (_this *DBQuery[T]) InsertModel(data *T) (int64, error) {

	return _this.InsertRecord(data, false, nil)
}
func (_this *DBQuery[T]) InsertOrUpdateModel(data *T) (int64, error) {

	return _this.InsertOrUpdateRecord(data, false, nil)
}

/*#PHPARG=[ Object ];*/
func (_this *DBQuery[T]) InsertRecord(data *T, bInsertID bool, fields *[]string) (int64, error) {

	if data == nil {
		return 0, nil
	}
	reflectData := reflect.ValueOf(data).Elem();
	//fldID := reflectData.FieldByName( "ID")
	//fldID.SetInt( 1 )

	return _this._InsertRecordByReflectValue( _this.tableNameOrig, reflectData, bInsertID, fields );
}

func (_this *DBQuery[T]) InsertOrUpdateRecord(model *T, bInsertID bool, fields *[]string) (int64, error) {

	if model == nil {
		return 0, nil
	}
	

	if( (*model).GetID() == 0 ){

		reflectData := reflect.ValueOf(model).Elem();
		return _this._InsertRecordByReflectValue( _this.tableNameOrig, reflectData, bInsertID, fields );
	}else{

		var arr = []*T{}
		Arr_Append( &arr, model )

		var err = _this._updateBulkRecords( &arr, fields );
		return (*model).GetID(), err
	}
}

func (_this *DBQuery[T]) InsertOrUpdateModels(models []*T) ( error) {

	return _this.InsertOrUpdateRecords( models, nil);
}

func (_this *DBQuery[T]) InsertOrUpdateRecords(models []*T, fields []string) ( error) {

	var arrIns = []*T{}
	var arrUpd = []*T{}

	for _, model := range(models) {

		if( (*model).GetID() == 0 ){

			Arr_Append( &arrIns, model)
		}else{

			Arr_Append( &arrUpd, model)			
		}
	}

	if( len(arrIns) > 0 ){

		for _, model := range(arrIns){
			var _, err = _this.InsertRecord( model, false, &fields);
			if( err != nil){
				return err
			}
		}
	}
	if( len(arrUpd) > 0 ){

		var err = _this._updateBulkRecords( &arrUpd, &fields);
		if( err != nil){
			return err
		}
	}
	return nil

}


func (_this *DBQuery[T]) UpdateModels(  records *[]*T) error {

	return _this._updateBulkRecords(  records, nil);
}
func (_this *DBQuery[T]) UpdateModel( model *T) error {

	var arr = []*T{}
	Arr_Append( &arr, model )

	return _this._updateBulkRecords( &arr, nil);
}


func (_this *DBQuery[T]) DeleteRecords()  error {

	var sqlQuery    = _this._deleteRecords();

	var ctx = _this.tableInst.m_ctx

	ctx.currOperationDTime2 = time.Now()		
	dbResult1, err := _this.tableInst.m_ctx.Db.Exec(sqlQuery)	
	ctx.updateDeltaTime2()

	if( dbResult1 != nil && err == nil ){
		_this.clearCachedSyntax();
		return nil;
	}
	
	_this.checkMySqlError( sqlQuery, err );
	return err;
}


const COUNT_NAME   = "Count1";
type TGetCount struct {
	Generic_MODEL
	Count1 int32 // the same name as COUNT_NAME
}

func (_this *DBQuery[T])  GetCount( ) (int64, error){

	
	var sqlQuery     = _this._getCount( COUNT_NAME );

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult1, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult1 )
	ctx.updateDeltaTime2()	

	if( dbResult1 != nil && err == nil ){
		/*if( $bDonotCall )
		{
			$rec = _this._oneRecord( $mysqliResult );
			return I($rec[ $COUNT_NAME ]);                
		}
		else*/

		var tableCnt = (new ( DBTable[TGetCount])).Constr("", "", _this.tableInst.m_ctx)
		
		var ret, err = tableCnt.Qry("").SingleDataInt( dbResult1, COUNT_NAME );
		if( err == nil ){

			return ret, nil;
		}
	}
	
	//if( $bDonotCall )
	//    _this.checkMySqlError1( $sqlQuery );
	//else
		_this.checkMySqlError( sqlQuery, err );
	return 0, err;
}

/*#PHPARG=[ INT ];*/
func (_this *DBQuery[T])  GetDistinctCount( fields []string) (int64,error){

	var sqlQuery     = _this._getDistinctCount( COUNT_NAME, fields );

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()			
	dbResult1, err := _this.tableInst.m_ctx.Db.Query(sqlQuery)
	defer queryClose( dbResult1 )
	ctx.updateDeltaTime2()	

	if( dbResult1 != nil && err == nil ){
		/*if( $bDonotCall )
		{
			$rec = _this._oneRecord( $mysqliResult );
			return I($rec[ $COUNT_NAME ]);                
		}
		else*/
		var tableCnt = (new ( DBTable[TGetCount])).Constr("", "", _this.tableInst.m_ctx)
		var ret, err = tableCnt.Qry("").SingleDataInt( dbResult1, COUNT_NAME );
		if( err == nil ){

			return ret, nil;
		}
	}
	
	//if( $bDonotCall )
	//    _this.checkMySqlError1( $sqlQuery );
	//else
		_this.checkMySqlError( sqlQuery, err );
	return 0, err;
}
