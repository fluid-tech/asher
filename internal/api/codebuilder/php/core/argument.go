package core


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
	return a.value
}