package context

import "asher/internal/impl/laravel/5.8/handler/generator"

type Transactor struct {
	BaseContext
	transactorGenerators map[string]*generator.TransactorGenerator
}

func NewTransactorContext() *Transactor {
	return &Transactor{
		transactorGenerators: make(map[string]*generator.TransactorGenerator),
	}
}

/**
Store a TransactorInfo instance.
*/
func (transactor *Transactor) AddToCtx(key string, value interface{}) {
	transactor.transactorGenerators[key] = value.(*generator.TransactorGenerator)
}

/**
Fetches a TransactorInfo instance
The user of this method must cast and fetch appropriate data
*/
func (transactor *Transactor) GetCtx(key string) interface{} {
	return transactor.transactorGenerators[key]
}
