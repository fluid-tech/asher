package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type ReturnStatement struct {
	api.TabbedUnit
	tabs      int
	statement string
}

func NewReturnStatement(stmt string) *ReturnStatement {
	return &ReturnStatement{
		statement: stmt,
	}
}

func (r *ReturnStatement) SetNumTabs(tabs int) {
	r.tabs = tabs
}


func (r *ReturnStatement) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(r.tabs)), "return ", r.statement, ";")
	return builder.String()
}
