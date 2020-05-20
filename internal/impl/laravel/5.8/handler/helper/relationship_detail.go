package helper

import (
	"asher/internal/api/codebuilder/php/core"
)

const (
	HasOne   = 1
	HasManny = 2
)

type RelationshipDetail struct {
	relationshipType int

	Function         *core.Function
}

func NewRelationshipDetail(relationShipType int) *RelationshipDetail {
	return &RelationshipDetail{
		relationshipType: relationShipType,
		Function:         nil,
	}
}

func (relationshipDetail *RelationshipDetail) RelationshipType() int {
	return relationshipDetail.relationshipType
}
