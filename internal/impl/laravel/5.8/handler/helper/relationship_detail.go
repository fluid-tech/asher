package helper

import (
	"asher/internal/api/codebuilder/php/core"
	"fmt"
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

/**
Just For Printing Purpose.
*/
func (relationshipDetail *RelationshipDetail) String() {
	fmt.Println("RelationShipType : ", relationshipDetail.RelationshipType)
	fmt.Println("Function : ", relationshipDetail.Function.String())
	fmt.Println("Foreign Key : ", relationshipDetail.ForeignKey)
	fmt.Println("ReferenceModel Name : ", relationshipDetail.ReferencingModel)
}
