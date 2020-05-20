package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
	"fmt"
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
	if tempValue.Rest {
		controllerHandler.HandleRestController(identifier, value)
	}

	var emitFiles []*api.EmitterFile
	routeFiles, _ := controllerHandler.HandleRoutes(identifier, value)
	emitFiles = append(emitFiles, routeFiles...)
	return emitFiles, nil
}

func (controllerHandler *ControllerHandler) HandleRestController(
	identifier string, value interface{}) (api.EmitterFile, error) {
	controller := value.(models.Controller)

	controllerGenerator := generator.NewControllerGenerator(controller)
	controllerGenerator.SetIdentifier(identifier)
	context.GetFromRegistry("controller").AddToCtx(identifier+`RestController`, controllerGenerator)
	tempRouteEmitter := core.NewPhpEmitterFile("", "", nil, 1)
	return tempRouteEmitter, nil
}

func (controllerHandler *ControllerHandler) HandleTransactor(
	identifier string, value interface{}) ([]api.EmitterFile, error) {
	fileToEmitt := []api.EmitterFile{}
	transactorGenerator := generator.NewTransactorGenerator()
	transactorGenerator.SetIdentifier(identifier)

	context.GetFromRegistry("controller").AddToCtx(identifier+"transactor", transactorGenerator)
	tempRouteEmitter := core.NewPhpEmitterFile("", "", nil, 1)

	fileToEmitt = append(fileToEmitt, tempRouteEmitter)
	return fileToEmitt, nil
}

func (controllerHandler *ControllerHandler) HandleMutator(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {
	return nil, nil
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
