package context

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
)

type Route struct {
	BaseContext
	routeGenerators map[string]*generator.QueryGenerator
}

func NewRouteContext() *Route {
	return &Route{}
}

/**
Store a MigrationInfo instance.
*/
func (route *Route) AddToCtx(key string, value interface{})  {
	route.routeGenerators[key] = value.(*generator.QueryGenerator)
}

/**
Fetches a MigrationInfo instance
The user of this method must cast and fetch appropriate data
*/
func (route *Route) GetCtx(key string) interface{} {
	return route.routeGenerators[key]
}




