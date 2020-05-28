package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type ForEach struct {
	api.TabbedUnit
	tabs       int
	Condition  string
	Statements []api.TabbedUnit
}

func NewForEach() *ForEach {
	return &ForEach{
		tabs:       0,
		Statements: []api.TabbedUnit{},
	}
}

func (forEach *ForEach) SetNumTabs(tabs int) {
	forEach.tabs = tabs
}

func (forEach *ForEach) AddStatement(statement api.TabbedUnit) *ForEach {
	forEach.Statements = append(forEach.Statements, statement)
	return forEach
}

func (forEach *ForEach) AddStatements(statements []api.TabbedUnit) *ForEach {
	forEach.Statements = append(forEach.Statements, statements...)
	return forEach
}

func (forEach *ForEach) AddCondition(unit string) *ForEach {
	forEach.Condition = unit
	return forEach
}

func (forEach *ForEach) String() string {
	var builder strings.Builder
	tabbedString := api.TabbedString(uint(forEach.tabs))

	fmt.Fprint(&builder, tabbedString, "foreach ( "+forEach.Condition+") {\n")

	for _, element := range forEach.Statements {
		element.SetNumTabs(forEach.tabs + 1)
		fmt.Fprint(&builder, element.String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, "}\n")
	return builder.String()
}
