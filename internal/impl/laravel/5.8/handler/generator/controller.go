package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
)

type ControllerGenerator struct {
	classBuilder		interfaces.Class
	identifier 			string
	imports 			[]string
}

func NewControllerGenerator() *ControllerGenerator  {
	return &ControllerGenerator{
		classBuilder: builder.NewClassBuilder(),
		identifier: "",
		imports: []string{},

	}
}

func (conGen *ControllerGenerator) AddImports(units []string)  *ControllerGenerator {
	conGen.imports = append(conGen.imports, units...)
	return conGen
}

func (conGen *ControllerGenerator) AddSimpleStatement(
	identifier string) *api.TabbedUnit  {
	statement:= api.TabbedUnit(core.NewSimpleStatement(identifier))
	return &statement
}

func (conGen *ControllerGenerator) AddFunctionCall(
	identifier string) *api.TabbedUnit  {
	functionCallStatement := api.TabbedUnit(core.NewFunctionCall(identifier))
	return &functionCallStatement
}

func (conGen *ControllerGenerator) AddReturnStatement(
	identifier string) *api.TabbedUnit {
	returnStatement := api.TabbedUnit(core.NewReturnStatement(identifier))
	return &returnStatement
}

func (conGen *ControllerGenerator) AddParameter(
	identifier string) *api.TabbedUnit  {
	parameter := api.TabbedUnit(core.NewParameter(identifier))
	return &parameter
}

func (conGen *ControllerGenerator) SetIdentifier(identifier string)  {
	conGen.identifier  = identifier
}

// Creating create function builder for Rest controller
func (conGen *ControllerGenerator) AddCreateFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier+`Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$`+loweCamelCaseIdentifier+` = $this->`+transactorVariableName+`->create`)
	functionCallStatement.AddArg(conGen.AddParameter("Auth::id()"))
	functionCallStatement.AddArg(conGen.AddParameter("$request->all()"))
	createCallStatement := api.TabbedUnit(functionCallStatement)

	returnStatement := conGen.AddReturnStatement("$"+loweCamelCaseIdentifier)

	createFunctionStatement := []*api.TabbedUnit{
		&createCallStatement,
		returnStatement,
	}

	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("create").
		SetVisibility("public").AddArgument("Request $request").
		AddStatements(createFunctionStatement).GetFunction())
	return conGen
}

// Creating update function builder for controller
func (conGen *ControllerGenerator) AddUpdateFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier+`Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$`+loweCamelCaseIdentifier+` = $this->`+transactorVariableName+`->update`)
	functionCallStatement.AddArg(conGen.AddParameter("Auth::id()"))
	functionCallStatement.AddArg(conGen.AddParameter("$request->all()"))
	updateCallStatement := api.TabbedUnit(functionCallStatement)

	returnStatement := conGen.AddReturnStatement("$"+loweCamelCaseIdentifier)

	updateFunctionStatement := []*api.TabbedUnit{
		&updateCallStatement,
		returnStatement,
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("update").
		SetVisibility("public").AddArgument("Request $request").
		AddStatements(updateFunctionStatement).GetFunction())
	return conGen
}

// Creating delete function builder for controller
func (conGen *ControllerGenerator) AddDeleteFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier+`Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$`+loweCamelCaseIdentifier+` = $this->`+transactorVariableName+`->delete`)
	functionCallStatement.AddArg(conGen.AddParameter("$id"))
	functionCallStatement.AddArg(conGen.AddParameter("$request->user->id"))
	deleteCallStatement := api.TabbedUnit(functionCallStatement)
	returnStatement := conGen.AddReturnStatement("$"+loweCamelCaseIdentifier)

	deleteFunctionArgument := []string{
		"Request $request",
		"$id",
	}
	deleteFunctionStatement := []*api.TabbedUnit{
		&deleteCallStatement,
		returnStatement,
	}

	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("delete").
		SetVisibility("public").AddArguments(deleteFunctionArgument).
		AddStatements(deleteFunctionStatement).GetFunction())
	return conGen
}

// Creating findById function builder for controller
func (conGen *ControllerGenerator) AddFindByIdFunction() *ControllerGenerator {

	//creating Try Block
	returnTryStatement := []*api.TabbedUnit{
		conGen.AddReturnStatement(
			`response(['data' => `+conGen.identifier+`::findOrFail($id)])`),
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("findById").
		AddArgument("$id").SetVisibility("public").AddStatements(returnTryStatement).GetFunction())
	return conGen
}


// Creating getAll function builder for controller
func (conGen *ControllerGenerator) AddGetAllFunction() *ControllerGenerator  {
	returnStatement := core.NewReturnStatement(conGen.identifier+ "::all()")
	tabbedUnitReturnStatement := api.TabbedUnit(returnStatement)
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().
		SetName("getAll").SetVisibility("public").
		AddStatement(&tabbedUnitReturnStatement).GetFunction())
	return conGen
}

func (conGen *ControllerGenerator) AddMemberInClass(
	visibility string,identifier string)  *ControllerGenerator {
	variable := api.TabbedUnit(core.NewVarDeclaration(visibility, identifier))
	conGen.classBuilder.AddMember(&variable)
	return conGen
}


// Creates a constructor of controller with Query and Mutator Class injected of the current model
func (conGen *ControllerGenerator) AddConstructor() *ControllerGenerator  {
	lowerCamelIdentifier := strcase.ToLowerCamel(conGen.identifier)
	queryVariableName := lowerCamelIdentifier+`Query`
	transactorVariableName := lowerCamelIdentifier+`Transactor`
	constructorArguments := []string{
		conGen.identifier+`Query $`+queryVariableName,
		conGen.identifier+`Transactor $`+ transactorVariableName,
	}

	conGen.AddMemberInClass("private", queryVariableName)
	conGen.AddMemberInClass("private", transactorVariableName)
	constructorStatements := []*api.TabbedUnit{
		conGen.AddSimpleStatement(
			"$this->"+queryVariableName+" = "+queryVariableName),
		conGen.AddSimpleStatement(
			"$this->"+ transactorVariableName +" = "+ transactorVariableName),
	}
	conGen.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return conGen
}


//Builds the RestController Class
func (conGen *ControllerGenerator) BuildRestController() *core.Class {
	className := conGen.identifier + "RestController"
	namespace := `App\Http\Controllers`
	extends  := "Controller"

	restControllerImports := []string{
		`App\`+conGen.identifier,
		`App\Transactors\`+conGen.identifier+`Transactor`,
		`App\Query\`+conGen.identifier+`Query`,
		`Illuminate\Http\Request`,
	}

	conGen.AddImports(restControllerImports)

	//Adding functions in the controller
	conGen.AddConstructor()
	conGen.AddCreateFunction()
	conGen.AddUpdateFunction()
	conGen.AddDeleteFunction()
	conGen.AddFindByIdFunction()
	conGen.AddGetAllFunction()

	conGen.classBuilder.SetName(className).
		SetPackage(namespace).
		SetExtends(extends).
		AddImports(conGen.imports)
	return conGen.classBuilder.GetClass()
}

func (conGen *ControllerGenerator) String() string  {
	return conGen.BuildRestController().String()
}


