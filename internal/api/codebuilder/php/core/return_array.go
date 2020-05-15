package core

import (
	"fmt"
	"strings"
)

type ReturnArray struct{
	TabbedUnit
	tabs int
	statements []string
}

func NewReturnArray(arr []string) *ReturnArray {
	return &ReturnArray{
		statements: arr,
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
	fmt.Fprint(&builder, TabbedString(uint(r.tabs)), "return [\n", strings.Join(r.statements, ",\n"), "];\n")
	return builder.String()
}
/**
Appends to return statements the given array
 */
func (r *ReturnArray) Append(arrayContent []string){
	r.statements = append(r.statements, arrayContent...)
}