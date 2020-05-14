package context

type BaseContext interface{
	AddToCtx(key string, value interface{})
	GetCtx(key string) interface{}
}
