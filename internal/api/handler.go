package api

type Handler interface {
	Handle(value interface{}) (EmitterFile, error)
}