package builder

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type Function struct {
	interfaces.Function
	function *core.Function
}

/**
Returns a Function Builder instance
*/
func NewFunctionBuilder() *Function {
	return &Function{
		function: core.NewFunction(),
	}
}

func NewFunctionBuilderFromFunction(function *core.Function) *Function {
	return &Function{
		function: function,
	}
}

func (f *Function) SetName(name string) interfaces.Function {
	f.function.Name = name
	return f
}

func (f *Function) AddArgument(name string) interfaces.Function {
	return f.AddArguments([]string{name})
}

func (f *Function) AddArguments(args []string) interfaces.Function {
	f.function.Arguments = append(f.function.Arguments, args...)
	return f
}

func (f *Function) AddStatement(statement *api.TabbedUnit) interfaces.Function {
	return f.AddStatements([]*api.TabbedUnit{statement})
}

func (f *Function) AddStatements(statements []*api.TabbedUnit) interfaces.Function {
	f.function.Statements = append(f.function.Statements, statements...)
	return f
}

func (f *Function) SetVisibility(vis string) interfaces.Function {
	f.function.Visibility = vis
	return f
}

func (f *Function) GetFunction() *core.Function {
	return f.function
}
