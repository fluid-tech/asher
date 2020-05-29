package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type VarAssignment struct {
	api.TabbedUnit
	tabs       int
	Visibility string
	Identifier string
	Rhs        string
}

func NewVarAssignment(visibility string, id string, rhs string) *VarAssignment {
	return &VarAssignment{
		tabs:       0,
		Visibility: visibility,
		Identifier: id,
		Rhs:        rhs,
	}
}

func (v *VarAssignment) SetNumTabs(tabs int) {
	v.tabs = tabs
}

func (v *VarAssignment) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, api.TabbedString(uint(v.tabs)),
		v.Visibility, " $", v.Identifier, " = ", v.Rhs, ";")
	return builder.String()
}
