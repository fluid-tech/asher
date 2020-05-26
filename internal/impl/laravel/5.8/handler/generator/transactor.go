package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

const namespace = `App\Transactors`

//TEMP VARIABLE
var queryVariableName string
var mutatorVariableName string
var superConstructorCall *core.FunctionCall

type TransactorGenerator struct {
	classBuilder   interfaces.Class
	identifier     string
	imports        []string
	transactorType string
	parentConstructorCallArgs []api.TabbedUnit
	constructorStatements []api.TabbedUnit
	transactorMembers []api.TabbedUnit
}

func NewTransactorGenerator(identifier string, transactorType string) *TransactorGenerator {
	return &TransactorGenerator{
		classBuilder:   builder.NewClassBuilder(),
		identifier:     identifier,
		imports:        []string{},
		transactorType: transactorType,
		parentConstructorCallArgs: []api.TabbedUnit{},
		constructorStatements: []api.TabbedUnit{},
		transactorMembers: []api.TabbedUnit{},
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

func (transactorGenerator *TransactorGenerator) AddParentConstructorCallArgs(
	parameter *core.Parameter) *TransactorGenerator{

	transactorGenerator.parentConstructorCallArgs =append(transactorGenerator.parentConstructorCallArgs, parameter)
	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) AddTransactorMember(
	member api.TabbedUnit) *TransactorGenerator{

	transactorGenerator.transactorMembers =append(transactorGenerator.transactorMembers, member)
	return transactorGenerator
}






/**
Sets the type of the transactor
Parameters:
	- identifier: string
Sample Usage:
	- SetTransactorType("default") or SetTransactorType("file") or SetTransactorType("image")
*/
func (transactorGenerator *TransactorGenerator) SetTransactorType(identifier string) *TransactorGenerator {
	transactorGenerator.transactorType = identifier
	return transactorGenerator
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

func (transactorGenerator *TransactorGenerator)addDefaults() *TransactorGenerator{
	/*Default Imports*/
	transactorImports := []string{
		`App\Query\` + transactorGenerator.identifier + `Query`,
		`App\Transactors\Mutations\` + transactorGenerator.identifier + `Mutator`,
	}
	transactorGenerator.imports = append(transactorImports, transactorGenerator.imports...)


	/*Default CLASS MEMBERS*/
	className := transactorGenerator.identifier + "Transactor"
	transactorGenerator.transactorMembers = append([]api.TabbedUnit{core.NewSimpleStatement(
		"private const CLASS_NAME = '" + className + "'")},
		transactorGenerator.transactorMembers...)

	/*Default parent Constructor CALL Arguments*/
	lowerCamelIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)
	queryVariableName = lowerCamelIdentifier + `Query`
	mutatorVariableName = lowerCamelIdentifier + `Mutator`

	transactorGenerator.parentConstructorCallArgs = append([]api.TabbedUnit{
		core.NewParameter("$"+queryVariableName),
		core.NewParameter("$"+mutatorVariableName),
		core.NewParameter(`"id"`)},
		transactorGenerator.parentConstructorCallArgs...
	)


	/*DEFAULT CONSTRUCTOR STATEMENTS*/
	superConstructorCall = core.NewFunctionCall("parent::__construct")

	transactorGenerator.constructorStatements = append(transactorGenerator.constructorStatements,
		superConstructorCall,
		core.NewSimpleStatement("$this->className = self::CLASS_NAME"),
		)

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

	transactorGenerator.addDefaults()

	className := transactorGenerator.identifier + "Transactor"
	transactorGenerator.classBuilder.SetName(className).SetPackage(namespace).
		SetExtends(strcase.ToCamel(transactorGenerator.transactorType)+"Transactor")

	/*IMPORTS*/
	transactorGenerator.classBuilder.AddImports(transactorGenerator.imports)

	/*MEMBERS*/
	transactorGenerator.classBuilder.AddMembers(transactorGenerator.transactorMembers)

	/*CALL TO SUPER CONSTRUCTOR*/
	superConstructorCall.AddArgs(transactorGenerator.parentConstructorCallArgs)

	/*CONSTRUCTOR*/

	constructorFuncBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArguments([]string{transactorGenerator.identifier + `Query $` + queryVariableName,
		transactorGenerator.identifier + `Mutator $` + mutatorVariableName}).
		AddStatements(transactorGenerator.constructorStatements)

	transactorGenerator.classBuilder.AddFunction(constructorFuncBuilder.GetFunction())


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
