package atomicsql

import (
	sql "database/sql"
)

func Arr_Contains[T comparable](arr *[]T, value T) bool {

	for _, elem := range *arr {
		if elem == value {
			return true
		}
	}
	return false
}

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

func Arr_InsertAtIndex[T any](a *[]T, index int, value T) {

	if len(*a) == index { // nil or empty slice or after last element

		*a = append(*a, value)

	} else {
		*a = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
		(*a)[index] = value
	}
}

func Arr_Slice[T any](a *[]T, idx int, count int) []T {
	return (*a)[idx:count]
}

func Arr_Append[T any](a *[]T, value ...T) {

	*a = append(*a, value...)

}

func Arr_AddRange[T any](a *[]T, a2 *[]T) {

	*a = append(*a, (*a2)...)
}
func Arr_RemoveAt[T any](a *[]T, index int) {

	//var aa = make([]T, lenArr - 1);
	//var a1 *[]T = &aa;

	//memmove(unsafe.Pointer(&a1), unsafe.Pointer(&a), unsafe.Sizeof((a1) ))
	*a = append((*a)[:index], (*a)[index+1:]...)

}

func Arr_Remove[T comparable](a *[]T, elem T) {

	for i, other := range *a {

		if other == elem {
			*a = append((*a)[:i], (*a)[i+1:]...)
			return
		}
	}

}

func Str_SubString(s string, index int, len int) string {
	return s[index : index+len]
}

func Util_CloneMapString[T any](dict *map[string]T) map[string]T {

	var ret = map[string]T{}
	for key, val := range *dict {
		ret[key] = val
	}
	return ret
}
func Util_FromMapToArray[T any](dict *map[string]T) []T {

	var ret = []T{}
	for _, val := range *dict {
		Arr_Append(&ret, val)
	}
	return ret
}

func Null_String(s string ) sql.NullString{

	return sql.NullString{ String: s, Valid: true};
}