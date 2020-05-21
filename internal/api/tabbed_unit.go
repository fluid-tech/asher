package api

import "strings"

type TabbedUnit interface {
	SetNumTabs(tabs int)
	Id() string // something that uniquely identifies this tabbed unit, in cases of vars it could be the var name
	String() string
}

func TabbedString(numTabs uint) string {
	if numTabs > 0 {
		return strings.Repeat(" ", int(numTabs<<2))
	}
	return ""
}
