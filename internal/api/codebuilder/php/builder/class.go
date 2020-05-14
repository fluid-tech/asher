package builder

import (
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type Class struct {
	interfaces.Class
	class core.Class
}

func NewClassBuilder() *Class {
	return &Class{
		class: core.Class{},
	}
}

/**
Sets the name of this class
*/
func (klass *Class) SetName(name string) interfaces.Class {
	klass.class.Name = name
	return klass
}

/**
Adds a TabbedUnit to this class
*/
func (klass *Class) AddMembers(units []*core.TabbedUnit) interfaces.Class {
	klass.class.Members = append(klass.class.Members, units...)
	return klass
}

/**
Sets the extends param of this class
*/
func (klass *Class) SetExtends(extendsClass string) interfaces.Class {
	klass.class.Extends = extendsClass
	return klass
}

/**
Adds a method to the class. ENSURE the first method in this list is the constructor
*/
func (klass *Class) AddFunction(function *core.Function) interfaces.Class {
	klass.class.Functions = append(klass.class.Functions, function)
	return klass
}

/**
Adds an interface to the implements field
*/
func (klass *Class) AddInterface(ifName string) interfaces.Class {
	klass.class.ImplementsList = append(klass.class.ImplementsList, ifName)
	return klass
}

/**
Adds imports to this class
*/
func (klass *Class) AddImport(imports []string) interfaces.Class {
	klass.class.Imports = append(klass.class.Imports, imports...)
	return klass
}

/**
Sets the package name for this class
*/
func (klass *Class) SetPackage(pkg string) interfaces.Class {
	klass.class.Package = pkg
	return klass
}

func (klass *Class) GetClass() core.Class {
	return klass.class
}
