package api

type Handler interface {
	/**
	Callback triggered by an emitter to process data
	*/
	Handle(value interface{}) (EmitterFile, error)
}
