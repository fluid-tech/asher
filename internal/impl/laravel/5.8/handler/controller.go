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

func (controllerHandler *ControllerHandler) Handle(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {

	tempValue := value.(models.Controller)
	fileToEmit := []*api.EmitterFile{}
	if tempValue.Rest {
		fileToEmit = append(fileToEmit,controllerHandler.HandleRestController(identifier, value)...)
	}
	fileToEmit = append(fileToEmit, controllerHandler.HandleMutator(identifier)...)
	fileToEmit = append(fileToEmit, controllerHandler.HandleTransactor(identifier)...)
	fileToEmit = append(fileToEmit, controllerHandler.HandleRestController(identifier,value)...)


	routeFiles, _ := controllerHandler.HandleRoutes(identifier, value)
	fileToEmit = append(fileToEmit, routeFiles...)
	return fileToEmit, nil
}

func (controllerHandler *ControllerHandler) HandleRestController(
	identifier string, value interface{}) ([]*api.EmitterFile) {

	fileToEmiit := []*api.EmitterFile{}
	controller := value.(models.Controller)

	conGen := generator.NewControllerGenerator()
	conGen.SetIdentifier(identifier)
	if controller.HttpMethods != nil {
		if len(controller.HttpMethods) >= 0 {
			conGen.AddConstructorFunction()
			for _, element := range controller.HttpMethods {
				switch strings.ToLower(element) {
				case "post":
					conGen.AddCreateFunction()
				case "get":
					conGen.AddFindByIdFunction().AddGetAllFunction()
				case "put":
					conGen.AddUpdateFunction()
				case "delete":
					conGen.AddDeleteFunction()
				}
			}
		}
	} else {
		conGen.AddAllFunctionsInController()
	}
	context.GetFromRegistry("controller").AddToCtx(identifier+`RestController`, conGen)
	conGenRef := api.Generator(conGen)
	tempRouteEmitter := api.EmitterFile(
		core.NewPhpEmitterFile(identifier+"RestController",  api.ControllerPath, &conGenRef, api.Controller))
	fileToEmiit = append(fileToEmiit, &tempRouteEmitter)
	return fileToEmiit
}

func (controllerHandler *ControllerHandler) HandleTransactor(identifier string) ([]*api.EmitterFile) {
	filesToEmmit := []*api.EmitterFile{}
	transactorGen := generator.NewTransactorGenerator()
	transactorGen.SetIdentifier(identifier)

	context.GetFromRegistry("controller").AddToCtx(identifier+"Transactor", transactorGen)
	transactorGenRef := api.Generator(transactorGen)
	tempRouteEmitter := api.EmitterFile(core.NewPhpEmitterFile(identifier+"Transactor",  api.TransactorPath, &transactorGenRef, api.Transactor))

	filesToEmmit = append(filesToEmmit, &tempRouteEmitter)
	return filesToEmmit
}

func (controllerHandler *ControllerHandler) HandleMutator(identifier string) ([]*api.EmitterFile) {
	filesToEmit := []*api.EmitterFile{}
	mutatorGen := generator.NewMutatorGenerator()
	mutatorGen.SetIdentifier(identifier)

	context.GetFromRegistry("controller").AddToCtx(identifier+"Mutator", mutatorGen)
	mutatorGenRef := api.Generator(mutatorGen)
	tempRouteEmitter := api.EmitterFile(
		core.NewPhpEmitterFile(identifier+"Mutator",  api.MutatorPath, &mutatorGenRef, api.Mutator))

	filesToEmit = append(filesToEmit, &tempRouteEmitter)
	return filesToEmit
}

func (controllerHandler *ControllerHandler) HandleQuery(
	identifier string, value interface{}) ([]api.EmitterFile, error) {

	var filesToEmit []api.EmitterFile

	queryGenerator := generator.NewQueryGenerator(identifier, true)

	context.GetFromRegistry("query").AddToCtx(identifier, queryGenerator)

	emitFile := core.NewPhpEmitterFile(identifier+"Query.php", "/app/query", queryGenerator.String(), 1)
	filesToEmit = append(filesToEmit, emitFile)

	return filesToEmit, nil

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
func (controllerHandler *ControllerHandler) HandleRoutes(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	var filesToEmit []*api.EmitterFile
	controller := value.(models.Controller)
	addRouteToEmmitFiles := false
	gen := context.GetFromRegistry("route").GetCtx("api")
	if gen == nil {
		addRouteToEmmitFiles = true
		context.GetFromRegistry("route").AddToCtx("api", context.NewRouteContext())
		gen = context.GetFromRegistry("route").GetCtx("api")
	}
	apiGenerator := gen.(generator.QueryGenerator)
	apiGenerator.AddDefaultRestRoutes(identifier, controller)

	if addRouteToEmmitFiles {
		emitFile := api.EmitterFile(core.NewPhpEmitterFile("asher_api.php", "/routes", nil, 1))
		filesToEmit = append(filesToEmit, &emitFile)
	}

	return filesToEmit, nil
}
