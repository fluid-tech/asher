package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type Function struct {
	api.TabbedUnit
	tabs       int
	Name       string
	Visibility string
	Static     bool
	Arguments  []string
	Statements []api.TabbedUnit
}

func NewFunction() *Function {
	return &Function{
		TabbedUnit: nil,
		tabs:       0,
		Name:       "",
		Visibility: "",
		Static:     false,
		Arguments:  []string{},
		Statements: []api.TabbedUnit{},
	}
}

func (f *Function) SetNumTabs(tabs int) {
	f.tabs = tabs
}

func (f *Function) String() string {
	var builder strings.Builder
	tabbedString := api.TabbedString(uint(f.tabs))
	fmt.Fprint(&builder, tabbedString, f.Visibility, staticStr(f.Static), " function ", f.Name, "(")
	fmt.Fprint(&builder, strings.Join(f.Arguments, ", "), ") {\n")
	for _, element := range f.Statements {
		element.SetNumTabs(f.tabs + 1)
		fmt.Fprint(&builder, element.String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, "}\n\n")
	return builder.String()
}

func staticStr(isStatic bool) string {
	if isStatic {
		return " static"
	}
	return ""
}

/**
Append Statement
*/
func (f *Function) AppendStatement(unit api.TabbedUnit) {
	f.Statements = append(f.Statements, unit)
}

/**
Appends Argument
*/
func (f *Function) AppendArgument(unit string) {
	f.Arguments = append(f.Arguments, unit)
}
