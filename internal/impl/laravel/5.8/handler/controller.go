package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/models"
)

type ControllerHandler struct {
	api.Handler
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{}
}

func (controllerHandler *ControllerHandler) Handle(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	var emitFiles []*api.EmitterFile
	routeFiles, _ := controllerHandler.HandleRoutes(identifier,value)
	emitFiles = append(emitFiles,routeFiles...)
}

func (controllerHandler *ControllerHandler) HandleController(identifier string, value interface{}) ([]*api.EmitterFile, error) {

}

func (controllerHandler *ControllerHandler) HandleTransactor(identifier string, value interface{}) ([]*api.EmitterFile, error) {

}

func (controllerHandler *ControllerHandler) HandleMutator(identifier string, value interface{}) ([]*api.EmitterFile, error) {

}

func (controllerHandler *ControllerHandler) HandleQuery(identifier string, value interface{}) ([]*api.EmitterFile, error) {

}

func (controllerHandler *ControllerHandler) HandleRoutes(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	var filesToEmit []*api.EmitterFile
	//controller := value.(models.Controller)
	addRouteToEmmitFiles := false
	gen := context.GetFromRegistry("route").GetCtx("api")
	if gen == nil{
		addRouteToEmmitFiles = true
		context.GetFromRegistry("route").AddToCtx("api", context.NewRouteContext())
		gen = context.GetFromRegistry("route").GetCtx("api")
	}
	apiGenerator := gen.(generator.RouteGenerator)
	apiGenerator.AddDefaultRestRoutes(identifier)

	if addRouteToEmmitFiles {
		emitFile:= api.EmitterFile(core.NewPhpEmitterFile("asher_api.php","/routes",nil,1))
		filesToEmit=append(filesToEmit,&emitFile)
	}

	return filesToEmit, nil
}


