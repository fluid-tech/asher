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
const RequestUserId = "$request->user->id"
const Transactor = "Transactor"
const Query = "Query"
const ThisAssignFmt = "$this->%s = $%s"
const ResponseHelperCreateFmt = `ResponseHelper::create($%s)`
const ResponseHelperDeleteFmt = `ResponseHelper::delete($%s)`
const ResponseHelperUpdateFmt = `ResponseHelper::update($%s)`
const ResponseHelperSuccess = `ResponseHelper::success`
const CreateCallFmt = `$%s = $this->%s->create`
const DeleteCallFmt = `$%s = $this->%s->delete`
const UpdateCallFmt = `$%s = $this->%s->update`
const FindByIdCallFmt = `%s($this->%s->findById(%s))`
const GetAllCall = `%s($this->%s->paginate())`
const ResponseHelperImport = `App\Helpers\ResponseHelper`
const RequestImport = `Illuminate\Http\Request`

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
	conGen.queryVariableName = conGen.lowerCamelCaseIdentifier + Query
	conGen.transactorVariableName = conGen.lowerCamelCaseIdentifier + Transactor
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
		fmt.Sprintf(CreateCallFmt, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName)).
		AddArg(core.NewParameter(AuthID)).AddArg(core.NewParameter(RequestAll))

	responseHelperCreate := core.NewReturnStatement(fmt.Sprintf(ResponseHelperCreateFmt, conGen.lowerCamelCaseIdentifier))

	createFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		responseHelperCreate,
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
		fmt.Sprintf(UpdateCallFmt, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName))

	functionCallStatement.AddArg(core.NewParameter(AuthID))
	functionCallStatement.AddArg(core.NewParameter(RequestAll))

	responseHelperUpdate := core.NewReturnStatement(fmt.Sprintf(ResponseHelperUpdateFmt, conGen.lowerCamelCaseIdentifier))

	updateFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		responseHelperUpdate,
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
		fmt.Sprintf(DeleteCallFmt, conGen.lowerCamelCaseIdentifier, conGen.transactorVariableName))
	functionCallStatement.AddArg(core.NewParameter(Id)).
		AddArg(core.NewParameter(RequestUserId))

	responseHelperDelete := core.NewReturnStatement(fmt.Sprintf(ResponseHelperDeleteFmt, conGen.lowerCamelCaseIdentifier))

	deleteFunctionStatement := []api.TabbedUnit{
		functionCallStatement,
		responseHelperDelete,
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
	responseHelperSuccess := []api.TabbedUnit{
		core.NewReturnStatement(fmt.Sprintf(
			FindByIdCallFmt, ResponseHelperSuccess, conGen.queryVariableName, Id)),
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName(FindByIdMethod).
		AddArgument(Id).SetVisibility(VisibilityPublic).AddStatements(responseHelperSuccess).GetFunction())
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
	responseHelperSuccess := core.NewReturnStatement(
		fmt.Sprintf(GetAllCall, ResponseHelperSuccess, conGen.queryVariableName))
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().
		SetName(GetAllMethod).SetVisibility(VisibilityPublic).
		AddStatement(responseHelperSuccess).GetFunction())
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
		fmt.Sprintf(QueryObjectFmt, conGen.identifier, conGen.queryVariableName),
		fmt.Sprintf(TransactorObjectFmt, conGen.identifier, conGen.transactorVariableName),
	}

	conGen.classBuilder.AddMember(core.NewVarDeclaration(VisibilityPrivate, conGen.queryVariableName))
	conGen.classBuilder.AddMember(core.NewVarDeclaration(VisibilityPrivate, conGen.transactorVariableName))

	constructorStatements := []api.TabbedUnit{
		core.NewSimpleStatement(fmt.Sprintf(ThisAssignFmt, conGen.queryVariableName, conGen.queryVariableName)),
		core.NewSimpleStatement(fmt.Sprintf(ThisAssignFmt, conGen.transactorVariableName, conGen.transactorVariableName)),
	}

	conGen.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility(VisibilityPublic).SetName(CallConstructor).
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
	- controllerGeneratorObject.AddFunctionsInController([]string{"Post"})
*/
func (conGen *ControllerGenerator) AddFunctionsInController(methods []string) {
	if methods != nil && len(methods) > 0 {
		conGen.AddConstructor()
		for _, element := range methods {
			switch strings.ToUpper(element) {
			case HttpPost:
				conGen.AddCreate()
			case HttpGet:
				conGen.AddFindById().AddGetAll()
			case HttpPut:
				conGen.AddUpdate()
			case HttpDelete:
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
		fmt.Sprintf(`App\%s`, conGen.identifier),
		fmt.Sprintf(`App\Transactors\%s%s`, conGen.identifier, Transactor),
		fmt.Sprintf(`App\Query\%s%s`, conGen.identifier, Query),
		RequestImport,
		ResponseHelperImport,
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
