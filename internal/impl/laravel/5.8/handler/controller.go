package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"strings"
)

type ControllerHandler struct {
	api.Handler
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{}
}

/**
Depending upon the type of the controller its equivalent handler will be called
Parameters:
	- identifier: name of the model for which routes are to be generated
	-value : configuration for controller
Returns:
	- array of emitter files, error
*/
func (controllerHandler *ControllerHandler) Handle(
	identifier string, value interface{}) ([]api.EmitterFile, error) {

	controllerConfig := value.(models.Controller)
	var fileToEmit []api.EmitterFile

	if controllerConfig.Rest {
		fileToEmit = append(fileToEmit, controllerHandler.handleRestController(identifier, value)...)
	}

	return fileToEmit, nil
}

/**
IF the type of controller is REST in controller conifg then this method will be called
By default it will be using transactor patter
Internally It will
1.handleTransactor, which will create transactor for the model
2.handleMutator, which will create mutator for the model
3.handleQuery, which will create query for the model
4.handleRoutes, crete routes for the supported methods which will be added in routes/api.php
to read more about transactor pattern read docs about transactor patter
Parameters:
	- modelName: name of the model for which routes are to be generated
	-controller : configuration for controller
Returns:
	- array of emmitter files
*/
func (controllerHandler *ControllerHandler) handleRestController(
	identifier string, value interface{}) []api.EmitterFile {

	var fileToEmit []api.EmitterFile
	controllerConfig := value.(models.Controller)

	conGen := generator.NewControllerGenerator()
	conGen.SetIdentifier(identifier)

	conGen.AddFunctionsInController(controllerConfig.HttpMethods)

	context.GetFromRegistry("controller").AddToCtx(identifier, conGen)

	controllerEmitterFile := core.NewPhpEmitterFile(identifier+"RestController.php", api.ControllerPath, conGen, api.Controller)

	transactorEmitterFile := controllerHandler.handleTransactor(identifier, controllerConfig)
	mutatorEmitterFile := controllerHandler.handleMutator(identifier)
	queryEmitterFile := controllerHandler.handleQuery(identifier)
	routeEmitterFile := controllerHandler.handleRoutes(identifier, controllerConfig)

	fileToEmit = append(fileToEmit, controllerEmitterFile,transactorEmitterFile,mutatorEmitterFile,queryEmitterFile)
	if routeEmitterFile != nil {
		fileToEmit = append(fileToEmit, routeEmitterFile)
	}
	return fileToEmit
}

/*make methods private direct cast and pass in function call cast everything to * which ypu are fecthing from context*/
func (controllerHandler *ControllerHandler) handleTransactor(identifier string,
	controllerConfig models.Controller) api.EmitterFile {

	var transactorEmmiterFile api.EmitterFile

	modelGen := context.GetFromRegistry("model").GetCtx(identifier).(*generator.ModelGenerator)
	migrationGen := context.GetFromRegistry("migration").GetCtx(identifier).(*generator.MigrationGenerator)
	/*SWITCH CASE FOR TYPE OF TRANSACTOR*/

	switch strings.ToLower(controllerConfig.Type) {
	case "file":
		transactorEmmiterFile = controllerHandler.handleFileTransactor(identifier, controllerConfig, modelGen, migrationGen)
	case "image":
		transactorEmmiterFile = controllerHandler.handleImageTransactor(identifier, controllerConfig, modelGen, migrationGen)
	default:
		transactorEmmiterFile = controllerHandler.handleDefaultTransactor(identifier, controllerConfig)
	}

	return transactorEmmiterFile
}

func (controllerHandler *ControllerHandler) handleDefaultTransactor(identifier string,
	controllerConfig models.Controller) api.EmitterFile {
	//controller := value.(models.Controller)

	transactorGen := generator.NewTransactorGenerator(identifier, "Base")

	context.GetFromRegistry("transactor").AddToCtx(identifier, transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor.php", api.TransactorPath, transactorGen, api.Transactor)

	return transactorEmitter
}

func (controllerHandler *ControllerHandler) handleFileTransactor(identifier string, controllerConfig models.Controller,
	modelGen *generator.ModelGenerator, migrationGen *generator.MigrationGenerator) api.EmitterFile {

	/*MODEL AND MIGRATION UPDATES*/
	modelGen.AddCreateValidationRule("file_urls", "sometimes|required", "").
		AddCreateValidationRule("file_urls.urls", "array", "")

	modelGen.AddUpdateValidationRule("file_urls", "sometimes|required", "").
		AddUpdateValidationRule("file_urls.urls", "array", "")

	modelGen.AddFillable("file_urls")

	/*TODO DATA type should be configurable in 2nd iteration*/
	migrationGen.AddColumn(core.NewSimpleStatement(`$table->longText('file_urls')->nullable()`))

	transactorGen := generator.NewTransactorGenerator(identifier,"file")

	transactorGen.AppendImports([]string{`App\Helpers\BaseFileUploadHelper`}).
		AddParentConstructorCallArgs(core.NewParameter(
			`new BaseFileUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES,"png")`)).
		AddTransactorMember(core.NewSimpleStatement(`private const BASE_PATH = "`+
			strings.ToLower(identifier)+`"`)).
		AddTransactorMember(core.NewSimpleStatement(
			"public const IMAGE_VALIDATION_RULES =" +
				" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))


	context.GetFromRegistry("transactor").AddToCtx(identifier, transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor.php", api.TransactorPath, transactorGen, api.Transactor)

	return transactorEmitter
}

/*Depeneding upon file type
image */

func (controllerHandler *ControllerHandler) handleImageTransactor(identifier string, controllerConfig models.Controller,
	modelGen *generator.ModelGenerator, migrationGen *generator.MigrationGenerator) api.EmitterFile {
	//controller := value.(models.Controller)

	/*controller argumtn something while merging with master*/
	modelGen.AddCreateValidationRule("file_urls", "sometimes|required", "").
		AddCreateValidationRule("file_urls.urls", "array", "")

	modelGen.AddUpdateValidationRule("file_urls", "sometimes|required", "").
		AddUpdateValidationRule("file_urls.urls", "array", "")

	modelGen.AddFillable("file_urls")

	migrationGen.AddColumn(core.NewSimpleStatement(`$table->longText('file_urls')->nullable()`))

	/**/
	transactorGen := generator.NewTransactorGenerator(identifier, "image")
	transactorGen.AppendImports([]string{`App\Helpers\ImageUploadHelper`}).
		AddParentConstructorCallArgs(core.NewParameter(
			`new ImageUploadHelper(self::BASE_PATH, self::IMAGE_VALIDATION_RULES)`)).
		AddTransactorMember(core.NewSimpleStatement(`private const BASE_PATH = "`+
			strings.ToLower(identifier)+`"`)).
		AddTransactorMember(core.NewSimpleStatement(
			"public const IMAGE_VALIDATION_RULES =" +
				" array(\n        'file' => 'required|mimes:jpeg,jpg,png|max:3000'\n    )"))

	context.GetFromRegistry("transactor").AddToCtx(identifier, transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor.php", api.TransactorPath, transactorGen, api.Transactor)

	return transactorEmitter
}

func (controllerHandler *ControllerHandler) handleMutator(identifier string) api.EmitterFile {

	mutatorGen := generator.NewMutatorGenerator()
	mutatorGen.SetIdentifier(identifier)

	context.GetFromRegistry("mutator").AddToCtx(identifier, mutatorGen)
	mutatorEmitter :=
		core.NewPhpEmitterFile(identifier+"Mutator.php", api.MutatorPath, mutatorGen, api.Mutator)

	return mutatorEmitter
}

func (controllerHandler *ControllerHandler) handleQuery(
	identifier string) api.EmitterFile {


	queryGenerator := generator.NewQueryGenerator(identifier, true)

	context.GetFromRegistry("query").AddToCtx(identifier, queryGenerator)

	emitFile := core.NewPhpEmitterFile(identifier+"Query.php", api.QueryPath, queryGenerator, api.Query)


	return emitFile
}

/**
Checks if the asher_api.php is created or not
if it is not created this method will create it and will add it to the context to be used by other controllers
to add their routes
Parameters:
	- identifier: name of the model for which controller is generator
Returns:
	- if asher_api.php file already exists it returns blank *api.EmitterFile array
		else it will create and return the file in the array
*/
func (controllerHandler *ControllerHandler) handleRoutes(identifier string, controllerConfig models.Controller) api.EmitterFile {


	addRouteToEmmitFiles := false

	gen := context.GetFromRegistry("route").GetCtx("api").(*generator.RouteGenerator)

	if gen == nil {
		addRouteToEmmitFiles = true
		context.GetFromRegistry("route").AddToCtx("api", generator.NewRouteGenerator())
		gen = context.GetFromRegistry("route").GetCtx("api").(*generator.RouteGenerator)
	}

	gen.AddDefaultRestRoutes(identifier, controllerConfig.HttpMethods)

	if addRouteToEmmitFiles {
		return core.NewPhpEmitterFile("asher_api.php", api.RouteFilePath, gen, api.RouterFile)
	}

	return nil
}
