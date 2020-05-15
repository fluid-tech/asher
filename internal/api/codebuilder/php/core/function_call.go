package core

import (
	"fmt"
	"strings"
)

type FunctionCall struct {
	TabbedUnit
	tabs int
	Def  string
	Args []*TabbedUnit
}

func NewClosure(def string) *FunctionCall {
	return &FunctionCall{
		Def:        def,
		Args:       []*TabbedUnit{},
	}
}

func (c *FunctionCall) SetNumTabs(tabs int) {
	c.tabs = tabs
}

func (c *FunctionCall) Id() string {
	return c.Def
}

/**
Adds a tabbed unit to the args list
Returns the current instance so that you can chain it
 */
func (c *FunctionCall) AddArg(unit *TabbedUnit) *FunctionCall {
	c.Args = append(c.Args, unit)
	return c
}

func (c *FunctionCall) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, TabbedString(uint(c.tabs)), c.Def, "(")
	argLen := len(c.Args)
	for i, element := range c.Args{
		(*element).SetNumTabs(c.tabs + 1)
		fmt.Fprintf(&builder , (*element).String())
		if i != argLen - 1 {
			fmt.Fprint(&builder, ", ")
		}
		fmt.Fprintf(&builder, "\n")
	}
	fmt.Fprintf(&builder, ");")
	return builder.String()
}
