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



func (_this *DBQueryNoSql[T]) _RtmWhereEq( model *T, fieldName string, operands []any, field2 string, bNot bool) bool {

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

func (_this *DBQueryNoSql[T])  _getDistinctRTM( fields []string, models []*T) []*T{

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

func _Aggregate_doRuntime_NSql[T IGeneric_MODEL, V IGeneric_MODEL]( _this *DBQueryNoSql[T], models []*T )  ([]*V, error) {

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
			if( fldT.Anonymous){
				continue;
			}
			if( fldT.Type.Kind() == reflect.Array ||
			    fldT.Type.Kind() == reflect.Slice){
				
				// the groups fields are (name, []type)
				Arr_Append( &fieldsGroupBy, fldT.Name );	// for time1 and Money
			}else{

				// the common fields are the fields that  are similar as (name, type)
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
				key += s + "_";
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
			}
			Arr_Append( &dict[ key ].modelsT, modelT);					
		}
	}
	var arr []*V = []*V{};	
	
	//do this for all keys in dictionary. 
	// the key is made by all common fields using underrscode between them
	for _, it := range dict {
		
		var arrLen = len(it.modelsT);
		for _, fldTName := range fieldsGroupBy {

			//put in each field of modelV the array of elem from all modelsT for each corespondent Fields
			// for ex:
			// for field F1, create an array. extract the values from all modelsT for field F1. 
			// arr = { modelT[0].F1, .., modelT[n-1].F1 }
			// -> set the field modelV.F1 = arr
			var reflectValV = reflect.ValueOf( it.modelV ).Elem() //V
			if( !reflectValV.IsValid() ){
				continue;
			}
			var fldV = reflectValV.FieldByName( fldTName );

			var fldVIndirect =  reflect.Indirect(fldV)
			var reflArr = reflect.MakeSlice( fldVIndirect.Type(), 0, arrLen );

			for iElem := 0; iElem < arrLen; iElem++{

				// var val = it.ModelsT[iElem] 
				var valModel = reflect.ValueOf( it.modelsT[iElem] ).Elem();
				var valFld   = valModel.FieldByName( fldTName );

				reflArr = reflect.Append( reflArr, valFld );
			}

			fldV.Set( reflArr );
		}
		Arr_Append(&arr, it.modelV)
	}
	if( errTxt != ""){
		return arr, fmt.Errorf(errTxt);
	}else{
		return arr, nil;
	}		
}

func (_this *DBQueryNoSql[T]) getModel_FieldValueS( modelT *T, modelV any, fieldName string, bDoSet bool) (string, bool){

	var reflectValT = reflect.ValueOf( modelT ).Elem()
	var reflectValV = reflect.ValueOf( modelV ).Elem()

	if( reflectValT.IsValid() && reflectValV.IsValid() ){

		var fldT = reflectValT.FieldByName( fieldName );
		var fldV = reflectValV.FieldByName( fieldName );

		
		var fldTNullable = _getNullableField( fldT );
		if( fldTNullable != nil && fldTNullable.Valid.Bool() ){
			fldT = fldTNullable.Value
		}

		if( bDoSet){
			var fldVNullable = _getNullableField( fldV );
			if( fldVNullable != nil /*&& fldVNullable.Valid.Bool()*/ ){
				fldV = fldVNullable.Value

				if( fldTNullable != nil && bDoSet ){

					var fValT = fldTNullable.Valid
					var fValV = fldVNullable.Valid
					fValV.Set( fValT );
				}
			}
		}

		var fieldInfoTypeT = fldT.Type()
		var fieldInfoTypeV = fldV.Type()

		if  fieldInfoTypeT == reflect.TypeOf((*int16)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int32)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int64)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int)(nil)).Elem() ||
	   		fieldInfoTypeT == reflect.TypeOf((*int8)(nil)).Elem() {	
				
				if( bDoSet){
					if  !(fieldInfoTypeV == reflect.TypeOf((*int16)(nil)).Elem() ||
						fieldInfoTypeV == reflect.TypeOf((*int32)(nil)).Elem() ||
						fieldInfoTypeV == reflect.TypeOf((*int64)(nil)).Elem() ||
						fieldInfoTypeV == reflect.TypeOf((*int)(nil)).Elem() ||
						fieldInfoTypeV == reflect.TypeOf((*int8)(nil)).Elem()) {
							return "", false
					}				
					fldV.SetInt( fldT.Int() )
				}
				
				return fmt.Sprintf( "%d", fldT.Int() ), true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*bool)(nil)).Elem() ){

			if( bDoSet){

				if( !(fieldInfoTypeV == reflect.TypeOf((*bool)(nil)).Elem()) ){
					return "", false
				}
			
				fldV.SetBool( fldT.Bool() )
			}

			return IFF( fldT.Bool(), "true", "false"), true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*float32)(nil)).Elem() ||
			fieldInfoTypeT == reflect.TypeOf((*float64)(nil)).Elem() ){

				if( bDoSet){

					if( !(fieldInfoTypeV == reflect.TypeOf((*float32)(nil)).Elem() ||
						fieldInfoTypeV == reflect.TypeOf((*float64)(nil)).Elem() ) ){
						return "", false
					}
				
					fldV.SetFloat( fldT.Float() )
				}
				return fmt.Sprintf( "%f", fldT.Float() ), true;
		} else 
		if(  fieldInfoTypeT == reflect.TypeOf((*string)(nil)).Elem() ) {

			if( bDoSet){

				if( !(fieldInfoTypeV == reflect.TypeOf((*string)(nil)).Elem() ) ) {
					return "", false
				}			
				fldV.SetString( fldT.String() )
			}
			return fldT.String(), true
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*time.Time)(nil)).Elem() ){

			if( bDoSet){

				if( fieldInfoTypeV == reflect.TypeOf((*time.Time)(nil)).Elem() ){
					return "", false
				}			
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
			
			if( bDoSet){

				if( fieldInfoTypeV == reflect.TypeOf((*[]uint8)(nil)).Elem() ){
					return "", false
				}				
				fldV.SetBytes( fldT.Bytes() )
			}
			return txt, true
		}else
		{
			if fieldInfoTypeT == fieldInfoTypeV {
				if( bDoSet){
					fldV.Set( fldT )
				}
				return "", true
			}
		}		
	}
	return "", false
}

/*
func (_this *DBQueryNoSql[T]) getModel_FieldValueAsArr( reflValueArrayModelsT reflect.Value, modelV any, fieldName string) bool{
	
	//var reflectValT = reflect.ValueOf( modelsT )
	var reflectValV = reflect.ValueOf( modelV ).Elem() //V

	if(  reflectValV.IsValid() ){

		//var fldT = reflectValT.FieldByName( fieldName );
		var fldV = reflectValV.FieldByName( fieldName );

		if( !fldV.CanSet() ){
			return false;
		}

		//because modelT is *[]*T
		//fldT = fldT.Elem().Elem();
		
		//var fldTNullable = _this._getNullableField( fldT );
		//if( fldTNullable != nil && fldTNullable.Valid.Bool() ){
		//	fldT = fldTNullable.Value
		//}

		//var fieldInfoTypeT = reflectValT.Type()
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
				fldV.Set( reflectValT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]bool)(nil)).Elem() ){

			if( !(fieldInfoTypeV == reflect.TypeOf((*[]bool)(nil)).Elem()) ){
				return false
			}
			fldV.Set( reflectValT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]float32)(nil)).Elem() ||
			fieldInfoTypeT == reflect.TypeOf((*[]float64)(nil)).Elem() ){

				if( !(fieldInfoTypeV == reflect.TypeOf((*[]float32)(nil)).Elem() ||
					  fieldInfoTypeV == reflect.TypeOf((*[]float64)(nil)).Elem() ) ){
					return false
				}
				fldV.Set( reflectValT );
				return true;
		} else 
		if(  fieldInfoTypeT == reflect.TypeOf((*[]string)(nil)).Elem() ) {

			if( !(fieldInfoTypeV == reflect.TypeOf((*[]string)(nil)).Elem() ) ) {
				return false
			}
			fldV.Set( reflectValT );
				return true;
		} else 
		if( fieldInfoTypeT == reflect.TypeOf((*[]time.Time)(nil)).Elem() ){

			if( fieldInfoTypeV == reflect.TypeOf((*[]time.Time)(nil)).Elem() ){
				return false
			}
			fldV.Set( reflectValT );
			return true;
		}else
		if fieldInfoTypeT == reflect.TypeOf((*[][]uint8)(nil)).Elem() {

			if( fieldInfoTypeV == reflect.TypeOf((*[][]uint8)(nil)).Elem() ){
				return false
			}
			fldV.Set( reflectValT );
				return true;
		}else	
		if reflectValT == fldV {

			fldV.Set( reflectValT );
				return true;
		}			
	}
	return  true
}
*/

func (_this *DBQueryNoSql[T]) getValueI( modelV any, fieldName string ) (int64, error){

	var ret, err = _this.getValueS(modelV, fieldName);
	if( err != nil ){
		return 0, err
	}
	var i, err2 = strconv.Atoi( ret );
	return int64(i), err2
}
func (_this *DBQueryNoSql[T]) getValueS( modelV any, fieldName string ) (string, error){

	var reflectValV = reflect.ValueOf( modelV ).Elem()

	if( reflectValV.IsValid() ){

		var fldV = reflectValV.FieldByName( fieldName );

		var fldVNullable = _getNullableField( fldV );
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


func (_this *DBQueryNoSql[T]) rtm_getModelsAsDicts( models *[]*T, fields []string ) []*TModel[T] {

	var arr = []*TModel[T]{}

	for iElem := 0; iElem < len(*models); iElem++ {
		
		var model = (*models)[ iElem ];
		var dict = map[string]string{}

		for iFld := 0; iFld < len(fields); iFld++ {
		
			var fieldName = fields[iFld];
			var valS, _ = _this.getModel_FieldValueS( model, model, fieldName, false);
			dict[ fieldName ] = valS;
		}
		var elem = new (TModel[T]);
		elem.model =  model;
		elem.dict  = dict;
		Arr_Append( &arr, elem);
	}
	
	return arr;
}

func (_this *DBQueryNoSql[T]) rtm_updateModelsFromDicts( models1 *[]*TModel[T] ) []*T {

	var arr = []*T{}

	for iElem := 0; iElem < len(*models1); iElem++ {

		Arr_Append( &arr, (*models1)[iElem].model );
	}
	return arr;
}