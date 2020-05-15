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
	identifier string	// name of the model
}

func NewAuditColInputFromType(auditCol bool, softDeletes bool, timestamp bool, id string) *AuditColInput {
	return &AuditColInput{
		identifier: id,
		columnType: formatColumnType(auditCol, softDeletes, timestamp),
	}
}

func (input *AuditColInput) isAuditColSet() bool {
	return input.columnType & 3 == 3
}

func (input *AuditColInput) isSoftDeletesSet() bool {
	return input.columnType & 4 == 4
}

func (input *AuditColInput) isTimestampSet() bool {
	return input.columnType & 8 == 8
}

/**
Returns a slice containing a list of columns to be appended to the fillable array
 */
func (input *AuditColInput) getFillableArray() []string {
	var arr []string
	if  input.isAuditColSet() {
		arr = append(arr, `"created_by"`, `"updated_by"`)
	}
	if input.isTimestampSet() {
		arr = append(arr, `"created_at"`, `"updated_at"`)
	}
	if input.isSoftDeletesSet() {
		arr = append(arr, `"deleted_at"`)
	}
	return arr
}


func NewAuditColHandler() *AuditCol {
	return &AuditCol{}
}

func (auditColHandler *AuditCol) Handle(identifier string, value interface{}) ([]*api.EmitterFile, error) {
	input := value.(*AuditColInput)
	// todo handle errors
	auditColHandler.handleFillable(input)
	return []*api.EmitterFile{}, nil
}


func (auditColHandler *AuditCol) handleFillable(input *AuditColInput) error {
	modelClass := context.GetFromRegistry("model").GetCtx(input.identifier).(*core.Class)
	if modelClass != nil {
		element, err := modelClass.FindInMembers("fillable")
		if err != nil {
			return err
		}
		// todo add validation rules in the CREATE_VALIDATION_RULES and UPDATE_VALIDATION_RULES array
		arrayAssignment := (*element).(*core.ArrayAssignment)
		arrayAssignment.Rhs = append(arrayAssignment.Rhs, input.getFillableArray()...)
		return nil
	}
	return errors.New(fmt.Sprintf("model class %s not found", input.identifier))
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
/**
Returns a bit mask of the booleans provided
auditCols has a value of 3
softDeletes - 4
timestamp - 8
if all are set then the integer would represent 15
 */
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
