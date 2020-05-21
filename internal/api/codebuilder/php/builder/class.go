package builder

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type Class struct {
	interfaces.Class
	class *core.Class
}

/**
Creates a new instance of this builder, with a new core.Class
*/
func NewClassBuilder() *Class {
	return &Class{
		class: core.NewClass(),
	}
}

/**
Creates a new instance of this builder, with an existing core.Class
*/
func NewClassBuilderFromClass(class *core.Class) *Class {
	return &Class{
		class: class,
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
Add a TabbedUnit to this this class's members list
*/
func (klass *Class) AddMember(unit api.TabbedUnit) interfaces.Class {
	return klass.AddMembers([]api.TabbedUnit{unit})
}

/**
Adds a list of TabbedUnit to this class's members list
*/
func (klass *Class) AddMembers(units []api.TabbedUnit) interfaces.Class {
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
func (klass *Class) AddImports(imports []string) interfaces.Class {
	klass.class.Imports = append(klass.class.Imports, imports...)
	return klass
}

/**
Add import to this class
*/
func (klass *Class) AddImport(importPkg string) interfaces.Class {
	return klass.AddImports([]string{importPkg})
}

/**
Sets the package name for this class
*/
func (klass *Class) SetPackage(pkg string) interfaces.Class {
	klass.class.Package = pkg
	return klass
}

func (klass *Class) GetClass() *core.Class {
	return klass.class
}
