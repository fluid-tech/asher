package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

var dbBeginTransaction = core.NewSimpleStatement("DB::beginTransaction()")
var dbCommit = core.NewSimpleStatement("DB::commit()")

type TransactorGenerator struct {
	classBuilder interfaces.Class
	identifier   string
	imports      []string
	transactorType string
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

	parentConstructorCall := core.NewFunctionCall("parent::__construct").
		AddArg(core.NewParameter("$" + queryVariableName + ", $" + mutatorVariableName + ", 'id'"))

	constructorStatements := []api.TabbedUnit{
		parentConstructorCall,
		core.NewSimpleStatement("$this->className = self::CLASS_NAME"),
	}

	transactorGenerator.classBuilder.AddFunction(builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return transactorGenerator
}

//func (transactorGenerator *TransactorGenerator) AddCreateFunction() *TransactorGenerator {
//	loweCamelCaseIdentifier := strcase.ToLowerCamel(transactorGenerator.identifier)
//
//	createFunctionCall := core.NewSimpleStatement(`$`+loweCamelCaseIdentifier+" = parent::create($createById, $args)")
//	createCallStatement := api.TabbedUnit(createFunctionCall)
//
//	returnStatement := transactorGenerator.addReturnStatement("$" + loweCamelCaseIdentifier)
//
//	createFunctionStatement := []*api.TabbedUnit {
//		&dbBeginTransaction,
//		&createCallStatement,
//		&dbCommit,
//		returnStatement,
//	}
//	//var hasOneStatements = transactorGenerator.CheckHasOneStatment()
//
//	//for _, element := range hasOneStatements {
//	//	createFunctionStatement = append(createFunctionStatement, element)
//	//}
//
//	//var hasmanyStatement = transactorGenerator.CheckHasManyStatements()
//	//for _, element := range hasmanyStatement {
//	//	createFunctionStatement = append(createFunctionStatement, element)
//	//}
//	transactorGenerator.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("create").
//		SetVisibility("public").AddArgument("$createById, $args, $extraArgs = null").
//		AddStatements(createFunctionStatement).GetFunction())
//	return transactorGenerator
//}

//func (transactorGenerator *TransactorGenerator) CheckHasManyStatements() []*api.TabbedUnit  {
//	hasManyStatements := []*api.TabbedUnit{}
//	modelContext := context.GetFromRegistry("model").GetCtx(transactorGenerator.identifier)
//	modelGenerator := modelContext.(*ModelGenerator)
//	for _,element := range  modelGenerator.relationshipDetails {
//		switch element.RelationshipType() {
//		case helper.HasManny:
//			statements := transactorGenerator.AddHasManyStatements("Comment:blog_id,id")
//			hasManyStatements = append(hasManyStatements, statements...)
//		}
//	}
//	return hasManyStatements
//}
//
//func (transactorGenerator *TransactorGenerator) AddHasManyStatements(hasMany string) []*api.TabbedUnit  {
//	statements := []*api.TabbedUnit{}
//	hasManyStrings := strings.Split(hasMany, `:`)
//	loweCamelCaseIdentifier := strcase.ToLowerCamel(hasManyStrings[0])
//
//	transactorVariableName := loweCamelCaseIdentifier + `Transactor`
//
//	keys := strings.Split(hasManyStrings[1], `,`)
//
//	functionCallStatement := core.NewFunctionCall(
//		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->create`)
//	functionCallStatement.AddArg(transactorGenerator.addParameter("$createdById"))
//	functionCallStatement.AddArg(transactorGenerator.addParameter("$args"))
//	apiFunctionCall := api.TabbedUnit(functionCallStatement)
//
//	forEachStatement := api.TabbedUnit(core.NewForEach().AddCondition(`$args['`+keys[0]+`'] as $`+keys[0]).
//		AddStatement(transactorGenerator.addSimpleStatement(`$args['`+ keys[0]+`'] = `+keys[0])).
//		AddStatement(&apiFunctionCall))
//	statements = append(statements, &forEachStatement)
//	return statements
//}
//
//
//func (transactorGenerator *TransactorGenerator) CheckHasOneStatment() []*api.TabbedUnit {
//	hasOneStatements := []*api.TabbedUnit{}
//	modelContext := context.GetFromRegistry("model").GetCtx(transactorGenerator.identifier)
//	modelGenerator := modelContext.(*ModelGenerator)
//	for _,element := range  modelGenerator.relationshipDetails {
//		switch element.RelationshipType() {
//		case helper.HasOne:
//			statements := transactorGenerator.AddHasOneStatements("Admin:user_id,id")
//			 hasOneStatements = append(hasOneStatements, statements...)
//
//		}
//	}
//	return hasOneStatements
//}
//func (transactorGenerator *TransactorGenerator) AddHasOneStatements(hasOne string) []*api.TabbedUnit  {
//	statements := []*api.TabbedUnit{}
//	hasOneStrings := strings.Split(hasOne, `:`)
//	loweCamelCaseIdentifier := strcase.ToLowerCamel(hasOneStrings[0])
//	transactorVariableName := loweCamelCaseIdentifier + `Transactor`
//	keys := strings.Split(hasOneStrings[1], `,`)
//	functionCallStatement := core.NewFunctionCall(
//		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->create`)
//	functionCallStatement.AddArg(transactorGenerator.addParameter("$createdById"))
//	functionCallStatement.AddArg(transactorGenerator.addParameter("$args"))
//	apiFunctionCall := api.TabbedUnit(functionCallStatement)
//
//	hasOneStatement := transactorGenerator.addSimpleStatement(`$args[`+keys[0]+`] = `+loweCamelCaseIdentifier+`->`+keys[1])
//
//
//	statements = append(statements, &apiFunctionCall)
//	statements = append(statements, hasOneStatement)
//	return statements
//}



/**
Main Function To be called when we want to build the transactor
Returns:
	- Return instance of core.Class
Sample Usage:
	- transactorGeneratorObject.BuildRestTransactor()
*/
func (transactorGenerator *TransactorGenerator) BuildTransactor() *core.Class {
	const namespace = `App\Transactors`
	var extendsTransactor string

	switch transactorGenerator.transactorType {
		case "default":
			extendsTransactor = "BaseTransactor"
		case "file" :
			extendsTransactor = "FileTransactor"
		case "image":
			extendsTransactor = "ImageTransactor"
	}


	transactorImports := []string{
		`App\Query\` + transactorGenerator.identifier + `Query`,
		`App\Transactors\Mutations\` + transactorGenerator.identifier + `Mutator`,
	}

	className := transactorGenerator.identifier + "Transactor"

	transactorGenerator.AppendImports(transactorImports)
	transactorGenerator.AddConstructorFunction()


	transactorGenerator.classBuilder.AddMember(core.NewSimpleStatement(
		"private const CLASS_NAME = '" + className + "'")).SetName(className).
		SetExtends(extendsTransactor).SetPackage(namespace).AddImports(transactorGenerator.imports)

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
