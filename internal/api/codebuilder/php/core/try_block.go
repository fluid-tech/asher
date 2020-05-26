package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type TryBlock struct {
	api.TabbedUnit
	tabs              int
	Statements        []api.TabbedUnit
	CatchBlock        []CatchBlock
	FinallyStatements []api.TabbedUnit
}

func NewTryBlock() *TryBlock {
	return &TryBlock{
		tabs:              0,
		Statements:        []api.TabbedUnit{},
		CatchBlock:        []CatchBlock{},
		FinallyStatements: []api.TabbedUnit{},
	}
}

func (tryBlock *TryBlock) Id() string {
	return "try"
}

func (tryBlock *TryBlock) SetNumTabs(tabs int) {
	tryBlock.tabs = tabs
}


func (tryBlock *TryBlock) String() string {
	var builder strings.Builder
	tabbedString := api.TabbedString(uint(tryBlock.tabs))

	fmt.Fprint(&builder, tabbedString, "try {\n")
	for _, element := range tryBlock.Statements {
		element.SetNumTabs(tryBlock.tabs + 1)
		fmt.Fprint(&builder, element.String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, "}\n")

	for _, element := range tryBlock.CatchBlock {
		element.SetNumTabs(tryBlock.tabs)
		fmt.Fprint(&builder, element.String())
	}

	if len(tryBlock.FinallyStatements) > 0 {

		fmt.Fprint(&builder, tabbedString, "finally {\n")
		for _, element := range tryBlock.FinallyStatements {
			element.SetNumTabs(tryBlock.tabs + 1)
			fmt.Fprint(&builder, element.String(), "\n")
		}
		fmt.Fprint(&builder, tabbedString, "}\n")
	}
	return builder.String()
}
