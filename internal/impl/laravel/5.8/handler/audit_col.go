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

func (auditColHandler *AuditCol) Handle(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	input := value.(*helper.AuditColInput)
	// todo handle errors
	err := auditColHandler.handleModel(identifier, input)
	if err != nil{
		return nil, err
	}
	err = auditColHandler.handleMigration(identifier, input)
	if err != nil{
		return nil, err
	}
	return []*api.EmitterFile{}, nil
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
	modelGenerator := context.GetFromRegistry("model").GetCtx(identifier).(*generator.ModelGenerator)
	if modelGenerator != nil {
		modelGenerator.SetTimestamps(input.IsTimestampSet())
		modelGenerator.SetSoftDeletes(input.IsSoftDeletesSet())
		modelGenerator.SetAuditCols(input.IsAuditColSet())
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
	migGen := context.GetFromRegistry("migration").GetCtx(identifier).(*generator.MigrationGenerator)
	if migGen != nil {
		migGen.SetAuditCols(input.IsAuditColSet())
		migGen.SetTimestamps(input.IsTimestampSet())
		migGen.SetSoftDeletes(input.IsSoftDeletesSet())
		migGen.SetPkCol(input.PkColVal)
	}
	return errors.New(fmt.Sprintf("migration class %s not found", identifier))

}
