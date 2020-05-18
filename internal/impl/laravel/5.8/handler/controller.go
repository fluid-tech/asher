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
	routeFiles, _ := controllerHandler.HandleRoutes(identifier,value)
	emitFiles = append(emitFiles,routeFiles...)
	return  emitFiles,nil
}

func (controllerHandler *ControllerHandler) HandleRestController(
	identifier string, value interface{}) (api.EmitterFile, error) {
	generatorObject := generator.NewControllerGenerator().BuildRestController().String()
	fmt.Printf(generatorObject)
	tempRouteEmitter := api.EmitterFile(core.NewPhpEmitterFile("","",nil,1))
	return tempRouteEmitter, nil
}

func (controllerHandler *ControllerHandler) HandleTransactor(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {
	return nil,nil
}

func (controllerHandler *ControllerHandler) HandleMutator(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {
	return nil,nil
}

func (controllerHandler *ControllerHandler) HandleQuery(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {
	return nil,nil
}

func (controllerHandler *ControllerHandler) HandleRoutes(
	identifier string, value interface{}) ([]*api.EmitterFile, error) {
	var emitFiles []*api.EmitterFile
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
		tempRouteEmitter := api.EmitterFile(core.NewPhpEmitterFile("","",nil,1))
		emitFiles=append(emitFiles, &tempRouteEmitter)
	}

	return emitFiles, nil
}
