package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"errors"
	"fmt"
)

const CreatedBy = "$table->%s('created_by')"
const UpdatedBy = "$table->%s('updated_by')->nullable()"
const Timestamp = `$table->timestamps();`
const SoftDeletes = `$table->softDeletes();`

type AuditCol struct {
	api.Handler
}

type AuditColInput struct {
	/**
	The column type specifies the action this handler performs
	For
	 	1 	- for CreatedBy,
		2 	- for UpdatedBy,
		3 	- for Both Created And UpdatedBy or formally AuditCols
		4 	- for Soft Deletes
		7 	- for CreatedBy, UpdatedBy, SoftDeletes
		8 	- for Timestamp
		15	- for CreatedBy, UpdatedBy, SoftDeletes, Timestamp
	*/
	columnType int
	value      bool
}

func NewAuditColInputFromType(auditCol bool, softDeletes bool, timestamp bool, value bool) *AuditColInput {
	return &AuditColInput{
		columnType: formatColumnType(auditCol, softDeletes, timestamp),
		value:      value,
	}
}

func NewAuditColHandler() *AuditCol {
	return &AuditCol{}
}

func (auditColHandler *AuditCol) Handle(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	requiresAuditCols := value.(bool)
	if requiresAuditCols {
		auditColHandler.handleFillable(identifier)
		auditColHandler.handleMigration(identifier)

	}
	return []*api.EmitterFile{}, nil
}

func (auditColHandler *AuditCol) handleFillable(identifier string) error {
	modelClass := context.GetFromRegistry("model").GetCtx(identifier).(*core.Class)
	if modelClass != nil {
		element, err := modelClass.FindInMembers("fillable")
		if err != nil {
			return err
		}
		// todo add validation rules in the CREATE_VALIDATION_RULES and UPDATE_VALIDATION_RULES array
		arrayAssignment := (*element).(*core.ArrayAssignment)
		arrayAssignment.Rhs = append(arrayAssignment.Rhs, "created_by", "updated_by")
		return nil
	}
	return errors.New(fmt.Sprintf("model class %s not found", identifier))
}

func (auditColHandler *AuditCol) handleMigration(identifier string) error {
	migrationInfo := context.GetFromRegistry("migration").GetCtx(identifier).(*context.MigrationInfo)
	if migrationInfo != nil {
		primaryKeyCol := getPrimaryKeyString(migrationInfo.PrimaryKeyCol)

		migrationClass := migrationInfo.Class
		function, err := migrationClass.FindInFunctions("up")
		if err != nil {
			return err
		}

		auditColHandler.appendToFunction(function, primaryKeyCol, 0)
		auditColHandler.appendToFunction(function, primaryKeyCol, 1)
		return nil

	}
	return errors.New(fmt.Sprintf("model class %s not found", identifier))

}

func (auditColHandler *AuditCol) appendToFunction(function *core.Function, primaryKeyCol string, auditColType int) {
	auditColumn := CreatedBy
	if auditColType == 1 {
		auditColumn = UpdatedBy
	}
	str := fmt.Sprintf(auditColumn, primaryKeyCol)
	_, err := function.FindStatement(str)
	if err != nil {
		simple := core.NewSimpleStatement(str)
		stmt := core.TabbedUnit(simple)
		function.Statements = append(function.Statements, &stmt)
	}
}

func getPrimaryKeyString(primaryKeyCol string) string {
	switch primaryKeyCol {
	case "unsignedBigInteger":
		return "unsignedBigInteger"
	case "bigInteger":
		return "bigInteger"

	}
	return "unsignedBigInteger"
}

func formatColumnType(auditCol bool, softDeletes bool, timestamp bool) int {
	colType := 0
	if auditCol {
		colType |= 3
	}
	if softDeletes {
		colType |= 4
	}
	if timestamp {
		colType |= 8
	}
	return colType
}
