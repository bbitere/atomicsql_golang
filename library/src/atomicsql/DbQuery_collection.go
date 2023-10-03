package atomicsql

import (
	//"reflect"

	//"unsafe"

	"fmt"
	"strconv"

	//"math"
	//Sort "sort"
	"time"
	//"unsafe"

	//reflect
	//Str "strings"

	//log "log"
	"reflect"
)



func (_this *DBQuery[T]) _RtmWhereEq( model *T, fieldName string, operands []any, field2 string, bNot bool) bool {

	var valFld, _ = _this.getModel_FieldValueS( model, model, fieldName, false);

	for i := 0; i < len(operands); i ++{

		var valOperand, _ = _this.getValueS( operands[i], fieldName);

		if( bNot){
			if( valFld == valOperand ){
				return false;
			}
		}else{
			if( valFld == valOperand ){
				return true;
			}
		}
	}
	if( bNot){
		return true;
	}else{
		return false;
	}
}

func (_this *DBQuery[T])  _getDistinctRTM( fields []string, models []*T) []*T{

	var dict = make( map[string] *T);

	for i := 0; i < len(models); i++{

		var model = models[i];

		var key = "";
		for iFld := 0; iFld < len(models); iFld++ {

			var fldTName = fields[ iFld ];

			var s, _ = _this.getModel_FieldValueS( model, model, fldTName, false );
			key += s;
		}
		dict[ key ] = model;
	}
	var arr = Util_FromMapToArray( &dict );
	return arr;
}

func _Aggregate_doRuntime[T IGeneric_MODEL, V IGeneric_MODEL]( _this *DBQuery[T], models []*T )  ([]*V, error) {

	type TAggrElem struct{
		modelsT []*T;
		modelV  *V;
	}
	var dict = make( map[string] *TAggrElem);

	var fieldsCommon = []string{};
	var fieldsGroupBy = []string{};

	var modelV1 *V = new(V)

	var reflectVal = reflect.TypeOf( modelV1 ).Elem()
	{
		for iFld := 0; iFld < reflectVal.NumField(); iFld++ {

			var fldT = reflectVal.Field( iFld );
			if( fldT.Type.Kind() == reflect.Array ||
			    fldT.Type.Kind() == reflect.Slice){
				Arr_Append( &fieldsGroupBy, fldT.Name );	// for time1 and Money
			}else{
				Arr_Append( &fieldsCommon, fldT.Name );	//for RoleName
			}			
		}		
	}
	var errTxt = "";

	for i := 0; i < len(models); i++{

		var modelT = models[i];
		var modelV  = new(V)

		var key = "";
		for iFld := 0; iFld < len(fieldsCommon); iFld++ {

			var fldTName = fieldsCommon[ iFld ];

			var s,bIsSet = _this.getModel_FieldValueS( modelT, modelV, fldTName, true);
			if( bIsSet){
				key += s;
			}else{
				errTxt += fmt.Sprintf("Model cannot set field %s ", fldTName)
			}
		}

		if( errTxt == ""){

			if( dict[key] == nil){

				dict[ key ] = &TAggrElem{
					modelsT: []*T{},
					modelV: modelV,
				};
			}else{
				Arr_Append( &dict[ key ].modelsT, modelT);
			}		
		}
	}
	var arr []*V = []*V{};	
	
	for _, it := range dict {

		for _, fldTName := range fieldsGroupBy {

			var bIsSet = _this.getModel_FieldValueAsArr( &it.modelsT, it.modelV, fldTName );
			if( !bIsSet){
				errTxt += fmt.Sprintf("Model cannot set field %s ", fldTName)
			}			
		}
		Arr_Append(&arr, modelV1)
	}
	if( errTxt != ""){
		return arr, fmt.Errorf(errTxt);
	}else{
		return arr, nil;
	}	
	
}


func (_this *DBQuery[T]) getModel_FieldValueS( modelT *T, modelV any, fieldName string, bDoSet bool) (string, bool){

	var reflectValT = reflect.ValueOf( modelT ).Elem()
	var reflectValV = reflect.ValueOf( modelV ).Elem()

	if( reflectValT.IsValid() && reflectValV.IsValid() ){

		var fldT = reflectValT.FieldByName( fieldName );
		var fldV = reflectValV.FieldByName( fieldName );

		var fldTNullable = _this._getNullableField( fldT );
		if( fldTNullable != nil && fldTNullable.Valid.Bool() ){
			fldT = fldTNullable.Value
		}

		var fldVNullable = _this._getNullableField( fldV );
		if( fldVNullable != nil && fldVNullable.Valid.Bool() ){
			fldV = fldVNullable.Value
		}

		var fieldInfoTypeT = fldT.Type()
		var fieldInfoTypeV = fldV.Type()

		if  fieldInfoTypeT == reflect.TypeOf((*int16)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int32)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int64)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int8)(nil)).Elem() {	
				
				if  !(fieldInfoTypeV == reflect.TypeOf((*int16)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*int32)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*int64)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*int)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*int8)(nil)).Elem()) {
						return "", false
				}
				if( bDoSet){
					fldV.SetInt( fldT.Int() )
				}
				
				return fmt.Sprintf( "%d", fldT.Int() ), true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*bool)(nil)).Elem() ){

			if( !(fieldInfoTypeV == reflect.TypeOf((*bool)(nil)).Elem()) ){
				return "", false
			}
			if( bDoSet){
				fldV.SetBool( fldT.Bool() )
			}

			return IFF( fldT.Bool(), "true", "false"), true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*float32)(nil)).Elem() ||
			fieldInfoTypeT == reflect.TypeOf((*float64)(nil)).Elem() ){

				if( !(fieldInfoTypeV == reflect.TypeOf((*float32)(nil)).Elem() ||
					  fieldInfoTypeV == reflect.TypeOf((*float64)(nil)).Elem() ) ){
					return "", false
				}
				if( bDoSet){
					fldV.SetFloat( fldT.Float() )
				}
				return fmt.Sprintf( "%f", fldT.Float() ), true;
		} else 
		if(  fieldInfoTypeT == reflect.TypeOf((*string)(nil)).Elem() ) {

			if( !(fieldInfoTypeV == reflect.TypeOf((*string)(nil)).Elem() ) ) {
				return "", false
			}
			if( bDoSet){
				fldV.SetString( fldT.String() )
			}
			return fldT.String(), true
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*time.Time)(nil)).Elem() ){

			if( fieldInfoTypeV == reflect.TypeOf((*time.Time)(nil)).Elem() ){
				return "", false
			}
			if( bDoSet){
				fldV.SetString( fldT.String() )
			}
			return fldT.String(), true
		}else
		if fieldInfoTypeT == reflect.TypeOf((*[]uint8)(nil)).Elem() {

			var slice = fldT.Bytes()

			var txt = "";
			for i := 0; i < len(slice); i++{
				txt += fmt.Sprintf( "%d,", slice[i]);
			}
			
			if( fieldInfoTypeV == reflect.TypeOf((*[]uint8)(nil)).Elem() ){
				return "", false
			}
			if( bDoSet){
				fldV.SetBytes( fldT.Bytes() )
			}
			return txt, true
		}		
	}
	return "", true
}


func (_this *DBQuery[T]) getModel_FieldValueAsArr( modelT *[]*T, modelV any, fieldName string) bool{

	var reflectValT = reflect.ValueOf( modelT ).Elem()
	var reflectValV = reflect.ValueOf( modelV ).Elem()

	if( reflectValT.IsValid() && reflectValV.IsValid() ){

		var fldT = reflectValT.FieldByName( fieldName );
		var fldV = reflectValV.FieldByName( fieldName );

		if( !fldV.CanSet() ){
			return false;
		}

		//because modelT is *[]*T
		//fldT = fldT.Elem().Elem();
		
		var fldTNullable = _this._getNullableField( fldT );
		if( fldTNullable != nil && fldTNullable.Valid.Bool() ){
			fldT = fldTNullable.Value
		}

		var fldVNullable = _this._getNullableField( fldV );
		if( fldVNullable != nil && fldVNullable.Valid.Bool() ){
			fldV = fldVNullable.Value
		}
		
		var fieldInfoTypeT = fldT.Type()
		var fieldInfoTypeV = fldV.Type()

		if  fieldInfoTypeT == reflect.TypeOf((*[]int16)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*[]int32)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*[]int64)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*[]int)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*[]int8)(nil)).Elem() {	
				
				if  !(fieldInfoTypeV == reflect.TypeOf((*[]int16)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*[]int32)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*[]int64)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*[]int)(nil)).Elem() ||
					fieldInfoTypeV == reflect.TypeOf((*[]int8)(nil)).Elem()) {
						return false
				}
				fldV.Set( fldT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]bool)(nil)).Elem() ){

			if( !(fieldInfoTypeV == reflect.TypeOf((*[]bool)(nil)).Elem()) ){
				return false
			}
			fldV.Set( fldT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]float32)(nil)).Elem() ||
			fieldInfoTypeT == reflect.TypeOf((*[]float64)(nil)).Elem() ){

				if( !(fieldInfoTypeV == reflect.TypeOf((*[]float32)(nil)).Elem() ||
					  fieldInfoTypeV == reflect.TypeOf((*[]float64)(nil)).Elem() ) ){
					return false
				}
				fldV.Set( fldT );
				return true;
		} else 
		if(  fieldInfoTypeT == reflect.TypeOf((*[]string)(nil)).Elem() ) {

			if( !(fieldInfoTypeV == reflect.TypeOf((*[]string)(nil)).Elem() ) ) {
				return false
			}
			fldV.Set( fldT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]time.Time)(nil)).Elem() ){

			if( fieldInfoTypeV == reflect.TypeOf((*[]time.Time)(nil)).Elem() ){
				return false
			}
			fldV.Set( fldT );
			return true;
		}else
		if fieldInfoTypeT == reflect.TypeOf((*[][]uint8)(nil)).Elem() {

			if( fieldInfoTypeV == reflect.TypeOf((*[][]uint8)(nil)).Elem() ){
				return false
			}
			fldV.Set( fldT );
				return true;
		}				
	}
	return  true
}


func (_this *DBQuery[T]) getValueI( modelV any, fieldName string ) (int64, error){

	var ret, err = _this.getValueS(modelV, fieldName);
	if( err != nil ){
		return 0, err
	}
	var i, err2 = strconv.Atoi( ret );
	return int64(i), err2
}
func (_this *DBQuery[T]) getValueS( modelV any, fieldName string ) (string, error){

	var reflectValV = reflect.ValueOf( modelV ).Elem()

	if( reflectValV.IsValid() ){

		var fldV = reflectValV.FieldByName( fieldName );

		var fldVNullable = _this._getNullableField( fldV );
		if( fldVNullable != nil && fldVNullable.Valid.Bool() ){
			fldV = fldVNullable.Value
		}

		var fieldInfoTypeV = fldV.Type()

		if  fieldInfoTypeV == reflect.TypeOf((*int16)(nil)).Elem() ||
	   		fieldInfoTypeV == reflect.TypeOf((*int32)(nil)).Elem() ||
	   		fieldInfoTypeV == reflect.TypeOf((*int64)(nil)).Elem() ||
	   		fieldInfoTypeV == reflect.TypeOf((*int)(nil)).Elem() ||
	   		fieldInfoTypeV == reflect.TypeOf((*int8)(nil)).Elem() {	
				
				return fmt.Sprintf( "%d", fldV.Int() ), nil;
		} else 
		if( fieldInfoTypeV == reflect.TypeOf((*bool)(nil)).Elem() ){

			return IFF( fldV.Bool(), "true", "false"), nil;
		} else 
		if( fieldInfoTypeV == reflect.TypeOf((*float32)(nil)).Elem() ||
			fieldInfoTypeV == reflect.TypeOf((*float64)(nil)).Elem() ){

				return fmt.Sprintf( "%f", fldV.Float() ), nil;
		} else 
		if(  fieldInfoTypeV == reflect.TypeOf((*string)(nil)).Elem() ) {

			return fldV.String(), nil
		} else 
		if( fieldInfoTypeV == reflect.TypeOf((*time.Time)(nil)).Elem() ){

			return fldV.String(), nil
		}else
		if fieldInfoTypeV == reflect.TypeOf((*[]uint8)(nil)).Elem() {

			var slice = fldV.Bytes()

			var txt = "";
			for i := 0; i < len(slice); i++{
				txt += fmt.Sprintf( "%d,", slice[i]);
			}
			
			return txt, nil
		}		
	}
	return "", fmt.Errorf("field '%s' not convert to data type", fieldName)
}


