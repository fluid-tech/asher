package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
)

const MutatorExtends = `BaseMutator`
const MutatorNamespace = `App\Transactors\Mutations`

type MutatorGenerator struct {
	api.Generator
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
Returns:
	- Return instance of MutatorGenerator
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (mutatorGenerator *MutatorGenerator) SetIdentifier(identifier string) *MutatorGenerator {
	mutatorGenerator.identifier = identifier
	return mutatorGenerator
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
Adds CallConstructor in the mutator
Returns:
	- Return instance of MutatorGenerator
Sample Usage:
	- mutatorGeneratorObject.AddConstructor()
*/
func (mutatorGenerator *MutatorGenerator) AddConstructor() *MutatorGenerator {

	parentConstructorCall := core.NewFunctionCall(CallParentConstructor).AddArg(core.NewParameter(
		fmt.Sprintf(`'App\%s', 'id'`, mutatorGenerator.identifier)))

	mutatorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility(VisibilityPublic).SetName(CallConstructor).
			AddStatement(parentConstructorCall).GetFunction())
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

	var className = fmt.Sprintf("%sMutator", mutatorGenerator.identifier)

	mutatorGenerator.AddConstructor()

	mutatorGenerator.classBuilder.SetName(className).
		SetExtends(MutatorExtends).SetPackage(MutatorNamespace).AddImports(mutatorGenerator.imports)
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
