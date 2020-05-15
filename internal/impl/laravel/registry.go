package laravel

import (
	"asher/internal/api"
	"asher/internal/impl/laravel/5.8/handler"
)

var handlerRegistry = map[string]api.Handler{
	"auditCols": handler.NewAuditColHandler(),

}

/**
Returns a Handler instance from the registry.
This method exists only to avoid writes to this map from outside this package */
func GetFromRegistry(key string) api.Handler {
	return handlerRegistry[key]
}
