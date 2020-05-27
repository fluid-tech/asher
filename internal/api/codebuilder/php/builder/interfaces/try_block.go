package interfaces

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
)

type TryBlock interface {

	//Appends a statement in the statement list of try
	AddStatement(statement api.TabbedUnit) TryBlock

	//Appends statement in the statement list of try
	AddStatements(statements []api.TabbedUnit) TryBlock

	//Appends a finally statement in the finally statement list of try
	AddFinallyStatement(statement api.TabbedUnit) TryBlock

	//Appends finally statement in the finally statement list of try
	AddFinallyStatements(statement []api.TabbedUnit) TryBlock

	//Appends a catchBlock in the catch block list of try
	AddCatchBlock(block *core.CatchBlock) TryBlock

	// returns the tryblock object
	GetTryBlock() *core.TryBlock
}
