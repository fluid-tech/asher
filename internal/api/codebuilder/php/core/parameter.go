package core

import "asher/internal/api"

type Parameter struct {
	api.TabbedUnit
	tabs  int
	value string
}

func NewParameter(arg string) *Parameter {
	return &Parameter{
		value: arg,
	}
}

func (a *Parameter) SetNumTabs(tabs int) {
	a.tabs = tabs
}

func (a *Parameter) Id() string {
	return a.value
}

func (a *Parameter) String() string {
	return a.value
}
