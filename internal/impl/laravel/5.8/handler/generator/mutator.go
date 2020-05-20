package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
)



type MutatorGenerator struct {
	classBuilder		interfaces.Class
	identifier 			string
	imports				[]string
}

func NewMutatorGenerator() *MutatorGenerator {
	return &MutatorGenerator{
		classBuilder: builder.NewClassBuilder(),
		identifier: "",
		imports: []string{},
	}
}

func (mutatorGenerator *MutatorGenerator) SetIdentifier(identifier string)  {
	mutatorGenerator.identifier = identifier
}

func (mutatorGenerator *MutatorGenerator) AddSimpleStatement(identifier string) *api.TabbedUnit  {
	statement:= api.TabbedUnit(core.NewSimpleStatement(identifier))
	return &statement
}

func (mutatorGenerator *MutatorGenerator) AddParameter(identifier string) *api.TabbedUnit  {
	parameter := api.TabbedUnit(core.NewParameter(identifier))
	return &parameter
}

func (mutatorGenerator *MutatorGenerator) AppendImports(imports []string)  *MutatorGenerator {
	mutatorGenerator.imports = append(mutatorGenerator.imports, imports...)
	return mutatorGenerator
}

func (mutatorGenerator *MutatorGenerator) AddConstructorFunction() *MutatorGenerator {
	parentConstructorCall := api.TabbedUnit(core.NewFunctionCall("parent::__construct").
		AddArg(mutatorGenerator.AddParameter(
			`'App\`+mutatorGenerator.identifier+`', 'id'`)))

	constructorStatements := []*api.TabbedUnit{
		&parentConstructorCall,
	}

	mutatorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddStatements(constructorStatements).GetFunction())
	return mutatorGenerator
}



func (mutatorGenerator *MutatorGenerator) BuildMutator() *core.Class  {
	var extends = `BaseMutator`
	var namespace = `App\Transactors\Mutations`
	className := mutatorGenerator.identifier+"Mutator"

	mutatorGenerator.AddConstructorFunction()

	mutatorGenerator.classBuilder.SetName(className).
		SetExtends(extends).SetPackage(namespace).AddImports(mutatorGenerator.imports)
	return mutatorGenerator.classBuilder.GetClass()
}

func (mutatorGenerator *MutatorGenerator) String() string {
	return mutatorGenerator.BuildMutator().String()
}