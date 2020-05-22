package core

import (
	"asher/internal/api"
	"fmt"
	"sort"
	"strings"
)

type ReturnArray struct {
	api.TabbedUnit
	tabs       int
	Statements []string
}

func NewReturnArrayFromMapRaw(arr map[string]string) *ReturnArray {
	return &ReturnArray{
		Statements: rawConvertMapToStringAssociativeArray(arr),
	}
}

func NewReturnArrayFromMap(arr map[string]string) *ReturnArray {
	return &ReturnArray{
		Statements: convertMapToStringAssociativeArray(arr),
	}
}

func NewReturnArray(arr []string) *ReturnArray {
	return &ReturnArray{
		Statements: arr,
	}
}

func (r *ReturnArray) SetNumTabs(tabs int) {
	r.tabs = tabs
}

func (r *ReturnArray) Id() string {
	return "return"
}

func (r *ReturnArray) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(r.tabs)), "return [\n", strings.Join(r.Statements, ",\n"), "];")
	return builder.String()
}

/**
Appends to return statements the given array
*/
func (r *ReturnArray) Append(arrayContent []string) {
	r.Statements = append(r.Statements, arrayContent...)
}

func convertMapToStringAssociativeArray(rulesMap map[string]string) []string {
	var returnVal []string
	keys := sortedKeysFromMap(rulesMap)
	for _, key := range keys {
		returnVal = append(returnVal, fmt.Sprintf(`"%s" => "%s"`, key, rulesMap[key]))
	}
	return returnVal
}

/*
 Fetches keys from a map and sorts them in ascending order.
 Parameters
 -	baseMap[string]string - The map whose keys are to be sorted and retured
 Returns
 - []string - A slice of keys sorted in the ascending order present in the map
 Usage
 myKeySlice := sortedKeysFromMap(someMap)
 */
func sortedKeysFromMap(baseMap map[string]string) []string {
	var keys []string
	for key := range baseMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}


func rawConvertMapToStringAssociativeArray(rulesMap map[string]string) []string {
	var returnVal []string
	keys := sortedKeysFromMap(rulesMap)
	for _, key := range keys {
		returnVal = append(returnVal, `'` + key + "' => " + rulesMap[key])
	}
	return returnVal
}