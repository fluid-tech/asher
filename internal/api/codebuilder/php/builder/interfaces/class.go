package interfaces

import "asher/internal/api/codebuilder/php/core"

type Class interface {
	SetName(className string) Class
	AddMembers(members []*core.TabbedUnit) Class
	AddMember(member *core.TabbedUnit) Class
	SetExtends(extendsClass string) Class
	AddFunction(function *core.Function) Class
	AddFunctions(functions []*core.Function) Class
	AddInterface(ifName string) Class
	AddImports(imports []string) Class
	AddImport(imports string) Class
	SetPackage(pkg string) Class
	GetClass() *core.Class
}