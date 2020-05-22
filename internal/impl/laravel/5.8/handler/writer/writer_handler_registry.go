package writer

import "asher/internal/api"

// A static object of WriterHandlerRegistry to maintain Singleton Pattern.
var registryObj *WriterHandlerRegistry = nil

// Holds a collection of handlers supported by this framework.
var handlersMap = map[int]api.WriterHandler{
	api.Model: NewMigrationWriterHandler(),
}

/**
Laravel's implementation for WriterHandlerRegistry. For more details [@see api/writer_handler_registry.go].
*/
type WriterHandlerRegistry struct {
	api.WriterHandlerRegistry
	handlers map[int]api.WriterHandler
}

/**
 Creates a single instance of the WriterHandlerRegistry.
 Returns:
	- instance of WriterHandlerRegistry.
*/
func NewWriterHandlerRegistry() *WriterHandlerRegistry {
	// We'll create a new instance of WriterHandlerRegistry, if no instance has been created yet.
	if registryObj == nil {
		registryObj = &WriterHandlerRegistry{
			handlers: handlersMap,
		}
	}
	return registryObj
}

/**
 Fetches an instance of the corresponding handler for the given handlerType. For more details
 [@see api/writer_handler_registry.go].
 Parameters:
	- handlerType: type of the handlers that is required. Value must be one of the supported types in api/emitter_file.
 Example:
	- registry.GetFromRegistry(api.Migration)
*/
func (registry *WriterHandlerRegistry) GetFromRegistry(handlerType int) api.WriterHandler {
	return registry.handlers[handlerType]
}


const PhpTag = "<?php\n"
const PhpExtension = ".php"
