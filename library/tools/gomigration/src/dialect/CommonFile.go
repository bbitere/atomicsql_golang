package gomigration_dialect

// EOuputLang defines the constant GoLang
const EOuputLang string = "go"

type TELangSql struct {
	PostgresSql string
	MySql       string
	MsSql       string
	Oracle      string
}
var ELangSql = TELangSql {
	PostgresSql : "postgres",
	MySql       : "mysql",
	MsSql       : "mssql",
	Oracle      : "oracle",
}

// DbColumn represents a database column
type DbColumn struct {
	LangName     string
	LangName2    string
	LangType     string
	SqlName      string
	IsIdentity   bool
	IsNullable   bool
	SqlType      string
	ForeignKey   *DbTable
}

// InitSql initializes DbColumn with SQL information
func (c *DbColumn) InitSql(sqlName, sqlType string, isNullable bool) *DbColumn {
	c.SqlName = sqlName
	c.SqlType = sqlType
	c.IsIdentity = false
	c.IsNullable = isNullable
	return c
}

// InitSqlPrimary initializes DbColumn for a primary key
func (c *DbColumn) InitSqlPrimary(sqlName string) *DbColumn {
	c.SqlName = sqlName
	c.SqlType = ""
	c.IsIdentity = true
	c.IsNullable = false
	return c
}

// InitLangSql initializes DbColumn with both language and SQL information
func (c *DbColumn) InitLangSql(langName, langName2, langType, sqlName, 
    sqlType string, isNullable, isIdentity bool, foreignKey *DbTable) *DbColumn {
        
	c.LangName = langName
	c.LangName2 = langName2
	c.LangType = langType
	c.IsIdentity = isIdentity

	c.SqlName = sqlName
	c.SqlType = sqlType
	c.IsNullable = isNullable
	c.ForeignKey = foreignKey

	return c
}

// DbTable represents a database table
type DbTable struct {
	Schema              string
	LangTableNameModel string
	SqlTableNameModel   string
	Columns             []*DbColumn
	PrimaryColumn       *DbColumn
	Json                string
}

// GetPluralTableNameModel returns the plural form of the table name
func (t *DbTable) GetPluralTableNameModel() string {
	return t.LangTableNameModel
}

// InitSql initializes DbTable with SQL information
func (t *DbTable) InitSql(sqlName string, primaryKey *DbColumn) *DbTable {
	t.SqlTableNameModel = sqlName
	t.LangTableNameModel = sqlName
	t.PrimaryColumn = primaryKey
	return t
}

// FKRootTgt represents the root and target tables and columns for a foreign key
type FKRootTgt struct {
	TableRoot  *DbTable
	ColumnRoot *DbColumn
	TableTgt   *DbTable
}

// NewFKRootTgt creates a new instance of FKRootTgt
func NewFKRootTgt(tableRoot *DbTable, columnRoot *DbColumn, tableTgt *DbTable) *FKRootTgt {
	return &FKRootTgt{
		TableRoot:  tableRoot,
		ColumnRoot: columnRoot,
		TableTgt:   tableTgt,
	}
}
