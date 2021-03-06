package handler

import (
	"asher/internal/api"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/generator"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"errors"
	"fmt"
)

const CreatedBy = "$table->%s('created_by')"
const UpdatedBy = "$table->%s('updated_by')->nullable()"

const FillableIdentifier = "fillable"
const CreateValidationRulesIdentifier = "getCreateValidationRules"
const UpdateValidationRulesIdentifier = "getUpdateValidationRules"
const UserModelIdentifier = "user"

type AuditCol struct {
	api.Handler
}

func NewAuditColHandler() *AuditCol {
	return &AuditCol{}
}

func (auditColHandler *AuditCol) Handle(identifier string, value interface{}) ([]api.EmitterFile, error) {
	input := value.(*helper.AuditColInput)
	// todo handle errors
	err := auditColHandler.handleModel(identifier, input)
	if err != nil {
		return nil, err
	}
	err = auditColHandler.handleMigration(identifier, input)
	if err != nil {
		return nil, err
	}
	return []api.EmitterFile{}, nil
}

/**
 Orchestrates methods that adds data to the model class
 Parameters
	- identifier:	string 					The model this handler is working on
	- input:		*helperAuditColInput 	The input given to this handler
 Returns error if the model with the provided identifier could not be found
 in the modelContext
*/
func (auditColHandler *AuditCol) handleModel(identifier string, input *helper.AuditColInput) error {
	retrievedModelGenerator := context.GetFromRegistry(context.Model).GetCtx(identifier)
	if retrievedModelGenerator != nil {
		auditColModel := generator.NewAuditColModel(retrievedModelGenerator.(*generator.ModelGenerator))
		auditColModel.SetTimestamps(input.IsTimestampSet())
		auditColModel.SetAuditCol(input.IsAuditColSet())
		auditColModel.SetSoftDeletes(input.IsSoftDeletesSet())
		return nil
	}
	return errors.New(fmt.Sprintf("model class %s not found", identifier))
}

/**
 Orchestrates methods that adds columns to the migration class
 Parameters
	- identifier:	string 					The model this handler is working on
	- input:		*helperAuditColInput 	The input given to this handler
 Returns an error if a migration with the provided identifier could not be found
 in the modelContext
*/
func (auditColHandler *AuditCol) handleMigration(identifier string, input *helper.AuditColInput) error {
	migGen := context.GetFromRegistry(context.Migration).GetCtx(identifier)
	if migGen != nil {
		auditColMigGen := generator.NewAuditColMigration(migGen.(*generator.MigrationGenerator))
		auditColMigGen.SetTimestamps(input.IsTimestampSet())
		auditColMigGen.SetPkCol(input.PkColVal)
		auditColMigGen.SetAuditCols(input.IsAuditColSet())
		auditColMigGen.SetSoftDeletes(input.IsSoftDeletesSet())
		return nil
	}
	return errors.New(fmt.Sprintf("migration class %s not found", identifier))

}
