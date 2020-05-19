package helper

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
)


/**
Converts SimpleStatement to TabbedUnit
*/
func TabbedUnitForSimpleStatement(unit *core.SimpleStatement) *api.TabbedUnit  {
	statement:= api.TabbedUnit(unit)
	return &statement
}

/**
Converts Function to TabbedUnit
*/
func TabbedUnitForFunction(unit *core.Function) *api.TabbedUnit {
	functionStatement := api.TabbedUnit(unit)
	return &functionStatement
}

/**
Converts VarDeclaration to TabbedUnit
*/
func TabbedUnitForVarDeclaration(unit *core.VarDeclaration) *api.TabbedUnit {
	varDeclarationStatement := api.TabbedUnit(unit)
	return &varDeclarationStatement
}


/**
Converts ArrayAssignment to TabbedUnit
*/
func TabbedUnitForArrayAssignment(unit *core.ArrayAssignment) *api.TabbedUnit {
	arrayAssignmentStatement := api.TabbedUnit(unit)
	return &arrayAssignmentStatement
}

/**
Converts PhpEmitterFile to EmitterFile
*/
func EmitterFileForPhpEmitterFiler(unit *core.PhpEmitterFile) *api.EmitterFile  {
	unitRef := api.EmitterFile(unit)
	return &unitRef
}


/**
Converts FunctionCall to TabbedUnit
*/
func TabbedUnitForFunctionCall(unit *core.FunctionCall) *api.TabbedUnit  {
	functionCallStatement := api.TabbedUnit(unit)
	return &functionCallStatement
}

/**
Converts ReturnStatement to TabbedUnit
*/
func TabbedUnitForReturnStatement(unit *core.ReturnStatement) *api.TabbedUnit {
	returnStatement := api.TabbedUnit(unit)
	return &returnStatement
}

/**
Converts Parameter to TabbedUnit
*/
func TabbedUnitForParameter(unit *core.Parameter) *api.TabbedUnit  {
	parameter := api.TabbedUnit(unit)
	return &parameter
}

