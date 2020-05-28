package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

const ControllerNamespace = `App\Http\Controllers\Api`
const ControllerExtends = "Controller"
const Request = "Request $request"
const AuthID = "Auth::id()"
const RequestAll = "$request->all()"
const Id = "$id"

type ControllerGenerator struct {
	api.Generator
	classBuilder             interfaces.Class
	identifier               string
	imports                  []string
	queryVariableName        string
	transactorVariableName   string
	lowerCamelCaseIdentifier string
}

func NewControllerGenerator() *ControllerGenerator {
	return &ControllerGenerator{
		classBuilder: builder.NewClassBuilder(),
		identifier:   "",
		imports:      []string{},
	}
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
func (conGen *ControllerGenerator) AppendImports(units []string) *ControllerGenerator {
	conGen.imports = append(conGen.imports, units...)
	return conGen
}

/**
Sets the identifier of the current class
Parameters:
	- identifier: string
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (conGen *ControllerGenerator) SetIdentifier(identifier string) *ControllerGenerator {
	conGen.identifier = identifier
	conGen.lowerCamelCaseIdentifier = strcase.ToLowerCamel(conGen.identifier)
	conGen.queryVariableName = conGen.lowerCamelCaseIdentifier + `Query`
	conGen.transactorVariableName = conGen.lowerCamelCaseIdentifier + `Transactor`
	return conGen
}

/**
Add Create Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddCreate()
*/
func (conGen *ControllerGenerator) AddCreate() *ControllerGenerator {

	functionCallStatement := core.NewFunctionCall(
		fmt.Sprintf(`$%s = $this->%s->create`, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName)).
		AddArg(core.NewParameter(AuthID)).AddArg(core.NewParameter(RequestAll))

	returnStatement := core.NewReturnStatement(fmt.Sprintf(`ResponseHelper::create($%s)`, conGen.lowerCamelCaseIdentifier))

	createFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		returnStatement,
	}

	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName(CreateMethod).
		SetVisibility(VisibilityPublic).AddArgument(Request).
		AddStatements(createFunctionStatement).GetFunction())
	return conGen
}

/**
Add Update Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddUpdate()
*/
func (conGen *ControllerGenerator) AddUpdate() *ControllerGenerator {

	functionCallStatement := core.NewFunctionCall(
		fmt.Sprintf(`$%s = $this->%s->update`, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName))

	functionCallStatement.AddArg(core.NewParameter(AuthID))
	functionCallStatement.AddArg(core.NewParameter(RequestAll))

	returnStatement := core.NewReturnStatement(fmt.Sprintf(`ResponseHelper::update($%s)`, conGen.lowerCamelCaseIdentifier))

	updateFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		returnStatement,
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName(UpdateMethod).
		SetVisibility(VisibilityPublic).AddArgument(Request).
		AddStatements(updateFunctionStatement).GetFunction())
	return conGen
}

/**
Add Delete Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddDelete()
*/
func (conGen *ControllerGenerator) AddDelete() *ControllerGenerator {

	functionCallStatement := core.NewFunctionCall(
		fmt.Sprintf(`$%s = $this->%s->delete`, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName))
	functionCallStatement.AddArg(core.NewParameter(Id)).
		AddArg(core.NewParameter("$request->user->id"))

	returnStatement := core.NewReturnStatement(fmt.Sprintf(`ResponseHelper::delete($%s)`, conGen.lowerCamelCaseIdentifier))

	deleteFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		returnStatement,
	}

	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName(DeleteMethod).
		SetVisibility(VisibilityPublic).AddArgument(Request).AddArgument(Id).
		AddStatements(deleteFunctionStatement).GetFunction())
	return conGen
}

/**
Add FindById Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddFindById()
*/
func (conGen *ControllerGenerator) AddFindById() *ControllerGenerator {
	returnTryStatement := []api.TabbedUnit{
		core.NewReturnStatement(fmt.Sprintf(
			`ResponseHelper::success($this->%s->findById(%s))`, conGen.queryVariableName, Id)),
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName(FindByIdMethod).
		AddArgument(Id).SetVisibility(VisibilityPublic).AddStatements(returnTryStatement).GetFunction())
	return conGen
}

/**
Add GetAll Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddGetAll()
*/
func (conGen *ControllerGenerator) AddGetAll() *ControllerGenerator {
	returnStatement := core.NewReturnStatement(
		fmt.Sprintf(`ResponseHelper::success($this->%s->paginate())`, conGen.queryVariableName))
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().
		SetName(GetAllMethod).SetVisibility(VisibilityPublic).
		AddStatement(returnStatement).GetFunction())
	return conGen
}

/**
Adds Constructor in the controller with Query and Transactor Injected of the currentController
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddConstructor()
*/
func (conGen *ControllerGenerator) AddConstructor() *ControllerGenerator {
	constructorArguments := []string{
		conGen.identifier + `Query $` + conGen.queryVariableName,
		conGen.identifier + `Transactor $` + conGen.transactorVariableName,
	}

	conGen.classBuilder.AddMember(core.NewVarDeclaration("private", conGen.queryVariableName))
	conGen.classBuilder.AddMember(core.NewVarDeclaration("private", conGen.transactorVariableName))

	constructorStatements := []api.TabbedUnit{
		core.NewSimpleStatement("$this->" + conGen.queryVariableName + " = $" + conGen.queryVariableName),
		core.NewSimpleStatement("$this->" + conGen.transactorVariableName + " = $" + conGen.transactorVariableName),
	}

	conGen.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility(VisibilityPublic).SetName(Constructor).
			AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return conGen
}

/**
Simply adds all the methods required inside the controller

Sample Usage:
	- controllerGeneratorObject.AddAllRESTMethods()
*/
func (conGen *ControllerGenerator) AddAllRESTMethods() {
	conGen.AddConstructor()
	conGen.AddCreate()
	conGen.AddUpdate()
	conGen.AddDelete()
	conGen.AddFindById()
	conGen.AddGetAll()
}

/**
Checks which all methods to be added in the controller
Parameters:
	- methods: string array of the methods allowed in the rest controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddFunctionsInController([]string{"POST"})
*/
func (conGen *ControllerGenerator) AddFunctionsInController(methods []string) {
	if methods != nil && len(methods) > 0 {
		conGen.AddConstructor()
		for _, element := range methods {
			switch strings.ToUpper(element) {
			case POST:
				conGen.AddCreate()
			case GET:
				conGen.AddFindById().AddGetAll()
			case PUT:
				conGen.AddUpdate()
			case DELETE:
				conGen.AddDelete()
			}
		}
	} else {
		conGen.AddAllRESTMethods()
	}
}

/**
Main Function To be called when we want to build the controller
Parameter:
	- controller configuration for controller
Returns:
	- Return instance of core.Class
Sample Usage:
	- controllerGeneratorObject.BuildRestController()
*/
func (conGen *ControllerGenerator) BuildRestController() *core.Class {
	className := fmt.Sprintf("%sRestController", conGen.identifier)

	restControllerImports := []string{
		`App\` + conGen.identifier,
		fmt.Sprintf(`App\Transactors\%sTransactor`, conGen.identifier),
		fmt.Sprintf(`App\Query\%sQuery`, conGen.identifier),
		`Illuminate\Http\Request`,
		`App\Helpers\ResponseHelper`,
	}

	conGen.AppendImports(restControllerImports)

	//Adding functions in the controller

	conGen.classBuilder.SetName(className).
		SetPackage(ControllerNamespace).
		SetExtends(ControllerExtends).
		AddImports(conGen.imports)
	return conGen.classBuilder.GetClass()
}

/**
Returns:
	- Return string object of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.String()
*/
func (conGen *ControllerGenerator) String() string {
	return conGen.BuildRestController().String()
}
