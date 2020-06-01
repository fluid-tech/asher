package handler

import (
	"asher/internal/api"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"asher/internal/models"
	"errors"
	"fmt"
	"strings"
)

type RelationshipHandler struct {
	api.Handler
}

func NewRelationshipHandler() *RelationshipHandler {
	return &RelationshipHandler{}
}

/**
This is The main Entry point for the Handler. This method retrieves the ModelGenerator from the context
and initiates the RelationshipModel with the fetched ModelGenerator.
it iterates over the array of hasMany and hasOne Relationship and process the relation and retrieves
tableName, PrimaryKey, ForeignKey from getRelationshipKeys method and calls addRelationshipToModel Method
which returns RelationshipDetail Obj which is then stored inside relation context here.
Parameters:
	- currentTableName: Name of Current Model
	- relations: String of all Relation reference of models.Relation
Returns:
	- nil nil
*/
func (relationshipHandler *RelationshipHandler) Handle(currentTableName string, relations interface{}) ([]api.EmitterFile,
	error) {

	modelGenerator := context.GetFromRegistry(context.Model).GetCtx(currentTableName)
	if modelGenerator != nil {
		relationshipModelGenerator := generator.NewRelationshipModel(modelGenerator.(*generator.ModelGenerator))
		myRelations := relations.(models.Relation)

		for _, rel := range myRelations.HasMany {
			var referenceTableName, foreignKey, primaryKey, err = getRelationshipKeys(rel, currentTableName)
			if err != nil {
				return nil, err
			}
			relationshipDetailObj, _ := relationshipModelGenerator.AddRelationshipToModel(helper.HasMany, currentTableName,
				referenceTableName, foreignKey, primaryKey)
			context.GetFromRegistry(context.Relation).AddToCtx(currentTableName, relationshipDetailObj)
		}

		for _, rel := range myRelations.HasOne {
			var referenceTableName, foreignKey, primaryKey, err = getRelationshipKeys(rel, currentTableName)
			if err != nil {
				return nil, err
			}
			relationshipDetailObj, _ := relationshipModelGenerator.AddRelationshipToModel(helper.HasOne, currentTableName,
				referenceTableName, foreignKey, primaryKey)
			context.GetFromRegistry(context.Relation).AddToCtx(currentTableName, relationshipDetailObj)
		}
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("model class %s not found", currentTableName))
}

/**
This method generates referenceTableName, foreignKey, primaryKey form the Relation string
as per laravel requirement.
*/
func getRelationshipKeys(relation string, currentTableName string) (string, string, string, error) {
	if relation != "" {
		var referenceTableName, foreignKey, primaryKey string
		splittedArray := strings.Split(relation, ":")
		if len(splittedArray) == 1 {
			referenceTableName = splittedArray[0]
			foreignKey = strings.ToLower(currentTableName) + "_id"
			primaryKey = "id"
		} else if len(splittedArray) == 2 {
			referenceTableName = splittedArray[0]
			pkFkSplitter := strings.Split(splittedArray[1], ",")
			if len(pkFkSplitter) == 2 {
				foreignKey = pkFkSplitter[0]
				primaryKey = pkFkSplitter[1]
			} else if len(pkFkSplitter) == 1 {
				foreignKey = pkFkSplitter[0]
				primaryKey = "id"
			}
		} else {
			return "", "", "", errors.New("rule Cannot be blank")
		}
		return referenceTableName, foreignKey, primaryKey, nil
	}
	return "", "", "", errors.New("relation String Cannot be Blank")
}
