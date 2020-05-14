package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler/context"
	"errors"
	"fmt"
)

type AuditCol struct {
	api.Handler
}

const CreatedBy = "$table->%s('created_by')->nullable()"
const UpdatedBy = "$table->%s('updated_by')->nullable()"

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
	if modelClass != nil{
		element, err := modelClass.FindInMembers("fillable")
		if err != nil {
			// todo return an err
			return err
		}
		arrayAssignment := (*element).(*core.ArrayAssignment)
		arrayAssignment.Rhs = append(arrayAssignment.Rhs, "created_by", "updated_by")
		return nil
	}
	return errors.New(fmt.Sprintf("model class %s not found", identifier))
}

func (auditColHandler *AuditCol) handleMigration(identifier string) error {
	migrationInfo := context.GetFromRegistry("migration").GetCtx(identifier).(*context.MigrationInfo)
	if migrationInfo != nil{
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
		simple := core.GetSimpleStatement(str)
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
