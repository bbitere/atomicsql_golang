package atomicsql

import (
	//"reflect"

	//"unsafe"

	sql "database/sql"
	"fmt"
	"math"
	Sort "sort"
	"time"
	"unsafe"

	//reflect
	Str "strings"

	log "log"
	"reflect"
	//"sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/common/arrays"
	//m "sourcerer.slotmonitor.ro/nevada/frontend.git/pkg/newton_models"
	//fmt "fmt"
)

const ATTR_ATOMICSQL_COPY_MODEL string = "atomicsql:\"copy-model\"";


const Generic_MODEL_Name = "Generic_MODEL"
const ORM_NAME string = "gatomicsql"
const SQL_ITEM_DEF string = "itm"
const SQL_ITEM_DEF_SQ string = "itmq"
const SQL_ITEM_DEF_Aggr string = "itmAggr"

const CONCAT_FIELDS string = "."
const DEF_TABLE_COLUMN_ID = "ID"
const DEF_TABLE_ROW_IDX = "ROW_IDX"

//{#@field@#}
//{@@staticKey@@}
const PREFIX_FIELDS = "{#@";
const POSTFIX_FIELDS = "@#}";


const SQL_POSTGRESS_RND_UUID = "gen_random_uuid()"

var gFOREIGN_KEYS map[string]([]string) = map[string]([]string){

	"user": []string{"user", "userrole"},
}

func debugger(){

	panic("debugger")
}
func array_push[T any](arr *[]T, e T) {

	Arr_Append(arr, e)
}
func Str_Index(data string, find string, idx int) int {

	return Str.Index(data[idx:], find)
}

func isset(data any) bool {

	if data != nil {
		return true
	}
	return false
}


func S(s any) string {

	return fmt.Sprintf("%s", s)
}
func IFFs(b bool, s1 string, s2 string) string {
	if b {
		return s1
	} else {
		return s2
	}
}
func IFF[T any](b bool, s1 T, s2 T) T {
	if b {
		return s1
	} else {
		return s2
	}
}

func str_replace(replaced string, replace string, target string) string {
	return Str.Replace(target, replaced, replace, -1)
}

func str_replaceN(replaced string, replace string, target string, nTimes int) string {
	return Str.Replace(target, replaced, replace, nTimes)
}

func substr(s string, idx int, len int) string {
	return s[idx : len]
}

/*#PHPARG=[ String ];*/
func _getDotPrefix( /*#String*/ text string) string {
	index := Str.Index(text, ".")
	if index >= 0 {
		return substr(text, 0, index)
	}
	return text
}

func _removeQuote( /*#String*/ data string) string {
	if Str.Index(data, "'") == 0 && Str_Index(data, "'", 1) == len(data)-1 {
		return substr(data, 1, len(data)-2)
	}
	return data
}

//------------------------------------------------------------------------

type DataOrderByFields struct {
	data map[string]string
}

func (_this DataOrderByFields) Constr(data map[string]string) {

	_this.data = data
}

func (_this *DBQuery[T]) _generateSqlSourceOfData() string {

	//if( _this.querySelectNewRecord_Text != "" && _this.querySelectNewRecord_isSubQuery ){
	//	return fmt.Sprintf( "(%s)", _this.querySelectNewRecord_Text);
	//}else
	if  _this.parentQuery != nil {
		//_this.querySelectNewRecord_Text != "" &&
		//_this.lamdaSelectNewRecord != "" {

			var sql = _this.parentQuery._generateSelectSql( 
				//_this.querySelectNewRecord_Text, 
				//_this.lamdaSelectNewRecord, 
				"","",
				true, 
				_this.selectSqlFields)
			return fmt.Sprintf(`( %s )`, sql)
	} else 
	if _this.tableName != "" {
		var tableName  = _this._quoteTable( _this.tableName )
		//schemaName := _this._quoteTable( _this.schemaTable )
		return tableName;
	} else {
		_this.errorRet = fmt.Errorf( "error-no-table: query has no table source" );
		return "/*error-no-table*/"
	}
}

func (_this *DBQuery[T]) _generateSelectSql(
	selectFields string, ITEM string /*#BOOL*/, bLimit bool, select_sqlFields []string) string {

	var table = _this._generateSqlSourceOfData()
	//sqlQuery += fmt.Sprintf(`SELECT %s FROM %s %s`, selectFields, table, ITEM)

	var joinTxt = "";
	if _this.getJoins() != nil {
		joins1 := _this.getJoins()._joins
		for _, element := range joins1 {

			joinTxt += element.getSqlTxt( ITEM )
		}
	}
	//sqlQuery += joinTxt

	var sqlQueryWhere = "";
	if _this.whereTxt != "" {
		sqlQueryWhere = fmt.Sprintf("WHERE %s", _this.whereTxt)
	}

	var sqlQueryGroupBy = "";
	if _this.excludeLangFieldsFromGroupBy != nil && select_sqlFields != nil {
		sqlQueryGroupBy = _this._Aggregate_generateGroupBySql( select_sqlFields, _this.excludeLangFieldsFromGroupBy ) 
	}
	var sqlQueryOrderBy = "";
	if _this.orderBy != "" {
		sqlQueryOrderBy = _this.orderBy
	}

	var sqlQueryLimit = "";
	if bLimit && _this.limit != "" {
		sqlQueryLimit = _this.limit
	}
	

	if( selectFields != ""){

		var localItm = _this.m_SQL_ITEM_DEF;
		var sqlQuerySelect = fmt.Sprintf(`SELECT %s FROM %s %s %s %s %s %s %s`, 
									selectFields, table, localItm, joinTxt, 
									sqlQueryWhere, sqlQueryGroupBy, sqlQueryOrderBy, sqlQueryLimit);
		
		return sqlQuerySelect

	}else{
	
		if( _this.querySelectNewRecord_Text != ""  ) {

			var localItm = _this.m_SQL_ITEM_DEF;
			var sqlBase = fmt.Sprintf(`SELECT %s FROM %s %s %s %s %s %s %s`, 
							_this.querySelectNewRecord_Text, table, localItm, joinTxt, 
							sqlQueryWhere, sqlQueryGroupBy, sqlQueryOrderBy, sqlQueryLimit);
			
			//var sqlQuerySelect = fmt.Sprintf(`SELECT %s FROM  (%s) %s`, selectFields, sqlBase, ITEM)
			return sqlBase
		} else {

			var localItm = _this.m_SQL_ITEM_DEF;
			var sqlBase = fmt.Sprintf(`SELECT %s.* FROM %s %s %s %s %s %s`, 
						localItm, table, localItm, joinTxt, sqlQueryWhere, sqlQueryGroupBy, sqlQueryLimit);
			
			//WHERE (itm1."isActive"=true) )"
			//var sqlQuerySelect = fmt.Sprintf(`SELECT %s FROM  (%s) %s`, selectFields, sqlBase, ITEM)
			return sqlBase
		}
	}
}

func (_this *DBQuery[T]) _Select_getSqlFields(  	
	selectSqlFields map[string]string,
	//excludedFields []string
	 ) []string {

	var retFields = []string{}

	for _, val := range(selectSqlFields) {

		Arr_Append( &retFields, val );
	}
	return retFields;
}

//is a test to exclude the sum and min max from aggregator
func _Select_getSqlFields1[T IGeneric_MODEL, V IGeneric_MODEL]( _this *DBQuery[T], 	
	selectSqlFields map[string]string,
	excludedFields []string ) []string {

	var retFields = []string{}

	//var ctx = _this.tableInst.m_ctx;
	var varV V;
	var reflV = reflect.TypeOf(varV);

	for i:= 0; i < reflV.NumField(); i++{

		var fldT = reflV.Field( i )

		var typeTName= fldT.Type.Name();

		if( typeTName == Generic_MODEL_Name ||
			fldT.Type.Kind() == reflect.Pointer ){//the foreignKey pointer should be excluded
			continue;
		}
		var fld, has = selectSqlFields[ fldT.Name ];
		if( has){
			
			if( _existInListString(excludedFields, fld)){
				continue;
			}
			if( _existInListString(excludedFields, fld)){
				continue;
			}
			Arr_Append( &retFields, fld );
		}
	}
	return retFields;
}
func _existInListString(list []string, item string) bool {

	for j:= 0; j < len(list); j++{
		if( list[j] == item){return true}
	}
	return false;	
}

func _Aggregate_generateSql[T IGeneric_MODEL, V IGeneric_MODEL]( _this *DBQuery[T], ITEM string ) (string, []string) {

	var ctx = _this.tableInst.m_ctx;
	var varT T;
	var varV V;
	var reflT = reflect.TypeOf(varT);
	var reflV = reflect.TypeOf(varV);

	var dictFld = map[string]reflect.StructField{};

	var tableDef map[string] string =  nil
	if( _this.tableInst != nil ){
		tableDef = ctx.SCHEMA_SQL_Columns[ _this.tableInst.m_langName ];
	}

	for i:= 0; i < reflT.NumField(); i++{

		var fldT = reflT.Field( i )

		var typeTName= fldT.Type.Name();

		if( typeTName == Generic_MODEL_Name ||
			fldT.Type.Kind() == reflect.Pointer ){//the foreignKey pointer should be excluded
			continue;
		}

		dictFld[ fldT.Name ] = fldT;
	}
	var selectFields = []string{};
	var groupByFields = []string{};
	var excludeGroupByLangFields = []string{};

	for i:= 0; i < reflV.NumField(); i++{
		
		var fldV = reflV.Field( i )

		var typeVName= fldV.Type.Name();
		if( typeVName == Generic_MODEL_Name ||
			fldV.Type.Kind() == reflect.Pointer ){//the foreignKey pointer should be excluded
			continue;
		}

		var fldT, has = dictFld[ fldV.Name ];
		if( !has ){
			_this.errorRet = fmt.Errorf( "incompatibile struct: %s - %s. Field not found %s",reflV.Name(), reflT.Name(), fldV.Name );
			return "",nil;
		}
		var typeTName= fldT.Type.Name();
		
		var sqlColumnName = fldV.Name;
		if( tableDef != nil ){
			var hasCol = false;
			sqlColumnName, hasCol = tableDef[ fldV.Name ];
			if( !hasCol){
				_this.errorRet = fmt.Errorf("internal error: Not found column %s.", fldV.Name );
				return "",nil;
			}
		}
		Arr_Append( &selectFields, _this._quoteField( sqlColumnName ) );
		
		if( 
		 ( fldV.Type.Kind() == reflect.Slice || 
		   fldV.Type.Kind() == reflect.Array) ){
			typeVName = "[]" + fldV.Type.Elem().Name()
		}

		if( typeVName == typeTName ){
			//the same type is the field that will group
			Arr_Append( &groupByFields, _this._quoteField( sqlColumnName ) );
		}else
		if( typeVName == "[]"+typeTName){
			
			Arr_Append( &excludeGroupByLangFields, fldV.Name );			
			//here are the fields that will be aggregate
		}else{
			_this.errorRet = fmt.Errorf("incompatibile struct: %s - %s. For field %s the types are incompatible ", reflV.Name(), reflT.Name(), fldV.Name );
			return "",nil;
		}
	}
	if( len( selectFields) == len( groupByFields) ){

		_this.errorRet = fmt.Errorf("incompatibile struct for aggregate: %s - %s. No fields to aggregate. you must have minimum 1 field with same name and different type ([]Type) in both structs ", reflV.Name(), reflT.Name() );
		return "",nil;
	}
	if( len( groupByFields) == 0  ){

		_this.errorRet = fmt.Errorf("incompatibile struct for aggregate: %s - %s. No fields to group by. you must have minimum 1 field with same name and type in both structs ", reflV.Name(), reflT.Name() );
		return "",nil;
	}

	var sqlSelectFields    = Str.Join( selectFields, ", ") ;
	//var sqlGrouppingFields = Str.Join( groupByFields, ", ") ;

	//var table = _this._generateSqlSourceOfData();
	//var sqlQuery = "";
	//sqlQuery += fmt.Sprintf(`SELECT %s FROM %s %s GROUP BY %s`, sqlSelectFields, table, ITEM, sqlGrouppingFields)
	//sqlQuery += fmt.Sprintf(`%s FROM %s %s`, sqlSelectFields, table, ITEM)
	//return sqlQuery, excludeGroupByLangFields;

	return sqlSelectFields, excludeGroupByLangFields
}

func (_this *DBQuery[T]) _Aggregate_generateGroupBySql( selectFields []string, excludeFields []string ) string {	

	var grouppingFields = []string{};
	for i := 0; i < len(selectFields); i++{

		var fldName = selectFields[i];
		var bExclude = false;
		for j := 0; j < len(excludeFields); j++{
			if( fldName == excludeFields[j]){ bExclude = true; break;}
		}
		if( !bExclude ){
			Arr_Append( &grouppingFields, fldName);
		}
	}
	var sqlGrouppingFields =  Str.Join( grouppingFields, ", ") ;

	var sqlQuery = fmt.Sprintf(`GROUP BY %s`, sqlGrouppingFields)
	return sqlQuery
}

/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _getRows(
	bDistinct bool, fields []string, bAddFieldsInSelect bool) string {

	var sourceQry = _this._generateSqlSourceOfData( )

	var sqlQuery = ""
	var ITEM = _this.m_SQL_ITEM_DEF

	var joinTxt = ""
	var arrFieldsSql = []string{}

	//c#
	//this.whereTxt = _whereFilterSearch(this.whereTxt);
	//this.whereTxt = _whereFilter(this.whereTxt);

	if fields != nil {
		for _, fld := range fields {

			var fldDef = _this._quoteTableField1(fld, true, _this.getJoins(), !bDistinct && bAddFieldsInSelect)
			if fldDef != "" { //it is possible sometime the defs of view to be empty

				Arr_Append(&arrFieldsSql, fldDef)
			}
		}
	}

	if _this.getJoins() != nil {
		joins1 := _this.getJoins()._joins
		for _, join := range joins1 {

			joinTxt += join.getSqlTxt(ITEM)
		}
	}

	if fields != nil {
		//bFirst := true
		sqlQuery = ""

		var sqlQueryFields = Str.Join(arrFieldsSql, ", ")

		if bDistinct {
			sqlQuery = fmt.Sprintf("SELECT DISTINCT %s FROM %s %s", sqlQueryFields, sourceQry, ITEM)
		} else {
			if !bAddFieldsInSelect {
				sqlQuery = fmt.Sprintf("SELECT %s FROM %s %s", sqlQueryFields, sourceQry, ITEM)
			} else {
				sqlQuery = fmt.Sprintf("SELECT %s.*, %s FROM %s %s", ITEM, sqlQueryFields, sourceQry, ITEM)
			}
		}
	} else {
		if( _this.querySelectNewRecord_Text != ""  ){
			/*_this.lamdaSelectNewRecord != ""*/ 

			sqlQuery += fmt.Sprintf("SELECT %s FROM %s %s",  _this.querySelectNewRecord_Text, sourceQry, ITEM )
		}else
		if joinTxt != "" { //pt ca altfel da eroare, left joinul randeaza si ID-ul de la alte tabele si se suprascrie
			sqlQuery += fmt.Sprintf("SELECT %s.* FROM %s %s", ITEM, sourceQry, ITEM)
		} else {
			sqlQuery += fmt.Sprintf("SELECT * FROM %s %s", sourceQry, ITEM)
		}
	}

	sqlQuery += joinTxt
	if _this.whereTxt != "" {
		sqlQuery += fmt.Sprintf(" WHERE %s", _this.whereTxt)
	}
	
	if _this.excludeLangFieldsFromGroupBy != nil &&  _this.selectSqlFields != nil {
		sqlQuery += " " + _this._Aggregate_generateGroupBySql( _this.selectSqlFields, _this.excludeLangFieldsFromGroupBy ) 
	}

	if _this.orderBy != "" {
		sqlQuery += " " + _this.orderBy
	}
	if _this.limit != "" {
		sqlQuery += " " + _this.limit
	}

	return sqlQuery
}

/*#PHPARG=[ Array< T > ];*/
func (_this *DBQuery[T]) _arrayRecordsAny(sqlResult *sql.Rows, fnNewInst func()any) ([]any, error) {
	//_this.clearCachedSyntax();
	var retList = []any{}

	
	for sqlResult.Next() {

		var model any = fnNewInst()
		var err1 = _this.readModelSqlResult(sqlResult, model)
		if( err1 != nil) {
			return  nil, err1
		}
		Arr_Append(&retList, model)
	}
	//_this.tableInst.m_ctx.updateDeltaTime();
	return retList, nil
}


/*#PHPARG=[ Array< T > ];*/
func (_this *DBQuery[T]) _arrayRecords(sqlResult *sql.Rows, fields []string) ([]*T, error) {
	//_this.clearCachedSyntax();
	retList := []*T{}
	
	for sqlResult.Next() {

		model := new(T)
		var err1 = _this.readRecordSqlResult( sqlResult, model, fields )
		if( err1 != nil) {
			return  nil, err1
		}
		Arr_Append(&retList, model)
	}
	return retList, nil
}

/*#PHPARG=[ Array< T > ];*/
func (_this *DBQuery[T]) _arrayModels(sqlResult *sql.Rows) ([]*T, error) {
	//_this.clearCachedSyntax();
	retList := []*T{}
	
	for sqlResult.Next() {

		model := new(T)
		var err1 = _this.readModelSqlResult( sqlResult, model )
		if( err1 != nil) {
			return  nil, err1
		}
		Arr_Append(&retList, model)
	}
	return retList, nil
}

func (_this *DBQuery[T])  result_getModelHeaderColumn( model any) []interface{}{

	//var strAttr = string( ATTR_ATOMICSQL_COPY_MODEL );
	var reflVal  = reflect.ValueOf(model).Elem()
	var reflType = reflect.TypeOf(model).Elem();

	numCols := reflVal.NumField()
	var columns  []interface{}

	for i := 0; i < numCols; i++ {

		var field  = reflVal.Field(i)

		var fldType = reflType.Field( i )
		var nameFld = fldType.Name
		if( nameFld == ""){}

		var nameTypeFld = field.Type().Name();
		if( nameTypeFld == Generic_MODEL_Name ){ 
			continue
		}
		if( field.Type().Kind() == reflect.Pointer){ 
			continue
		}		
		if( fldType.Tag != "" && string(fldType.Tag) == ATTR_ATOMICSQL_COPY_MODEL){

			// SELECT( x=> {User = *x.user; ...}
			var cols = _this.result_getModelHeaderColumn( field.Addr().Interface() )
			Arr_AddRange( &columns, &cols)
		}else{
			Arr_Append( &columns, field.Addr().Interface() )
		}
	}
	return columns;
}

func (_this *DBQuery[T])  readModelSqlResult(rows *sql.Rows, model any) error{

	var columns = _this.result_getModelHeaderColumn( model);
	var err = rows.Scan(columns...)
	return err;
}

func (_this *DBQuery[T])  result_getRecordHeaderColumn( model any, fields []string) []interface{}{

	//var strAttr = string( ATTR_ATOMICSQL_COPY_MODEL );
	var reflVal  = reflect.ValueOf(model).Elem()
	var reflType = reflect.TypeOf(model).Elem();

	numCols := reflVal.NumField()
	var columns  []interface{}

	for i := 0; i < numCols; i++ {

		var field  = reflVal.Field(i)

		var fldType = reflType.Field( i )
		var nameFld = fldType.Name
		if( nameFld == ""){}
		
		if( fields != nil){
			if( !Arr_Contains( &fields, nameFld) ){
				continue;
			}
		}

		var nameTypeFld = field.Type().Name();
		if( nameTypeFld == Generic_MODEL_Name ){ 
			continue
		}
		if( field.Type().Kind() == reflect.Pointer){ 
			continue
		}
		
		if( fldType.Tag != "" && string(fldType.Tag) == ATTR_ATOMICSQL_COPY_MODEL){

			// SELECT( x=> {User = *x.user; ...}
			var cols = _this.result_getRecordHeaderColumn( field.Addr().Interface(), fields )
			Arr_AddRange( &columns, &cols)
		}else{
			Arr_Append( &columns, field.Addr().Interface() )
		}
	}
	return columns;
	
}

func (_this *DBQuery[T])  readRecordSqlResult(rows *sql.Rows, model any, fields []string) error{

	var columns = _this.result_getRecordHeaderColumn( model, fields);
	err := rows.Scan(columns...)
	return err;
}
//----------------------------------------------------------------------------
func (_this *DBQuery[T]) _arrayOfSingleField(sqlResult *sql.Rows, fieldName string) []string {

	var retList = []string{}

	var model = new(T)
	for sqlResult.Next() {

		var strVal = readRecordSqlResult_Readfield(*sqlResult, model, fieldName)
		//err = sqlResult.Scan(&user.ID, &user.Username, &user.Password, &user.Tel)
		Arr_Append(&retList, strVal)
	}
	return retList
}

func readRecordSqlResult_Readfield[T IGeneric_MODEL](rows sql.Rows, model *T, fieldName string) string {

	s := reflect.ValueOf(model).Elem()
	//numCols := s.NumField()
	columns := make([]interface{}, 1)

	field := s.FieldByName(fieldName)
	columns[0] = field.Addr().Interface()

	err := rows.Scan(columns...)
	if err != nil {

	}
	return field.String()
}

/*#PHPARG=[ T ];*/
func (_this *DBQuery[T]) _oneModel( /*#mysqli_result*/ sqlResult *sql.Rows) (*T, error) {
	
	for sqlResult.Next() {

		var model = new(T)
		//var columns = _this.result_getHeaderColumn( model);
		var err = _this.readModelSqlResult( sqlResult, model)
		if( err != nil){
			return nil, err
		}
		return model, nil
	}
	return nil, nil
}
func (_this *DBQuery[T]) _oneRecord( /*#mysqli_result*/ sqlResult *sql.Rows, fields[]string) (*T, error) {
	//_this.clearCachedSyntax();	

	for sqlResult.Next() {

		var record = new(T)
		//var columns = _this.result_getHeaderColumn( model);
		err := _this.readRecordSqlResult( sqlResult, record, fields )
		if( err != nil){
			return nil, err
		}

		return record, nil
	}
	return nil, nil
}

func (_this *DBQuery[T]) setLimit(queryOffset int, querySize int) *DBQuery[T] {

	queryOffsetI := (queryOffset)
	querySizeI := (querySize)

	if querySizeI > 0 {
		_this.limit = fmt.Sprintf(" LIMIT %d OFFSET %d", querySizeI, queryOffsetI)
	} else if queryOffsetI > 0 {
		_this.limit = fmt.Sprintf(" LIMIT %d  OFFSET %d", 4000000000000, queryOffsetI)
	}
	return _this
}

/*
func (_this *DBQuery[T]) singleDataI(mysqliResult sql.Rows, field string) int {
	_this.clearCachedSyntax()
	model := _this._oneRecord(mysqliResult)

	return int(model[field])
}


func (_this *DBQuery[T]) singleDataF(mysqliResult sql.Rows, field string) float64 {
	_this.clearCachedSyntax()
	model := _this._oneRecord(mysqliResult)

	return float64(model[field]);
}
*/

/*#PHPARG=[ VOID];*/
func (_this *DBQuery[T]) checkMySqlError1( /*#String*/ sqlQuery string) {
	//_this.clearCached();
	//errMsg = mysqli_error( _this.g_DB.connection );
	//_this.last_error = "errMsg. GENERATED BY:sqlQuery";
	//UtilLog::logMsg("checkMySqlError() : errMsg [sqlQuery]");
}

/*#PHPARG=[ VOID];*/
func (_this *DBQuery[T]) checkMySqlError( /*#String*/ sqlQuery string, err error) {
	//_this.clearCachedSyntax();
	//_this.checkMySqlError1( sqlQuery);
	
	if( err != nil){
		_this.tableInst.m_ctx.hasError = err;
		log.Printf("sql query error: %s %s", sqlQuery, err.Error() )
	}
}

/*#PHPARG=[ String ];*/
/**
* for deleteRecors() I need to rebuild again because the DELETE is not support ITEM.
* @param type query
* @return string
 */
func (_this *DBQuery[T]) generateSqlText( /*#DBSqlQuery< T >*/ query *DBSqlQuery[T]) string {

	if query.fnWhere != nil {

		query1 := _this._whereGeneric( query.fnWhere ) //"(opText1) AND (opText2)" );
		return query1.getText()
	}
	_op := query.m_op
	if _op != "" {
		if _op == "AND" || _op == "OR" {
			var arr []string = []string{}
			for iQuery := 0; iQuery < len(query.m_listOperands); iQuery++ {

				txt1 := fmt.Sprintf("(%s)", _this.generateSqlText(query.m_listOperands[iQuery]))
				array_push(&arr, txt1)
			}
			var _op1 = fmt.Sprintf(" %s ", _op)
			ret := Str.Join(arr, _op1)
			return fmt.Sprintf(" ( %s ) ", ret)
		} else if _op == "IN" || _op == "NOT_IN" {
			var op1 = _this._quoteTableField(query.m_field1, false, _this.getJoins())

			var where1 = IFFs(_op == "NOT_IN",
				fmt.Sprintf("%s NOT IN (", op1),
				fmt.Sprintf("%s IN (", op1))
			bFirst := true

			for iQuery := 0; iQuery < len(query.m_listOperandsStr); iQuery++ {
				//foreach( query.m_listOperandsStr as opIn )

				opIn := query.m_listOperandsStr[iQuery]
				op := _this._quote(opIn, nil)

				where1 += fmt.Sprintf("%s%s", IFFs(bFirst, "", ", "), op)
				bFirst = false
			}
			where1 += " )"

			return where1
		} else {
			if query.m_field1 != "" {
				if query.m_field2 != "" {
					o1 := _this._quoteTableField(S(query.m_field1), false, _this.getJoins())
					o2 := _this._quoteTableField(S(query.m_field2), false, _this.getJoins())
					return fmt.Sprintf("%s %s %s", o1, _op, o2)
				} else {
					o1 := _this._quoteTableField(S(query.m_field1), false, _this.getJoins())
					o2 := _this._quote(query.m_operand2, nil)
					return fmt.Sprintf("%s %s %s", o1, _op, o2)
				}
			} else {
				if query.m_field2 != "" {
					o1 := _this._quote(query.m_operand1,nil)
					o2 := _this._quoteTableField(S(query.m_field2), false, _this.getJoins())
					return fmt.Sprintf("%s %s %s", o1, _op, o2)
				} else {
					o1 := _this._quote(query.m_operand1, nil)
					o2 := _this._quote(query.m_operand2, nil)
					return fmt.Sprintf("%s %s %s", o1, _op, o2)
				}
			}
		}
	} else {
		return query.text
	}
}



func (_this *DBQuery[T]) get_SQL_ITEM_DEF() string {
	//if( _this.pivotProvider )
	//{
	//    return SQL_ITEM_DEF_J;
	//}
	return _this.m_SQL_ITEM_DEF
}

func (_this *DBQuery[T]) getJoins() *DBSqlJoinCollection {
	//if( _this.pivotProvider )
	//    return _this.pivotProvider.joins;

	return _this.joins
}

func (_this *DBSqlJoinCollection) createJoinCollection(){

	if( _this._joins == nil ){
		_this._joins = make(map[string]*DBSqlJoin)
	}
}





func (_this *DBSqlJoinCollection) addJoin(
	ctx *DBContextBase,
	thisQuery IQueryBase,
	/*#String*/ itemTxt string,
	/*#String*/ FK_id string /*#String*/, FK string /*#String*/, itemFK string,
	/*#String*/ out_tableDBName *string) (string, error) {
	

	var table = ""
	var tableFK = ""
	var ID  = DEF_TABLE_COLUMN_ID

	valFk, has := ctx.FOREIGN_KEYS[FK];
	if has {

		tableFK = valFk.RootFldFk_sqlName
		table   = valFk.TgtTable_sqlName
		ID      = valFk.TgtFldID_sqlName
		*out_tableDBName = table
	}else{
		return "", fmt.Errorf("%s: internal error. Foreign map is not defined", ORM_NAME)
	}

	_this.createJoinCollection();

	var joinItem, hasJoin = _this._joins[FK_id+":"+itemFK]
	if hasJoin {
		return joinItem.nameItem, nil
	}

	j := new(DBSqlJoin)
	_this._joins[FK_id+":"+itemFK] = j

	/*if( pivotTableName && pivotTableName == table )
	{   //pt ca pivot va deveni ITEM
		j.nameItem = SQL_ITEM_DEF;
		j.sql = "";//pt ca se duplica
		return j.nameItem;
	}else*/
	{
		nameItem := fmt.Sprintf("%s%d", SQL_ITEM_DEF, len(_this._joins))
		j.nameItem = nameItem
		j.sql = fmt.Sprintf(" LEFT JOIN %s %s ON %s.%s = %s.%s ",
							thisQuery._quoteTable(table), nameItem, nameItem, thisQuery._quoteField(ID), itemFK, thisQuery._quoteField(tableFK) )
		return nameItem, nil
	}
}

func (_this *DBQuery[T]) newJoinCollection(){

	_this.joins = (new (DBSqlJoinCollection)).Constr();
	//_this.joins.createJoinCollection();
}

func (_this *DBQuery[T]) clearCachedSyntax() {

	_this.tableInst.m_ctx.clearCachedSyntax();
}

func (_this *DBQuery[T]) eqFO( /*#T_FieldName*/ field string, operand1 any, bNot bool) *DBSqlQuery[T] {

	var op1 = _this._quoteTableField(field, false, _this.getJoins())
	//OP1 = _this._uniformizeField( table, op1 );

	if operand1 == nil {

		if !bNot {
			var retSqlQuery = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s IS NULL ", op1))
			return retSqlQuery
		} else {
			var retSqlQuery = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s IS NOT NULL ", op1))
			return retSqlQuery
		}
	} else {
		op2 := _this._quote(operand1, nil)

		var ret *DBSqlQuery[T] = nil
		if !bNot {
			ret = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s = %s", op1, op2))
			ret.m_op = "="
		} else {
			ret = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s <> %s", op1, op2))
			ret.m_op = "<>"
		}
		ret.m_field1 = field
		ret.m_operand2 = operand1
		return ret
	}
}

func (_this *DBQuery[T]) eqFF( /*#T_FieldName*/ field string, field2 string, bNot bool) *DBSqlQuery[T] {

	var op1 = _this._quoteTableField(field, false, _this.getJoins())
	var op2 = _this._quoteTableField(field2, false, _this.getJoins())

	var ret *DBSqlQuery[T] = nil
	if !bNot {
		ret = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s = %s", op1, op2))
		ret.m_op = "="
	} else {
		ret = (new(DBSqlQuery[T])).Constr(fmt.Sprintf("%s <> %s", op1, op2))
		ret.m_op = "<>"
	}
	ret.m_field1 = field
	ret.m_field2 = field2
	return ret
}

/*#PHPARG=[ VOID ];*/
func (_this *DBQuery[T]) setForeignKeys( /*#HashMap<String>*/ withForeignKeys map[string]string) {
	_this.withForeignKeys = withForeignKeys
}

//---------------------------------------------------------------------------------------
func (_this *DBQuery[T]) _InsertRecordByReflectValue(	
			langTableName string, reflValue reflect.Value, 
			bInsertID bool, fields *[]string) (int64, error) {

	var ctx = _this.tableInst.m_ctx

	tableData := _this.tableInst.m_ctx.SCHEMA_SQL[langTableName]

	primaryLangKey := tableData.PrimaryColumnLangName
	if primaryLangKey == "" {
		primaryLangKey = DEF_TABLE_COLUMN_ID
	}
	

	sqlQuery, err := _this._insertRecord(langTableName, reflValue, bInsertID, fields)
	if sqlQuery != "" && err == nil {

		if( _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql){

			ctx.currOperationDTime2 = time.Now()
			dbResult1, err := _this.tableInst.m_ctx.Exec(sqlQuery)
			ctx.updateDeltaTime2()

			if dbResult1 != nil && err == nil {
				
				fld := reflect.ValueOf(reflValue).FieldByName(primaryLangKey)
				if !bInsertID {

					lastID, err := _this.getLastInsertedRowID1(dbResult1) //_this.tableName, primarySqlKey)
					if err == nil {
						fld.SetInt(lastID)
					}
				}
				ctx.updateDeltaTime();
				return fld.Int(), nil
			}
		}else{
		//	if( _this.tableInst.m_ctx.Dialect == ESqlDialect.Postgress ||
		//	   _this.tableInst.m_ctx.Dialect == ESqlDialect.MSSQL){	

			ctx.currOperationDTime2 = time.Now()		
			dbResultRows, err := _this.tableInst.m_ctx.Query(sqlQuery)
			defer queryClose( dbResultRows )
			ctx.updateDeltaTime2()	

			if dbResultRows != nil && err == nil {
				
				v := reflValue;
				fldValue := v.FieldByName(primaryLangKey)
				if !bInsertID {

					//lastID := fldValue.Int()

					lastID, err := _this.getLastInsertedRowsID(dbResultRows) //_this.tableName, primarySqlKey)
					if err == nil {
						if( fldValue.CanSet()){

							fldValue.SetInt(lastID)
						}						
					}
				}
				ctx.updateDeltaTime();
				return fldValue.Int(), nil
			}
		}		
	}

	_this.checkMySqlError(sqlQuery, err)
	return 0, err
}

func queryClose(result* sql.Rows){
	if( result != nil){
		result.Close()
	}
}

//---------------------------------------------------------------------------------------
/*#PHPARG=[ String ];*/
func (_this *DBQuery[T]) _insertRecord( langTableName string, modelValue reflect.Value, 
	/*#BOOL*/ bInsertID bool, fields *[]string) (string,error) {
	//g_DB           = _this.g_DB;
	
	ctx := _this.tableInst.m_ctx
	fieldsSchema := ctx.SCHEMA_SQL[langTableName]

	primaryKeySql := fieldsSchema.PrimaryColumnSqlName
	modelTableName := fieldsSchema.SqlTableName//_this.tableName

	bFirst := true
	sqlFieldsName := ""
	sqlFieldsData := ""
	
	//schemaTable := fieldsSchema.schemaTable
	//fieldsSchema = _this.g_DB.SCHEMA[ _this.tableNameOrig ];
	//arrFieldsSchema := array_keys( fieldsSchema );
	//if( fields ){
	//	arrFieldsSchema = fields;
	//}

	//foreach( modelData as columnTable => value )
	for _, columnTable := range fieldsSchema.Columns {

		//var modelInfo = reflect.ValueOf( modelData)
		//for iFld :=0; iFld < numFields; iFld++{
		//var fieldInfo = reflect.ValueOf( modelData).Field( iFld ); //isset()

		if columnTable.LangName == "" {
			continue
		}

		var reflctModel = modelValue;//reflect.ValueOf(*modelData)
		var fieldInfo = reflctModel.FieldByName(columnTable.LangName) //isset()

		if !bInsertID && columnTable.IsPrimary {
			//do not insert in udate cmd the primary key
			continue
		}
		//if( !columnTable.isNullable && (columnTable.sqlType == "uuid" || columnTable.sqlType == "UUID")){
		//	continue // insert the UUID created by the dfault key generator
		//}
		//typeVal = fieldsSchema[columnTable];
		var valSql string = ""
		var err error
		//var valSet bool = false;
		//var unusedDateTime time.Time

		//var fieldInfoType = fieldInfo.Type();
		//var ss = fieldInfoType.Name();
		//fmt.Print(ss);

		
		valSql, err = _this._insertRecord_setFieldGeneral( fieldInfo, reflctModel, &columnTable, 
							ctx, bInsertID, fields )
		if( err != nil){
			return "",err
		}
		

		if !bFirst {
			sqlFieldsName += ", "
			sqlFieldsData += ", "
		}
		bFirst = false
		columnName := columnTable.SqlName // eok ? sau trebuia langName ?

		if _this.withForeignKeys != nil {
			//fkField = Utils_getOriginalSourceForForeignKey( withForeignKeys, columnTable);
			fkField := columnName
			if fkField != "" && isset(_this.withForeignKeys[fkField]) {

				fkValue := _this.withForeignKeys[fkField]
				fkField := _this._quoteField(fkField)
				fkTable := _this._quoteTable(_getDotPrefix(fkValue))
				fkTable_FK := _this._quoteTableField(fkValue, false, nil)

				DEF_TABLE_COLUMN_ID := DEF_TABLE_COLUMN_ID

				sqlFieldsName += fkField
				sqlFieldsData += fmt.Sprintf("( SELECT %s FROM %s WHERE %s = %s LIMIT 1 )",
					DEF_TABLE_COLUMN_ID, fkTable, fkTable_FK, valSql)
			} else {

				sqlFieldsName += fmt.Sprintf("%s.%s",
					_this._quoteField(modelTableName),
					_this._quoteField(columnName))
				sqlFieldsData += valSql
			}
		} else {
			sqlFieldsName += fmt.Sprintf("%s",
				//_this._quoteTable(modelTableName),
				_this._quoteField(columnName))
			sqlFieldsData += valSql
		}
	}

	sqlQuery := ""
	if( ctx.Dialect == ESqlDialect.Postgress ){

		sqlQuery = fmt.Sprintf(`INSERT INTO %s
				 (%s) VALUES( %s ) 
				 RETURNING %s`, 
				 _this._quoteTable(modelTableName), 
				 sqlFieldsName, sqlFieldsData, 
				 _this._quoteField(primaryKeySql) )
				 
	}else
	if( ctx.Dialect == ESqlDialect.MsSql ){
		
		sqlQuery = fmt.Sprintf(`INSERT INTO %s
				 (%s) VALUES( %s ) LAST_INSERT_ID()`, 
				 _this._quoteTable(modelTableName),
				 sqlFieldsName, sqlFieldsData)
	}else
	{
	//if( ctx.Dialect == ESqlDialect.MYSql )
	sqlQuery = fmt.Sprintf(`INSERT INTO %s
				 (%s) VALUES( %s )`, 
				 _this._quoteTable(modelTableName), 
				 sqlFieldsName, sqlFieldsData)
	}

	return sqlQuery, nil
}

func (_this *DBQuery[T]) _insertRecord_setFieldGeneral( 
	fieldInfo reflect.Value,
	reflctModel reflect.Value,
	columnTable *TSqlColumnDef,
	ctx *DBContextBase,
	/*#BOOL*/ bInsertID bool, fields *[]string,

	) (string,error) {

	var fieldInfoType = fieldInfo.Type();
	var valSql = ""
	var err error	

	var fldNullable *_TNullableField

	if fieldInfoType.Kind() == reflect.Slice {

		var typeSqlElement = Str.ReplaceAll( columnTable.SqlType,"[]", "" )

		var valLen = fieldInfo.Len();
		var lstElem  = [] string{}

		for idx := 0; idx < valLen; idx++ {
			var elem = fieldInfo.Index( idx )

			elemType := elem
			//var ss = elemType.Kind()
			//fmt.Println(ss)

			var valSqlElem, err1 = _this._insertRecord_setFieldGeneral( 
				elemType,
				reflctModel,
				columnTable,
				ctx,
				bInsertID , fields );
			if( err1 != nil){
				return "", err1
			}

			Arr_Append( &lstElem, valSqlElem )			
		}
		valSql1 := Str.Join( lstElem, ", " );
		valSql = fmt.Sprintf( "ARRAY[ %s ]::%s[]", valSql1, typeSqlElement );
	}else{
	{
		fldNullable = _this._getNullableField( fieldInfo);
		if( fldNullable != nil){
		{
			var bIsNotNull = fldNullable.Valid.Bool()
			if( !bIsNotNull ){ 

				valSql = ctx.LangDB.VALUE_NULL

				if( columnTable.ForeignKeyLangName != nil ){//if the field is null, should be inserted

					var err error
					var FK_ID int64 = 0
		
					valSql, err, FK_ID = _this._insertRecord_setField_ForeignKey( 
									fieldInfo, reflctModel, columnTable, ctx,
									bInsertID, fields, true )
					if( err != nil){
						return "", err
					}
					if( fldNullable.Value.CanSet() ){
						fldNullable.Value.SetInt( FK_ID )
						fldNullable.Valid.SetBool( true )
					}
				}
			}else{
				
				valSql, err = _this._insertRecord_setField( fldNullable.Value, reflctModel, columnTable, 
						ctx, bInsertID, fields )	
				if( err != nil){
					return "",err
				}
			}
		}}else{

			valSql, err = _this._insertRecord_setField( fieldInfo, reflctModel, columnTable, 
								ctx, bInsertID, fields )
			if( err != nil){
				return "",err
			}
		}
	}}
	return valSql, nil
}

type  _TNullableField struct{
	Valid reflect.Value
	Value reflect.Value

}

func (_this *DBQuery[T]) _getNullableField( 
	 fieldInfo reflect.Value ) *_TNullableField{

	var fieldInfoType = fieldInfo.Type();

	if( fieldInfoType == reflect.TypeOf((*sql.NullBool)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullTime)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullString)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullByte)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullInt16)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullInt32)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullInt64)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*sql.NullFloat64)(nil)).Elem()   ) {
	{
		fldValid:= fieldInfo.FieldByName("Valid");
		nullableFieldName := ""

		switch fieldInfoType {
			case reflect.TypeOf((*sql.NullBool)(nil)).Elem():{				
				nullableFieldName = "Bool";
			}
			case reflect.TypeOf((*sql.NullTime)(nil)).Elem():{				
				nullableFieldName = "Time";
			}
			case reflect.TypeOf((*sql.NullString)(nil)).Elem():{				
				nullableFieldName = "String";
			}
			case reflect.TypeOf((*sql.NullByte)(nil)).Elem():{				
				nullableFieldName = "Byte";
			}
			case reflect.TypeOf((*sql.NullInt16)(nil)).Elem():{				
				nullableFieldName = "Int16";
			}
			case reflect.TypeOf((*sql.NullInt32)(nil)).Elem():{				
				nullableFieldName = "Int32";
			}
			case reflect.TypeOf((*sql.NullInt64)(nil)).Elem():{				
				nullableFieldName = "Int64";
			}
			case reflect.TypeOf((*sql.NullFloat64)(nil)).Elem():{				
				nullableFieldName = "Float64";
			}
		}
		fldNullValue:= fieldInfo.FieldByName(nullableFieldName );
		return &_TNullableField{ Valid:fldValid, Value:fldNullValue}
	}}
	return nil
}

type _TRetForeignKey_getID struct{
	ID int64
	tableFk string
	ptrFkItf reflect.Value
} 

func (_this *DBQuery[T])  getModel_ForeignKey_getID(
	ctx *DBContextBase, columnTable *TSqlColumnDef, reflctModel reflect.Value) *_TRetForeignKey_getID{

	var tableFk = columnTable.ForeignKeyLangName[1];
	primaryKey := ctx.SCHEMA_SQL[tableFk].PrimaryColumnLangName

	fldFkName := columnTable.ForeignKeyLangName[0];
	reflectFk := reflctModel.FieldByName(fldFkName)
	ptrFkItf  := reflect.Indirect(reflectFk);//
	
	if( !(ptrFkItf.CanAddr() && ptrFkItf.IsValid()) ){
		return nil
	}

	fldPrimary := ptrFkItf.FieldByName(primaryKey)						
	var ID = fldPrimary.Int()

	return &_TRetForeignKey_getID{ ID:ID, tableFk:tableFk, ptrFkItf:ptrFkItf }
}

func (_this *DBQuery[T]) _insertRecord_setField_ForeignKey( 
	fieldInfo reflect.Value,
	reflctModel reflect.Value,
	columnTable *TSqlColumnDef,
	ctx *DBContextBase,
	/*#BOOL*/ bInsertID bool, fields *[]string,
	fieldIsNullable bool,

	) (string, error, int64) {

	//var valSql = valFK_ID//fmt.Sprintf("%d", fieldInfo.Int())
	
	var valSql = ""
	var retGetID = _this.getModel_ForeignKey_getID(ctx, columnTable, reflctModel)
	
	if( retGetID == nil ){
		if( fieldIsNullable ){
			return ctx.LangDB.VALUE_NULL, nil, 0
		}else{
			return "", fmt.Errorf("cannot access address of FK"), 0
		}
	}

	var ID = retGetID.ID
	valSql = fmt.Sprintf("%d", ID)

	if( valSql == "0") {
		//I must insert also the fk in DB

		lastFkID, err1 :=_this._InsertRecordByReflectValue(
			retGetID.tableFk, retGetID.ptrFkItf, bInsertID, fields)
		if( err1 != nil)	{
			return "", err1, 0
		}
		/*
		colType := reflctModel.FieldByName(columnTable.langName)
		if( colType.CanSet() ){
			if( fieldIsNullable){
				colType.SetInt( int64(lastFkID) )
			}else{
				colType.SetInt( int64(lastFkID) )	
			}
		}*/
		
		valSql = fmt.Sprintf("%d", lastFkID )
		//ptrFkItf.Pointer()
		//lastFkID := InsertRecord(data *T, bInsertID bool, fields []string) (int, error) {		
		return valSql, nil, lastFkID
	}
	return valSql, nil, 0
}


func (_this *DBQuery[T]) _insertRecord_setField( 
	fieldInfo reflect.Value,
	reflctModel reflect.Value,
	columnTable *TSqlColumnDef,
	ctx *DBContextBase,
	/*#BOOL*/ bInsertID bool, fields *[]string,

	) (string,error) {

	valSql := ""
	fieldInfoType := fieldInfo.Type()
	
	if fieldInfoType == reflect.TypeOf((*int16)(nil)).Elem() ||
	   fieldInfoType == reflect.TypeOf((*int32)(nil)).Elem() ||
	   fieldInfoType == reflect.TypeOf((*int64)(nil)).Elem() ||
	   fieldInfoType == reflect.TypeOf((*int)(nil)).Elem() ||
	   fieldInfoType == reflect.TypeOf((*int8)(nil)).Elem() {
		
		if( columnTable.ForeignKeyLangName != nil ){

			valSql = fmt.Sprintf("%d", fieldInfo.Int())
			var err error
			var FK_ID int64 = 0

			if( valSql == "0" ){ //if is 0 should be inserted
				valSql, err, FK_ID = _this._insertRecord_setField_ForeignKey( 
								fieldInfo, reflctModel, columnTable, ctx,
								bInsertID, fields, false )
				if( err != nil){
					return "", err
				}
				if( fieldInfo.CanSet() ){
						fieldInfo.SetInt( FK_ID )	
				}
			}
			
		}else{
			valSql = fmt.Sprintf("%d", fieldInfo.Int())
		}
	} else 
	if fieldInfoType == reflect.TypeOf((*bool)(nil)).Elem() {

		var value = fieldInfo.Bool()
		if value == true /*|| value == "true"*/ {
			valSql = ctx.LangDB.VALUE_TRUE
		} else if value == false /*|| value == "false"*/ {
			valSql = ctx.LangDB.VALUE_FALSE
		}
	} else 
	if  fieldInfoType == reflect.TypeOf((*float32)(nil)).Elem() ||
		fieldInfoType == reflect.TypeOf((*float64)(nil)).Elem() {

		valSql = fmt.Sprintf("%f", fieldInfo.Float())
	} else if fieldInfoType == reflect.TypeOf((*string)(nil)).Elem() {

		valSql = _this._quote(fieldInfo.String(), columnTable)
	} else if fieldInfoType == reflect.TypeOf((*time.Time)(nil)).Elem() {
		valSql = _this._quote(fieldInfo.String(), columnTable)
	}else
	if fieldInfoType == reflect.TypeOf((*[]uint8)(nil)).Elem() {

		var slice = fieldInfo.Bytes()
		valSql = _this._quote(slice, columnTable)
	}else{
		//ceva nu e in ordine, ori e fk
		valSql = "";
		
	}	
	return valSql, nil
}

func (_this *DBQuery[T]) getLastInsertedRowID1(result sql.Result) (int64, error) {

	if _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql {
		
		return result.LastInsertId()
	}else {
		//query = db.Rebind(query)
		//var id int
		//err = db.Get(&id, query, args...)      
	}
	return 0, nil
}

func (_this *DBQuery[T]) getLastInsertedRowID(result *sql.Row) (int64, error) {
	
	if _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql {
		return 0, nil;
		//return result.LastInsertId()
	}else {
		var id int64 = 0;
		err := result.Scan(&id);
		return id, err
	}
}
func (_this *DBQuery[T]) getLastInsertedRowsID(result *sql.Rows) (int64, error) {

	if _this.tableInst.m_ctx.Dialect == ESqlDialect.MySql {
		return 0, nil;
		//return result.LastInsertId()
	}else {
		var id int64 = 0;

		for result.Next() {

			err := result.Scan(&id);
			return id, err
		}
	}
	return 0, nil
}


func (_this *DBQuery[T]) _deleteRecords() string {

	ctx := _this.tableInst.m_ctx
	var modelTableName =_this.tableName;
	var modelSchema =_this.schemaTable;
	
	if(_this.m_queryAND != nil) {

		_this.m_SQL_ITEM_DEF = modelTableName;
		//_this.getJoins().createJoinCollection();

		_this.whereTxt =_this.generateSqlText(_this.m_queryAND );
	}else{
		_this.whereTxt = "1=1"
	}

	var bHasJoins bool = _this.getJoins() != nil && _this.getJoins()._joins != nil;

	var sqlQuery = ""

	if( ctx.Dialect == ESqlDialect.MySql){
	
		sqlQuery = fmt.Sprintf(" DELETE %s FROM %s ", modelTableName, modelTableName);
	}else
	if( ctx.Dialect == ESqlDialect.Postgress){

		if( bHasJoins ){

			/*
			DELETE FROM  "User" AS u 
			USING 
				(SELECT "ID", "RoleName" from "UserRole"  ) AS ur      
			WHERE 
				ur."ID" = u."UserRole_ID"
			AND ( ur."RoleName" = 'aa' )
			*/
			sqlQuery = fmt.Sprintf(" DELETE %s FROM %s.%s ", 
							_this.m_SQL_ITEM_DEF,
							_this._quoteField(modelSchema), 
							_this._quoteField(modelTableName) );
		}else{
			sqlQuery = fmt.Sprintf(" DELETE FROM %s.%s ", 
							_this._quoteField(modelSchema), 
							_this._quoteField(modelTableName) );
		}
	}else
	if( ctx.Dialect == ESqlDialect.MsSql){

		sqlQuery = fmt.Sprintf(" DELETE FROM %s.%s ", 
						_this._quoteField(modelSchema), 
						_this._quoteField(modelTableName) );
	}
	
	if(_this.getJoins() != nil && _this.getJoins()._joins != nil ){

		for _, join := range( _this.getJoins()._joins ) {
			sqlQuery += join.getSqlTxt("");
		}
	}
	
	if(_this.whereTxt != "" ){
		sqlQuery += " WHERE " +_this.whereTxt;
	}else
	{   //not allowed to delete without WHERE. IT could be a bug. that erase entire table
		return "";
	}
	
	return sqlQuery;
}


func (_this *DBQuery[T]) _getCount( /*#String*/ fldName string) string {
	
	var sqlQuery = _this._generateSelectSql( fmt.Sprintf( "COUNT(*) AS %s", fldName), SQL_ITEM_DEF, true, nil);
	
	return sqlQuery;
}

/*#PHPARG=[ String];*/
func (_this *DBQuery[T]) _getDistinctCount( fldName string, fields []string  ) string {
	
	var table = _this._generateSqlSourceOfData();
	var ITEM = SQL_ITEM_DEF;
	var sqlQuery = "";
	var bSelectDD = false;

	if( fields != nil ){
		var fieldQList = "";
		for _, field := range(fields) {

			var fieldQ  = _this._quoteTableField( field, false, _this.getJoins() );
			fieldQList += fieldQ;
		}   
		bSelectDD = true;
		sqlQuery = fmt.Sprintf( "SELECT COUNT(*) AS %s FROM ( SELECT DISTINCT %s FROM %s %s ",
										fldName, fieldQList, table, ITEM );
	} else {
		bSelectDD = true;
		sqlQuery = fmt.Sprintf( "SELECT COUNT(*) AS %s FROM ( SELECT DISTINCT * FROM %s %s ",
										fldName, table, ITEM );
	}
	
	if( _this.getJoins() != nil && _this.getJoins()._joins != nil){

		for _, join:= range(_this.getJoins()._joins) {
			sqlQuery += join.getSqlTxt(ITEM);
		}
	}
	if( _this.whereTxt != ""){
		sqlQuery += " WHERE " + _this.whereTxt;
	}
	if( _this.limit != ""){
		sqlQuery += " " + _this.limit;
	}
	if( bSelectDD ){
		sqlQuery += " ) ctxi";    
	}
	return sqlQuery;
}


func (_this *DBQuery[T]) collectForeignKeysID( arrForeignKeys []string, structDef TDefIncludeRelation){

	var reflectVal = structDef.ValueDef
	//var reflectType = reflect.TypeOf( structDef )
	var numOfFields = reflectVal.NumField()

	//var table = '';

	for iFld := 0; iFld < numOfFields; iFld ++ {

		var fld = reflectVal.Field( iFld )
		if( fld.Type().Kind() == reflect.Struct ){
			continue;
		}else{

			Arr_Append( &arrForeignKeys, fld.String() )
		}
	}
}
	


func (_this *DBQuery[T]) InFO( field string, operandsIn []any, 
		bNOT bool, bSpaceOrNull bool, fieldFK string) *DBSqlQuery[T]{
	 //bSpaceOrNull - este legat de faptul ca se raporteaza null ca si '' in getFilter()
	var txtFieldFK = "";

	var ctx = _this.tableInst.m_ctx
	
	if( len(operandsIn) == 0 ){

		var where1 	= fmt.Sprintf( "%s = %s", ctx.LangDB.VALUE_FALSE, ctx.LangDB.VALUE_TRUE)
		var ret 	= (new (DBSqlQuery[T])).Constr( where1 );
		ret.m_field1          = field;
		ret.m_listOperandsStr = operandsIn;
		if( bNOT ){

			ret.text 	= fmt.Sprintf( "%s = %s", ctx.LangDB.VALUE_TRUE, ctx.LangDB.VALUE_TRUE)
			ret.m_op = "="
		}else{
			ret.m_op = "=";
		}
		return ret
	}else{

		//table = _this.tableName;
		//_this.getJoins() = new DBSqlJoinCollection();
		var op1   = _this._quoteTableField( field, false, _this.getJoins() ); 

		var where1 = ""
		if( bNOT ){
			where1  = fmt.Sprintf( " %s NOT IN (", op1)
		}else{
			where1  = fmt.Sprintf( " %s IN (", op1);
		}
		
		var bFirst = true;
		if( operandsIn != nil){
			for _, opIn := range(operandsIn){
			
				var op = _this._quote( opIn, nil); 
				if( bSpaceOrNull && (opIn == "" ) ){   // sa accepte si null. field in ("", null) - nu merge, pt ca null e tratat special
					var fieldFKSql =  _this._quoteTableField( fieldFK, false, _this.getJoins() ); 
					txtFieldFK = fmt.Sprintf( "%s is null", fieldFKSql);
				}

				if( bFirst ){
					where1 +=  op;
				}else{
					where1 += ", " + op;
				}
				bFirst = false;
			}  
		}else{
			where1 += "''";//I need to put x in ( '' )  
		}
		where1 += " )";

		if( txtFieldFK != "" && !bNOT ){
			where1 = fmt.Sprintf("( %s OR %s )", where1, txtFieldFK);
		}
		
		var ret = (new (DBSqlQuery[T])).Constr( where1 );
		ret.m_field1          = field;
		ret.m_listOperandsStr = operandsIn;
		if( bNOT ){

			ret.m_op = "NOT_IN"
		}else{
			ret.m_op = "IN";
		}
		return ret;
	}
}




func convertToTemplateT[T IGeneric_MODEL] ( models[]any) []*T {

	var arr = []*T{}

	for i := 0; i < len(models); i++{

		Arr_Append( &arr, models[i].(*T) )
	}
	return arr;
}

/**
 convert 
*/
func (_this *DBQuery[T]) _getModelRelations(
	includeRelDefs []*TDefIncludeRelation, fnNewModel func()any) ([]any, error) {

	var arrInstModel = []any{}

	sqlQuery := _this._getRows(false, nil, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()		
	dbResult, err := _this.tableInst.m_ctx.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if dbResult != nil && err == nil {
	
		_this.clearCachedSyntax()
		if( fnNewModel == nil ){

			arrInstModel, err = _this._arrayRecordsAny(dbResult, func()any{return new (T)})
			if( err != nil){
				return nil, err
			}
		}else{
			arrInstModel, err = _this._arrayRecordsAny(dbResult, fnNewModel)
			if( err != nil){
				return nil, err
			}
		}		
	}else{
		
		_this.checkMySqlError(sqlQuery, err)
		return nil, err
	}

	var dictIncludedRel = make( map[string] *[]*TDefIncludeRelation )

	//compute for all includeRelDefs
	for i:= 0; i< len(includeRelDefs); i++ {

		var includeRelDef= includeRelDefs[i]
		//var currTable = ctx.SCHEMA_SQL_BySqlName[ _this.tableInst.m_sqlName ]
		if( includeRelDef.ValueDef.NumField() <= 1){
			continue 
		}
		var reflectFld = includeRelDef.ValueDef.Field(1)// primul e Generic_Def
		var valCol     = reflectFld.String()

		//first set the FK, PathFull, 
		_this._getModelRelations_setFKRel(includeRelDef, valCol, nil);
		//after I have FK . set the dictionary
		_this._getModelRelations_setFKRel(includeRelDef, valCol, &dictIncludedRel); 
	}

	// structDefs order by rank
	Sort.Slice( includeRelDefs, func(i, j int) bool {
				return includeRelDefs[i].RankFK < includeRelDefs[j].RankFK
			})
	
	//var ctx = _this.tableInst.m_ctx

	//var arrFKeys = []string{}
	for _, includeRelDef := range(includeRelDefs) {

		var ids = []any{}
		
		var fkTag = includeRelDef.KeyFK
		var fk, hasFk = ctx.FOREIGN_KEYS[ fkTag ]
		if( !hasFk ){
			continue
		}

		
		if( fk.RootTable_sqlName != _this.tableName){
			continue; //I will do only for includeRelDef that have my table
		}
		var table = fk.TgtTable_sqlName;
		var table_ID = fk.TgtFldID_sqlName;
		
		var langTableRelation_ID  = fk.RootFldFk_langName
		var lang2TableRelation_ID = fk.RootFldFk_lang2Name

		for _, elem := range(arrInstModel){
		
			var id = _this._getFK_IDs(elem, lang2TableRelation_ID)
			if( id != NULL_FK_ID ){
				Arr_Append( &ids, any(id) )
			}
		}
		///_this.collectForeignKeysID( arrFKeys, structDef)

		newQuery, err := _this._changeTable(table);
		if( err != nil){
			return nil, err
		}
		var includeRelDefListFinal = [] *TDefIncludeRelation{}
		var includeRelDefList, has = dictIncludedRel[includeRelDef.PathFK]
		if( !has ){
			
			includeRelDefList = nil
		}else{
			//copy data here
			includeRelDefListFinal = *includeRelDefList
			/*
			for iLst =0; iLst < len(includeRelDefList); iLst++{

				var new = new (m.TDefIncludeRelation)
				//copy entire data
				*new = *includeRelDefList[iLst] 
				new.RankFK -= 1;
			}*/
		}

		modelsLvl1, err := newQuery.WhereIn( table_ID, ids ).
						_getModelRelations( includeRelDefListFinal, includeRelDef.FnNewInst );
		if( err != nil){
			return nil, err
		}
		_this._replaceFKeyValues( langTableRelation_ID, lang2TableRelation_ID, arrInstModel, 
								table_ID, modelsLvl1);
	}

	_this.tableInst.m_ctx.updateDeltaTime();

	return arrInstModel, err;
}

func (_this *DBQuery[T]) _getModelRelations_setFKRel(
	includeRelDef *TDefIncludeRelation, valCol string, 
	dictIncludedRel* map[string] *[]*TDefIncludeRelation ){

	//valCol = userRole_ID.ID
	var valColParts = Str.Split( valCol, CONCAT_FIELDS)
	//fkTag = user.userRole_ID

	
	var fkField = ""
	var fkTag = _this.tableName 
	for j := 0; j < len(valColParts)-1; j++{

		fkTag  += CONCAT_FIELDS + valColParts[j];		
		fkField = valColParts[j];

		if( dictIncludedRel != nil && fkTag != includeRelDef.PathFK){
			// add in this map, all relations
			var elList, has = (*dictIncludedRel)[fkTag]
			if( !has ){
				elList = new([]*TDefIncludeRelation)
				(*dictIncludedRel)[fkTag] = elList
			}
			Arr_Append( elList, includeRelDef)
		}
	}
	
	includeRelDef.PathFK = fkTag
	includeRelDef.KeyFK = _this.tableName + CONCAT_FIELDS + fkField
	includeRelDef.RankFK = len(valColParts)
}

func (_this *DBQuery[T]) _getFK_IDs(elem any, lang2TableRelation_ID string )int64{

	var reflectVal = reflect.ValueOf( elem ).Elem()
	if( reflectVal.CanAddr()){

		var fld = reflectVal.FieldByName( lang2TableRelation_ID )
		if( fld.IsValid() ){

			var fldNullable = _this._getNullableField( fld);
			if( fldNullable != nil ){
				return fldNullable.Value.Int()
			}

			return fld.Int();
		}else{			
			debugger()
		}
	}
	return NULL_FK_ID
}

func (_this *DBQuery[T])  _replaceFKeyValues( 
	langModelTableRelation_ID string, 
	lang2ModelTableRelation_ID string, 
	models []any,

	langRelTable_ID string,
	relations[]any ){

	var dictRelations = make(map[int64]any)

	for i := 0; i < len(relations); i++{
		var valID = reflect.ValueOf( relations[i] ).Elem().FieldByName(langRelTable_ID).Int()

		dictRelations[ valID ] = relations[i]
	}

	var _models = (models);
	for i := 0; i < len(_models); i++{
		
		//var reflectValue =  reflect.ValueOf( _models[i] ).Elem();
		//var fkID = reflectValue.FieldByName(lang2ModelTableRelation_ID).Int()
		var fkID = _this._getFK_IDs( _models[i], lang2ModelTableRelation_ID  );
		if( fkID != NULL_FK_ID){

			var ptr, has = dictRelations[ fkID ];
			if( has ){
				var reflectVal = reflect.ValueOf(ptr)
				var reflectValue =  reflect.ValueOf( _models[i] ).Elem();
				reflectValue.FieldByName(langModelTableRelation_ID).Set(reflectVal)
			}
		}
	}
}


func (_this *DBQuery[T]) _changeTable(tableName string) (*DBQuery[IGeneric_MODEL], error) {

	
	var table, has = _this.tableInst.m_ctx.AllTables[ tableName ]
	if( has ){
		var query = (new (DBQuery[IGeneric_MODEL])).Constr(table)

		return query, nil
	}
	return nil, fmt.Errorf("table %s not found", tableName)
}


func (_this *DBQuery[T]) _updateBulkRecords(  records *[]*T, fields*[]string) error {
        
	var ctx            = _this.tableInst.m_ctx;
	var modelTableName = _this.tableName;
	
	var tableDef, _ = ctx.SCHEMA_SQL[ _this.tableNameOrig ];
	var primaryKeySql = tableDef.PrimaryColumnSqlName
	var primaryKeyLang  = tableDef.PrimaryColumnLangName
	var dictFieldsSchema = tableDef.getDictColumnByLangName()
	
	var sqlQuery = "";

	for _, modelData:= range( *records )  {
	
		var bValFirst = true;
		var sqlFieldName_value = "";

		for langFieldName, columnDef:= range( *dictFieldsSchema ){
		
			if( fields != nil && Arr_Contains( fields, columnDef.SqlName) ){
				continue
			}
			if( columnDef.ForeignKeyLangName != nil){

				//cand fac update la un model, setez model.FK_ID = model.FKID.ID
				
				var reflectModel   = reflect.ValueOf( modelData ).Elem();
				//var colTable   = ctx.SCHEMA_SQL[];

				var retDataGetID = _this.getModel_ForeignKey_getID( ctx, &columnDef, reflectModel )
				if( retDataGetID == nil){
					continue;
				}
				var ptr_ID = retDataGetID.ID
				if( ptr_ID == 0){
					//asta inseamna ca relatia a fost inserata noua

					//obtin ptr la User.UserRoleID*
					var lang1FieldName = columnDef.ForeignKeyLangName[0];
					var langTableName = columnDef.ForeignKeyLangName[1];

					var reflFldPtr = reflectModel.FieldByName( lang1FieldName ).Elem()
					if( reflFldPtr.CanAddr() && reflFldPtr.IsValid() ){

						var err error 
						ptr_ID, err = _this._InsertRecordByReflectValue( langTableName, reflFldPtr, false, nil) 
						if( ptr_ID == 0 || err != nil){
							return err
						}
					}
				}

				var fld = reflectModel.FieldByName( columnDef.LangName )

				var fldNullable = _this._getNullableField( fld );
				if( fldNullable != nil && fldNullable.Valid.Bool() ){
					fld = fldNullable.Value

					fldNullable.Valid.SetBool( ptr_ID != 0  )							
				}

				var fieldInfoType = fld.Type()

				if  fieldInfoType == reflect.TypeOf((*int16)(nil)).Elem() ||
					fieldInfoType == reflect.TypeOf((*int32)(nil)).Elem() ||
					fieldInfoType == reflect.TypeOf((*int64)(nil)).Elem() ||
					fieldInfoType == reflect.TypeOf((*int)(nil)).Elem() ||
					fieldInfoType == reflect.TypeOf((*int8)(nil)).Elem() {		
						fld.SetInt( ptr_ID )		
				} 
			}
			var value = _this.getModel_FieldValue( modelData, langFieldName, &columnDef )
			if( ( (primaryKeyLang == langFieldName)) ){ 
				//do not insert in udate cmd the primary key
				continue;
			}
			if( !bValFirst){
				sqlFieldName_value += ", ";
			}

			var sqlVal = fmt.Sprintf( "%s = %s ", _this._quoteField(columnDef.SqlName), _this._quote( value, &columnDef) );
			sqlFieldName_value += sqlVal;
			bValFirst = false;
		}
		var ID = reflect.ValueOf( modelData).Elem().FieldByName(primaryKeyLang).Int();

		sqlQuery += fmt.Sprintf( 
			`UPDATE %s SET %s WHERE %s = %d %s`, _this._quoteTable(modelTableName), 
					sqlFieldName_value, _this._quoteField(primaryKeySql), ID, ctx.LangDB.END_COMMAND );
	}		
	
	ctx.currOperationDTime2 = time.Now()			
	dbResult, err := _this.tableInst.m_ctx.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()	

	if dbResult != nil && err == nil {

		_this.clearCachedSyntax()
		//return _this._arrayOfSingleField(dbResult, field)
		return nil
	}		
	
	_this.checkMySqlError( sqlQuery, err);
		
	return err;
}

func (_this *DBQuery[T]) getModel_FieldValue(model *T, fieldName string, columnTable *TSqlColumnDef) any{

	var reflectVal = reflect.ValueOf( model).Elem()
	if( reflectVal.IsValid() ){

		var fld = reflectVal.FieldByName( fieldName );

		var fldNullable = _this._getNullableField( fld);
		if( fldNullable != nil && fldNullable.Valid.Bool() ){
			fld = fldNullable.Value
		}

		var fieldInfoType = fld.Type()

		if  fieldInfoType == reflect.TypeOf((*int16)(nil)).Elem() ||
	   		fieldInfoType == reflect.TypeOf((*int32)(nil)).Elem() ||
	   		fieldInfoType == reflect.TypeOf((*int64)(nil)).Elem() ||
	   		fieldInfoType == reflect.TypeOf((*int)(nil)).Elem() ||
	   		fieldInfoType == reflect.TypeOf((*int8)(nil)).Elem() {		
				return fld.Int()		
		} else 
		if fieldInfoType == reflect.TypeOf((*bool)(nil)).Elem() {

			return fld.Bool()			
		} else 
		if  fieldInfoType == reflect.TypeOf((*float32)(nil)).Elem() ||
			fieldInfoType == reflect.TypeOf((*float64)(nil)).Elem() {
				return fld.Float()
		} else 
		if fieldInfoType == reflect.TypeOf((*string)(nil)).Elem() {
			return fld.String()
		} else 
		if fieldInfoType == reflect.TypeOf((*time.Time)(nil)).Elem() {
			return fld.String()
		}else
		if fieldInfoType == reflect.TypeOf((*[]uint8)(nil)).Elem() {

			var slice = fld.Bytes()
			
			return slice
			/*
			var typeSqlElement = Str.ReplaceAll( columnTable.sqlType,"[]", "" )
				
			var valSql = fmt.Sprintf( "ARRAY[]::%s[]", typeSqlElement)
			return valSql;
			*/
		}
		
	}
	return nil
}

type TFuncStatic struct {
	varName string
	varValue any
}
func toany( val *int, typeVar string ) (int, any){

	var sizeInt = unsafe.Sizeof(val)
	var cellInt = 1;//consider 64 bits
	if( sizeInt == 4){
		cellInt = 2
	}
	if( sizeInt == 8){
		cellInt = 1
	}

	switch( typeVar ){

		case "int16": 	return 1, *((*int16)(( unsafe.Pointer(val) )) );
		case "int32": 	return 1, *((*int32)(( unsafe.Pointer(val) )) );
		case "int": 	return 1, *val;
		case "int64": 	return cellInt, *((*int64)(( unsafe.Pointer(val) )) );
		case "float32": return 1, *((*float32)(( unsafe.Pointer(val) )) );
		case "float64":	return cellInt, *((*float64)(( unsafe.Pointer(val) )) );
		case "bool": 	return 1, *((*bool)(( unsafe.Pointer(val) )) );
		case "string": {

			var tt string = "";
			var sizeOfString = unsafe.Sizeof(tt)
			return int(math.Round(float64(sizeOfString)/float64(sizeInt)) ), 
					*((*string)(( unsafe.Pointer(val) )) );
		}
		case "time": 	{

			var tt time.Time;
			var sizeOfTime = unsafe.Sizeof(tt)
			return int(math.Round(float64(sizeOfTime)/float64(sizeInt)) ), 
					*((*time.Time)(( unsafe.Pointer(val) )) );
		}
	}
	return 1, nil	
}
func  (_this *DBQuery[T])  _extractStaticVarFromFunc(  
	ptr_fnWhere unsafe.Pointer, 
	externalVarsSignature []TExternVar ) map[string]any{

	var dictVar = map[string]any{}

	type TT1 struct {
		f  *uintptr
		i1  int
		i2  int
		i3  int
		i4  int
		i5  int
		i6  int
		i7  int
		i8  int
		i9  int
		i10 int
		i11 int
		i12 int
	}

	type TT struct {
		f *TT1
	}
	type TTP *TT

	var ptr = TTP(ptr_fnWhere)
	var off = 0;

	for i := 0; i < len( externalVarsSignature ); i++ {

		var varType = externalVarsSignature[i].VarType;
		var idx = 0
		var val any = nil

		switch( off ){
			case 0:{
				idx, val = toany(&ptr.f.i1, varType);
			}		
			case 1:{
				idx, val = toany(&ptr.f.i2, varType);
			}
			case 2:{
				idx, val = toany(&ptr.f.i3, varType);
			}
			case 3:{
				idx, val = toany(&ptr.f.i4, varType);
			}
			case 4:{
				idx, val = toany(&ptr.f.i5, varType);
			}
			case 5:{
				idx, val = toany(&ptr.f.i6, varType);
			}
			case 6:{
				idx, val = toany(&ptr.f.i7, varType);
			}
			case 7:{
				idx, val = toany(&ptr.f.i8, varType);
			}
			case 8:{
				idx, val = toany(&ptr.f.i9, varType);
			}
			case 9:{
				idx, val = toany(&ptr.f.i10, varType);
			}
			case 10:{
				idx, val = toany(&ptr.f.i11, varType);
			}
			case 11:{
				idx, val = toany(&ptr.f.i12, varType);
			}
		}

		off += idx
		//arrays.Append( &arr, val )
		dictVar[externalVarsSignature[i].VarName ] = val
	}

	return dictVar;
}

func (_this *DBQuery[T]) _whereEq(field string, operand any, field2 string, bNot bool) *DBQuery[T] {

	var queryVal *DBSqlQuery[T] = nil

	if field2 == "" {
		queryVal = _this.eqFO(field, operand, bNot)
	} else {
		queryVal = _this.eqFF(field, field2, bNot)
	}

	if _this.whereTxt != "" {
		_this.whereTxt += " AND "
	}
	_this.whereTxt += queryVal.getText()

	if _this.m_queryAND == nil {

		_this.m_queryAND = (new(DBSqlQuery[T])).Constr("") //"(opText1) AND (opText2)" );
		_this.m_queryAND.m_op = "AND"
		_this.m_queryAND.m_listOperands = []*DBSqlQuery[T]{}
	}
	array_push(&_this.m_queryAND.m_listOperands, queryVal)
	return _this
}	

func (_this *DBQuery[T])  getSqlNativeMethod( compiledQry TCompiledSqlQuery, ptr_fnWhere unsafe.Pointer, excludedLangFields []string ) (string, []string){

	var statics = _this._extractStaticVarFromFunc( ptr_fnWhere, compiledQry.ExternVar )
	var fields     = compiledQry.Fields;
	var sql        = compiledQry.CompiledQuery;
	
	for staticKey, staticVal := range(statics) {

		//if( gettype(staticVal) == "object" ){
		//	continue;
		//}
		
		if( staticVal == true ){

			sql = str_replace( "{@@"+staticKey+"@@}", "(1=1)", sql );
		}else
		if( staticVal == false ){

			sql = str_replace( "{@@"+staticKey+"@@}", "(1=0)", sql );
		}else{

			var staticVal1 = _this._quote( staticVal, nil );
			sql = str_replace( "{@@"+staticKey+"@@}", staticVal1, sql );
		}
	}
	
	var PREFIX_FIELDS_len  = len( PREFIX_FIELDS);
	var POSTFIX_FIELDS_len = len( POSTFIX_FIELDS);

	var selectAggregatedFields  = map[string]string  {}// = arrays.CloneMapString( &compiledQry.SelectSqlFields);
	

	for fldLangName, field_ := range(fields) {

		var field		= field_;
		var len_field   = len(field);
		var fldClean   	= substr( field, PREFIX_FIELDS_len, len_field - POSTFIX_FIELDS_len );
		var itm        	= _this._quoteTableField( fldClean, false, _this.getJoins() );

		if( compiledQry.SelectSqlFields != nil ){
			
			for _, fieldExpr := range(compiledQry.SelectSqlFields) {

				if( !Str.Contains( fieldExpr, field) ){
					continue;
				}
				var fieldExpr1  = str_replace( field, itm, fieldExpr );
				if( fieldExpr1 != fieldExpr){
					
					if( !_existInListString( excludedLangFields, fldLangName) ){
						selectAggregatedFields[fldLangName] = fieldExpr1;
					}
				}
			}
		}
		sql = str_replace( field, itm, sql );
	}
	if( selectAggregatedFields != nil && len(selectAggregatedFields) > 0 ){
		return sql, Util_FromMapToArray( &selectAggregatedFields );
	}
	return sql, nil
}


func (_this *DBQuery[T])  _whereGeneric(  fnWhere func(x *T) bool ) *DBSqlQuery[T]{

	var ctx = _this.tableInst.m_ctx
	//foreach( SQL_WHERE_QUERIES as file =>sqlQueries )
	var sqlQueries = ctx.CompiledSqlQueries;
	
	var fullTag =  _this.myTag + "-" + _this.subTag;
	var query, hasQuery = sqlQueries[ fullTag ];
	if( hasQuery ){

		var sql,_ = _this.getSqlNativeMethod( query, unsafe.Pointer(&fnWhere), nil );
		
		var ret = (new(DBSqlQuery[T])).Constr( sql );//"(opText1) AND (opText2)" );
		ret.fnWhere = fnWhere;
		
		if( _this.m_queryAND == nil ){

			_this.m_queryAND      = (new (DBSqlQuery[T])).Constr( "" );//"(opText1) AND (opText2)" );
			_this.m_queryAND.m_op = "AND";
			_this.m_queryAND.m_listOperands   = []*DBSqlQuery[T]{}
		}
		
		array_push( &_this.m_queryAND.m_listOperands, ret );
		
		return ret;
	}
	log.Printf("DBQuery::where() not found signature, tag: %s! Recompile the project, to regenerate schema", fullTag)
	//UtilLog::errorMsg("DBSqlProvider::where() not found signature! Recompile the project, to regenerate DBSchemaAdapter_MySqlProc.gen.php");
	return nil;
}


func _Select_query[T IGeneric_MODEL, V IGeneric_MODEL]( _this *DBQuery[T], fnSelect func(x *T) *V ) *DBQuery[V]{

	var ctx = _this.tableInst.m_ctx
	//foreach( SQL_WHERE_QUERIES as file =>sqlQueries )
	var sqlQueries = ctx.CompiledSqlQueries;
	
	var fullTag =  _this.myTag + "-" + _this.subTag;
	var compiledDataQuery, hasQuery = sqlQueries[ fullTag ];
	if( hasQuery ){

		//var safe_SQL_ITEM_DEF= _this.m_SQL_ITEM_DEF;
		//_this.m_SQL_ITEM_DEF = ctx.newSQL_ITEM(SQL_ITEM_DEF_SQ);

		//var oldJoins = _this.joins;
		//var newQuery = (new(DBSqlQuery[T])).Constr( sql );

		var tbl1 = (new(DBTable[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)
	
		var query = (new(DBQuery[V])).Constr(tbl1);
		query.myTag = _this.myTag;
		query.parentQuery = _this;//.cloneQuery_GenModel();
		
		//query.lamdaSelectNewRecord = _this.m_SQL_ITEM_DEF;
		query.excludeLangFieldsFromGroupBy = _this.excludeLangFieldsFromGroupBy;
		_this.excludeLangFieldsFromGroupBy = nil;//move in SELECT , the groupping part
		query.newJoinCollection();
		query.m_SQL_ITEM_DEF = ctx.newSQL_ITEM(SQL_ITEM_DEF_SQ);

		var sql, _selectSqlFields= query.getSqlNativeMethod( compiledDataQuery, unsafe.Pointer(&fnSelect), query.excludeLangFieldsFromGroupBy );
		query.querySelectNewRecord_Text = sql;
		query.querySelectNewRecord_isAgregator =false;
		//query.joins = _this.joins;
		//_this.joins = oldJoins;

		if( query.excludeLangFieldsFromGroupBy != nil){

			query.selectSqlFields = _selectSqlFields
			//query.selectSqlFields = _this._Select_getSqlFields1( compiledDataQuery.SelectSqlFields, query.excludeSqlFieldFromGroupBy );
		}
		//query.tablePhpModelName    = tablePhpModelName;
		
		//_this.m_SQL_ITEM_DEF = safe_SQL_ITEM_DEF;
		
		return query;
	}
	log.Printf("DBQuery::select() not found signature, tag: %s! Recompile the project, to regenerate schema", fullTag)
	return nil;
}