package api

import "strings"

type TabbedUnit interface {
	SetNumTabs(tabs int)
	String() string
}

func TabbedString(numTabs uint) string {
	if numTabs > 0 {
		return strings.Repeat(" ", int(numTabs<<2))
	}
	return ""
}
