package context

import "asher/internal/impl/laravel/5.8/handler/generator"

type Controller struct {
	BaseContext
	controllerGenerators map[string]*generator.ControllerGenerator
}

func NewControllerContext() *Controller {
	return &Controller{}
}

/**
Store a ControllerInfo instance.
*/
func (controller *Controller) AddToCtx(key string, value interface{})  {
	controller.controllerGenerators[key] = value.(*generator.ControllerGenerator)
}

/**
Fetches a ControllerInfo instance
The user of this method must cast and fetch appropriate data
*/
func (controller *Controller) GetCtx(key string) interface{} {
	return controller.controllerGenerators[key]
}
