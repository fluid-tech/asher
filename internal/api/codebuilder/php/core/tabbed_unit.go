package php

import "strings"

type TabbedUnit interface {
	ToString() string
}

func TabbedString(numTabs uint) string  {
	return strings.Repeat(" ", int(numTabs << 4))
}
