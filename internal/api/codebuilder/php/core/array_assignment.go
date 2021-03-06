package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type ArrayAssignment struct {
	api.TabbedUnit
	tabs       int
	Visibility string
	Identifier string
	Rhs        []string
}

func NewArrayAssignment(visibility string, identifier string, rhs []string) *ArrayAssignment {
	return &ArrayAssignment{
		Visibility: visibility,
		Identifier: identifier,
		Rhs:        rhs,
	}
}

func (arr *ArrayAssignment) SetNumTabs(tabs int) {
	arr.tabs = tabs
}

func (arr *ArrayAssignment) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(arr.tabs)),
		arr.Visibility, " $", arr.Identifier, " = [", strings.Join(arr.Rhs, ", \n"), "\n];\n")
	return builder.String()
}
