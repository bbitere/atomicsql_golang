package atomicsql

import (
	"testing"
)

func Test1(t *testing.T) {

	var vv TTime
	var xx = vv.Day()
	if( xx != 1 ){
		t.Errorf("day passed")
		return;
	}
	
	//return t.Deadline()
}