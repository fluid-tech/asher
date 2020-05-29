package context

const (
	Migration  = "migration"
	Model      = "model"
	Route      = "route"
	Controller = "controller"
	Mutator    = "mutator"
	Transactor = "transactor"
	Query      = "query"
	Relation   = "relation"
)

var registry = map[string]BaseContext{
	Migration:  NewGenericContext(),
	Model:      NewGenericContext(),
	Route:      NewGenericContext(),
	Controller: NewGenericContext(),
	Mutator:    NewGenericContext(),
	Transactor: NewGenericContext(),
	Query:      NewGenericContext(),
	Relation:   NewGenericContext(),
}

/*
Fetches a BaseContext implementation from the registry. This method
exists only to avoid writes to this map from outside this package
*/
func GetFromRegistry(key string) BaseContext {
	return registry[key]
}
