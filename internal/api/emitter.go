package api

// Emits level 1 nodes in the JSON
type Emitter interface {
	Emit(key string, value interface{})
}
