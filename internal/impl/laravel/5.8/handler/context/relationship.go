package context

import (
	"asher/internal/impl/laravel/5.8/handler/helper"
)

type Relationship struct {
	BaseContext
	relationshipDetailMap map[string][]*helper.RelationshipDetail
}

func NewRelationshipContext() *Relationship {
	return &Relationship{relationshipDetailMap: make(map[string][]*helper.RelationshipDetail)}
}

func (m *Relationship) AddToCtx(key string, value interface{}) {
	relationshipDetail := value.(*helper.RelationshipDetail)
	arrayOfRelationshipDetail := m.relationshipDetailMap[key]
	m.relationshipDetailMap[key] = append(arrayOfRelationshipDetail, relationshipDetail)
}

func (m *Relationship) GetCtx(key string) interface{} {
	return m.relationshipDetailMap[key]
}
