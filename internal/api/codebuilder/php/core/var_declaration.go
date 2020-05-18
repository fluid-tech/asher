package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type VarDeclaration struct {
	api.TabbedUnit
	tabs       int
	Visibility string
	Identifier string
}

func NewVarDeclaration(visibility string, id string) *VarDeclaration {
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

func (v *VarDeclaration) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(v.tabs)),
		v.Visibility, " $", v.Identifier, ";")
	return builder.String()
}
