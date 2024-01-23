package gomigration_utils

import (
	//"fmt"
	//"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

// Nop function
func Utils_Utils_Nop() {}

// GetFileInfoName function
func Utils_GetFileInfoName(filePath string) string {

    var fileName = path.Base(filePath);
    var ext = path.Ext(filePath)

	if ext == "" {
		return fileName;
	}
	return strings.Replace(fileName, ext, "", 1);
}
func Utils_GetFileInfoExt(filePath string) string {

	var ext = path.Ext(filePath)
	return ext;
}

// NormalizePath function
func Utils_NormalizePath(path string) string {
	if path == "" {
		return path
	}
	return strings.Replace(path, "/", "\\", -1)
}

// NormalizeName function
func Utils_NormalizeName(name string) string {
	
    var outName     = strings.Replace(name, "-", "_", -1)
    var _, err = strconv.Atoi(outName); 

	
	if outName == "in" ||
		outName == "for" ||
		outName == "if" ||
		outName == "while" ||
		outName == "do" ||
		outName == "switch" ||
		err == nil {
		outName = "_" + outName
	}
	return outName
}

// GetListFromDict function
func Utils_GetListFromDict[T any](dict map[string]*T) [] *T {
	var lst []*T
	if dict == nil {
		return nil
	}

	for _, value := range dict {
		lst = append(lst, value)
	}
	return lst
}

// GetDictFromList function
func Utils_GetDictFromList[T any](list []*T, fn func(*T) string) map[string]*T {

	dict := make(map[string]*T)
	if list == nil {
		return nil
	}

	for _, it := range list {
		key := fn(it)
		dict[key] = it
	}
	return dict
}

// GetDictionaryDifference function
func Utils_GetDictionaryDifference[T any](dicB map[string] *T, dicA map[string] *T) map[string] *T {
	dict := make(map[string] *T)
	// Walk A, and if any of the entries are not
	// in B, add them to the result dictionary.

	for key, value := range dicB {
		if _, ok := dicA[key]; !ok {
			dict[key] = value
		}
	}
	return dict
}

// GetDictionaryUnion function
func Utils_GetDictionaryUnion[T any, V any](dicB map[string] *T, dicA map[string] *V,
	) map[string]struct{ Item1 *T; Item2 *V } {

	type TT = struct{ Item1 *T; Item2 *V };
	dict := make(map[string] TT )
	// Walk A, and if any of the entries are not
	// in B, add them to the result dictionary.

	for key, value := range dicA {
		if val, ok := dicB[key]; ok {
			
			dict[key] = TT{ Item1: val, Item2: value};
			//map[string] *T{"ValueB": val, "ValueA": value}
		}
	}
	return dict
}

func Strings_Index(s string, substrArg string, index int) int{

	var s1 = strings.Index( s[index:], substrArg);
	return s1; 
}

const POSTFIX_FOREIGNKEY = "_ID";

func Utils_ConvertToIdentGoLang(ident string, removeUnderscore bool) string {
	var ident1 string
	var nextUpper, isUpper bool

	if ident == "user_psw" {
		// Utils.Nop() equivalent in Go (no-op function)
	}

	for i := 0; i < len(ident); i++ {
		ch := string(ident[i])
		var ch1 string
		if i+1 < len(ident) {
			ch1 = string(ident[i+1])
		}

		if ch == "_" && removeUnderscore &&
			((ident[i:] == POSTFIX_FOREIGNKEY) ||
				(ch1 != "_" && unicode.ToUpper(rune(ch1[0])) != rune(ch1[0]) && ch1 != "")) {
			ch = ""
			nextUpper = true
		} else if ch == " " {
			ch = "_"
		}

		if isUpper {
			ident1 += strings.ToUpper(ch)
		} else {
			ident1 += ch
		}
		isUpper = nextUpper

		nextUpper = false
	}

	if ident1 == "User_psw" {
		// Utils.Nop() equivalent in Go (no-op function)
	}
	return ident1
}