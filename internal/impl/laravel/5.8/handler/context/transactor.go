package context

import "asher/internal/impl/laravel/5.8/handler/generator"

type Transactor struct {
	BaseContext
	transactorGenerators map[string]*generator.TransactorGenerator
}

func NewTransactorContext() *Controller {
	return &Controller{}
}

/**
Store a ControllerInfo instance.
*/
func (transactor *Transactor) AddToCtx(key string, value interface{})  {
	transactor.transactorGenerators[key] = value.(*generator.TransactorGenerator)
}

/**
Fetches a ControllerInfo instance
The user of this method must cast and fetch appropriate data
*/
func (transactor *Transactor) GetCtx(key string) interface{} {
	return transactor.transactorGenerators[key]
}

