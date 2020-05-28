package context

import "asher/internal/impl/laravel/5.8/handler/generator"

type Query struct {
	BaseContext
	queryGenerators map[string]*generator.QueryGenerator
}

func NewQueryContext() *Query {
	return &Query{
		queryGenerators: make(map[string]*generator.QueryGenerator),
	}
}

/**
Store a ControllerInfo instance.
*/
func (query *Query) AddToCtx(key string, value interface{}) {
	query.queryGenerators[key] = value.(*generator.QueryGenerator)
}

/**
Fetches a ControllerInfo instance
The user of this method must cast and fetch appropriate data
*/
func (query *Query) GetCtx(key string) interface{} {
	return query.queryGenerators[key]
}
