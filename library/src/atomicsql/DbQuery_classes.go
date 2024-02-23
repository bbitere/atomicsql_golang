package atomicsql

// ------------------------------------------------------

// internal use
type DBSqlJoin struct {
	nameItem string
	sql      string //left join Table item on item.ID = item2.FK

}

/*#PHPARG=[ String ];*/
func (_this *DBSqlJoin) getSqlTxt(itm string) string {
	return _this.sql
}
func (_this *DBSqlJoin) Constr() *DBSqlJoin {
	_this.nameItem = ""
	_this.sql = ""
	return _this
}

// internal use
type DBSqlJoinCollection struct {

	/*#PHPARG=[ Array< DBSqlJoin >];*/
	_joins 			map[string]*DBSqlJoin
	_joins_ordered 	[]string
}

func (_this *DBSqlJoinCollection) Constr() *DBSqlJoinCollection {
	_this._joins = nil //make(map[string]*DBSqlJoin)
	return _this
}

func (_this *DBSqlJoinCollection) getJoinsSqlQuery(ITEM string) string {

	if( _this._joins == nil ){
		return "";
	}
	var sqlQuery = "";	
	//for _, join := range _this._joins {
	for _, joinKey := range _this._joins_ordered {
		
		var join = _this._joins[ joinKey ]

		sqlQuery += join.getSqlTxt(ITEM)
	}

	return sqlQuery;
}

type DBSqlQuery[T IGeneric_MODEL] struct {
	text string

	m_op       string
	m_field1   string
	m_field2   string
	m_operand1 string
	m_operand2 any
	/*#PHPARG=[ Delegate1< BOOL, T> ];*/
	fnWhere  func(x *T) bool
	fnWhereS func( x *T, q IDBQuery) bool

	/*#PHPARG=[ Array<DBSqlQuery> ];*/
	m_listOperands []*DBSqlQuery[T]

	/*#PHPARG=[ Array<String> ];*/
	m_listOperandsStr []any
}

func (_this *DBSqlQuery[T]) Constr( /*#String*/ text string) *DBSqlQuery[T] {

	_this.text = text
	return _this
}

func (_this *DBSqlQuery[T]) cloneSqlQuery_GenModel() *DBSqlQuery[IGeneric_MODEL] {

	var newQ = (new(DBSqlQuery[IGeneric_MODEL])).Constr(_this.text)
	newQ.text = _this.text

	newQ.m_op = _this.m_op
	newQ.m_field1 = _this.m_field1
	newQ.m_field2 = _this.m_field2
	newQ.m_operand1 = _this.m_operand1
	newQ.m_operand2 = _this.m_operand2
	newQ.m_listOperandsStr = _this.m_listOperandsStr

	//fnWhere
	//m_listOperands
	return newQ
}

/*#PHPARG=[ String ];*/
func (_this *DBSqlQuery[T]) getText() string {
	return _this.text
}

type RuntimeCollection[T IGeneric_MODEL] struct {
	structDefs       []*TDefIncludeRelation
	modelsCollection []*T
}

type IRuntimeCollection[T IGeneric_MODEL] interface {
	DeleteModels(models []*T) bool
	InsertModels(models []*T)
	InsertOrUpdateModels(models []*T)
	GetModels() []*T
}

type TModel[T IGeneric_MODEL] struct {
	model *T
	dict  map[string]string
}

type RuntimeQuery[T IGeneric_MODEL] struct {
	structDefs []*TDefIncludeRelation
	models     []*T

	// the collection where can we can insert, update, delete
	collection *IRuntimeCollection[T]
}

func (_this *RuntimeQuery[T]) Constr(models []*T, structDefs []*TDefIncludeRelation, collection *IRuntimeCollection[T]) *RuntimeQuery[T] {

	_this.models = models
	_this.structDefs = structDefs
	_this.collection = collection
	return _this
}
