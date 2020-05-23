package context

var registry = map[string]BaseContext{
	"migration": NewMigrationContext(),
	"model":     NewModelContext(),
	"relation":  NewRelationshipContext(),
}

/*
Fetches a BaseContext implementation from the registry. This method
exists only to avoid writes to this map from outside this package
*/
func GetFromRegistry(key string) BaseContext {
	return registry[key]
}
