package context

import "asher/internal/impl/laravel/5.8/handler/generator"

type Mutator struct {
	BaseContext
	mutatorGenerators map[string]*generator.MutatorGenerator
}

func NewMutatorContext() *Mutator {
	return &Mutator{}
}

/**
Store a MutatorInfo instance.
*/
func (mutator *Mutator) AddToCtx(key string, value interface{})  {
	mutator.mutatorGenerators[key] = value.(*generator.MutatorGenerator)
}

/**
Fetches a MutatorInfo instance
The user of this method must cast and fetch appropriate data
*/
func (mutator *Mutator) GetCtx(key string) interface{} {
	return mutator.mutatorGenerators[key]
}

