package interfaces

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
)

type Function interface {
	// sets the name of the function
	SetName(name string) Function

	//adds an arg to the args list array
	AddArgument(name string) Function

	// appends args list arr with given array
	AddArguments(args []string) Function

	// adds a statement to the statement list
	AddStatement(unit api.TabbedUnit) Function

	// appends statements to the statements arr
	AddStatements(units []api.TabbedUnit) Function

	// Sets the visibility of the method
	SetVisibility(vis string) Function

	// returns the function
	GetFunction() *core.Function
}
