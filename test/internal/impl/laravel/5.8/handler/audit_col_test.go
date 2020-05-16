package handler

import (
	"asher/internal/api"
	"asher/internal/api/codebuilder/php/builder"
	"asher/internal/api/codebuilder/php/core"
	"asher/internal/impl/laravel/5.8/handler"
	"asher/internal/impl/laravel/5.8/handler/context"
	"asher/internal/impl/laravel/5.8/handler/helper"
	"fmt"
	"reflect"
	"testing"
)

func TestModel(t *testing.T) {
	var table = []*struct {
		in  [][]string
		out [][]string
	}{
		{genModelTest("Hello", true, true, true, t), [][]string{arr, updateTest1, createTest1}},
		{genModelTest("Rnadom", true, true, false, t), [][]string{brr, updateTest1, createTest1}},
		{genModelTest("Random", true, false, false, t), [][]string{crr, updateTest3, createTest1}},
		{genModelTest("Hell4", false, true, true, t), [][]string{drr, updateTest4, createTest4}},
		{genModelTest("Hell1", false, false, false, t), [][]string{frr, createTest4, createTest4}},
	}

	for i, element := range table {
		for j, inner := range element.in {
			if !reflect.DeepEqual(inner, element.out[j]) {
				t.Errorf("in test case %d , %d expected '%s' found '%s'", i, j, inner, element.out[j])
			}
		}

	}
}

func genModelTest(className string, auditCol bool, softDeletes bool, timestamp bool, t *testing.T) [][]string {
	klass := buildClassWithArrayDecl(className)
	// adding klass to model registry
	modelRegistry := context.GetFromRegistry("model")
	modelRegistry.AddToCtx(className, klass)

	// calling handle for fillable
	auditColHandler := handler.NewAuditColHandler()
	auditColHandler.Handle(className, helper.NewAuditColInputFromType(auditCol, softDeletes, timestamp))

	if softDeletes {
		_, err := klass.FindMember("use SoftDeletes")
		if err != nil {
			t.Error("Couldnt find member use SoftDeletes")
		}
	}

	if timestamp {
		_, err := klass.FindMember("timestamps")
		if err != nil {
			t.Error("Couldnt find member timestamps")
		}
	}

	return [][]string{
		getFillableRhs(klass, t),
		getReturnArrayStatementsForFunction(handler.UpdateValidationRulesIdentifier, klass, t),
		getReturnArrayStatementsForFunction(handler.CreateValidationRulesIdentifier, klass, t),
	}

}

func getReturnArrayStatementsForFunction(funcName string, klass *core.Class, t *testing.T) []string {
	updateValRulesFunc, err := klass.FindFunction(funcName)
	if err != nil {
		t.Error("couldnt find function getUpdateValidationRules")
	}
	updateRetStmt, err := updateValRulesFunc.FindById("return")
	if err != nil {
		t.Error("Couldnt find return in updateValidationRules")
	}
	return (*updateRetStmt).(*core.ReturnArray).Statements
}

func getFillableRhs(klass *core.Class, t *testing.T) []string {
	element, err := klass.FindMember("fillable")
	if err != nil {
		t.Error("fillable not found in klass")
	}
	return (*element).(*core.ArrayAssignment).Rhs
}

func buildClassWithArrayDecl(className string) *core.Class {
	assigmentSS := api.TabbedUnit(core.NewSimpleStatement("$this->fullyQualifiedModel = $fullyQualifiedModel"))
	functionBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("__construct").
		AddArgument("string $fullyQualifiedModel").
		AddStatement(&assigmentSS)

	returnArray := api.TabbedUnit(core.NewReturnArray([]string{`"user_id" => "required|exists:users,id"`, `"password" => "min:8|max20"`}))
	returnArray2 := api.TabbedUnit(core.NewReturnArray([]string{`"user_id" => "required|exists:users,id"`, `"password" => "min:8|max20"`}))
	getCreateRules := builder.NewFunctionBuilder().SetVisibility("public").SetName("getCreateValidationRules").
		AddStatement(&returnArray)

	getUpdateRules := builder.NewFunctionBuilder().SetVisibility("public").SetName("getUpdateValidationRules").
		AddStatement(&returnArray2)

	member := api.TabbedUnit(core.GetVarDeclaration("private", "fullyQualifiedModel"))
	rhs := []string{`"user_id"`, `"password"`}
	arrayAssignmentMember := api.TabbedUnit(core.NewArrayAssignment("public", "fillable", rhs))
	klass := builder.NewClassBuilder().SetName(className).SetExtends("BaseMutator").
		AddFunction(functionBuilder.GetFunction()).AddMember(&member).AddMember(&arrayAssignmentMember).
		SetPackage("App").AddImport(`Illuminate\Database\Eloquent\Model`).
		AddFunction(getCreateRules.GetFunction()).AddFunction(getUpdateRules.GetFunction())

	return klass.GetClass()
}

func TestMigration(t *testing.T) {

	var table = []struct {
		in  []bool
		out []bool
	}{
		{genMigTest("TTT", true, true, true, t), []bool{true, true, true, true}},
		{genMigTest("TT", true, false, true, t), []bool{true, true, false, true}},
		{genMigTest("T", false, true, true, t), []bool{false, false, true, true}},
		{genMigTest("TTTA", true, true, false, t), []bool{true, true, true, false}},
	}

	for i, element := range table {
		if !reflect.DeepEqual(element.in, element.out) {
			t.Errorf("failed %d expected %t found %t", i, element.in, element.out)
		}
	}

}

func genMigTest(className string, auditCol bool, softDeletes bool, timestamp bool, t *testing.T) []bool {
	klass := genMigrationClass(className)
	context.GetFromRegistry("migration").AddToCtx(className, klass)

	// calling handle for fillable
	auditColHandler := handler.NewAuditColHandler()
	auditColHandler.Handle(className, helper.NewAuditColInputFromType(auditCol, softDeletes, timestamp))

	function, err := klass.FindFunction("up")
	if err != nil {
		t.Error("cant find method up")
	}
	upFunc, err := function.FindById("Schema::create")
	if err != nil {
		t.Error("cant find method call Schema::create")
	}

	up := (*upFunc).(*core.FunctionCall)
	fun, err := up.FindById("anon")
	if err != nil {
		t.Error("cant find anon")
	}
	anonFunction := (*fun).(*core.Function)
	createdBy := fmt.Sprintf(handler.CreatedBy, "unsignedBigInteger")
	updatedBy := fmt.Sprintf(handler.UpdatedBy, "unsignedBigInteger")

	return []bool{
		runTestFor(auditCol, createdBy, anonFunction, t),
		runTestFor(auditCol, updatedBy, anonFunction, t),
		runTestFor(softDeletes, handler.SoftDeletes, anonFunction, t),
		runTestFor(timestamp, handler.Timestamp, anonFunction, t),
	}

}

func runTestFor(condition bool, stmt string, anon *core.Function, t *testing.T) bool {
	if condition {
		_, err := anon.FindById(stmt)
		if err != nil {
			t.Errorf("%s failed", stmt)
			return false
		}
		return true
	}
	return false
}

/**
Returns a half baked migration class with up method only
*/
func genMigrationClass(id string) *core.Class {

	stmt := api.TabbedUnit(core.NewSimpleStatement("$table->bigIncrements('id')"))
	anon := api.TabbedUnit(builder.NewFunctionBuilder().AddArgument("Blueprint $table").
		AddStatement(&stmt).GetFunction())
	funcCall := api.TabbedUnit(core.NewFunctionCall("Schema::create").AddArg(&anon))

	funcBuilder := builder.NewFunctionBuilder().SetVisibility("public").SetName("up").
		AddStatement(&funcCall)

	return builder.NewClassBuilder().SetName(id).AddFunction(funcBuilder.GetFunction()).GetClass()

}

/*** CONSTANT STRINGS/ ARRAYS ***/

var arr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`, `"deleted_at"`}
var brr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`, `"deleted_at"`}
var crr = []string{`"user_id"`, `"password"`, `"created_by"`, `"updated_by"`}
var drr = []string{`"user_id"`, `"password"`, `"deleted_at"`}
var frr = []string{`"user_id"`, `"password"`}

var updateTest1 = []string{
	`"user_id" => "required|exists:users,id"`,
	`"password" => "min:8|max20"`,
	`"updated_by" => "required|exists:users,id"`,
	`"deleted_at" => "required|date_format:Y-m-d H:i:s"`,
}
var createTest1 = []string{
	`"user_id" => "required|exists:users,id"`,
	`"password" => "min:8|max20"`,
	`"created_by" => "required|exists:users,id"`,
}

var updateTest3 = []string{
	`"user_id" => "required|exists:users,id"`,
	`"password" => "min:8|max20"`,
	`"updated_by" => "required|exists:users,id"`,
}

var updateTest4 = []string{
	`"user_id" => "required|exists:users,id"`,
	`"password" => "min:8|max20"`,
	`"deleted_at" => "required|date_format:Y-m-d H:i:s"`,
}
var createTest4 = []string{
	`"user_id" => "required|exists:users,id"`,
	`"password" => "min:8|max20"`,
}
