package core

import (
	"fmt"
	"strings"
)

type SimpleStatement struct {
	TabbedUnit
	numTabs         int
	SimpleStatement string
}

func GetSimpleStatement(simpleStatement string) *SimpleStatement {
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
	fmt.Fprint(&builder, TabbedString(uint(stmt.numTabs)), stmt.SimpleStatement, ";\n")
	return builder.String()
}
