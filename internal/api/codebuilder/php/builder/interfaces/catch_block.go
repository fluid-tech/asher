package interfaces

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
)

type CatchBlock interface {

	//Adds an argument in the catch
	AddArgument(unit string) CatchBlock

	//Appends a statement in the statement list of try
	AddStatement(statement api.TabbedUnit) CatchBlock

	//Appends statement in the statement list of try
	AddStatements(statements []api.TabbedUnit) CatchBlock

	// returns the catchBlock object
	GetCatchBlock() *core.CatchBlock
}
