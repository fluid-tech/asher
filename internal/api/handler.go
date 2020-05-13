package api

type Handler interface {
	Handle(key string , value interface{}) (fileName string, contents []string)
}