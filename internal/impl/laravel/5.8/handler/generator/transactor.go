package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)



type TransactorGenerator struct {
	classBuilder		interfaces.Class
	identifier 			string
	imports				[]string
}

func NewTransactorGenerator() *TransactorGenerator  {
	return &TransactorGenerator{
		classBuilder: builder.NewClassBuilder(),
		identifier: "",
		imports: []string{},
	}
}

func (transactorGenerator *TransactorGenerator) SetIdentifier(identifier string) *TransactorGenerator {
	transactorGenerator.identifier = identifier
	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) AddSimpleStatement(identifier string) *api.TabbedUnit  {
	statement:= api.TabbedUnit(core.NewSimpleStatement(identifier))
	return &statement
}

func (transactorGenerator *TransactorGenerator) AddParameter(identifier string) *api.TabbedUnit  {
	parameter := api.TabbedUnit(core.NewParameter(identifier))
	return &parameter
}

func (transactorGenerator *TransactorGenerator) AddConstructorFunction() *TransactorGenerator  {
	lowerCamelIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)
	queryVariableName := lowerCamelIdentifier+`Query`
	mutatorVariableName := lowerCamelIdentifier+`Mutator`

	constructorArguments := []string{
		transactorGenerator.identifier+`Query $`+queryVariableName,
		transactorGenerator.identifier+`Mutator $`+mutatorVariableName,
	}

	parentConstructorCall := api.TabbedUnit(core.NewFunctionCall("parent::__construct").
		AddArg(transactorGenerator.AddParameter(
			"$"+queryVariableName+", $"+mutatorVariableName+", 'id'")))

	constructorStatements := []*api.TabbedUnit{
		&parentConstructorCall,
		transactorGenerator.AddSimpleStatement("$this->className = self::CLASS_NAME"),

	}


	transactorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) AppendImports(imports []string)  *TransactorGenerator{
	transactorGenerator.imports = append(transactorGenerator.imports, imports...)
	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) BuildTransactor() *core.Class  {
	const extends = `BaseTransactor`
	const namespace = `App\Transactors`

	transactorImports := []string{
		`App\Query\`+transactorGenerator.identifier+`Query`,
		`App\Transactors\Mutations\`+transactorGenerator.identifier+`Mutator`,
	}

	className := transactorGenerator.identifier+"Transactor"
	transactorGenerator.AppendImports(transactorImports)
	transactorGenerator.AddConstructorFunction()

	transactorGenerator.classBuilder.AddMember(transactorGenerator.AddSimpleStatement(
		"private const CLASS_NAME = '"+className+"'")).SetName(className).
		SetExtends(extends).SetPackage(namespace).AddImports(transactorGenerator.imports)
	return transactorGenerator.classBuilder.GetClass()
}

func (transactorGenerator *TransactorGenerator) String() string {
	return transactorGenerator.BuildTransactor().String()
}