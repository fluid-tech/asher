package context

import (
	"asher/internal/impl/laravel/5.8/handler/generator"
)

type Model struct {
	BaseContext
	modelMap map[string]*generator.ModelGenerator
}

func NewModelContext() *Model {
	return &Model{modelMap: make(map[string]*generator.ModelGenerator)}
}

func (m *Model) AddToCtx(key string, value interface{}) {
	m.modelMap[key] = value.(*generator.ModelGenerator)
}

func (m *Model) GetCtx(key string) interface{} {
	return m.modelMap[key]
}
