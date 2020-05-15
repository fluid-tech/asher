package core

import (
	"fmt"
	"strings"
)

type ReturnStatement struct{
	TabbedUnit
	tabs int
	statement string
}

func NewReturnStatement(stmt string) *ReturnStatement {
	return &ReturnStatement{
		statement:  stmt,
	}
}

func (r *ReturnStatement) SetNumTabs(tabs int) {
	r.tabs = tabs
}

func (r *ReturnStatement) Id() string {
	return "return"
}

func (r *ReturnStatement) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, TabbedString(uint(r.tabs)), "return ", r.statement, ";\n")
	return builder.String()
}