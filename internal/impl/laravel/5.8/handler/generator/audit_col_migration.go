package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/helper"
)

type AuditColMigration struct {
	migrationGen *MigrationGenerator
	pkColVal     string
}



func NewAuditColGenerator(generator *MigrationGenerator) *AuditColMigration {
	return &AuditColMigration{
		migrationGen: generator,
		pkColVal:     DefaultColVal,
	}
}

/**
 Sets the AuditCols field of this generator. During build adds string such as
 `$table->unsignedBigInteger('created_by');` and `$table->unsignedBigInteger('updated_by');`
 to the migration. The method name changes depending upon the col type used in the users table.
 NOTE- If this is set and SetPkCol is not used `unsignedBigInteger` is used as the default.
 Parameters
	- auditCols:	bool 		If set this generator adds created by and updated by cols to the migration file
 Returns:
	- instance of the generator object
 Example:
	- builder.SetAuditCols(true)
*/
func (auditColGen *AuditColMigration) SetAuditCols(auditCols bool) *AuditColMigration {
	if auditCols {
		cbstr := helper.ColTypeSwitcher(auditColGen.pkColVal, CreatedByStr, []string{})
		upstr := helper.ColTypeSwitcher(auditColGen.pkColVal, UpdatedByStr, []string{}) + "->nullable()"
		auditColGen.migrationGen.AddColumns([]*core.SimpleStatement{
			core.NewSimpleStatement(cbstr), core.NewSimpleStatement(upstr),
		})
	}
	return auditColGen
}

/**
 Sets the pkCol field of this generator. During build adds string such as
 `$table->unsignedBigInteger('created_by');` and `$table->unsignedBigInteger('updated_by');`
 NOTE if this is not set the default value (`unsignedBigInteger`) is used.
 Parameters
	- pkColType:	string		The primary key col type of users table
 Returns:
	- instance of the generator object
 Example:
	- builder.SetPkCol(true)
*/
func (auditColGen *AuditColMigration) SetPkCol(pkCol string) *AuditColMigration {
	auditColGen.pkColVal = pkCol
	return auditColGen
}

/**
 Sets the timestamps field of this generator. During build adds the string `$table->softDeletes();`
 to the migration.
 Returns:
	- instance of the generator object
 Example:
	- builder.SetSoftDeletes(true)
*/
func (auditColGen *AuditColMigration) SetSoftDeletes(softDeletes bool) *AuditColMigration {
	if softDeletes {
		auditColGen.migrationGen.AddColumn(*core.NewSimpleStatement(SoftDeletesCol))
	}
	return auditColGen
}

/**
 Sets the timestamps field of this generator. During build adds the string `$table->timestamps();`
 to the migration.
 Returns:
	- instance of the generator object
 Example:
	- builder.SetTimestamps(true)
*/
func (auditColGen *AuditColMigration) SetTimestamps(timestamps bool) *AuditColMigration {
	if timestamps {
		auditColGen.migrationGen.AddColumns([]*core.SimpleStatement{
			core.NewSimpleStatement(TimestampCol),
		})
	}
	return auditColGen
}