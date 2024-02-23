package atomicsql

import (
	sql "database/sql"
	fmt "fmt"
	log "log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var DBConnectionInst DBContextBase

var static_db *sql.DB

func StaticOpenDB(connStr TConnectionString, maxIdle int, maxOpen int) (*sql.DB, error) {

	var dataSource = ""
	var sqlLang = ""
	if connStr.SqlLang == ESqlDialect.Postgres {

		sqlLang = "postgres"
		dataSource = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			connStr.Host, connStr.Port, connStr.User, connStr.Password, connStr.DbName)
	} else if connStr.SqlLang == ESqlDialect.MySql {

		sqlLang = "mysql"
		dataSource = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			connStr.User,
			connStr.Password,
			connStr.Host,
			connStr.Port,
			connStr.DbName,
		)
	} else if connStr.SqlLang == ESqlDialect.MsSql {

		sqlLang = "mssql"
		dataSource = fmt.Sprintf(
			"Data Source=%s;"+
				"Initial Catalog=%s;"+
				"User id=%s;"+
				"password=%s;",
			connStr.Host,
			connStr.DbName,
			connStr.User,
			connStr.Password)
	} else {
		dataSource = ""
	}

	db, err := sql.Open(sqlLang, dataSource) // "&charset=utf8mb4,utf8"
	if err != nil {

		log.Printf(fmt.Sprintf("Failed to open sql connection to '%s' err:%v", dataSource, err))
		return nil, err
	}
	///defer db.Close()

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Printf(fmt.Sprintf("Failed to ping to sql connection to '%s' err:%v", dataSource, err))
		return nil, err
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	//static_db = db
	return db, err
}

func OpenDB(connStr TConnectionString, maxIdle int, maxOpen int) (*DBContextBase, error) {

	var db, err = StaticOpenDB(connStr, maxIdle, maxOpen)
	if err != nil {
		return nil, err
	}

	ctxBase := new(DBContextBase)
	ctxBase.ConnectionString = connStr
	ctxBase.Dialect = connStr.SqlLang
	ctxBase.Db = db

	return ctxBase, nil
}
