package atomicNSql

import (
	"context"

	fmt "fmt"
	//log "log"
	time "time"

	atomicsql "github.com/bbitere/atomicsql_golang.git/src/atomicsql"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)


func StaticOpenDB_NoSql( ctxBase *DBContextBaseNoSql, connStr atomicsql.TConnectionString)  error {

	var dataSource = ""
	//var sqlLang = ""
	var clientOptions *options.ClientOptions = nil;

	if connStr.SqlLang == atomicsql.ESqlDialect.MongoDB {

		var port = connStr.Port;
		if( port == 0){
			port = 27017;
		}
		

		//sqlLang = "mongo";
		if( connStr.User != ""){
			
			dataSource = fmt.Sprintf(			
				"mongodb://%s:%s@%s:%d/?authSource=%s",
				connStr.User,
				connStr.Password,
				connStr.Host,
				port,
				connStr.DbName,			
				)						
		}else{
			dataSource = fmt.Sprintf(			
				"mongodb://%s:%d/?authSource=%s",				
				connStr.Host,
				port,
				connStr.DbName,			
				)
		}
		clientOptions = options.Client().ApplyURI( dataSource )

		
	} else {
		dataSource = ""
	}

	var err error
	var client *mongo.Client = nil
	client, err = mongo.Connect( context.Background(), clientOptions )
	if err != nil {
		return err
	}

	// Ping the MongoDB server to check if the connection is established
	//ctx, cancel := context.WithTimeout(context.Background(), mm.timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	ctxBase.ClientDB = client;
	ctxBase.Database = client.Database( connStr.DbName );
	
	ctxBase.ConnectionString = connStr
	ctxBase.Dialect = connStr.SqlLang
	
	return err
}





func OpenDB_NoSql(connStr atomicsql.TConnectionString, maxIdle int, maxOpen int) (*DBContextBaseNoSql, error) {

	if( connStr.SqlLang == atomicsql.ESqlDialect.MongoDB){

		ctxBaseNoSql := new(DBContextBaseNoSql)

		var err = StaticOpenDB_NoSql(ctxBaseNoSql, connStr)
		if err != nil {
			return nil, err
		}

		return ctxBaseNoSql, nil
	}
	return nil,nil
}
