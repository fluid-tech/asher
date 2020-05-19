package handler
//
//import (
//	"asher/internal/api"
//	"asher/internal/api/codebuilder/php/core"
//	"asher/internal/impl/laravel/5.8/handler/context"
//	"asher/internal/impl/laravel/5.8/handler/helper"
//	"errors"
//	"fmt"
//)
//
//const CreatedBy = "$table->%s('created_by')"
//const UpdatedBy = "$table->%s('updated_by')->nullable()"
//const Timestamp = `$table->timestamps()`
//const SoftDeletes = `$table->softDeletes()`
//
//const FillableIdentifier = "fillable"
//const CreateValidationRulesIdentifier = "getCreateValidationRules"
//const UpdateValidationRulesIdentifier = "getUpdateValidationRules"
//const UserModelIdentifier = "user"
//
//type AuditCol struct {
//	api.Handler
//}
//
//func NewAuditColHandler() *AuditCol {
//	return &AuditCol{}
//}
//
//func (auditColHandler *AuditCol) Handle(identifier string, value interface{}) ([]*api.EmitterFile, error) {
//	input := value.(*helper.AuditColInput)
//	// todo handle errors
//	auditColHandler.handleModel(identifier, input)
//	auditColHandler.handleMigration(identifier, input)
//	return []*api.EmitterFile{}, nil
//}
//
///**
//Function orchestrates methods that adds data to the model class
//*/
//func (auditColHandler *AuditCol) handleModel(identifier string, input *helper.AuditColInput) error {
//	modelClass := context.GetFromRegistry("model").GetCtx(identifier).(*core.Class)
//	if modelClass != nil {
//		auditColHandler.handleTimestamp(modelClass, input.IsTimestampSet())
//		auditColHandler.handleSoftDeletes(modelClass, input.IsSoftDeletesSet())
//		auditColHandler.handleFillable(modelClass, input.GetFillableArray())
//		auditColHandler.handleValidationRules(UpdateValidationRulesIdentifier, modelClass, input.GetUpdateValidationRules())
//		auditColHandler.handleValidationRules(CreateValidationRulesIdentifier, modelClass, input.GetCreateValidationRules())
//		return nil
//	}
//	return errors.New(fmt.Sprintf("model class %s not found", identifier))
//}
//
//func (auditColHandler *AuditCol) handleTimestamp(modelClass *core.Class, isTimeStampSet bool) {
//	// adding timestamps true
//	if isTimeStampSet {
//		tab := api.TabbedUnit(core.NewVarAssignment("public", "timestamps", "true"))
//		modelClass.AppendMember(&tab)
//	}
//}
//
//func (auditColHandler *AuditCol) handleSoftDeletes(modelClass *core.Class, isSoftDeleteSet bool) {
//	// adding use SoftDeletes
//	if isSoftDeleteSet {
//		tab := api.TabbedUnit(core.NewSimpleStatement("use SoftDeletes"))
//		modelClass.AppendMember(&tab)
//	}
//}
//
//func (auditColHandler *AuditCol) handleFillable(modelClass *core.Class, fillableArray []string) error {
//	element, err := modelClass.FindMember(FillableIdentifier)
//	if err != nil {
//		return err
//	}
//	arrayAssignment := (*element).(*core.ArrayAssignment)
//	arrayAssignment.Rhs = append(arrayAssignment.Rhs, fillableArray...)
//	return nil
//}
//
//func (auditColHandler *AuditCol) handleValidationRules(currentIdentifier string, klass *core.Class, arrayToAppend []string) error {
//	// adding validation rules for audit cols
//
//	function, err := klass.FindFunction(currentIdentifier)
//	if err != nil {
//		return err
//	}
//	// this is an assumption that this method returns an array in the first line itself
//	// ie the first element is a ReturnArray
//	returnStmt, err := function.FindById("return")
//	if err != nil {
//		return err
//	}
//
//	returnArray := (*returnStmt).(*core.ReturnArray)
//	returnArray.Append(arrayToAppend)
//	return nil
//}
//
//func (auditColHandler *AuditCol) handleMigration(identifier string, input *helper.AuditColInput) error {
//	migrationCtx := context.GetFromRegistry("migration")
//	migrationInfo := migrationCtx.GetCtx(identifier).(*context.MigrationInfo)	if migrationInfo != nil {
//
//		primaryKeyCol := getLaravelColString(migrationCtx.GetCtx("user").(*context.MigrationInfo))
//
//		function, err := getSchemaCreateMethod(migrationInfo)
//		if err != nil {
//			return err
//		}
//		if input.IsAuditColSet() {
//			auditColHandler.appendToFunction(function, fmt.Sprintf(CreatedBy, primaryKeyCol))
//			auditColHandler.appendToFunction(function, fmt.Sprintf(UpdatedBy, primaryKeyCol))
//		}
//		if input.IsSoftDeletesSet() {
//			auditColHandler.appendToFunction(function, SoftDeletes)
//		}
//		if input.IsTimestampSet() {
//			auditColHandler.appendToFunction(function, Timestamp)
//		}
//
//	}
//	return errors.New(fmt.Sprintf("model class %s not found", identifier))
//
//}
//
//func (auditColHandler *AuditCol) appendToFunction(function *core.Function, auditColStr string) {
//	_, err := function.FindStatement(auditColStr)
//	if err != nil {
//		// no record of statement found inserting
//		simple := core.NewSimpleStatement(auditColStr)
//		stmt := api.TabbedUnit(simple)
//		function.AppendStatement(&stmt)
//	}
//}
//
//func getLaravelColString(userMigration *context.MigrationInfo) string {
//	primaryKeyStr := ""
//	if userMigration != nil && len(userMigration.PrimaryKeyCol) > 0 {
//		// using the first argument as the type of primaryKeyCol
//		primaryKeyStr = userMigration.PrimaryKeyCol[0]
//	}
//	return keyColToLaravelString(primaryKeyStr)
//
//}
//
//func getSchemaCreateMethod(migrationInfo *context.MigrationInfo) (*core.Function, error) {
//	migrationClass := migrationInfo.Class
//	function, err := migrationClass.Build().FindFunction("up")
//	if err != nil {
//		return nil, err
//	}
//
//	// finding nested function Call schema::create
//	upFunc, err := function.FindById("Schema::create")
//	if err != nil {
//		return nil, err
//	}
//	upFun, err := (*upFunc).(*core.FunctionCall).FindById("anon")
//	if err != nil {
//		return nil, err
//	}
//	return (*upFun).(*core.Function), nil
//}
//
//func keyColToLaravelString(primaryKeyCol string) string {
//	// Todo add UUID
//	switch primaryKeyCol {
//	case "unsignedBigInteger":
//		return "unsignedBigInteger"
//	case "bigInteger":
//		return "bigInteger"
//	}
//	return "unsignedBigInteger"
//}
