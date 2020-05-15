package core

import (
	"fmt"
	"strings"
)


type Argument struct{
	TabbedUnit
	tabs int
	value string
}

func NewArgument(arg string) *Argument {
	return &Argument{
		value:  arg,
	}
}

func (a *Argument) SetNumTabs(tabs int) {
	a.tabs = tabs
}

func (a *Argument) Id() string {
	return a.value
}

func (a *Argument) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, TabbedString(uint(a.tabs)), a.value)
	return builder.String()
}