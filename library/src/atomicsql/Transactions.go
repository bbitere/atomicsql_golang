package atomicsql

import (
	//"reflect"
	context "context"
	sql "database/sql"
	"fmt"
)
type Transaction struct{
	tx 	*sql.Tx
	err error
}

func (_this *DBContextBase) BeginTx() (*Transaction, error){

	ctx := context.TODO()
	var tx, err = _this.Db.BeginTx( ctx, nil)
	var t = Transaction{ tx: tx, err: err};

	return &t, err
}

func (_this *DBContextBase) CommitOrRollBack(t *Transaction) ( error){
	
	var cnt = len( _this.stackTransactions );
	if( t != _this.stackTransactions[ cnt-1]  ){
		return fmt.Errorf("transaction is not the top transaction");
	}

	Arr_RemoveAt( &_this.stackTransactions, cnt-1);

	if( _this.hasError != nil){

		var err = t.tx.Rollback();
		return err
	}else{
		var err = t.tx.Commit();
		return err
	}
}

