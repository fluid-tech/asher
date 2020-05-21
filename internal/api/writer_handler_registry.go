package api

/**
Maintains a collection of handlers of all types. Each and every framework must have an implementation for this
interface. This would be used by the Writer to fetch handler for the required framework. It is recommended to use a map
to store the WriterHandlers with the handlerType as the key.
***********************************************************************************************************************
Note: All the implementation must strictly follow singleton pattern for the implementation of this interface.
*/
type WriterHandlerRegistry interface {
	/**
	 Fetches the corresponding WriterHandler for the given handlerType.
	 Parameters:
		- handlerType: the type of the handler required. The value must be one of the FileTypes specified in the
					   emitter_file. To view all the supported types [@see /api/emitter_file.go],
	 Returns:
		- instance of the corresponding WriterHandler for the given handlerType or nil if no such handlerType is found.
	 Example:
		- GetFromRegistry(api.Model)
	*/
	GetFromRegistry(handlerType int) WriterHandler
}
