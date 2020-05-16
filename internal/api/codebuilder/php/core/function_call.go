package core

import (
	"errors"
	"fmt"
	"strings"
)

type FunctionCall struct {
	TabbedUnit
	tabs int
	Def  string
	Args []*TabbedUnit
}

func NewFunctionCall(def string) *FunctionCall {
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

func (c *FunctionCall) FindById(id string) (*TabbedUnit, error) {
	for _, element := range c.Args {
		if  (*element).Id() == id {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("cant find statement with identifier %s", id))
}

func (c *FunctionCall) String() string {
	var builder strings.Builder
	fmt.Fprint(&builder, TabbedString(uint(c.tabs)), c.Def, "(")
	argLen := len(c.Args)
	for i, element := range c.Args{
		//fmt.Println(reflect.ValueOf((*element).(*Function)).)

		fmt.Fprintf(&builder , (*element).String())
		if i != argLen - 1 {
			fmt.Fprint(&builder, ", ")
		}
	}
	fmt.Fprintf(&builder, ");")
	return builder.String()
}
