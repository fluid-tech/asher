package core

import (
	"fmt"
	"strings"
)

type ArrayAssignment struct {
	TabbedUnit
	tabs       int
	Visibility string
	Identifier string
	Rhs        []string
}

func (arr *ArrayAssignment) SetNumTabs(tabs int) {
	arr.tabs = tabs
}

func (arr *ArrayAssignment) Id() string {
	return arr.Identifier
}

func (arr *ArrayAssignment) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, TabbedString(uint(arr.tabs)),
		arr.Visibility, " ", arr.Identifier, " = ", strings.Join(arr.Rhs, ", \n"), ";\n")
	return builder.String()
}
