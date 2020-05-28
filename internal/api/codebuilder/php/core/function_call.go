package core

import (
	"asher/internal/api"
	"fmt"
	"strings"
)

type FunctionCall struct {
	api.TabbedUnit
	tabs int
	Def  string
	Args []api.TabbedUnit
}

func NewFunctionCall(def string) *FunctionCall {
	return &FunctionCall{
		Def:  def,
		Args: []api.TabbedUnit{},
	}
}

func (c *FunctionCall) SetNumTabs(tabs int) {
	c.tabs = tabs
}

/**
Adds a tabbed unit to the args list
Returns the current instance so that you can chain it
*/
func (c *FunctionCall) AddArg(unit api.TabbedUnit) *FunctionCall {
	c.Args = append(c.Args, unit)
	return c
}

func (c *FunctionCall) AddArgs(unit []api.TabbedUnit) *FunctionCall {
	c.Args = append(c.Args, unit...)
	return c
}

func (c *FunctionCall) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, api.TabbedString(uint(c.tabs)), c.Def, "(")
	argLen := len(c.Args)
	for i, element := range c.Args {
		fmt.Fprintf(&builder, element.String())
		if i != argLen-1 {
			fmt.Fprint(&builder, ", ")
		}
	}
	fmt.Fprintf(&builder, ");")
	return builder.String()
}
