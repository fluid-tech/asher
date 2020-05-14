package api

// Emits level 1 nodes in the JSON
type Emitter interface {
	/**
	Callback triggered by AsherWalker.
	value is usually the data stored in a model object
	*/
	Emit(value interface{})
	/**
	Callback triggered by AsherWalker once it emits all nodes.
	This contains the path and the contents of the files to write
	*/
	GetFileMap() map[string]*EmitterFile
}
