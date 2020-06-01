package laravel

import (
	"asher/internal/api"
	"asher/internal/impl/laravel/5.8/handler"
)

const (
	HandlerAuditCols  = "auditCols"
	HandlerController = "controller"
	HandlerColumns    = "columns"
	HandlerRelation   = "relation"
)

var handlerRegistry = map[string]api.Handler{
	HandlerAuditCols:  handler.NewAuditColHandler(),
	HandlerController: handler.NewControllerHandler(),
	HandlerColumns:    handler.NewColumnHandler(),
	HandlerRelation:   handler.NewRelationshipHandler(),
}

/**
Returns a Handler instance from the registry.
This method exists only to avoid writes to this map from outside this package */
func GetFromRegistry(key string) api.Handler {
	return handlerRegistry[key]
}
