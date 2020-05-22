package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"flag"
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
		fileToEmit = append(fileToEmit,controllerHandler.HandleRestController(identifier, value)...)
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
func (controllerHandler *ControllerHandler) HandleRestController(
	identifier string, value interface{}) []api.EmitterFile {

	var fileToEmit []api.EmitterFile
	controllerConfig := value.(models.Controller)

	conGen := generator.NewControllerGenerator()
	conGen.SetIdentifier(identifier)

	conGen.AddFunctionsInController(controllerConfig.HttpMethods)

	context.GetFromRegistry("controller").AddToCtx(identifier+`RestController`, conGen)

	controllerEmitterFile := core.NewPhpEmitterFile(identifier+"RestController",  api.ControllerPath, conGen, api.Controller)
	transactorEmitterFiles := controllerHandler.HandleTransactor(identifier,value)
	mutatorEmitterFiles := controllerHandler.HandleMutator(identifier)
	queryEmitterFiles := controllerHandler.HandleQuery(identifier)
	routeEmitterFiles := controllerHandler.HandleRoutes(identifier,value)

	fileToEmit = append(fileToEmit, controllerEmitterFile)
	fileToEmit = append(fileToEmit, transactorEmitterFiles...)
	fileToEmit = append(fileToEmit, mutatorEmitterFiles...)
	fileToEmit = append(fileToEmit, queryEmitterFiles...)
	fileToEmit = append(fileToEmit, routeEmitterFiles...)


	return fileToEmit
}

func (controllerHandler *ControllerHandler) HandleTransactor(identifier string, value interface{}) []api.EmitterFile {
	controllerConfig := value.(models.Controller)
	var filesToEmmit []api.EmitterFile
	var transactorEmmiterFile []api.EmitterFile


	modelGen := context.GetFromRegistry("model").GetCtx(identifier).(generator.ModelGenerator)
	migrationGen := context.GetFromRegistry("migration").GetCtx(identifier).(generator.MigrationGenerator)
	/*SWITCH CASE FOR TYPE OF TRANSACTOR*/

	switch controllerConfig.Type {
		case "file":
			transactorEmmiterFile = controllerHandler.HandleFileTransactor(identifier,value,modelGen,migrationGen)
		case "image":
			transactorEmmiterFile = controllerHandler.HandleImageTransactor(identifier,value,modelGen,migrationGen)
		default:
			transactorEmmiterFile = controllerHandler.HandleDefaultTransactor(identifier,value,modelGen,migrationGen)
	}
	filesToEmmit = append(filesToEmmit,transactorEmmiterFile...)

	return filesToEmmit
}

func (controllerHandler *ControllerHandler) HandleDefaultTransactor(identifier string, value interface{},
modelGen generator.ModelGenerator, migrationGen generator.MigrationGenerator) []api.EmitterFile {
	controller := value.(models.Controller)
	var filesToEmmit []api.EmitterFile

	transactorGen := generator.NewTransactorGenerator()
	transactorGen.SetIdentifier(identifier)
	transactorGen.SetTransactorType("default")

	context.GetFromRegistry("transactor").AddToCtx(identifier+"Transactor", transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor",  api.TransactorPath, transactorGen, api.Transactor)

	filesToEmmit = append(filesToEmmit, transactorEmitter)
	return filesToEmmit
}

func (controllerHandler *ControllerHandler) HandleFileTransactor(identifier string, value interface{},
modelGen generator.ModelGenerator, migrationGen generator.MigrationGenerator) []api.EmitterFile {
	controller := value.(models.Controller)
	var filesToEmmit []api.EmitterFile


	modelGen.AddCreateValidationRule("file_urls", "required").
		AddUpdateValidationRule("file_urls.urls", "array")

	modelGen.AddUpdateValidationRule("file_urls", "required").
		AddUpdateValidationRule("file_urls.urls", "array")

	modelGen.AddFillable("fileurls")

	migrationGen.AddColumn(core.NewSimpleStatement(`$table->longText('img_urls')->nullable();`))

	transactorGen := generator.NewTransactorGenerator()
	transactorGen.SetIdentifier(identifier)
	transactorGen.SetTransactorType("file")

	context.GetFromRegistry("transactor").AddToCtx(identifier+"Transactor", transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor",  api.TransactorPath, transactorGen, api.Transactor)

	filesToEmmit = append(filesToEmmit, transactorEmitter)
	return filesToEmmit
}

func (controllerHandler *ControllerHandler) HandleImageTransactor(identifier string, value interface{},
modelGen generator.ModelGenerator, migrationGen generator.MigrationGenerator) []api.EmitterFile {
	controller := value.(models.Controller)
	var filesToEmmit []api.EmitterFile

	modelGen.AddCreateValidationRule("img_urls", "required").
		AddUpdateValidationRule("img_urls.urls", "array")

	modelGen.AddUpdateValidationRule("img_urls", "required").
		AddUpdateValidationRule("img_urls.urls", "array")


	modelGen.AddFillable("imgurls")
	migrationGen.AddColumn("")

	transactorGen := generator.NewTransactorGenerator()
	transactorGen.SetIdentifier(identifier)
	transactorGen.SetTransactorType("image")

	context.GetFromRegistry("transactor").AddToCtx(identifier+"Transactor", transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(identifier+"Transactor",  api.TransactorPath, transactorGen, api.Transactor)

	filesToEmmit = append(filesToEmmit, transactorEmitter)
	return filesToEmmit
}

func (controllerHandler *ControllerHandler) HandleMutator(identifier string) ([]api.EmitterFile) {
	var filesToEmit []api.EmitterFile
	mutatorGen := generator.NewMutatorGenerator()
	mutatorGen.SetIdentifier(identifier)

	context.GetFromRegistry("controller").AddToCtx(identifier+"Mutator", mutatorGen)
	mutatorEmitter :=
		core.NewPhpEmitterFile(identifier+"Mutator",  api.MutatorPath, mutatorGen, api.Mutator)

	filesToEmit = append(filesToEmit, mutatorEmitter)
	return filesToEmit
}

func (controllerHandler *ControllerHandler) HandleQuery(
	identifier string) []api.EmitterFile {

	var filesToEmit []api.EmitterFile

	queryGenerator := generator.NewQueryGenerator(identifier, true)

	context.GetFromRegistry("query").AddToCtx(identifier, queryGenerator)

	emitFile := core.NewPhpEmitterFile(identifier+"Query.php", "/app/query", queryGenerator, 1)
	filesToEmit = append(filesToEmit, emitFile)

	return filesToEmit
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
func (controllerHandler *ControllerHandler) HandleRoutes(identifier string, value interface{}) []api.EmitterFile {
	var filesToEmit []api.EmitterFile
	controller := value.(models.Controller)
	addRouteToEmmitFiles := false
	gen := context.GetFromRegistry("route").GetCtx("api")
	if gen == nil {
		addRouteToEmmitFiles = true
		context.GetFromRegistry("route").AddToCtx("api", context.NewRouteContext())
		gen = context.GetFromRegistry("route").GetCtx("api")
	}
	apiGenerator := gen.(generator.RouteGenerator)
	apiGenerator.AddDefaultRestRoutes(identifier, controller)

	if addRouteToEmmitFiles {
		emitFile := core.NewPhpEmitterFile("asher_api.php", "/routes", apiGenerator, 1)
		filesToEmit = append(filesToEmit, emitFile)
	}

	return filesToEmit
}
