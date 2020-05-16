package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type SimpleStatement struct {
	api.TabbedUnit
	numTabs         int
	SimpleStatement string
}

func NewSimpleStatement(simpleStatement string) *SimpleStatement {
	return &SimpleStatement{
		SimpleStatement: simpleStatement,
	}
}

func (stmt *SimpleStatement) SetNumTabs(tabs int) {
	stmt.numTabs = tabs
}

func (stmt *SimpleStatement) Id() string {
	return stmt.SimpleStatement
}

func (stmt *SimpleStatement) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(stmt.numTabs)), stmt.SimpleStatement, ";")
	return builder.String()
}
