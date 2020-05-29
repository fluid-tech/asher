package context

const (
	ContextMigration  = "migration"
	ContextModel      = "model"
	ContextRoute      = "route"
	ContextController = "controller"
	ContextMutator    = "mutator"
	ContextTransactor = "transactor"
	ContextQuery      = "query"
	ContextRelation   = "relation"
)

var registry = map[string]BaseContext{
	ContextMigration:  NewGenericContext(),
	ContextModel:      NewGenericContext(),
	ContextRoute:      NewGenericContext(),
	ContextController: NewGenericContext(),
	ContextMutator:    NewGenericContext(),
	ContextTransactor: NewGenericContext(),
	ContextQuery:      NewGenericContext(),
	ContextRelation:   NewGenericContext(),
}

/*
Fetches a BaseContext implementation from the registry. This method
exists only to avoid writes to this map from outside this package
*/
func GetFromRegistry(key string) BaseContext {
	return registry[key]
}
