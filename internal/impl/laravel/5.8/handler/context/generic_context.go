package context

type GenericContext struct {
	BaseContext
	genericMap map[string]interface{}
}

/**
 Creates a New instance of a GenericContext that holds an interface
 Returns:
	- A pointer to A GenericContext
 Usage:
	- ctx := NewGenericContext()
*/
func NewGenericContext() *GenericContext {
	return &GenericContext{
		genericMap: make(map[string]interface{}),
	}
}

func (g *GenericContext) AddToCtx(key string, value interface{}) {
	g.genericMap[key] = value
}

func (g *GenericContext) GetCtx(key string) interface{} {
	return g.genericMap[key]
}
