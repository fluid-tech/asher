package core

import (
	"asher/internal/api"
	"errors"
	"fmt"
	"strings"
)

type Class struct {
	api.TabbedUnit
	Tabs           int
	Name           string
	Extends        string
	ImplementsList []string
	Members        []*api.TabbedUnit
	Functions      []*Function
	Package        string
	Imports        []string
}

func NewClass() *Class {
	return &Class{
		Tabs:           0,
		Name:           "",
		Extends:        "",
		ImplementsList: []string{},
		Members:        []*api.TabbedUnit{},
		Functions:      []*Function{},
		Package:        "",
		Imports:        []string{},
	}
}

func (klass *Class) Id() string {
	return klass.Name
}

func (klass *Class) SetNumTabs(tabs int) {
	klass.Tabs = tabs
}

func (klass Class) String() string {
	var sb strings.Builder
	klass.handlePackage(&sb)
	klass.handleImports(&sb)
	fmt.Fprint(&sb, api.TabbedString(uint(klass.Tabs)))
	fmt.Fprint(&sb, "class ", klass.Name)
	klass.handleExtends(&sb)
	klass.handleImplementsList(&sb)
	fmt.Fprint(&sb, " {\n")
	klass.handleMembers(&sb)
	klass.handleFunctions(&sb)
	fmt.Fprint(&sb,  api.TabbedString(uint(klass.Tabs)), "}\n")

	return sb.String()
}

/**
Searches for the identifier in the members list
*/
func (klass *Class) FindMember(identifier string) (*api.TabbedUnit, error) {
	for _, element := range klass.Members {
		if (*element).Id() == identifier {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldnt find member with identifier %s", identifier))
}

/**
Searches for the identifier in the functions list
*/
func (klass *Class) FindFunction(identifier string) (*Function, error) {
	for _, element := range klass.Functions {
		if element.Id() == identifier {
			return element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldnt find function with identifier %s", identifier))
}

/**
Appends a tabbed unit to the members list
 */
func (klass *Class) AppendMember(unit *api.TabbedUnit){
	klass.Members = append(klass.Members, unit)
}

func (klass *Class) handlePackage(builder *strings.Builder) {
	if klass.Package != "" {
		fmt.Fprint(builder, "namespace ", klass.Package, ";\n\n")
	}
}

func (klass *Class) handleImports(builder *strings.Builder) {
	if len(klass.Imports) > 0 {
		for _, element := range klass.Imports {
			fmt.Fprint(builder, "use ", element, ";\n")
		}
		fmt.Fprint(builder, "\n")
	}
}

func (klass *Class) handleExtends(builder *strings.Builder) {
	if klass.Extends != "" {
		fmt.Fprint(builder, " extends ", klass.Extends)
	}
}

func (klass *Class) handleMembers(builder *strings.Builder) {
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

func (klass *Class) handleImplementsList(builder *strings.Builder) {
	if len(klass.ImplementsList) > 0 {
		_, err := fmt.Fprint(builder, " implements ", strings.Join(klass.ImplementsList[:], ", "))
		if err != nil {
			fmt.Println("Error encounted ", err)
		}
	}
}

func (klass *Class) handleFunctions(builder *strings.Builder) {
	if len(klass.Functions) > 0 {
		for _, element := range klass.Functions {
			element.SetNumTabs(klass.Tabs + 1)
			fmt.Fprint(builder, element, "\n")
		}
	}
}
