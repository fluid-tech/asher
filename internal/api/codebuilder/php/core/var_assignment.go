package core

import (
	"fmt"
	"strings"
)

type VarAssignment struct {
	TabbedUnit
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

func (v *VarAssignment) Id() string {
	return v.Identifier
}

func (v *VarAssignment) String() string  {
	var builder strings.Builder
	fmt.Fprintf(&builder, TabbedString(uint(v.tabs)),
		v.Visibility, " $", v.Identifier, " = ", v.Rhs, ";\n")
	return builder.String()
}
