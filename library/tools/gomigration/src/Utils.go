package main

import (
	//"fmt"
	//"os"
	"path"
	"strconv"
	"strings"
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
func Utils_GetListFromDict(dict map[string]interface{}) []interface{} {
	var lst []interface{}
	if dict == nil {
		return nil
	}

	for _, value := range dict {
		lst = append(lst, value)
	}
	return lst
}

// GetDictFromList function
func Utils_GetDictFromList(list []interface{}, fn func(interface{}) string) map[string]interface{} {
	dict := make(map[string]interface{})
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
func Utils_GetDictionaryDifference(dicB map[string]interface{}, dicA map[string]interface{}) map[string]interface{} {
	dict := make(map[string]interface{})
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
func Utils_GetDictionaryUnion(dicB map[string]interface{}, dicA map[string]interface{}) map[string]interface{} {
	dict := make(map[string]interface{})
	// Walk A, and if any of the entries are not
	// in B, add them to the result dictionary.

	for key, value := range dicA {
		if val, ok := dicB[key]; ok {
			dict[key] = map[string]interface{}{"ValueB": val, "ValueA": value}
		}
	}
	return dict
}

func strings_Index(s string, substrArg string, index int) int{

	var s1 = strings.Index( s[index:], substrArg);
	return s1; 
}