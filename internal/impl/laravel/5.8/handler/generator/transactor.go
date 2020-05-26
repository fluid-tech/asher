package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

const namespace = `App\Transactors`

/*Variables used between functions*/
var queryVariableName string
var mutatorVariableName string
var superConstructorCall *core.FunctionCall

type TransactorGenerator struct {
	classBuilder   interfaces.Class
	identifier     string
	imports        []string
	classToExtend string /*Base,File,Image*/
	transactorMembers []api.TabbedUnit
	constructorStatements []api.TabbedUnit
	parentConstructorCallArgs []api.TabbedUnit

}

func NewTransactorGenerator(identifier string, classToExtend string) *TransactorGenerator {
	return &TransactorGenerator{
		classBuilder:   builder.NewClassBuilder(),
		identifier:     identifier,
		imports:        []string{},
		classToExtend: classToExtend,
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

/*FUNCTIONS TO BE USED BY HANDLERS*/


/**
Add statements inside the transactor constructor
Parameters:
	- statement: api.TabbedUnit
Returns:
	- instance of Transactor Generator object
*/
func (transactorGenerator *TransactorGenerator) AddConstructorStatement(
	statement api.TabbedUnit) *TransactorGenerator{
	transactorGenerator.constructorStatements =append(transactorGenerator.constructorStatements, statement)
	return transactorGenerator
}

/**
Add the parameters to the parent::constructor function call
Every type of transactor requires different sets of parameters according to their extended class
Parameters:
	- parameter: *core.Parameter
Returns:
	- instance of Transactor Generator object
*/

func (transactorGenerator *TransactorGenerator) AddParentConstructorCallArgs(
	parameter *core.Parameter) *TransactorGenerator{
	transactorGenerator.parentConstructorCallArgs =append(transactorGenerator.parentConstructorCallArgs, parameter)
	return transactorGenerator
}

/**
Add Member To the Transactor class
Parameters:
	- member : api.TabbedUnit
Returns:
	- instance of Transactor Generator object
*/
func (transactorGenerator *TransactorGenerator) AddTransactorMember(
	member api.TabbedUnit) *TransactorGenerator{

	transactorGenerator.transactorMembers =append(transactorGenerator.transactorMembers, member)
	return transactorGenerator
}


/**
Appends import to the Transactor file
Parameters:
	- units: string array of the import
Returns:
	- instance of Transactor Generator object
Sample Usage:
	- AppendImport([]string{"App\User",})
*/
func (transactorGenerator *TransactorGenerator) AppendImports(units []string) *TransactorGenerator {
	transactorGenerator.imports = append(transactorGenerator.imports, units...)
	return transactorGenerator
}



/**
There are some common imports, class members and parameters to the parent constructor call
addDefaults Method inserts the common things in each of the array so that they will be already present during Build call
NOTE:This will prepend the values to each array as they are the ones which will come first and then the values added from
the handler
Prepend is required because defaults are added after the handler inserts the values
*/
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
		SetExtends(strcase.ToCamel(transactorGenerator.classToExtend)+"Transactor")

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
