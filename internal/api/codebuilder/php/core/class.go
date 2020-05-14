package core

import (
	"errors"
	"fmt"
	"strings"
)

type Class struct {
	TabbedUnit
	Tabs           int
	Name           string
	Extends        string
	ImplementsList []string
	Members        []*TabbedUnit
	Functions      []*Function
	Package        string
	Imports        []string
}

func (klass Class) Id() string {
	return klass.Name
}

func (klass Class) SetNumTabs(tabs int) {
	klass.Tabs = tabs
}

func (klass Class) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, TabbedString(uint(klass.Tabs)))
	fmt.Fprint(&sb, "class ", klass.Name, " extends ", klass.Extends)
	klass.handleImplementsList(&sb)
	fmt.Fprint(&sb, " { \n")
	klass.handleMembers(&sb)
	klass.handleFunctions(&sb)
	fmt.Fprint(&sb, "\n\n", TabbedString(uint(klass.Tabs)), "}")

	return sb.String()
}

func (klass Class) handleMembers(builder *strings.Builder) {
	if len(klass.Members) > 0 {
		for _, element := range klass.Members {
			(*element).SetNumTabs(klass.Tabs + 1)
			_, err := fmt.Fprint(builder, (*element).String(), "\n")
			if err != nil {
				fmt.Println("Error encounted ", err)
			}
		}
	}
}

func (klass Class) handleImplementsList(builder *strings.Builder) {
	if len(klass.ImplementsList) > 0 {
		_, err := fmt.Fprint(builder, " implements ", strings.Join(klass.ImplementsList[:], ", "))
		if err != nil {
			fmt.Println("Error encounted ", err)
		}
	}
}

func (klass Class) handleFunctions(builder *strings.Builder) {
	if len(klass.Functions) > 0 {
		for _, element := range klass.Functions {
			element.SetNumTabs(klass.Tabs + 1)
			_, err := fmt.Fprint(builder, element, "\n")
			if err != nil {
				fmt.Println("Error encounted ", err)
			}
		}
	}
}

func (klass *Class) FindInMembers(identifier string) (*TabbedUnit, error) {
	for _, element := range klass.Members {
		if (*element).Id() == identifier {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldnt find member with identifier %s", identifier))
}

func (klass *Class) FindInFunctions(identifier string) (*Function, error) {
	for _, element := range klass.Functions {
		if element.Id() == identifier {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldnt find function with identifier %s", identifier))
}
