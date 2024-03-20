package atomicsql

import (
	sql "database/sql"
	time "time"
)

// utility method: array: contains
func Arr_Contains[T comparable](arr *[]T, value T) bool {

	for _, elem := range *arr {
		if elem == value {
			return true
		}
	}
	return false
}

// utility method: array: indexof
func Arr_IndexOf[T comparable](arr *[]T, index int, value T) int {

	var _arr = *arr
	var len1 = len(_arr)

	for i := 0; i < len1; i++ {

		if _arr[i] == value {
			return i
		}
	}
	return -1
}

// utility method: array: insert at index
func Arr_InsertAtIndex[T any](a *[]T, index int, value T) {

	if len(*a) == index { // nil or empty slice or after last element

		*a = append(*a, value)

	} else {
		*a = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
		(*a)[index] = value		
	}
}

// utility method: array: insert at create a slice
func Arr_Slice[T any](a *[]T, idx int, count int) []T {
	return (*a)[idx:count]
}

// utility method: array: append items
func Arr_Append[T any](a *[]T, value ...T) {

	*a = append(*a, value...)

}

// utility method: array: add range
func Arr_AddRange[T any](a *[]T, a2 *[]T) {

	*a = append(*a, (*a2)...)
}

// utility method: array: remove at index
func Arr_RemoveAt[T any](a *[]T, index int) {

	//var aa = make([]T, lenArr - 1);
	//var a1 *[]T = &aa;

	//memmove(unsafe.Pointer(&a1), unsafe.Pointer(&a), unsafe.Sizeof((a1) ))
	*a = append((*a)[:index], (*a)[index+1:]...)

}

// utility method: array: remove object
func Arr_Remove[T comparable](a *[]T, elem T) {

	for i, other := range *a {

		if other == elem {
			*a = append((*a)[:i], (*a)[i+1:]...)
			return
		}
	}

}

// utility method: string: substring
func Str_SubString(s string, index int, len int) string {
	return s[index : index+len]
}

// utility method: clone map[string][T]
func Util_CloneMapString[T any](dict *map[string]T) map[string]T {

	var ret = map[string]T{}
	for key, val := range *dict {
		ret[key] = val
	}
	return ret
}
type Aaa struct {m []string }
// utility method: convert from map to array. (values only)
func Util_FromMapToArray[T any](dict *map[string]T) []T {

	var ret = []T{}
	for _, val := range *dict {
		Arr_Append(&ret, val)
	}
	
	return ret
}

// utility method: convert from map to array. (keys only)
func Util_FromMapKeysToArray[T any](dict *map[string]T) []string {

	var ret = []string{}
	for key := range *dict {
		Arr_Append(&ret, key)
	}
	return ret
}

// Used to replace operator ?:
func IFF[T any](b bool, s1 T, s2 T) T {
	if b {
		return s1
	} else {
		return s2
	}
}


// In the next example you faild to do correct the things
//  var user2 = m.User{UserName: UserName2, Money: UserMoney, 
//                     UserRoleID: &m.UserRole{ RoleName: RoleNameAdmin, IsActive: true, 
//                		RoleStatusID:  &m.StatusRole{ StatusName: { String: StatusNameActive} }}};
// because  StatusName.String = StatusNameActive, and the developer omit to set Status.Valid =true, the code will have an issue
// for this you need to write the code like this:
//  var user2 = m.User{UserName: UserName2, Money: UserMoney, 
//                     UserRoleID: &m.UserRole{ RoleName: RoleNameAdmin, IsActive: true, 
//                		RoleStatusID:  &m.StatusRole{ StatusName: atmsql.Null_String( StatusNameActive ) }}};
func Null_String(s string ) sql.NullString{

	return sql.NullString{ String: s, Valid: true};
}

// see the comments for Null_String()
func Null_Int64(s int64) sql.NullInt64{

	return sql.NullInt64{ Int64: s, Valid: true};
}
// see the comments for Null_String()
func Null_Int32(s int32 ) sql.NullInt32{

	return sql.NullInt32{ Int32: s, Valid: true};
}

// see the comments for Null_String()
func Null_Int16(s int16 ) sql.NullInt16{

	return sql.NullInt16{ Int16: s, Valid: true};
}

// see the comments for Null_String()
func Null_Byte(s byte) sql.NullByte{

	return sql.NullByte{ Byte: s, Valid: true};
}


// see the comments for Null_String()
func Null_Bool(s bool ) sql.NullBool{

	return sql.NullBool{ Bool: s, Valid: true};
}

// see the comments for Null_String()
func Null_Float64(s float64 ) sql.NullFloat64{

	return sql.NullFloat64{ Float64: s, Valid: true};
}
// see the comments for Null_String()
func Null_Time(s time.Time ) sql.NullTime{

	return sql.NullTime{ Time: s, Valid: true};
}


func Utils_getMapCounter( dictMap *map[string]interface{}) int {

	var num = 0;
	for _, elem := range(*dictMap) {

		num++;
		if( &elem != nil ){}
	}
	return num;
}

func Utils_existInListString(list []string, item string) bool {

	for j := 0; j < len(list); j++ {
		if list[j] == item {
			return true
		}
	}
	return false
}

func Util_Nop(){
	
}