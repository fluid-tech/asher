package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"github.com/iancoleman/strcase"
	"go/ast"
	"strings"
)

var db_begin_transaction = api.TabbedUnit(core.NewSimpleStatement("DB::beginTransaction()"))
var db_commit = api.TabbedUnit(core.NewSimpleStatement("DB::commit()"))

type TransactorGenerator struct {
	classBuilder interfaces.Class
	identifier   string
	imports      []string
	constructorArguments []string
	constructorStatement []*api.TabbedUnit
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
Adds a Member in class
Parameters:
	- visibility: string
	- identifier: string
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- addMemberInClass("public", "variableName")
*/
func (transactorGenerator *TransactorGenerator) addMemberInClass(visibility string, identifier string) *TransactorGenerator {
	variable := api.TabbedUnit(core.NewVarDeclaration(visibility, identifier))
	transactorGenerator.classBuilder.AddMember(&variable)
	return transactorGenerator
}

/**
Adds a Return Statement
Parameters:
	- identifier: string
Returns:
	- Return instance of TabbedUnit
Sample Usage:
	- addSimpleStatement("Just A Simple Statement String")
*/
func (transactorGenerator *TransactorGenerator) addReturnStatement(identifier string) *api.TabbedUnit {
	returnStatement := api.TabbedUnit(core.NewReturnStatement(identifier))
	return &returnStatement
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
	transactorGenerator.constructorArguments = append(transactorGenerator.constructorArguments, constructorArguments...)

	parentConstructorCall := api.TabbedUnit(core.NewFunctionCall("parent::__construct").
		AddArg(transactorGenerator.addParameter("$" + queryVariableName + ", $" + mutatorVariableName + ", 'id'")))

	constructorStatements := []*api.TabbedUnit{
		&parentConstructorCall,
		transactorGenerator.addSimpleStatement("$this->className = self::CLASS_NAME"),
	}
	transactorGenerator.constructorStatement = append(transactorGenerator.constructorStatement, constructorStatements...)

	transactorGenerator.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddArguments(transactorGenerator.constructorArguments).AddStatements(transactorGenerator.constructorStatement).GetFunction())
	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) AddCreateFunction() *TransactorGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)

	createFunctionCall := core.NewFunctionCall(`$`+loweCamelCaseIdentifier+"v= parent::create")
	createFunctionCall.AddArg(transactorGenerator.addParameter("$createdById"))
	createFunctionCall.AddArg(transactorGenerator.addParameter("$args"))
	createCallStatement := api.TabbedUnit(createFunctionCall)

	returnStatement := transactorGenerator.addReturnStatement("$" + loweCamelCaseIdentifier)

	var createFunctionStatement []*api.TabbedUnit
	var hasOneStatements = transactorGenerator.CheckHasOneStatment()
	createFunctionStatement = append(createFunctionStatement, &db_begin_transaction)
	for _, element := range hasOneStatements {
		createFunctionStatement = append(createFunctionStatement, element)
	}
	createFunctionStatement = append(createFunctionStatement, &createCallStatement)
	var hasmanyStatement = transactorGenerator.CheckHasManyStatements()
	for _, element := range hasmanyStatement {
		createFunctionStatement = append(createFunctionStatement, element)
	}
	createFunctionStatement = append(createFunctionStatement, &db_commit)
	createFunctionStatement = append(createFunctionStatement, returnStatement)


	transactorGenerator.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("create").
		SetVisibility("public").AddArgument("Request $request").
		AddStatements(createFunctionStatement).GetFunction())


	return transactorGenerator
}

func (transactorGenerator *TransactorGenerator) CheckHasManyStatements() []*api.TabbedUnit  {
	hasManyStatements := []*api.TabbedUnit{}
	modelContext := context.GetFromRegistry("model").GetCtx(transactorGenerator.identifier)
	modelGenerator := modelContext.(*ModelGenerator)
	for _,element := range  modelGenerator.relationshipDetails {
		switch element.RelationshipType() {
		case helper.HasManny:
			statements := transactorGenerator.AddHasManyStatements("Comment:blog_id,id")
			hasManyStatements = append(hasManyStatements, statements...)
		}
	}
	return hasManyStatements
}

func (transactorGenerator *TransactorGenerator) AddHasManyStatements(hasMany string) []*api.TabbedUnit  {
	statements := []*api.TabbedUnit{}
	hasManyStrings := strings.Split(hasMany, `:`)
	loweCamelCaseIdentifier := strcase.ToLowerCamel(hasManyStrings[0])

	transactorVariableName := loweCamelCaseIdentifier + `Transactor`

	keys := strings.Split(hasManyStrings[1], `,`)

	functionCallStatement := core.NewFunctionCall(
		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->create`)
	functionCallStatement.AddArg(transactorGenerator.addParameter("$createdById"))
	functionCallStatement.AddArg(transactorGenerator.addParameter("$args"))
	apiFunctionCall := api.TabbedUnit(functionCallStatement)

	forEachStatement := api.TabbedUnit(core.NewForEach().AddCondition(`$args['`+keys[0]+`'] as $`+keys[0]).
		AddStatement(transactorGenerator.addSimpleStatement(`$args['`+keys[0]+`'] = `+keys[0])).
		AddStatement(&apiFunctionCall))
	statements = append(statements, &forEachStatement)
	return statements
}


func (transactorGenerator *TransactorGenerator) CheckHasOneStatment() []*api.TabbedUnit {
	hasOneStatements := []*api.TabbedUnit{}
	modelContext := context.GetFromRegistry("model").GetCtx(transactorGenerator.identifier)
	modelGenerator := modelContext.(*ModelGenerator)
	for _,element := range  modelGenerator.relationshipDetails {
		switch element.RelationshipType() {
		case helper.HasOne:
			statements := transactorGenerator.AddHasOneStatements("Admin:user_id,id")
			 hasOneStatements = append(hasOneStatements, statements...)

		}
	}
	return hasOneStatements
}
func (transactorGenerator *TransactorGenerator) AddHasOneStatements(hasOne string) []*api.TabbedUnit  {
	statements := []*api.TabbedUnit{}
	hasOneStrings := strings.Split(hasOne, `:`)
	loweCamelCaseIdentifier := strcase.ToLowerCamel(hasOneStrings[0])
	transactorVariableName := loweCamelCaseIdentifier + `Transactor`
	keys := strings.Split(hasOneStrings[1], `,`)
	functionCallStatement := core.NewFunctionCall(
		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->create`)
	functionCallStatement.AddArg(transactorGenerator.addParameter("$createdById"))
	functionCallStatement.AddArg(transactorGenerator.addParameter("$args"))
	apiFunctionCall := api.TabbedUnit(functionCallStatement)

	hasOneStatement := transactorGenerator.addSimpleStatement(`$args[`+keys[0]+`] = `+loweCamelCaseIdentifier+`->`+keys[1])


	statements = append(statements, &apiFunctionCall)
	statements = append(statements, hasOneStatement)
	return statements
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

	transactorGenerator.CheckForRelationships()

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
