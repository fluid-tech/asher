package core

import (
	"fmt"
	"regexp"
	"strings"
)

type Function struct{
	TabbedUnit
	tabs int
	Name string
	Visibility string
	Arguments []string
	Statements []*TabbedUnit
}

func NewFunction() *Function {
	return &Function{
		TabbedUnit: nil,
		tabs:       0,
		Name:       "",
		Visibility: "",
		Arguments:  []string{},
		Statements: []*TabbedUnit{},
	}
}

func (f *Function) SetNumTabs(tabs int){
	f.tabs = tabs
}

func (f *Function) Id() string {
	return f.Name
}

func (f *Function) String() string {
	var builder strings.Builder
	tabbedString :=  TabbedString(uint(f.tabs))
	fmt.Fprint(&builder, tabbedString, f.Visibility, " function ", f.Name, "(")

	fmt.Fprint(&builder, strings.Join(f.Arguments, ", "),") {\n")
	for _, element := range f.Statements{
		(*element).SetNumTabs(f.tabs + 1)
		fmt.Fprint(&builder, (*element).String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, "}\n\n")
	return builder.String()
}

func (f *Function) FindStatement(pattern string) (*TabbedUnit, error) {
	var err error
	for _, element := range f.Statements {
		found, err := regexp.Match(pattern, []byte((*element).Id()))
		if  err == nil && found {
			return element, nil
		}
	}
	return nil, err
}