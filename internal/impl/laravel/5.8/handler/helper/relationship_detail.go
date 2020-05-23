package helper

import (
	"asher/internal/api/codebuilder/php/core"
)

const (
	HasOne  = 1
	HasMany = 2
)

type RelationshipDetail struct {
	RelationshipType int
	Function         *core.Function
	ForeignKey       string
	ReferencingModel string

}

// TODO: Dhano search Relationship context for order key returns array of RelationshipDetails
func NewRelationshipDetail(relationShipType int, function *core.Function, foreignKey string, referencingModel string) *RelationshipDetail {
	return &RelationshipDetail{
		RelationshipType: relationShipType,
		Function:         function,
		ForeignKey:       foreignKey,
		ReferencingModel: referencingModel,
	}
}

func (relationshipDetail *RelationshipDetail) GetRelationshipType() int {
	return relationshipDetail.RelationshipType
}
