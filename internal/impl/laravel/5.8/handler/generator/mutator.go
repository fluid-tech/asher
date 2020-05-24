package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)

type MutatorGenerator struct {
	classBuilder interfaces.Class
	identifier   string
	imports      []string
}

func NewMutatorGenerator() *MutatorGenerator {
	return &MutatorGenerator{
		classBuilder: builder.NewClassBuilder(),
		identifier:   "",
		imports:      []string{},
	}
}

/**
Sets the identifier of the current class
Parameters:
	- identifier: string
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (mutatorGenerator *MutatorGenerator) SetIdentifier(identifier string) {
	mutatorGenerator.identifier = identifier
}

/**
Adds a Simple Statement
Parameters:
	- identifier: string
Returns:
	- Return instance of TabbedUnit
Sample Usage:
	- addSimpleStatement("Just A Simple Statement String")
*/
func (mutatorGenerator *MutatorGenerator) addSimpleStatement(identifier string) *api.TabbedUnit {
	statement := api.TabbedUnit(core.NewSimpleStatement(identifier))
	return &statement
}

/**
Adds a Parameter
Parameters:
	- identifier: string
Returns:
	- Return instance of TabbedUnit
Sample Usage:
	- addParameter("id")
*/
func (mutatorGenerator *MutatorGenerator) addParameter(identifier string) *api.TabbedUnit {
	parameter := api.TabbedUnit(core.NewParameter(identifier))
	return &parameter
}

/**
Appends import to the controller file
Parameters:
	- units: string array of the import
Returns:
	- instance of ControllerGenerator object
Sample Usage:
	- AppendImport([]string{"App\User",})
*/
func (mutatorGenerator *MutatorGenerator) AppendImports(imports []string) *MutatorGenerator {
	mutatorGenerator.imports = append(mutatorGenerator.imports, imports...)
	return mutatorGenerator
}

/**
Adds Constructor in the mutator
Returns:
	- Return instance of MutatorGenerator
Sample Usage:
	- mutatorGeneratorObject.AddConstructorFunction()
*/
func (mutatorGenerator *MutatorGenerator) AddConstructorFunction() *MutatorGenerator {

	parentConstructorCall := api.TabbedUnit(
		core.NewFunctionCall("parent::__construct").AddArg(core.NewParameter(
			`'App\` + mutatorGenerator.identifier + `', 'id'`)))

	constructorStatements := []api.TabbedUnit{
		parentConstructorCall,
	}

	mutatorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddStatements(constructorStatements).GetFunction())
	return mutatorGenerator
}

/**
Main Function To be called when we want to build the mutator
Returns:
	- Return instance of core.Class
Sample Usage:
	- mutatorGeneratorObject.BuildMutator()
*/
func (mutatorGenerator *MutatorGenerator) BuildMutator() *core.Class {
	var extends = `BaseMutator`
	var namespace = `App\Transactors\Mutations`
	className := mutatorGenerator.identifier + "Mutator"

	mutatorGenerator.AddConstructorFunction()

	mutatorGenerator.classBuilder.SetName(className).
		SetExtends(extends).SetPackage(namespace).AddImports(mutatorGenerator.imports)
	return mutatorGenerator.classBuilder.GetClass()
}

/**
Returns:
	- Return string object of MutatorGenerator
Sample Usage:
	- mutatorGeneratorObject.String()
*/
func (mutatorGenerator *MutatorGenerator) String() string {
	return mutatorGenerator.BuildMutator().String()
}
