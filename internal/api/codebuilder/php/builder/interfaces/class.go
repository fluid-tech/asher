package interfaces

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
)

type Class interface {
	SetName(className string) Class
	AddMembers(members []*api.TabbedUnit) Class
	AddMember(member *api.TabbedUnit) Class
	SetExtends(extendsClass string) Class
	AddFunction(function *core.Function) Class
	AddInterface(ifName string) Class
	AddImports(imports []string) Class
	AddImport(imports string) Class
	SetPackage(pkg string) Class
	GetClass() *core.Class
}