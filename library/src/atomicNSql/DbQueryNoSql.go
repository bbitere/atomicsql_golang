package atomicNSql

import (
	//"reflect"
	//"unsafe"

	//sql "database/sql"
	context "context"
	fmt "fmt"
	log "log"
	reflect "reflect"
	"sort"

	atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"

	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	//primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// --------------------------------------------------------------------------------------------------------

// --------------------------------------------------------------------------------------------------------

// DBQueryNoSql is the struct that will do the magic in atomicsql.
//
// context.Table.Qry(""); ->return DBQueryNoSql
//
// DBQueryNoSql can put in a sequence a lot of another methods to set attributes, the last will be a method that return or set the data.
type DBQueryNoSql[T atomicsql.IGeneric_MODEL] struct {
	//DBSqlQueryNoSql[T]
	//rows []*T

	//tableName string;
	/*#PHPARG=[ DBSqlProvider<T> ];*/
	parentQuery atomicsql.IDBQuery //*DBQueryNoSql[T]

	/*#PHPARG=[ DBSqlQueryNoSql<T> ];*/
	//querySelectNewRecord any;//*DBSqlQueryNoSql[T]
	querySelectNewRecord_Text        string //*DBSqlQueryNoSql[T]
	querySelectNewRecord_isAgregator bool

	/*#PHPARG=[ String ];*/
	//lamdaSelectNewRecord string

	//what fields need to be excluded from GroupBy
	excludeLangFieldsFromGroupBy []string
	//what fields need to be in SELECT in clause from GroupBy
	selectSqlFields []string

	
	//withForeignKeys map[string]string //foreignkeys

	tableInst *DBTableNoSql[T] //table DBTableNoSql used in GoLang

	///*#PHPARG=[ String ];*/
	//tablePhpModelName string

	tableName            string //sqlName
	schemaTable          string
	tableNameOrig        string //langName
	tableNameOrig_nonTmp string

	fnNewInstance func(bFull bool) any // create a new model

	//clone_sqlText  string
	whereTxt       string
	//limit          string
	//orderBy        string
	last_error     string //= "";
	//m_SQL_ITEM_DEF string
	myTag          string // memorize the tag to be used with Where() and Select()
	subTag         string
	pRTM           *atomicsql.RuntimeQuery[T]

	dictSubQueryStrs     map[string]string
	parentContainerQuery atomicsql.IDBQuery
	currentSubQueryID    string

	
	//m_queryAND *DBSqlQueryNoSql[T]


	sortCriteria bson.M;
	findOptions *options.FindOptions;
	
	
	errorRet error
}


// internal user
func (_this *DBQueryNoSql[T]) Constr(tableInst *DBTableNoSql[T]) *DBQueryNoSql[T] {

	//_this.rows = []*T{}

	_this.sortCriteria = bson.M{};
	_this.findOptions = options.Find()

	_this.last_error = ""
	//_this.m_SQL_ITEM_DEF = SQL_ITEM_DEF
	_this.tableInst = tableInst
	_this.tableName = tableInst.getSqlName()
	_this.schemaTable = tableInst.getSchemaTable()
	_this.tableNameOrig = tableInst.getLangName()
	//_this.lamdaSelectNewRecord = ""

	//_this.joins = (new(DBSqlJoinCollection)).Constr()
	//_this.joins._joins = make(map[string]*DBSqlJoin)

	//_this.m_SQL_ITEM_DEF = SQL_ITEM_DEF
	//_this.m_queryAND = nil
	_this.whereTxt = ""
	//_this.limit = ""
	//_this.orderBy = ""
	//_this.withForeignKeys = nil
	_this.pRTM = nil

	return _this
}

func (_this *DBQueryNoSql[T]) SetSubQueryString(key string, sqlQuery string) {

	_this.dictSubQueryStrs[key] = sqlQuery
}

func (_this *DBQueryNoSql[T]) GetTagID() string {

	return _this.myTag
}
func (_this *DBQueryNoSql[T]) IsRTM() bool {

	return _this.pRTM != nil
}

func (_this *DBQueryNoSql[T]) cloneQuery_Empty() *DBQueryNoSql[T] {

	newQuery := (new(DBQueryNoSql[T])).Constr(_this.tableInst) //_this.tableInst.m_DBSqlProvider );

	newQuery.tableNameOrig = _this.tableInst.m_sqlName
	newQuery.tableNameOrig_nonTmp = ""
	newQuery.tableInst = _this.tableInst

	newQuery.tableName = _this.tableName
	//newQuery.getJoins()._joins = _this.getJoins()._joins

	//newQuery.clone_sqlText = _this.clone_sqlText
	//newQuery.m_queryAND = _this.m_queryAND
	//newQuery.whereTxt = _this.whereTxt
	//newQuery.limit = _this.limit
	//newQuery.orderBy = _this.orderBy
	//newQuery.withForeignKeys = _this.withForeignKeys

	newQuery.parentQuery = _this.parentQuery // ???
	//newQuery.querySelectNewRecord_Text = _this.querySelectNewRecord_Text
	//newQuery.querySelectNewRecord_isAgregator = _this.querySelectNewRecord_isAgregator

	//newQuery.lamdaSelectNewRecord = _this.lamdaSelectNewRecord;

	return newQuery
}

func (_this *DBQueryNoSql[T]) cloneQuery() *DBQueryNoSql[T] {

	newQuery := (new(DBQueryNoSql[T])).Constr(_this.tableInst) //_this.tableInst.m_DBSqlProvider );

	newQuery.tableNameOrig = _this.tableInst.m_sqlName
	newQuery.tableNameOrig_nonTmp = ""
	newQuery.tableInst = _this.tableInst

	newQuery.tableName = _this.tableName
	//newQuery.getJoins()._joins = _this.getJoins()._joins

	//newQuery.clone_sqlText = _this.clone_sqlText
	//newQuery.m_queryAND = _this.m_queryAND
	newQuery.whereTxt = _this.whereTxt

	newQuery.sortCriteria = Util_cloneMap( _this.sortCriteria );
	newQuery.findOptions = options.MergeFindOptions( newQuery.findOptions );

	//newQuery.orderBy = _this.orderBy
	//newQuery.withForeignKeys = _this.withForeignKeys

	newQuery.parentQuery = _this.parentQuery //.cloneQuery();
	newQuery.querySelectNewRecord_Text = _this.querySelectNewRecord_Text
	newQuery.querySelectNewRecord_isAgregator = _this.querySelectNewRecord_isAgregator

	//newQuery.lamdaSelectNewRecord = _this.lamdaSelectNewRecord;

	return newQuery
}

func (_this *DBQueryNoSql[T]) cloneQuery_GenModel() *DBQueryNoSql[atomicsql.IGeneric_MODEL] {

	var newTable = _this.tableInst.cloneTable_GenModel()
	var newQuery = (new(DBQueryNoSql[atomicsql.IGeneric_MODEL])).Constr(newTable) //_this.tableInst.m_DBSqlProvider );

	newQuery.tableNameOrig = _this.tableInst.m_sqlName
	newQuery.tableNameOrig_nonTmp = ""
	newQuery.tableInst = newTable

	newQuery.tableName = _this.tableName
	//newQuery.getJoins()._joins = _this.getJoins()._joins

	//if _this.m_queryAND != nil {
	//	newQuery.m_queryAND = _this.m_queryAND.cloneSqlQuery_GenModel()
	//}

	//newQuery.clone_sqlText = _this.clone_sqlText
	newQuery.whereTxt = _this.whereTxt
	newQuery.sortCriteria = Util_cloneMap( _this.sortCriteria );
	newQuery.findOptions = options.MergeFindOptions( _this.findOptions );
	
	//newQuery.withForeignKeys = _this.withForeignKeys

	newQuery.parentQuery = _this.parentQuery //.cloneQuery();
	newQuery.querySelectNewRecord_Text = _this.querySelectNewRecord_Text
	newQuery.querySelectNewRecord_isAgregator = _this.querySelectNewRecord_isAgregator
	//newQuery.lamdaSelectNewRecord = _this.lamdaSelectNewRecord;
	newQuery.excludeLangFieldsFromGroupBy = _this.excludeLangFieldsFromGroupBy

	return newQuery
}

/*
func (_this *DBQueryNoSql[T]) Select[V any](fnSelect func(x *T) *V) *DBQueryNoSql[V] {

	return _selectNewRecord[T]( _this, fnSelect );
}
//*/

// Select() - Projects each element of a sequence  into a new form.
//
// Let's look to next example. First define a local struct vUser1.
//
// If you wish to extend a model using User as a base, you can use the annotation 'atomicsql:"copy-model"' as shown in the example below.
//
//	import m "models"
//
//	type vUser1 struct {
//
//		m.User	`atomicsql:"copy-model"`  	// extends vUser1 struct with struct m.User
//		UserRole string						// add an field that is the forkeignkey UserRoleID.
//	}
//
// And after, look to the call.
// Select contain 2 parameters
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to vUser1 for each model that the sequence will return.
//
//	ex:
//	usrs4, err := orm.Select( ctx.User.Qry("evcy59").
//
//	Where(func(x *m.User) bool {
//
//		return x.UserRoleID.IsActive == true
//	}),
//
//	func (x *m.User ) *vUser1 {
//
//		return &vUser1{
//
//			User:     *x,						//return the original m.User struct
//			UserRole: x.UserRoleID.RoleName,	//add UserRole field
//		}
//
//	}).GetModels();
//
// **NOTE**: pay attention to Qry("evcy59")
//
// - "evcy59" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the Select() argument not placed outside
func SelectN[T atomicsql.IGeneric_MODEL, V atomicsql.IGeneric_MODEL](
	sequence *DBQueryNoSql[T],
	fnSelect func(x *T) *V,
) *DBQueryNoSql[V] {

	if sequence.pRTM != nil {

		var _this = sequence
		var tbl1 = (new(DBTableNoSql[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1)

		var arr = []*V{}
		for _, itm := range _this.pRTM.GetModels() {

			Arr_Append(&arr, fnSelect(itm))
		}
		query.pRTM = (new(atomicsql.RuntimeQuery[V])).Constr(arr, _this.pRTM.GetStructDefs(), nil)

		return query
	} else {

		//TODO
		//sequence.subTag = tag_Select + sequence.tableInst.m_ctx.getSubTag()
		return nil// bbitere:NOD _Select_query(sequence, fnSelect)
	}
}

// Select() - Projects each element of a sequence  into a new form.
//
// Let's look to next example. First define a local struct vUser1.
//
// If you wish to extend a model using User as a base, you can use the annotation 'atomicsql:"copy-model"' as shown in the example below.
//
//	import m "models"
//
//	type vUser1 struct {
//
//		m.User	`atomicsql:"copy-model"`  	// extends vUser1 struct with struct m.User
//		UserRole string						// add an field that is the forkeignkey UserRoleID.
//	}
//
// And after, look to the call.
// Select contain 2 parameters
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to vUser1 for each model that the sequence will return.
//
//	ex:
//	usrs4, err := orm.Select( ctx.User.Qry("evcy59").
//
//	Where(func(x *m.User) bool {
//
//		return x.UserRoleID.IsActive == true
//	}),
//
//	func (x *m.User ) *vUser1 {
//
//		return &vUser1{
//
//			User:     *x,						//return the original m.User struct
//			UserRole: x.UserRoleID.RoleName,	//add UserRole field
//		}
//
//	}).GetModels();
//
// **NOTE**: pay attention to Qry("evcy59")
//
// - "evcy59" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the Select() argument not placed outside
func SelectSubQN[T atomicsql.IGeneric_MODEL, V atomicsql.IGeneric_MODEL](
	sequence *DBQueryNoSql[T],
	fnSelect func(x *T, q atomicsql.IDBQuery) *V,
) *DBQueryNoSql[V] {

	if sequence.pRTM != nil {

		var _this = sequence
		var tbl1 = (new(DBTableNoSql[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1)

		var arr = []*V{}
		for _, itm := range _this.pRTM.GetModels() {

			Arr_Append(&arr, fnSelect(itm, _this))
		}
		query.pRTM = (new(atomicsql.RuntimeQuery[V])).Constr(arr, _this.pRTM.GetStructDefs(), nil)

		return query
	} else {

		//sequence.subTag = tag_SelectSubQ + sequence.tableInst.m_ctx.getSubTag()
		return nil// bbitere:NOD _Select_querySubQ(sequence, fnSelect)
	}
}

// Aggregate() - Applies an accumulator function over a sequence.
//
// Agregate( sequence) - receive a sequence and return an altered sequence.
//
// From this reason - Aggregate() can be used only in tandem/pairs with Select()
//
//	Select( Aggregate[ m.User, TUserAggr ]( sequence ), fnSelect );
//
// from this reason you need to look first to Select() method documentation
//
// So Let's look to next example: User struct is defined from DataBase and it is exported as struct in golang.
//
//	type User struct /*atomicsql-table:"user"*/ {
//
//		orm.Generic_MODEL
//
//		ID                  int32                         `json:"ID,omitempty"`
//		UserName            string                        `json:"userName"`
//		UUID                string                        `json:"UUID"`
//		UserPsw             string                        `json:"user_psw"`
//		UserRoleID          *UserRole                     `json:"-"`
//		UserRole_ID         sql.NullInt32                 `json:"userRole_ID"`
//		Misc1               []uint8                       `json:"misc1"`
//		Time1               sql.NullTime                  `json:"time1"`
//		Money               float64                       `json:"money"`
//
//	}
//
// First we need to declare the aggregator struct TUserAggr and the translate / select struct TUserView
//
// Lets see that in next section:
//
//	 import (
//	 m "models"
//	 orm "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
//	 )
//	 type TUserAggr struct {
//
//	 	orm.Generic_MODEL	// all struct should have the basestruct orm.Generic_MODEL
//	 	UserRoleID          *m.UserRole			//should be identic like in m.User
//	 	UserRole_ID         sql.NullInt32		//should be identic like in m.User
//
//			// remark that the fields with [] will be aggregated
//	 	Time1               []sql.NullTime	//should have the name and the basic type like in m.User
//	 	Money               []float64 		//should have the name and the basic type like in m.User
//	 }
//	 type TUserView struct{
//
//	 	orm.Generic_MODEL	// all struct should have the basestruct orm.Generic_MODEL
//	 	UserRoleName        string				// the RoleName that categorize the data list.
//	 	MinTime1           	sql.NullTime		// the result of aggregated Time1
//	 	SumMoney            float64 			// the result of sum of Money
//	 }
//
// You need to understand the main point how to creeate TUserAggr struct
// column Time, Money are the columns that will be aggreegate
// they need to have the same name and basic type like m.User struct (see the first example)
// the only difference is that they are array
//
// the another columns invovled are : UserRoleID, UserRole_ID and they should be identically like in m.User struct
//
// So, the aggregator will convert m.User structs list in TUserView list.
//
//	usrs, err := orm.Select(
//					orm.Aggregate[ m.User, TUserAggr ]( ctx.User.Qry("evcy58").
//						Where(func(x *m.User) bool {
//							return x.UserRoleID.IsActive == true
//						}),
//					),
//					func (x *TUserAggr ) *TUserView {
//						return &TUserView{
//							UserRoleName: x.UserRoleID.RoleName,	// the foreignkey x.UserRoleID.RoleName will help to categorize the elems in sequence
//							MinTime1: orm.Sql_MinDateN( x.Time1 ),	//x.Time1 is an array and will be aggregate as MinDate
//							SumMoney: orm.Sql_SumF64( x.Money ),	//x.Money is an array and will be aggregate as Sum
//						}
//					}).GetModels();
//
// **NOTE**: pay attention to Qry("evcy58")
//
// - "evcy58" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the Select() argument not placed outside
func AggregateNSql[T atomicsql.IGeneric_MODEL, V atomicsql.IGeneric_MODEL](
	_this *DBQueryNoSql[T],
) *DBQueryNoSql[V] {

	if _this.pRTM != nil {

		var tbl1 = (new(DBTableNoSql[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1)

		var arr, err2 = _Aggregate_doRuntime_NSql[T, V](_this, _this.pRTM.models)
		if err2 != nil {
			_this.checkMySqlError("Error in _Aggregate_doRuntime. Structs are not compatible", err2)
		}

		query.pRTM = (new(RuntimeQuery[V])).Constr(arr, _this.pRTM.structDefs, nil)
		return query
	} else {

		//var ctx = _this.tableInst.m_ctx
		//var safe_SQL_ITEM_DEF = _this.m_SQL_ITEM_DEF;

		var tbl1 = (new(DBTableNoSql[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1)

		query.myTag = _this.myTag
		//_this.excludeLangFieldsFromGroupBy	= excludeFromGroupBy
		query.parentQuery = _this //.cloneQuery_GenModel();
		//query.querySelectNewRecord = (new (DBSqlQueryNoSql[V])).Constr(sql);

		query.querySelectNewRecord_isAgregator = true
		//query.m_SQL_ITEM_DEF = ctx.newSQL_ITEM(SQL_ITEM_DEF_Aggr)
		//query.lamdaSelectNewRecord 	= _this.m_SQL_ITEM_DEF;

		var sql, excludeFromGroupBy = _Aggregate_generateSql_NoSql[T, V](_this)
		query.querySelectNewRecord_Text = sql
		query.excludeLangFieldsFromGroupBy = excludeFromGroupBy
		//query.tablePhpModelName    = tablePhpModelName;

		//_this.m_SQL_ITEM_DEF = safe_SQL_ITEM_DEF;
		return query
	}
}

/*
func Aggregate1[T IGeneric_MODEL, V IGeneric_MODEL](
	_this *DBQueryNoSql[T],
	fnAggregate func (x []*T) *V,
	fields ... string,
	 )  *DBQueryNoSql[V] {

	if( _this.pRTM != nil ){

		var tbl1 = (new(DBTableNoSql[V])).Constr(
			_this.tableInst.m_sqlName,
			_this.tableInst.m_langName,
			_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1);

		var arr = []*V{}
		for _, itm := range( _this.pRTM.models ) {

			Arr_Append( &arr, _Aggregate_doRuntime[T, V]( itm ) );
		}
		query.pRTM = (new (RuntimeQuery[V])).Constr( arr, _this.pRTM.structDefs )
		return query;
	}else {

		var ctx = _this.tableInst.m_ctx
		//var safe_SQL_ITEM_DEF = _this.m_SQL_ITEM_DEF;

		var tbl1 = (new(DBTableNoSql[V])).Constr(
		_this.tableInst.m_sqlName,
		_this.tableInst.m_langName,
		_this.tableInst.m_ctx)

		var query = (new(DBQueryNoSql[V])).Constr(tbl1);

		query.myTag = _this.myTag;
		//_this.excludeLangFieldsFromGroupBy	= excludeFromGroupBy
		query.parentQuery = _this;//.cloneQuery_GenModel();
		//query.querySelectNewRecord = (new (DBSqlQueryNoSql[V])).Constr(sql);

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
}
*/

// GetValueString() - Return a value from the sequence using fnSelect for first element
//
// Example:
//
//	usrRoleName, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValueString( func(x *m.User) String {return x.UserRoleID.RoleName;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `string` for first model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValueString() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValueString(fnSelect func(x *T) string) (string, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			return fnSelect(_this.pRTM.models[0]), nil
		}
		return "", nil
	} else {

		sequence.subTag = tag_SelectValue + sequence.tableInst.m_ctx.getSubTag()
		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())

		var queryValue, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if queryValue != nil && dbResult1 != nil && err == nil {

			var ret, err1 = queryValue.singleDataS(dbResult1, TGetValueModel_VALUE)
			if err1 == nil {
				return ret, nil
			}
			return "", err1
		}
		return "", err
	}
}

// GetValueInt() - Return a value from the sequence using fnSelect for first element
//
// Example:
//
//	usrStatusID, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValueInt( func(x *m.User) int64 {return x.UserRoleID.Status_ID;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `int64` for first model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValueInt() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValueInt(fnSelect func(x *T) int64) (int64, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			return fnSelect(_this.pRTM.models[0]), nil
		}
		return 0, nil
	} else {

		sequence.subTag = tag_SelectValue + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var queryValue, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if queryValue != nil && dbResult1 != nil && err == nil {

			var ret, err1 = queryValue.singleDataInt(dbResult1, TGetValueModel_VALUE)
			if err1 == nil {
				return ret, nil
			}
			return 0, err1
		}
		return 0, err
	}
}

// GetValueFloat() - Return a value from the sequence using fnSelect for first element
//
// # Select contain 2 parameters
//
// Example:
//
//	usrData, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValueFloat64( func(x *m.User) float64 {return x.UserRoleID.RoleDataFloat;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `float64` for first model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValueFloat() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValueFloat(fnSelect func(x *T) float64) (float64, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			return fnSelect(_this.pRTM.models[0]), nil
		}
		return 0, nil
	} else {

		sequence.subTag = tag_SelectValue + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var queryValue, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if queryValue != nil && dbResult1 != nil && err == nil {

			var ret, err1 = queryValue.singleDataFloat(dbResult1, TGetValueModel_VALUE)
			if err1 == nil {
				return ret, nil
			}
			return 0, err1
		}
		return 0, err
	}
}

// GetValueBool() - Return a value from the sequence using fnSelect for first element
//
// Example:
//
//	usrActive, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValueBool( func(x *m.User) bool {return x.UserRoleID.IsActive;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `bool` for first model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValueBool() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValueBool(fnSelect func(x *T) bool) (bool, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			return fnSelect(_this.pRTM.models[0]), nil
		}
		return false, nil
	} else {

		sequence.subTag = tag_SelectValue + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var queryValue, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if queryValue != nil && dbResult1 != nil && err == nil {

			var ret, err1 = queryValue.singleDataBool(dbResult1, TGetValueModel_VALUE)
			if err1 == nil {
				return ret, nil
			}
			return false, err1
		}
		return false, err
	}
}

// GetValuesString() - Return values from the sequence using fnSelect for each elements
//
// Example:
//
//	usrRoles, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValuesString( func(x *m.User) String {return x.UserRoleID.RoleName;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `string` for each model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValuesString() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValuesString(fnSelect func(x *T) string) ([]string, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			var arr = []string{}

			for i := 0; i < len(arr); i++ {
				Arr_Append(&arr, fnSelect(_this.pRTM.models[i]))
			}
			return arr, nil
		}
		return nil, nil
	} else {

		sequence.subTag = tag_SelectValues + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var query, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if query != nil && dbResult1 != nil && err == nil {

			var ret, err4 = query._arrayOfSingleFieldString(dbResult1, TGetValueModel_VALUE)
			return ret, err4
		}
		return nil, err
	}
}

// GetValuesInt() - Return values from the sequence using fnSelect for each elements
//
// Example:
//
//	usrStatusesID, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValuesInt( func(x *m.User) int64 {return x.UserRoleID.Status_ID;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `int64` for each model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValuesInt() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValuesInt(fnSelect func(x *T) int64) ([]int64, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			var arr = []int64{}

			for i := 0; i < len(arr); i++ {
				Arr_Append(&arr, fnSelect(_this.pRTM.models[i]))
			}
			return arr, nil
		}
		return nil, nil
	} else {

		sequence.subTag = tag_SelectValues + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var query, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if query != nil && dbResult1 != nil && err == nil {

			var ret, err4 = query._arrayOfSingleFieldInt(dbResult1, TGetValueModel_VALUE)
			return ret, err4
		}
		return nil, err
	}
}

// GetValuesFloat() - Return values from the sequence using fnSelect for each elements
//
// Example:
//
//	usrStatusesID, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValuesFloat( func(x *m.User) float64 {return x.UserRoleID.DataFloat;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `float64` for each model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValuesFloat() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValuesFloat(fnSelect func(x *T) float64) ([]float64, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			var arr = []float64{}

			for i := 0; i < len(arr); i++ {
				Arr_Append(&arr, fnSelect(_this.pRTM.models[i]))
			}
			return arr, nil
		}
		return nil, nil
	} else {

		sequence.subTag = tag_SelectValues + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var query, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if query != nil && dbResult1 != nil && err == nil {

			var ret, err4 = query._arrayOfSingleFieldFloat(dbResult1, TGetValueModel_VALUE)
			return ret, err4
		}
		return nil, err
	}
}

// GetValuesBool() - Return values from the sequence using fnSelect for each elements
//
// Example:
//
//	usrStatusesID, err := ctx.User.Qry("tag1").Where(...).Order(..).GetValuesBool( func(x *m.User) bool {return x.UserRoleID.DataFloat;});
//
// In the upper example, we have :
//
// 1. the sequcence. ctx.User.Qry().Where().Order().. etc
//
// 2 the literal function fnSelect that will convert from User to `bool` for each model that the sequence will return.
//
// **NOTE**: pay attention to Qry("tag1")
//
// - "tag1" it is an unique tag per application that help to retrive the sql associated code with this instruction
//
// **NOTE2**: pay attention: literal function fnSelect and sequence should be stacked in the GetValuesBool() argument not placed outside
// ...
func (_this *DBQueryNoSql[T]) GetValuesBool(fnSelect func(x *T) bool) ([]bool, error) {

	var sequence = _this
	if sequence.pRTM != nil {

		var _this = sequence
		if len(_this.pRTM.models) > 0 {

			var arr = []bool{}

			for i := 0; i < len(arr); i++ {
				Arr_Append(&arr, fnSelect(_this.pRTM.models[i]))
			}
			return arr, nil
		}
		return nil, nil
	} else {

		sequence.subTag = tag_SelectValues + sequence.tableInst.m_ctx.getSubTag()

		var dbResult1 *mongo.Cursor = nil
		defer dbResult1.Close(context.Background())
		var query, err = _SelectValue_queryNSql(sequence, fnSelect, &dbResult1)
		if query != nil && dbResult1 != nil && err == nil {

			var ret, err4 = query._arrayOfSingleFieldBool(dbResult1, TGetValueModel_VALUE)
			return ret, err4
		}
		return nil, err
	}
}

// WhereEq() is a limited filter function. the limitation is because have only 1 condition. For more conditions use Where()
//
// it can apply an additional filter, between a field and a value
//
// Ex: context.Table.WhereEq( context.User_.Name, userName). ...
//
// in this example the whereEq add a condition: User_.Name == userName
func (_this *DBQueryNoSql[T]) WhereEq(field string, operand any) *DBQueryNoSql[T] {

	//TODO
	return _this;
}

// WhereNotEq() is a limited filter function. the limitation is because have only 1 condition. For more conditions use Where()
//
// it can apply an additional filter, between a field and a value
//
//	ex:
//
//	context.Table.WhereNotEq( context.User_.Name, userName). ...
//
// in this example the whereEq add a condition: User_.Name != userName
func (_this *DBQueryNoSql[T]) WhereNotEq(field string, operand any) *DBQueryNoSql[T] {

	//TODO
	return _this
}

// WhereEqFld() is a limited filter function. the limitation is because have only 1 condition. For more conditions use Where()
//
// it can apply an additional filter, between a field and another field
//
//	ex:
//
//	context.Table.WhereEqFld( context.User_.Name, context.User_.Name1). ...
//
// in this example the WhereEqFld add a condition: User_.Name == User_.Name2
func (_this *DBQueryNoSql[T]) WhereEqFld(field string, field2 string) *DBQueryNoSql[T] {

	//TODO
	return _this
}

// WhereNotEqFld() is a limited filter function. the limitation is because have only 1 condition. For more conditions use Where()
//
// it can apply an additional filter, between a field and another field
//
//	ex:
//
//	context.Table.WhereNotEqFld( context.User_.Name, context.User_.Name1). ...
//
// in this example the WhereNotEqFld add a condition: User_.Name == User_.Name2
func (_this *DBQueryNoSql[T]) WhereNotEqFld(field string, field2 string) *DBQueryNoSql[T] {

	//TODO
	return _this
}

func (_this *DBQueryNoSql[T]) WhereIn(field string, operandsIn []any) *DBQueryNoSql[T] {

	//TODO
	return _this
}

/*#PHPARG=[ DBSqlProvider< T > ];*/
/**  @return DBSqlProvider  */
func (_this *DBQueryNoSql[T]) whereNotIn(field string, operandsIn []any) *DBQueryNoSql[T] {

	//TODO
	return _this
}

// Where() is an unlimited conditioning function. It allows you to add any desired filtering condition, including complex subconditions, even a subquery condition
//
// Ex:
//
//	context.Table.Where( func(x *Table)bool{
//
//	return Sql_IIF( x.Relation != nil, x.RelationID.Name, "") != "admin" &&
//	             (val == nil || x.Relation_ID == val)
//	}
//
// in this example the Where() add a condition :
// Let see the SQL code:
//
//	IsNull( User.RelationID.Name, "") AND (val is null OR User.Relation_ID = val)
//
// or you can add a subquery condition
//
// Let's see an example of subquery condition (***This feature is not Implemented yet.***):
// Ex:
//
//	context.Table.Qry("label").Where( func(x *Table)bool{
//
//	    var count, _ = context.Table2.Qry("").WhereEq( context.Table2_.Field1, "val2").GetCount();
//	    return Sql_IIF( x.Relation != nil, count, 0) > 5 &&
//	             (val == nil || x.Relation_ID == val)
//	}
func (_this *DBQueryNoSql[T]) Where(fnWhere func(x *T) bool) *DBQueryNoSql[T] {

	return _this._whereSubQuery(nil, fnWhere)
}

// WhereSubQ() is an unlimited conditioning function. It allows you to add any desired filtering condition, including complex subconditions, even a subquery condition but also having subqueries
//
// Ex:
// import (
//
//	atmsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"
//	atmfunc "github.com/bbitere/atomicsql_golang.git/src/atomicsql_func"
//
// )
//
//		context.Table.Qry("label1").WhereSubQ( func(x *Table, q atmsql.IDBQuery)bool{
//
//		    var ids, _ = context.Table2.QryS("ids",q).WhereEq( context.Table2_.Field1, "val2").GetRowsAsFieldInt( context.Table2_.ID );
//		    return atmfunc.Sql_ArrayContain( ids, context.Table.ForeignKey_ID);
//		}
//	 you observe that in this example inside we are using
//	 QryS("ids",q), where the label is the name of variable ids.
//
// Let see the SQL code:
//
//	WHERE User.RelationID IN ( SELECT ID FROM Table2 WHERE Field1 = 'val2' )
func (_this *DBQueryNoSql[T]) WhereSubQ(fnWhereS func(x *T, q atomicsql.IDBQuery) bool) *DBQueryNoSql[T] {

	if _this.pRTM == nil {
		//collect subquery strings
		_this.dictSubQueryStrs = make(map[string]string)
		var model = _this.generateFullModel()
		fnWhereS(model, _this)
	}

	return _this._whereSubQuery(fnWhereS, nil)
}

func (_this *DBQueryNoSql[T]) _whereSubQuery(
	fnWhereS func(x *T, q atomicsql.IDBQuery) bool,
	fnWhere func(x *T) bool) *DBQueryNoSql[T] {

	if _this.pRTM != nil {

		var arr = []*T{}
		for _, itm := range _this.pRTM.models {

			if fnWhereS != nil {
				if fnWhereS(itm, _this) {
					Arr_Append(&arr, itm)
				}
			} else {
				if fnWhere(itm) {
					Arr_Append(&arr, itm)
				}
			}
		}
		_this.pRTM.models = arr

	} else {

		//TODO
	}

	return _this
}

// Return a slice of models from sequence.
//
//	ex:
//
//	var elems = context.Table.Where().GetModels()
func (_this *DBQueryNoSql[T]) GetModels() ([]*T, error) {

	return _this.GetRecords(nil)
}

// Return a slice of records from sequence
//
//	ex:
//
//	var elems = context.Table.Where().GetRecords([]string{"ID", "Name"})
//
// ___________________________________________________
//
//	**NOTE**:the diference between Model and Record is that :
//
//	the record can have some fields that are not updated from database.
//	it is returned from methods (GetRecords, GetRecord, etc) that select only some fields that will be filled with database values, and another fields with invalid values (non selected fields).
//
// Using records should be done with precaution, because can impact the app flow
func (_this *DBQueryNoSql[T]) GetRecords(fields []string) ([]*T, error) {

	if _this.pRTM != nil {
		return _this.pRTM.models, _this.errorRet
	}
	
	var err error = nil

	_this.checkMySqlError1( err)
	return nil, err
}

// Return the first model from sequence. if sequence is empty the returned value is nil
//
//	ex:
//
//	var elem = context.Table.Where().GetFirstModel()
func (_this *DBQueryNoSql[T]) GetFirstModel() (*T, error) {

	if _this.pRTM != nil {
		if len(_this.pRTM.models) > 0 {
			return _this.pRTM.models[0], _this.errorRet
		} else {
			return nil, nil
		}
	}

	var err error = nil

	_this.checkMySqlError1( err)
	return nil, err
	
}

// Return the first record from sequence. If sequence is empty the returned value is nil
//
//	ex:
//
//	var elem = context.Table.Where().GetFirstRecord([]string{"ID", "Name"})
//
// ___________________________________________________
//
//	**NOTE**:the diference between Model and Record is that :
//
//	the record can have some fields that are not updated from database.
//	it is returned from methods (GetRecords, GetRecord, etc) that select only some fields that will be filled with database values, and another fields with invalid values (non selected fields).
//
// **NOTE**: Using records should be done with precaution, because can impact the app flow
func (_this *DBQueryNoSql[T]) GetFirstRecord(fields []string) (*T, error) {

	if _this.pRTM != nil {
		if len(_this.pRTM.models) > 0 {
			return _this.pRTM.models[0], _this.errorRet
		} else {
			return nil, nil
		}
	}

	var err error = nil

	_this.checkMySqlError1( err)
	return nil, err
	
}

/*
// Return a slice of distinct models from sequence.
//
//  ex:
//
//  var elems = context.Table.Where().GetDistinctModels( []string{ context.Table_.Field1, context.Table_.Field2 })
func (_this *DBQueryNoSql[T]) GetDistinctModels( fields []string ) ([]*T, error) {

	if( _this.pRTM != nil ){
		return _this._getDistinctRTM(fields, _this.pRTM.models), _this.errorRet;
	}

	sqlQuery := _this._getRows(true, fields, false)

	var ctx = _this.tableInst.m_ctx
	ctx.currOperationDTime2 = time.Now()
	dbResult, err := _this.tableInst.m_ctx.Query(sqlQuery)
	defer queryClose( dbResult )
	ctx.updateDeltaTime2()

	if err == nil {
		_this.clearCachedSyntax()
		return _this._arrayModels( dbResult )
	}

	_this.checkMySqlError(sqlQuery, err)
	return nil, err
}
*/

// Return a slice of distinct records from sequence. The uniqueness is provided by pair of fields arg
//
//	ex:
//
//	var elems = context.Table.Where().GetDistinctRecords([]string{"ID", "Name"})
//
// ___________________________________________________
//
//	**NOTE**:the diference between Model and Record is that :
//
//	the record can have some fields that are not updated from database.
//	it is returned from methods (GetRecords, GetRecord, etc) that select only some fields that will be filled with database values, and another fields with invalid values (non selected fields).
//
// **NOTE**: Using records should be done with precaution, because can impact the app flow
func (_this *DBQueryNoSql[T]) GetDistinctRecords(fields []string) ([]*T, error) {

	if _this.pRTM != nil {
		return _this._getDistinctRTM(fields, _this.pRTM.models), _this.errorRet
	}

	var err error = nil

	_this.checkMySqlError1( err)
	return nil, err
}

// Return the first model from sequence.
//
// this model will have set having the Relation pointer for a foreignkey.
//
// context.Table_ is a avatar of Table that contains the fields definitions
// and for the Relation (Table.RelationID) the Def() method will contains
// the definitions of import the foreignkey
//
//	 ex:
//
//	 var elem, err = context.Table.Where().GetFirstModelRel( context.Table_.RelationID.Def() )
//
//	    if(  elem != nill && elem.RelationID != nil ){
//
//				fmt.Print( elems.RelationID.Name )
//
//		}
//
// if this will missing, the returned model will have the relation pointer = nil. (Table.RelationID = nil)
//
// Please check the info
func (_this *DBQueryNoSql[T]) GetFirstModelRel(structDefs ...*TDefIncludeRelation) (*T, error) {

	if _this.pRTM != nil {
		if len(_this.pRTM.models) > 0 {
			return _this.pRTM.models[0], _this.errorRet
		} else {
			return nil, nil
		}
	}

	var err error = nil

	_this.checkMySqlError1( err)
	return nil, err
}

// Return a slice of models from sequence.
//
// this model will have set the Relation pointer to a foreignkey table.
//
// context.Table_ is a avatar of Table that contains the fields definitions
// and for the Relation (Table.RelationID) the Def() method will contains
// the definitions of import the foreignkey
//
//	 ex:
//
//	 var elems, err = context.Table.Where().GetModelsRel( context.Table_.RelationID.Def() )
//
//	    if( len(elems) > 0 && elems[0].RelationID != nil ){
//
//				fmt.Print( elems[0].RelationID.Name )
//
//		}
//
// if this will missing, the returned model will have the relation pointer = nil. (Table.RelationID = nil)
func (_this *DBQueryNoSql[T]) GetModelsRel(structDefs ...*TDefIncludeRelation) ([]*T, error) {

	if _this.pRTM != nil {
		return _this.pRTM.models, _this.errorRet
	}

	//TODO

	return nil, nil
}
func (_this *DBQueryNoSql[T]) _getModelsRel(structDefs []*TDefIncludeRelation) ([]*T, error) {

	//TODO

	return nil, nil
}

// return a single string from a specific field, (field arg is to determine this field )
//
// this is useful when we want to obtain a property value from a props table
//
// ex:
// Let consider a Table struct having 2 columns
//
//	type Table struct{
//	    propName  string  	// the name of Property
//	    propValue string	// the value of Property. here can be any type, not only string
//	}
//
//	var propString, err = context.Table.Qry("").WhereEq( context.Table_.propName, nameValue ).GetSingleDataString(context.Table_.propValue);
//
// this is echivalent with:
//
//	var modelTable, err = context.Table.Qry("").WhereEq( context.Table_.propName, nameValue ).GetFirstRecord( []String{context.Table_.propValue});
//	var propString = modelTable.propValue;
//
// So, the second case it is a bit complex, we prefer the first one.
func (_this *DBQueryNoSql[T]) GetSingleDataString(fieldName string) (string, error) {

	if _this.pRTM != nil {
		if _this.errorRet != nil {
			return "", _this.errorRet
		}
		if len(_this.pRTM.models) > 0 {
			return _this.getValueS(_this.pRTM.models[0], fieldName)
		} else {
			return "", nil
		}
	}

	//TODO
	var err error = nil
	_this.checkMySqlError1( err)
	return "", err
}

// return a single integernce from a specific field, (field arg is to determine this field )
//
// this is useful when we want to obtain a property value from a props table
//
// ex:
// Let consider a Table struct having 2 columns
//
//	type Table struct{
//	    propName  string  	// the name of Property
//	    propValue int	// the value of Property. here can be any type, not only string
//	}
//
//	var propString, err = context.Table.Qry("").WhereEq( context.Table_.propName, nameValue ).GetSingleDataI(context.Table_.propValue);
//
// this is echivalent with:
//
//	var modelTable, err = context.Table.Qry("").WhereEq( context.Table_.propName, nameValue ).GetFirstRecord( []String{context.Table_.propValue});
//	var propString = modelTable.propValue;
//
// So, the second case it is a bit complex, we prefer the first one.
func (_this *DBQueryNoSql[T]) GetSingleDataInt(sqlResult *mongo.Cursor, fieldName string) (int64, error) {

	if _this.pRTM != nil {
		if _this.errorRet != nil {
			return 0, _this.errorRet
		}
		if len(_this.pRTM.models) > 0 {
			return _this.getValueI(_this.pRTM.models[0], fieldName)
		} else {
			return 0, nil
		}
	}

	//TODO
	var err error = nil
	_this.checkMySqlError1( err)
	return 0, err
}

// return an array with data for all elements of sequence from a specific field, (field arg is to determine this field )
//
// this is useful when we want to obtain the Ids- of sequence
//
//	ex:
//
//	var ids, err = context.Table.Qry("").Where( .. ).GetRowsAsFieldString( context.Table_.UUID);
//
// this is echivalent with:
//
//	var models, err = context.Table.Qry("").Where( .. ).GetRecords( []string{context.Table_.UUID});
//	var ids = []String{}
//	for model := range models {
//	    Arr_apend( &ids, model.UUID );
//	}
func (_this *DBQueryNoSql[T]) GetRowsAsFieldString(fieldName string) ([]string, error) {

	if _this.pRTM != nil {

		if _this.errorRet != nil {
			return nil, _this.errorRet
		}
		var arr = []string{}
		for i := 0; i < len(_this.pRTM.models); i++ {
			var val, err = _this.getValueS(_this.pRTM.models[i], fieldName)
			if err == nil {
				Arr_Append(&arr, val)
			} else {
				return nil, err
			}
		}
		return arr, nil
	}
	//TODO
	var err error = nil
	_this.checkMySqlError1( err)
	return nil, err
}

// return an array with data for all elements of sequence from a specific field, (field arg is to determine this field )
//
// this is useful when we want to obtain the Ids- of sequence
//
//	ex:
//
//	var ids, err = context.Table.Qry("").Where( .. ).GetRowsAsFieldInt( context.Table_.ID);
//
// this is echivalent with:
//
//	var models, err = context.Table.Qry("").Where( .. ).GetRecords( []string{context.Table_.ID});
//	var ids = []int64{}
//	for model := range models {
//	    Arr_apend( &ids, int64(model.ID) );
//	}
func (_this *DBQueryNoSql[T]) GetRowsAsFieldInt(fieldName string) ([]int64, error) {

	if _this.pRTM != nil {

		if _this.errorRet != nil {
			return nil, _this.errorRet
		}
		var arr = []int64{}
		for i := 0; i < len(_this.pRTM.models); i++ {
			var val, err = _this.getValueI(_this.pRTM.models[i], fieldName)
			if err == nil {
				Arr_Append(&arr, val)
			} else {
				return nil, err
			}
		}
		return arr, nil
	}
	//TODO
	var err error = nil
	_this.checkMySqlError1( err)
	return nil, err
}

func (_this *DBQueryNoSql[T]) Sqlquery_GetRowsAsFieldInt(fieldName string) string {

	//TODO	
	return ""
}

// Sorts the elements of a sequence in ascending or descending order, using multiple orderFields arg.
//
// Parameters:
//
// - `orderFields`: is a dictionary that have for each field the asc or desc attribute
// Example:
// var orderFields = DataOrderByFields;
// orderFields.SetDictionary(
//
//				ctx.Table_.Field1, atomicsql.ESortField.Asc,
//				ctx.Table_.Field2, atomicsql.ESortField.Desc,
//			)
//	 var rows = ctx.Table.Qry("").OrderByFields( &orderFields ).GetModels();
//
// the rows are ordered by Table.Field1 Asc and after Table.Field2 Desc
func (_this *DBQueryNoSql[T]) OrderByFields(orderFields *atomicsql.DataOrderByFields) *DBQueryNoSql[T] {

	if _this.pRTM != nil {

		var fields = Util_FromMapKeysToArray(&orderFields.data)
		var models = _this.rtm_getModelsAsDicts(&_this.pRTM.models, fields)

		sort.Slice(models,
			func(i int, j int) bool {

				for itm := 0; itm < len(fields); itm++ {

					var fieldName = fields[itm]
					var order1 = orderFields.data[fieldName]

					var v1 = models[i].dict[fieldName]
					var v2 = models[j].dict[fieldName]
					if order1 == ESortField.Desc {

						if v1 > v2 {
							return true
						}
					} else if order1 == ESortField.Asc {

						if v1 < v2 {
							return true
						}
					}
				}
				return false
			})
		_this.pRTM.models = _this.rtm_updateModelsFromDicts(&models)
		return _this
	} else {

		if orderFields == nil || orderFields.data == nil {
			return _this
		}

		//for

		//TODO
		return _this
	}
}

// Sorts the elements of a sequence in ascending order, using field arg.
// Example:
//
//	var rows = ctx.Table.Qry("").OrderAsc(ctx.Table_.Field1).GetModels();
//
// the rows are ordered Ascendent by Table.Field1
func (_this *DBQueryNoSql[T]) OrderAsc(field string) *DBQueryNoSql[T] {

	if _this.pRTM != nil {

		var models = _this.rtm_getModelsAsDicts(&_this.pRTM.models, []string{field})
		sort.Slice(models,
			func(i int, j int) bool {
				var v1 = models[i].dict[field]
				var v2 = models[j].dict[field]
				return v1 < v2
			})
		_this.pRTM.models = _this.rtm_updateModelsFromDicts(&models)
		return _this
	} else {
		
		_this.sortCriteria[field] = 1
		_this.findOptions.SetSort( _this.sortCriteria )
	
		return _this
	}
}

// Sorts the elements of a sequence in descending order, using field arg.
// Example:
//
//	var rows = ctx.Table.Qry("").OrderDesc(ctx.Table_.Field1).GetModels();
//
// the rows are ordered Descending by Table.Field1
func (_this *DBQueryNoSql[T]) OrderDesc(field string) *DBQueryNoSql[T] {

	if _this.pRTM != nil {

		var models = _this.rtm_getModelsAsDicts(&_this.pRTM.models, []string{field})
		sort.Slice(models,
			func(i int, j int) bool {
				var v1 = models[i].dict[field]
				var v2 = models[j].dict[field]
				return v1 > v2
			})
		_this.pRTM.models = _this.rtm_updateModelsFromDicts(&models)
		return _this
	} else {

		_this.sortCriteria[field] = -1
		_this.findOptions.SetSort( _this.sortCriteria )
		return _this
	}
}

// Insert model arg 'model'
//
//	ex:
//
//	context.Table.Qry("").InsertRecord( model)
func (_this *DBQueryNoSql[T]) InsertModel(model *T) (int64, error) {

	return _this.InsertRecord(model, false, nil)
}

// Insert models or Update all data in model arg 'model'
//
//	ex:
//
//	context.Table.Qry("").InsertOrUpdateModel( record)
func (_this *DBQueryNoSql[T]) InsertOrUpdateModel(data *T) (int64, error) {

	return _this.InsertOrUpdateRecord(data, false, nil)
}

// Insert models or Update all data in model arg 'model'
//
//	ex:
//
//	context.Table.Qry("").UpdateModels( records)
//
// you can use fields to select only same fields for update. For insertion, this arg is ingnored
func (_this *DBQueryNoSql[T]) InsertOrUpdateRecord(model *T, bInsertID bool, fields *[]string) (int64, error) {

	if _this.pRTM != nil {
		if _this.pRTM.collection != nil {

			var models = []*T{model}
			(*_this.pRTM.collection).InsertOrUpdateModels(models)
			return 1, nil
		}
		return 1, nil
	}

	if model == nil {
		return 0, nil
	}

	if (*model).GetID() == 0 {

		reflectData := reflect.ValueOf(model).Elem()
		return _this._InsertRecordByReflectValue(_this.tableNameOrig, reflectData, bInsertID, fields)
	} else {

		//var arr = []*T{}
		//Arr_Append(&arr, model)

		//var err = _this._updateBulkRecords(&arr, fields)
		//return (*model).GetID(), err
		return (*model).GetID(), nil
	}
}

// Insert model arg 'model'
//
//	ex:
//
//	context.Table.Qry("").InsertRecord( record)
//
// you can use fields to select only same fields for update. For insertsion, this arg is ingnored
func (_this *DBQueryNoSql[T]) InsertRecord(data *T, bInsertID bool, fields *[]string) (int64, error) {

	if _this.pRTM != nil {
		if _this.pRTM.collection != nil {

			var models = []*T{data}
			(*_this.pRTM.collection).InsertModels(models)
			return 1, nil
		}
		return 1, nil
	}

	if data == nil {
		return 0, nil
	}
	var _, id, reflectData = _this._getNameAndID(data)
	if id != 0 {
		return id, fmt.Errorf("the model is already inserted. Detache it first")
	}
	//fldID.SetInt( 1 )

	return _this._InsertRecordByReflectValue(_this.tableNameOrig, reflectData, bInsertID, fields)
}

// Insert models or Update all data in models list arg 'models'
//
// ex: context.Table.Qry("").UpdateModels( records)
func (_this *DBQueryNoSql[T]) InsertOrUpdateModels(models []*T) error {

	return _this.InsertOrUpdateRecords(models, nil)
}

// Insert models or Update all data in models list arg 'models'
//
//	ex:
//
//	context.Table.Qry("").UpdateModels( records)
//
// you can use fields to select only same fields for update. For insertsion, this arg is ingnored
func (_this *DBQueryNoSql[T]) InsertOrUpdateRecords(models []*T, fields []string) error {

	if _this.pRTM != nil {
		Arr_AddRange(&_this.pRTM.models, &models)
		return nil
	}

	var arrIns = []*T{}
	var arrUpd = []*T{}

	for _, model := range models {

		if (*model).GetID() == 0 {

			Arr_Append(&arrIns, model)
		} else {

			Arr_Append(&arrUpd, model)
		}
	}

	if len(arrIns) > 0 {

		for _, model := range arrIns {
			var _, err = _this.InsertRecord(model, false, &fields)
			if err != nil {
				return err
			}
		}
	}
	if len(arrUpd) > 0 {

		//TODO
	}
	return nil

}

// Update all data in models list arg 'records'
//
//	ex:
//
//	context.Table.Qry("").UpdateModels( records)
func (_this *DBQueryNoSql[T]) UpdateModels(records *[]*T) error {

	//TODO
	return nil
}

// Update all data in model arg 'model'
//
//	ex:
//
//	context.Table.Qry("").UpdateModels( records)
func (_this *DBQueryNoSql[T]) UpdateModel(model *T) error {

	var arr = []*T{}
	Arr_Append(&arr, model)

	//return _this._updateBulkRecords(&arr, nil)
	//TODO
	return nil
}

// Delete all models selected in curent sequence (using Where(), WhereEq(), WhereNotEq() ) from database or collection.
//
// You can mix it with a filter condition [atomicsql.Where]() or [atomicsql.WhereEq]() or [atomicsql.WhereNotEq]()
//
//	ex:
//
//	context.Table.Qry("").WhereEq("field", "value").DeleteModels()
//
// or
//
//	context.Table.Qry("").WhereEq( context.Table_.ID, value_ID ).DeleteModels()
func (_this *DBQueryNoSql[T]) DeleteModels() error {

	if _this.pRTM != nil {

		if _this.pRTM.collection != nil {

			var ret = (*_this.pRTM.collection).DeleteModels(_this.pRTM.models)
			if !ret {
				return fmt.Errorf("elem not inserted")
			}
			return nil
		} else {

			var nameID = ""
			var arrIDs = []any{}
			for _, model := range _this.pRTM.models {

				var name_ID, id, _ = _this._getNameAndID(model)
				nameID = name_ID
				Arr_Append(&arrIDs, any(id))
			}
			return _this.WhereIn(nameID, arrIDs)._deleteModels()
		}
	} else {

		return _this._deleteModels()
	}
}

func (_this *DBQueryNoSql[T]) _deleteModels() error {

	//TODO
	var err error = nil

	_this.checkMySqlError1( err)
	return  err
}

// Delete a model from database or collection
//
//	ex:
//
//	context.Table.Qry("").DeleteModel( model )
func (_this *DBQueryNoSql[T]) DeleteModel(model *T) error {

	if _this.pRTM != nil {

		if _this.pRTM.collection != nil {

			var ret = (*_this.pRTM.collection).DeleteModels(_this.pRTM.models)
			if !ret {
				return fmt.Errorf("elem not inserted")
			}
			return nil
		} else {

			//TODO
			///var name_ID, id, _ = _this._getNameAndID(model)
			///return _this.WhereEq(name_ID, id)._deleteModels()
		}
	}

	//TODO
	var err error = nil

	_this.checkMySqlError1( err)
	return err
}

const COUNT_NAME = "Count1"

type TGetCount struct {
	atomicsql.Generic_MODEL
	Count1 int32 // the same name as COUNT_NAME
}

const TGetValueModel_VALUE = "Value1"

type TGetValueModel[V comparable] struct {
	atomicsql.Generic_MODEL
	Value1 V // the same name as TGetValueModel_VALUE
}

// return the number of elements in a sequence
//
// You can mix it with a filter condition [atomicsql.Where]() or [atomicsql.WhereEq]() or [atomicsql.WhereNotEq]()
//
//	ex:
//
//	context.Table.Qry("").WhereEq("field", "value").GetCount()
func (_this *DBQueryNoSql[T]) GetCount() (int64, error) {

	if _this.pRTM != nil {
		return int64(len(_this.pRTM.models)), nil
	}

	
	//TODO
	var err error = nil

	_this.checkMySqlError1( err)
	return 0, err
}

// return the number of distinct elements in a sequence, find the distinct elemnts using field
//
// You can mix it with a filter condition [atomicsql.Where]() or [atomicsql.WhereEq]() or [atomicsql.WhereNotEq]()
// optional you can specify the fields where to select the distinction between elems
//
//	ex:
//
//	var count, err = context.Table.Qry("").WhereEq( ctx.Table_.field, "value").GetDistinct1Count(ctx.Table_.field)
func (_this *DBQueryNoSql[T]) GetDistinct1Count(field string) (int64, error) {

	return _this.GetDistinctCount([]string{field})
}

// return the number of distinct elements in a sequence, find the distinct elemnts using fields
//
// You can mix it with a filter condition [atomicsql.Where]() or [atomicsql.WhereEq]() or [atomicsql.WhereNotEq]()
// optional you can specify the fields where to select the distinction between elems
//
//	ex:
//
//	var count, err = context.Table.Qry("").WhereEq( ctx.Table_.field, "value").GetDistinct1Count([]string{ctx.Table_.field})
func (_this *DBQueryNoSql[T]) GetDistinctCount(fields []string) (int64, error) {

	if _this.pRTM != nil {
		return int64(len(_this._getDistinctRTM(fields, _this.pRTM.models))), nil
	}

	if fields == nil || len(fields) == 0 {
		return 0, fmt.Errorf("arg fields is empty")
	}

	var err error = nil
	_this.checkMySqlError1( err)
	return 0, err
}

// ToRTM() - is a method that switch the execution of DB query in golang code, from that point forward.
//
// why this?
//
// Because sometime the speed of DB query can do switched to be test in golang code,
// and compare the speed. you can do it easily if you switch the flag "bRuntime"
//
// Because the ORM engine and scan compiler tool is still in development and might have occasional issues during compilation,
// it can be helpful to switch to runtime execution temporarily
//
// Ex
//
//	 var models = context.Table.Qry("tag1").Where( func(x *Table) bool{
//			return  x.proj_id == proj_id &&
//	 			((x.Field.Field2.FieldName == val && x.Field.Field3.FieldName2 == val2 ) ||
//	  			 IIF( x.Field.Field4.Name != nil, val4, val3) )
//	             )
//		            }).GetModels()
//
// # And Lets suppose that the scanner compile crash when it try to parse this complex syntax
//
// But, Your project must continue to run, not to be stopped, and the ORM blamed.
//
// So let see what you need to do: You can translate it as:
//
//	 var models = context.Table.Qry("tag1").WhereEq( context.Table_.proj_id, proj_id ).  // do a smart filter to reduce the rows transfered to golang app
//	 				ToRTM( RunAsRTM, context.Table_.Field.Field2.Def(), context.Table_.Field.Field3.Def(), context.Table_.Field.Field4.Def() ).  //pass the foreign keys definitions, in order to load the relations
//	 				Where( func(x *Table) bool{  // execute the where in golang code.
//							return
//	 						(x.Field.Field2.FieldName == val && x.Field.Field3.FieldName2 == val2 ) ||
//	  			 			IIF( x.Field.Field4.Name != nil, val4, val3) )
//		            })
//	    GetModels()
//
// if Flag RunAsRTM = true, the execution will be passed in golang code.
//
// After the ORM team will fix the crash of scanner compiler tool,
//
// you can switch back RunAsRTM = false, and the execution of the query will be done on DB server.
func (_this *DBQueryNoSql[T]) ToRTM(bRuntime bool, structDefs ...*atomicsql.TDefIncludeRelation) *DBQueryNoSql[T] {

	if bRuntime {

		var models, _ = _this._getModelsRel(structDefs)
		_this.pRTM = (new(atomicsql.RuntimeQuery[T])).Constr(models, structDefs, nil)
	}
	return _this
}


// SetLimit() - set the pagination of query
//
// Example:
//
//	qry1, err := ctx.User.Qry("tag1").Where(...).OrderAsc(ctx._User.UserName).SetLimit( 100, 10).GetModels();
//
// This query will take the rows from offset = 100 and with pagination = 10
// ...
func (_this *DBQueryNoSql[T]) SetLimit( queryOffset int, querySize int) *DBQueryNoSql[T] {

	//_this.setLimit( queryOffset, querySize);
	return _this
}

// CloneQry() - Clone the current query
// Why we add this?
// 
// Example:
//
//	qry1, err := ctx.User.Qry("tag1").Where(...).CloneQry();
//  var cnt,err1 = qry1.Count();
//  var rows,err2 = qry1.OrderAsc(ctx._User.UserName).SetLimit( 0, 10).GetModels();
//
// ...
func (_this *DBQueryNoSql[T]) CloneQry() (*DBQueryNoSql[T], error) {

	var newQry = _this.cloneQuery_Empty();

	//var sqlQuery = _this._generateSelectSql( "", SQL_ITEM_DEF, true, nil)
	//newQry.clone_sqlText = sqlQuery;
	newQry.findOptions = options.MergeFindOptions( _this.findOptions );

	return newQry, nil
}

//--------------------------------------------------------------------------------------------------------------
// internal
//--------------------------------------------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _getNameAndID(model *T) (string, int64, reflect.Value) {

	var reflectData = reflect.ValueOf(model).Elem()
	var nameID = _this.tableInst.m_ctx.SCHEMA_SQL_BySqlName[_this.tableName].PrimaryColumnLangName
	var fldTID = reflectData.FieldByName(nameID)
	var id = fldTID.Int()

	return nameID, id, reflectData
}


func (_this *DBQueryNoSql[T]) clearCachedSyntax() {

	_this.tableInst.m_ctx.clearCachedSyntax()
}

func (_this *DBQueryNoSql[T]) _generateSelectSql(
	selectFields string, ITEM string /*#BOOL*/, bLimit bool, select_sqlFields []string) string {
		return "";
}

/*#PHPARG=[ VOID];*/
func (_this *DBQueryNoSql[T]) checkMySqlError( /*#String*/ sqlQuery any, err error) {
	//_this.clearCachedSyntax();
	//_this.checkMySqlError1( sqlQuery);

	if err != nil {
		_this.tableInst.m_ctx.hasError = err
		log.Printf("sql query error: %s %s", sqlQuery, err.Error())
	}
	_this.clearCachedSyntax()
}
func (_this *DBQueryNoSql[T]) checkMySqlError1( err error) {
	//_this.clearCachedSyntax();
	//_this.checkMySqlError1( sqlQuery);

	if err != nil {
		_this.tableInst.m_ctx.hasError = err
		log.Printf("sql query error: %s %s", "", err.Error())
	}
	_this.clearCachedSyntax()
}

func _SelectValue_queryNSql[T atomicsql.IGeneric_MODEL, V comparable](
	_this *DBQueryNoSql[T], fnSelect func(x *T) V, refDbResult1 **mongo.Cursor) (*DBQueryNoSql[TGetValueModel[V]], error) {

	return nil, nil
}

func (_this *DBQueryNoSql[T]) singleDataS(dbResult *mongo.Cursor, fieldName string) (string, error) {

	_this.clearCachedSyntax()
	var model = new(T)
	for dbResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(dbResult, model, fieldName)
		return value.String(), err
	}
	return "", nil
	/*
		model, err := _this._oneRecord(dbResult, []string{fieldName} )

		if model != nil && err == nil{
			val := reflect.ValueOf(model).Elem().FieldByName(fieldName)
			return val.String(), nil
		}
		return "", err
	*/
}

func (_this *DBQueryNoSql[T]) singleDataInt(dbResult *mongo.Cursor, fieldName string) (int64, error) {

	_this.clearCachedSyntax()
	var model = new(T)
	for dbResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(dbResult, model, fieldName)
		return value.Int(), err
	}
	return 0, nil
}

func (_this *DBQueryNoSql[T]) singleDataFloat(dbResult *mongo.Cursor, fieldName string) (float64, error) {

	_this.clearCachedSyntax()
	var model = new(T)
	for dbResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(dbResult, model, fieldName)
		return value.Float(), err
	}
	return 0, nil
}
func (_this *DBQueryNoSql[T]) singleDataBool(dbResult *mongo.Cursor, fieldName string) (bool, error) {

	_this.clearCachedSyntax()
	var model = new(T)
	for dbResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(dbResult, model, fieldName)
		return value.Bool(), err
	}
	return false, nil
}

// ----------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _arrayOfSingleFieldString(sqlResult *mongo.Cursor, fieldName string) ([]string, error) {

	_this.clearCachedSyntax()
	var retList = []string{}
	var model = new(T)
	for sqlResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(sqlResult, model, fieldName)
		if err != nil {
			return nil, err
		}
		Arr_Append(&retList, value.String())
	}
	return retList, nil
}

// ----------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _arrayOfSingleFieldInt(sqlResult *mongo.Cursor, fieldName string) ([]int64, error) {

	_this.clearCachedSyntax()
	var retList = []int64{}
	var model = new(T)
	for sqlResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(sqlResult, model, fieldName)
		if err != nil {
			return nil, err
		}
		Arr_Append(&retList, value.Int())
	}
	return retList, nil
}

// ----------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _arrayOfSingleFieldFloat(sqlResult *mongo.Cursor, fieldName string) ([]float64, error) {

	_this.clearCachedSyntax()
	var retList = []float64{}
	var model = new(T)
	for sqlResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(sqlResult, model, fieldName)
		if err != nil {
			return nil, err
		}
		Arr_Append(&retList, value.Float())
	}
	return retList, nil
}

// ----------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _arrayOfSingleFieldBool(sqlResult *mongo.Cursor, fieldName string) ([]bool, error) {

	_this.clearCachedSyntax()
	var retList = []bool{}
	var model = new(T)
	for sqlResult.Next() {

		var value, err = readRecordSqlResult_ReadfieldValue(sqlResult, model, fieldName)
		if err != nil {
			return nil, err
		}
		Arr_Append(&retList, value.Bool())
	}
	return retList, nil
}

func (_this *DBQueryNoSql[T]) generateFullModel() *T {

	var ctx = _this.tableInst.m_ctx
	var def, has = ctx.DictTablesIncludeRelDefs[_this.tableName]
	if has {
		var model = (*def).FnNewInst(true).(*T)
		return model
	} else {
		var model *T = new(T)
		return model
	}
}

// ---------------------------------------------------------------------------------------
func (_this *DBQueryNoSql[T]) _InsertRecordByReflectValue(
	langTableName string, reflValue reflect.Value,
	bInsertID bool, fields *[]string) (int64, error) {

		return 0, nil
}

func _Aggregate_generateSql_NoSql[T atomicsql.IGeneric_MODEL, V atomicsql.IGeneric_MODEL](_this *DBQueryNoSql[T]) (string, []string) {

	return "", nil
}