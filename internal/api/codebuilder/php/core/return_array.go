package core

import (
	"fmt"
	"strings"
)

type ReturnArray struct{
	TabbedUnit
	tabs int
	Statements []string
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
	fmt.Fprint(&builder, TabbedString(uint(r.tabs)), "return [\n", strings.Join(r.Statements, ",\n"), "];\n")
	return builder.String()
}
/**
Appends to return statements the given array
 */
func (r *ReturnArray) Append(arrayContent []string){
	r.Statements = append(r.Statements, arrayContent...)
}