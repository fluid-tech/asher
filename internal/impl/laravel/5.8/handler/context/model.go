package context

import "asher/internal/api/codebuilder/php/core"

type Model struct {
	BaseContext
	modelMap map[string]*core.Class
}

func NewModelContext() *Model {
	return &Model{modelMap: make(map[string]*core.Class)}
}

func (m *Model) AddToCtx(key string, value interface{}) {
	m.modelMap[key] = value.(*core.Class)
}

func (m *Model) GetCtx(key string) interface{} {
	return m.modelMap[key]
}
