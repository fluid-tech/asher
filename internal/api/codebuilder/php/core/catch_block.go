package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type CatchBlock struct {
	api.TabbedUnit
	tabs            int
	CatchArg        string
	CatchStatements []*api.TabbedUnit
}

func NewCatchBlock() *CatchBlock {
	return &CatchBlock{
		tabs:            0,
		CatchArg:        "",
		CatchStatements: []*api.TabbedUnit{},
	}
}

func (catchBlock *CatchBlock) AddArgument(unit string) *CatchBlock {
	catchBlock.CatchArg = unit
	return catchBlock
}

func (catchBlock *CatchBlock) AddStatement(statement *api.TabbedUnit) *CatchBlock {
	catchBlock.CatchStatements = append(catchBlock.CatchStatements, statement)
	return catchBlock
}

func (catchBlock *CatchBlock) AddStatements(statements []*api.TabbedUnit) *CatchBlock {
	catchBlock.CatchStatements = append(catchBlock.CatchStatements, statements...)
	return catchBlock
}

func (catchBlock *CatchBlock) Id() string {
	return "catch"
}

func (catchBlock *CatchBlock) SetNumTabs(tabs int) {
	catchBlock.tabs = tabs
}

func (catchBlock *CatchBlock) String() string {
	var builder strings.Builder
	tabbedString := api.TabbedString(uint(catchBlock.tabs))
	fmt.Fprint(&builder, tabbedString, " catch ( ")
	fmt.Fprint(&builder, catchBlock.CatchArg+" "+") { \n")
	for _, element := range catchBlock.CatchStatements {
		(*element).SetNumTabs(catchBlock.tabs)
		fmt.Fprint(&builder, tabbedString, (*element).String(), "\n")
	}
	fmt.Fprint(&builder, tabbedString, " }\n")
	return builder.String()
}
