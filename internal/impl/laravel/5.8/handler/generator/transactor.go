package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

type TransactorGenerator struct {
	classBuilder interfaces.Class
	identifier   string
	imports      []string
}

func NewTransactorGenerator() *TransactorGenerator {
	return &TransactorGenerator{
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
func (transactorGenerator *TransactorGenerator) SetIdentifier(identifier string) *TransactorGenerator {
	transactorGenerator.identifier = identifier
	return transactorGenerator
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
func (transactorGenerator *TransactorGenerator) addSimpleStatement(identifier string) *api.TabbedUnit {
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
func (transactorGenerator *TransactorGenerator) addParameter(identifier string) *api.TabbedUnit {
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
func (transactorGenerator *TransactorGenerator) AppendImports(units []string) *TransactorGenerator {
	transactorGenerator.imports = append(transactorGenerator.imports, units...)
	return transactorGenerator
}


/**
Adds Constructor in the Transactor with Query and Mutator Injected of the currentModel
Returns:
	- Return instance of TransactorGenerator
Sample Usage:
	- transactorGeneratorObject.AddConstructorFunction()
*/
func (transactorGenerator *TransactorGenerator) AddConstructorFunction() *TransactorGenerator {
	lowerCamelIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)
	queryVariableName := lowerCamelIdentifier + `Query`
	mutatorVariableName := lowerCamelIdentifier + `Mutator`

	constructorArguments := []string{
		transactorGenerator.identifier + `Query $` + queryVariableName,
		transactorGenerator.identifier + `Mutator $` + mutatorVariableName,
	}

	parentConstructorCall := api.TabbedUnit(core.NewFunctionCall("parent::__construct").
		AddArg(transactorGenerator.addParameter("$" + queryVariableName + ", $" + mutatorVariableName + ", 'id'")))

	constructorStatements := []*api.TabbedUnit{
		&parentConstructorCall,
		transactorGenerator.addSimpleStatement("$this->className = self::CLASS_NAME"),
	}

	transactorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return transactorGenerator
}

/**
Main Function To be called when we want to build the transactor
Returns:
	- Return instance of core.Class
Sample Usage:
	- transactorGeneratorObject.BuildRestTransactor()
*/
func (transactorGenerator *TransactorGenerator) BuildTransactor() *core.Class {
	const extends = `BaseTransactor`
	const namespace = `App\Transactors`

	transactorImports := []string{
		`App\Query\` + transactorGenerator.identifier + `Query`,
		`App\Transactors\Mutations\` + transactorGenerator.identifier + `Mutator`,
	}

	className := transactorGenerator.identifier + "Transactor"
	transactorGenerator.AppendImports(transactorImports)
	transactorGenerator.AddConstructorFunction()

	transactorGenerator.classBuilder.AddMember(transactorGenerator.addSimpleStatement(
		"private const CLASS_NAME = '" + className + "'")).SetName(className).
		SetExtends(extends).SetPackage(namespace).AddImports(transactorGenerator.imports)
	return transactorGenerator.classBuilder.GetClass()
}

/**
Returns:
	- Return string object of TransactorGenerator
Sample Usage:
	- transactorGeneratorObject.String()
*/
func (transactorGenerator *TransactorGenerator) String() string {
	return transactorGenerator.BuildTransactor().String()
}
