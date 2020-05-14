package api

type Handler interface {
	/**
	Callback triggered by an emitter to process data
	*/
	Handle(objectIdentifier string, value interface{}) ([]*EmitterFile, error)
}
