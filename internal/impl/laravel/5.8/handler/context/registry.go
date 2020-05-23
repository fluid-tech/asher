package context

var registry = map[string]BaseContext{
	"migration":  NewMigrationContext(),
	"model":      NewModelContext(),
	"route":      NewRouteContext(),
	"controller": NewControllerContext(),
	"mutator":    NewMutatorContext(),
	"transactor": NewTransactorContext(),
	"query":      NewQueryContext(),
}

/*
Fetches a BaseContext implementation from the registry. This method
exists only to avoid writes to this map from outside this package
*/
func GetFromRegistry(key string) BaseContext {
	return registry[key]
}
