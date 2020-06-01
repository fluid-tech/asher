package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"errors"
	"fmt"
)

type RelationshipModel struct {
	api.Generator
	modGen *ModelGenerator
}

func NewRelationshipModel(modelGenerator *ModelGenerator) *RelationshipModel {
	return &RelationshipModel{
		modGen: modelGenerator,
	}
}

/**
This Method creates a new RelationshipDetail Object and add it to the ModelGenerator
Parameters:
	- relationshipType: Sets if relation is hasMay or hasOne
	- currentTableName: Name of the Model within which this relation exists.
	- referenceTableName: Name of the Function that is being generated.
	- foreignKey: The ForeignKey to which it refer in ReferenceTable.
	- primaryKey: The PrimaryKey in Current Table.
Returns:
	- instance of the *RelationshipDetail
Example:
	- Input: AddRelationshipToModel(1, 'Orders', 'OrderProducts', 'order_id', 'id')
*/
func (relationshipModel *RelationshipModel) AddRelationshipToModel(relationshipType int, referenceTableName string, foreignKey string, primaryKey string) (*helper.RelationshipDetail, error) {

	generatedFunction, err := relationshipModel.buildRelationship(relationshipType, referenceTableName, foreignKey,
		primaryKey)
	if err != nil {
		return nil, err
	}
	relationshipDetailObj := helper.NewRelationshipDetail(relationshipType, generatedFunction, foreignKey,
		referenceTableName)
	relationshipModel.modGen.classBuilder.AddFunction(generatedFunction)
	return relationshipDetailObj, nil

}

/**
Generates a hasMay, hasOne Relationship Function that is to be appended to the Model File.
Parameters:
	- relationshipType: Sets if relation is hasMay or hasOne
	- referenceTableName: Name of the Function that is being generated.
	- foreignKey: The ForeignKey to which it refer in ReferenceTable.
	- primaryKey: The PrimaryKey in Current Table.
Returns:
	- instance of the *core.Function
Example:
	- Input: buildRelationship(1, 'OrderProducts', 'order_id', 'id')
	- Output: output.String() gives :
		function OrderProducts() {
			return $this->hasMany('App\OrderProducts', 'order_id', 'id')
		}
*/
func (relationshipModel *RelationshipModel) buildRelationship(relationshipType int, referenceTableName string,
	foreignKey string, primaryKey string) (*core.Function, error) {
	relation := ""
	if relationshipType == helper.HasMany {
		relation = "hasMany"
	} else if relationshipType == helper.HasOne {
		relation = "hasOne"
	} else {
		return nil, errors.New("This type of relation is not supported [relationshipType Number: ]" + string(relationshipType))
	}
	returnStatementStringFormatter := fmt.Sprintf(`return $this->%s('App\%s','%s','%s')`, relation,
		referenceTableName, foreignKey, primaryKey)
	statement := core.NewSimpleStatement(returnStatementStringFormatter)
	return builder.NewFunctionBuilder().SetName(referenceTableName).AddStatement(statement).GetFunction(), nil
}
