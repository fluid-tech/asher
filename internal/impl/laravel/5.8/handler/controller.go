package handler
import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"errors"
	"fmt"
	"strings"
)

const controllerFileName = "%sRestController"
const transactorFileName = "%sTransactor"
const mutatorFileName = "%sMutator"
const queryFileName = "%sQuery"
const apiContext = "api"

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
func (controllerHandler *ControllerHandler) Handle(identifier string, value interface{}) ([]api.EmitterFile, error) {

	controllerConfig := value.(models.Controller)
	var fileToEmit []api.EmitterFile

	if controllerConfig.Rest {
		restControllerEmitterFiles, err := controllerHandler.handleRestController(identifier, controllerConfig)
		if err != nil {
			return nil,err
		}
		fileToEmit = append(fileToEmit, restControllerEmitterFiles...)
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
func (controllerHandler *ControllerHandler) handleRestController(identifier string,
	controllerConfig models.Controller) ([]api.EmitterFile,error) {


	var fileToEmit []api.EmitterFile
	conGen := generator.NewControllerGenerator()


	conGen.SetIdentifier(identifier).
		AddFunctionsInController(controllerConfig.HttpMethods)
	context.GetFromRegistry(context.Controller).AddToCtx(identifier, conGen)
	controllerEmitterFile := core.NewPhpEmitterFile(fmt.Sprintf(controllerFileName,identifier),
		api.ControllerPath, conGen, api.Controller)
	transactorEmitterFile,err := controllerHandler.handleTransactor(identifier, controllerConfig.Type)
	if err != nil {
		return nil, err
	}
	mutatorEmitterFile := controllerHandler.handleMutator(identifier)
	queryEmitterFile := controllerHandler.handleQuery(identifier)
	routeEmitterFile := controllerHandler.handleRoutes(identifier, controllerConfig.HttpMethods)
	fileToEmit = append(fileToEmit, controllerEmitterFile, transactorEmitterFile, mutatorEmitterFile, queryEmitterFile)
	/*AS ROUTE FILE IS EMITTED ONLY ONCE IE FOR THE FIRST TIME AFTER THAT IT IS ONLY USED*/
	if routeEmitterFile != nil {
		fileToEmit = append(fileToEmit, routeEmitterFile)
	}

	return fileToEmit,nil
}


/*REFER DOCS FOR TYPE OF TRANSACTOR*/
func (controllerHandler *ControllerHandler) handleTransactor(identifier string,
	controllerType string) (api.EmitterFile,error) {

	var transactorEmmiterFile api.EmitterFile
	retrievedModelGen := context.GetFromRegistry(context.Model).GetCtx(identifier)
	if retrievedModelGen == nil{
		return nil,errors.New(fmt.Sprintf("model class %s not found", identifier))
	}
	modelGen := retrievedModelGen.(*generator.ModelGenerator)


	retrievedMigGen := context.GetFromRegistry(context.Migration).GetCtx(identifier)
	if retrievedMigGen == nil{
		return nil,errors.New(fmt.Sprintf("migration class %s not found", identifier))
	}
	migrationGen := retrievedMigGen.(*generator.MigrationGenerator)

	/*SWITCH CASE FOR TYPE OF TRANSACTOR*/
	switch strings.ToLower(controllerType) {
	case "file":
		transactorEmmiterFile = controllerHandler.handleFileTransactor(identifier, modelGen, migrationGen)
	case "image":
		transactorEmmiterFile = controllerHandler.handleImageTransactor(identifier, modelGen, migrationGen)
	default:
		transactorEmmiterFile = controllerHandler.handleDefaultTransactor(identifier)
	}

	return transactorEmmiterFile,nil
}


func (controllerHandler *ControllerHandler) handleDefaultTransactor(identifier string) api.EmitterFile {
	//controller := value.(models.Controller)

	transactorGen := generator.NewTransactorGenerator("Base").SetIdentifier(identifier)

	context.GetFromRegistry(context.Transactor).AddToCtx(identifier, transactorGen)
	transactorEmitter := core.NewPhpEmitterFile(fmt.Sprintf(transactorFileName,identifier), api.TransactorPath,
		transactorGen, api.Transactor)

	return transactorEmitter
}

func (controllerHandler *ControllerHandler) handleFileTransactor(identifier string, modelGen *generator.ModelGenerator,
	migrationGen *generator.MigrationGenerator) api.EmitterFile {

	/*MODEL AND MIGRATION UPDATES*/
	generator.NewTransactorModel(modelGen).AddFileUrlsToFillAbles().AddFileUrlsValidationRules()
	generator.NewTransactorMigration(migrationGen).AddMigrationForFileUrls()

	/*BUILDING OF TRANSACTOR*/
	transactorGen := generator.NewTransactorGenerator("file").SetIdentifier(identifier)
	generator.NewFileTransactor(transactorGen).AddDefaults()

	context.GetFromRegistry(context.Transactor).AddToCtx(identifier, transactorGen)

	transactorEmitter := core.NewPhpEmitterFile(fmt.Sprintf(transactorFileName,identifier), api.TransactorPath,
		transactorGen, api.Transactor)

	return transactorEmitter
}

func (controllerHandler *ControllerHandler) handleImageTransactor(identifier string, modelGen *generator.ModelGenerator,
	migrationGen *generator.MigrationGenerator) api.EmitterFile {
	//controller := value.(models.Controller)

	/*MODEL AND MIGRATION UPDATES*/
	generator.NewTransactorModel(modelGen).AddFileUrlsToFillAbles().AddFileUrlsValidationRules()
	generator.NewTransactorMigration(migrationGen).AddMigrationForFileUrls()

	/*TRANSACTOR GENERATION*/
	transactorGen := generator.NewTransactorGenerator("image").SetIdentifier(identifier)
	generator.NewImageTransactor(transactorGen).AddDefaults()
	context.GetFromRegistry(context.Transactor).AddToCtx(identifier, transactorGen)
	transactorEmitter := core.NewPhpEmitterFile(fmt.Sprintf(transactorFileName,identifier),
		api.TransactorPath, transactorGen,
		api.Transactor)
	return transactorEmitter
}


func (controllerHandler *ControllerHandler) handleMutator(identifier string) api.EmitterFile {


	mutatorGen := generator.NewMutatorGenerator()
	mutatorGen.SetIdentifier(identifier)

	context.GetFromRegistry(context.Mutator).AddToCtx(identifier, mutatorGen)
	mutatorEmitter :=
		core.NewPhpEmitterFile(fmt.Sprintf(mutatorFileName,identifier), api.MutatorPath, mutatorGen, api.Mutator)

	return mutatorEmitter
}

func (controllerHandler *ControllerHandler) handleQuery(identifier string) api.EmitterFile {

	queryGenerator := generator.NewQueryGenerator(true).SetIdentifier(identifier)

	context.GetFromRegistry(context.Query).AddToCtx(identifier, queryGenerator)
	emitFile := core.NewPhpEmitterFile(fmt.Sprintf(queryFileName,identifier), api.QueryPath, queryGenerator, api.Query)

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
func (controllerHandler *ControllerHandler) handleRoutes(identifier string, httpMethods []string) api.EmitterFile {

	var actualGenerator *generator.RouteGenerator
	addRouteToEmmitFiles := false
	gen := context.GetFromRegistry(context.Route).GetCtx(apiContext)

	if gen == nil {
		addRouteToEmmitFiles = true
		actualGenerator = generator.NewRouteGenerator()
		context.GetFromRegistry(context.Route).AddToCtx(apiContext, actualGenerator)
	} else {
		actualGenerator = gen.(*generator.RouteGenerator)
	}
	actualGenerator.AddDefaultRestRoutes(identifier, httpMethods)
	if addRouteToEmmitFiles {
		return core.NewPhpEmitterFile("asher_api.php", api.RouteFilePath, actualGenerator, api.RouterFile)
	}

	return nil
}
