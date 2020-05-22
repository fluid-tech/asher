package generator

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/helper"
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
	- Input: buildRelationshipFunction(1, 'Orders', 'OrderProducts', 'order_id', 'id')
*/
func (relationshipModel *RelationshipModel) AddRelationshipToModel(relationshipType int, currentTableName string, referenceTableName string, foreignKey string, primaryKey string) *helper.RelationshipDetail {

	generatedFunction := relationshipModel.buildRelationshipFunction(relationshipType, referenceTableName, foreignKey, primaryKey)
	relationshipDetailObj := helper.NewRelationshipDetail(relationshipType, generatedFunction, foreignKey, referenceTableName)
	relationshipModel.modGen.classBuilder.AddFunction(generatedFunction)
	return relationshipDetailObj

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
	- Input: buildRelationshipFunction(1, 'OrderProducts', 'order_id', 'id')
	- Output: output.String() gives :
		function OrderProducts() {
			return $this->hasMany('App\OrderProducts', 'order_id', 'id')
		}
*/
func (relationshipModel *RelationshipModel) buildRelationshipFunction(relationshipType int, referenceTableName string, foreignKey string, primaryKey string) *core.Function {
	relation := ""
	if relationshipType == helper.HasManny {
		relation = "hasMany"
	} else if relationshipType == helper.HasOne {
		relation = "hasOne"
	} else {
		panic("This type of relation is not supported [relationshipType Number: ]" + string(relationshipType))
	}
	returnStatementStringFormatter := fmt.Sprintf(`return $this->%s('App\%s','%s','%s')`, relation, referenceTableName, foreignKey, primaryKey)
	statement := api.TabbedUnit(core.NewSimpleStatement(returnStatementStringFormatter))
	return builder.NewFunctionBuilder().SetName(referenceTableName).AddStatement(statement).GetFunction()
}
