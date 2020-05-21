package api

/**
 A writer will be used by AsherWalker to write the files for the given respective EmitterFiles.
 The main task of Walker is to identify the type of EmitterFile and delegate it to it's respective handlers.
 These handlers will then write the content of those EmitterFile after performing some pre-processing, if required.
 */
type Writer struct {
	registry
}

/**
 Used to walk through all the given emitterFiles and delegate them to their respective WriterHandlers.
 Parameters:
	- emitterFiles: An collection of EmitterFile that needs to be written on the current project.
 Returns:
 	- true if all the given EmitterFiles were written successfully on the current project.
*/
Walk(emitterFiles [] EmitterFile)	bool
