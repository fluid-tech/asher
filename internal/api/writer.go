package api

/**
A writer will be used by AsherWalker to write the files for the given respective EmitterFiles.
The main task of Walker is to identify the type of EmitterFile and delegate it to it's respective handlers.
These handlers will then write the content of those EmitterFile after performing some pre-processing, if required.
*/
type Writer struct {
	registry WriterHandlerRegistry
}

/**
 Creates a new instance of Writer with the given WriterHandlerRegistry.
 Parameters:
	- 	registry: instance of the WriterHandlerRegistry.
 Returns:
	- A new instance of the Writer with the given WriterHandlerRegistry.
*/
func NewWriter(registry WriterHandlerRegistry) *Writer {
	return &Writer{
		registry: registry,
	}
}

/**
 Used to walk through all the given emitterFiles and delegate them to their respective WriterHandlers.
 Parameters:
	- emitterFiles: An collection of EmitterFile that needs to be written on the current project.
 Returns:
 	- true if all the given EmitterFiles were written successfully on the current project.
*/
func (writer *Writer) Walk(emitterFiles []*EmitterFile) bool {
	for _, emitterFile := range emitterFiles {
		writerHandler := writer.registry.GetFromRegistry((*emitterFile).FileType())
		// A nil writerHandler means that the framework doesn;t support the given type of WriterHandler
		if writerHandler == nil {
			return false
		}

		// Delegating the emitterFile to the respective WriterHandler
		writerHandler.BeforeHandle(*emitterFile)
		status := writerHandler.Handle(*emitterFile)
		writerHandler.AfterHandle(*emitterFile)

		// This means that no bytes were written and the operation failed.
		if !status {
			return false
		}
	}
	return true
}
