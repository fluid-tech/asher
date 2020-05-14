package core

import (
	"fmt"
	"strings"
)

type VarDeclaration struct {
	TabbedUnit
	tabs       int
	Visibility string
	Identifier string
}

func GetVarDeclaration(visibility string, id string) *VarDeclaration {
	return &VarDeclaration{
		tabs:       0,
		Visibility: visibility,
		Identifier: id,
	}
}


func (v *VarDeclaration) SetNumTabs(tabs int) {
	v.tabs = tabs
}

func (v *VarDeclaration) Id() string {
	return v.Identifier
}

func (v *VarDeclaration) String() string  {
	var builder strings.Builder
	fmt.Fprintf(&builder, TabbedString(uint(v.tabs)),
		v.Visibility, " $", v.Identifier, ";\n")
	return builder.String()
}