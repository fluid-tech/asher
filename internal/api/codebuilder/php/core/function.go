package core

import (
	"asher/internal/api"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Function struct{
	api.TabbedUnit
	tabs int
	Name string
	Visibility string
	Static bool
	Arguments []string
	Statements []*api.TabbedUnit
}

func NewFunction() *Function {
	return &Function{
		TabbedUnit: nil,
		tabs:       0,
		Name:       "",
		Visibility: "",
		Static: false,
		Arguments:  []string{},
		Statements: []*api.TabbedUnit{},
	}
}

func (f *Function) SetNumTabs(tabs int){
	f.tabs = tabs
}

func (f *Function) Id() string {
	if f.Name == "" {
		return "anon"
	}
	return f.Name
}

func (f *Function) String() string {
	var builder strings.Builder
	tabbedString :=  api.TabbedString(uint(f.tabs))
	fmt.Fprint(&builder, tabbedString, f.Visibility, getStaticValue(f.Static), " function ", f.Name, "(")

	fmt.Fprint(&builder, strings.Join(f.Arguments, ", "),") {\n")
	for _, element := range f.Statements{
		(*element).SetNumTabs(f.tabs + 1)
		fmt.Fprint(&builder, (*element).String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, "}\n\n")
	return builder.String()
}

func getStaticValue(isStatic bool) string {
	if isStatic {
		return " static"
	}
	return ""
}

/**
Finds a statement with a regex
 */
func (f *Function) FindStatement(pattern string) (*api.TabbedUnit, error) {
	for _, element := range f.Statements {
		found, err := regexp.Match(pattern, []byte((*element).Id()))
		if  err == nil && found {
			return element, nil
		}
	}
	return nil, errors.New("couldnt find pattern")
}

/**
Finds a tabbed unit by id
 */
func (f *Function) FindById(id string) (*api.TabbedUnit, error) {
	for _, element := range f.Statements {
		if  (*element).Id() == id {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("cant find statement with identifier %s", id))
}

/**
Append Statement
 */
func (f *Function) AppendStatement(unit *api.TabbedUnit) {
	f.Statements = append(f.Statements, unit)
}