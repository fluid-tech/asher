package generator

import "asher/internal/api/codebuilder/php/core"

type AuditColModel struct {
	modelGen *ModelGenerator
}

func NewAuditColModel(generator *ModelGenerator) *AuditColModel {
	return &AuditColModel{
		modelGen: generator,
	}
}

/**
 During build adds `public $timestamps = true` to the model if true.
 Parameters
	- auditCols:	  timestamp 		If set this generator adds created by and updated by cols to the migration file
 Returns:
	- instance of the generator object
 Example:
	- builder.SetTimestamps(true)
*/
func (auditColModel *AuditColModel) SetTimestamps(timestamp bool) *AuditColModel {
	if timestamp {
		auditColModel.modelGen.classBuilder.AddMember(core.NewSimpleStatement(DefaultTimestampStr))
	}
	return auditColModel
}

/**
 During build adds `Use SoftDeletes` to the model if true
 Parameters:
	- softDeletes:		bool		If set adds `Use SoftDeletes` to this model
 Returns:
	- instance of this generator
 Example:
	- builder.SetSoftDeletes(true)
*/
func (auditColModel *AuditColModel) SetSoftDeletes(softDeletes bool) *AuditColModel {
	if softDeletes {
		auditColModel.modelGen.classBuilder.AddMember(core.NewSimpleStatement(UseSoftDeletesStr))
		// todo make date format configurable
		auditColModel.modelGen.AddUpdateValidationRule(DeletedAtStr, DeletedAtValidationRule)
		auditColModel.modelGen.AddFillable(DeletedAtStr)
	}
	return auditColModel
}

/**
 Adds the AuditCols field of to the fillable array of this model and adds them to create and update
 validation rules array.
 Parameters
	- auditCols:		bool 		If set this generator adds created by and updated by cols to the migration file
 Returns:
	- instance of the generator object
 Example:
	- builder.SetAuditCols(true)
*/
func (auditColModel *AuditColModel) SetAuditCol(auditCol bool) *AuditColModel {
	if auditCol {
		auditColModel.modelGen.AddCreateValidationRule(CreatedByStr, DefaultAuditColValidation)
		auditColModel.modelGen.AddUpdateValidationRule(UpdatedByStr, DefaultAuditColValidation)
		auditColModel.modelGen.AddFillable(CreatedByStr).AddFillable(UpdatedByStr)
	}
	return auditColModel
}
