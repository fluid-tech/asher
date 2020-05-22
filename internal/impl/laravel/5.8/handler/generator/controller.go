package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/builder/interfaces"
	"asher/internal/api/codebuilder/php/core"
	"github.com/iancoleman/strcase"
	"strings"
)

type ControllerGenerator struct {
	classBuilder interfaces.Class
	identifier   string
	imports      []string

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
Adds a Simple Statement
Parameters:
	- identifier: string
Returns:
	- Return instance of TabbedUnit
Sample Usage:
	- addSimpleStatement("Just A Simple Statement String")
*/
func (conGen *ControllerGenerator) addSimpleStatement(identifier string) *api.TabbedUnit {
	statement := api.TabbedUnit(core.NewSimpleStatement(identifier))
	return &statement
}

/**
Adds a FunctionCall String
Parameters:
	- identifier: string
Returns:
	- Return instance of TabbedUnit
Sample Usage:
	- addFunctionCall("create")
	-  o/p: create();
*/
func (conGen *ControllerGenerator) addFunctionCall(identifier string) *api.TabbedUnit {
	functionCallStatement := api.TabbedUnit(core.NewFunctionCall(identifier))
	return &functionCallStatement
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
func (conGen *ControllerGenerator) addReturnStatement(identifier string) *api.TabbedUnit {
	returnStatement := api.TabbedUnit(core.NewReturnStatement(identifier))
	return &returnStatement
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
func (conGen *ControllerGenerator) addParameter(identifier string) *api.TabbedUnit {
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
func (conGen *ControllerGenerator) addMemberInClass(visibility string, identifier string) *ControllerGenerator {
	variable := api.TabbedUnit(core.NewVarDeclaration(visibility, identifier))
	conGen.classBuilder.AddMember(&variable)
	return conGen
}

/**
Sets the identifier of the current class
Parameters:
	- identifier: string
Sample Usage:
	- SetIdentifier("ClassName")
*/
func (conGen *ControllerGenerator) SetIdentifier(identifier string) {
	conGen.identifier = identifier
}

/**
Add Create Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddCreateFunction()
*/
func (conGen *ControllerGenerator) AddCreateFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier + `Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->create`)
	functionCallStatement.AddArg(conGen.addParameter("Auth::id()"))
	functionCallStatement.AddArg(conGen.addParameter("$request->all()"))
	createCallStatement := api.TabbedUnit(functionCallStatement)

	returnStatement := conGen.addReturnStatement("$" + loweCamelCaseIdentifier)

	createFunctionStatement := []*api.TabbedUnit{
		&createCallStatement,
		returnStatement,
	}

	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("create").
		SetVisibility("public").AddArgument("Request $request").
		AddStatements(createFunctionStatement).GetFunction())
	return conGen
}

/**
Add Update Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddUpdateFunction()
*/
func (conGen *ControllerGenerator) AddUpdateFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier + `Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->update`)
	functionCallStatement.AddArg(conGen.addParameter("Auth::id()"))
	functionCallStatement.AddArg(conGen.addParameter("$request->all()"))
	updateCallStatement := api.TabbedUnit(functionCallStatement)

	returnStatement := conGen.addReturnStatement("$" + loweCamelCaseIdentifier)

	updateFunctionStatement := []*api.TabbedUnit{
		&updateCallStatement,
		returnStatement,
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("update").
		SetVisibility("public").AddArgument("Request $request").
		AddStatements(updateFunctionStatement).GetFunction())
	return conGen
}

/**
Add Delete Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddDeleteFunction()
*/
func (conGen *ControllerGenerator) AddDeleteFunction() *ControllerGenerator {
	loweCamelCaseIdentifier := strcase.ToLowerCamel(conGen.identifier)
	transactorVariableName := loweCamelCaseIdentifier + `Transactor`

	functionCallStatement := core.NewFunctionCall(
		`$` + loweCamelCaseIdentifier + ` = $this->` + transactorVariableName + `->delete`)
	functionCallStatement.AddArg(conGen.addParameter("$id"))
	functionCallStatement.AddArg(conGen.addParameter("$request->user->id"))
	deleteCallStatement := api.TabbedUnit(functionCallStatement)
	returnStatement := conGen.addReturnStatement("$" + loweCamelCaseIdentifier)

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

/**
Add FindById Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddFindByIdFunction()
*/
func (conGen *ControllerGenerator) AddFindByIdFunction() *ControllerGenerator {
	queryVariableName := strcase.ToLowerCamel(conGen.identifier) + `Query`
	returnTryStatement := []*api.TabbedUnit{
		conGen.addReturnStatement(`response(['data' => $this->` + queryVariableName + `->findById($id)])`),
	}
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().SetName("findById").
		AddArgument("$id").SetVisibility("public").AddStatements(returnTryStatement).GetFunction())
	return conGen
}

/**
Add GetAll Function of REST in the controller
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddGetAllFunction()
*/
func (conGen *ControllerGenerator) AddGetAllFunction() *ControllerGenerator {
	queryVariableName := strcase.ToLowerCamel(conGen.identifier) + `Query`
	returnStatement := core.NewReturnStatement(`$this->` + queryVariableName + `->datatables()`)
	tabbedUnitReturnStatement := api.TabbedUnit(returnStatement)
	conGen.classBuilder.AddFunction(builder.NewFunctionBuilder().
		SetName("getAll").SetVisibility("public").
		AddStatement(&tabbedUnitReturnStatement).GetFunction())
	return conGen
}

/**
Adds Constructor in the controller with Query and Transactor Injected of the currentController
Returns:
	- Return instance of ControllerGenerator
Sample Usage:
	- controllerGeneratorObject.AddConstructorFunction()
*/
func (conGen *ControllerGenerator) AddConstructorFunction() *ControllerGenerator {
	lowerCamelIdentifier := strcase.ToLowerCamel(conGen.identifier)
	queryVariableName := lowerCamelIdentifier + `Query`
	transactorVariableName := lowerCamelIdentifier + `Transactor`
	constructorArguments := []string{
		conGen.identifier + `Query $` + queryVariableName,
		conGen.identifier + `Transactor $` + transactorVariableName,
	}

	conGen.addMemberInClass("private", queryVariableName)
	conGen.addMemberInClass("private", transactorVariableName)

	constructorStatements := []*api.TabbedUnit{
		conGen.addSimpleStatement("$this->" + queryVariableName + " = $" + queryVariableName),
		conGen.addSimpleStatement("$this->" + transactorVariableName + " = $" + transactorVariableName),
	}
	conGen.classBuilder.AddFunction(
		builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
			AddArguments(constructorArguments).AddStatements(constructorStatements).GetFunction())
	return conGen
}

func (conGen *ControllerGenerator) AddAllRESTFunctionsInController() {
	conGen.AddConstructorFunction()
	conGen.AddCreateFunction()
	conGen.AddUpdateFunction()
	conGen.AddDeleteFunction()
	conGen.AddFindByIdFunction()
	conGen.AddGetAllFunction()
}

func (conGen *ControllerGenerator) AddFunctionsInController(methods []string) {
	if methods != nil {
		if len(methods) >= 0 {
			conGen.AddConstructorFunction()
			for _, element := range methods {
				switch strings.ToLower(element) {
				case "POST":
					conGen.AddCreateFunction()
				case "GET":
					conGen.AddFindByIdFunction().AddGetAllFunction()
				case "PUT":
					conGen.AddUpdateFunction()
				case "DELETE":
					conGen.AddDeleteFunction()
				}
			}
		}
	} else {
		conGen.AddAllRESTFunctionsInController()
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
	className := conGen.identifier + "RestController"
	namespace := `App\Http\Controllers`
	extends := "Controller"

	restControllerImports := []string{
		`App\` + conGen.identifier,
		`App\Transactors\` + conGen.identifier + `Transactor`,
		`App\Query\` + conGen.identifier + `Query`,
		`Illuminate\Http\Request`,
	}

	conGen.AppendImports(restControllerImports)

	//Adding functions in the controller

	conGen.classBuilder.SetName(className).
		SetPackage(namespace).
		SetExtends(extends).
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
